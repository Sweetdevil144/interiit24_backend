package cache

import (
	"container/list"
)

var Store map[string]string
var Store_queue *list.List

func Add(key, value string) error {
	const (
		maxSize = 100
	)
	_, found := Store[key]
	if Store_queue.Len() >= maxSize-1 && !found {
		elem := Store_queue.Front()
		delete(Store, elem.Value.(string))
		Store_queue.Remove(Store_queue.Front())
	}
	Store[key] = value
	return nil
}

func Get(key string) (string,bool) {
	value,found:=Store[key]
	return value,found
}

func Init() {
	Store=map[string]string{}
	Store_queue = list.New()
	Store_queue.Init()
}
