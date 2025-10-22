// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Heading] class.
var (
	HeadingClass     _HeadingClass
	HeadingClassOnce sync.Once
)

func getHeadingClass() _HeadingClass {
	HeadingClassOnce.Do(func() {
		HeadingClass = _HeadingClass{objc.GetClass("CLHeading")}
	})
	return HeadingClass
}

type _HeadingClass struct {
	class objc.Class
}

// An interface definition for the [Heading] class.
type IHeading interface {
	objectivec.IObject
	HeadingAccuracy() unsafe.Pointer
	MagneticHeading() unsafe.Pointer
	Timestamp() foundation.NSDate
	TrueHeading() unsafe.Pointer
	X() HeadingComponentValue
	Y() HeadingComponentValue
	Z() HeadingComponentValue
}

// The orientation of the user’s device, relative to true or magnetic north.
//
// A object contains computed values for the device’s azimuth (orientation) relative to true or magnetic north. It also includes the raw data for the three-dimensional vector used to compute those values. A navigation app might use the information to rotate a map so that it reflects the direction that the user is facing. Typically, you don’t create instances of this class yourself, nor do you subclass it. Instead, you receive instances of this class through the delegate assigned to the object whose method you called.


// The orientation of the user’s device, relative to true or magnetic north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLHeading

type Heading struct {
	objectivec.Object
}

// HeadingFrom constructs a [Heading] from an unsafe.Pointer.
//
// The orientation of the user’s device, relative to true or magnetic north.
func HeadingFrom(ptr unsafe.Pointer) Heading {
	return Heading{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (hc _HeadingClass) Alloc() Heading {
	rv := objc.Send[Heading](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (hc _HeadingClass) New() Heading {
	rv := objc.Send[Heading](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ Heading) Init() Heading {
	rv := objc.Send[Heading](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ Heading) Autorelease() Heading {
	rv := objc.Send[Heading](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewHeading creates a new Heading instance.
func NewHeading() Heading {
	return getHeadingClass().New()
}



// The maximum deviation (measured in degrees) between the reported heading and the true geomagnetic heading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLHeading/headingAccuracy

func (h_ Heading) HeadingAccuracy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("headingAccuracy"))
	return rv
}


// The heading (measured in degrees) relative to magnetic north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLHeading/magneticHeading

func (h_ Heading) MagneticHeading() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("magneticHeading"))
	return rv
}


// The time at which this heading was determined.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLHeading/timestamp

func (h_ Heading) Timestamp() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](h_.ID, objc.Sel("timestamp"))
	return rv
}


// The heading (measured in degrees) relative to true north.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLHeading/trueHeading

func (h_ Heading) TrueHeading() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](h_.ID, objc.Sel("trueHeading"))
	return rv
}


// The geomagnetic data (measured in microteslas) for the x-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLHeading/x

func (h_ Heading) X() HeadingComponentValue {
	rv := objc.Send[HeadingComponentValue](h_.ID, objc.Sel("x"))
	return rv
}


// The geomagnetic data (measured in microteslas) for the y-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLHeading/y

func (h_ Heading) Y() HeadingComponentValue {
	rv := objc.Send[HeadingComponentValue](h_.ID, objc.Sel("y"))
	return rv
}


// The geomagnetic data (measured in microteslas) for the z-axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLHeading/z

func (h_ Heading) Z() HeadingComponentValue {
	rv := objc.Send[HeadingComponentValue](h_.ID, objc.Sel("z"))
	return rv
}



