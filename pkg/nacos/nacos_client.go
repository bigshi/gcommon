/**
 * Create Time:2024/1/12
 * User: luchao
 * Email: lcmusic1994@gmail.com
 */

package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/qionggemens/gcommon/pkg/gutil"
	"strconv"
	"strings"
)

var clientCfg *constant.ClientConfig
var serverCfg []constant.ServerConfig

// getClientConfig
//
//	@Description: 获取客户端配置
//	@param namespaceId
//	@param moduleId
//	@return *constant.ClientConfig
func getClientConfig(namespaceId string, dataId string) *constant.ClientConfig {
	if clientCfg != nil {
		return clientCfg
	}
	logDir := fmt.Sprintf("/tmp/nacos/log/%s/%s", namespaceId, dataId)
	cacheDir := fmt.Sprintf("/tmp/nacos/cache/%s/%s", namespaceId, dataId)
	clientCfg = constant.NewClientConfig(
		constant.WithNamespaceId(namespaceId),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(logDir),
		constant.WithCacheDir(cacheDir),
		constant.WithLogLevel("debug"),
		constant.WithUsername(gutil.GetNacosUsername()),
		constant.WithPassword(gutil.GetNacosPassword()),
	)
	return clientCfg
}

// getServerConfig
//
//	@Description: 获取服务端配置
//	@return []constant.ServerConfig
func getServerConfig() []constant.ServerConfig {
	if serverCfg != nil {
		return serverCfg
	}
	nacosAddr := gutil.GetNacosAddr()
	addr := strings.Split(nacosAddr, ":")
	port, _ := strconv.ParseUint(addr[1], 10, 64)
	serverCfg = []constant.ServerConfig{
		*constant.NewServerConfig(addr[0], port),
	}
	return serverCfg
}

// initClient
//
//	@Description: 获取客户端配置
//	@param namespaceId
//	@param moduleId
//	@return *constant.ClientConfig
func initClient(namespaceId string, dataId string) (config_client.IConfigClient, naming_client.INamingClient, error) {
	logDir := fmt.Sprintf("/tmp/nacos/log/%s/%s", namespaceId, dataId)
	cacheDir := fmt.Sprintf("/tmp/nacos/cache/%s/%s", namespaceId, dataId)
	clientCfg := constant.NewClientConfig(
		constant.WithNamespaceId(namespaceId),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir(logDir),
		constant.WithCacheDir(cacheDir),
		constant.WithLogLevel("debug"),
	)
	nacosAddr := gutil.GetNacosAddr()
	addr := strings.Split(nacosAddr, ":")
	port, _ := strconv.ParseUint(addr[1], 10, 64)
	serverCfg := []constant.ServerConfig{
		*constant.NewServerConfig(addr[0], port),
	}
	nacosClientParam := vo.NacosClientParam{
		ClientConfig:  clientCfg,
		ServerConfigs: serverCfg,
	}
	configClient, err := clients.NewConfigClient(nacosClientParam)
	if err != nil {
		return nil, nil, err
	}
	namingClient, err := clients.NewNamingClient(nacosClientParam)
	return configClient, namingClient, nil
}
