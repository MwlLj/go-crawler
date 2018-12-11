package common

type ICrawler interface {
}

type IQueueCallback interface {
	OnPrepareQueuePop(url []byte)
}

type IQueueMgr interface {
	NewQueue(callback IQueueCallback) IQueue
}

type IQueue interface {
}
