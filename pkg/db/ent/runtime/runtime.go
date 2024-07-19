// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/apppool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/schema"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	apppoolMixin := schema.AppPool{}.Mixin()
	apppool.Policy = privacy.NewPolicies(apppoolMixin[0], schema.AppPool{})
	apppool.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := apppool.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	apppoolMixinFields0 := apppoolMixin[0].Fields()
	_ = apppoolMixinFields0
	apppoolMixinFields1 := apppoolMixin[1].Fields()
	_ = apppoolMixinFields1
	apppoolFields := schema.AppPool{}.Fields()
	_ = apppoolFields
	// apppoolDescCreatedAt is the schema descriptor for created_at field.
	apppoolDescCreatedAt := apppoolMixinFields0[0].Descriptor()
	// apppool.DefaultCreatedAt holds the default value on creation for the created_at field.
	apppool.DefaultCreatedAt = apppoolDescCreatedAt.Default.(func() uint32)
	// apppoolDescUpdatedAt is the schema descriptor for updated_at field.
	apppoolDescUpdatedAt := apppoolMixinFields0[1].Descriptor()
	// apppool.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	apppool.DefaultUpdatedAt = apppoolDescUpdatedAt.Default.(func() uint32)
	// apppool.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	apppool.UpdateDefaultUpdatedAt = apppoolDescUpdatedAt.UpdateDefault.(func() uint32)
	// apppoolDescDeletedAt is the schema descriptor for deleted_at field.
	apppoolDescDeletedAt := apppoolMixinFields0[2].Descriptor()
	// apppool.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	apppool.DefaultDeletedAt = apppoolDescDeletedAt.Default.(func() uint32)
	// apppoolDescEntID is the schema descriptor for ent_id field.
	apppoolDescEntID := apppoolMixinFields1[1].Descriptor()
	// apppool.DefaultEntID holds the default value on creation for the ent_id field.
	apppool.DefaultEntID = apppoolDescEntID.Default.(func() uuid.UUID)
	// apppoolDescAppID is the schema descriptor for app_id field.
	apppoolDescAppID := apppoolFields[0].Descriptor()
	// apppool.DefaultAppID holds the default value on creation for the app_id field.
	apppool.DefaultAppID = apppoolDescAppID.Default.(func() uuid.UUID)
	// apppoolDescPoolID is the schema descriptor for pool_id field.
	apppoolDescPoolID := apppoolFields[1].Descriptor()
	// apppool.DefaultPoolID holds the default value on creation for the pool_id field.
	apppool.DefaultPoolID = apppoolDescPoolID.Default.(func() uuid.UUID)
	coinMixin := schema.Coin{}.Mixin()
	coin.Policy = privacy.NewPolicies(coinMixin[0], schema.Coin{})
	coin.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := coin.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	coinMixinFields0 := coinMixin[0].Fields()
	_ = coinMixinFields0
	coinMixinFields1 := coinMixin[1].Fields()
	_ = coinMixinFields1
	coinFields := schema.Coin{}.Fields()
	_ = coinFields
	// coinDescCreatedAt is the schema descriptor for created_at field.
	coinDescCreatedAt := coinMixinFields0[0].Descriptor()
	// coin.DefaultCreatedAt holds the default value on creation for the created_at field.
	coin.DefaultCreatedAt = coinDescCreatedAt.Default.(func() uint32)
	// coinDescUpdatedAt is the schema descriptor for updated_at field.
	coinDescUpdatedAt := coinMixinFields0[1].Descriptor()
	// coin.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	coin.DefaultUpdatedAt = coinDescUpdatedAt.Default.(func() uint32)
	// coin.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	coin.UpdateDefaultUpdatedAt = coinDescUpdatedAt.UpdateDefault.(func() uint32)
	// coinDescDeletedAt is the schema descriptor for deleted_at field.
	coinDescDeletedAt := coinMixinFields0[2].Descriptor()
	// coin.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	coin.DefaultDeletedAt = coinDescDeletedAt.Default.(func() uint32)
	// coinDescEntID is the schema descriptor for ent_id field.
	coinDescEntID := coinMixinFields1[1].Descriptor()
	// coin.DefaultEntID holds the default value on creation for the ent_id field.
	coin.DefaultEntID = coinDescEntID.Default.(func() uuid.UUID)
	// coinDescPoolID is the schema descriptor for pool_id field.
	coinDescPoolID := coinFields[0].Descriptor()
	// coin.DefaultPoolID holds the default value on creation for the pool_id field.
	coin.DefaultPoolID = coinDescPoolID.Default.(func() uuid.UUID)
	// coinDescCoinTypeID is the schema descriptor for coin_type_id field.
	coinDescCoinTypeID := coinFields[1].Descriptor()
	// coin.DefaultCoinTypeID holds the default value on creation for the coin_type_id field.
	coin.DefaultCoinTypeID = coinDescCoinTypeID.Default.(func() uuid.UUID)
	// coinDescCoinType is the schema descriptor for coin_type field.
	coinDescCoinType := coinFields[2].Descriptor()
	// coin.DefaultCoinType holds the default value on creation for the coin_type field.
	coin.DefaultCoinType = coinDescCoinType.Default.(string)
	// coinDescFeeRatio is the schema descriptor for fee_ratio field.
	coinDescFeeRatio := coinFields[3].Descriptor()
	// coin.DefaultFeeRatio holds the default value on creation for the fee_ratio field.
	coin.DefaultFeeRatio = coinDescFeeRatio.Default.(decimal.Decimal)
	// coinDescFixedRevenueAble is the schema descriptor for fixed_revenue_able field.
	coinDescFixedRevenueAble := coinFields[4].Descriptor()
	// coin.DefaultFixedRevenueAble holds the default value on creation for the fixed_revenue_able field.
	coin.DefaultFixedRevenueAble = coinDescFixedRevenueAble.Default.(bool)
	// coinDescLeastTransferAmount is the schema descriptor for least_transfer_amount field.
	coinDescLeastTransferAmount := coinFields[5].Descriptor()
	// coin.DefaultLeastTransferAmount holds the default value on creation for the least_transfer_amount field.
	coin.DefaultLeastTransferAmount = coinDescLeastTransferAmount.Default.(decimal.Decimal)
	// coinDescBenefitIntervalSeconds is the schema descriptor for benefit_interval_seconds field.
	coinDescBenefitIntervalSeconds := coinFields[6].Descriptor()
	// coin.DefaultBenefitIntervalSeconds holds the default value on creation for the benefit_interval_seconds field.
	coin.DefaultBenefitIntervalSeconds = coinDescBenefitIntervalSeconds.Default.(uint32)
	// coinDescRemark is the schema descriptor for remark field.
	coinDescRemark := coinFields[7].Descriptor()
	// coin.DefaultRemark holds the default value on creation for the remark field.
	coin.DefaultRemark = coinDescRemark.Default.(string)
	fractionMixin := schema.Fraction{}.Mixin()
	fraction.Policy = privacy.NewPolicies(fractionMixin[0], schema.Fraction{})
	fraction.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := fraction.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	fractionMixinFields0 := fractionMixin[0].Fields()
	_ = fractionMixinFields0
	fractionMixinFields1 := fractionMixin[1].Fields()
	_ = fractionMixinFields1
	fractionFields := schema.Fraction{}.Fields()
	_ = fractionFields
	// fractionDescCreatedAt is the schema descriptor for created_at field.
	fractionDescCreatedAt := fractionMixinFields0[0].Descriptor()
	// fraction.DefaultCreatedAt holds the default value on creation for the created_at field.
	fraction.DefaultCreatedAt = fractionDescCreatedAt.Default.(func() uint32)
	// fractionDescUpdatedAt is the schema descriptor for updated_at field.
	fractionDescUpdatedAt := fractionMixinFields0[1].Descriptor()
	// fraction.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	fraction.DefaultUpdatedAt = fractionDescUpdatedAt.Default.(func() uint32)
	// fraction.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	fraction.UpdateDefaultUpdatedAt = fractionDescUpdatedAt.UpdateDefault.(func() uint32)
	// fractionDescDeletedAt is the schema descriptor for deleted_at field.
	fractionDescDeletedAt := fractionMixinFields0[2].Descriptor()
	// fraction.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	fraction.DefaultDeletedAt = fractionDescDeletedAt.Default.(func() uint32)
	// fractionDescEntID is the schema descriptor for ent_id field.
	fractionDescEntID := fractionMixinFields1[1].Descriptor()
	// fraction.DefaultEntID holds the default value on creation for the ent_id field.
	fraction.DefaultEntID = fractionDescEntID.Default.(func() uuid.UUID)
	// fractionDescAppID is the schema descriptor for app_id field.
	fractionDescAppID := fractionFields[0].Descriptor()
	// fraction.DefaultAppID holds the default value on creation for the app_id field.
	fraction.DefaultAppID = fractionDescAppID.Default.(func() uuid.UUID)
	// fractionDescUserID is the schema descriptor for user_id field.
	fractionDescUserID := fractionFields[1].Descriptor()
	// fraction.DefaultUserID holds the default value on creation for the user_id field.
	fraction.DefaultUserID = fractionDescUserID.Default.(func() uuid.UUID)
	// fractionDescOrderUserID is the schema descriptor for order_user_id field.
	fractionDescOrderUserID := fractionFields[2].Descriptor()
	// fraction.DefaultOrderUserID holds the default value on creation for the order_user_id field.
	fraction.DefaultOrderUserID = fractionDescOrderUserID.Default.(func() uuid.UUID)
	// fractionDescWithdrawState is the schema descriptor for withdraw_state field.
	fractionDescWithdrawState := fractionFields[3].Descriptor()
	// fraction.DefaultWithdrawState holds the default value on creation for the withdraw_state field.
	fraction.DefaultWithdrawState = fractionDescWithdrawState.Default.(string)
	// fractionDescWithdrawAt is the schema descriptor for withdraw_at field.
	fractionDescWithdrawAt := fractionFields[4].Descriptor()
	// fraction.DefaultWithdrawAt holds the default value on creation for the withdraw_at field.
	fraction.DefaultWithdrawAt = fractionDescWithdrawAt.Default.(uint32)
	// fractionDescPromisePayAt is the schema descriptor for promise_pay_at field.
	fractionDescPromisePayAt := fractionFields[5].Descriptor()
	// fraction.DefaultPromisePayAt holds the default value on creation for the promise_pay_at field.
	fraction.DefaultPromisePayAt = fractionDescPromisePayAt.Default.(uint32)
	// fractionDescMsg is the schema descriptor for msg field.
	fractionDescMsg := fractionFields[6].Descriptor()
	// fraction.DefaultMsg holds the default value on creation for the msg field.
	fraction.DefaultMsg = fractionDescMsg.Default.(string)
	fractionruleMixin := schema.FractionRule{}.Mixin()
	fractionrule.Policy = privacy.NewPolicies(fractionruleMixin[0], schema.FractionRule{})
	fractionrule.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := fractionrule.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	fractionruleMixinFields0 := fractionruleMixin[0].Fields()
	_ = fractionruleMixinFields0
	fractionruleMixinFields1 := fractionruleMixin[1].Fields()
	_ = fractionruleMixinFields1
	fractionruleFields := schema.FractionRule{}.Fields()
	_ = fractionruleFields
	// fractionruleDescCreatedAt is the schema descriptor for created_at field.
	fractionruleDescCreatedAt := fractionruleMixinFields0[0].Descriptor()
	// fractionrule.DefaultCreatedAt holds the default value on creation for the created_at field.
	fractionrule.DefaultCreatedAt = fractionruleDescCreatedAt.Default.(func() uint32)
	// fractionruleDescUpdatedAt is the schema descriptor for updated_at field.
	fractionruleDescUpdatedAt := fractionruleMixinFields0[1].Descriptor()
	// fractionrule.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	fractionrule.DefaultUpdatedAt = fractionruleDescUpdatedAt.Default.(func() uint32)
	// fractionrule.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	fractionrule.UpdateDefaultUpdatedAt = fractionruleDescUpdatedAt.UpdateDefault.(func() uint32)
	// fractionruleDescDeletedAt is the schema descriptor for deleted_at field.
	fractionruleDescDeletedAt := fractionruleMixinFields0[2].Descriptor()
	// fractionrule.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	fractionrule.DefaultDeletedAt = fractionruleDescDeletedAt.Default.(func() uint32)
	// fractionruleDescEntID is the schema descriptor for ent_id field.
	fractionruleDescEntID := fractionruleMixinFields1[1].Descriptor()
	// fractionrule.DefaultEntID holds the default value on creation for the ent_id field.
	fractionrule.DefaultEntID = fractionruleDescEntID.Default.(func() uuid.UUID)
	// fractionruleDescPoolCoinTypeID is the schema descriptor for pool_coin_type_id field.
	fractionruleDescPoolCoinTypeID := fractionruleFields[0].Descriptor()
	// fractionrule.DefaultPoolCoinTypeID holds the default value on creation for the pool_coin_type_id field.
	fractionrule.DefaultPoolCoinTypeID = fractionruleDescPoolCoinTypeID.Default.(func() uuid.UUID)
	// fractionruleDescWithdrawInterval is the schema descriptor for withdraw_interval field.
	fractionruleDescWithdrawInterval := fractionruleFields[1].Descriptor()
	// fractionrule.DefaultWithdrawInterval holds the default value on creation for the withdraw_interval field.
	fractionrule.DefaultWithdrawInterval = fractionruleDescWithdrawInterval.Default.(uint32)
	// fractionruleDescMinAmount is the schema descriptor for min_amount field.
	fractionruleDescMinAmount := fractionruleFields[2].Descriptor()
	// fractionrule.DefaultMinAmount holds the default value on creation for the min_amount field.
	fractionrule.DefaultMinAmount = fractionruleDescMinAmount.Default.(decimal.Decimal)
	// fractionruleDescWithdrawRate is the schema descriptor for withdraw_rate field.
	fractionruleDescWithdrawRate := fractionruleFields[3].Descriptor()
	// fractionrule.DefaultWithdrawRate holds the default value on creation for the withdraw_rate field.
	fractionrule.DefaultWithdrawRate = fractionruleDescWithdrawRate.Default.(decimal.Decimal)
	gooduserMixin := schema.GoodUser{}.Mixin()
	gooduser.Policy = privacy.NewPolicies(gooduserMixin[0], schema.GoodUser{})
	gooduser.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := gooduser.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	gooduserMixinFields0 := gooduserMixin[0].Fields()
	_ = gooduserMixinFields0
	gooduserMixinFields1 := gooduserMixin[1].Fields()
	_ = gooduserMixinFields1
	gooduserFields := schema.GoodUser{}.Fields()
	_ = gooduserFields
	// gooduserDescCreatedAt is the schema descriptor for created_at field.
	gooduserDescCreatedAt := gooduserMixinFields0[0].Descriptor()
	// gooduser.DefaultCreatedAt holds the default value on creation for the created_at field.
	gooduser.DefaultCreatedAt = gooduserDescCreatedAt.Default.(func() uint32)
	// gooduserDescUpdatedAt is the schema descriptor for updated_at field.
	gooduserDescUpdatedAt := gooduserMixinFields0[1].Descriptor()
	// gooduser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	gooduser.DefaultUpdatedAt = gooduserDescUpdatedAt.Default.(func() uint32)
	// gooduser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	gooduser.UpdateDefaultUpdatedAt = gooduserDescUpdatedAt.UpdateDefault.(func() uint32)
	// gooduserDescDeletedAt is the schema descriptor for deleted_at field.
	gooduserDescDeletedAt := gooduserMixinFields0[2].Descriptor()
	// gooduser.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	gooduser.DefaultDeletedAt = gooduserDescDeletedAt.Default.(func() uint32)
	// gooduserDescEntID is the schema descriptor for ent_id field.
	gooduserDescEntID := gooduserMixinFields1[1].Descriptor()
	// gooduser.DefaultEntID holds the default value on creation for the ent_id field.
	gooduser.DefaultEntID = gooduserDescEntID.Default.(func() uuid.UUID)
	// gooduserDescRootUserID is the schema descriptor for root_user_id field.
	gooduserDescRootUserID := gooduserFields[0].Descriptor()
	// gooduser.DefaultRootUserID holds the default value on creation for the root_user_id field.
	gooduser.DefaultRootUserID = gooduserDescRootUserID.Default.(func() uuid.UUID)
	// gooduserDescName is the schema descriptor for name field.
	gooduserDescName := gooduserFields[1].Descriptor()
	// gooduser.DefaultName holds the default value on creation for the name field.
	gooduser.DefaultName = gooduserDescName.Default.(string)
	// gooduserDescPoolCoinTypeID is the schema descriptor for pool_coin_type_id field.
	gooduserDescPoolCoinTypeID := gooduserFields[2].Descriptor()
	// gooduser.DefaultPoolCoinTypeID holds the default value on creation for the pool_coin_type_id field.
	gooduser.DefaultPoolCoinTypeID = gooduserDescPoolCoinTypeID.Default.(func() uuid.UUID)
	// gooduserDescReadPageLink is the schema descriptor for read_page_link field.
	gooduserDescReadPageLink := gooduserFields[3].Descriptor()
	// gooduser.DefaultReadPageLink holds the default value on creation for the read_page_link field.
	gooduser.DefaultReadPageLink = gooduserDescReadPageLink.Default.(string)
	orderuserMixin := schema.OrderUser{}.Mixin()
	orderuser.Policy = privacy.NewPolicies(orderuserMixin[0], schema.OrderUser{})
	orderuser.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := orderuser.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	orderuserMixinFields0 := orderuserMixin[0].Fields()
	_ = orderuserMixinFields0
	orderuserMixinFields1 := orderuserMixin[1].Fields()
	_ = orderuserMixinFields1
	orderuserFields := schema.OrderUser{}.Fields()
	_ = orderuserFields
	// orderuserDescCreatedAt is the schema descriptor for created_at field.
	orderuserDescCreatedAt := orderuserMixinFields0[0].Descriptor()
	// orderuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderuser.DefaultCreatedAt = orderuserDescCreatedAt.Default.(func() uint32)
	// orderuserDescUpdatedAt is the schema descriptor for updated_at field.
	orderuserDescUpdatedAt := orderuserMixinFields0[1].Descriptor()
	// orderuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderuser.DefaultUpdatedAt = orderuserDescUpdatedAt.Default.(func() uint32)
	// orderuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderuser.UpdateDefaultUpdatedAt = orderuserDescUpdatedAt.UpdateDefault.(func() uint32)
	// orderuserDescDeletedAt is the schema descriptor for deleted_at field.
	orderuserDescDeletedAt := orderuserMixinFields0[2].Descriptor()
	// orderuser.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	orderuser.DefaultDeletedAt = orderuserDescDeletedAt.Default.(func() uint32)
	// orderuserDescEntID is the schema descriptor for ent_id field.
	orderuserDescEntID := orderuserMixinFields1[1].Descriptor()
	// orderuser.DefaultEntID holds the default value on creation for the ent_id field.
	orderuser.DefaultEntID = orderuserDescEntID.Default.(func() uuid.UUID)
	// orderuserDescGoodUserID is the schema descriptor for good_user_id field.
	orderuserDescGoodUserID := orderuserFields[0].Descriptor()
	// orderuser.DefaultGoodUserID holds the default value on creation for the good_user_id field.
	orderuser.DefaultGoodUserID = orderuserDescGoodUserID.Default.(func() uuid.UUID)
	// orderuserDescUserID is the schema descriptor for user_id field.
	orderuserDescUserID := orderuserFields[1].Descriptor()
	// orderuser.DefaultUserID holds the default value on creation for the user_id field.
	orderuser.DefaultUserID = orderuserDescUserID.Default.(func() uuid.UUID)
	// orderuserDescAppID is the schema descriptor for app_id field.
	orderuserDescAppID := orderuserFields[2].Descriptor()
	// orderuser.DefaultAppID holds the default value on creation for the app_id field.
	orderuser.DefaultAppID = orderuserDescAppID.Default.(func() uuid.UUID)
	// orderuserDescName is the schema descriptor for name field.
	orderuserDescName := orderuserFields[3].Descriptor()
	// orderuser.DefaultName holds the default value on creation for the name field.
	orderuser.DefaultName = orderuserDescName.Default.(string)
	// orderuserDescProportion is the schema descriptor for proportion field.
	orderuserDescProportion := orderuserFields[4].Descriptor()
	// orderuser.DefaultProportion holds the default value on creation for the proportion field.
	orderuser.DefaultProportion = orderuserDescProportion.Default.(decimal.Decimal)
	// orderuserDescRevenueAddress is the schema descriptor for revenue_address field.
	orderuserDescRevenueAddress := orderuserFields[5].Descriptor()
	// orderuser.DefaultRevenueAddress holds the default value on creation for the revenue_address field.
	orderuser.DefaultRevenueAddress = orderuserDescRevenueAddress.Default.(string)
	// orderuserDescReadPageLink is the schema descriptor for read_page_link field.
	orderuserDescReadPageLink := orderuserFields[6].Descriptor()
	// orderuser.DefaultReadPageLink holds the default value on creation for the read_page_link field.
	orderuser.DefaultReadPageLink = orderuserDescReadPageLink.Default.(string)
	// orderuserDescAutoPay is the schema descriptor for auto_pay field.
	orderuserDescAutoPay := orderuserFields[7].Descriptor()
	// orderuser.DefaultAutoPay holds the default value on creation for the auto_pay field.
	orderuser.DefaultAutoPay = orderuserDescAutoPay.Default.(bool)
	poolMixin := schema.Pool{}.Mixin()
	pool.Policy = privacy.NewPolicies(poolMixin[0], schema.Pool{})
	pool.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := pool.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	poolMixinFields0 := poolMixin[0].Fields()
	_ = poolMixinFields0
	poolMixinFields1 := poolMixin[1].Fields()
	_ = poolMixinFields1
	poolFields := schema.Pool{}.Fields()
	_ = poolFields
	// poolDescCreatedAt is the schema descriptor for created_at field.
	poolDescCreatedAt := poolMixinFields0[0].Descriptor()
	// pool.DefaultCreatedAt holds the default value on creation for the created_at field.
	pool.DefaultCreatedAt = poolDescCreatedAt.Default.(func() uint32)
	// poolDescUpdatedAt is the schema descriptor for updated_at field.
	poolDescUpdatedAt := poolMixinFields0[1].Descriptor()
	// pool.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	pool.DefaultUpdatedAt = poolDescUpdatedAt.Default.(func() uint32)
	// pool.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	pool.UpdateDefaultUpdatedAt = poolDescUpdatedAt.UpdateDefault.(func() uint32)
	// poolDescDeletedAt is the schema descriptor for deleted_at field.
	poolDescDeletedAt := poolMixinFields0[2].Descriptor()
	// pool.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	pool.DefaultDeletedAt = poolDescDeletedAt.Default.(func() uint32)
	// poolDescEntID is the schema descriptor for ent_id field.
	poolDescEntID := poolMixinFields1[1].Descriptor()
	// pool.DefaultEntID holds the default value on creation for the ent_id field.
	pool.DefaultEntID = poolDescEntID.Default.(func() uuid.UUID)
	// poolDescMiningpoolType is the schema descriptor for miningpool_type field.
	poolDescMiningpoolType := poolFields[0].Descriptor()
	// pool.DefaultMiningpoolType holds the default value on creation for the miningpool_type field.
	pool.DefaultMiningpoolType = poolDescMiningpoolType.Default.(string)
	// poolDescName is the schema descriptor for name field.
	poolDescName := poolFields[1].Descriptor()
	// pool.DefaultName holds the default value on creation for the name field.
	pool.DefaultName = poolDescName.Default.(string)
	// poolDescSite is the schema descriptor for site field.
	poolDescSite := poolFields[2].Descriptor()
	// pool.DefaultSite holds the default value on creation for the site field.
	pool.DefaultSite = poolDescSite.Default.(string)
	// poolDescLogo is the schema descriptor for logo field.
	poolDescLogo := poolFields[3].Descriptor()
	// pool.DefaultLogo holds the default value on creation for the logo field.
	pool.DefaultLogo = poolDescLogo.Default.(string)
	// poolDescDescription is the schema descriptor for description field.
	poolDescDescription := poolFields[4].Descriptor()
	// pool.DefaultDescription holds the default value on creation for the description field.
	pool.DefaultDescription = poolDescDescription.Default.(string)
	rootuserMixin := schema.RootUser{}.Mixin()
	rootuser.Policy = privacy.NewPolicies(rootuserMixin[0], schema.RootUser{})
	rootuser.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := rootuser.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	rootuserMixinFields0 := rootuserMixin[0].Fields()
	_ = rootuserMixinFields0
	rootuserMixinFields1 := rootuserMixin[1].Fields()
	_ = rootuserMixinFields1
	rootuserFields := schema.RootUser{}.Fields()
	_ = rootuserFields
	// rootuserDescCreatedAt is the schema descriptor for created_at field.
	rootuserDescCreatedAt := rootuserMixinFields0[0].Descriptor()
	// rootuser.DefaultCreatedAt holds the default value on creation for the created_at field.
	rootuser.DefaultCreatedAt = rootuserDescCreatedAt.Default.(func() uint32)
	// rootuserDescUpdatedAt is the schema descriptor for updated_at field.
	rootuserDescUpdatedAt := rootuserMixinFields0[1].Descriptor()
	// rootuser.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	rootuser.DefaultUpdatedAt = rootuserDescUpdatedAt.Default.(func() uint32)
	// rootuser.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	rootuser.UpdateDefaultUpdatedAt = rootuserDescUpdatedAt.UpdateDefault.(func() uint32)
	// rootuserDescDeletedAt is the schema descriptor for deleted_at field.
	rootuserDescDeletedAt := rootuserMixinFields0[2].Descriptor()
	// rootuser.DefaultDeletedAt holds the default value on creation for the deleted_at field.
	rootuser.DefaultDeletedAt = rootuserDescDeletedAt.Default.(func() uint32)
	// rootuserDescEntID is the schema descriptor for ent_id field.
	rootuserDescEntID := rootuserMixinFields1[1].Descriptor()
	// rootuser.DefaultEntID holds the default value on creation for the ent_id field.
	rootuser.DefaultEntID = rootuserDescEntID.Default.(func() uuid.UUID)
	// rootuserDescName is the schema descriptor for name field.
	rootuserDescName := rootuserFields[0].Descriptor()
	// rootuser.DefaultName holds the default value on creation for the name field.
	rootuser.DefaultName = rootuserDescName.Default.(string)
	// rootuserDescPoolID is the schema descriptor for pool_id field.
	rootuserDescPoolID := rootuserFields[1].Descriptor()
	// rootuser.DefaultPoolID holds the default value on creation for the pool_id field.
	rootuser.DefaultPoolID = rootuserDescPoolID.Default.(func() uuid.UUID)
	// rootuserDescEmail is the schema descriptor for email field.
	rootuserDescEmail := rootuserFields[2].Descriptor()
	// rootuser.DefaultEmail holds the default value on creation for the email field.
	rootuser.DefaultEmail = rootuserDescEmail.Default.(string)
	// rootuserDescAuthToken is the schema descriptor for auth_token field.
	rootuserDescAuthToken := rootuserFields[3].Descriptor()
	// rootuser.DefaultAuthToken holds the default value on creation for the auth_token field.
	rootuser.DefaultAuthToken = rootuserDescAuthToken.Default.(string)
	// rootuserDescAuthTokenSalt is the schema descriptor for auth_token_salt field.
	rootuserDescAuthTokenSalt := rootuserFields[4].Descriptor()
	// rootuser.DefaultAuthTokenSalt holds the default value on creation for the auth_token_salt field.
	rootuser.DefaultAuthTokenSalt = rootuserDescAuthTokenSalt.Default.(string)
	// rootuserDescAuthed is the schema descriptor for authed field.
	rootuserDescAuthed := rootuserFields[5].Descriptor()
	// rootuser.DefaultAuthed holds the default value on creation for the authed field.
	rootuser.DefaultAuthed = rootuserDescAuthed.Default.(bool)
	// rootuserDescRemark is the schema descriptor for remark field.
	rootuserDescRemark := rootuserFields[6].Descriptor()
	// rootuser.DefaultRemark holds the default value on creation for the remark field.
	rootuser.DefaultRemark = rootuserDescRemark.Default.(string)
}

const (
	Version = "v0.11.3" // Version of ent codegen.
)
