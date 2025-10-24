// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class hideRegistration */


/* debug [class_header]: Header for hideRegistration */
// The class instance for the [hideRegistration] class.
var (
	HideRegistrationClass     _hideRegistrationClass
	HideRegistrationClassOnce sync.Once
)

func gethideRegistrationClass() _hideRegistrationClass {
	HideRegistrationClassOnce.Do(func() {
		HideRegistrationClass = _hideRegistrationClass{objc.GetClass("hideRegistration")}
	})
	return HideRegistrationClass
}

type _hideRegistrationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for hideRegistration */
// An interface definition for the [hideRegistration] class.
type IhideRegistration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for hideRegistration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for hideRegistration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for hideRegistration */
// Alloc allocates a new instance without initialization.
func (hc _hideRegistrationClass) Alloc() hideRegistration {
	rv := objc.Send[hideRegistration](objc.ID(hc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (hc _hideRegistrationClass) New() hideRegistration {
	rv := objc.Send[hideRegistration](objc.ID(hc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (h_ hideRegistration) Init() hideRegistration {
	rv := objc.Send[hideRegistration](h_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h_ hideRegistration) Autorelease() hideRegistration {
	rv := objc.Send[hideRegistration](h_.ID, objc.Sel("autorelease"))
	return rv
}

// NewhideRegistration creates a new hideRegistration instance.
func NewhideRegistration() hideRegistration {
	return gethideRegistrationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for hideRegistration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/hideRegistration-c.ivar
type hideRegistration struct {
	objectivec.Object
}

// hideRegistrationFrom constructs a [hideRegistration] from an unsafe.Pointer.
func hideRegistrationFrom(ptr unsafe.Pointer) hideRegistration {
	return hideRegistration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for hideRegistration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for hideRegistration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for hideRegistration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for hideRegistration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for hideRegistration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class hideRegistration */



