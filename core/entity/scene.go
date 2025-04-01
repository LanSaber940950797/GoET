package entity

type IScene interface {
	GetFiber() *Fiber
	GetSceneType() uint64
}

type Scene struct {
	Entity
	Fiber     *Fiber
	SceneType uint64
}
