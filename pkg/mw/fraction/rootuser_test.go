package fraction

import (
	"context"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	poolmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/stretchr/testify/assert"
)

var rootuserRet = &npool.RootUser{
	EntID:          uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	Name:           pools.RandomPoolUserNameForTest(),
	Email:          "gggo@go.go",
	AuthToken:      "7ecdq1fosdsfcruypom2otsn7hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73",
	Authed:         true,
	Remark:         "asdfaf",
}

var rootuserReq = &npool.RootUserReq{
	EntID:     &rootuserRet.EntID,
	Email:     &rootuserRet.Email,
	Name:      &rootuserRet.Name,
	AuthToken: &rootuserRet.AuthToken,
	Authed:    &rootuserRet.Authed,
	Remark:    &rootuserRet.Remark,
}

func createRootUser(t *testing.T) {
	poolH, err := pool.NewHandler(context.Background(),
		pool.WithConds(&poolmw.Conds{
			MiningpoolType: &v1.Uint32Val{
				Op:    cruder.EQ,
				Value: uint32(rootuserRet.MiningpoolType),
			},
		}),
		pool.WithLimit(2),
		pool.WithOffset(0),
	)
	assert.Nil(t, err)

	poolInfos, _, err := poolH.GetPools(context.Background())
	assert.Nil(t, err)
	if !assert.NotEqual(t, len(poolInfos), 0) {
		return
	}

	rootuserRet.PoolID = poolInfos[0].EntID
	rootuserReq.PoolID = &poolInfos[0].EntID

	handler, err := rootuser.NewHandler(
		context.Background(),
		rootuser.WithEntID(rootuserReq.EntID, true),
		rootuser.WithPoolID(rootuserReq.PoolID, true),
		rootuser.WithName(rootuserReq.Name, true),
		rootuser.WithEmail(rootuserReq.Email, true),
		rootuser.WithAuthToken(rootuserReq.AuthToken, true),
		rootuser.WithAuthed(rootuserReq.Authed, true),
		rootuser.WithRemark(rootuserReq.Remark, true),
	)
	assert.Nil(t, err)

	err = handler.CreateRootUser(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetRootUser(context.Background())
	if assert.Nil(t, err) {
		rootuserRet.UpdatedAt = info.UpdatedAt
		rootuserRet.CreatedAt = info.CreatedAt
		rootuserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		rootuserRet.AuthToken = info.AuthToken
		rootuserRet.ID = info.ID
		rootuserRet.EntID = info.EntID
		assert.Equal(t, info, rootuserRet)
	}
}

func deleteRootUser(t *testing.T) {
	handler, err := rootuser.NewHandler(
		context.Background(),
		rootuser.WithID(&rootuserRet.ID, true),
		rootuser.WithEntID(&rootuserRet.EntID, true),
	)
	assert.Nil(t, err)
	err = handler.DeleteRootUser(context.Background())
	assert.Nil(t, err)
}
