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

func newLogLogin(db *gorm.DB, opts ...gen.DOOption) logLogin {
	_logLogin := logLogin{}

	_logLogin.logLoginDo.UseDB(db, opts...)
	_logLogin.logLoginDo.UseModel(&model.LogLogin{})

	tableName := _logLogin.logLoginDo.TableName()
	_logLogin.ALL = field.NewAsterisk(tableName)
	_logLogin.ID = field.NewInt64(tableName, "id")
	_logLogin.Username = field.NewString(tableName, "username")
	_logLogin.Status = field.NewInt32(tableName, "status")
	_logLogin.Ipaddr = field.NewString(tableName, "ipaddr")
	_logLogin.LoginLocation = field.NewString(tableName, "login_location")
	_logLogin.Browser = field.NewString(tableName, "browser")
	_logLogin.Os = field.NewString(tableName, "os")
	_logLogin.Platform = field.NewString(tableName, "platform")
	_logLogin.CreateBy = field.NewString(tableName, "create_by")
	_logLogin.UpdateBy = field.NewString(tableName, "update_by")
	_logLogin.Remark = field.NewString(tableName, "remark")
	_logLogin.Msg = field.NewString(tableName, "msg")
	_logLogin.LoginTime = field.NewTime(tableName, "login_time")
	_logLogin.CreatedAt = field.NewTime(tableName, "created_at")
	_logLogin.UpdatedAt = field.NewTime(tableName, "updated_at")
	_logLogin.DeletedAt = field.NewField(tableName, "deleted_at")

	_logLogin.fillFieldMap()

	return _logLogin
}

type logLogin struct {
	logLoginDo logLoginDo

	ALL           field.Asterisk
	ID            field.Int64  // 主键id
	Username      field.String // 用户名
	Status        field.Int32  // 1=正常 2=异常
	Ipaddr        field.String // ip地址
	LoginLocation field.String // 归属地
	Browser       field.String // 浏览器
	Os            field.String // 系统
	Platform      field.String // 固件
	CreateBy      field.String // 创建人
	UpdateBy      field.String // 更新者
	Remark        field.String // 备注
	Msg           field.String // 消息
	LoginTime     field.Time   // 登录时间
	CreatedAt     field.Time   // 创建时间
	UpdatedAt     field.Time   // 更新时间
	DeletedAt     field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (l logLogin) Table(newTableName string) *logLogin {
	l.logLoginDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l logLogin) As(alias string) *logLogin {
	l.logLoginDo.DO = *(l.logLoginDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *logLogin) updateTableName(table string) *logLogin {
	l.ALL = field.NewAsterisk(table)
	l.ID = field.NewInt64(table, "id")
	l.Username = field.NewString(table, "username")
	l.Status = field.NewInt32(table, "status")
	l.Ipaddr = field.NewString(table, "ipaddr")
	l.LoginLocation = field.NewString(table, "login_location")
	l.Browser = field.NewString(table, "browser")
	l.Os = field.NewString(table, "os")
	l.Platform = field.NewString(table, "platform")
	l.CreateBy = field.NewString(table, "create_by")
	l.UpdateBy = field.NewString(table, "update_by")
	l.Remark = field.NewString(table, "remark")
	l.Msg = field.NewString(table, "msg")
	l.LoginTime = field.NewTime(table, "login_time")
	l.CreatedAt = field.NewTime(table, "created_at")
	l.UpdatedAt = field.NewTime(table, "updated_at")
	l.DeletedAt = field.NewField(table, "deleted_at")

	l.fillFieldMap()

	return l
}

func (l *logLogin) WithContext(ctx context.Context) *logLoginDo { return l.logLoginDo.WithContext(ctx) }

func (l logLogin) TableName() string { return l.logLoginDo.TableName() }

func (l logLogin) Alias() string { return l.logLoginDo.Alias() }

func (l *logLogin) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *logLogin) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 16)
	l.fieldMap["id"] = l.ID
	l.fieldMap["username"] = l.Username
	l.fieldMap["status"] = l.Status
	l.fieldMap["ipaddr"] = l.Ipaddr
	l.fieldMap["login_location"] = l.LoginLocation
	l.fieldMap["browser"] = l.Browser
	l.fieldMap["os"] = l.Os
	l.fieldMap["platform"] = l.Platform
	l.fieldMap["create_by"] = l.CreateBy
	l.fieldMap["update_by"] = l.UpdateBy
	l.fieldMap["remark"] = l.Remark
	l.fieldMap["msg"] = l.Msg
	l.fieldMap["login_time"] = l.LoginTime
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
	l.fieldMap["deleted_at"] = l.DeletedAt
}

