package entity

import (
	"GoET/core/idgenerater"
	"GoET/core/objectpool"
	"errors"
	"fmt"
	"reflect"
)

type EntityStatus byte

const (
	None       EntityStatus = 0
	IsFromPool EntityStatus = 1 << iota
	IsRegister
	IsComponent
	IsCreated
	IsNew
)

type Entity struct {
	id         int64
	instanceId int64
	status     EntityStatus
	parent     *Entity
	iScene     IScene
	children   map[int64]*Entity
	components map[reflect.Type]*Entity
}

func createEntity[T Entity](isFromPool bool) *T {
	var t *T
	if isFromPool {
		//todo 内存池
		return objectpool.Fetch[T]()
	} else {
		t = &T{}
	}

	return t
}

func (e *Entity) GetParent() *Entity {
	return e.parent
}

func (e *Entity) SetParent(parent *Entity) error {
	if parent == nil {
		return errors.New("cannot set parent to nil")
	}
	if parent == e {
		return errors.New("cannot set parent to self")
	}
	if parent.iScene == nil {
		return errors.New("cannot set parent because parent domain is nil")
	}

	if e.parent != nil {
		if e.parent == parent {
			fmt.Println("重复设置了Parent")
			return nil
		}
		e.parent.RemoveFromChildren(e)
	}

	e.parent = parent
	e.SetIsComponent(false)
	e.parent.addToChildren(e)

	if scene, ok := interface{}(e).(IScene); ok {
		scene.SetFiber(e.parent.iScene.Fiber())
		e.SetScene(scene)
	} else {
		e.SetScene(e.parent.iScene)
	}

	return nil
}

func (e *Entity) Scene() IScene {
	return e.iScene
}

func (e *Entity) SetScene(scene IScene) {
	if scene == nil {
		panic("scene is nil")
	}

	if e.iScene == scene {
		return
	}

	preScene := e.iScene
	e.iScene = scene
	if preScene == nil {
		if e.instanceId == 0 {
			e.instanceId = idgenerater.GenerateInstanceId()
		}
		e.SetIsRegister(true)
	}

	if e.children != nil {
		for _, child := range e.children {
			child.SetScene(scene)
		}
	}

	if e.components != nil {
		for _, component := range e.components {
			component.SetScene(scene)
		}
	}

	if !e.IsCreated() {
		e.SetIsCreated(true)
	}
}

func (e *Entity) IsDisposed() bool {
	return e.instanceId == 0
}

func (e *Entity) Dispose() {
	if e.IsDisposed() {
		return
	}

	e.SetIsRegister(false)
	e.instanceId = 0

	// 清理Children
	if e.children != nil {
		for _, child := range e.children {
			child.Dispose()
		}
		//e.children = nil
	}

	// 清理Component
	if e.components != nil {
		for _, component := range e.components {
			component.Dispose()
		}
		e.components = nil
	}

	// 触发Destroy事件
	if _, ok := interface{}(e).(IDestroy); ok {
		EntitySystemMgr.Destroy(e)
	}

	e.iScene = nil

	if e.parent != nil && !e.parent.IsDisposed() {
		if e.IsComponent() {
			e.parent.RemoveComponent(e)
		} else {
			e.parent.RemoveFromChildren(e)
		}
	}

	e.parent = nil
	isFromPool := e.IsFromPool()
	e.status = 0
	e.SetIsFromPool(isFromPool)
	objectpool.Release(e)
}

func (e *Entity) Parent() *Entity {
	return e.parent
}

func GetParent[T Entity](e *Entity) *T {
	return any(e.parent).(*T)
}
