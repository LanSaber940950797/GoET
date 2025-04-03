package entity

import (
	"GoET/core/actor"
	"GoET/core/options"
)

type Fiber struct {
	isDisposed   bool
	id           int32
	zone         int32
	root         *Scene
	entitySystem *EntitySystem
}

func NewFiber(id int32, zone int32, sceneType SceneType, name string) *Fiber {
	f := &Fiber{
		id:   id,
		zone: zone,
	}
	f.root = NewScene(f, int64(id), 1, sceneType, name)
	return f
}
func (f *Fiber) Address() actor.Address {
	return actor.NewAddress(f.Process(), f.id)
}

func (f *Fiber) Process() int32 {
	return options.GetProcess()
}

func (f *Fiber) Root() *Scene {
	return f.root
}

func (f *Fiber) Update() {

}

func (f *Fiber) IsDisposed() bool {
	return f.isDisposed
}

func (f *Fiber) Dispose() {
	if f.IsDisposed() {
		return
	}

	f.isDisposed = true
	f.root.Dispose()
}
