package gooduser

import (
	"context"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"

	"github.com/stretchr/testify/assert"
)

var rootuserRet = &npool.RootUser{
	EntID:          uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	Email:          "gggo@go.go",
	AuthToken:      "7ecdq1fosdsfcruypom2otsn8hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73",
	Authed:         true,
	Remark:         "asdfaf",
}

var rootuserReq = &npool.RootUserReq{
	EntID:     &rootuserRet.EntID,
	Email:     &rootuserRet.Email,
	AuthToken: &rootuserRet.AuthToken,
	Authed:    &rootuserRet.Authed,
	Remark:    &rootuserRet.Remark,
}

func createRootUser(t *testing.T) {
	name, err := f2pool.RandomF2PoolUser(8)
	assert.Nil(t, err)

	rootuserRet.Name = name
	rootuserReq.Name = &name

	handler, err := rootuser.NewHandler(
		context.Background(),
		rootuser.WithEntID(rootuserReq.EntID, true),
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
