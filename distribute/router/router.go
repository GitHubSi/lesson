package router

import (
	"fmt"
	"sync"
)

const (
	DistributeChannelNum = 128
	OperationRemove      = 1
	OperationAdd         = 2
)

type DistributeMsg struct {
	LectureId    int
	OperatorType int
}

type DistributeGo struct {

	// Identity string， 用来区分K1_K2的，所有K1_K2的都有一个专门的channel来处理
	Identity string

	// 监听的消息队列
	MsgChannel chan *DistributeMsg
}

func (distributeGo *DistributeGo) Run() {
	go func() {
		fmt.Println(distributeGo.Identity, "=================")

		defer func() {
			if r := recover(); r != nil {
				fmt.Println(r)
			}
		}()

		for {
			msg := <-distributeGo.MsgChannel

			// REMOVE操作, 对应OLD的情况
			if msg.OperatorType == OperationRemove {
				fmt.Println("Remove")
			}

			// ADD操作，对应NEW的情况
			if msg.OperatorType == OperationAdd {
				fmt.Println("Add")
			}
		}
	}()
}

// 直接声明全局变量
var GlobalDistribute = Distribute{
	Router: make(map[string]*DistributeGo),
}

type Distribute struct {

	// 读写锁
	RouterLock sync.Mutex

	// 按照Key的纬度就声明Goroutine，如果不行，调整成k1、k2的方式
	Router map[string]*DistributeGo
}

func (dist *Distribute) HandleChannel(identityKey string, lectureId int, operationType int) {

	// 加锁，Map操作，防止并发
	dist.RouterLock.Lock()
	defer dist.RouterLock.Unlock()

	msg := &DistributeMsg{
		LectureId:    lectureId,
		OperatorType: operationType,
	}

	distributeGo, isOk := dist.Router[identityKey]
	if !isOk {
		distributeGo = &DistributeGo{
			MsgChannel: make(chan *DistributeMsg, DistributeChannelNum),
			Identity:   identityKey,
		}
		dist.Router[identityKey] = distributeGo
		distributeGo.Run()
	}

	distributeGo.MsgChannel <- msg
}
