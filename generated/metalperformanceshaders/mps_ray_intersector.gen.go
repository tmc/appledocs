// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	CullMode() CullMode get set /* not a class type */
	SetCullMode(value CullMode get set /* not a class type */)
	FrontFacingWinding() Winding get set /* not a class type */
	SetFrontFacingWinding(value Winding get set /* not a class type */)
	IntersectionDataType() IntersectionDataType get set /* not a class type */
	SetIntersectionDataType(value IntersectionDataType get set /* not a class type */)
	IntersectionStride() objectivec.IObject
	SetIntersectionStride(value objectivec.IObject)
	RayDataType() RayDataType get set /* not a class type */
	SetRayDataType(value RayDataType get set /* not a class type */)
	RayMaskOptions() RayMaskOptions get set /* not a class type */
	SetRayMaskOptions(value RayMaskOptions get set /* not a class type */)
	RayStride() objectivec.IObject
	SetRayStride(value objectivec.IObject)
	BoundingBoxIntersectionTestType() BoundingBoxIntersectionTestType get set /* not a class type */
	SetBoundingBoxIntersectionTestType(value BoundingBoxIntersectionTestType get set /* not a class type */)
	TriangleIntersectionTestType() TriangleIntersectionTestType get set /* not a class type */
	SetTriangleIntersectionTestType(value TriangleIntersectionTestType get set /* not a class type */)
	RayIndexDataType() DataType get set /* not a class type */
	SetRayIndexDataType(value DataType get set /* not a class type */)
	RayMask() objectivec.IObject
	SetRayMask(value objectivec.IObject)
	RayMaskOperator() RayMaskOperator get set /* not a class type */
	SetRayMaskOperator(value RayMaskOperator get set /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _RayIntersectorClass) Alloc() RayIntersector {
	rv := objc.Send[RayIntersector](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998438-initwithcoder
func NewRayIntersectorWithCoderDevice(aDecoder foundation.Coder, device unsafe.Pointer) RayIntersector {
	instance := getRayIntersectorClass().Alloc()
	rv := objc.Send[RayIntersector](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998439-initwithdevice
func NewRayIntersectorWithDevice(device unsafe.Pointer) RayIntersector {
	instance := getRayIntersectorClass().Alloc()
	rv := objc.Send[RayIntersector](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}






















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998433-cullmode
func (r_ RayIntersector) CullMode() CullMode get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("cullMode"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998433-cullmode
func (r_ RayIntersector) SetCullMode(value CullMode get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setCullMode:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998437-frontfacingwinding
func (r_ RayIntersector) FrontFacingWinding() Winding get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("frontFacingWinding"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998437-frontfacingwinding
func (r_ RayIntersector) SetFrontFacingWinding(value Winding get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setFrontFacingWinding:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998440-intersectiondatatype
func (r_ RayIntersector) IntersectionDataType() IntersectionDataType get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("intersectionDataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998440-intersectiondatatype
func (r_ RayIntersector) SetIntersectionDataType(value IntersectionDataType get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIntersectionDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998441-intersectionstride
func (r_ RayIntersector) IntersectionStride() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("intersectionStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998441-intersectionstride
func (r_ RayIntersector) SetIntersectionStride(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIntersectionStride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998443-raydatatype
func (r_ RayIntersector) RayDataType() RayDataType get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("rayDataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998443-raydatatype
func (r_ RayIntersector) SetRayDataType(value RayDataType get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998444-raymaskoptions
func (r_ RayIntersector) RayMaskOptions() RayMaskOptions get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("rayMaskOptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998444-raymaskoptions
func (r_ RayIntersector) SetRayMaskOptions(value RayMaskOptions get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayMaskOptions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998445-raystride
func (r_ RayIntersector) RayStride() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("rayStride"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/2998445-raystride
func (r_ RayIntersector) SetRayStride(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayStride:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/3013799-boundingboxintersectiontesttype
func (r_ RayIntersector) BoundingBoxIntersectionTestType() BoundingBoxIntersectionTestType get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("boundingBoxIntersectionTestType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/3013799-boundingboxintersectiontesttype
func (r_ RayIntersector) SetBoundingBoxIntersectionTestType(value BoundingBoxIntersectionTestType get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setBoundingBoxIntersectionTestType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/3013800-triangleintersectiontesttype
func (r_ RayIntersector) TriangleIntersectionTestType() TriangleIntersectionTestType get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("triangleIntersectionTestType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/3013800-triangleintersectiontesttype
func (r_ RayIntersector) SetTriangleIntersectionTestType(value TriangleIntersectionTestType get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setTriangleIntersectionTestType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/3131884-rayindexdatatype
func (r_ RayIntersector) RayIndexDataType() DataType get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("rayIndexDataType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/3131884-rayindexdatatype
func (r_ RayIntersector) SetRayIndexDataType(value DataType get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayIndexDataType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/3152591-raymask
func (r_ RayIntersector) RayMask() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](r_.ID, objc.Sel("rayMask"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/3152591-raymask
func (r_ RayIntersector) SetRayMask(value objectivec.IObject) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayMask:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/3242876-raymaskoperator
func (r_ RayIntersector) RayMaskOperator() RayMaskOperator get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("rayMaskOperator"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrayintersector/3242876-raymaskoperator
func (r_ RayIntersector) SetRayMaskOperator(value RayMaskOperator get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRayMaskOperator:"), value)
}







