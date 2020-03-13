package main

// @Time : 2020年3月13日19:35:35
// @Author : Lemyhello
// @Desc: 展示如何使用Pool连接池拿到tcp各种应用实例

import (
	"fmt"
	"github.com/streadway/amqp"
	"net"
	"time"
	"tcp-pool/Pool"
)

var (
	mqurl = "amqp://root:root@127.0.0.1:5672/test" //根据实际情况填写mq配置连接
	mqPool Pool.Pool
)

func main()  {
	rabbitmq()
	mysql()
	redis()
	//拿到一个连接
	mq,_ := mqPool.Get()
	//实例化对象
	mqconn :=mq.(*amqp.Connection)
	//将连接放回连接池中
	defer mqPool.Put(mq)
	//开始操作rabbitmq...
	mqconn.Channel()
	//do something....
}

//rabbitmq rabbitmq连接池
func rabbitmq()  {
	//factory 创建连接的方法
	factory := func() (interface{}, error) { return amqp.Dial(mqurl) }
	//close 关闭连接的方法
	close := func(v interface{}) error { return v.(net.Conn).Close() }
	//创建一个连接池： 初始化2，最大连接5，空闲连接数是4
	poolConfig := &Pool.Config{
		InitialCap: 2,
		MaxIdle:    5,
		MaxCap:     4,
		Factory:    factory,
		Close:      close,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	mqPool, _ = Pool.NewChannelPool(poolConfig)
	//从连接池中取得一个连接
	//v, err := p.Get()
	//do something
	//conn :=v.(*amqp.Connection)
	//将连接放回连接池中
	//p.Put(v)
	//释放连接池中的所有连接
	//p.Release()
	//查看当前连接中的数量
	current := mqPool.Len()
	fmt.Println("len=", current)
	return
}

//mysql mysql连接池
func mysql()  {

}

//redis redis连接池
func redis()  {

}