// Code generated from Apple documentation for StoreKitTest. DO NOT EDIT.

package storekittest

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKAdTestSession */

/* debug [class_header]: Header for SKAdTestSession */
// The class instance for the [AdTestSession] class.
var (
	AdTestSessionClass     _AdTestSessionClass
	AdTestSessionClassOnce sync.Once
)

func getAdTestSessionClass() _AdTestSessionClass {
	AdTestSessionClassOnce.Do(func() {
		AdTestSessionClass = _AdTestSessionClass{objc.GetClass("SKAdTestSession")}
	})
	return AdTestSessionClass
}

type _AdTestSessionClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for AdTestSession */
// An interface definition for the [AdTestSession] class.
type IAdTestSession interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for AdTestSession */
	// properties:
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for AdTestSession */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for AdTestSession */
// Alloc allocates a new instance without initialization.
func (ac _AdTestSessionClass) Alloc() AdTestSession {
	rv := objc.Send[AdTestSession](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AdTestSessionClass) New() AdTestSession {
	rv := objc.Send[AdTestSession](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AdTestSession) Init() AdTestSession {
	rv := objc.Send[AdTestSession](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AdTestSession) Autorelease() AdTestSession {
	rv := objc.Send[AdTestSession](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAdTestSession creates a new AdTestSession instance.
func NewAdTestSession() AdTestSession {
	return getAdTestSessionClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for AdTestSession */
// The class you use to test ad impressions and postbacks in Xcode.
//
// Use the class to test your implementations of SKAdNetwork. Create one instance of this class to use in multiple test cases. The instance represents a test session, and holds a set of test postbacks. Use to create test postbacks. Call to add test postbacks to the test session. The test session deletes the postbacks from the instance after you call .

// The class you use to test ad impressions and postbacks in Xcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKitTest/SKAdTestSession
type AdTestSession struct {
	objectivec.Object
}

// AdTestSessionFrom constructs a [AdTestSession] from an unsafe.Pointer.
//
// The class you use to test ad impressions and postbacks in Xcode.
func AdTestSessionFrom(ptr unsafe.Pointer) AdTestSession {
	return AdTestSession{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for AdTestSession */
/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for AdTestSession */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for AdTestSession */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for AdTestSession */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for AdTestSession */
/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKAdTestSession */
