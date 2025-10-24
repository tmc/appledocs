// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class trustUsesSystemKeychain */


/* debug [class_header]: Header for trustUsesSystemKeychain */
// The class instance for the [trustUsesSystemKeychain] class.
var (
	TrustUsesSystemKeychainClass     _trustUsesSystemKeychainClass
	TrustUsesSystemKeychainClassOnce sync.Once
)

func gettrustUsesSystemKeychainClass() _trustUsesSystemKeychainClass {
	TrustUsesSystemKeychainClassOnce.Do(func() {
		TrustUsesSystemKeychainClass = _trustUsesSystemKeychainClass{objc.GetClass("trustUsesSystemKeychain")}
	})
	return TrustUsesSystemKeychainClass
}

type _trustUsesSystemKeychainClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for trustUsesSystemKeychain */
// An interface definition for the [trustUsesSystemKeychain] class.
type ItrustUsesSystemKeychain interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for trustUsesSystemKeychain */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for trustUsesSystemKeychain */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for trustUsesSystemKeychain */
// Alloc allocates a new instance without initialization.
func (tc _trustUsesSystemKeychainClass) Alloc() trustUsesSystemKeychain {
	rv := objc.Send[trustUsesSystemKeychain](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _trustUsesSystemKeychainClass) New() trustUsesSystemKeychain {
	rv := objc.Send[trustUsesSystemKeychain](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustUsesSystemKeychain) Init() trustUsesSystemKeychain {
	rv := objc.Send[trustUsesSystemKeychain](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustUsesSystemKeychain) Autorelease() trustUsesSystemKeychain {
	rv := objc.Send[trustUsesSystemKeychain](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustUsesSystemKeychain creates a new trustUsesSystemKeychain instance.
func NewtrustUsesSystemKeychain() trustUsesSystemKeychain {
	return gettrustUsesSystemKeychainClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for trustUsesSystemKeychain */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustUsesSystemKeychain-c.ivar
type trustUsesSystemKeychain struct {
	objectivec.Object
}

// trustUsesSystemKeychainFrom constructs a [trustUsesSystemKeychain] from an unsafe.Pointer.
func trustUsesSystemKeychainFrom(ptr unsafe.Pointer) trustUsesSystemKeychain {
	return trustUsesSystemKeychain{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for trustUsesSystemKeychain *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for trustUsesSystemKeychain */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for trustUsesSystemKeychain */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for trustUsesSystemKeychain */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for trustUsesSystemKeychain */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class trustUsesSystemKeychain */



