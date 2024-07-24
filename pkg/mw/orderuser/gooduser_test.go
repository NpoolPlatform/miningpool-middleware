package orderuser

import (
	"context"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mpbasetypes "github.com/NpoolPlatform/message/npool/basetypes/miningpool/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	coinmw "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/coin"
	npool "github.com/NpoolPlatform/message/npool/miningpool/mw/v1/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/mw/gooduser"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var gooduserRet = &npool.GoodUser{
	EntID:          uuid.NewString(),
	RootUserID:     rootuserRet.EntID,
	MiningpoolType: mpbasetypes.MiningpoolType_F2Pool,
}

var gooduserReq = &npool.GoodUserReq{
	EntID:      &gooduserRet.EntID,
	RootUserID: &gooduserRet.RootUserID,
}

func createGoodUser(t *testing.T) {
	coinH, err := coin.NewHandler(context.Background(),
		coin.WithConds(&coinmw.Conds{
			PoolID: &basetypes.StringVal{
				Op:    cruder.EQ,
				Value: rootuserRet.PoolID,
			},
		}),
		coin.WithOffset(0),
		coin.WithLimit(2),
	)
	assert.Nil(t, err)

	coinInfos, _, err := coinH.GetCoins(context.Background())
	assert.Nil(t, err)

	if len(coinInfos) == 0 {
		return
	}

	for _, coinInfo := range coinInfos {
		gooduserReq.CoinTypeIDs = append(gooduserReq.CoinTypeIDs, coinInfo.CoinTypeID)
	}

	handler, err := gooduser.NewHandler(
		context.Background(),
		gooduser.WithEntID(gooduserReq.EntID, true),
		gooduser.WithRootUserID(gooduserReq.RootUserID, true),
		gooduser.WithCoinTypeIDs(gooduserReq.CoinTypeIDs, true),
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
		gooduserRet.PoolID = info.PoolID
		gooduserRet.MiningpoolTypeStr = info.MiningpoolTypeStr
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
