/**
 * Create Time:2023/5/19
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package configcenter

import (
	"fmt"
	"github.com/qionggemens/gcommon/pkg/configcenter"
	"os"
	"testing"
)

func TestMd5(t *testing.T) {
	os.Setenv("NACOS_IP", "192.168.88.42")
	err := configcenter.LoadYamlConfig("", "messaging.yaml")
	fmt.Println(err)
	fmt.Println(configcenter.GetBool("jwt.check.enabled"))
	fmt.Println(configcenter.GetString("jwt.secret.app"))

}
