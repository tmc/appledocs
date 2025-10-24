// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKSmartCardSlotManager */


/* debug [class_header]: Header for TKSmartCardSlotManager */
// The class instance for the [TKSmartCardSlotManager] class.
var (
	TKSmartCardSlotManagerClass     _TKSmartCardSlotManagerClass
	TKSmartCardSlotManagerClassOnce sync.Once
)

func getTKSmartCardSlotManagerClass() _TKSmartCardSlotManagerClass {
	TKSmartCardSlotManagerClassOnce.Do(func() {
		TKSmartCardSlotManagerClass = _TKSmartCardSlotManagerClass{objc.GetClass("TKSmartCardSlotManager")}
	})
	return TKSmartCardSlotManagerClass
}

type _TKSmartCardSlotManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardSlotManager */
// An interface definition for the [TKSmartCardSlotManager] class.
type ITKSmartCardSlotManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKSmartCardSlotManager */
	// properties:
	SlotNames() []string
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardSlotManager */
	// methods:
	GetSlotWithNameReply(name objc.IObject /* cross-framework: NSString */, reply unsafe.Pointer)
	SlotNamed(name objc.IObject /* cross-framework: NSString */) ITKSmartCardSlot
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardSlotManager */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardSlotManagerClass) Alloc() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardSlotManagerClass) New() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardSlotManager) Init() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardSlotManager) Autorelease() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardSlotManager creates a new TKSmartCardSlotManager instance.
func NewTKSmartCardSlotManager() TKSmartCardSlotManager {
	return getTKSmartCardSlotManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardSlotManager */
// An interface to all available smart card reader slots.
//
// Get a list of all known smart card reader slots in the system using the property, and access individual slots by name using the method.


// An interface to all available smart card reader slots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager
type TKSmartCardSlotManager struct {
	objectivec.Object
}

// TKSmartCardSlotManagerFrom constructs a [TKSmartCardSlotManager] from an unsafe.Pointer.
//
// An interface to all available smart card reader slots.
func TKSmartCardSlotManagerFrom(ptr unsafe.Pointer) TKSmartCardSlotManager {
	return TKSmartCardSlotManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardSlotManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardSlotManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardSlotManager */

// The shared singleton Smart Card reader slot manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/default
func (tc _TKSmartCardSlotManagerClass) DefaultManager() TKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](objc.ID(tc.class), objc.Sel("defaultManager"))
	return rv
}/* debug [class_properties_class/property]: defaultManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardSlotManager */

// Asynchronously calls a block with a Smart Card reader slot for a specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/getSlot(withName:reply:)
func (t_ TKSmartCardSlotManager) GetSlotWithNameReply(name objc.IObject /* cross-framework: NSString */, reply unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("getSlotWithName:reply:"), name, reply)
}/* debug [instance_methods/method]: GetSlotWithNameReply */


// Returns the Smart Card slot with a given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/slotNamed(_:)
func (t_ TKSmartCardSlotManager) SlotNamed(name objc.IObject /* cross-framework: NSString */) ITKSmartCardSlot {
	rv := objc.Send[TKSmartCardSlot](t_.ID, objc.Sel("slotNamed:"), name)
	return rv
}/* debug [instance_methods/method]: SlotNamed */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardSlotManager */

// The shared singleton Smart Card reader slot manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/default
func (t_ TKSmartCardSlotManager) DefaultManager() ITKSmartCardSlotManager {
	rv := objc.Send[TKSmartCardSlotManager](t_.ID, objc.Sel("defaultManager"))
	return rv
}/* debug [instance_properties/getter]: defaultManager */


// A list of identifiers for all the Smart Card reader slots available to the system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardSlotManager/slotNames
func (t_ TKSmartCardSlotManager) SlotNames() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("slotNames"))
	return rv
}/* debug [instance_properties/getter]: slotNames */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardSlotManager */


