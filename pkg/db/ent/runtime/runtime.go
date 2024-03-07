// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"context"

	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/schema"
	"github.com/google/uuid"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
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
	// coinDescSite is the schema descriptor for site field.
	coinDescSite := coinFields[1].Descriptor()
	// coin.DefaultSite holds the default value on creation for the site field.
	coin.DefaultSite = coinDescSite.Default.(string)
	// coinDescFeeRate is the schema descriptor for fee_rate field.
	coinDescFeeRate := coinFields[4].Descriptor()
	// coin.DefaultFeeRate holds the default value on creation for the fee_rate field.
	coin.DefaultFeeRate = coinDescFeeRate.Default.(float32)
	// coinDescRemark is the schema descriptor for remark field.
	coinDescRemark := coinFields[6].Descriptor()
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
	// gooduserDescHashRate is the schema descriptor for hash_rate field.
	gooduserDescHashRate := gooduserFields[6].Descriptor()
	// gooduser.DefaultHashRate holds the default value on creation for the hash_rate field.
	gooduser.DefaultHashRate = gooduserDescHashRate.Default.(float64)
	// gooduserDescReadPageLink is the schema descriptor for read_page_link field.
	gooduserDescReadPageLink := gooduserFields[7].Descriptor()
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
	// orderuserDescCompensationTime is the schema descriptor for compensation_time field.
	orderuserDescCompensationTime := orderuserFields[7].Descriptor()
	// orderuser.DefaultCompensationTime holds the default value on creation for the compensation_time field.
	orderuser.DefaultCompensationTime = orderuserDescCompensationTime.Default.(uint32)
	// orderuserDescRevenueAddress is the schema descriptor for revenue_address field.
	orderuserDescRevenueAddress := orderuserFields[8].Descriptor()
	// orderuser.DefaultRevenueAddress holds the default value on creation for the revenue_address field.
	orderuser.DefaultRevenueAddress = orderuserDescRevenueAddress.Default.(string)
	// orderuserDescReadPageLink is the schema descriptor for read_page_link field.
	orderuserDescReadPageLink := orderuserFields[10].Descriptor()
	// orderuser.DefaultReadPageLink holds the default value on creation for the read_page_link field.
	orderuser.DefaultReadPageLink = orderuserDescReadPageLink.Default.(string)
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
}

const (
	Version = "v0.11.3" // Version of ent codegen.
)