func (l logLogin) clone(db *gorm.DB) logLogin {
	l.logLoginDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l logLogin) replaceDB(db *gorm.DB) logLogin {
	l.logLoginDo.ReplaceDB(db)
	return l
}

type logLoginDo struct{ gen.DO }

func (l logLoginDo) Debug() *logLoginDo {
	return l.withDO(l.DO.Debug())
}

func (l logLoginDo) WithContext(ctx context.Context) *logLoginDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l logLoginDo) ReadDB() *logLoginDo {
	return l.Clauses(dbresolver.Read)
}

func (l logLoginDo) WriteDB() *logLoginDo {
	return l.Clauses(dbresolver.Write)
}

func (l logLoginDo) Session(config *gorm.Session) *logLoginDo {
	return l.withDO(l.DO.Session(config))
}

func (l logLoginDo) Clauses(conds ...clause.Expression) *logLoginDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l logLoginDo) Returning(value interface{}, columns ...string) *logLoginDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l logLoginDo) Not(conds ...gen.Condition) *logLoginDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l logLoginDo) Or(conds ...gen.Condition) *logLoginDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l logLoginDo) Select(conds ...field.Expr) *logLoginDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l logLoginDo) Where(conds ...gen.Condition) *logLoginDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l logLoginDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *logLoginDo {
	return l.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (l logLoginDo) Order(conds ...field.Expr) *logLoginDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l logLoginDo) Distinct(cols ...field.Expr) *logLoginDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l logLoginDo) Omit(cols ...field.Expr) *logLoginDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l logLoginDo) Join(table schema.Tabler, on ...field.Expr) *logLoginDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l logLoginDo) LeftJoin(table schema.Tabler, on ...field.Expr) *logLoginDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l logLoginDo) RightJoin(table schema.Tabler, on ...field.Expr) *logLoginDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l logLoginDo) Group(cols ...field.Expr) *logLoginDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l logLoginDo) Having(conds ...gen.Condition) *logLoginDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l logLoginDo) Limit(limit int) *logLoginDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l logLoginDo) Offset(offset int) *logLoginDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l logLoginDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *logLoginDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l logLoginDo) Unscoped() *logLoginDo {
	return l.withDO(l.DO.Unscoped())
}

func (l logLoginDo) Create(values ...*model.LogLogin) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l logLoginDo) CreateInBatches(values []*model.LogLogin, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l logLoginDo) Save(values ...*model.LogLogin) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l logLoginDo) First() (*model.LogLogin, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.LogLogin), nil
	}
}

func (l logLoginDo) Take() (*model.LogLogin, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.LogLogin), nil
	}
}

func (l logLoginDo) Last() (*model.LogLogin, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.LogLogin), nil
	}
}

func (l logLoginDo) Find() ([]*model.LogLogin, error) {
	result, err := l.DO.Find()
	return result.([]*model.LogLogin), err
}

func (l logLoginDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.LogLogin, err error) {
	buf := make([]*model.LogLogin, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l logLoginDo) FindInBatches(result *[]*model.LogLogin, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l logLoginDo) Attrs(attrs ...field.AssignExpr) *logLoginDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l logLoginDo) Assign(attrs ...field.AssignExpr) *logLoginDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l logLoginDo) Joins(fields ...field.RelationField) *logLoginDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l logLoginDo) Preload(fields ...field.RelationField) *logLoginDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l logLoginDo) FirstOrInit() (*model.LogLogin, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.LogLogin), nil
	}
}

func (l logLoginDo) FirstOrCreate() (*model.LogLogin, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.LogLogin), nil
	}
}

func (l logLoginDo) FindByPage(offset int, limit int) (result []*model.LogLogin, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l logLoginDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l logLoginDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l logLoginDo) Delete(models ...*model.LogLogin) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *logLoginDo) withDO(do gen.Dao) *logLoginDo {
	l.DO = *do.(*gen.DO)
	return l
}
