package pipeline

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/infraboard/keyauth/pkg/token"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/types/ftime"
	"github.com/rs/xid"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func NewCreatePipelineRequest() *CreatePipelineRequest {
	return &CreatePipelineRequest{}
}

// NewQueryPipelineRequest 查询book列表
func NewQueryPipelineRequest() *QueryPipelineRequest {
	return &QueryPipelineRequest{
		Page: &request.NewDefaultPageRequest().PageRequest,
	}
}

func (req *CreatePipelineRequest) Validate() error {
	if len(req.Stages) == 0 {
		return fmt.Errorf("no stages")
	}
	return validate.Struct(req)
}

func (req *CreatePipelineRequest) EnsureStep() {
	for m := range req.Stages {
		s := req.Stages[m]
		s.Id = int32(m) + 1
		for n := range s.Steps {
			t := s.Steps[n]
			t.Id = int32(n) + 1
			t.Status = NewDefaultStepStatus()
		}
	}
}

func LoadPipelineFromBytes(payload []byte) (*Pipeline, error) {
	ins := NewDefaultPipeline()

	// 解析Value
	if err := json.Unmarshal(payload, ins); err != nil {
		return nil, fmt.Errorf("unmarshal step error, vaule(%s) %s", string(payload), err)
	}

	// 校验合法性
	if err := ins.Validate(); err != nil {
		return nil, err
	}

	return ins, nil
}

func NewDefaultPipeline() *Pipeline {
	return &Pipeline{
		Status: &PipelineStatus{},
	}
}

func NewDefaultPipelineStatus() *PipelineStatus {
	return &PipelineStatus{}
}

func NewPipeline(req *CreatePipelineRequest) (*Pipeline, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 补充step id
	req.EnsureStep()

	p := &Pipeline{
		Id:          xid.New().String(),
		CreateAt:    ftime.Now().Timestamp(),
		Name:        req.Name,
		With:        req.With,
		Mount:       req.Mount,
		Tags:        req.Tags,
		Description: req.Description,
		On:          req.On,
		Stages:      req.Stages,
		Status:      &PipelineStatus{},
	}
	return p, nil
}

// MakeObjectKey 构建etcd对应的key
// 例如: inforboard/workflow/service/node/node-01
func (p *Pipeline) MakeObjectKey() string {
	return fmt.Sprintf("%s/%s/%s", EtcdPipelinePrefix(), p.Namespace, p.Id)
}

func (p *Pipeline) UpdateOwner(tk *token.Token) {
	p.CreateBy = tk.Account
	p.Domain = tk.Domain
	p.Namespace = tk.Namespace
}

func (p *Pipeline) Validate() error {
	return validate.Struct(p)
}

func (p *Pipeline) UpdateStep(s *Step) error {
	ns, id := s.GetNamespace(), s.GetPipelineID()
	if ns != p.Namespace || id != p.Id {
		return fmt.Errorf("this step not match this pipeline, id or namespace is not correct")
	}

	stage, err := p.GetStageByNumber(s.GetPipelineStageNumber())
	if err != nil {
		return fmt.Errorf("get stage error, %s", err)
	}

	return stage.UpdateStep(s)
}

func (p *Pipeline) AddStage(item *Stage) {
	p.Stages = append(p.Stages, item)
}

func (p *Pipeline) GetStageByNumber(number int32) (*Stage, error) {
	if int(number) > len(p.Stages) || number <= 0 {
		return nil, fmt.Errorf("number range 1 ~ %d", len(p.Stages))
	}

	return p.Stages[number-1], nil
}

func (p *Pipeline) LastStage() *Stage {
	total := len(p.Stages)
	if total == 0 {
		return nil
	}

	return p.Stages[total-1]
}

func (p *Pipeline) IsComplete() bool {
	return p.Status.Status.Equal(PIPELINE_STATUS_COMPLETE)
}

func (s *Pipeline) IsScheduled() bool {
	if s.Status == nil {
		return false
	}

	return s.Status.SchedulerNode != ""
}

func (p *Pipeline) IsRunning() bool {
	if p.Status == nil {
		return false
	}
	return p.Status.Status.Equal(PIPELINE_STATUS_EXECUTING)
}

