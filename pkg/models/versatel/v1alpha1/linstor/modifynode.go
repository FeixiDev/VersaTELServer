package linstor

import (
	"context"

	"github.com/LINBIT/golinstor/client"
	log "github.com/sirupsen/logrus"
)

// ModifyNode 修改节点类型
func ModifyNode(ctx context.Context, c *client.Client, nodename, nodeType string) error {
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

	return err

	// 返回成功提示
	// successMessage := fmt.Sprintf("节点 %s 的类型已成功修改为 %s", nodename, nodeType)
	// return successMessage

	// 开始以为是get函数问题，所以采用了getall 后遍历的方法。但实际上是不需要的
	// fmt.Printf("执行修改节点类型方法")
	// nodes, err := c.Nodes.GetAll(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, node := range nodes {
	// 	fmt.Printf("node.Name: %s\n", node.Name)
	// 	fmt.Printf("nodename: %s\n", nodename)
	// 	fmt.Printf("nodeType: %s", nodeType)
	// 	if node.Name == nodename {
	// 		fmt.Printf("要修改的节点信息：%+v\n", node)
	// 		// 构建包含修改信息的对象
	// 		nodeModify := client.NodeModify{
	// 			NodeType: nodeType,
	// 		}

	// 		// 调用 LINSTOR API 更新节点信息
	// 		err = c.Nodes.Modify(ctx, nodename, nodeModify)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		fmt.Printf("修改后的节点信息：%+v\n", node)
	// 		fmt.Printf("节点 %s 的类型已成功修改为 %s\n", nodename, nodeType)

	// 		break
	// 	}
	// }
	// return nil
}
