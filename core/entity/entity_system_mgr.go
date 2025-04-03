package entity

type IAwake interface {
}

type IUpdate interface {
}

type IDestroy interface {
}

type IAwakeSystem interface {
	Awake()
}

type IUpdateSystem interface {
	Update()
}

type IDestroySystem interface {
	Destroy()
}

type EntitySystemManager struct {
}

var EntitySystemMgr = &EntitySystemManager{}

func (mgr *EntitySystemManager) Awake(e *Entity) {

}

func (mgr *EntitySystemManager) Update(e *Entity) {

}

func (mgr *EntitySystemManager) Destroy(e *Entity) {

}
func (mgr *EntitySystemManager) Register(e *Entity) {

}
