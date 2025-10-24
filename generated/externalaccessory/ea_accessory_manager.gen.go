// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EAAccessoryManager */


/* debug [class_header]: Header for EAAccessoryManager */
// The class instance for the [EAAccessoryManager] class.
var (
	EAAccessoryManagerClass     _EAAccessoryManagerClass
	EAAccessoryManagerClassOnce sync.Once
)

func getEAAccessoryManagerClass() _EAAccessoryManagerClass {
	EAAccessoryManagerClassOnce.Do(func() {
		EAAccessoryManagerClass = _EAAccessoryManagerClass{objc.GetClass("EAAccessoryManager")}
	})
	return EAAccessoryManagerClass
}

type _EAAccessoryManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EAAccessoryManager */
// An interface definition for the [EAAccessoryManager] class.
type IEAAccessoryManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EAAccessoryManager */
	// properties:
	ConnectedAccessories() []EAAccessory
	EAAccessoryKey() objc.IObject /* cross-framework: NSString */
	EAAccessorySelectedKey() objc.IObject /* cross-framework: NSString */
	EABluetoothAccessoryPickerErrorDomain() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EAAccessoryManager */
	// methods:
	RegisterForLocalNotifications()
	UnregisterForLocalNotifications()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EAAccessoryManager */
// Alloc allocates a new instance without initialization.
func (ec _EAAccessoryManagerClass) Alloc() EAAccessoryManager {
	rv := objc.Send[EAAccessoryManager](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EAAccessoryManagerClass) New() EAAccessoryManager {
	rv := objc.Send[EAAccessoryManager](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EAAccessoryManager) Init() EAAccessoryManager {
	rv := objc.Send[EAAccessoryManager](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EAAccessoryManager) Autorelease() EAAccessoryManager {
	rv := objc.Send[EAAccessoryManager](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEAAccessoryManager creates a new EAAccessoryManager instance.
func NewEAAccessoryManager() EAAccessoryManager {
	return getEAAccessoryManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EAAccessoryManager */
// The object you use to identify connected accessories, and begin delivery of connection and disconnection notifications.
//
// An object coordinates the attached accessories for an iOS-based device. Use the shared accessory manager to retrieve a list of connected accessories, and start and stop the delivery of connection and disconnection notifications.


// The object you use to identify connected accessories, and begin delivery of connection and disconnection notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessoryManager
type EAAccessoryManager struct {
	objectivec.Object
}

// EAAccessoryManagerFrom constructs a [EAAccessoryManager] from an unsafe.Pointer.
//
// The object you use to identify connected accessories, and begin delivery of connection and disconnection notifications.
func EAAccessoryManagerFrom(ptr unsafe.Pointer) EAAccessoryManager {
	return EAAccessoryManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EAAccessoryManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EAAccessoryManager */

// Returns the shared accessory manager object for the iOS-based device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessoryManager/shared()
func (ec _EAAccessoryManagerClass) SharedAccessoryManager() EAAccessoryManager {
	rv := objc.Send[EAAccessoryManager](objc.ID(ec.class), objc.Sel("sharedAccessoryManager"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedAccessoryManager) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EAAccessoryManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EAAccessoryManager */

// Begins the delivery of accessory-related notifications to the current application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessoryManager/registerForLocalNotifications()
func (e_ EAAccessoryManager) RegisterForLocalNotifications() {
	objc.Send[objc.ID](e_.ID, objc.Sel("registerForLocalNotifications"))
}/* debug [instance_methods/method]: RegisterForLocalNotifications */


// Stops the delivery of accessory-related notifications to the current application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessoryManager/unregisterForLocalNotifications()
func (e_ EAAccessoryManager) UnregisterForLocalNotifications() {
	objc.Send[objc.ID](e_.ID, objc.Sel("unregisterForLocalNotifications"))
}/* debug [instance_methods/method]: UnregisterForLocalNotifications */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EAAccessoryManager */

// The accessory objects corresponding to the list of currently connected accessories.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAAccessoryManager/connectedAccessories
func (e_ EAAccessoryManager) ConnectedAccessories() []EAAccessory {
	rv := objc.Send[[]EAAccessory](e_.ID, objc.Sel("connectedAccessories"))
	return rv
}/* debug [instance_properties/getter]: connectedAccessories */


// A key that indicates the accessory object whose status changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eaaccessorykey
func (e_ EAAccessoryManager) EAAccessoryKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("EAAccessoryKey"))
	return rv
}/* debug [instance_properties/getter]: EAAccessoryKey */


// A key that indicates the accessory object that the user selected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eaaccessoryselectedkey
func (e_ EAAccessoryManager) EAAccessorySelectedKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("EAAccessorySelectedKey"))
	return rv
}/* debug [instance_properties/getter]: EAAccessorySelectedKey */


// The domain for errors passed to a Bluetooth picker completion block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/externalaccessory/eabluetoothaccessorypickererrordomain
func (e_ EAAccessoryManager) EABluetoothAccessoryPickerErrorDomain() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](e_.ID, objc.Sel("EABluetoothAccessoryPickerErrorDomain"))
	return rv
}/* debug [instance_properties/getter]: EABluetoothAccessoryPickerErrorDomain */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EAAccessoryManager */


