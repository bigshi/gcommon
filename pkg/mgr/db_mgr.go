/**
 * Create Time:2022/12/8
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package mgr

import (
	"database/sql"
	"errors"
	"github.com/golang/glog"
	"github.com/qionggemens/gcommon/pkg/configcenter"
	"time"
)

func NewDB() (*sql.DB, error) {
	mysqlUrl := configcenter.GetString("db.url", "")
	maxOpenConns := configcenter.GetInt("db.max.open_conns", 10)
	minIdleConns := configcenter.GetInt("db.min.idle_conns", 2)
	maxLifeTime := configcenter.GetInt64("db.max.life_time", 60)

	if mysqlUrl == "" {
		glog.Errorf("InitDB fail - msg:mysqlUrl is empty")
		return nil, errors.New("InitDB fail")
	}
	db, err := sql.Open("mysql", mysqlUrl)
	if err != nil {
		glog.Errorf("InitDB fail - msg:%s", err.Error())
		return nil, errors.New("InitDB fail")
	}
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(minIdleConns)
	db.SetConnMaxLifetime(time.Duration(maxLifeTime) * time.Second)
	glog.Infoln("init db success")
	return db, nil
}
