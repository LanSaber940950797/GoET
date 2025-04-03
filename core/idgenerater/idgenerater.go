package idgenerater

import (
	"GoET/core/options"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const (
	MaxZone   = 1024
	Mask14bit = 0x3fff
	Mask30bit = 0x3fffffff
	Mask20bit = 0xfffff
	Epoch2022 = 1640995200 // 2022-01-01 00:00:00 UTC in seconds
)

type IdStruct struct {
	Process uint16 // 14bit
	Time    uint32 // 30bit
	Value   uint32 // 20bit
}

func (i IdStruct) ToLong() int64 {
	var result uint64
	result |= uint64(i.Process)
	result <<= 30
	result |= uint64(i.Time)
	result <<= 20
	result |= uint64(i.Value)
	return int64(result)
}

func NewIdStruct(time uint32, process uint16, value uint32) IdStruct {
	return IdStruct{
		Process: process,
		Time:    time,
		Value:   value,
	}
}

func NewIdStructFromLong(id int64) IdStruct {
	result := uint64(id)
	value := uint32(result & Mask20bit)
	result >>= 20
	time := uint32(result & Mask30bit)
	result >>= 30
	process := uint16(result & Mask14bit)
	return IdStruct{
		Process: process,
		Time:    time,
		Value:   value,
	}
}

func (i IdStruct) String() string {
	return fmt.Sprintf("process: %d, time: %d, value: %d", i.Process, i.Time, i.Value)
}

type InstanceIdStruct struct {
	Time  uint32 // 32bit
	Value uint32 // 32bit
}

func (i InstanceIdStruct) ToLong() int64 {
	var result uint64
	result |= uint64(i.Time)
	result <<= 32
	result |= uint64(i.Value)
	return int64(result)
}

func NewInstanceIdStruct(time uint32, value uint32) InstanceIdStruct {
	return InstanceIdStruct{
		Time:  time,
		Value: value,
	}
}

func NewInstanceIdStructFromLong(id int64) InstanceIdStruct {
	result := uint64(id)
	value := uint32(result & 0xffffffff)
	result >>= 32
	time := uint32(result & 0xffffffff)
	return InstanceIdStruct{
		Time:  time,
		Value: value,
	}
}

func (i InstanceIdStruct) String() string {
	return fmt.Sprintf("time: %d, value: %d", i.Time, i.Value)
}

type IdGenerater struct {
	epoch2022       int64
	value           int32
	instanceIdValue int32
	mu              sync.Mutex
}

var epoch1970 = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
var idGenerater = &IdGenerater{
	epoch2022:       Epoch2022 - epoch1970,
	value:           0,
	instanceIdValue: 0,
}

func (g *IdGenerater) timeSince2022() uint32 {
	now := time.Now().Unix()
	return uint32(now - g.epoch2022)
}

func (g *IdGenerater) generateId() int64 {
	time := g.timeSince2022()
	var v int32
	g.mu.Lock()
	if g.value > Mask20bit-1 {
		g.value = 0
	}
	v = g.value
	g.value++
	g.mu.Unlock()

	idStruct := NewIdStruct(time, uint16(options.GetProcess()), uint32(v))
	return idStruct.ToLong()
}

func (g *IdGenerater) generateInstanceId() int64 {
	time := g.timeSince2022()
	v := uint32(atomic.AddInt32(&g.instanceIdValue, 1))
	instanceIdStruct := NewInstanceIdStruct(time, v)
	return instanceIdStruct.ToLong()
}

func GenerateId() int64 {
	return idGenerater.generateId()
}

func GenerateInstanceId() int64 {
	return idGenerater.generateInstanceId()
}
