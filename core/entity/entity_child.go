package entity

import "GoET/core/idgenerater"

func AddChild[T Entity](self *Entity, isFromPool bool) *T {
	t := createEntity[T](isFromPool)
	e := any(t).(*Entity)
	e.id = idgenerater.GenerateId()
	e.SetParent(self)
	EntitySystemMgr.Awake(e)
	return t
}

func AddChildWithId[T Entity](self *Entity, id int64, isFromPool bool) *T {
	t := createEntity[T](isFromPool)
	e := any(t).(*Entity)
	e.id = id
	e.SetParent(self)
	//如果有Awake
	EntitySystemMgr.Awake(e)
	return t
}

func GetChild[T Entity](e *Entity, id int64) *T {
	if e.children == nil {
		return nil
	}
	c := e.children[id]
	return any(c).(*T)
}

func (e *Entity) RemoveChild(id int64) {
	if e.children == nil {
		return
	}
	c := e.children[id]
	if c == nil {
		return
	}
	delete(e.children, id)
	c.Dispose()
}

func (e *Entity) addToChildren(entity *Entity) {
	e.children[entity.id] = entity
}

func (e *Entity) RemoveFromChildren(entity *Entity) {
	delete(e.children, entity.id)
	// if len(e.children) == 0 {
	// 	e.children = nil
	// }
}
