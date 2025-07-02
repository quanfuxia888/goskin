package common

import (
	"fmt"
	"sync"
)

var (
	codeRegistry = make(map[int]string)
	codeLock     sync.RWMutex
)

func RegisterCode(code int, msgKey string) {
	codeLock.Lock()
	defer codeLock.Unlock()

	if _, exists := codeRegistry[code]; exists {
		panic(fmt.Sprintf("duplicate error code: %d", code))
	}
	codeRegistry[code] = msgKey
}

// 获取注册 key（用于文档/多语言）
func GetMsgKey(code int) string {
	codeLock.RLock()
	defer codeLock.RUnlock()
	return codeRegistry[code]
}

// 获取全部注册项
func GetAllCodes() map[int]string {
	codeLock.RLock()
	defer codeLock.RUnlock()
	result := make(map[int]string)
	for k, v := range codeRegistry {
		result[k] = v
	}
	return result
}
