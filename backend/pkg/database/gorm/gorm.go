package gorm

import (
	"bookcabin-backend/config"
	"bookcabin-backend/pkg/logger"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

type GormDB struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

func NewGormDB(conf *config.Config, log *logrus.Logger) *GormDB {
	gormLogger := logger.NewCustomGormLogger(log)

	gormConf := &gorm.Config{
		SkipDefaultTransaction: true,
		DryRun:                 conf.Database.DryRun,
		PrepareStmt:            true,
		Logger:                 gormLogger,
	}

	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("pkg/database/%s", conf.Database.File)), gormConf)

	if err != nil {
		panic(fmt.Errorf("failed open database connection: %v", err))
	}

	sqldb, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("connection refused error: %v", err))
	}

	sqldb.SetMaxIdleConns(conf.Database.MaxIdleCons)
	sqldb.SetMaxOpenConns(conf.Database.MaxOpenCons)
	sqldb.SetConnMaxIdleTime(time.Duration(conf.Database.ConnMaxIdleTime) * time.Minute)
	sqldb.SetConnMaxLifetime(time.Duration(conf.Database.ConnMaxLifetime) * time.Minute)

	if err := sqldb.Ping(); err != nil {
		panic(fmt.Errorf("ping database got failed: %v", err))
	}

	return &GormDB{db, sqldb}
}

func (g *GormDB) SqlDB() *sql.DB {
	return g.sqlDB
}

func (g *GormDB) DB() *gorm.DB {
	return g.db
}

// Close current db connection. If database connection is not an io.Closer, returns an error.
func (g *GormDB) Close() {
	err := g.sqlDB.Close()

	if err != nil {
		panic(fmt.Errorf("failed close database connection: %v", err))
	}
}