func (p *Pipeline) NextFlowNumber() int64 {
	if p.Status == nil {
		return 1
	}

	return p.Status.GetCurrentFlow() + 1
}

func (p *Pipeline) CurrentFlowNumber() int64 {
	if p.Status == nil {
		return 0
	}

	return p.Status.GetCurrentFlow()
}

func (p *Pipeline) incFlow() {
	p.Status.CurrentFlow++
}

func (p *Pipeline) Run() {
	if p.Status == nil {
		p.Status = NewDefaultPipelineStatus()
	}
	p.Status.Status = PIPELINE_STATUS_EXECUTING
	p.Status.StartAt = ftime.Now().Timestamp()
}

func (p *Pipeline) Complete() {
	if p.Status == nil {
		p.Status = NewDefaultPipelineStatus()
	}
	p.Status.Status = PIPELINE_STATUS_COMPLETE
	p.Status.EndAt = ftime.Now().Timestamp()
}

func (p *Pipeline) GetCurrentFlow() *Flow {
	flow := p.CurrentFlowNumber()

	// pipeline还没有运行
	if flow == 0 {
		return nil
	}

	for i := range p.Stages {
		stage := p.Stages[i]
		if f := stage.GetFlow(flow); f != nil {
			return f
		}
	}

	return nil
}

func (p *Pipeline) GetNextFlow() *Flow {
	for i := range p.Stages {
		stage := p.Stages[i]

		// stage 通过了继续搜索下一个stage
		if stage.IsPassed() {
			continue
		}

		// 如果stage中断后 就没有下一步了
		if stage.IsBreakNow() {
			return nil
		}

		steps := stage.NextStep()
		for i := range steps {
			step := steps[i]
			steps[i].BuildKey(p.Namespace, p.Id, stage.Id)
			step.setFlowNumber(p.NextFlowNumber())
		}

		if len(steps) > 0 {
			p.incFlow()
		}
		return NewFlow(p.CurrentFlowNumber(), steps)
	}
	return nil
}

// 只有上一个flow执行完成后, 才会有下个fow
// 注意: 多个并行的任务是不能跨stage同时执行的
//      也就是说stage一定是串行的
func (p *Pipeline) NextStep() (steps []*Step, isComplete bool) {
	if f := p.GetCurrentFlow(); f != nil {
		// 如果有flow中断, pipeline 提前结束
		if f.IsBreak() {
			isComplete = true
			return
		}

		// 如果flow没有pass 说明还是在运行中, 不需要调度下以组step
		if !f.IsPassed() {
			return
		}
	}

	f := p.GetNextFlow()
	if f == nil {
		p.Complete()
		isComplete = true
		return
	}

	

	steps = f.items
	return
}

func (t *Pipeline) ShortDescribe() string {
	return fmt.Sprintf("%s[%s]", t.Name, t.Id)
}

func (t *Pipeline) ScheduledNodeName() string {
	return t.Status.SchedulerNode
}

func (t *Pipeline) SetScheduleNode(nodeName string) {
	t.Status.SchedulerNode = nodeName
}

func (p *Pipeline) EtcdObjectKey() string {
	return fmt.Sprintf("%s/%s/%s", EtcdPipelinePrefix(), p.Namespace, p.Id)
}

func (s *Pipeline) MatchScheduler(schedulerName string) bool {
	if s.Status == nil {
		return false
	}
	return s.Status.SchedulerNode == schedulerName
}

// NewPipelineSet todo
func NewPipelineSet() *PipelineSet {
	return &PipelineSet{
		Items: []*Pipeline{},
	}
}

func (s *PipelineSet) Add(item *Pipeline) {
	s.Items = append(s.Items, item)
}

// NewQueryPipelineRequest 查询book列表
func NewDescribePipelineRequestWithID(id string) *DescribePipelineRequest {
	return &DescribePipelineRequest{
		Id: id,
	}
}

// NewDeletePipelineRequestWithID 查询book列表
func NewDeletePipelineRequestWithID(id string) *DeletePipelineRequest {
	return &DeletePipelineRequest{
		Id: id,
	}
}

func NewQueryPipelineOptions() *QueryPipelineOptions {
	return &QueryPipelineOptions{}
}

// QueryPipelineTaskOptions ListPipeline 查询条件
type QueryPipelineOptions struct {
	Node string
}
