// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKSmartCardTokenRegistrationManager */


/* debug [class_header]: Header for TKSmartCardTokenRegistrationManager */
// The class instance for the [TKSmartCardTokenRegistrationManager] class.
var (
	TKSmartCardTokenRegistrationManagerClass     _TKSmartCardTokenRegistrationManagerClass
	TKSmartCardTokenRegistrationManagerClassOnce sync.Once
)

func getTKSmartCardTokenRegistrationManagerClass() _TKSmartCardTokenRegistrationManagerClass {
	TKSmartCardTokenRegistrationManagerClassOnce.Do(func() {
		TKSmartCardTokenRegistrationManagerClass = _TKSmartCardTokenRegistrationManagerClass{objc.GetClass("TKSmartCardTokenRegistrationManager")}
	})
	return TKSmartCardTokenRegistrationManagerClass
}

type _TKSmartCardTokenRegistrationManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardTokenRegistrationManager */
// An interface definition for the [TKSmartCardTokenRegistrationManager] class.
type ITKSmartCardTokenRegistrationManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKSmartCardTokenRegistrationManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardTokenRegistrationManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardTokenRegistrationManager */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardTokenRegistrationManagerClass) Alloc() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardTokenRegistrationManagerClass) New() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardTokenRegistrationManager) Init() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardTokenRegistrationManager) Autorelease() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardTokenRegistrationManager creates a new TKSmartCardTokenRegistrationManager instance.
func NewTKSmartCardTokenRegistrationManager() TKSmartCardTokenRegistrationManager {
	return getTKSmartCardTokenRegistrationManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardTokenRegistrationManager */
// Provides a centralized management system for registering and unregistering smartcards using their token IDs.
//
// keeps its itself accessible via Keychain and system will automatically invoke an NFC slot when a cryptographic operation is required and asks to provide the registered card.


// Provides a centralized management system for registering and unregistering smartcards using their token IDs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenRegistrationManager
type TKSmartCardTokenRegistrationManager struct {
	objectivec.Object
}

// TKSmartCardTokenRegistrationManagerFrom constructs a [TKSmartCardTokenRegistrationManager] from an unsafe.Pointer.
//
// Provides a centralized management system for registering and unregistering smartcards using their token IDs.
func TKSmartCardTokenRegistrationManagerFrom(ptr unsafe.Pointer) TKSmartCardTokenRegistrationManager {
	return TKSmartCardTokenRegistrationManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardTokenRegistrationManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardTokenRegistrationManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardTokenRegistrationManager */

// Default instance of registration manager
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardTokenRegistrationManager/default
func (tc _TKSmartCardTokenRegistrationManagerClass) DefaultManager() TKSmartCardTokenRegistrationManager {
	rv := objc.Send[TKSmartCardTokenRegistrationManager](objc.ID(tc.class), objc.Sel("defaultManager"))
	return rv
}/* debug [class_properties_class/property]: defaultManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardTokenRegistrationManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardTokenRegistrationManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardTokenRegistrationManager */


