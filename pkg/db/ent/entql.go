// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/apppool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/pool"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 8)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   apppool.Table,
			Columns: apppool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: apppool.FieldID,
			},
		},
		Type: "AppPool",
		Fields: map[string]*sqlgraph.FieldSpec{
			apppool.FieldCreatedAt: {Type: field.TypeUint32, Column: apppool.FieldCreatedAt},
			apppool.FieldUpdatedAt: {Type: field.TypeUint32, Column: apppool.FieldUpdatedAt},
			apppool.FieldDeletedAt: {Type: field.TypeUint32, Column: apppool.FieldDeletedAt},
			apppool.FieldEntID:     {Type: field.TypeUUID, Column: apppool.FieldEntID},
			apppool.FieldAppID:     {Type: field.TypeUUID, Column: apppool.FieldAppID},
			apppool.FieldPoolID:    {Type: field.TypeUUID, Column: apppool.FieldPoolID},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   coin.Table,
			Columns: coin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: coin.FieldID,
			},
		},
		Type: "Coin",
		Fields: map[string]*sqlgraph.FieldSpec{
			coin.FieldCreatedAt:              {Type: field.TypeUint32, Column: coin.FieldCreatedAt},
			coin.FieldUpdatedAt:              {Type: field.TypeUint32, Column: coin.FieldUpdatedAt},
			coin.FieldDeletedAt:              {Type: field.TypeUint32, Column: coin.FieldDeletedAt},
			coin.FieldEntID:                  {Type: field.TypeUUID, Column: coin.FieldEntID},
			coin.FieldPoolID:                 {Type: field.TypeUUID, Column: coin.FieldPoolID},
			coin.FieldCoinTypeID:             {Type: field.TypeUUID, Column: coin.FieldCoinTypeID},
			coin.FieldCoinType:               {Type: field.TypeString, Column: coin.FieldCoinType},
			coin.FieldRevenueType:            {Type: field.TypeString, Column: coin.FieldRevenueType},
			coin.FieldFeeRatio:               {Type: field.TypeOther, Column: coin.FieldFeeRatio},
			coin.FieldFixedRevenueAble:       {Type: field.TypeBool, Column: coin.FieldFixedRevenueAble},
			coin.FieldLeastTransferAmount:    {Type: field.TypeOther, Column: coin.FieldLeastTransferAmount},
			coin.FieldBenefitIntervalSeconds: {Type: field.TypeUint32, Column: coin.FieldBenefitIntervalSeconds},
			coin.FieldRemark:                 {Type: field.TypeString, Column: coin.FieldRemark},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   fraction.Table,
			Columns: fraction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fraction.FieldID,
			},
		},
		Type: "Fraction",
		Fields: map[string]*sqlgraph.FieldSpec{
			fraction.FieldCreatedAt:     {Type: field.TypeUint32, Column: fraction.FieldCreatedAt},
			fraction.FieldUpdatedAt:     {Type: field.TypeUint32, Column: fraction.FieldUpdatedAt},
			fraction.FieldDeletedAt:     {Type: field.TypeUint32, Column: fraction.FieldDeletedAt},
			fraction.FieldEntID:         {Type: field.TypeUUID, Column: fraction.FieldEntID},
			fraction.FieldAppID:         {Type: field.TypeUUID, Column: fraction.FieldAppID},
			fraction.FieldUserID:        {Type: field.TypeUUID, Column: fraction.FieldUserID},
			fraction.FieldOrderUserID:   {Type: field.TypeUUID, Column: fraction.FieldOrderUserID},
			fraction.FieldWithdrawState: {Type: field.TypeString, Column: fraction.FieldWithdrawState},
			fraction.FieldWithdrawAt:    {Type: field.TypeUint32, Column: fraction.FieldWithdrawAt},
			fraction.FieldPromisePayAt:  {Type: field.TypeUint32, Column: fraction.FieldPromisePayAt},
			fraction.FieldMsg:           {Type: field.TypeString, Column: fraction.FieldMsg},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   fractionrule.Table,
			Columns: fractionrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fractionrule.FieldID,
			},
		},
		Type: "FractionRule",
		Fields: map[string]*sqlgraph.FieldSpec{
			fractionrule.FieldCreatedAt:        {Type: field.TypeUint32, Column: fractionrule.FieldCreatedAt},
			fractionrule.FieldUpdatedAt:        {Type: field.TypeUint32, Column: fractionrule.FieldUpdatedAt},
			fractionrule.FieldDeletedAt:        {Type: field.TypeUint32, Column: fractionrule.FieldDeletedAt},
			fractionrule.FieldEntID:            {Type: field.TypeUUID, Column: fractionrule.FieldEntID},
			fractionrule.FieldPoolCoinTypeID:   {Type: field.TypeUUID, Column: fractionrule.FieldPoolCoinTypeID},
			fractionrule.FieldWithdrawInterval: {Type: field.TypeUint32, Column: fractionrule.FieldWithdrawInterval},
			fractionrule.FieldMinAmount:        {Type: field.TypeOther, Column: fractionrule.FieldMinAmount},
			fractionrule.FieldWithdrawRate:     {Type: field.TypeOther, Column: fractionrule.FieldWithdrawRate},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   gooduser.Table,
			Columns: gooduser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: gooduser.FieldID,
			},
		},
		Type: "GoodUser",
		Fields: map[string]*sqlgraph.FieldSpec{
			gooduser.FieldCreatedAt:      {Type: field.TypeUint32, Column: gooduser.FieldCreatedAt},
			gooduser.FieldUpdatedAt:      {Type: field.TypeUint32, Column: gooduser.FieldUpdatedAt},
			gooduser.FieldDeletedAt:      {Type: field.TypeUint32, Column: gooduser.FieldDeletedAt},
			gooduser.FieldEntID:          {Type: field.TypeUUID, Column: gooduser.FieldEntID},
			gooduser.FieldRootUserID:     {Type: field.TypeUUID, Column: gooduser.FieldRootUserID},
			gooduser.FieldName:           {Type: field.TypeString, Column: gooduser.FieldName},
			gooduser.FieldPoolCoinTypeID: {Type: field.TypeUUID, Column: gooduser.FieldPoolCoinTypeID},
			gooduser.FieldReadPageLink:   {Type: field.TypeString, Column: gooduser.FieldReadPageLink},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   orderuser.Table,
			Columns: orderuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: orderuser.FieldID,
			},
		},
		Type: "OrderUser",
		Fields: map[string]*sqlgraph.FieldSpec{
			orderuser.FieldCreatedAt:      {Type: field.TypeUint32, Column: orderuser.FieldCreatedAt},
			orderuser.FieldUpdatedAt:      {Type: field.TypeUint32, Column: orderuser.FieldUpdatedAt},
			orderuser.FieldDeletedAt:      {Type: field.TypeUint32, Column: orderuser.FieldDeletedAt},
			orderuser.FieldEntID:          {Type: field.TypeUUID, Column: orderuser.FieldEntID},
			orderuser.FieldGoodUserID:     {Type: field.TypeUUID, Column: orderuser.FieldGoodUserID},
			orderuser.FieldUserID:         {Type: field.TypeUUID, Column: orderuser.FieldUserID},
			orderuser.FieldAppID:          {Type: field.TypeUUID, Column: orderuser.FieldAppID},
			orderuser.FieldName:           {Type: field.TypeString, Column: orderuser.FieldName},
			orderuser.FieldProportion:     {Type: field.TypeOther, Column: orderuser.FieldProportion},
			orderuser.FieldRevenueAddress: {Type: field.TypeString, Column: orderuser.FieldRevenueAddress},
			orderuser.FieldReadPageLink:   {Type: field.TypeString, Column: orderuser.FieldReadPageLink},
			orderuser.FieldAutoPay:        {Type: field.TypeBool, Column: orderuser.FieldAutoPay},
		},
	}
	graph.Nodes[6] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   pool.Table,
			Columns: pool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: pool.FieldID,
			},
		},
		Type: "Pool",
		Fields: map[string]*sqlgraph.FieldSpec{
			pool.FieldCreatedAt:      {Type: field.TypeUint32, Column: pool.FieldCreatedAt},
			pool.FieldUpdatedAt:      {Type: field.TypeUint32, Column: pool.FieldUpdatedAt},
			pool.FieldDeletedAt:      {Type: field.TypeUint32, Column: pool.FieldDeletedAt},
			pool.FieldEntID:          {Type: field.TypeUUID, Column: pool.FieldEntID},
			pool.FieldMiningpoolType: {Type: field.TypeString, Column: pool.FieldMiningpoolType},
			pool.FieldName:           {Type: field.TypeString, Column: pool.FieldName},
			pool.FieldSite:           {Type: field.TypeString, Column: pool.FieldSite},
			pool.FieldLogo:           {Type: field.TypeString, Column: pool.FieldLogo},
			pool.FieldDescription:    {Type: field.TypeString, Column: pool.FieldDescription},
		},
	}
	graph.Nodes[7] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   rootuser.Table,
			Columns: rootuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: rootuser.FieldID,
			},
		},
		Type: "RootUser",
		Fields: map[string]*sqlgraph.FieldSpec{
			rootuser.FieldCreatedAt:     {Type: field.TypeUint32, Column: rootuser.FieldCreatedAt},
			rootuser.FieldUpdatedAt:     {Type: field.TypeUint32, Column: rootuser.FieldUpdatedAt},
			rootuser.FieldDeletedAt:     {Type: field.TypeUint32, Column: rootuser.FieldDeletedAt},
			rootuser.FieldEntID:         {Type: field.TypeUUID, Column: rootuser.FieldEntID},
			rootuser.FieldName:          {Type: field.TypeString, Column: rootuser.FieldName},
			rootuser.FieldPoolID:        {Type: field.TypeUUID, Column: rootuser.FieldPoolID},
			rootuser.FieldEmail:         {Type: field.TypeString, Column: rootuser.FieldEmail},
			rootuser.FieldAuthToken:     {Type: field.TypeString, Column: rootuser.FieldAuthToken},
			rootuser.FieldAuthTokenSalt: {Type: field.TypeString, Column: rootuser.FieldAuthTokenSalt},
			rootuser.FieldAuthed:        {Type: field.TypeBool, Column: rootuser.FieldAuthed},
			rootuser.FieldRemark:        {Type: field.TypeString, Column: rootuser.FieldRemark},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (apq *AppPoolQuery) addPredicate(pred func(s *sql.Selector)) {
	apq.predicates = append(apq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the AppPoolQuery builder.
func (apq *AppPoolQuery) Filter() *AppPoolFilter {
	return &AppPoolFilter{config: apq.config, predicateAdder: apq}
}

// addPredicate implements the predicateAdder interface.
func (m *AppPoolMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the AppPoolMutation builder.
func (m *AppPoolMutation) Filter() *AppPoolFilter {
	return &AppPoolFilter{config: m.config, predicateAdder: m}
}

// AppPoolFilter provides a generic filtering capability at runtime for AppPoolQuery.
type AppPoolFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *AppPoolFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *AppPoolFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(apppool.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *AppPoolFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(apppool.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *AppPoolFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(apppool.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *AppPoolFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(apppool.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *AppPoolFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(apppool.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *AppPoolFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(apppool.FieldAppID))
}

// WherePoolID applies the entql [16]byte predicate on the pool_id field.
func (f *AppPoolFilter) WherePoolID(p entql.ValueP) {
	f.Where(p.Field(apppool.FieldPoolID))
}

// addPredicate implements the predicateAdder interface.
func (cq *CoinQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CoinQuery builder.
func (cq *CoinQuery) Filter() *CoinFilter {
	return &CoinFilter{config: cq.config, predicateAdder: cq}
}

// addPredicate implements the predicateAdder interface.
func (m *CoinMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CoinMutation builder.
func (m *CoinMutation) Filter() *CoinFilter {
	return &CoinFilter{config: m.config, predicateAdder: m}
}

// CoinFilter provides a generic filtering capability at runtime for CoinQuery.
type CoinFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *CoinFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *CoinFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(coin.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CoinFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(coin.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CoinFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(coin.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CoinFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(coin.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *CoinFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(coin.FieldEntID))
}

// WherePoolID applies the entql [16]byte predicate on the pool_id field.
func (f *CoinFilter) WherePoolID(p entql.ValueP) {
	f.Where(p.Field(coin.FieldPoolID))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *CoinFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(coin.FieldCoinTypeID))
}

// WhereCoinType applies the entql string predicate on the coin_type field.
func (f *CoinFilter) WhereCoinType(p entql.StringP) {
	f.Where(p.Field(coin.FieldCoinType))
}

// WhereRevenueType applies the entql string predicate on the revenue_type field.
func (f *CoinFilter) WhereRevenueType(p entql.StringP) {
	f.Where(p.Field(coin.FieldRevenueType))
}

// WhereFeeRatio applies the entql other predicate on the fee_ratio field.
func (f *CoinFilter) WhereFeeRatio(p entql.OtherP) {
	f.Where(p.Field(coin.FieldFeeRatio))
}

// WhereFixedRevenueAble applies the entql bool predicate on the fixed_revenue_able field.
func (f *CoinFilter) WhereFixedRevenueAble(p entql.BoolP) {
	f.Where(p.Field(coin.FieldFixedRevenueAble))
}

// WhereLeastTransferAmount applies the entql other predicate on the least_transfer_amount field.
func (f *CoinFilter) WhereLeastTransferAmount(p entql.OtherP) {
	f.Where(p.Field(coin.FieldLeastTransferAmount))
}

// WhereBenefitIntervalSeconds applies the entql uint32 predicate on the benefit_interval_seconds field.
func (f *CoinFilter) WhereBenefitIntervalSeconds(p entql.Uint32P) {
	f.Where(p.Field(coin.FieldBenefitIntervalSeconds))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *CoinFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(coin.FieldRemark))
}

// addPredicate implements the predicateAdder interface.
func (fq *FractionQuery) addPredicate(pred func(s *sql.Selector)) {
	fq.predicates = append(fq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the FractionQuery builder.
func (fq *FractionQuery) Filter() *FractionFilter {
	return &FractionFilter{config: fq.config, predicateAdder: fq}
}

// addPredicate implements the predicateAdder interface.
func (m *FractionMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the FractionMutation builder.
func (m *FractionMutation) Filter() *FractionFilter {
	return &FractionFilter{config: m.config, predicateAdder: m}
}

// FractionFilter provides a generic filtering capability at runtime for FractionQuery.
type FractionFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *FractionFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *FractionFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(fraction.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *FractionFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(fraction.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *FractionFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(fraction.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *FractionFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(fraction.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *FractionFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(fraction.FieldEntID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *FractionFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(fraction.FieldAppID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *FractionFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(fraction.FieldUserID))
}

// WhereOrderUserID applies the entql [16]byte predicate on the order_user_id field.
func (f *FractionFilter) WhereOrderUserID(p entql.ValueP) {
	f.Where(p.Field(fraction.FieldOrderUserID))
}

// WhereWithdrawState applies the entql string predicate on the withdraw_state field.
func (f *FractionFilter) WhereWithdrawState(p entql.StringP) {
	f.Where(p.Field(fraction.FieldWithdrawState))
}

// WhereWithdrawAt applies the entql uint32 predicate on the withdraw_at field.
func (f *FractionFilter) WhereWithdrawAt(p entql.Uint32P) {
	f.Where(p.Field(fraction.FieldWithdrawAt))
}

// WherePromisePayAt applies the entql uint32 predicate on the promise_pay_at field.
func (f *FractionFilter) WherePromisePayAt(p entql.Uint32P) {
	f.Where(p.Field(fraction.FieldPromisePayAt))
}

// WhereMsg applies the entql string predicate on the msg field.
func (f *FractionFilter) WhereMsg(p entql.StringP) {
	f.Where(p.Field(fraction.FieldMsg))
}

// addPredicate implements the predicateAdder interface.
func (frq *FractionRuleQuery) addPredicate(pred func(s *sql.Selector)) {
	frq.predicates = append(frq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the FractionRuleQuery builder.
func (frq *FractionRuleQuery) Filter() *FractionRuleFilter {
	return &FractionRuleFilter{config: frq.config, predicateAdder: frq}
}

// addPredicate implements the predicateAdder interface.
func (m *FractionRuleMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the FractionRuleMutation builder.
func (m *FractionRuleMutation) Filter() *FractionRuleFilter {
	return &FractionRuleFilter{config: m.config, predicateAdder: m}
}

// FractionRuleFilter provides a generic filtering capability at runtime for FractionRuleQuery.
type FractionRuleFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *FractionRuleFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *FractionRuleFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(fractionrule.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *FractionRuleFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(fractionrule.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *FractionRuleFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(fractionrule.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *FractionRuleFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(fractionrule.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *FractionRuleFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(fractionrule.FieldEntID))
}

// WherePoolCoinTypeID applies the entql [16]byte predicate on the pool_coin_type_id field.
func (f *FractionRuleFilter) WherePoolCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(fractionrule.FieldPoolCoinTypeID))
}

// WhereWithdrawInterval applies the entql uint32 predicate on the withdraw_interval field.
func (f *FractionRuleFilter) WhereWithdrawInterval(p entql.Uint32P) {
	f.Where(p.Field(fractionrule.FieldWithdrawInterval))
}

// WhereMinAmount applies the entql other predicate on the min_amount field.
func (f *FractionRuleFilter) WhereMinAmount(p entql.OtherP) {
	f.Where(p.Field(fractionrule.FieldMinAmount))
}

// WhereWithdrawRate applies the entql other predicate on the withdraw_rate field.
func (f *FractionRuleFilter) WhereWithdrawRate(p entql.OtherP) {
	f.Where(p.Field(fractionrule.FieldWithdrawRate))
}

// addPredicate implements the predicateAdder interface.
func (guq *GoodUserQuery) addPredicate(pred func(s *sql.Selector)) {
	guq.predicates = append(guq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the GoodUserQuery builder.
func (guq *GoodUserQuery) Filter() *GoodUserFilter {
	return &GoodUserFilter{config: guq.config, predicateAdder: guq}
}

// addPredicate implements the predicateAdder interface.
func (m *GoodUserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the GoodUserMutation builder.
func (m *GoodUserMutation) Filter() *GoodUserFilter {
	return &GoodUserFilter{config: m.config, predicateAdder: m}
}

// GoodUserFilter provides a generic filtering capability at runtime for GoodUserQuery.
type GoodUserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *GoodUserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *GoodUserFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(gooduser.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *GoodUserFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(gooduser.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *GoodUserFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(gooduser.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *GoodUserFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(gooduser.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *GoodUserFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(gooduser.FieldEntID))
}

// WhereRootUserID applies the entql [16]byte predicate on the root_user_id field.
func (f *GoodUserFilter) WhereRootUserID(p entql.ValueP) {
	f.Where(p.Field(gooduser.FieldRootUserID))
}

// WhereName applies the entql string predicate on the name field.
func (f *GoodUserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(gooduser.FieldName))
}

// WherePoolCoinTypeID applies the entql [16]byte predicate on the pool_coin_type_id field.
func (f *GoodUserFilter) WherePoolCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(gooduser.FieldPoolCoinTypeID))
}

// WhereReadPageLink applies the entql string predicate on the read_page_link field.
func (f *GoodUserFilter) WhereReadPageLink(p entql.StringP) {
	f.Where(p.Field(gooduser.FieldReadPageLink))
}

// addPredicate implements the predicateAdder interface.
func (ouq *OrderUserQuery) addPredicate(pred func(s *sql.Selector)) {
	ouq.predicates = append(ouq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the OrderUserQuery builder.
func (ouq *OrderUserQuery) Filter() *OrderUserFilter {
	return &OrderUserFilter{config: ouq.config, predicateAdder: ouq}
}

// addPredicate implements the predicateAdder interface.
func (m *OrderUserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the OrderUserMutation builder.
func (m *OrderUserMutation) Filter() *OrderUserFilter {
	return &OrderUserFilter{config: m.config, predicateAdder: m}
}

// OrderUserFilter provides a generic filtering capability at runtime for OrderUserQuery.
type OrderUserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *OrderUserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *OrderUserFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(orderuser.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *OrderUserFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(orderuser.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *OrderUserFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(orderuser.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *OrderUserFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(orderuser.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *OrderUserFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(orderuser.FieldEntID))
}

// WhereGoodUserID applies the entql [16]byte predicate on the good_user_id field.
func (f *OrderUserFilter) WhereGoodUserID(p entql.ValueP) {
	f.Where(p.Field(orderuser.FieldGoodUserID))
}

// WhereUserID applies the entql [16]byte predicate on the user_id field.
func (f *OrderUserFilter) WhereUserID(p entql.ValueP) {
	f.Where(p.Field(orderuser.FieldUserID))
}

// WhereAppID applies the entql [16]byte predicate on the app_id field.
func (f *OrderUserFilter) WhereAppID(p entql.ValueP) {
	f.Where(p.Field(orderuser.FieldAppID))
}

// WhereName applies the entql string predicate on the name field.
func (f *OrderUserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(orderuser.FieldName))
}

// WhereProportion applies the entql other predicate on the proportion field.
func (f *OrderUserFilter) WhereProportion(p entql.OtherP) {
	f.Where(p.Field(orderuser.FieldProportion))
}

// WhereRevenueAddress applies the entql string predicate on the revenue_address field.
func (f *OrderUserFilter) WhereRevenueAddress(p entql.StringP) {
	f.Where(p.Field(orderuser.FieldRevenueAddress))
}

// WhereReadPageLink applies the entql string predicate on the read_page_link field.
func (f *OrderUserFilter) WhereReadPageLink(p entql.StringP) {
	f.Where(p.Field(orderuser.FieldReadPageLink))
}

// WhereAutoPay applies the entql bool predicate on the auto_pay field.
func (f *OrderUserFilter) WhereAutoPay(p entql.BoolP) {
	f.Where(p.Field(orderuser.FieldAutoPay))
}

// addPredicate implements the predicateAdder interface.
func (pq *PoolQuery) addPredicate(pred func(s *sql.Selector)) {
	pq.predicates = append(pq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the PoolQuery builder.
func (pq *PoolQuery) Filter() *PoolFilter {
	return &PoolFilter{config: pq.config, predicateAdder: pq}
}

// addPredicate implements the predicateAdder interface.
func (m *PoolMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the PoolMutation builder.
func (m *PoolMutation) Filter() *PoolFilter {
	return &PoolFilter{config: m.config, predicateAdder: m}
}

// PoolFilter provides a generic filtering capability at runtime for PoolQuery.
type PoolFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *PoolFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[6].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *PoolFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(pool.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *PoolFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(pool.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *PoolFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(pool.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *PoolFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(pool.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *PoolFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(pool.FieldEntID))
}

// WhereMiningpoolType applies the entql string predicate on the miningpool_type field.
func (f *PoolFilter) WhereMiningpoolType(p entql.StringP) {
	f.Where(p.Field(pool.FieldMiningpoolType))
}

// WhereName applies the entql string predicate on the name field.
func (f *PoolFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(pool.FieldName))
}

// WhereSite applies the entql string predicate on the site field.
func (f *PoolFilter) WhereSite(p entql.StringP) {
	f.Where(p.Field(pool.FieldSite))
}

// WhereLogo applies the entql string predicate on the logo field.
func (f *PoolFilter) WhereLogo(p entql.StringP) {
	f.Where(p.Field(pool.FieldLogo))
}

// WhereDescription applies the entql string predicate on the description field.
func (f *PoolFilter) WhereDescription(p entql.StringP) {
	f.Where(p.Field(pool.FieldDescription))
}

// addPredicate implements the predicateAdder interface.
func (ruq *RootUserQuery) addPredicate(pred func(s *sql.Selector)) {
	ruq.predicates = append(ruq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the RootUserQuery builder.
func (ruq *RootUserQuery) Filter() *RootUserFilter {
	return &RootUserFilter{config: ruq.config, predicateAdder: ruq}
}

// addPredicate implements the predicateAdder interface.
func (m *RootUserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the RootUserMutation builder.
func (m *RootUserMutation) Filter() *RootUserFilter {
	return &RootUserFilter{config: m.config, predicateAdder: m}
}

// RootUserFilter provides a generic filtering capability at runtime for RootUserQuery.
type RootUserFilter struct {
	predicateAdder
	config
}

// Where applies the entql predicate on the query filter.
func (f *RootUserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[7].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql uint32 predicate on the id field.
func (f *RootUserFilter) WhereID(p entql.Uint32P) {
	f.Where(p.Field(rootuser.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *RootUserFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(rootuser.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *RootUserFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(rootuser.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *RootUserFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(rootuser.FieldDeletedAt))
}

// WhereEntID applies the entql [16]byte predicate on the ent_id field.
func (f *RootUserFilter) WhereEntID(p entql.ValueP) {
	f.Where(p.Field(rootuser.FieldEntID))
}

// WhereName applies the entql string predicate on the name field.
func (f *RootUserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(rootuser.FieldName))
}

// WherePoolID applies the entql [16]byte predicate on the pool_id field.
func (f *RootUserFilter) WherePoolID(p entql.ValueP) {
	f.Where(p.Field(rootuser.FieldPoolID))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *RootUserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(rootuser.FieldEmail))
}

// WhereAuthToken applies the entql string predicate on the auth_token field.
func (f *RootUserFilter) WhereAuthToken(p entql.StringP) {
	f.Where(p.Field(rootuser.FieldAuthToken))
}

// WhereAuthTokenSalt applies the entql string predicate on the auth_token_salt field.
func (f *RootUserFilter) WhereAuthTokenSalt(p entql.StringP) {
	f.Where(p.Field(rootuser.FieldAuthTokenSalt))
}

// WhereAuthed applies the entql bool predicate on the authed field.
func (f *RootUserFilter) WhereAuthed(p entql.BoolP) {
	f.Where(p.Field(rootuser.FieldAuthed))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *RootUserFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(rootuser.FieldRemark))
}
