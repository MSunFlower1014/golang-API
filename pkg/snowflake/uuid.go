package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/prometheus/common/log"
)

var node *snowflake.Node

func init() {
	// Node需要唯一
	temp, err := snowflake.NewNode(1)
	if err != nil {
		log.Infof("snowflake init err : %v", err)
	}
	node = temp
}

func GetUID() uint {
	return uint(node.Generate().Int64())
}
