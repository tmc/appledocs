// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CMMagnetometerData */


/* debug [class_header]: Header for CMMagnetometerData */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MagnetometerData */
// An interface definition for the [MagnetometerData] class.
type IMagnetometerData interface {
	ILogItem
	
/* debug [class_interface_properties]: Properties for MagnetometerData */
	// properties:
	MagneticField() objc.IObject /* cross-framework: CMMagneticField */
	MagnetometerData() ICMMagnetometerData
	SetMagnetometerData(value ICMMagnetometerData)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MagnetometerData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MagnetometerData */
// Alloc allocates a new instance without initialization.
func (mc _MagnetometerDataClass) Alloc() MagnetometerData {
	rv := objc.Send[MagnetometerData](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MagnetometerData */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MagnetometerData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MagnetometerData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MagnetometerData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MagnetometerData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MagnetometerData */

// Returns the magnetic field measured by the magnetometer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMMagnetometerData/magneticField
func (m_ MagnetometerData) MagneticField() objc.IObject /* cross-framework: CMMagneticField */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("magneticField"))
	return rv
}/* debug [instance_properties/getter]: magneticField */


// The latest sample of magnetometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/magnetometerdata
func (m_ MagnetometerData) MagnetometerData() ICMMagnetometerData {
	rv := objc.Send[MagnetometerData](m_.ID, objc.Sel("magnetometerData"))
	return rv
}/* debug [instance_properties/getter]: magnetometerData */


// The latest sample of magnetometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/magnetometerdata
func (m_ MagnetometerData) SetMagnetometerData(value ICMMagnetometerData) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMagnetometerData:"), value)
}/* debug [instance_properties/setter]: magnetometerData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMMagnetometerData */



