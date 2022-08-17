package singleton

type HungrySingleton struct {
}

var hungrySingletonInstance *HungrySingleton

func init() {
	hungrySingletonInstance = &HungrySingleton{}
}

func (singleton *HungrySingleton) GetInstance() *HungrySingleton {
	return hungrySingletonInstance
}
