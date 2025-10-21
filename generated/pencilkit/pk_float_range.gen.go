// Code generated from Apple documentation for PencilKit. DO NOT EDIT.

package pencilkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FloatRange] class.
var (
	FloatRangeClass     _FloatRangeClass
	FloatRangeClassOnce sync.Once
)

func getFloatRangeClass() _FloatRangeClass {
	FloatRangeClassOnce.Do(func() {
		FloatRangeClass = _FloatRangeClass{objc.GetClass("PKFloatRange")}
	})
	return FloatRangeClass
}

type _FloatRangeClass struct {
	class objc.Class
}

// An interface definition for the [FloatRange] class.
type IFloatRange interface {
	objectivec.IObject
}

// A utility class that represents range components of a stroke.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKFloatRange
type FloatRange struct {
	objectivec.Object
}

// FloatRangeFrom constructs a [FloatRange] from an unsafe.Pointer.
//
// A utility class that represents range components of a stroke.
func FloatRangeFrom(ptr unsafe.Pointer) FloatRange {
	return FloatRange{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FloatRangeClass) Alloc() FloatRange {
	rv := objc.Send[FloatRange](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FloatRangeClass) New() FloatRange {
	rv := objc.Send[FloatRange](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FloatRange) Init() FloatRange {
	rv := objc.Send[FloatRange](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FloatRange) Autorelease() FloatRange {
	rv := objc.Send[FloatRange](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFloatRange creates a new FloatRange instance.
func NewFloatRange() FloatRange {
	return getFloatRangeClass().New()
}


// A utility class used to contain ranges returned by the PKStroke API.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKFloatRange/initWithLowerBound:upperBound:
func NewFloatRangeWithLowerBoundUpperBound(lowerBound float64, upperBound float64) FloatRange {
	instance := getFloatRangeClass().Alloc()
	rv := objc.Send[FloatRange](instance.ID, objc.Sel("initWithLowerBound:upperBound:"), lowerBound, upperBound)
	rv.Autorelease()
	return rv
}


// A floating point value that represents the lower bound of the range.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKFloatRange/lowerBound
func (f_ FloatRange) LowerBound() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("lowerBound"))
	return rv
}

// A floating point value that represents the upper bound of the range.
//
// [Full Topic]: https://developer.apple.com/documentation/PencilKit/PKFloatRange/upperBound
func (f_ FloatRange) UpperBound() float64 {
	rv := objc.Send[float64](f_.ID, objc.Sel("upperBound"))
	return rv
}


