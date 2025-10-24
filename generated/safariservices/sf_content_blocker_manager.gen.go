// Code generated from Apple documentation for SafariServices. DO NOT EDIT.

package safariservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SFContentBlockerManager */


/* debug [class_header]: Header for SFContentBlockerManager */
// The class instance for the [SFContentBlockerManager] class.
var (
	SFContentBlockerManagerClass     _SFContentBlockerManagerClass
	SFContentBlockerManagerClassOnce sync.Once
)

func getSFContentBlockerManagerClass() _SFContentBlockerManagerClass {
	SFContentBlockerManagerClassOnce.Do(func() {
		SFContentBlockerManagerClass = _SFContentBlockerManagerClass{objc.GetClass("SFContentBlockerManager")}
	})
	return SFContentBlockerManagerClass
}

type _SFContentBlockerManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SFContentBlockerManager */
// An interface definition for the [SFContentBlockerManager] class.
type ISFContentBlockerManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SFContentBlockerManager */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SFContentBlockerManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SFContentBlockerManager */
// Alloc allocates a new instance without initialization.
func (sc _SFContentBlockerManagerClass) Alloc() SFContentBlockerManager {
	rv := objc.Send[SFContentBlockerManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SFContentBlockerManagerClass) New() SFContentBlockerManager {
	rv := objc.Send[SFContentBlockerManager](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SFContentBlockerManager) Init() SFContentBlockerManager {
	rv := objc.Send[SFContentBlockerManager](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SFContentBlockerManager) Autorelease() SFContentBlockerManager {
	rv := objc.Send[SFContentBlockerManager](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFContentBlockerManager creates a new SFContentBlockerManager instance.
func NewSFContentBlockerManager() SFContentBlockerManager {
	return getSFContentBlockerManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SFContentBlockerManager */
// A class that your app uses to interact with a content blocker extension.
//
// Use this class to determine the state of your content blocker and reload the content-blocking rules used by Safari.


// A class that your app uses to interact with a content blocker extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerManager
type SFContentBlockerManager struct {
	objectivec.Object
}

// SFContentBlockerManagerFrom constructs a [SFContentBlockerManager] from an unsafe.Pointer.
//
// A class that your app uses to interact with a content blocker extension.
func SFContentBlockerManagerFrom(ptr unsafe.Pointer) SFContentBlockerManager {
	return SFContentBlockerManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SFContentBlockerManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SFContentBlockerManager */

// Determines the state of your content blocker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerManager/getStateOfContentBlocker(withIdentifier:completionHandler:)
func (sc _SFContentBlockerManagerClass) GetStateOfContentBlockerWithIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("getStateOfContentBlockerWithIdentifier:completionHandler:"), identifier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=GetStateOfContentBlockerWithIdentifierCompletionHandler) */


// Tells Safari to reload the specified extension’s content-blocking rules.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/SafariServices/SFContentBlockerManager/reloadContentBlocker(withIdentifier:completionHandler:)
func (sc _SFContentBlockerManagerClass) ReloadContentBlockerWithIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(sc.class), objc.Sel("reloadContentBlockerWithIdentifier:completionHandler:"), identifier, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ReloadContentBlockerWithIdentifierCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SFContentBlockerManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SFContentBlockerManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SFContentBlockerManager */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SFContentBlockerManager */



