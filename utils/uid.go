package utils

import "github.com/bwmarrin/snowflake"

func GetUid() (uid string, err error) {
	//生成uid
	node, err := snowflake.NewNode(1)
	uid = node.Generate().String()
	return uid, err
}
