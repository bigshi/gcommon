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
)

func TestGetBool(t *testing.T) {
	t.Logf("fdafdsf")
	err := glog.SetLogPath("/tmp/logs/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	glog.Infof("dsafdaf")
	glog.Flush()
}
