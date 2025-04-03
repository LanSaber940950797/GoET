package actor

type Address struct {
	Process int32
	Fiber   int32
}

func NewAddress(process int32, fiber int32) Address {
	return Address{
		Process: process,
		Fiber:   fiber,
	}
}

func (a Address) Equal(b Address) bool {
	return a.Process == b.Process && a.Fiber == b.Fiber
}

type ActorId struct {
	address    Address
	instanceId int64
}

func (a ActorId) Process() int32 {
	return a.address.Process
}

func (a ActorId) Fiber() int32 {
	return a.address.Fiber
}

func (a ActorId) InstanceId() int64 {
	return a.instanceId
}

func (a ActorId) SetProcess(process int32) {
	a.address.Process = process
}

func (a ActorId) SetFiber(fiber int32) {
	a.address.Fiber = fiber
}

func (a ActorId) SetInstanceId(instanceId int64) {
	a.instanceId = instanceId
}

func NewActorIdByAddress(address Address, instanceId int64) ActorId {
	return ActorId{address: address, instanceId: instanceId}
}

func NewActorId(process int32, fiber int32, instanceId int64) ActorId {
	return ActorId{address: NewAddress(process, fiber), instanceId: instanceId}
}
