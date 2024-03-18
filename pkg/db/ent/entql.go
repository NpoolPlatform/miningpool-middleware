// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/coin"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fraction"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/fractionrule"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/gooduser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/orderuser"
	"github.com/NpoolPlatform/miningpool-middleware/pkg/db/ent/rootuser"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 6)}
	graph.Nodes[0] = &sqlgraph.Node{
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
			coin.FieldCreatedAt:        {Type: field.TypeUint32, Column: coin.FieldCreatedAt},
			coin.FieldUpdatedAt:        {Type: field.TypeUint32, Column: coin.FieldUpdatedAt},
			coin.FieldDeletedAt:        {Type: field.TypeUint32, Column: coin.FieldDeletedAt},
			coin.FieldEntID:            {Type: field.TypeUUID, Column: coin.FieldEntID},
			coin.FieldMiningpoolType:   {Type: field.TypeString, Column: coin.FieldMiningpoolType},
			coin.FieldSite:             {Type: field.TypeString, Column: coin.FieldSite},
			coin.FieldCoinType:         {Type: field.TypeString, Column: coin.FieldCoinType},
			coin.FieldRevenueTypes:     {Type: field.TypeJSON, Column: coin.FieldRevenueTypes},
			coin.FieldFeeRate:          {Type: field.TypeFloat32, Column: coin.FieldFeeRate},
			coin.FieldFixedRevenueAble: {Type: field.TypeBool, Column: coin.FieldFixedRevenueAble},
			coin.FieldRemark:           {Type: field.TypeString, Column: coin.FieldRemark},
			coin.FieldThreshold:        {Type: field.TypeFloat32, Column: coin.FieldThreshold},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
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
			fraction.FieldOrderUserID:   {Type: field.TypeUUID, Column: fraction.FieldOrderUserID},
			fraction.FieldWithdrawState: {Type: field.TypeString, Column: fraction.FieldWithdrawState},
			fraction.FieldWithdrawTime:  {Type: field.TypeUint32, Column: fraction.FieldWithdrawTime},
			fraction.FieldPayTime:       {Type: field.TypeUint32, Column: fraction.FieldPayTime},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
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
			fractionrule.FieldMiningpoolType:   {Type: field.TypeString, Column: fractionrule.FieldMiningpoolType},
			fractionrule.FieldCoinType:         {Type: field.TypeString, Column: fractionrule.FieldCoinType},
			fractionrule.FieldWithdrawInterval: {Type: field.TypeUint32, Column: fractionrule.FieldWithdrawInterval},
			fractionrule.FieldMinAmount:        {Type: field.TypeFloat32, Column: fractionrule.FieldMinAmount},
			fractionrule.FieldWithdrawRate:     {Type: field.TypeFloat32, Column: fractionrule.FieldWithdrawRate},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
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
			gooduser.FieldGoodID:         {Type: field.TypeUUID, Column: gooduser.FieldGoodID},
			gooduser.FieldRootUserID:     {Type: field.TypeUUID, Column: gooduser.FieldRootUserID},
			gooduser.FieldName:           {Type: field.TypeString, Column: gooduser.FieldName},
			gooduser.FieldMiningpoolType: {Type: field.TypeString, Column: gooduser.FieldMiningpoolType},
			gooduser.FieldCoinType:       {Type: field.TypeString, Column: gooduser.FieldCoinType},
			gooduser.FieldHashRate:       {Type: field.TypeFloat32, Column: gooduser.FieldHashRate},
			gooduser.FieldReadPageLink:   {Type: field.TypeString, Column: gooduser.FieldReadPageLink},
			gooduser.FieldRevenueType:    {Type: field.TypeString, Column: gooduser.FieldRevenueType},
		},
	}
	graph.Nodes[4] = &sqlgraph.Node{
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
			orderuser.FieldRootUserID:     {Type: field.TypeUUID, Column: orderuser.FieldRootUserID},
			orderuser.FieldGoodUserID:     {Type: field.TypeUUID, Column: orderuser.FieldGoodUserID},
			orderuser.FieldOrderID:        {Type: field.TypeUUID, Column: orderuser.FieldOrderID},
			orderuser.FieldName:           {Type: field.TypeString, Column: orderuser.FieldName},
			orderuser.FieldMiningpoolType: {Type: field.TypeString, Column: orderuser.FieldMiningpoolType},
			orderuser.FieldCoinType:       {Type: field.TypeString, Column: orderuser.FieldCoinType},
			orderuser.FieldProportion:     {Type: field.TypeFloat32, Column: orderuser.FieldProportion},
			orderuser.FieldRevenueAddress: {Type: field.TypeString, Column: orderuser.FieldRevenueAddress},
			orderuser.FieldReadPageLink:   {Type: field.TypeString, Column: orderuser.FieldReadPageLink},
			orderuser.FieldAutoPay:        {Type: field.TypeBool, Column: orderuser.FieldAutoPay},
		},
	}
	graph.Nodes[5] = &sqlgraph.Node{
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
			rootuser.FieldCreatedAt:      {Type: field.TypeUint32, Column: rootuser.FieldCreatedAt},
			rootuser.FieldUpdatedAt:      {Type: field.TypeUint32, Column: rootuser.FieldUpdatedAt},
			rootuser.FieldDeletedAt:      {Type: field.TypeUint32, Column: rootuser.FieldDeletedAt},
			rootuser.FieldEntID:          {Type: field.TypeUUID, Column: rootuser.FieldEntID},
			rootuser.FieldName:           {Type: field.TypeString, Column: rootuser.FieldName},
			rootuser.FieldMiningpoolType: {Type: field.TypeString, Column: rootuser.FieldMiningpoolType},
			rootuser.FieldEmail:          {Type: field.TypeString, Column: rootuser.FieldEmail},
			rootuser.FieldAuthToken:      {Type: field.TypeString, Column: rootuser.FieldAuthToken},
			rootuser.FieldAuthed:         {Type: field.TypeBool, Column: rootuser.FieldAuthed},
			rootuser.FieldRemark:         {Type: field.TypeString, Column: rootuser.FieldRemark},
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
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
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

// WhereMiningpoolType applies the entql string predicate on the miningpool_type field.
func (f *CoinFilter) WhereMiningpoolType(p entql.StringP) {
	f.Where(p.Field(coin.FieldMiningpoolType))
}

// WhereSite applies the entql string predicate on the site field.
func (f *CoinFilter) WhereSite(p entql.StringP) {
	f.Where(p.Field(coin.FieldSite))
}

// WhereCoinType applies the entql string predicate on the coin_type field.
func (f *CoinFilter) WhereCoinType(p entql.StringP) {
	f.Where(p.Field(coin.FieldCoinType))
}

// WhereRevenueTypes applies the entql json.RawMessage predicate on the revenue_types field.
func (f *CoinFilter) WhereRevenueTypes(p entql.BytesP) {
	f.Where(p.Field(coin.FieldRevenueTypes))
}

// WhereFeeRate applies the entql float32 predicate on the fee_rate field.
func (f *CoinFilter) WhereFeeRate(p entql.Float32P) {
	f.Where(p.Field(coin.FieldFeeRate))
}

// WhereFixedRevenueAble applies the entql bool predicate on the fixed_revenue_able field.
func (f *CoinFilter) WhereFixedRevenueAble(p entql.BoolP) {
	f.Where(p.Field(coin.FieldFixedRevenueAble))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *CoinFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(coin.FieldRemark))
}

// WhereThreshold applies the entql float32 predicate on the threshold field.
func (f *CoinFilter) WhereThreshold(p entql.Float32P) {
	f.Where(p.Field(coin.FieldThreshold))
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
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
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

// WhereOrderUserID applies the entql [16]byte predicate on the order_user_id field.
func (f *FractionFilter) WhereOrderUserID(p entql.ValueP) {
	f.Where(p.Field(fraction.FieldOrderUserID))
}

// WhereWithdrawState applies the entql string predicate on the withdraw_state field.
func (f *FractionFilter) WhereWithdrawState(p entql.StringP) {
	f.Where(p.Field(fraction.FieldWithdrawState))
}

// WhereWithdrawTime applies the entql uint32 predicate on the withdraw_time field.
func (f *FractionFilter) WhereWithdrawTime(p entql.Uint32P) {
	f.Where(p.Field(fraction.FieldWithdrawTime))
}

// WherePayTime applies the entql uint32 predicate on the pay_time field.
func (f *FractionFilter) WherePayTime(p entql.Uint32P) {
	f.Where(p.Field(fraction.FieldPayTime))
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
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
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

// WhereMiningpoolType applies the entql string predicate on the miningpool_type field.
func (f *FractionRuleFilter) WhereMiningpoolType(p entql.StringP) {
	f.Where(p.Field(fractionrule.FieldMiningpoolType))
}

// WhereCoinType applies the entql string predicate on the coin_type field.
func (f *FractionRuleFilter) WhereCoinType(p entql.StringP) {
	f.Where(p.Field(fractionrule.FieldCoinType))
}

// WhereWithdrawInterval applies the entql uint32 predicate on the withdraw_interval field.
func (f *FractionRuleFilter) WhereWithdrawInterval(p entql.Uint32P) {
	f.Where(p.Field(fractionrule.FieldWithdrawInterval))
}

// WhereMinAmount applies the entql float32 predicate on the min_amount field.
func (f *FractionRuleFilter) WhereMinAmount(p entql.Float32P) {
	f.Where(p.Field(fractionrule.FieldMinAmount))
}

// WhereWithdrawRate applies the entql float32 predicate on the withdraw_rate field.
func (f *FractionRuleFilter) WhereWithdrawRate(p entql.Float32P) {
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
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
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

// WhereGoodID applies the entql [16]byte predicate on the good_id field.
func (f *GoodUserFilter) WhereGoodID(p entql.ValueP) {
	f.Where(p.Field(gooduser.FieldGoodID))
}

// WhereRootUserID applies the entql [16]byte predicate on the root_user_id field.
func (f *GoodUserFilter) WhereRootUserID(p entql.ValueP) {
	f.Where(p.Field(gooduser.FieldRootUserID))
}

// WhereName applies the entql string predicate on the name field.
func (f *GoodUserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(gooduser.FieldName))
}

// WhereMiningpoolType applies the entql string predicate on the miningpool_type field.
func (f *GoodUserFilter) WhereMiningpoolType(p entql.StringP) {
	f.Where(p.Field(gooduser.FieldMiningpoolType))
}

// WhereCoinType applies the entql string predicate on the coin_type field.
func (f *GoodUserFilter) WhereCoinType(p entql.StringP) {
	f.Where(p.Field(gooduser.FieldCoinType))
}

// WhereHashRate applies the entql float32 predicate on the hash_rate field.
func (f *GoodUserFilter) WhereHashRate(p entql.Float32P) {
	f.Where(p.Field(gooduser.FieldHashRate))
}

// WhereReadPageLink applies the entql string predicate on the read_page_link field.
func (f *GoodUserFilter) WhereReadPageLink(p entql.StringP) {
	f.Where(p.Field(gooduser.FieldReadPageLink))
}

// WhereRevenueType applies the entql string predicate on the revenue_type field.
func (f *GoodUserFilter) WhereRevenueType(p entql.StringP) {
	f.Where(p.Field(gooduser.FieldRevenueType))
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
		if err := schemaGraph.EvalP(schemaGraph.Nodes[4].Type, p, s); err != nil {
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

// WhereRootUserID applies the entql [16]byte predicate on the root_user_id field.
func (f *OrderUserFilter) WhereRootUserID(p entql.ValueP) {
	f.Where(p.Field(orderuser.FieldRootUserID))
}

// WhereGoodUserID applies the entql [16]byte predicate on the good_user_id field.
func (f *OrderUserFilter) WhereGoodUserID(p entql.ValueP) {
	f.Where(p.Field(orderuser.FieldGoodUserID))
}

// WhereOrderID applies the entql [16]byte predicate on the order_id field.
func (f *OrderUserFilter) WhereOrderID(p entql.ValueP) {
	f.Where(p.Field(orderuser.FieldOrderID))
}

// WhereName applies the entql string predicate on the name field.
func (f *OrderUserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(orderuser.FieldName))
}

// WhereMiningpoolType applies the entql string predicate on the miningpool_type field.
func (f *OrderUserFilter) WhereMiningpoolType(p entql.StringP) {
	f.Where(p.Field(orderuser.FieldMiningpoolType))
}

// WhereCoinType applies the entql string predicate on the coin_type field.
func (f *OrderUserFilter) WhereCoinType(p entql.StringP) {
	f.Where(p.Field(orderuser.FieldCoinType))
}

// WhereProportion applies the entql float32 predicate on the proportion field.
func (f *OrderUserFilter) WhereProportion(p entql.Float32P) {
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
		if err := schemaGraph.EvalP(schemaGraph.Nodes[5].Type, p, s); err != nil {
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

// WhereMiningpoolType applies the entql string predicate on the miningpool_type field.
func (f *RootUserFilter) WhereMiningpoolType(p entql.StringP) {
	f.Where(p.Field(rootuser.FieldMiningpoolType))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *RootUserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(rootuser.FieldEmail))
}

// WhereAuthToken applies the entql string predicate on the auth_token field.
func (f *RootUserFilter) WhereAuthToken(p entql.StringP) {
	f.Where(p.Field(rootuser.FieldAuthToken))
}

// WhereAuthed applies the entql bool predicate on the authed field.
func (f *RootUserFilter) WhereAuthed(p entql.BoolP) {
	f.Where(p.Field(rootuser.FieldAuthed))
}

// WhereRemark applies the entql string predicate on the remark field.
func (f *RootUserFilter) WhereRemark(p entql.StringP) {
	f.Where(p.Field(rootuser.FieldRemark))
}
