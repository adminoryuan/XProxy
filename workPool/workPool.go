package workpool

// 协程池 对 协程做一个限定
type WrokPool struct {
	Max int //同时执行的最大数量

	WorkChan chan Task //接收执行任务
}
type Task func()
