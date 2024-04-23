package orderuser

import (
	"context"
	"testing"

	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"

	"github.com/stretchr/testify/assert"
)

var gooduserRet = &npool.GoodUser{
	EntID:          uuid.NewString(),
	RootUserID:     rootuserRet.EntID,
	MiningpoolType: basetypes.MiningpoolType_F2Pool,
	CoinType:       basetypes.CoinType_BitCoin,
	HashRate:       5.0,
	RevenueType:    basetypes.RevenueType_FPPS,
}

var gooduserReq = &npool.GoodUserReq{
	EntID:          &gooduserRet.EntID,
	RootUserID:     &gooduserRet.RootUserID,
	MiningpoolType: &gooduserRet.MiningpoolType,
	CoinType:       &gooduserRet.CoinType,
	HashRate:       &gooduserRet.HashRate,
	RevenueType:    &gooduserRet.RevenueType,
}

func createGoodUser(t *testing.T) {
	handler, err := gooduser.NewHandler(
		context.Background(),
		gooduser.WithEntID(gooduserReq.EntID, true),
		gooduser.WithRootUserID(gooduserReq.RootUserID, true),
		gooduser.WithMiningpoolType(gooduserReq.MiningpoolType, true),
		gooduser.WithCoinType(gooduserReq.CoinType, true),
		gooduser.WithHashRate(gooduserReq.HashRate, true),
		gooduser.WithRevenueType(gooduserReq.RevenueType, true),
	)
	if !assert.Nil(t, err) {
		return
	}

	err = handler.CreateGoodUser(context.Background())
	if !assert.Nil(t, err) {
		return
	}

	info, err := handler.GetGoodUser(context.Background())
	if assert.Nil(t, err) {
		gooduserRet.UpdatedAt = info.UpdatedAt
		gooduserRet.CreatedAt = info.CreatedAt
		gooduserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		gooduserRet.CoinTypeStr = info.CoinTypeStr
		gooduserRet.RevenueTypeStr = info.RevenueTypeStr
		gooduserRet.ID = info.ID
		gooduserRet.EntID = info.EntID
		gooduserRet.Name = info.Name
		gooduserRet.ReadPageLink = info.ReadPageLink
		assert.Equal(t, info, gooduserRet)
	}
}

func deleteGoodUser(t *testing.T) {
	handler, err := gooduser.NewHandler(
		context.Background(),
		gooduser.WithID(&gooduserRet.ID, true),
		gooduser.WithEntID(&gooduserRet.EntID, true),
	)
	assert.Nil(t, err)
	err = handler.DeleteGoodUser(context.Background())
	assert.Nil(t, err)
}