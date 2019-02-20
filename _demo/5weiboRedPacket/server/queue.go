package server

import (
	"math/rand"
	"time"
)

type task struct {
	id       uint32
	callback chan uint
}

// 单任务
//var chTasks = make(chan task)

// 多任务
const taskNum = 16

var chTaskList = make([]chan task, taskNum)

func fetchPackageListMoney(chTasks chan task) {
	for {
		t := <-chTasks
		id := t.id

		list1, ok := packageList.Load(id)

		if ok && list1 != nil {
			list := list1.([]uint)
			l := len(list)

			// 分配随机数
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			i := r.Intn(l)
			money := list[i]

			// 更新红包列表中的信息
			if l > 1 {
				if i == l-1 {
					packageList.Store(uint32(id), list[:i])
				} else if i == 0 {
					packageList.Store(uint32(id), list[1:])
				} else {
					packageList.Store(uint32(id), append(list[:i], list[i+1:]...))
				}
			} else {
				packageList.Delete(uint32(id))
			}

			t.callback <- money
		} else {
			t.callback <- 0
		}
	}
}
