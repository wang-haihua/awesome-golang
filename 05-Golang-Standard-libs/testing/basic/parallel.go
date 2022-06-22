package basic

import "sync"

var (
	data   = make(map[string]string)
	locker sync.RWMutex
)

func WriteToMap(key, val string) {
	locker.Lock()
	defer locker.Unlock()
	data[key] = val
}

func ReadFromMap(key string) string {
	locker.RLock()
	defer locker.RUnlock()
	return data[key]
}
