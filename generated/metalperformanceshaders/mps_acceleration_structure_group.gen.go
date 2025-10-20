// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AccelerationStructureGroup] class.
var (
	AccelerationStructureGroupClass     _AccelerationStructureGroupClass
	AccelerationStructureGroupClassOnce sync.Once
)

func getAccelerationStructureGroupClass() _AccelerationStructureGroupClass {
	AccelerationStructureGroupClassOnce.Do(func() {
		AccelerationStructureGroupClass = _AccelerationStructureGroupClass{objc.GetClass("MPSAccelerationStructureGroup")}
	})
	return AccelerationStructureGroupClass
}

type _AccelerationStructureGroupClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructureGroup] class.
type IAccelerationStructureGroup interface {
	objectivec.IObject
}

// A group of acceleration structures.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureGroup
type AccelerationStructureGroup struct {
	objectivec.Object
}

// AccelerationStructureGroupFrom constructs a [AccelerationStructureGroup] from an unsafe.Pointer.
//
// A group of acceleration structures.
func AccelerationStructureGroupFrom(ptr unsafe.Pointer) AccelerationStructureGroup {
	return AccelerationStructureGroup{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureGroupClass) Alloc() AccelerationStructureGroup {
	rv := objc.Send[AccelerationStructureGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureGroupClass) New() AccelerationStructureGroup {
	rv := objc.Send[AccelerationStructureGroup](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructureGroup) Init() AccelerationStructureGroup {
	rv := objc.Send[AccelerationStructureGroup](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructureGroup) Autorelease() AccelerationStructureGroup {
	rv := objc.Send[AccelerationStructureGroup](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructureGroup creates a new AccelerationStructureGroup instance.
func NewAccelerationStructureGroup() AccelerationStructureGroup {
	return getAccelerationStructureGroupClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructureGroup/init(device:)
func NewAccelerationStructureGroupWithDevice(device objc.ID) AccelerationStructureGroup {
	instance := getAccelerationStructureGroupClass().Alloc()
	rv := objc.Send[AccelerationStructureGroup](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}
