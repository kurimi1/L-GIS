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

func newKttz(db *gorm.DB) kttz {
	_kttz := kttz{}

	_kttz.kttzDo.UseDB(db)
	_kttz.kttzDo.UseModel(&model.Kttz{})

	tableName := _kttz.kttzDo.TableName()
	_kttz.ALL = field.NewField(tableName, "*")
	_kttz.KCAAA = field.NewString(tableName, "KCAAA")
	_kttz.PKGKI = field.NewInt32(tableName, "PKGKI")
	_kttz.PKGKIA = field.NewInt32(tableName, "PKGKIA")
	_kttz.PKGKDA = field.NewString(tableName, "PKGKDA")
	_kttz.PKGKJ = field.NewString(tableName, "PKGKJ")
	_kttz.PKGKGA = field.NewString(tableName, "PKGKGA")
	_kttz.PKGKGB = field.NewString(tableName, "PKGKGB")
	_kttz.PKGKGC = field.NewString(tableName, "PKGKGC")
	_kttz.PKGKGF = field.NewString(tableName, "PKGKGF")
	_kttz.PKGKFA = field.NewFloat32(tableName, "PKGKFA")
	_kttz.PKGKFB = field.NewFloat32(tableName, "PKGKFB")
	_kttz.PKCDA = field.NewFloat32(tableName, "PKCDA")
	_kttz.PKGADH = field.NewString(tableName, "PKGADH")
	_kttz.PKGKQ = field.NewString(tableName, "PKGKQ")
	_kttz.KCAPA = field.NewString(tableName, "KCAPA")
	_kttz.KCAPB = field.NewString(tableName, "KCAPB")
	_kttz.KWBH = field.NewString(tableName, "KWBH")
	_kttz.PKCDD = field.NewString(tableName, "PKCDD")
	_kttz.PKGKU = field.NewString(tableName, "PKGKU")
	_kttz.PKGKPH = field.NewString(tableName, "PKGKPH")
	_kttz.PKGKTA = field.NewString(tableName, "PKGKTA")
	_kttz.PKGKT = field.NewString(tableName, "PKGKT")
	_kttz.数据 = field.NewString(tableName, "数据")

	_kttz.fillFieldMap()

	return _kttz
}

type kttz struct {
	kttzDo kttzDo

	ALL    field.Field
	KCAAA  field.String
	PKGKI  field.Int32
	PKGKIA field.Int32
	PKGKDA field.String
	PKGKJ  field.String
	PKGKGA field.String
	PKGKGB field.String
	PKGKGC field.String
	PKGKGF field.String
	PKGKFA field.Float32
	PKGKFB field.Float32
	PKCDA  field.Float32
	PKGADH field.String
	PKGKQ  field.String
	KCAPA  field.String
	KCAPB  field.String
	KWBH   field.String
	PKCDD  field.String
	PKGKU  field.String
	PKGKPH field.String
	PKGKTA field.String
	PKGKT  field.String
	数据     field.String

	fieldMap map[string]field.Expr
}

func (k kttz) Table(newTableName string) *kttz {
	k.kttzDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k kttz) As(alias string) *kttz {
	k.kttzDo.DO = *(k.kttzDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *kttz) updateTableName(table string) *kttz {
	k.ALL = field.NewField(table, "*")
	k.KCAAA = field.NewString(table, "KCAAA")
	k.PKGKI = field.NewInt32(table, "PKGKI")
	k.PKGKIA = field.NewInt32(table, "PKGKIA")
	k.PKGKDA = field.NewString(table, "PKGKDA")
	k.PKGKJ = field.NewString(table, "PKGKJ")
	k.PKGKGA = field.NewString(table, "PKGKGA")
	k.PKGKGB = field.NewString(table, "PKGKGB")
	k.PKGKGC = field.NewString(table, "PKGKGC")
	k.PKGKGF = field.NewString(table, "PKGKGF")
	k.PKGKFA = field.NewFloat32(table, "PKGKFA")
	k.PKGKFB = field.NewFloat32(table, "PKGKFB")
	k.PKCDA = field.NewFloat32(table, "PKCDA")
	k.PKGADH = field.NewString(table, "PKGADH")
	k.PKGKQ = field.NewString(table, "PKGKQ")
	k.KCAPA = field.NewString(table, "KCAPA")
	k.KCAPB = field.NewString(table, "KCAPB")
	k.KWBH = field.NewString(table, "KWBH")
	k.PKCDD = field.NewString(table, "PKCDD")
	k.PKGKU = field.NewString(table, "PKGKU")
	k.PKGKPH = field.NewString(table, "PKGKPH")
	k.PKGKTA = field.NewString(table, "PKGKTA")
	k.PKGKT = field.NewString(table, "PKGKT")
	k.数据 = field.NewString(table, "数据")

	k.fillFieldMap()

	return k
}

func (k *kttz) WithContext(ctx context.Context) *kttzDo { return k.kttzDo.WithContext(ctx) }

func (k kttz) TableName() string { return k.kttzDo.TableName() }

func (k kttz) Alias() string { return k.kttzDo.Alias() }

func (k *kttz) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *kttz) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 23)
	k.fieldMap["KCAAA"] = k.KCAAA
	k.fieldMap["PKGKI"] = k.PKGKI
	k.fieldMap["PKGKIA"] = k.PKGKIA
	k.fieldMap["PKGKDA"] = k.PKGKDA
	k.fieldMap["PKGKJ"] = k.PKGKJ
	k.fieldMap["PKGKGA"] = k.PKGKGA
	k.fieldMap["PKGKGB"] = k.PKGKGB
	k.fieldMap["PKGKGC"] = k.PKGKGC
	k.fieldMap["PKGKGF"] = k.PKGKGF
	k.fieldMap["PKGKFA"] = k.PKGKFA
	k.fieldMap["PKGKFB"] = k.PKGKFB
	k.fieldMap["PKCDA"] = k.PKCDA
	k.fieldMap["PKGADH"] = k.PKGADH
	k.fieldMap["PKGKQ"] = k.PKGKQ
	k.fieldMap["KCAPA"] = k.KCAPA
	k.fieldMap["KCAPB"] = k.KCAPB
	k.fieldMap["KWBH"] = k.KWBH
	k.fieldMap["PKCDD"] = k.PKCDD
	k.fieldMap["PKGKU"] = k.PKGKU
	k.fieldMap["PKGKPH"] = k.PKGKPH
	k.fieldMap["PKGKTA"] = k.PKGKTA
	k.fieldMap["PKGKT"] = k.PKGKT
	k.fieldMap["数据"] = k.数据
}

func (k kttz) clone(db *gorm.DB) kttz {
	k.kttzDo.ReplaceDB(db)
	return k
}

type kttzDo struct{ gen.DO }

func (k kttzDo) Debug() *kttzDo {
	return k.withDO(k.DO.Debug())
}

func (k kttzDo) WithContext(ctx context.Context) *kttzDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k kttzDo) ReadDB(ctx context.Context) *kttzDo {
	return k.WithContext(ctx).Clauses(dbresolver.Read)
}

func (k kttzDo) WriteDB(ctx context.Context) *kttzDo {
	return k.WithContext(ctx).Clauses(dbresolver.Write)
}

func (k kttzDo) Clauses(conds ...clause.Expression) *kttzDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k kttzDo) Returning(value interface{}, columns ...string) *kttzDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k kttzDo) Not(conds ...gen.Condition) *kttzDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k kttzDo) Or(conds ...gen.Condition) *kttzDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k kttzDo) Select(conds ...field.Expr) *kttzDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k kttzDo) Where(conds ...gen.Condition) *kttzDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k kttzDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *kttzDo {
	return k.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (k kttzDo) Order(conds ...field.Expr) *kttzDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k kttzDo) Distinct(cols ...field.Expr) *kttzDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k kttzDo) Omit(cols ...field.Expr) *kttzDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k kttzDo) Join(table schema.Tabler, on ...field.Expr) *kttzDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k kttzDo) LeftJoin(table schema.Tabler, on ...field.Expr) *kttzDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k kttzDo) RightJoin(table schema.Tabler, on ...field.Expr) *kttzDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k kttzDo) Group(cols ...field.Expr) *kttzDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k kttzDo) Having(conds ...gen.Condition) *kttzDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k kttzDo) Limit(limit int) *kttzDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k kttzDo) Offset(offset int) *kttzDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k kttzDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *kttzDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k kttzDo) Unscoped() *kttzDo {
	return k.withDO(k.DO.Unscoped())
}

func (k kttzDo) Create(values ...*model.Kttz) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k kttzDo) CreateInBatches(values []*model.Kttz, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k kttzDo) Save(values ...*model.Kttz) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k kttzDo) First() (*model.Kttz, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kttz), nil
	}
}

func (k kttzDo) Take() (*model.Kttz, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kttz), nil
	}
}

func (k kttzDo) Last() (*model.Kttz, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kttz), nil
	}
}

func (k kttzDo) Find() ([]*model.Kttz, error) {
	result, err := k.DO.Find()
	return result.([]*model.Kttz), err
}

func (k kttzDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Kttz, err error) {
	buf := make([]*model.Kttz, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k kttzDo) FindInBatches(result *[]*model.Kttz, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k kttzDo) Attrs(attrs ...field.AssignExpr) *kttzDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k kttzDo) Assign(attrs ...field.AssignExpr) *kttzDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k kttzDo) Joins(fields ...field.RelationField) *kttzDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k kttzDo) Preload(fields ...field.RelationField) *kttzDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k kttzDo) FirstOrInit() (*model.Kttz, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kttz), nil
	}
}

func (k kttzDo) FirstOrCreate() (*model.Kttz, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kttz), nil
	}
}

func (k kttzDo) FindByPage(offset int, limit int) (result []*model.Kttz, count int64, err error) {
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

func (k kttzDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k *kttzDo) withDO(do gen.Dao) *kttzDo {
	k.DO = *do.(*gen.DO)
	return k
}
