/**
 * Create Time:2023/10/31
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package glog

import (
	"fmt"
	"github.com/qionggemens/gcommon/pkg/glog"
	"testing"
	"time"
)

func TestGetBool(t *testing.T) {
	t.Logf("fdafdsf")
	err := glog.SetLogPath("/tmp/logs/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 2000; i++ {
		glog.Infof("dsafdaf")
		glog.Errorf("43143- %d", i)
	}
	time.Sleep(1 * time.Minute)
}
