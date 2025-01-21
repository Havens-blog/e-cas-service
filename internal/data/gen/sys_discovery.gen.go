// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/Havens-blog/e-cas-service/internal/data/model"
)

func newSysDiscovery(db *gorm.DB, opts ...gen.DOOption) sysDiscovery {
	_sysDiscovery := sysDiscovery{}

	_sysDiscovery.sysDiscoveryDo.UseDB(db, opts...)
	_sysDiscovery.sysDiscoveryDo.UseModel(&model.SysDiscovery{})

	tableName := _sysDiscovery.sysDiscoveryDo.TableName()
	_sysDiscovery.ALL = field.NewAsterisk(tableName)
	_sysDiscovery.ID = field.NewInt32(tableName, "id")
	_sysDiscovery.Name = field.NewString(tableName, "name")
	_sysDiscovery.Picture = field.NewString(tableName, "picture")
	_sysDiscovery.Rank = field.NewInt32(tableName, "rank")
	_sysDiscovery.Link = field.NewString(tableName, "link")
	_sysDiscovery.Status = field.NewInt32(tableName, "status")

	_sysDiscovery.fillFieldMap()

	return _sysDiscovery
}

type sysDiscovery struct {
	sysDiscoveryDo sysDiscoveryDo

	ALL     field.Asterisk
	ID      field.Int32
	Name    field.String // 名称
	Picture field.String // 发现页图片
	Rank    field.Int32  // 排序
	Link    field.String // 发现页链接
	Status  field.Int32  // 发现页状态, 1-正常，2-异常

	fieldMap map[string]field.Expr
}

func (s sysDiscovery) Table(newTableName string) *sysDiscovery {
	s.sysDiscoveryDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sysDiscovery) As(alias string) *sysDiscovery {
	s.sysDiscoveryDo.DO = *(s.sysDiscoveryDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sysDiscovery) updateTableName(table string) *sysDiscovery {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.Name = field.NewString(table, "name")
	s.Picture = field.NewString(table, "picture")
	s.Rank = field.NewInt32(table, "rank")
	s.Link = field.NewString(table, "link")
	s.Status = field.NewInt32(table, "status")

	s.fillFieldMap()

	return s
}

func (s *sysDiscovery) WithContext(ctx context.Context) *sysDiscoveryDo {
	return s.sysDiscoveryDo.WithContext(ctx)
}

func (s sysDiscovery) TableName() string { return s.sysDiscoveryDo.TableName() }

func (s sysDiscovery) Alias() string { return s.sysDiscoveryDo.Alias() }

func (s *sysDiscovery) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sysDiscovery) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 6)
	s.fieldMap["id"] = s.ID
	s.fieldMap["name"] = s.Name
	s.fieldMap["picture"] = s.Picture
	s.fieldMap["rank"] = s.Rank
	s.fieldMap["link"] = s.Link
	s.fieldMap["status"] = s.Status
}

func (s sysDiscovery) clone(db *gorm.DB) sysDiscovery {
	s.sysDiscoveryDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sysDiscovery) replaceDB(db *gorm.DB) sysDiscovery {
	s.sysDiscoveryDo.ReplaceDB(db)
	return s
}

type sysDiscoveryDo struct{ gen.DO }

func (s sysDiscoveryDo) Debug() *sysDiscoveryDo {
	return s.withDO(s.DO.Debug())
}

func (s sysDiscoveryDo) WithContext(ctx context.Context) *sysDiscoveryDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sysDiscoveryDo) ReadDB() *sysDiscoveryDo {
	return s.Clauses(dbresolver.Read)
}

func (s sysDiscoveryDo) WriteDB() *sysDiscoveryDo {
	return s.Clauses(dbresolver.Write)
}

func (s sysDiscoveryDo) Session(config *gorm.Session) *sysDiscoveryDo {
	return s.withDO(s.DO.Session(config))
}

func (s sysDiscoveryDo) Clauses(conds ...clause.Expression) *sysDiscoveryDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sysDiscoveryDo) Returning(value interface{}, columns ...string) *sysDiscoveryDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sysDiscoveryDo) Not(conds ...gen.Condition) *sysDiscoveryDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sysDiscoveryDo) Or(conds ...gen.Condition) *sysDiscoveryDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sysDiscoveryDo) Select(conds ...field.Expr) *sysDiscoveryDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sysDiscoveryDo) Where(conds ...gen.Condition) *sysDiscoveryDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sysDiscoveryDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *sysDiscoveryDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s sysDiscoveryDo) Order(conds ...field.Expr) *sysDiscoveryDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sysDiscoveryDo) Distinct(cols ...field.Expr) *sysDiscoveryDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sysDiscoveryDo) Omit(cols ...field.Expr) *sysDiscoveryDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sysDiscoveryDo) Join(table schema.Tabler, on ...field.Expr) *sysDiscoveryDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sysDiscoveryDo) LeftJoin(table schema.Tabler, on ...field.Expr) *sysDiscoveryDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sysDiscoveryDo) RightJoin(table schema.Tabler, on ...field.Expr) *sysDiscoveryDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sysDiscoveryDo) Group(cols ...field.Expr) *sysDiscoveryDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sysDiscoveryDo) Having(conds ...gen.Condition) *sysDiscoveryDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sysDiscoveryDo) Limit(limit int) *sysDiscoveryDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sysDiscoveryDo) Offset(offset int) *sysDiscoveryDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sysDiscoveryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *sysDiscoveryDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sysDiscoveryDo) Unscoped() *sysDiscoveryDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sysDiscoveryDo) Create(values ...*model.SysDiscovery) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sysDiscoveryDo) CreateInBatches(values []*model.SysDiscovery, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sysDiscoveryDo) Save(values ...*model.SysDiscovery) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sysDiscoveryDo) First() (*model.SysDiscovery, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDiscovery), nil
	}
}

func (s sysDiscoveryDo) Take() (*model.SysDiscovery, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDiscovery), nil
	}
}

func (s sysDiscoveryDo) Last() (*model.SysDiscovery, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDiscovery), nil
	}
}

func (s sysDiscoveryDo) Find() ([]*model.SysDiscovery, error) {
	result, err := s.DO.Find()
	return result.([]*model.SysDiscovery), err
}

func (s sysDiscoveryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SysDiscovery, err error) {
	buf := make([]*model.SysDiscovery, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sysDiscoveryDo) FindInBatches(result *[]*model.SysDiscovery, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sysDiscoveryDo) Attrs(attrs ...field.AssignExpr) *sysDiscoveryDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sysDiscoveryDo) Assign(attrs ...field.AssignExpr) *sysDiscoveryDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sysDiscoveryDo) Joins(fields ...field.RelationField) *sysDiscoveryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sysDiscoveryDo) Preload(fields ...field.RelationField) *sysDiscoveryDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sysDiscoveryDo) FirstOrInit() (*model.SysDiscovery, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDiscovery), nil
	}
}

func (s sysDiscoveryDo) FirstOrCreate() (*model.SysDiscovery, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SysDiscovery), nil
	}
}

func (s sysDiscoveryDo) FindByPage(offset int, limit int) (result []*model.SysDiscovery, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sysDiscoveryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sysDiscoveryDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sysDiscoveryDo) Delete(models ...*model.SysDiscovery) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sysDiscoveryDo) withDO(do gen.Dao) *sysDiscoveryDo {
	s.DO = *do.(*gen.DO)
	return s
}
