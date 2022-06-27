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

func newKcj(db *gorm.DB) kcj {
	_kcj := kcj{}

	_kcj.kcjDo.UseDB(db)
	_kcj.kcjDo.UseModel(&model.Kcj{})

	tableName := _kcj.kcjDo.TableName()
	_kcj.ALL = field.NewField(tableName, "*")
	_kcj.KCAAA = field.NewString(tableName, "KCAAA")
	_kcj.SWJCC = field.NewString(tableName, "SWJCC")
	_kcj.SWNB = field.NewString(tableName, "SWNB")
	_kcj.GCAB = field.NewString(tableName, "GCAB")
	_kcj.SWJDAC = field.NewFloat64(tableName, "SWJDAC")
	_kcj.SWAIA = field.NewString(tableName, "SWAIA")
	_kcj.GCCBA = field.NewString(tableName, "GCCBA")
	_kcj.GCBI = field.NewString(tableName, "GCBI")
	_kcj.JJGLDA = field.NewString(tableName, "JJGLDA")
	_kcj.JJGLDB = field.NewString(tableName, "JJGLDB")
	_kcj.BCB = field.NewString(tableName, "BCB")
	_kcj.BPJ = field.NewInt32(tableName, "BPJ")
	_kcj.BPWDX = field.NewString(tableName, "BPWDX")
	_kcj.DDBWDX = field.NewString(tableName, "DDBWDX")
	_kcj.数据 = field.NewString(tableName, "数据")

	_kcj.fillFieldMap()

	return _kcj
}

type kcj struct {
	kcjDo kcjDo

	ALL    field.Field
	KCAAA  field.String
	SWJCC  field.String
	SWNB   field.String
	GCAB   field.String
	SWJDAC field.Float64
	SWAIA  field.String
	GCCBA  field.String
	GCBI   field.String
	JJGLDA field.String
	JJGLDB field.String
	BCB    field.String
	BPJ    field.Int32
	BPWDX  field.String
	DDBWDX field.String
	数据     field.String

	fieldMap map[string]field.Expr
}

func (k kcj) Table(newTableName string) *kcj {
	k.kcjDo.UseTable(newTableName)
	return k.updateTableName(newTableName)
}

func (k kcj) As(alias string) *kcj {
	k.kcjDo.DO = *(k.kcjDo.As(alias).(*gen.DO))
	return k.updateTableName(alias)
}

func (k *kcj) updateTableName(table string) *kcj {
	k.ALL = field.NewField(table, "*")
	k.KCAAA = field.NewString(table, "KCAAA")
	k.SWJCC = field.NewString(table, "SWJCC")
	k.SWNB = field.NewString(table, "SWNB")
	k.GCAB = field.NewString(table, "GCAB")
	k.SWJDAC = field.NewFloat64(table, "SWJDAC")
	k.SWAIA = field.NewString(table, "SWAIA")
	k.GCCBA = field.NewString(table, "GCCBA")
	k.GCBI = field.NewString(table, "GCBI")
	k.JJGLDA = field.NewString(table, "JJGLDA")
	k.JJGLDB = field.NewString(table, "JJGLDB")
	k.BCB = field.NewString(table, "BCB")
	k.BPJ = field.NewInt32(table, "BPJ")
	k.BPWDX = field.NewString(table, "BPWDX")
	k.DDBWDX = field.NewString(table, "DDBWDX")
	k.数据 = field.NewString(table, "数据")

	k.fillFieldMap()

	return k
}

func (k *kcj) WithContext(ctx context.Context) *kcjDo { return k.kcjDo.WithContext(ctx) }

func (k kcj) TableName() string { return k.kcjDo.TableName() }

func (k kcj) Alias() string { return k.kcjDo.Alias() }

func (k *kcj) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := k.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (k *kcj) fillFieldMap() {
	k.fieldMap = make(map[string]field.Expr, 15)
	k.fieldMap["KCAAA"] = k.KCAAA
	k.fieldMap["SWJCC"] = k.SWJCC
	k.fieldMap["SWNB"] = k.SWNB
	k.fieldMap["GCAB"] = k.GCAB
	k.fieldMap["SWJDAC"] = k.SWJDAC
	k.fieldMap["SWAIA"] = k.SWAIA
	k.fieldMap["GCCBA"] = k.GCCBA
	k.fieldMap["GCBI"] = k.GCBI
	k.fieldMap["JJGLDA"] = k.JJGLDA
	k.fieldMap["JJGLDB"] = k.JJGLDB
	k.fieldMap["BCB"] = k.BCB
	k.fieldMap["BPJ"] = k.BPJ
	k.fieldMap["BPWDX"] = k.BPWDX
	k.fieldMap["DDBWDX"] = k.DDBWDX
	k.fieldMap["数据"] = k.数据
}

func (k kcj) clone(db *gorm.DB) kcj {
	k.kcjDo.ReplaceDB(db)
	return k
}

type kcjDo struct{ gen.DO }

func (k kcjDo) Debug() *kcjDo {
	return k.withDO(k.DO.Debug())
}

func (k kcjDo) WithContext(ctx context.Context) *kcjDo {
	return k.withDO(k.DO.WithContext(ctx))
}

func (k kcjDo) ReadDB(ctx context.Context) *kcjDo {
	return k.WithContext(ctx).Clauses(dbresolver.Read)
}

func (k kcjDo) WriteDB(ctx context.Context) *kcjDo {
	return k.WithContext(ctx).Clauses(dbresolver.Write)
}

func (k kcjDo) Clauses(conds ...clause.Expression) *kcjDo {
	return k.withDO(k.DO.Clauses(conds...))
}

func (k kcjDo) Returning(value interface{}, columns ...string) *kcjDo {
	return k.withDO(k.DO.Returning(value, columns...))
}

func (k kcjDo) Not(conds ...gen.Condition) *kcjDo {
	return k.withDO(k.DO.Not(conds...))
}

func (k kcjDo) Or(conds ...gen.Condition) *kcjDo {
	return k.withDO(k.DO.Or(conds...))
}

func (k kcjDo) Select(conds ...field.Expr) *kcjDo {
	return k.withDO(k.DO.Select(conds...))
}

func (k kcjDo) Where(conds ...gen.Condition) *kcjDo {
	return k.withDO(k.DO.Where(conds...))
}

func (k kcjDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *kcjDo {
	return k.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (k kcjDo) Order(conds ...field.Expr) *kcjDo {
	return k.withDO(k.DO.Order(conds...))
}

func (k kcjDo) Distinct(cols ...field.Expr) *kcjDo {
	return k.withDO(k.DO.Distinct(cols...))
}

func (k kcjDo) Omit(cols ...field.Expr) *kcjDo {
	return k.withDO(k.DO.Omit(cols...))
}

func (k kcjDo) Join(table schema.Tabler, on ...field.Expr) *kcjDo {
	return k.withDO(k.DO.Join(table, on...))
}

func (k kcjDo) LeftJoin(table schema.Tabler, on ...field.Expr) *kcjDo {
	return k.withDO(k.DO.LeftJoin(table, on...))
}

func (k kcjDo) RightJoin(table schema.Tabler, on ...field.Expr) *kcjDo {
	return k.withDO(k.DO.RightJoin(table, on...))
}

func (k kcjDo) Group(cols ...field.Expr) *kcjDo {
	return k.withDO(k.DO.Group(cols...))
}

func (k kcjDo) Having(conds ...gen.Condition) *kcjDo {
	return k.withDO(k.DO.Having(conds...))
}

func (k kcjDo) Limit(limit int) *kcjDo {
	return k.withDO(k.DO.Limit(limit))
}

func (k kcjDo) Offset(offset int) *kcjDo {
	return k.withDO(k.DO.Offset(offset))
}

func (k kcjDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *kcjDo {
	return k.withDO(k.DO.Scopes(funcs...))
}

func (k kcjDo) Unscoped() *kcjDo {
	return k.withDO(k.DO.Unscoped())
}

func (k kcjDo) Create(values ...*model.Kcj) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Create(values)
}

func (k kcjDo) CreateInBatches(values []*model.Kcj, batchSize int) error {
	return k.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (k kcjDo) Save(values ...*model.Kcj) error {
	if len(values) == 0 {
		return nil
	}
	return k.DO.Save(values)
}

func (k kcjDo) First() (*model.Kcj, error) {
	if result, err := k.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kcj), nil
	}
}

func (k kcjDo) Take() (*model.Kcj, error) {
	if result, err := k.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kcj), nil
	}
}

func (k kcjDo) Last() (*model.Kcj, error) {
	if result, err := k.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kcj), nil
	}
}

func (k kcjDo) Find() ([]*model.Kcj, error) {
	result, err := k.DO.Find()
	return result.([]*model.Kcj), err
}

func (k kcjDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Kcj, err error) {
	buf := make([]*model.Kcj, 0, batchSize)
	err = k.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (k kcjDo) FindInBatches(result *[]*model.Kcj, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return k.DO.FindInBatches(result, batchSize, fc)
}

func (k kcjDo) Attrs(attrs ...field.AssignExpr) *kcjDo {
	return k.withDO(k.DO.Attrs(attrs...))
}

func (k kcjDo) Assign(attrs ...field.AssignExpr) *kcjDo {
	return k.withDO(k.DO.Assign(attrs...))
}

func (k kcjDo) Joins(fields ...field.RelationField) *kcjDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Joins(_f))
	}
	return &k
}

func (k kcjDo) Preload(fields ...field.RelationField) *kcjDo {
	for _, _f := range fields {
		k = *k.withDO(k.DO.Preload(_f))
	}
	return &k
}

func (k kcjDo) FirstOrInit() (*model.Kcj, error) {
	if result, err := k.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kcj), nil
	}
}

func (k kcjDo) FirstOrCreate() (*model.Kcj, error) {
	if result, err := k.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Kcj), nil
	}
}

func (k kcjDo) FindByPage(offset int, limit int) (result []*model.Kcj, count int64, err error) {
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

func (k kcjDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = k.Count()
	if err != nil {
		return
	}

	err = k.Offset(offset).Limit(limit).Scan(result)
	return
}

func (k *kcjDo) withDO(do gen.Dao) *kcjDo {
	k.DO = *do.(*gen.DO)
	return k
}