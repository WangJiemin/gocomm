package gocomm

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

func ZKConnect(zkServer []string) *zk.Conn {
	c, _, err := zk.Connect(zkServer, time.Second) //*10)
	if err != nil {
		panic(err)
	}
	return c
}

func ZKGet(c *zk.Conn, path string) []byte {
	res, _, err := c.Get(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("zk get res:%s\n",string(res))
	return res
}

func ZKPut(c *zk.Conn, path string, data []byte) {
	fmt.Printf("zk path:%s\n", path)
	exist, _, _, err := c.ExistsW(path)
	if err == nil && exist {
		fmt.Printf("zk exist,set value,%s\n", string(data))
		c.Set(path, data, 1)
	} else {
		fmt.Printf("zk not exist,create value,%s\n", string(data))
		c.Create(path, data, 0, zk.WorldACL(zk.PermAll))
	}
}
