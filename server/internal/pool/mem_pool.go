package pool

import (
	"sync"
)

// 定义 2^n 的分级桶，覆盖常用的网络包大小（128B 到 64KB）
// 你可以根据压测情况自行删减或增加桶
var (
	pool128 = sync.Pool{New: func() any { return make([]byte, 128) }}
	pool512 = sync.Pool{New: func() any { return make([]byte, 512) }}
	pool1K  = sync.Pool{New: func() any { return make([]byte, 1024) }}
	pool4K  = sync.Pool{New: func() any { return make([]byte, 4096) }}
	pool16K = sync.Pool{New: func() any { return make([]byte, 16384) }}
	pool64K = sync.Pool{New: func() any { return make([]byte, 65536) }}
)

// GetBytes 根据传入的 size，自动去匹配最接近且够用的常驻桶
func GetBytes(size int) []byte {
	if size <= 128 {
		return pool128.Get().([]byte)
	}
	if size <= 512 {
		return pool512.Get().([]byte)
	}
	if size <= 1024 {
		return pool1K.Get().([]byte)
	}
	if size <= 4096 {
		return pool4K.Get().([]byte)
	}
	if size <= 16384 {
		return pool16K.Get().([]byte)
	}
	if size <= 65536 {
		return pool64K.Get().([]byte)
	}
	// 超过最大桶的包（比如大于64KB的极端包），直接现场 make，不入池，交给 GC
	return make([]byte, size)
}

// PutBytes 识别切片的真实容量（cap），精准放回对应的桶里
func PutBytes(buf []byte) {
	c := cap(buf)
	if c >= 65536 {
		pool64K.Put(buf)
		return
	}
	if c >= 16384 {
		pool16K.Put(buf)
		return
	}
	if c >= 4096 {
		pool4K.Put(buf)
		return
	}
	if c >= 1024 {
		pool1K.Put(buf)
		return
	}
	if c >= 512 {
		pool512.Put(buf)
		return
	}
	if c >= 128 {
		pool128.Put(buf)
		return
	}
}

var WorkTaskPool = sync.Pool{
	New: func() interface{} {
		return &WorkTask{}
	},
}

func GetWorkTask() *WorkTask {
	return WorkTaskPool.Get().(*WorkTask)
}

func PutWorkTask(task *WorkTask) {
	if task.Body != nil {
		PutBytes(task.Body)
	}
	task.ConnID = 0
	task.CmdID = 0
	task.DataLen = 0
	task.Body = nil
	task.Conn = nil
	WorkTaskPool.Put(task)
}
