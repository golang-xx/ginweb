package logic

import (
	"context"
	"fmt"

	"ginweb02/internal/svc"
	"ginweb02/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type EdituserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEdituserLogic(ctx context.Context, svcCtx *svc.ServiceContext) EdituserLogic {
	return EdituserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EdituserLogic) Edituser(req types.UserUpdateReq) (*types.UserOptResp, error) {
	// todo: add your logic here and delete this line
	fmt.Println("edit user")
	l.svcCtx.DbEngin.Updates(req)
	return &types.UserOptResp{
		Id: req.Id,Token: "333",
	}, nil
}
