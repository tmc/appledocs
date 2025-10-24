// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXCallDirectoryManager */


/* debug [class_header]: Header for CXCallDirectoryManager */
// The class instance for the [CXCallDirectoryManager] class.
var (
	CXCallDirectoryManagerClass     _CXCallDirectoryManagerClass
	CXCallDirectoryManagerClassOnce sync.Once
)

func getCXCallDirectoryManagerClass() _CXCallDirectoryManagerClass {
	CXCallDirectoryManagerClassOnce.Do(func() {
		CXCallDirectoryManagerClass = _CXCallDirectoryManagerClass{objc.GetClass("CXCallDirectoryManager")}
	})
	return CXCallDirectoryManagerClass
}

type _CXCallDirectoryManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXCallDirectoryManager */
// An interface definition for the [CXCallDirectoryManager] class.
type ICXCallDirectoryManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXCallDirectoryManager */
	// properties:
	CXErrorDomainCallDirectoryManager() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXCallDirectoryManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXCallDirectoryManager */
// Alloc allocates a new instance without initialization.
func (cc _CXCallDirectoryManagerClass) Alloc() CXCallDirectoryManager {
	rv := objc.Send[CXCallDirectoryManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXCallDirectoryManagerClass) New() CXCallDirectoryManager {
	rv := objc.Send[CXCallDirectoryManager](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallDirectoryManager) Init() CXCallDirectoryManager {
	rv := objc.Send[CXCallDirectoryManager](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallDirectoryManager) Autorelease() CXCallDirectoryManager {
	rv := objc.Send[CXCallDirectoryManager](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallDirectoryManager creates a new CXCallDirectoryManager instance.
func NewCXCallDirectoryManager() CXCallDirectoryManager {
	return getCXCallDirectoryManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXCallDirectoryManager */
// The programmatic interface to an object that manages a Call Directory app extension.


// The programmatic interface to an object that manages a Call Directory app extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager
type CXCallDirectoryManager struct {
	objectivec.Object
}

// CXCallDirectoryManagerFrom constructs a [CXCallDirectoryManager] from an unsafe.Pointer.
//
// The programmatic interface to an object that manages a Call Directory app extension.
func CXCallDirectoryManagerFrom(ptr unsafe.Pointer) CXCallDirectoryManager {
	return CXCallDirectoryManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXCallDirectoryManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXCallDirectoryManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXCallDirectoryManager */

// Returns the shared call directory manager instance for the app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/sharedInstance
func (cc _CXCallDirectoryManagerClass) SharedInstance() CXCallDirectoryManager {
	rv := objc.Send[CXCallDirectoryManager](objc.ID(cc.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_properties_class/property]: sharedInstance */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXCallDirectoryManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXCallDirectoryManager */

// Domain for errors when interacting with a call directory manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxerrordomaincalldirectorymanager
func (c_ CXCallDirectoryManager) CXErrorDomainCallDirectoryManager() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("CXErrorDomainCallDirectoryManager"))
	return rv
}/* debug [instance_properties/getter]: CXErrorDomainCallDirectoryManager */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXCallDirectoryManager */


