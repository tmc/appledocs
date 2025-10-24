// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class authenticationModuleEntries */


/* debug [class_header]: Header for authenticationModuleEntries */
// The class instance for the [authenticationModuleEntries] class.
var (
	AuthenticationModuleEntriesClass     _authenticationModuleEntriesClass
	AuthenticationModuleEntriesClassOnce sync.Once
)

func getauthenticationModuleEntriesClass() _authenticationModuleEntriesClass {
	AuthenticationModuleEntriesClassOnce.Do(func() {
		AuthenticationModuleEntriesClass = _authenticationModuleEntriesClass{objc.GetClass("authenticationModuleEntries")}
	})
	return AuthenticationModuleEntriesClass
}

type _authenticationModuleEntriesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for authenticationModuleEntries */
// An interface definition for the [authenticationModuleEntries] class.
type IauthenticationModuleEntries interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for authenticationModuleEntries */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for authenticationModuleEntries */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for authenticationModuleEntries */
// Alloc allocates a new instance without initialization.
func (ac _authenticationModuleEntriesClass) Alloc() authenticationModuleEntries {
	rv := objc.Send[authenticationModuleEntries](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _authenticationModuleEntriesClass) New() authenticationModuleEntries {
	rv := objc.Send[authenticationModuleEntries](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ authenticationModuleEntries) Init() authenticationModuleEntries {
	rv := objc.Send[authenticationModuleEntries](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ authenticationModuleEntries) Autorelease() authenticationModuleEntries {
	rv := objc.Send[authenticationModuleEntries](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewauthenticationModuleEntries creates a new authenticationModuleEntries instance.
func NewauthenticationModuleEntries() authenticationModuleEntries {
	return getauthenticationModuleEntriesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for authenticationModuleEntries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/authenticationModuleEntries-c.ivar
type authenticationModuleEntries struct {
	objectivec.Object
}

// authenticationModuleEntriesFrom constructs a [authenticationModuleEntries] from an unsafe.Pointer.
func authenticationModuleEntriesFrom(ptr unsafe.Pointer) authenticationModuleEntries {
	return authenticationModuleEntries{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for authenticationModuleEntries *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for authenticationModuleEntries */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for authenticationModuleEntries */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for authenticationModuleEntries */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for authenticationModuleEntries */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class authenticationModuleEntries */



