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
}

// An array instance that contains rasterization rates.
//
// The and properties of an point to instances that contains rasterization rates for the layer map. You can use array subscript syntax to access the samples. instances perform bounds checking on any accesses you make to their sample data.
//
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

// Alloc allocates a new instance without initialization.
func (rc _RasterizationRateSampleArrayClass) Alloc() RasterizationRateSampleArray {
	rv := objc.Send[RasterizationRateSampleArray](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The horizontal rasterization rates for the layer map’s rows.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/horizontal
func (r_ RasterizationRateSampleArray) Horizontal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("horizontal"))
	return rv
}


// SetHorizontal sets the value of the horizontal property.
// The horizontal rasterization rates for the layer map’s rows.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/horizontal
func (r_ RasterizationRateSampleArray) SetHorizontal(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHorizontal:"), value)
}

// The maximum number of rows and columns in the layer map.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/maxsamplecount
func (r_ RasterizationRateSampleArray) MaxSampleCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("maxSampleCount"))
	return rv
}


// SetMaxSampleCount sets the value of the maxSampleCount property.
// The maximum number of rows and columns in the layer map.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/maxsamplecount
func (r_ RasterizationRateSampleArray) SetMaxSampleCount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMaxSampleCount:"), value)
}

// The number of rows and columns in the layer map.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/samplecount
func (r_ RasterizationRateSampleArray) SampleCount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("sampleCount"))
	return rv
}


// SetSampleCount sets the value of the sampleCount property.
// The number of rows and columns in the layer map.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/samplecount
func (r_ RasterizationRateSampleArray) SetSampleCount(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSampleCount:"), value)
}

// The vertical rasterization rates for the layer map’s rows.
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/vertical
func (r_ RasterizationRateSampleArray) Vertical() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("vertical"))
	return rv
}


// SetVertical sets the value of the vertical property.
// The vertical rasterization rates for the layer map’s rows.

//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrasterizationratelayerdescriptor/vertical
func (r_ RasterizationRateSampleArray) SetVertical(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setVertical:"), value)
}



