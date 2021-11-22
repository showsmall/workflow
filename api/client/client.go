package client

import (
	kc "github.com/infraboard/keyauth/client"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
	"google.golang.org/grpc"

	"github.com/infraboard/workflow/api/app/action"
	"github.com/infraboard/workflow/api/app/application"
	"github.com/infraboard/workflow/api/app/deploy"
	"github.com/infraboard/workflow/api/app/pipeline"
	"github.com/infraboard/workflow/api/app/template"
)

var (
	client *Client
)

// SetGlobal todo
func SetGlobal(cli *Client) {
	client = cli
}

// C Global
func C() *Client {
	return client
}

// NewClient todo
func NewClient(conf *kc.Config) (*Client, error) {
	zap.DevelopmentSetup()
	log := zap.L()

	conn, err := grpc.Dial(conf.Address(), grpc.WithInsecure(), grpc.WithPerRPCCredentials(conf.Authentication))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		log:  log,
	}, nil
}

// Client 客户端
type Client struct {
	conn *grpc.ClientConn
	log  logger.Logger
}

// Example todo
func (c *Client) Application() application.ServiceClient {
	return application.NewServiceClient(c.conn)
}

// Example todo
func (c *Client) Pipeline() pipeline.ServiceClient {
	return pipeline.NewServiceClient(c.conn)
}

// Example todo
func (c *Client) Action() action.ServiceClient {
	return action.NewServiceClient(c.conn)
}

// Example todo
func (c *Client) Template() template.ServiceClient {
	return template.NewServiceClient(c.conn)
}

// Example todo
func (c *Client) Deploy() deploy.ServiceClient {
	return deploy.NewServiceClient(c.conn)
}
