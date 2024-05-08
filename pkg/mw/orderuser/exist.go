package orderuser

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	orderusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/apppool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
	orderuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
)

type existHandler struct {
	*Handler
	conds *orderusercrud.Conds
}

func (h *existHandler) existJoinCoinAndAppPool(s *sql.Selector) {
	guT := sql.Table(gooduser.Table)
	coinT := sql.Table(coin.Table)
	apppoolT := sql.Table(apppool.Table)

	s.LeftJoin(guT).On(
		s.C(orderuserent.FieldGoodUserID),
		guT.C(gooduser.FieldEntID),
	).LeftJoin(coinT).On(
		guT.C(gooduser.FieldCoinID),
		coinT.C(coin.FieldEntID),
	).LeftJoin(apppoolT).On(
		coinT.C(coin.FieldPoolID),
		apppoolT.C(apppool.FieldPoolID),
	).AppendSelect(
		coin.FieldCoinType,
		coin.FieldRevenueType,
		gooduser.FieldRootUserID,
		apppool.FieldAppID,
	)
}

func (h *Handler) checkAppAuth(ctx context.Context) (bool, error) {
	exist := false
	var err error
	existH := &existHandler{
		Handler: h,
		conds: &orderusercrud.Conds{
			GoodUserID: &cruder.Cond{
				Op:  cruder.EQ,
				Val: *h.GoodUserID,
			},
			AppID: &cruder.Cond{
				Op:  cruder.EQ,
				Val: *h.AppID,
			},
		},
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := orderusercrud.SetQueryConds(cli.OrderUser.Query(), existH.Conds)
		if err != nil {
			return err
		}

		stm.Modify(existH.existJoinCoinAndAppPool)

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistOrderUser(ctx context.Context) (bool, error) {
	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			OrderUser.
			Query().
			Where(
				orderuserent.EntID(*h.EntID),
				orderuserent.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistOrderUserConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := orderusercrud.SetQueryConds(cli.OrderUser.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
