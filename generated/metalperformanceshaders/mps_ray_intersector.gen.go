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

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/cullmode
func (r_ RayIntersector) CullMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("cullMode"))
	return rv
}


// SetCullMode sets the value of the cullMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/cullmode
func (r_ RayIntersector) SetCullMode(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCullMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/frontfacingwinding
func (r_ RayIntersector) FrontFacingWinding() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("frontFacingWinding"))
	return rv
}


// SetFrontFacingWinding sets the value of the frontFacingWinding property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/frontfacingwinding
func (r_ RayIntersector) SetFrontFacingWinding(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFrontFacingWinding:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/intersectiondatatype
func (r_ RayIntersector) IntersectionDataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("intersectionDataType"))
	return rv
}


// SetIntersectionDataType sets the value of the intersectionDataType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/intersectiondatatype
func (r_ RayIntersector) SetIntersectionDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIntersectionDataType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/intersectionstride
func (r_ RayIntersector) IntersectionStride() int {
	rv := objc.Send[int](r_.ID, objc.Sel("intersectionStride"))
	return rv
}


// SetIntersectionStride sets the value of the intersectionStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/intersectionstride
func (r_ RayIntersector) SetIntersectionStride(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIntersectionStride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raydatatype
func (r_ RayIntersector) RayDataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("rayDataType"))
	return rv
}


// SetRayDataType sets the value of the rayDataType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raydatatype
func (r_ RayIntersector) SetRayDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayDataType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/rayindexdatatype
func (r_ RayIntersector) RayIndexDataType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("rayIndexDataType"))
	return rv
}


// SetRayIndexDataType sets the value of the rayIndexDataType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/rayindexdatatype
func (r_ RayIntersector) SetRayIndexDataType(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayIndexDataType:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymask
func (r_ RayIntersector) RayMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("rayMask"))
	return rv
}


// SetRayMask sets the value of the rayMask property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymask
func (r_ RayIntersector) SetRayMask(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymaskoperator
func (r_ RayIntersector) RayMaskOperator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("rayMaskOperator"))
	return rv
}


// SetRayMaskOperator sets the value of the rayMaskOperator property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymaskoperator
func (r_ RayIntersector) SetRayMaskOperator(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayMaskOperator:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymaskoptions
func (r_ RayIntersector) RayMaskOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("rayMaskOptions"))
	return rv
}


// SetRayMaskOptions sets the value of the rayMaskOptions property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymaskoptions
func (r_ RayIntersector) SetRayMaskOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayMaskOptions:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raystride
func (r_ RayIntersector) RayStride() int {
	rv := objc.Send[int](r_.ID, objc.Sel("rayStride"))
	return rv
}


// SetRayStride sets the value of the rayStride property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raystride
func (r_ RayIntersector) SetRayStride(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayStride:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/triangleintersectiontesttype
func (r_ RayIntersector) TriangleIntersectionTestType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("triangleIntersectionTestType"))
	return rv
}


// SetTriangleIntersectionTestType sets the value of the triangleIntersectionTestType property.
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/triangleintersectiontesttype
func (r_ RayIntersector) SetTriangleIntersectionTestType(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTriangleIntersectionTestType:"), value)
}



