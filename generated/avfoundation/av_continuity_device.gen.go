// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVContinuityDevice */


/* debug [class_header]: Header for AVContinuityDevice */
// The class instance for the [ContinuityDevice] class.
var (
	ContinuityDeviceClass     _ContinuityDeviceClass
	ContinuityDeviceClassOnce sync.Once
)

func getContinuityDeviceClass() _ContinuityDeviceClass {
	ContinuityDeviceClassOnce.Do(func() {
		ContinuityDeviceClass = _ContinuityDeviceClass{objc.GetClass("AVContinuityDevice")}
	})
	return ContinuityDeviceClass
}

type _ContinuityDeviceClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ContinuityDevice */
// An interface definition for the [ContinuityDevice] class.
type IContinuityDevice interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ContinuityDevice */
	// properties:
	IsConnected() bool
	SetIsConnected(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ContinuityDevice */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ContinuityDevice */
// Alloc allocates a new instance without initialization.
func (cc _ContinuityDeviceClass) Alloc() ContinuityDevice {
	rv := objc.Send[ContinuityDevice](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ContinuityDeviceClass) New() ContinuityDevice {
	rv := objc.Send[ContinuityDevice](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ContinuityDevice) Init() ContinuityDevice {
	rv := objc.Send[ContinuityDevice](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ContinuityDevice) Autorelease() ContinuityDevice {
	rv := objc.Send[ContinuityDevice](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContinuityDevice creates a new ContinuityDevice instance.
func NewContinuityDevice() ContinuityDevice {
	return getContinuityDeviceClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ContinuityDevice */
// A class that represents a physical iOS device that’s nearby and can provide access to its cameras and microphones.
//
// Each continuity device instance represents another iOS device that’s nearby. Your app can access the other device’s cameras and microphones with its and properties, respectively.


// A class that represents a physical iOS device that’s nearby and can provide access to its cameras and microphones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContinuityDevice
type ContinuityDevice struct {
	objectivec.Object
}

// ContinuityDeviceFrom constructs a [ContinuityDevice] from an unsafe.Pointer.
//
// A class that represents a physical iOS device that’s nearby and can provide access to its cameras and microphones.
func ContinuityDeviceFrom(ptr unsafe.Pointer) ContinuityDevice {
	return ContinuityDevice{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ContinuityDevice *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ContinuityDevice */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ContinuityDevice */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ContinuityDevice */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ContinuityDevice */

// A Boolean value that indicates whether you can use the continuity device because it’s connected to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/isconnected
func (c_ ContinuityDevice) IsConnected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isConnected"))
	return rv
}/* debug [instance_properties/getter]: isConnected */


// A Boolean value that indicates whether you can use the continuity device because it’s connected to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcontinuitydevice/isconnected
func (c_ ContinuityDevice) SetIsConnected(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsConnected:"), value)
}/* debug [instance_properties/setter]: isConnected */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVContinuityDevice */


