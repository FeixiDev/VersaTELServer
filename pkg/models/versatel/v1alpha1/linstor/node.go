package linstor

import (
	"context"
	"fmt"
	"net"

	"github.com/LINBIT/golinstor/client"
	log "github.com/sirupsen/logrus"
)

func GetNodeData(ctx context.Context, c *client.Client) []map[string]string {
	nodes, err := c.Nodes.GetAll(ctx)
	resources, err := c.Resources.GetResourceView(ctx)
	nodesInfo := []map[string]string{}
	if err != nil {
		log.Fatal(err)
	}
	for _, node := range nodes {
		resNum := 0
		for _, res := range resources {
			if res.NodeName == node.Name {
				resNum++
			}
		}
		sps, err := c.Nodes.GetStoragePools(ctx, node.Name)
		if err != nil {
			log.Fatal(err)
		}
		defaultInterface := node.NetInterfaces[0]
		addr := fmt.Sprintf("%s:%d (%s)", defaultInterface.Address, defaultInterface.SatellitePort, defaultInterface.SatelliteEncryptionType)

		nodeInfo := map[string]string{
			"name":           node.Name,
			"nodeType":       node.Type,
			"resourceNum":    fmt.Sprintf("%d", resNum),
			"storagePoolNum": fmt.Sprintf("%d", len(sps)),
			"addr":           addr,
			"status":         node.ConnectionStatus,
		}
		nodesInfo = append(nodesInfo, nodeInfo)
	}
	return nodesInfo
}

func DescribeNode(ctx context.Context, c *client.Client, nodename string) error {
	_, err := c.Nodes.Get(ctx, nodename)
	return err
}

func CreateNode(ctx context.Context, c *client.Client, name, ip, nodeType string) error {
	// 将ip字符串转换为net.IP
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return fmt.Errorf("无效的IP地址：%s", ip)
	}
	netInterfaces := []client.NetInterface{client.NetInterface{Name: "default", Address: ipAddr, SatellitePort: 3366, SatelliteEncryptionType: "Plain"}}
	node := client.Node{Name: name, Type: nodeType, NetInterfaces: netInterfaces}
	err := c.Nodes.Create(ctx, node)
	return err
}

func DeleteNode(ctx context.Context, c *client.Client, nodename string) error {
	return c.Nodes.Delete(ctx, nodename)
}

// ModifyNode 修改节点类型
func ModifyNode(ctx context.Context, c *client.Client, nodename, nodeType string) (string, error) {
	// 获取节点信息
	// fmt.Printf("执行修改节点类型方法\n")
	node, err := c.Nodes.Get(ctx, nodename)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("要修改的节点信息：%+v\n", node)
	// fmt.Printf("node.Name: %s\n", node.Name)
	// fmt.Printf("nodename: %s\n", nodename)
	// fmt.Printf("nodeType: %s\n", nodeType)

	// 构建包含修改信息的对象
	nodeModify := client.NodeModify{
		NodeType: nodeType,
	}

	// 调用 LINSTOR API 更新节点信息
	err = c.Nodes.Modify(ctx, node.Name, nodeModify)
	if err != nil {
		log.Fatal(err)
	}

	successMessage := fmt.Sprintf("节点 %s 的类型已成功修改为 %s", nodename, nodeType)
	return successMessage, err
}
