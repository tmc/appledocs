// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [RasterizationRateSampleArray] class.
type IRasterizationRateSampleArray interface {
	objectivec.IObject
	

	// properties:
	Horizontal() IMTLRasterizationRateSampleArray
	SetHorizontal(value IMTLRasterizationRateSampleArray)
	MaxSampleCount() MTLSize
	SetMaxSampleCount(value MTLSize)
	SampleCount() MTLSize
	SetSampleCount(value MTLSize)
	Vertical() IMTLRasterizationRateSampleArray
	SetVertical(value IMTLRasterizationRateSampleArray)


	

	// methods:
	ObjectAtIndexedSubscript(index uint) foundation.Number
	SetObjectAtIndexedSubscript(value foundation.foundation.INSNumber, index uint)


}





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




















// Retrieves the sample value at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateSampleArray/objectAtIndexedSubscript:
func (r_ RasterizationRateSampleArray) ObjectAtIndexedSubscript(index uint) foundation.Number {
	rv := objc.Send[foundation.Number](r_.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}


// Stores a sample value at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRasterizationRateSampleArray/setObject:atIndexedSubscript:
func (r_ RasterizationRateSampleArray) SetObjectAtIndexedSubscript(value foundation.foundation.INSNumber, index uint) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setObject:atIndexedSubscript:"), value, index)
}







// The horizontal rasterization rates for the layer map’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/horizontal
func (r_ RasterizationRateSampleArray) Horizontal() IMTLRasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](r_.ID, objc.Sel("horizontal"))
	return rv
}


// The horizontal rasterization rates for the layer map’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/horizontal
func (r_ RasterizationRateSampleArray) SetHorizontal(value IMTLRasterizationRateSampleArray) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHorizontal:"), value)
}


// The maximum number of rows and columns in the layer map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/maxsamplecount
func (r_ RasterizationRateSampleArray) MaxSampleCount() MTLSize {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("maxSampleCount"))
	return rv
}


// The maximum number of rows and columns in the layer map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/maxsamplecount
func (r_ RasterizationRateSampleArray) SetMaxSampleCount(value MTLSize) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxSampleCount:"), value)
}


// The number of rows and columns in the layer map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/samplecount
func (r_ RasterizationRateSampleArray) SampleCount() MTLSize {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("sampleCount"))
	return rv
}


// The number of rows and columns in the layer map.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/samplecount
func (r_ RasterizationRateSampleArray) SetSampleCount(value MTLSize) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSampleCount:"), value)
}


// The vertical rasterization rates for the layer map’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/vertical
func (r_ RasterizationRateSampleArray) Vertical() IMTLRasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](r_.ID, objc.Sel("vertical"))
	return rv
}


// The vertical rasterization rates for the layer map’s rows.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/vertical
func (r_ RasterizationRateSampleArray) SetVertical(value IMTLRasterizationRateSampleArray) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVertical:"), value)
}








