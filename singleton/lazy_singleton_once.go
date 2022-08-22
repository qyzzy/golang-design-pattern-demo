package singleton

import "sync"

type LazySingletonOnce struct {
}

var (
	lazySingletonOnce *LazySingletonOnce
	once              = &sync.Once{}
)

func (singleton *LazySingletonOnce) GetInstance() *LazySingletonOnce {
	once.Do(func() {
		lazySingletonOnce = &LazySingletonOnce{}
	})
	return lazySingletonOnce
}
