/**
 * Create Time:2023/11/9
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package gutil

import (
	"fmt"
	"github.com/qionggemens/gcommon/pkg/gutil"
	"testing"
)

func TestRSAGenerate(t *testing.T) {
	prk, puk, err := gutil.RSAGenerate(1024)
	fmt.Println(err)

	fmt.Println(string(prk))
	fmt.Println(string(puk))
}
