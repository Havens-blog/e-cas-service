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

func newCloudregionsTbl(db *gorm.DB, opts ...gen.DOOption) cloudregionsTbl {
	_cloudregionsTbl := cloudregionsTbl{}

	_cloudregionsTbl.cloudregionsTblDo.UseDB(db, opts...)
	_cloudregionsTbl.cloudregionsTblDo.UseModel(&model.CloudregionsTbl{})

	tableName := _cloudregionsTbl.cloudregionsTblDo.TableName()
	_cloudregionsTbl.ALL = field.NewAsterisk(tableName)
	_cloudregionsTbl.CreatedAt = field.NewTime(tableName, "created_at")
	_cloudregionsTbl.UpdatedAt = field.NewTime(tableName, "updated_at")
	_cloudregionsTbl.UpdateVersion = field.NewInt32(tableName, "update_version")
	_cloudregionsTbl.DeletedAt = field.NewField(tableName, "deleted_at")
	_cloudregionsTbl.Deleted = field.NewBool(tableName, "deleted")
	_cloudregionsTbl.ID = field.NewString(tableName, "id")
	_cloudregionsTbl.Description = field.NewString(tableName, "description")
	_cloudregionsTbl.IsEmulated = field.NewBool(tableName, "is_emulated")
	_cloudregionsTbl.Name = field.NewString(tableName, "name")
	_cloudregionsTbl.Status = field.NewString(tableName, "status")
	_cloudregionsTbl.Progress = field.NewFloat32(tableName, "progress")
	_cloudregionsTbl.Enabled = field.NewBool(tableName, "enabled")
	_cloudregionsTbl.ExternalID = field.NewString(tableName, "external_id")
	_cloudregionsTbl.ImportedAt = field.NewTime(tableName, "imported_at")
	_cloudregionsTbl.Source = field.NewString(tableName, "source")
	_cloudregionsTbl.Latitude = field.NewFloat32(tableName, "latitude")
	_cloudregionsTbl.Longitude = field.NewFloat32(tableName, "longitude")
	_cloudregionsTbl.City = field.NewString(tableName, "city")
	_cloudregionsTbl.CountryCode = field.NewString(tableName, "country_code")
	_cloudregionsTbl.Environment = field.NewString(tableName, "environment")
	_cloudregionsTbl.Provider = field.NewString(tableName, "provider")

	_cloudregionsTbl.fillFieldMap()

	return _cloudregionsTbl
}

type cloudregionsTbl struct {
	cloudregionsTblDo cloudregionsTblDo

	ALL           field.Asterisk
	CreatedAt     field.Time
	UpdatedAt     field.Time
	UpdateVersion field.Int32
	DeletedAt     field.Field
	Deleted       field.Bool
	ID            field.String
	Description   field.String
	IsEmulated    field.Bool
	Name          field.String
	Status        field.String
	Progress      field.Float32
	Enabled       field.Bool
	ExternalID    field.String
	ImportedAt    field.Time
	Source        field.String
	Latitude      field.Float32
	Longitude     field.Float32
	City          field.String
	CountryCode   field.String
	Environment   field.String
	Provider      field.String

	fieldMap map[string]field.Expr
}

func (c cloudregionsTbl) Table(newTableName string) *cloudregionsTbl {
	c.cloudregionsTblDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cloudregionsTbl) As(alias string) *cloudregionsTbl {
	c.cloudregionsTblDo.DO = *(c.cloudregionsTblDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cloudregionsTbl) updateTableName(table string) *cloudregionsTbl {
	c.ALL = field.NewAsterisk(table)
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.UpdateVersion = field.NewInt32(table, "update_version")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Deleted = field.NewBool(table, "deleted")
	c.ID = field.NewString(table, "id")
	c.Description = field.NewString(table, "description")
	c.IsEmulated = field.NewBool(table, "is_emulated")
	c.Name = field.NewString(table, "name")
	c.Status = field.NewString(table, "status")
	c.Progress = field.NewFloat32(table, "progress")
	c.Enabled = field.NewBool(table, "enabled")
	c.ExternalID = field.NewString(table, "external_id")
	c.ImportedAt = field.NewTime(table, "imported_at")
	c.Source = field.NewString(table, "source")
	c.Latitude = field.NewFloat32(table, "latitude")
	c.Longitude = field.NewFloat32(table, "longitude")
	c.City = field.NewString(table, "city")
	c.CountryCode = field.NewString(table, "country_code")
	c.Environment = field.NewString(table, "environment")
	c.Provider = field.NewString(table, "provider")

	c.fillFieldMap()

	return c
}

func (c *cloudregionsTbl) WithContext(ctx context.Context) *cloudregionsTblDo {
	return c.cloudregionsTblDo.WithContext(ctx)
}

func (c cloudregionsTbl) TableName() string { return c.cloudregionsTblDo.TableName() }

func (c cloudregionsTbl) Alias() string { return c.cloudregionsTblDo.Alias() }

func (c *cloudregionsTbl) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cloudregionsTbl) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 21)
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["update_version"] = c.UpdateVersion
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["deleted"] = c.Deleted
	c.fieldMap["id"] = c.ID
	c.fieldMap["description"] = c.Description
	c.fieldMap["is_emulated"] = c.IsEmulated
	c.fieldMap["name"] = c.Name
	c.fieldMap["status"] = c.Status
	c.fieldMap["progress"] = c.Progress
	c.fieldMap["enabled"] = c.Enabled
	c.fieldMap["external_id"] = c.ExternalID
	c.fieldMap["imported_at"] = c.ImportedAt
	c.fieldMap["source"] = c.Source
	c.fieldMap["latitude"] = c.Latitude
	c.fieldMap["longitude"] = c.Longitude
	c.fieldMap["city"] = c.City
	c.fieldMap["country_code"] = c.CountryCode
	c.fieldMap["environment"] = c.Environment
	c.fieldMap["provider"] = c.Provider
}

func (c cloudregionsTbl) clone(db *gorm.DB) cloudregionsTbl {
	c.cloudregionsTblDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cloudregionsTbl) replaceDB(db *gorm.DB) cloudregionsTbl {
	c.cloudregionsTblDo.ReplaceDB(db)
	return c
}

type cloudregionsTblDo struct{ gen.DO }

func (c cloudregionsTblDo) Debug() *cloudregionsTblDo {
	return c.withDO(c.DO.Debug())
}

func (c cloudregionsTblDo) WithContext(ctx context.Context) *cloudregionsTblDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cloudregionsTblDo) ReadDB() *cloudregionsTblDo {
	return c.Clauses(dbresolver.Read)
}

func (c cloudregionsTblDo) WriteDB() *cloudregionsTblDo {
	return c.Clauses(dbresolver.Write)
}

func (c cloudregionsTblDo) Session(config *gorm.Session) *cloudregionsTblDo {
	return c.withDO(c.DO.Session(config))
}

func (c cloudregionsTblDo) Clauses(conds ...clause.Expression) *cloudregionsTblDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cloudregionsTblDo) Returning(value interface{}, columns ...string) *cloudregionsTblDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cloudregionsTblDo) Not(conds ...gen.Condition) *cloudregionsTblDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cloudregionsTblDo) Or(conds ...gen.Condition) *cloudregionsTblDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cloudregionsTblDo) Select(conds ...field.Expr) *cloudregionsTblDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cloudregionsTblDo) Where(conds ...gen.Condition) *cloudregionsTblDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cloudregionsTblDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *cloudregionsTblDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c cloudregionsTblDo) Order(conds ...field.Expr) *cloudregionsTblDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cloudregionsTblDo) Distinct(cols ...field.Expr) *cloudregionsTblDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cloudregionsTblDo) Omit(cols ...field.Expr) *cloudregionsTblDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cloudregionsTblDo) Join(table schema.Tabler, on ...field.Expr) *cloudregionsTblDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cloudregionsTblDo) LeftJoin(table schema.Tabler, on ...field.Expr) *cloudregionsTblDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cloudregionsTblDo) RightJoin(table schema.Tabler, on ...field.Expr) *cloudregionsTblDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cloudregionsTblDo) Group(cols ...field.Expr) *cloudregionsTblDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cloudregionsTblDo) Having(conds ...gen.Condition) *cloudregionsTblDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cloudregionsTblDo) Limit(limit int) *cloudregionsTblDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cloudregionsTblDo) Offset(offset int) *cloudregionsTblDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cloudregionsTblDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *cloudregionsTblDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cloudregionsTblDo) Unscoped() *cloudregionsTblDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cloudregionsTblDo) Create(values ...*model.CloudregionsTbl) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cloudregionsTblDo) CreateInBatches(values []*model.CloudregionsTbl, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cloudregionsTblDo) Save(values ...*model.CloudregionsTbl) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cloudregionsTblDo) First() (*model.CloudregionsTbl, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloudregionsTbl), nil
	}
}

func (c cloudregionsTblDo) Take() (*model.CloudregionsTbl, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloudregionsTbl), nil
	}
}

func (c cloudregionsTblDo) Last() (*model.CloudregionsTbl, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloudregionsTbl), nil
	}
}

func (c cloudregionsTblDo) Find() ([]*model.CloudregionsTbl, error) {
	result, err := c.DO.Find()
	return result.([]*model.CloudregionsTbl), err
}

func (c cloudregionsTblDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CloudregionsTbl, err error) {
	buf := make([]*model.CloudregionsTbl, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cloudregionsTblDo) FindInBatches(result *[]*model.CloudregionsTbl, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cloudregionsTblDo) Attrs(attrs ...field.AssignExpr) *cloudregionsTblDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cloudregionsTblDo) Assign(attrs ...field.AssignExpr) *cloudregionsTblDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cloudregionsTblDo) Joins(fields ...field.RelationField) *cloudregionsTblDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cloudregionsTblDo) Preload(fields ...field.RelationField) *cloudregionsTblDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cloudregionsTblDo) FirstOrInit() (*model.CloudregionsTbl, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloudregionsTbl), nil
	}
}

func (c cloudregionsTblDo) FirstOrCreate() (*model.CloudregionsTbl, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CloudregionsTbl), nil
	}
}

func (c cloudregionsTblDo) FindByPage(offset int, limit int) (result []*model.CloudregionsTbl, count int64, err error) {
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

func (c cloudregionsTblDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cloudregionsTblDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cloudregionsTblDo) Delete(models ...*model.CloudregionsTbl) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cloudregionsTblDo) withDO(do gen.Dao) *cloudregionsTblDo {
	c.DO = *do.(*gen.DO)
	return c
}
