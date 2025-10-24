// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class trustAccount */


/* debug [class_header]: Header for trustAccount */
// The class instance for the [trustAccount] class.
var (
	TrustAccountClass     _trustAccountClass
	TrustAccountClassOnce sync.Once
)

func gettrustAccountClass() _trustAccountClass {
	TrustAccountClassOnce.Do(func() {
		TrustAccountClass = _trustAccountClass{objc.GetClass("trustAccount")}
	})
	return TrustAccountClass
}

type _trustAccountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for trustAccount */
// An interface definition for the [trustAccount] class.
type ItrustAccount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for trustAccount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for trustAccount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for trustAccount */
// Alloc allocates a new instance without initialization.
func (tc _trustAccountClass) Alloc() trustAccount {
	rv := objc.Send[trustAccount](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _trustAccountClass) New() trustAccount {
	rv := objc.Send[trustAccount](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustAccount) Init() trustAccount {
	rv := objc.Send[trustAccount](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustAccount) Autorelease() trustAccount {
	rv := objc.Send[trustAccount](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustAccount creates a new trustAccount instance.
func NewtrustAccount() trustAccount {
	return gettrustAccountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for trustAccount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustAccount-c.ivar
type trustAccount struct {
	objectivec.Object
}

// trustAccountFrom constructs a [trustAccount] from an unsafe.Pointer.
func trustAccountFrom(ptr unsafe.Pointer) trustAccount {
	return trustAccount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for trustAccount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for trustAccount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for trustAccount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for trustAccount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for trustAccount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class trustAccount */



