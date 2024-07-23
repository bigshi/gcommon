/**
 * Create Time:2023/5/19
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package nacos

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/qionggemens/gcommon/pkg/nacos"
	"testing"
	"time"
)

func TestGetBool(t *testing.T) {
	fmt.Println(nacos.GetBool("jwt.check.enabled", false))
}

func TestGetString(t *testing.T) {
	fmt.Println(nacos.GetString("jwt.secret.app", ""))
}

func TestGetStrList(t *testing.T) {
	fmt.Println(nacos.GetStrList("spring.redis.cluster.nodes"))
	glog.Flush()
	time.Sleep(30 * time.Second)
}
