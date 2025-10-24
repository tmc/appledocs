// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class trustMetaAccount */


/* debug [class_header]: Header for trustMetaAccount */
// The class instance for the [trustMetaAccount] class.
var (
	TrustMetaAccountClass     _trustMetaAccountClass
	TrustMetaAccountClassOnce sync.Once
)

func gettrustMetaAccountClass() _trustMetaAccountClass {
	TrustMetaAccountClassOnce.Do(func() {
		TrustMetaAccountClass = _trustMetaAccountClass{objc.GetClass("trustMetaAccount")}
	})
	return TrustMetaAccountClass
}

type _trustMetaAccountClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for trustMetaAccount */
// An interface definition for the [trustMetaAccount] class.
type ItrustMetaAccount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for trustMetaAccount */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for trustMetaAccount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for trustMetaAccount */
// Alloc allocates a new instance without initialization.
func (tc _trustMetaAccountClass) Alloc() trustMetaAccount {
	rv := objc.Send[trustMetaAccount](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _trustMetaAccountClass) New() trustMetaAccount {
	rv := objc.Send[trustMetaAccount](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustMetaAccount) Init() trustMetaAccount {
	rv := objc.Send[trustMetaAccount](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustMetaAccount) Autorelease() trustMetaAccount {
	rv := objc.Send[trustMetaAccount](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustMetaAccount creates a new trustMetaAccount instance.
func NewtrustMetaAccount() trustMetaAccount {
	return gettrustMetaAccountClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for trustMetaAccount */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustMetaAccount-c.ivar
type trustMetaAccount struct {
	objectivec.Object
}

// trustMetaAccountFrom constructs a [trustMetaAccount] from an unsafe.Pointer.
func trustMetaAccountFrom(ptr unsafe.Pointer) trustMetaAccount {
	return trustMetaAccount{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for trustMetaAccount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for trustMetaAccount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for trustMetaAccount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for trustMetaAccount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for trustMetaAccount */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class trustMetaAccount */



