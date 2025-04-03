package objectpool

import (
	"reflect"
	"sync"
)

type IPool interface {
	IsFromPool() bool
	SetIsFromPool(bool)
}

type ObjectPool struct {
	Type reflect.Type
	Pool sync.Pool
}

var objectPools sync.Map

// New 创建一个新的对象池
// func New(obj interface{}) *ObjectPool {
// 	t := reflect.TypeOf(obj)
// 	pool := &ObjectPool{
// 		Type: t,
// 		Pool: sync.Pool{
// 			New: func() interface{} {
// 				return reflect.New(t).Interface()
// 			},
// 		},
// 	}
// 	objectPools.Store(t, pool)
// 	return pool
// }

// Fetch 从对象池中获取一个对象指针
func Fetch[T interface{}]() *T {
	t := reflect.TypeOf((*T)(nil)).Elem()
	poolInterface, ok := objectPools.Load(t)
	if !ok {
		// 如果对象池不存在，则创建一个新的对象池
		pool := &ObjectPool{
			Type: t,
			Pool: sync.Pool{
				New: func() interface{} {
					return reflect.New(t).Interface()
				},
			},
		}
		objectPools.Store(t, pool)
		poolInterface = pool
	}
	pool := poolInterface.(*ObjectPool)
	return pool.Pool.Get().(*T)
}

// Release 将对象归还到对象池
func Release(obj interface{}) {
	if ipool, ok := interface{}(obj).(IPool); ok {
		if !ipool.IsFromPool() {
			return
		}
		ipool.SetIsFromPool(false)
	}
	t := reflect.TypeOf(obj).Elem()
	poolInterface, ok := objectPools.Load(t)
	if !ok {
		panic("object pool not found for type " + t.String())
	}
	pool := poolInterface.(*ObjectPool)
	pool.Pool.Put(obj)
}

// func Release[T interface{}](obj *T) {
// 	t := reflect.TypeOf((*T)(nil)).Elem()
// 	poolInterface, ok := objectPools.Load(t)
// 	if !ok {
// 		panic("object pool not found for type " + t.String())
// 	}
// 	pool := poolInterface.(*ObjectPool)
// 	pool.Pool.Put(obj)
// }
