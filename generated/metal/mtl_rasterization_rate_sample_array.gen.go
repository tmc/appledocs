// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTLRasterizationRateSampleArray */


/* debug [class_header]: Header for MTLRasterizationRateSampleArray */
// The class instance for the [RasterizationRateSampleArray] class.
var (
	RasterizationRateSampleArrayClass     _RasterizationRateSampleArrayClass
	RasterizationRateSampleArrayClassOnce sync.Once
)

func getRasterizationRateSampleArrayClass() _RasterizationRateSampleArrayClass {
	RasterizationRateSampleArrayClassOnce.Do(func() {
		RasterizationRateSampleArrayClass = _RasterizationRateSampleArrayClass{objc.GetClass("MTLRasterizationRateSampleArray")}
	})
	return RasterizationRateSampleArrayClass
}

type _RasterizationRateSampleArrayClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RasterizationRateSampleArray */
// An interface definition for the [RasterizationRateSampleArray] class.
type IRasterizationRateSampleArray interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for RasterizationRateSampleArray */
	// properties:
	Horizontal() IMTLRasterizationRateSampleArray
	SetHorizontal(value IMTLRasterizationRateSampleArray)
	MaxSampleCount() objc.IObject /* cross-framework: MTLSize */
	SetMaxSampleCount(value objc.IObject /* cross-framework: MTLSize */)
	SampleCount() objc.IObject /* cross-framework: MTLSize */
	SetSampleCount(value objc.IObject /* cross-framework: MTLSize */)
	Vertical() IMTLRasterizationRateSampleArray
	SetVertical(value IMTLRasterizationRateSampleArray)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RasterizationRateSampleArray */
	// methods:
	ObjectAtIndexedSubscript(index uint) foundation.Number
	SetObjectAtIndexedSubscript(value objc.IObject /* cross-framework: NSNumber */, index uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RasterizationRateSampleArray */
// Alloc allocates a new instance without initialization.
func (rc _RasterizationRateSampleArrayClass) Alloc() RasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RasterizationRateSampleArrayClass) New() RasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RasterizationRateSampleArray) Init() RasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RasterizationRateSampleArray) Autorelease() RasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRasterizationRateSampleArray creates a new RasterizationRateSampleArray instance.
func NewRasterizationRateSampleArray() RasterizationRateSampleArray {
	return getRasterizationRateSampleArrayClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RasterizationRateSampleArray */
// An array instance that contains rasterization rates.
//
// The and properties of an point to instances that contains rasterization rates for the layer map. You can use array subscript syntax to access the samples. instances perform bounds checking on any accesses you make to their sample data.


// An array instance that contains rasterization rates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateSampleArray
type RasterizationRateSampleArray struct {
	objectivec.Object
}

// RasterizationRateSampleArrayFrom constructs a [RasterizationRateSampleArray] from an unsafe.Pointer.
//
// An array instance that contains rasterization rates.
func RasterizationRateSampleArrayFrom(ptr unsafe.Pointer) RasterizationRateSampleArray {
	return RasterizationRateSampleArray{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RasterizationRateSampleArray *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RasterizationRateSampleArray */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RasterizationRateSampleArray */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RasterizationRateSampleArray */

// Retrieves the sample value at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateSampleArray/objectAtIndexedSubscript:
func (r_ RasterizationRateSampleArray) ObjectAtIndexedSubscript(index uint) foundation.Number {
	rv := objc.Send[foundation.Number](r_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}/* debug [instance_methods/method]: ObjectAtIndexedSubscript */


// Stores a sample value at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateSampleArray/setObject:atIndexedSubscript:
func (r_ RasterizationRateSampleArray) SetObjectAtIndexedSubscript(value objc.IObject /* cross-framework: NSNumber */, index uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setObject:atIndexedSubscript:"), value, index)
}/* debug [instance_methods/method]: SetObjectAtIndexedSubscript */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RasterizationRateSampleArray */

// The horizontal rasterization rates for the layer map’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/horizontal
func (r_ RasterizationRateSampleArray) Horizontal() IMTLRasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](r_.ID, objc.Sel("horizontal"))
	return rv
}/* debug [instance_properties/getter]: horizontal */


// The horizontal rasterization rates for the layer map’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/horizontal
func (r_ RasterizationRateSampleArray) SetHorizontal(value IMTLRasterizationRateSampleArray) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHorizontal:"), value)
}/* debug [instance_properties/setter]: horizontal */


// The maximum number of rows and columns in the layer map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/maxsamplecount
func (r_ RasterizationRateSampleArray) MaxSampleCount() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("maxSampleCount"))
	return rv
}/* debug [instance_properties/getter]: maxSampleCount */


// The maximum number of rows and columns in the layer map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/maxsamplecount
func (r_ RasterizationRateSampleArray) SetMaxSampleCount(value objc.IObject /* cross-framework: MTLSize */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxSampleCount:"), value)
}/* debug [instance_properties/setter]: maxSampleCount */


// The number of rows and columns in the layer map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/samplecount
func (r_ RasterizationRateSampleArray) SampleCount() objc.IObject /* cross-framework: MTLSize */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("sampleCount"))
	return rv
}/* debug [instance_properties/getter]: sampleCount */


// The number of rows and columns in the layer map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/samplecount
func (r_ RasterizationRateSampleArray) SetSampleCount(value objc.IObject /* cross-framework: MTLSize */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSampleCount:"), value)
}/* debug [instance_properties/setter]: sampleCount */


// The vertical rasterization rates for the layer map’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/vertical
func (r_ RasterizationRateSampleArray) Vertical() IMTLRasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](r_.ID, objc.Sel("vertical"))
	return rv
}/* debug [instance_properties/getter]: vertical */


// The vertical rasterization rates for the layer map’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/vertical
func (r_ RasterizationRateSampleArray) SetVertical(value IMTLRasterizationRateSampleArray) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVertical:"), value)
}/* debug [instance_properties/setter]: vertical */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTLRasterizationRateSampleArray */



