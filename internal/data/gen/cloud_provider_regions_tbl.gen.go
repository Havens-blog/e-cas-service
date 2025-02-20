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

func newCloudProviderRegionsTbl(db *gorm.DB, opts ...gen.DOOption) cloudProviderRegionsTbl {
	_cloudProviderRegionsTbl := cloudProviderRegionsTbl{}

	_cloudProviderRegionsTbl.cloudProviderRegionsTblDo.UseDB(db, opts...)
	_cloudProviderRegionsTbl.cloudProviderRegionsTblDo.UseModel(&model.CloudProviderRegionsTbl{})

	tableName := _cloudProviderRegionsTbl.cloudProviderRegionsTblDo.TableName()
	_cloudProviderRegionsTbl.ALL = field.NewAsterisk(tableName)
	_cloudProviderRegionsTbl.CreatedAt = field.NewTime(tableName, "created_at")
	_cloudProviderRegionsTbl.UpdatedAt = field.NewTime(tableName, "updated_at")
	_cloudProviderRegionsTbl.UpdateVersion = field.NewInt32(tableName, "update_version")
	_cloudProviderRegionsTbl.DeletedAt = field.NewField(tableName, "deleted_at")
	_cloudProviderRegionsTbl.Deleted = field.NewBool(tableName, "deleted")
	_cloudProviderRegionsTbl.RowID = field.NewInt64(tableName, "row_id")
	_cloudProviderRegionsTbl.SyncStatus = field.NewString(tableName, "sync_status")
	_cloudProviderRegionsTbl.LastSync = field.NewTime(tableName, "last_sync")
	_cloudProviderRegionsTbl.LastSyncEndAt = field.NewTime(tableName, "last_sync_end_at")
	_cloudProviderRegionsTbl.CloudregionID = field.NewString(tableName, "cloudregion_id")
	_cloudProviderRegionsTbl.CloudproviderID = field.NewString(tableName, "cloudprovider_id")
	_cloudProviderRegionsTbl.Enabled = field.NewBool(tableName, "enabled")
	_cloudProviderRegionsTbl.SyncResults = field.NewString(tableName, "sync_results")
	_cloudProviderRegionsTbl.LastDeepSyncAt = field.NewTime(tableName, "last_deep_sync_at")
	_cloudProviderRegionsTbl.LastAutoSyncAt = field.NewTime(tableName, "last_auto_sync_at")

	_cloudProviderRegionsTbl.fillFieldMap()

	return _cloudProviderRegionsTbl
}

type cloudProviderRegionsTbl struct {
	cloudProviderRegionsTblDo cloudProviderRegionsTblDo

	ALL             field.Asterisk
	CreatedAt       field.Time
	UpdatedAt       field.Time
	UpdateVersion   field.Int32
	DeletedAt       field.Field
	Deleted         field.Bool
	RowID           field.Int64
	SyncStatus      field.String
	LastSync        field.Time
	LastSyncEndAt   field.Time
	CloudregionID   field.String
	CloudproviderID field.String
	Enabled         field.Bool
	SyncResults     field.String
	LastDeepSyncAt  field.Time
	LastAutoSyncAt  field.Time

	fieldMap map[string]field.Expr
}

func (c cloudProviderRegionsTbl) Table(newTableName string) *cloudProviderRegionsTbl {
	c.cloudProviderRegionsTblDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cloudProviderRegionsTbl) As(alias string) *cloudProviderRegionsTbl {
	c.cloudProviderRegionsTblDo.DO = *(c.cloudProviderRegionsTblDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cloudProviderRegionsTbl) updateTableName(table string) *cloudProviderRegionsTbl {
	c.ALL = field.NewAsterisk(table)
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.UpdateVersion = field.NewInt32(table, "update_version")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Deleted = field.NewBool(table, "deleted")
	c.RowID = field.NewInt64(table, "row_id")
	c.SyncStatus = field.NewString(table, "sync_status")
	c.LastSync = field.NewTime(table, "last_sync")
	c.LastSyncEndAt = field.NewTime(table, "last_sync_end_at")
	c.CloudregionID = field.NewString(table, "cloudregion_id")
	c.CloudproviderID = field.NewString(table, "cloudprovider_id")
	c.Enabled = field.NewBool(table, "enabled")
	c.SyncResults = field.NewString(table, "sync_results")
	c.LastDeepSyncAt = field.NewTime(table, "last_deep_sync_at")
	c.LastAutoSyncAt = field.NewTime(table, "last_auto_sync_at")

	c.fillFieldMap()

	return c
}

func (c *cloudProviderRegionsTbl) WithContext(ctx context.Context) *cloudProviderRegionsTblDo {
	return c.cloudProviderRegionsTblDo.WithContext(ctx)
}

func (c cloudProviderRegionsTbl) TableName() string { return c.cloudProviderRegionsTblDo.TableName() }

func (c cloudProviderRegionsTbl) Alias() string { return c.cloudProviderRegionsTblDo.Alias() }

func (c *cloudProviderRegionsTbl) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cloudProviderRegionsTbl) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 15)
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["update_version"] = c.UpdateVersion
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["deleted"] = c.Deleted
	c.fieldMap["row_id"] = c.RowID
	c.fieldMap["sync_status"] = c.SyncStatus
	c.fieldMap["last_sync"] = c.LastSync
	c.fieldMap["last_sync_end_at"] = c.LastSyncEndAt
	c.fieldMap["cloudregion_id"] = c.CloudregionID
	c.fieldMap["cloudprovider_id"] = c.CloudproviderID
	c.fieldMap["enabled"] = c.Enabled
	c.fieldMap["sync_results"] = c.SyncResults
	c.fieldMap["last_deep_sync_at"] = c.LastDeepSyncAt
	c.fieldMap["last_auto_sync_at"] = c.LastAutoSyncAt
}

func (c cloudProviderRegionsTbl) clone(db *gorm.DB) cloudProviderRegionsTbl {
	c.cloudProviderRegionsTblDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cloudProviderRegionsTbl) replaceDB(db *gorm.DB) cloudProviderRegionsTbl {
	c.cloudProviderRegionsTblDo.ReplaceDB(db)
	return c
}

type cloudProviderRegionsTblDo struct{ gen.DO }

func (c cloudProviderRegionsTblDo) Debug() *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Debug())
}

func (c cloudProviderRegionsTblDo) WithContext(ctx context.Context) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cloudProviderRegionsTblDo) ReadDB() *cloudProviderRegionsTblDo {
	return c.Clauses(dbresolver.Read)
}

func (c cloudProviderRegionsTblDo) WriteDB() *cloudProviderRegionsTblDo {
	return c.Clauses(dbresolver.Write)
}

func (c cloudProviderRegionsTblDo) Session(config *gorm.Session) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Session(config))
}

func (c cloudProviderRegionsTblDo) Clauses(conds ...clause.Expression) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cloudProviderRegionsTblDo) Returning(value interface{}, columns ...string) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cloudProviderRegionsTblDo) Not(conds ...gen.Condition) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cloudProviderRegionsTblDo) Or(conds ...gen.Condition) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cloudProviderRegionsTblDo) Select(conds ...field.Expr) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cloudProviderRegionsTblDo) Where(conds ...gen.Condition) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cloudProviderRegionsTblDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *cloudProviderRegionsTblDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c cloudProviderRegionsTblDo) Order(conds ...field.Expr) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cloudProviderRegionsTblDo) Distinct(cols ...field.Expr) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cloudProviderRegionsTblDo) Omit(cols ...field.Expr) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cloudProviderRegionsTblDo) Join(table schema.Tabler, on ...field.Expr) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cloudProviderRegionsTblDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cloudProviderRegionsTblDo) RightJoin(table schema.Tabler, on ...field.Expr) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cloudProviderRegionsTblDo) Group(cols ...field.Expr) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cloudProviderRegionsTblDo) Having(conds ...gen.Condition) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cloudProviderRegionsTblDo) Limit(limit int) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cloudProviderRegionsTblDo) Offset(offset int) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cloudProviderRegionsTblDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cloudProviderRegionsTblDo) Unscoped() *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cloudProviderRegionsTblDo) Create(values ...*model.CloudProviderRegionsTbl) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cloudProviderRegionsTblDo) CreateInBatches(values []*model.CloudProviderRegionsTbl, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cloudProviderRegionsTblDo) Save(values ...*model.CloudProviderRegionsTbl) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cloudProviderRegionsTblDo) First() (*model.CloudProviderRegionsTbl, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloudProviderRegionsTbl), nil
	}
}

func (c cloudProviderRegionsTblDo) Take() (*model.CloudProviderRegionsTbl, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloudProviderRegionsTbl), nil
	}
}

func (c cloudProviderRegionsTblDo) Last() (*model.CloudProviderRegionsTbl, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloudProviderRegionsTbl), nil
	}
}

func (c cloudProviderRegionsTblDo) Find() ([]*model.CloudProviderRegionsTbl, error) {
	result, err := c.DO.Find()
	return result.([]*model.CloudProviderRegionsTbl), err
}

func (c cloudProviderRegionsTblDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CloudProviderRegionsTbl, err error) {
	buf := make([]*model.CloudProviderRegionsTbl, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cloudProviderRegionsTblDo) FindInBatches(result *[]*model.CloudProviderRegionsTbl, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cloudProviderRegionsTblDo) Attrs(attrs ...field.AssignExpr) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cloudProviderRegionsTblDo) Assign(attrs ...field.AssignExpr) *cloudProviderRegionsTblDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cloudProviderRegionsTblDo) Joins(fields ...field.RelationField) *cloudProviderRegionsTblDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cloudProviderRegionsTblDo) Preload(fields ...field.RelationField) *cloudProviderRegionsTblDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cloudProviderRegionsTblDo) FirstOrInit() (*model.CloudProviderRegionsTbl, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloudProviderRegionsTbl), nil
	}
}

func (c cloudProviderRegionsTblDo) FirstOrCreate() (*model.CloudProviderRegionsTbl, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloudProviderRegionsTbl), nil
	}
}

func (c cloudProviderRegionsTblDo) FindByPage(offset int, limit int) (result []*model.CloudProviderRegionsTbl, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cloudProviderRegionsTblDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cloudProviderRegionsTblDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cloudProviderRegionsTblDo) Delete(models ...*model.CloudProviderRegionsTbl) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cloudProviderRegionsTblDo) withDO(do gen.Dao) *cloudProviderRegionsTblDo {
	c.DO = *do.(*gen.DO)
	return c
}
