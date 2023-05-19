/**
 * Create Time:2023/5/19
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package configcenter

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/qionggemens/gcommon/pkg/configcenter"
	"testing"
	"time"
)

func TestGetBool(t *testing.T) {
	fmt.Println(configcenter.GetBool("jwt.check.enabled"))
}

func TestGetString(t *testing.T) {
	fmt.Println(configcenter.GetString("jwt.secret.app"))
}

func TestGetStrList(t *testing.T) {
	fmt.Println(configcenter.GetStrList("spring.redis.cluster.nodes"))
	glog.Flush()
	time.Sleep(30 * time.Second)
}
