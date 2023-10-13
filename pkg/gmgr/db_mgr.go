/**
 * Create Time:2022/12/8
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gmgr

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang/glog"
	"github.com/qionggemens/gcommon/pkg/configcenter"
	"time"
)

func NewDB(tbName string) (*sql.DB, error) {
	mysqlUrl := configcenter.GetString(fmt.Sprintf("db.%s.mysql_url", tbName), "")
	maxOpenConns := configcenter.GetInt(fmt.Sprintf("db.%s.max_open_conns", tbName), 10)
	minIdleConns := configcenter.GetInt(fmt.Sprintf("db.%s.min_idle_conns", tbName), 2)
	maxLifeTime := configcenter.GetInt64(fmt.Sprintf("db.%s.max_life_time", tbName), 60)

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
