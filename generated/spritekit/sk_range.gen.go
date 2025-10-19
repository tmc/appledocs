// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKRange] class.
var sKRangeClass = _SKRangeClass{objc.GetClass("SKRange")}

type _SKRangeClass struct {
	class objc.Class
}

// An interface definition for the [SKRange] class.
type ISKRange interface {
	objectivec.IObject
}

// A definition of a range of floating-point values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange

type SKRange struct {
	objectivec.Object
}

// SKRangeFrom constructs a [SKRange] from an unsafe.Pointer.
//
// A definition of a range of floating-point values.
func SKRangeFrom(ptr unsafe.Pointer) SKRange {
	return SKRange{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (sc _SKRangeClass) Alloc() SKRange {
	rv := objc.Send[SKRange](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _SKRangeClass) New() SKRange {
	rv := objc.Send[SKRange](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKRange) Init() SKRange {
	rv := objc.Send[SKRange](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKRange) Autorelease() SKRange {
	rv := objc.Send[SKRange](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKRange creates a new SKRange instance.
func NewSKRange() SKRange {
	return sKRangeClass.New()
}


// Initializes a new range object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/init(lowerLimit:upperLimit:)
func NewSKRangeWithLowerLimitUpperLimit(lower float64, upper float64) SKRange {
	instance := sKRangeClass.Alloc()
	rv := objc.Send[SKRange](instance.ID, objc.Sel("initWithLowerLimit:upperLimit:"), lower, upper)
	rv.Autorelease()
	return rv
}
// Creates and initializes a new range object that specifies only a maximum value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/init(upperLimit:)
func NewRangeWithUpperLimit(upper float64) SKRange {
	rv := objc.Send[SKRange](objc.ID(sKRangeClass.class), objc.Sel("rangeWithUpperLimit:"), upper)
	rv.Autorelease()
	return rv
}
// Creates and initializes a new range object using a value and a maximum distance from that value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/init(value:variance:)
func NewRangeWithValueVariance(value float64, variance float64) SKRange {
	rv := objc.Send[SKRange](objc.ID(sKRangeClass.class), objc.Sel("rangeWithValue:variance:"), value, variance)
	rv.Autorelease()
	return rv
}
// Creates and initializes a new range object that specifies a constant value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/init(constantValue:)
func NewRangeWithConstantValue(value float64) SKRange {
	rv := objc.Send[SKRange](objc.ID(sKRangeClass.class), objc.Sel("rangeWithConstantValue:"), value)
	rv.Autorelease()
	return rv
}
// Creates and initializes a new range object that specifies only a minimum value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/init(lowerLimit:)
func NewRangeWithLowerLimit(lower float64) SKRange {
	rv := objc.Send[SKRange](objc.ID(sKRangeClass.class), objc.Sel("rangeWithLowerLimit:"), lower)
	rv.Autorelease()
	return rv
}


// Creates and initializes a new range object that specifies a constant value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/init(constantValue:)
func (sc _SKRangeClass) RangeWithConstantValue(value float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("rangeWithConstantValue:"), value)
	return rv
}
// Creates and initializes a new range object that specifies only a minimum value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/init(lowerLimit:)
func (sc _SKRangeClass) RangeWithLowerLimit(lower float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("rangeWithLowerLimit:"), lower)
	return rv
}
// Creates and initializes a new range object that specifies only a maximum value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/init(upperLimit:)
func (sc _SKRangeClass) RangeWithUpperLimit(upper float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("rangeWithUpperLimit:"), upper)
	return rv
}
// Creates and initializes a new range object using a value and a maximum distance from that value. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/init(value:variance:)
func (sc _SKRangeClass) RangeWithValueVariance(value float64, variance float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("rangeWithValue:variance:"), value, variance)
	return rv
}
// Creates and initializes a new range object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/rangeWithLowerLimit:upperLimit:
func (sc _SKRangeClass) RangeWithLowerLimitUpperLimit(lower float64, upper float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("rangeWithLowerLimit:upperLimit:"), lower, upper)
	return rv
}
// Creates and initializes a new range object that encompasses all possible values. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKRange/withNoLimits()
func (sc _SKRangeClass) RangeWithNoLimits() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("rangeWithNoLimits"))
	return rv
}

