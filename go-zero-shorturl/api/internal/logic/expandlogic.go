package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"shortenurl/api/internal/svc"
	"shortenurl/api/internal/types"
	"shortenurl/rpc/transform/transform"
)

type ExpandLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewExpandLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExpandLogic {
	return &ExpandLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExpandLogic) Expand(req *types.ExpandReq) (resp *types.ExpandResq, err error) {
	rpcResp, err := l.svcCtx.Transformer.Expand(l.ctx, &transform.ExpandReq{
		Shorten: req.Shorten,
	})
	if err != nil {
		return nil, err
	}

	return &types.ExpandResq{
		Url: rpcResp.Url,
	}, nil
}
