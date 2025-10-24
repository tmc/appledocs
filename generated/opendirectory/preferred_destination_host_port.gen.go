// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class preferredDestinationHostPort */


/* debug [class_header]: Header for preferredDestinationHostPort */
// The class instance for the [preferredDestinationHostPort] class.
var (
	PreferredDestinationHostPortClass     _preferredDestinationHostPortClass
	PreferredDestinationHostPortClassOnce sync.Once
)

func getpreferredDestinationHostPortClass() _preferredDestinationHostPortClass {
	PreferredDestinationHostPortClassOnce.Do(func() {
		PreferredDestinationHostPortClass = _preferredDestinationHostPortClass{objc.GetClass("preferredDestinationHostPort")}
	})
	return PreferredDestinationHostPortClass
}

type _preferredDestinationHostPortClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for preferredDestinationHostPort */
// An interface definition for the [preferredDestinationHostPort] class.
type IpreferredDestinationHostPort interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for preferredDestinationHostPort */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for preferredDestinationHostPort */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for preferredDestinationHostPort */
// Alloc allocates a new instance without initialization.
func (pc _preferredDestinationHostPortClass) Alloc() preferredDestinationHostPort {
	rv := objc.Send[preferredDestinationHostPort](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _preferredDestinationHostPortClass) New() preferredDestinationHostPort {
	rv := objc.Send[preferredDestinationHostPort](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ preferredDestinationHostPort) Init() preferredDestinationHostPort {
	rv := objc.Send[preferredDestinationHostPort](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ preferredDestinationHostPort) Autorelease() preferredDestinationHostPort {
	rv := objc.Send[preferredDestinationHostPort](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewpreferredDestinationHostPort creates a new preferredDestinationHostPort instance.
func NewpreferredDestinationHostPort() preferredDestinationHostPort {
	return getpreferredDestinationHostPortClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for preferredDestinationHostPort */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/preferredDestinationHostPort-c.ivar
type preferredDestinationHostPort struct {
	objectivec.Object
}

// preferredDestinationHostPortFrom constructs a [preferredDestinationHostPort] from an unsafe.Pointer.
func preferredDestinationHostPortFrom(ptr unsafe.Pointer) preferredDestinationHostPort {
	return preferredDestinationHostPort{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for preferredDestinationHostPort *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for preferredDestinationHostPort */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for preferredDestinationHostPort */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for preferredDestinationHostPort */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for preferredDestinationHostPort */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class preferredDestinationHostPort */



