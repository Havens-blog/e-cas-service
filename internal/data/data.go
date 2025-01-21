package data

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"github.com/Havens-blog/e-cas-service/configs/conf"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/consts"
	"github.com/Havens-blog/e-cas-service/internal/data/gen"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDataRepo,
	NewSysRoleRepo,
	NewSysMenuRepo,
	NewSysRoleMenuRepo,
	NewSysUserRepo,
)

type dataRepo struct {
	data *Data
	log  *log.Helper
}

type contextTxKey struct{}

func NewDataRepo(d *Data, lg log.Logger) biz.Repo {
	return &dataRepo{
		data: d,
		log:  log.NewHelper(lg),
	}
}

// Data .
type Data struct {
	db  *gen.Query
	rdb redis.Cmdable //*redis.Client
	es  *elasticsearch.Client
}

// NewData .
func NewData(c *conf.Data, logg log.Logger) (*Data, error) {
	mysqlConfig := mysql.Config{
		DSN:                       c.Mysql.Source, // DSN data source name
		DefaultStringSize:         191,            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,          // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		//Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true, //禁用外键约束
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //禁用表复数形式
			TablePrefix:   "",   //表前缀
		},
	})
	if err != nil {
		return nil, fmt.Errorf("MySQL启动异常: %s", err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	if err := AutoMigrate(db); err != nil {
		return nil, fmt.Errorf("autoMigrate failed: %s", err)
	}
	var rdb *redis.Client
	if c.Redis.Password == "" {
		rdb = redis.NewClient(&redis.Options{
			Addr: c.Redis.Addr,
			DB:   int(c.Redis.Db),
		})
	} else {
		rdb = redis.NewClient(&redis.Options{
			Addr:     c.Redis.Addr,
			Password: c.Redis.Password,
			DB:       int(c.Redis.Db),
		})
	}

	_, err = rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis connect ping failed: %s", err)
	}

	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: c.Es.Address,
		Username:  c.Es.Username,
		Password:  c.Es.Password,
	})
	if err != nil {
		return nil, fmt.Errorf("elasticsearch connect failed: %s", err)
	}

	consts.DB = db
	consts.RDB = rdb
	consts.Rdb = rdb
	consts.ES = es
	return &Data{db: gen.Use(db), rdb: rdb, es: es}, nil
}

func AutoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		return err
	}
	return nil
}

func (d *Data) Query(ctx context.Context) *gen.Query {
	tx, ok := ctx.Value(contextTxKey{}).(*gen.Query)
	if ok {
		return tx
	}
	return d.db
}

func buildLikeValue(key string) string {
	return fmt.Sprintf("%%%s%%", key)
}

func convertPageSize(page, size int32) (limit, offset int) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	limit = int(size)
	offset = int((page - 1) * size)
	return
}
