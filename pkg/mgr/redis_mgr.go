/**
 * Create Time:2023/2/7
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package mgr

import (
	"github.com/qionggemens/gcommon/pkg/configcenter"
	"github.com/redis/go-redis/v9"
	"strings"
)

func NewRedis() (*redis.ClusterClient, error) {
	redisAddr := configcenter.GetString("redis.addr", "")
	redisPwd := configcenter.GetString("redis.pwd", "123456")
	redisPoolSize := configcenter.GetInt("redis.pool.size", 10)
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(redisAddr, ","),
		Password: redisPwd,
		PoolSize: redisPoolSize,
	})
	return rdb, nil
}
