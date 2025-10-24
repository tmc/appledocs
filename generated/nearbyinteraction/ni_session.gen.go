// Code generated from Apple documentation for NearbyInteraction. DO NOT EDIT.

package nearbyinteraction

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/gamekit"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NISession */


/* debug [class_header]: Header for NISession */
// The class instance for the [NISession] class.
var (
	NISessionClass     _NISessionClass
	NISessionClassOnce sync.Once
)

func getNISessionClass() _NISessionClass {
	NISessionClassOnce.Do(func() {
		NISessionClass = _NISessionClass{objc.GetClass("NISession")}
	})
	return NISessionClass
}

type _NISessionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NISession */
// An interface definition for the [NISession] class.
type INISession interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NISession */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NISession */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NISession */
// Alloc allocates a new instance without initialization.
func (nc _NISessionClass) Alloc() NISession {
	rv := objc.Send[NISession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NISessionClass) New() NISession {
	rv := objc.Send[NISession](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NISession) Init() NISession {
	rv := objc.Send[NISession](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NISession) Autorelease() NISession {
	rv := objc.Send[NISession](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNISession creates a new NISession instance.
func NewNISession() NISession {
	return getNISessionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NISession */
// An object that identifies a unique connection between two peer devices.
//
// This class represents the central mechanism to interact with nearby objects, for example, a peer Apple device or third-party accessory. After creating an for a nearby object, the app interacts with the object by receiving callbacks. One session represents an interaction between the user and a single nearby object. To interact with multiple nearby objects, create a separate session for each. For more information, see .


// An object that identifies a unique connection between two peer devices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession
type NISession struct {
	objectivec.Object
}

// NISessionFrom constructs a [NISession] from an unsafe.Pointer.
//
// An object that identifies a unique connection between two peer devices.
func NISessionFrom(ptr unsafe.Pointer) NISession {
	return NISession{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NISession *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NISession */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NISession */

// An object that communicates the device’s supported framework features.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/deviceCapabilities
func (nc _NISessionClass) DeviceCapabilities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(nc.class), objc.Sel("deviceCapabilities"))
	return rv
}/* debug [class_properties_class/property]: deviceCapabilities */

// A Boolean value that indicates whether the device supports basic interaction-session functionality.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NearbyInteraction/NISession/isSupported
func (nc _NISessionClass) Supported() bool {
	rv := objc.Send[bool](objc.ID(nc.class), objc.Sel("supported"))
	return rv
}/* debug [class_properties_class/property]: supported */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NISession */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NISession */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NISession */


