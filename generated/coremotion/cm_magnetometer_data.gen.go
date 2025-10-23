// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MagnetometerData] class.
var (
	MagnetometerDataClass     _MagnetometerDataClass
	MagnetometerDataClassOnce sync.Once
)

func getMagnetometerDataClass() _MagnetometerDataClass {
	MagnetometerDataClassOnce.Do(func() {
		MagnetometerDataClass = _MagnetometerDataClass{objc.GetClass("CMMagnetometerData")}
	})
	return MagnetometerDataClass
}

type _MagnetometerDataClass struct {
	class objc.Class
}

// An interface definition for the [MagnetometerData] class.
type IMagnetometerData interface {
	ILogItem
	// properties:
	MagneticField() MagneticField /* not a class type */
	MagnetometerData() ICMMagnetometerData
	SetMagnetometerData(value ICMMagnetometerData)
	// methods:
}

// Measurements of the Earth’s magnetic field relative to the device.
//
// Your application can obtain samples of magnetometer measurements, as represented by instances of this class, from the block handler of the method or from the property of the class.


// Measurements of the Earth’s magnetic field relative to the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagnetometerData
type MagnetometerData struct {
	LogItem
}

// MagnetometerDataFrom constructs a [MagnetometerData] from an unsafe.Pointer.
//
// Measurements of the Earth’s magnetic field relative to the device.
func MagnetometerDataFrom(ptr unsafe.Pointer) MagnetometerData {
	return MagnetometerData{
		LogItem: LogItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MagnetometerDataClass) Alloc() MagnetometerData {
	rv := objc.Send[MagnetometerData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MagnetometerDataClass) New() MagnetometerData {
	rv := objc.Send[MagnetometerData](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MagnetometerData) Init() MagnetometerData {
	rv := objc.Send[MagnetometerData](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MagnetometerData) Autorelease() MagnetometerData {
	rv := objc.Send[MagnetometerData](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMagnetometerData creates a new MagnetometerData instance.
func NewMagnetometerData() MagnetometerData {
	return getMagnetometerDataClass().New()
}



// Returns the magnetic field measured by the magnetometer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagnetometerData/magneticField
func (m_ MagnetometerData) MagneticField() MagneticField /* not a class type */ {
	rv := objc.Send[MagneticField](m_.ID, objc.Sel("magneticField"))
	return rv
}


// The latest sample of magnetometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/magnetometerdata
func (m_ MagnetometerData) MagnetometerData() ICMMagnetometerData {
	rv := objc.Send[MagnetometerData](m_.ID, objc.Sel("magnetometerData"))
	return rv
}


// The latest sample of magnetometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/magnetometerdata
func (m_ MagnetometerData) SetMagnetometerData(value ICMMagnetometerData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMagnetometerData:"), value)
}



