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
	// properties:
	BoundingBoxIntersectionTestType() BoundingBoxIntersectionTestType /* not a class type */
	SetBoundingBoxIntersectionTestType(value BoundingBoxIntersectionTestType /* not a class type */)
	CullMode() CullMode /* not a class type */
	SetCullMode(value CullMode /* not a class type */)
	FrontFacingWinding() Winding /* not a class type */
	SetFrontFacingWinding(value Winding /* not a class type */)
	IntersectionDataType() IntersectionDataType /* not a class type */
	SetIntersectionDataType(value IntersectionDataType /* not a class type */)
	IntersectionStride() int
	SetIntersectionStride(value int)
	RayDataType() RayDataType /* not a class type */
	SetRayDataType(value RayDataType /* not a class type */)
	RayIndexDataType() DataType /* not a class type */
	SetRayIndexDataType(value DataType /* not a class type */)
	RayMask() unsafe.Pointer
	SetRayMask(value unsafe.Pointer)
	RayMaskOperator() RayMaskOperator /* not a class type */
	SetRayMaskOperator(value RayMaskOperator /* not a class type */)
	RayMaskOptions() RayMaskOptions /* not a class type */
	SetRayMaskOptions(value RayMaskOptions /* not a class type */)
	RayStride() int
	SetRayStride(value int)
	TriangleIntersectionTestType() TriangleIntersectionTestType /* not a class type */
	SetTriangleIntersectionTestType(value TriangleIntersectionTestType /* not a class type */)
	// methods:
}

// A kernel that performs intersection tests between rays and geometry.


// A kernel that performs intersection tests between rays and geometry.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/boundingBoxIntersectionTestType
func (r_ RayIntersector) BoundingBoxIntersectionTestType() BoundingBoxIntersectionTestType /* not a class type */ {
	rv := objc.Send[BoundingBoxIntersectionTestType](r_.ID, objc.Sel("boundingBoxIntersectionTestType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSRayIntersector/boundingBoxIntersectionTestType
func (r_ RayIntersector) SetBoundingBoxIntersectionTestType(value BoundingBoxIntersectionTestType /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBoundingBoxIntersectionTestType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/cullmode
func (r_ RayIntersector) CullMode() CullMode /* not a class type */ {
	rv := objc.Send[CullMode](r_.ID, objc.Sel("cullMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/cullmode
func (r_ RayIntersector) SetCullMode(value CullMode /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCullMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/frontfacingwinding
func (r_ RayIntersector) FrontFacingWinding() Winding /* not a class type */ {
	rv := objc.Send[Winding](r_.ID, objc.Sel("frontFacingWinding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/frontfacingwinding
func (r_ RayIntersector) SetFrontFacingWinding(value Winding /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFrontFacingWinding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/intersectiondatatype
func (r_ RayIntersector) IntersectionDataType() IntersectionDataType /* not a class type */ {
	rv := objc.Send[IntersectionDataType](r_.ID, objc.Sel("intersectionDataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/intersectiondatatype
func (r_ RayIntersector) SetIntersectionDataType(value IntersectionDataType /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIntersectionDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/intersectionstride
func (r_ RayIntersector) IntersectionStride() int {
	rv := objc.Send[int](r_.ID, objc.Sel("intersectionStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/intersectionstride
func (r_ RayIntersector) SetIntersectionStride(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIntersectionStride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raydatatype
func (r_ RayIntersector) RayDataType() RayDataType /* not a class type */ {
	rv := objc.Send[RayDataType](r_.ID, objc.Sel("rayDataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raydatatype
func (r_ RayIntersector) SetRayDataType(value RayDataType /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/rayindexdatatype
func (r_ RayIntersector) RayIndexDataType() DataType /* not a class type */ {
	rv := objc.Send[DataType](r_.ID, objc.Sel("rayIndexDataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/rayindexdatatype
func (r_ RayIntersector) SetRayIndexDataType(value DataType /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayIndexDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymask
func (r_ RayIntersector) RayMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("rayMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymask
func (r_ RayIntersector) SetRayMask(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymaskoperator
func (r_ RayIntersector) RayMaskOperator() RayMaskOperator /* not a class type */ {
	rv := objc.Send[RayMaskOperator](r_.ID, objc.Sel("rayMaskOperator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymaskoperator
func (r_ RayIntersector) SetRayMaskOperator(value RayMaskOperator /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayMaskOperator:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymaskoptions
func (r_ RayIntersector) RayMaskOptions() RayMaskOptions /* not a class type */ {
	rv := objc.Send[RayMaskOptions](r_.ID, objc.Sel("rayMaskOptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raymaskoptions
func (r_ RayIntersector) SetRayMaskOptions(value RayMaskOptions /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayMaskOptions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raystride
func (r_ RayIntersector) RayStride() int {
	rv := objc.Send[int](r_.ID, objc.Sel("rayStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/raystride
func (r_ RayIntersector) SetRayStride(value int) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayStride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/triangleintersectiontesttype
func (r_ RayIntersector) TriangleIntersectionTestType() TriangleIntersectionTestType /* not a class type */ {
	rv := objc.Send[TriangleIntersectionTestType](r_.ID, objc.Sel("triangleIntersectionTestType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/triangleintersectiontesttype
func (r_ RayIntersector) SetTriangleIntersectionTestType(value TriangleIntersectionTestType /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTriangleIntersectionTestType:"), value)
}



