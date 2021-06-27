// Code generated by protoc-gen-go-http. DO NOT EDIT.

package pipeline

import (
	http "github.com/infraboard/mcube/pb/http"
)

// HttpEntry todo
func HttpEntry() *http.EntrySet {
	set := &http.EntrySet{
		Items: []*http.Entry{
			{
				Path:         "/workflow.pipeline.Service/CreatePipeline",
				FunctionName: "CreatePipeline",
			},
			{
				Path:         "/workflow.pipeline.Service/QueryPipeline",
				FunctionName: "QueryPipeline",
			},
			{
				Path:         "/workflow.pipeline.Service/DescribePipeline",
				FunctionName: "DescribePipeline",
			},
			{
				Path:         "/workflow.pipeline.Service/DeletePipeline",
				FunctionName: "DeletePipeline",
			},
			{
				Path:         "/workflow.pipeline.Service/CreateAction",
				FunctionName: "CreateAction",
			},
			{
				Path:         "/workflow.pipeline.Service/QueryAction",
				FunctionName: "QueryAction",
			},
		},
	}
	return set
}
