// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RotationRateData] class.
var (
	RotationRateDataClass     _RotationRateDataClass
	RotationRateDataClassOnce sync.Once
)

func getRotationRateDataClass() _RotationRateDataClass {
	RotationRateDataClassOnce.Do(func() {
		RotationRateDataClass = _RotationRateDataClass{objc.GetClass("CMRotationRateData")}
	})
	return RotationRateDataClass
}

type _RotationRateDataClass struct {
	class objc.Class
}

// An interface definition for the [RotationRateData] class.
type IRotationRateData interface {
	ILogItem
	// properties:
	RotationRate() RotationRate /* not a class type */
	SetRotationRate(value RotationRate /* not a class type */)
	// methods:
}

// A data object that contains a single rotation-rate measurement.


// A data object that contains a single rotation-rate measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRotationRateData
type RotationRateData struct {
	LogItem
}

// RotationRateDataFrom constructs a [RotationRateData] from an unsafe.Pointer.
//
// A data object that contains a single rotation-rate measurement.
func RotationRateDataFrom(ptr unsafe.Pointer) RotationRateData {
	return RotationRateData{
		LogItem: LogItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc _RotationRateDataClass) Alloc() RotationRateData {
	rv := objc.Send[RotationRateData](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RotationRateDataClass) New() RotationRateData {
	rv := objc.Send[RotationRateData](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RotationRateData) Init() RotationRateData {
	rv := objc.Send[RotationRateData](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RotationRateData) Autorelease() RotationRateData {
	rv := objc.Send[RotationRateData](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRotationRateData creates a new RotationRateData instance.
func NewRotationRateData() RotationRateData {
	return getRotationRateDataClass().New()
}



// The rotation rate as measured by the device’s gyroscope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmgyrodata/rotationrate
func (r_ RotationRateData) RotationRate() RotationRate /* not a class type */ {
	rv := objc.Send[RotationRate](r_.ID, objc.Sel("rotationRate"))
	return rv
}


// The rotation rate as measured by the device’s gyroscope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmgyrodata/rotationrate
func (r_ RotationRateData) SetRotationRate(value RotationRate /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRotationRate:"), value)
}



