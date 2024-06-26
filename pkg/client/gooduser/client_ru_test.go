package gooduser

import (
	"context"
	"testing"

	poolmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/client/pool"
	rootuserclient "github.com/NpoolPlatform/miningpool-middleware/pkg/client/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/pools/f2pool"
	"github.com/stretchr/testify/assert"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/rootuser"

	"github.com/google/uuid"
)

var rootUserRet = &npool.RootUser{
	EntID:          uuid.NewString(),
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	Email:          "sssss@ss.com",
	AuthToken:      "7ecdq1fosdsfcruypom2otsn7hfr69azmqvh7v3zelol1ntsba85a1yvol66qp73",
	Authed:         true,
	Remark:         "sdfasdfasdf",
}

var rootUserReq = &npool.RootUserReq{
	EntID:     &rootUserRet.EntID,
	Email:     &rootUserRet.Email,
	AuthToken: &rootUserRet.AuthToken,
	Authed:    &rootUserRet.Authed,
	Remark:    &rootUserRet.Remark,
}

func createRootUser(t *testing.T) {
	poolInfos, _, err := pool.GetPools(context.Background(), &poolmw.Conds{
		MiningpoolType: &v1.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(ret.MiningpoolType),
		},
	}, 0, 2)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(poolInfos))

	rootUserRet.PoolID = poolInfos[0].EntID
	rootUserReq.PoolID = &poolInfos[0].EntID

	name, err := f2pool.RandomF2PoolUser(7)
	assert.Nil(t, err)

	rootUserRet.Name = name
	rootUserReq.Name = &name

	err = rootuserclient.CreateRootUser(context.Background(), rootUserReq)
	assert.Nil(t, err)

	info, err := rootuserclient.GetRootUser(context.Background(), *rootUserReq.EntID)
	if assert.Nil(t, err) {
		rootUserRet.CreatedAt = info.CreatedAt
		rootUserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		rootUserRet.UpdatedAt = info.UpdatedAt
		rootUserRet.ID = info.ID
		rootUserRet.AuthToken = info.AuthToken
		rootUserRet.EntID = info.EntID
		assert.Equal(t, rootUserRet, info)
	}
}

func deleteRootUser(t *testing.T) {
	exist, err := rootuserclient.ExistRootUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: rootUserRet.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	err = rootuserclient.DeleteRootUser(context.Background(), rootUserRet.ID, rootUserRet.EntID)
	assert.Nil(t, err)

	exist, err = rootuserclient.ExistRootUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: rootUserRet.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}
