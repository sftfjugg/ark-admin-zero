package job

import (
	"context"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysJobLogic {
	return &DeleteSysJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysJobLogic) DeleteSysJob(req *types.DeleteSysJobReq) error {
	userList, _ := l.svcCtx.SysUserModel.FindByCondition(l.ctx, "job_id", req.Id)
	if len(userList) != 0 {
		return errorx.NewDefaultError(errorx.DeleteJobErrorCode)
	}

	err := l.svcCtx.SysJobModel.Delete(l.ctx, req.Id)
	if err != nil {
		return errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	return nil
}
