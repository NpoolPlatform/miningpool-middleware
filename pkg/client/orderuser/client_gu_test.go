package orderuser

import (
	"context"
	"testing"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/client/coin"
	gooduserclient "github.com/NpoolPlatform/miningpool-middleware/pkg/client/gooduser"
	"github.com/google/uuid"

	coinmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	"github.com/stretchr/testify/assert"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
)

var goodUserRet = &npool.GoodUser{
	EntID:          uuid.NewString(),
	Name:           "fffff",
	RootUserID:     rootUserRet.EntID,
	PoolCoinTypeID: uuid.NewString(),
	MiningpoolType: mpbasetypes.MiningpoolType_F2Pool,
	CoinType:       basetypes.CoinType_CoinTypeBitCoin,
	ReadPageLink:   "fffff",
	RevenueType:    mpbasetypes.RevenueType_FPPS,
}

var goodUserReq = &npool.GoodUserReq{
	EntID:          &goodUserRet.EntID,
	PoolCoinTypeID: &goodUserRet.PoolCoinTypeID,
	Name:           &goodUserRet.Name,
	RootUserID:     &goodUserRet.RootUserID,
	ReadPageLink:   &goodUserRet.ReadPageLink,
}

func createGoodUser(t *testing.T) {
	coinInfos, _, err := coin.GetCoins(context.Background(), &coinmw.Conds{
		CoinType: &basetypes.Uint32Val{
			Op:    cruder.EQ,
			Value: uint32(ret.CoinType),
		},
		PoolID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: rootUserRet.PoolID,
		},
	}, 0, 2)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, len(coinInfos))

	goodUserRet.PoolCoinTypeID = coinInfos[0].EntID
	goodUserReq.PoolCoinTypeID = &coinInfos[0].EntID

	err = gooduserclient.CreateGoodUser(context.Background(), goodUserReq)
	assert.Nil(t, err)

	info, err := gooduserclient.GetGoodUser(context.Background(), *goodUserReq.EntID)
	if assert.Nil(t, err) {
		goodUserRet.Name = info.Name
		goodUserRet.ReadPageLink = info.ReadPageLink
		goodUserRet.CreatedAt = info.CreatedAt
		goodUserRet.PoolID = info.PoolID
		goodUserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
		goodUserRet.CoinTypeStr = info.CoinTypeStr
		goodUserRet.RevenueTypeStr = info.RevenueTypeStr
		goodUserRet.UpdatedAt = info.UpdatedAt
		goodUserRet.FeeRatio = info.FeeRatio
		goodUserRet.ID = info.ID
		goodUserRet.EntID = info.EntID
		assert.Equal(t, goodUserRet, info)
	}
}

func deleteGoodUser(t *testing.T) {
	exist, err := gooduserclient.ExistGoodUserConds(context.Background(), &npool.Conds{
		EntID: &basetypes.StringVal{
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
		EntID: &basetypes.StringVal{
			Op:    cruder.EQ,
			Value: goodUserRet.EntID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, false, exist)
	}
}
