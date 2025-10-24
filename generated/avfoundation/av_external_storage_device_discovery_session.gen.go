// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVExternalStorageDeviceDiscoverySession */


/* debug [class_header]: Header for AVExternalStorageDeviceDiscoverySession */
// The class instance for the [ExternalStorageDeviceDiscoverySession] class.
var (
	ExternalStorageDeviceDiscoverySessionClass     _ExternalStorageDeviceDiscoverySessionClass
	ExternalStorageDeviceDiscoverySessionClassOnce sync.Once
)

func getExternalStorageDeviceDiscoverySessionClass() _ExternalStorageDeviceDiscoverySessionClass {
	ExternalStorageDeviceDiscoverySessionClassOnce.Do(func() {
		ExternalStorageDeviceDiscoverySessionClass = _ExternalStorageDeviceDiscoverySessionClass{objc.GetClass("AVExternalStorageDeviceDiscoverySession")}
	})
	return ExternalStorageDeviceDiscoverySessionClass
}

type _ExternalStorageDeviceDiscoverySessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ExternalStorageDeviceDiscoverySession */
// An interface definition for the [ExternalStorageDeviceDiscoverySession] class.
type IExternalStorageDeviceDiscoverySession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ExternalStorageDeviceDiscoverySession */
	// properties:
	ExternalStorageDevices() []ExternalStorageDevice
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ExternalStorageDeviceDiscoverySession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ExternalStorageDeviceDiscoverySession */
// Alloc allocates a new instance without initialization.
func (ec _ExternalStorageDeviceDiscoverySessionClass) Alloc() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _ExternalStorageDeviceDiscoverySessionClass) New() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ ExternalStorageDeviceDiscoverySession) Init() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ ExternalStorageDeviceDiscoverySession) Autorelease() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewExternalStorageDeviceDiscoverySession creates a new ExternalStorageDeviceDiscoverySession instance.
func NewExternalStorageDeviceDiscoverySession() ExternalStorageDeviceDiscoverySession {
	return getExternalStorageDeviceDiscoverySessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ExternalStorageDeviceDiscoverySession */
// Informs your app when the external storage devices connect to and disconnect from the system.


// Informs your app when the external storage devices connect to and disconnect from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession
type ExternalStorageDeviceDiscoverySession struct {
	objectivec.Object
}

// ExternalStorageDeviceDiscoverySessionFrom constructs a [ExternalStorageDeviceDiscoverySession] from an unsafe.Pointer.
//
// Informs your app when the external storage devices connect to and disconnect from the system.
func ExternalStorageDeviceDiscoverySessionFrom(ptr unsafe.Pointer) ExternalStorageDeviceDiscoverySession {
	return ExternalStorageDeviceDiscoverySession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ExternalStorageDeviceDiscoverySession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ExternalStorageDeviceDiscoverySession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ExternalStorageDeviceDiscoverySession */

// A Boolean value that indicates whether the system supports external storage devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/isSupported
func (ec _ExternalStorageDeviceDiscoverySessionClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(ec.class), objc.Sel("supported"))
	return rv
}/* debug [class_properties_class/property]: supported */

// The system’s singleton device discovery session instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/shared
func (ec _ExternalStorageDeviceDiscoverySessionClass) SharedSession() ExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](objc.ID(ec.class), objc.Sel("sharedSession"))
	return rv
}/* debug [class_properties_class/property]: sharedSession */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ExternalStorageDeviceDiscoverySession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ExternalStorageDeviceDiscoverySession */

// An array of external storage devices the session updates as individual devices connect or disconnect from the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/externalStorageDevices
func (e_ ExternalStorageDeviceDiscoverySession) ExternalStorageDevices() []ExternalStorageDevice {
	rv := objc.Send[[]ExternalStorageDevice](e_.ID, objc.Sel("externalStorageDevices"))
	return rv
}/* debug [instance_properties/getter]: externalStorageDevices */


// A Boolean value that indicates whether the system supports external storage devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/isSupported
func (e_ ExternalStorageDeviceDiscoverySession) Supported() bool {
	rv := objc.Send[bool](e_.ID, objc.Sel("supported"))
	return rv
}/* debug [instance_properties/getter]: supported */


// The system’s singleton device discovery session instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDeviceDiscoverySession/shared
func (e_ ExternalStorageDeviceDiscoverySession) SharedSession() IAVExternalStorageDeviceDiscoverySession {
	rv := objc.Send[ExternalStorageDeviceDiscoverySession](e_.ID, objc.Sel("sharedSession"))
	return rv
}/* debug [instance_properties/getter]: sharedSession */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVExternalStorageDeviceDiscoverySession */



