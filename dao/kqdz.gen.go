// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"LGIS/model"
)

func newKqdz(db *gorm.DB) kqdz {
	_kqdz := kqdz{}

	_kqdz.kqdzDo.UseDB(db)
	_kqdz.kqdzDo.UseModel(&model.Kqdz{})

	tableName := _kqdz.kqdzDo.TableName()
	_kqdz.ALL = field.NewField(tableName, "*")
	_kqdz.KCAAA = field.NewString(tableName, "KCAAA")
	_kqdz.PKGKAA = field.NewString(tableName, "PKGKAA")
	_kqdz.DSB = field.NewString(tableName, "DSB")
	_kqdz.YSACAA = field.NewString(tableName, "YSACAA")
	_kqdz.YSACAB = field.NewString(tableName, "YSACAB")
	_kqdz.YSACC = field.NewString(tableName, "YSACC")
	_kqdz.GZAM = field.NewString(tableName, "GZAM")
	_kqdz.KCAJ = field.NewString(tableName, "KCAJ")
	_kqdz.DHBIA = field.NewString(tableName, "DHBIA")
	_kqdz.DHBHAB = field.NewFloat32(tableName, "DHBHAB")
	_kqdz.PKHFE = field.NewString(tableName, "PKHFE")
	_kqdz.PKHFB = field.NewString(tableName, "PKHFB")
	_kqdz.DHCAA = field.NewString(tableName, "DHCAA")
	_kqdz.DHC = field.NewString(tableName, "DHC")
	_kqdz.数据 = field.NewString(tableName, "数据")

	_kqdz.fillFieldMap()

	return _kqdz
}

type kqdz struct {
	kqdzDo kqdzDo

	ALL    field.Field
	KCAAA  field.String
	PKGKAA field.String
	DSB    field.String
	YSACAA field.String
	YSACAB field.String
	YSACC  field.String
	GZAM   field.String
	KCAJ   field.String
	DHBIA  field.String
	DHBHAB field.Float32
	PKHFE  field.String
	PKHFB  field.String
	DHCAA  field.String
	DHC    field.String
	数据     field.String

	fieldMap map[string]field.Expr
}

func (k kqdz) Table(newTableName string) *kqdz {
	k.kqdzDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k kqdz) As(alias string) *kqdz {
	k.kqdzDo.DO = *(k.kqdzDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *kqdz) updateTableName(table string) *kqdz {
	k.ALL = field.NewField(table, "*")
	k.KCAAA = field.NewString(table, "KCAAA")
	k.PKGKAA = field.NewString(table, "PKGKAA")
	k.DSB = field.NewString(table, "DSB")
	k.YSACAA = field.NewString(table, "YSACAA")
	k.YSACAB = field.NewString(table, "YSACAB")
	k.YSACC = field.NewString(table, "YSACC")
	k.GZAM = field.NewString(table, "GZAM")
	k.KCAJ = field.NewString(table, "KCAJ")
	k.DHBIA = field.NewString(table, "DHBIA")
	k.DHBHAB = field.NewFloat32(table, "DHBHAB")
	k.PKHFE = field.NewString(table, "PKHFE")
	k.PKHFB = field.NewString(table, "PKHFB")
	k.DHCAA = field.NewString(table, "DHCAA")
	k.DHC = field.NewString(table, "DHC")
	k.数据 = field.NewString(table, "数据")

	k.fillFieldMap()

	return k
}

func (k *kqdz) WithContext(ctx context.Context) *kqdzDo { return k.kqdzDo.WithContext(ctx) }

func (k kqdz) TableName() string { return k.kqdzDo.TableName() }

func (k kqdz) Alias() string { return k.kqdzDo.Alias() }

func (k *kqdz) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *kqdz) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 15)
	k.fieldMap["KCAAA"] = k.KCAAA
	k.fieldMap["PKGKAA"] = k.PKGKAA
	k.fieldMap["DSB"] = k.DSB
	k.fieldMap["YSACAA"] = k.YSACAA
	k.fieldMap["YSACAB"] = k.YSACAB
	k.fieldMap["YSACC"] = k.YSACC
	k.fieldMap["GZAM"] = k.GZAM
	k.fieldMap["KCAJ"] = k.KCAJ
	k.fieldMap["DHBIA"] = k.DHBIA
	k.fieldMap["DHBHAB"] = k.DHBHAB
	k.fieldMap["PKHFE"] = k.PKHFE
	k.fieldMap["PKHFB"] = k.PKHFB
	k.fieldMap["DHCAA"] = k.DHCAA
	k.fieldMap["DHC"] = k.DHC
	k.fieldMap["数据"] = k.数据
}

func (k kqdz) clone(db *gorm.DB) kqdz {
	k.kqdzDo.ReplaceDB(db)
	return k
}

type kqdzDo struct{ gen.DO }

func (k kqdzDo) Debug() *kqdzDo {
	return k.withDO(k.DO.Debug())
}

func (k kqdzDo) WithContext(ctx context.Context) *kqdzDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k kqdzDo) ReadDB(ctx context.Context) *kqdzDo {
	return k.WithContext(ctx).Clauses(dbresolver.Read)
}

func (k kqdzDo) WriteDB(ctx context.Context) *kqdzDo {
	return k.WithContext(ctx).Clauses(dbresolver.Write)
}

func (k kqdzDo) Clauses(conds ...clause.Expression) *kqdzDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k kqdzDo) Returning(value interface{}, columns ...string) *kqdzDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k kqdzDo) Not(conds ...gen.Condition) *kqdzDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k kqdzDo) Or(conds ...gen.Condition) *kqdzDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k kqdzDo) Select(conds ...field.Expr) *kqdzDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k kqdzDo) Where(conds ...gen.Condition) *kqdzDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k kqdzDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *kqdzDo {
	return k.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (k kqdzDo) Order(conds ...field.Expr) *kqdzDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k kqdzDo) Distinct(cols ...field.Expr) *kqdzDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k kqdzDo) Omit(cols ...field.Expr) *kqdzDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k kqdzDo) Join(table schema.Tabler, on ...field.Expr) *kqdzDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k kqdzDo) LeftJoin(table schema.Tabler, on ...field.Expr) *kqdzDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k kqdzDo) RightJoin(table schema.Tabler, on ...field.Expr) *kqdzDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k kqdzDo) Group(cols ...field.Expr) *kqdzDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k kqdzDo) Having(conds ...gen.Condition) *kqdzDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k kqdzDo) Limit(limit int) *kqdzDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k kqdzDo) Offset(offset int) *kqdzDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k kqdzDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *kqdzDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k kqdzDo) Unscoped() *kqdzDo {
	return k.withDO(k.DO.Unscoped())
}

func (k kqdzDo) Create(values ...*model.Kqdz) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k kqdzDo) CreateInBatches(values []*model.Kqdz, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k kqdzDo) Save(values ...*model.Kqdz) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k kqdzDo) First() (*model.Kqdz, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kqdz), nil
	}
}

func (k kqdzDo) Take() (*model.Kqdz, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kqdz), nil
	}
}

func (k kqdzDo) Last() (*model.Kqdz, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kqdz), nil
	}
}

func (k kqdzDo) Find() ([]*model.Kqdz, error) {
	result, err := k.DO.Find()
	return result.([]*model.Kqdz), err
}

func (k kqdzDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Kqdz, err error) {
	buf := make([]*model.Kqdz, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k kqdzDo) FindInBatches(result *[]*model.Kqdz, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k kqdzDo) Attrs(attrs ...field.AssignExpr) *kqdzDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k kqdzDo) Assign(attrs ...field.AssignExpr) *kqdzDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k kqdzDo) Joins(fields ...field.RelationField) *kqdzDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k kqdzDo) Preload(fields ...field.RelationField) *kqdzDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k kqdzDo) FirstOrInit() (*model.Kqdz, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kqdz), nil
	}
}

func (k kqdzDo) FirstOrCreate() (*model.Kqdz, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kqdz), nil
	}
}

func (k kqdzDo) FindByPage(offset int, limit int) (result []*model.Kqdz, count int64, err error) {
	result, err = k.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = k.Offset(-1).Limit(-1).Count()
	return
}

func (k kqdzDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k *kqdzDo) withDO(do gen.Dao) *kqdzDo {
	k.DO = *do.(*gen.DO)
	return k
}