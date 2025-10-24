// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ACAccountStore */


/* debug [class_header]: Header for ACAccountStore */
// The class instance for the [ACAccountStore] class.
var (
	ACAccountStoreClass     _ACAccountStoreClass
	ACAccountStoreClassOnce sync.Once
)

func getACAccountStoreClass() _ACAccountStoreClass {
	ACAccountStoreClassOnce.Do(func() {
		ACAccountStoreClass = _ACAccountStoreClass{objc.GetClass("ACAccountStore")}
	})
	return ACAccountStoreClass
}

type _ACAccountStoreClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ACAccountStore */
// An interface definition for the [ACAccountStore] class.
type IACAccountStore interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ACAccountStore */
	// properties:
	Accounts() objc.IObject /* cross-framework: NSArray */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ACAccountStore */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ACAccountStore */
// Alloc allocates a new instance without initialization.
func (ac _ACAccountStoreClass) Alloc() ACAccountStore {
	rv := objc.Send[ACAccountStore](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _ACAccountStoreClass) New() ACAccountStore {
	rv := objc.Send[ACAccountStore](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ACAccountStore) Init() ACAccountStore {
	rv := objc.Send[ACAccountStore](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ACAccountStore) Autorelease() ACAccountStore {
	rv := objc.Send[ACAccountStore](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewACAccountStore creates a new ACAccountStore instance.
func NewACAccountStore() ACAccountStore {
	return getACAccountStoreClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ACAccountStore */
// The object you use to request, manage, and store the user’s account information.
//
// The class provides an interface for accessing, managing, and storing accounts. To create and retrieve accounts from the Accounts database, you must create an object. Each object belongs to a single account store object.


// The object you use to request, manage, and store the user’s account information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore
type ACAccountStore struct {
	objectivec.Object
}

// ACAccountStoreFrom constructs a [ACAccountStore] from an unsafe.Pointer.
//
// The object you use to request, manage, and store the user’s account information.
func ACAccountStoreFrom(ptr unsafe.Pointer) ACAccountStore {
	return ACAccountStore{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ACAccountStore *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ACAccountStore */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ACAccountStore */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ACAccountStore */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ACAccountStore */

// The accounts managed by this account store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/accounts
func (a_ ACAccountStore) Accounts() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](a_.ID, objc.Sel("accounts"))
	return rv
}/* debug [instance_properties/getter]: accounts */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ACAccountStore */


