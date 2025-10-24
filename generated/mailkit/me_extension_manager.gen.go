// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEExtensionManager */


/* debug [class_header]: Header for MEExtensionManager */
// The class instance for the [MEExtensionManager] class.
var (
	MEExtensionManagerClass     _MEExtensionManagerClass
	MEExtensionManagerClassOnce sync.Once
)

func getMEExtensionManagerClass() _MEExtensionManagerClass {
	MEExtensionManagerClassOnce.Do(func() {
		MEExtensionManagerClass = _MEExtensionManagerClass{objc.GetClass("MEExtensionManager")}
	})
	return MEExtensionManagerClass
}

type _MEExtensionManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEExtensionManager */
// An interface definition for the [MEExtensionManager] class.
type IMEExtensionManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEExtensionManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEExtensionManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEExtensionManager */
// Alloc allocates a new instance without initialization.
func (mc _MEExtensionManagerClass) Alloc() MEExtensionManager {
	rv := objc.Send[MEExtensionManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MEExtensionManagerClass) New() MEExtensionManager {
	rv := objc.Send[MEExtensionManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEExtensionManager) Init() MEExtensionManager {
	rv := objc.Send[MEExtensionManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEExtensionManager) Autorelease() MEExtensionManager {
	rv := objc.Send[MEExtensionManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEExtensionManager creates a new MEExtensionManager instance.
func NewMEExtensionManager() MEExtensionManager {
	return getMEExtensionManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEExtensionManager */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEExtensionManager
type MEExtensionManager struct {
	objectivec.Object
}

// MEExtensionManagerFrom constructs a [MEExtensionManager] from an unsafe.Pointer.
func MEExtensionManagerFrom(ptr unsafe.Pointer) MEExtensionManager {
	return MEExtensionManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEExtensionManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEExtensionManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEExtensionManager/reloadContentBlocker(withIdentifier:completionHandler:)
func (mc _MEExtensionManagerClass) ReloadContentBlockerWithIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("reloadContentBlockerWithIdentifier:completionHandler:"), identifier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReloadContentBlockerWithIdentifierCompletionHandler) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MailKit/MEExtensionManager/reloadVisibleMessages(completionHandler:)
func (mc _MEExtensionManagerClass) ReloadVisibleMessagesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("reloadVisibleMessagesWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReloadVisibleMessagesWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEExtensionManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEExtensionManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEExtensionManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEExtensionManager */



