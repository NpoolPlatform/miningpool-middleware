package orderuser

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	orderusercrud "github.com/NpoolPlatform/miningpool-middleware/pkg/crud/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent"
	orderuserent "github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
)

type existHandler struct {
	*Handler
	checkAppAuthSQL string
}

func (h *existHandler) existJoinCoinAndAppPool() error {
	if h.AppID == nil || h.GoodUserID == nil {
		return wlog.Errorf("invalid appid or gooduserid")
	}

	sql := fmt.Sprintf(`SELECT DISTINCT 
	coin_type,revenue_type,root_user_id,t1.ent_id,app_id 
	FROM good_users AS t1 
	LEFT JOIN coins AS t2 ON t1.coin_id = t2.ent_id 
	LEFT JOIN app_pools AS t3 ON t2.pool_id = t3.pool_id 
	WHERE (t1.ent_id = '%v' AND t3.app_id = '%v') AND t3.deleted_at = 0;`,
		h.GoodUserID.String(),
		h.AppID.String(),
	)

	h.checkAppAuthSQL = sql
	return nil
}

func (h *Handler) checkAppAuth(ctx context.Context) (bool, error) {
	exist := false
	var err error
	existH := &existHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		err := existH.existJoinCoinAndAppPool()
		if err != nil {
			return err
		}

		rows, err := cli.QueryContext(ctx, existH.checkAppAuthSQL)
		if err != nil {
			return err
		}
		defer rows.Close()

		if err := rows.Err(); err != nil {
			return err
		}

		cols, err := rows.Columns()
		if err != nil {
			return err
		}
		exist = len(cols) > 0
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
