// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RayIntersector] class.
var (
	RayIntersectorClass     _RayIntersectorClass
	RayIntersectorClassOnce sync.Once
)

func getRayIntersectorClass() _RayIntersectorClass {
	RayIntersectorClassOnce.Do(func() {
		RayIntersectorClass = _RayIntersectorClass{objc.GetClass("MPSRayIntersector")}
	})
	return RayIntersectorClass
}

type _RayIntersectorClass struct {
	class objc.Class
}

// An interface definition for the [RayIntersector] class.
type IRayIntersector interface {
	IKernel
}

// A kernel that performs intersection tests between rays and geometry.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector
type RayIntersector struct {
	Kernel
}

// RayIntersectorFrom constructs a [RayIntersector] from an unsafe.Pointer.
//
// A kernel that performs intersection tests between rays and geometry.
func RayIntersectorFrom(ptr unsafe.Pointer) RayIntersector {
	return RayIntersector{
		Kernel: KernelFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RayIntersectorClass) Alloc() RayIntersector {
	rv := objc.Send[RayIntersector](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RayIntersectorClass) New() RayIntersector {
	rv := objc.Send[RayIntersector](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RayIntersector) Init() RayIntersector {
	rv := objc.Send[RayIntersector](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RayIntersector) Autorelease() RayIntersector {
	rv := objc.Send[RayIntersector](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRayIntersector creates a new RayIntersector instance.
func NewRayIntersector() RayIntersector {
	return getRayIntersectorClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/boundingBoxIntersectionTestType
func (r_ RayIntersector) BoundingBoxIntersectionTestType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("boundingBoxIntersectionTestType"))
	return rv
}


// SetBoundingBoxIntersectionTestType sets the value of the boundingBoxIntersectionTestType property.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/boundingBoxIntersectionTestType
func (r_ RayIntersector) SetBoundingBoxIntersectionTestType(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBoundingBoxIntersectionTestType:"), value)
}


