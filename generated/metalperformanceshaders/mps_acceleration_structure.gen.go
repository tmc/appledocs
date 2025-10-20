// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AccelerationStructure] class.
var (
	AccelerationStructureClass     _AccelerationStructureClass
	AccelerationStructureClassOnce sync.Once
)

func getAccelerationStructureClass() _AccelerationStructureClass {
	AccelerationStructureClassOnce.Do(func() {
		AccelerationStructureClass = _AccelerationStructureClass{objc.GetClass("MPSAccelerationStructure")}
	})
	return AccelerationStructureClass
}

type _AccelerationStructureClass struct {
	class objc.Class
}

// An interface definition for the [AccelerationStructure] class.
type IAccelerationStructure interface {
	IKernel
}

// The base class for data structures that are built over geometry and used to accelerate ray tracing.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure
type AccelerationStructure struct {
	Kernel
}

// AccelerationStructureFrom constructs a [AccelerationStructure] from an unsafe.Pointer.
//
// The base class for data structures that are built over geometry and used to accelerate ray tracing.
func AccelerationStructureFrom(ptr unsafe.Pointer) AccelerationStructure {
	return AccelerationStructure{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AccelerationStructureClass) Alloc() AccelerationStructure {
	rv := objc.Send[AccelerationStructure](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AccelerationStructureClass) New() AccelerationStructure {
	rv := objc.Send[AccelerationStructure](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerationStructure) Init() AccelerationStructure {
	rv := objc.Send[AccelerationStructure](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerationStructure) Autorelease() AccelerationStructure {
	rv := objc.Send[AccelerationStructure](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerationStructure creates a new AccelerationStructure instance.
func NewAccelerationStructure() AccelerationStructure {
	return getAccelerationStructureClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSAccelerationStructure/init(group:)
func NewAccelerationStructureWithGroup(group unsafe.Pointer) AccelerationStructure {
	instance := getAccelerationStructureClass().Alloc()
	rv := objc.Send[AccelerationStructure](instance.ID, objc.Sel("initWithGroup:"), group)
	rv.Autorelease()
	return rv
}



