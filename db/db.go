package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

const (
	DBHealthCheckInterval       = 10 * time.Second
	DefaultDBAllowedPacket      = 16 * 1024 * 1024
	DBConnHealthCheckTimeout    = 10 * time.Second
	DBConnHealthCheckRetryTimes = 3
)

var (
	errDBConnHealthCheckTimeout = errors.New("db connection health check timeout")
	errOrmNotInited             = errors.New("orm  not inited")
)

var (
	omOnce sync.Once
	gOrm   *gorm.DB
)

// Init init db
func Init() error {
	var err error
	omOnce.Do(func() {
		err = initDB()
		if err != nil {
			return
		}

		go dbConnectionHealthCheck()
	})

	return err
}

func initDB() error {
	dbSourceName := getDBSourceName()
	logs.Info("db source name ", dbSourceName)

	ormConf := &gorm.Config{}

	db, err := gorm.Open(mysql.Open(dbSourceName), ormConf)
	if err != nil {
		logs.Error("init db engine error", "err", err)
		return err
	}

	//// 设置数据库表不为复数
	//db.SingularTable(true)
	gOrm = db

	logs.Info("init db connection success")
	return nil
}

// getBSourceName get dbSourceName
func getDBSourceName() string {
	user := beego.AppConfig.String("db::user")
	passwd := beego.AppConfig.String("db::password")
	host := beego.AppConfig.String("db::host")
	port := beego.AppConfig.String("db::port")
	dbname := beego.AppConfig.String("db::database")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, passwd,
		host, port, dbname)
}

// connHealthChecker check db connection health
func connHealthChecker() error {
	ctx, cancel := context.WithTimeout(context.TODO(), DBConnHealthCheckTimeout)
	defer cancel()

	if gOrm == nil {
		logs.Error("orm is nil, maybe not inited")
		return errOrmNotInited
	}

	errChan := make(chan error)
	go func() {
		if err := gOrm.Exec("select 1").Error; err != nil {
			logs.Error("db connection health check error", "err", err)
		}
		errChan <- nil
	}()

	select {
	case err := <-errChan:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		logs.Error("db connection health check timeout")
		return errDBConnHealthCheckTimeout
	}
	return nil
}

// dbConnectionHealthCheck check db connection health with retry
func dbConnectionHealthCheck() {
	ticker := time.NewTicker(DBHealthCheckInterval)
	for range ticker.C {
		var err error
		for i := 0; i < DBConnHealthCheckRetryTimes; i++ {
			err = connHealthChecker()
			if err != nil {
				logs.Error("db health check error ,begin retry", "err", err)
				continue
			}
			break
		}
		if err != nil {
			logs.Error("db health check error after retry", "err", err)
		}
	}
}

func GetORM() *gorm.DB {
	if gOrm == nil {
		return nil
	}

	return gOrm
}
