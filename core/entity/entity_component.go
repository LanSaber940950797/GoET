package entity

import (
	"errors"
	"fmt"
	"reflect"
)

func AddComponent[T Entity](self *Entity, isFromPool bool) *T {
	return AddComponentWithId[T](self, self.id, isFromPool)
}

func GetComponent[T Entity](e *Entity) *T {
	if e.components == nil {
		return nil
	}

	comptype := reflect.TypeOf((*T)(nil)).Elem()
	c := e.components[comptype]
	return any(c).(*T)
}

func RemoveComponent[T Entity](e *Entity) {
	if e.IsDisposed() {
		return
	}
	if e.components == nil {
		return
	}
	comptype := reflect.TypeOf((*T)(nil)).Elem()
	if c, ok := e.components[comptype]; ok {
		e.removeFromComponents(c)
		c.Dispose()
	}

}

func (e *Entity) RemoveComponent(entity *Entity) {
	if e.IsDisposed() {
		return
	}
	if e.components == nil {
		return
	}
	if c, ok := e.components[reflect.TypeOf(entity).Elem()]; ok {
		if c.instanceId != entity.id {
			return
		}
		e.removeFromComponents(c)
		c.Dispose()
	}
}

func AddComponentWithId[T Entity](e *Entity, id int64, isFromPool bool) *T {
	comptype := reflect.TypeOf((*T)(nil)).Elem()
	if e.components != nil && e.components[comptype] != nil {
		panic("entity already has component")
	}
	t := createEntity[T](isFromPool)
	component := any(t).(*Entity)
	component.id = id
	component.componentParent(e)
	return t
}

func (e *Entity) componentParent(parent *Entity) error {
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
		e.parent.removeFromComponents(e)
	}

	e.parent = parent
	e.SetIsComponent(true)
	e.parent.addToComponents(e)

	if scene, ok := interface{}(e).(IScene); ok {
		scene.SetFiber(e.parent.iScene.Fiber())
		e.SetScene(scene)
	} else {
		e.SetScene(e.parent.iScene)
	}

	return nil
}

func (e *Entity) addToComponents(component *Entity) {
	t := reflect.TypeOf(component).Elem()
	e.components[t] = component
}

func (e *Entity) removeFromComponents(component *Entity) {
	t := reflect.TypeOf(component).Elem()
	delete(e.components, t)
}
