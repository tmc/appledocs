// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class preferredDestinationHostName */


/* debug [class_header]: Header for preferredDestinationHostName */
// The class instance for the [preferredDestinationHostName] class.
var (
	PreferredDestinationHostNameClass     _preferredDestinationHostNameClass
	PreferredDestinationHostNameClassOnce sync.Once
)

func getpreferredDestinationHostNameClass() _preferredDestinationHostNameClass {
	PreferredDestinationHostNameClassOnce.Do(func() {
		PreferredDestinationHostNameClass = _preferredDestinationHostNameClass{objc.GetClass("preferredDestinationHostName")}
	})
	return PreferredDestinationHostNameClass
}

type _preferredDestinationHostNameClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for preferredDestinationHostName */
// An interface definition for the [preferredDestinationHostName] class.
type IpreferredDestinationHostName interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for preferredDestinationHostName */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for preferredDestinationHostName */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for preferredDestinationHostName */
// Alloc allocates a new instance without initialization.
func (pc _preferredDestinationHostNameClass) Alloc() preferredDestinationHostName {
	rv := objc.Send[preferredDestinationHostName](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _preferredDestinationHostNameClass) New() preferredDestinationHostName {
	rv := objc.Send[preferredDestinationHostName](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ preferredDestinationHostName) Init() preferredDestinationHostName {
	rv := objc.Send[preferredDestinationHostName](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ preferredDestinationHostName) Autorelease() preferredDestinationHostName {
	rv := objc.Send[preferredDestinationHostName](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewpreferredDestinationHostName creates a new preferredDestinationHostName instance.
func NewpreferredDestinationHostName() preferredDestinationHostName {
	return getpreferredDestinationHostNameClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for preferredDestinationHostName */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostName-c.ivar
type preferredDestinationHostName struct {
	objectivec.Object
}

// preferredDestinationHostNameFrom constructs a [preferredDestinationHostName] from an unsafe.Pointer.
func preferredDestinationHostNameFrom(ptr unsafe.Pointer) preferredDestinationHostName {
	return preferredDestinationHostName{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for preferredDestinationHostName *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for preferredDestinationHostName */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for preferredDestinationHostName */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for preferredDestinationHostName */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for preferredDestinationHostName */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class preferredDestinationHostName */



