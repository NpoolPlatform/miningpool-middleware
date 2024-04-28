package fraction

import (
	"context"
	"testing"

	gooduserclient "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	v1 "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
)

var goodUserRet = &npool.GoodUser{
	EntID:          uuid.NewString(),
	Name:           "fffff",
	RootUserID:     rootUserRet.EntID,
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	CoinType:       basetypes.CoinType_BitCoin,
	HashRate:       66,
	ReadPageLink:   "fffff",
	RevenueType:    basetypes.RevenueType_FPPS,
}

var goodUserReq = &npool.GoodUserReq{
	EntID:        &goodUserRet.EntID,
	Name:         &goodUserRet.Name,
	RootUserID:   &goodUserRet.RootUserID,
	HashRate:     &goodUserRet.HashRate,
	ReadPageLink: &goodUserRet.ReadPageLink,
}

func createGoodUser(t *testing.T) {
	err := gooduserclient.CreateGoodUser(context.Background(), goodUserReq)
	assert.Nil(t, err)

	info, err := gooduserclient.GetGoodUser(context.Background(), *goodUserReq.EntID)
	if assert.Nil(t, err) {
		goodUserRet.Name = info.Name
		goodUserRet.ReadPageLink = info.ReadPageLink
		goodUserRet.CreatedAt = info.CreatedAt
		goodUserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		goodUserRet.CoinTypeStr = info.CoinTypeStr
		goodUserRet.RevenueTypeStr = info.RevenueTypeStr
		goodUserRet.UpdatedAt = info.UpdatedAt
		goodUserRet.ID = info.ID
		goodUserRet.EntID = info.EntID
		assert.Equal(t, goodUserRet, info)
	}
}

func deleteGoodUser(t *testing.T) {
	exist, err := gooduserclient.ExistGoodUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: goodUserRet.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, true, exist)
	}

	err = gooduserclient.DeleteGoodUser(context.Background(), goodUserRet.ID, goodUserRet.EntID)
	assert.Nil(t, err)

	exist, err = gooduserclient.ExistGoodUserConds(context.Background(), &npool.Conds{
		EntID: &v1.StringVal{
			Op:    cruder.EQ,
			Value: goodUserRet.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}
