package idgen

import (
	"errors"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10                      // 节点数
	seqBits     uint8 = 12                      // 1毫秒内可生成的id序号的二进制位数
	workerMax   int64 = -1 ^ (-1 << workerBits) // 节点ID的最大值，用于防止溢出
	seqMax      int64 = -1 ^ (-1 << seqBits)    // 同上，用来表示生成id序号的最大值
	timeShift   uint8 = workerBits + seqBits    // 时间戳向左的偏移量
	workerShift uint8 = seqBits                 // 节点ID向左的偏移量
	epoch       int64 = 1567906170596           // 开始运行时间
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	seq       int64
}

var (
	worker *Worker
)

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID excess of quantity")
	}
	return &Worker{
		timestamp: 0,
		workerId:  workerId,
		seq:       0,
	}, nil
}

func (w *Worker) Next() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.seq = (w.seq + 1) & seqMax
		if w.seq == 0 {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.seq = 0
	}
	w.timestamp = now
	ID := int64((now-epoch)<<timeShift | (w.workerId << workerShift) | (w.seq))
	return ID
}

func GetNextID() int64 {
	var err error
	if worker == nil {
		worker, err = NewWorker(1)
		if err != nil {
			panic(err)
		}
	}
	return worker.Next()
}
