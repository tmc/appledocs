// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFSafariExtensionManager */


/* debug [class_header]: Header for SFSafariExtensionManager */
// The class instance for the [SFSafariExtensionManager] class.
var (
	SFSafariExtensionManagerClass     _SFSafariExtensionManagerClass
	SFSafariExtensionManagerClassOnce sync.Once
)

func getSFSafariExtensionManagerClass() _SFSafariExtensionManagerClass {
	SFSafariExtensionManagerClassOnce.Do(func() {
		SFSafariExtensionManagerClass = _SFSafariExtensionManagerClass{objc.GetClass("SFSafariExtensionManager")}
	})
	return SFSafariExtensionManagerClass
}

type _SFSafariExtensionManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFSafariExtensionManager */
// An interface definition for the [SFSafariExtensionManager] class.
type ISFSafariExtensionManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFSafariExtensionManager */
	// properties:
	SFExtensionProfileKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFSafariExtensionManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFSafariExtensionManager */
// Alloc allocates a new instance without initialization.
func (sc _SFSafariExtensionManagerClass) Alloc() SFSafariExtensionManager {
	rv := objc.Send[SFSafariExtensionManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFSafariExtensionManagerClass) New() SFSafariExtensionManager {
	rv := objc.Send[SFSafariExtensionManager](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFSafariExtensionManager) Init() SFSafariExtensionManager {
	rv := objc.Send[SFSafariExtensionManager](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFSafariExtensionManager) Autorelease() SFSafariExtensionManager {
	rv := objc.Send[SFSafariExtensionManager](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSafariExtensionManager creates a new SFSafariExtensionManager instance.
func NewSFSafariExtensionManager() SFSafariExtensionManager {
	return getSFSafariExtensionManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFSafariExtensionManager */
// A class that your app uses to find out the current state of a Safari app extension.


// A class that your app uses to find out the current state of a Safari app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionManager
type SFSafariExtensionManager struct {
	objectivec.Object
}

// SFSafariExtensionManagerFrom constructs a [SFSafariExtensionManager] from an unsafe.Pointer.
//
// A class that your app uses to find out the current state of a Safari app extension.
func SFSafariExtensionManagerFrom(ptr unsafe.Pointer) SFSafariExtensionManager {
	return SFSafariExtensionManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFSafariExtensionManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFSafariExtensionManager */

// Gets the current state of the Safari app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFSafariExtensionManager/getStateOfSafariExtension(withIdentifier:completionHandler:)
func (sc _SFSafariExtensionManagerClass) GetStateOfSafariExtensionWithIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler func(unsafe.Pointer, unsafe.Pointer)) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getStateOfSafariExtensionWithIdentifier:completionHandler:"), identifier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetStateOfSafariExtensionWithIdentifierCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFSafariExtensionManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFSafariExtensionManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFSafariExtensionManager */

// A string the system uses as a key in a user info dictionary to identify a profile identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/safariservices/sfextensionprofilekey
func (s_ SFSafariExtensionManager) SFExtensionProfileKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("SFExtensionProfileKey"))
	return rv
}/* debug [instance_properties/getter]: SFExtensionProfileKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFSafariExtensionManager */



