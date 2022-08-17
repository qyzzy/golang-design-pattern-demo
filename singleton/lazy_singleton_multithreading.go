package singleton

import "sync"

type LazySingletonMultithreading struct {
}

var (
	lazeSingletonMultithreading *LazySingletonMultithreading
	mu                          sync.Mutex
)

// Double check locking
func (singleton *LazySingletonMultithreading) GetInstance() *LazySingletonMultithreading {
	if lazeSingletonMultithreading == nil {
		mu.Lock()
		defer mu.Unlock()
		if lazeSingletonMultithreading == nil {
			lazeSingletonMultithreading = &LazySingletonMultithreading{}
		}
	}
	return lazeSingletonMultithreading
}
