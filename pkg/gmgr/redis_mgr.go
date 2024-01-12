/**
 * Create Time:2023/2/7
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gmgr

import (
	"github.com/qionggemens/gcommon/pkg/nacos"
	"github.com/redis/go-redis/v9"
	"strings"
)

func NewRedis() (*redis.ClusterClient, error) {
	redisAddr := nacos.GetString("redis.addr", "")
	redisPwd := nacos.GetString("redis.pwd", "123456")
	redisPoolSize := nacos.GetInt("redis.pool.size", 10)
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    strings.Split(redisAddr, ","),
		Password: redisPwd,
		PoolSize: redisPoolSize,
	})
	return rdb, nil
}
