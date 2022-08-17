package singleton

type LazySingleton struct {
}

var lazySingleton *LazySingleton

func (singleton *LazySingleton) GetInstance() *LazySingleton {
	if lazySingleton == nil {
		lazySingleton = &LazySingleton{}
	}
	return lazySingleton
}
