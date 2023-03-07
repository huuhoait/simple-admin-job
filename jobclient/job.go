// Code generated by goctl. DO NOT EDIT.
// Source: job.proto

package jobclient

import (
	"context"

	"github.com/suyuan32/simple-admin-job/job"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp   = job.BaseIDResp
	BaseResp     = job.BaseResp
	BaseUUIDResp = job.BaseUUIDResp
	Empty        = job.Empty
	IDReq        = job.IDReq
	IDsReq       = job.IDsReq
	PageInfoReq  = job.PageInfoReq
	UUIDReq      = job.UUIDReq
	UUIDsReq     = job.UUIDsReq

	Job interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultJob struct {
		cli zrpc.Client
	}
)

func NewJob(cli zrpc.Client) Job {
	return &defaultJob{
		cli: cli,
	}
}

func (m *defaultJob) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := job.NewJobClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}
