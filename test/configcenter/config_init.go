/**
 * Create Time:2023/5/19
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package configcenter

import (
	"flag"
	"fmt"
	"github.com/qionggemens/gcommon/pkg/configcenter"
	"os"
)

func init() {
	flag.Set("log_dir", "/tmp/logs/adminside")
	os.Setenv("NACOS_IP", "192.168.88.42")
	err := configcenter.LoadYamlConfig("", "adminside.yaml")
	fmt.Println(err)
}
