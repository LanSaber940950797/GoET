package entity

func (e *Entity) IsFromPool() bool {
	return e.status&IsFromPool == IsFromPool
}

func (e *Entity) SetIsFromPool(value bool) {
	if value {
		e.status |= IsFromPool
	} else {
		e.status &^= IsFromPool
	}
}

func (e *Entity) IsRegister() bool {
	return e.status&IsRegister == IsRegister
}

func (e *Entity) SetIsRegister(value bool) {
	if value {
		e.status |= IsRegister
	} else {
		e.status &^= IsRegister
	}

	if value {
		e.registerSystem()
	}
}

func (e *Entity) registerSystem() {
	e.iScene.Fiber().entitySystem.registerSystem(e)
}

func (e *Entity) IsComponent() bool {
	return e.status&IsComponent == IsComponent
}

func (e *Entity) SetIsComponent(value bool) {
	if value {
		e.status |= IsComponent
	} else {
		e.status &^= IsComponent
	}
}

func (e *Entity) IsNew() bool {
	return e.status&IsNew == IsNew
}

func (e *Entity) SetIsNew(value bool) {
	if value {
		e.status |= IsNew
	} else {
		e.status &^= IsNew
	}
}

func (e *Entity) IsCreated() bool {
	return e.status&IsCreated == IsCreated
}

func (e *Entity) SetIsCreated(value bool) {
	if value {
		e.status |= IsCreated
	} else {
		e.status &^= IsCreated
	}
}
