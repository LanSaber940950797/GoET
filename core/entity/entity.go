package entity

import "sync/atomic"

var (
	idGener         uint64
	instanceIdGener uint64
)

type Entity struct {
	Id         uint64
	InstanceId uint64
	Children   map[uint64]*Entity
	Components map[uint64]*Entity
}

func CreateEntity[T Entity](isFromPool bool) *T {
	var t *T
	if isFromPool {
		//todo 内存池
	} else {
		t = &T{}
	}

	return t
}

func AddChild[T Entity](self *Entity, isFromPool bool) *T {
	t := CreateEntity[T](isFromPool)
	self.AddChild(any(t).(*Entity))
	return t
}

func AddChildWithId[T Entity](self *Entity, id uint64, isFromPool bool) *T {
	t := CreateEntity[T](isFromPool)
	any(t).(*Entity).Id = atomic.AddUint64(&idGener, 1)
	self.AddChild(any(t).(*Entity))
	return t
}

func AddComponent[T Entity](self *Entity, isFromPool bool) *T {
	t := CreateEntity[T](isFromPool)
	self.AddComponent(any(t).(*Entity))
	return t
}

func (e *Entity) AddChild(entity *Entity) {
	e.Children[entity.Id] = entity
}

func (e *Entity) RemoveChild(entity *Entity) {
	e.RemoveChildById(entity.Id)
}

func (e *Entity) RemoveChildById(id uint64) {
	delete(e.Children, id)
}

func (e *Entity) AddComponent(entity *Entity) {
	e.Children[entity.Id] = entity
}

func (e *Entity) RemoveComponent(entity *Entity) {
	delete(e.Components, entity.Id)
}
