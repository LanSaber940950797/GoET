package entity

type SceneType int64

const (
	SceneType_None SceneType = 0
	SceneType_Main SceneType = 1 << iota
	SceneType_NetInner
	SceneType_Realm
	SceneType_Gate
)

type IScene interface {
	Fiber() *Fiber
	SetFiber(*Fiber)
	SceneType() uint64
	SetSceneType(uint64)
}

type Scene struct {
	Entity
	fiber     *Fiber
	sceneType uint64

	Name string
}

func NewScene(fiber *Fiber, id int64, instanceId int64, sceneType SceneType, name string) *Scene {
	s := &Scene{}
	s.id = id
	s.instanceId = instanceId
	s.Name = name
	s.SetIsCreated(true)
	s.SetIsNew(true)
	s.fiber = fiber
	s.SetScene(s)
	s.SetIsRegister(true)
	return s
}

func (s *Scene) Fiber() *Fiber {
	return s.fiber
}

func (s *Scene) SetFiber(fiber *Fiber) {
	s.fiber = fiber
}

func (s *Scene) SceneType() uint64 {
	return s.sceneType
}

func (s *Scene) SetSceneType(sceneType uint64) {
	s.sceneType = sceneType
}
