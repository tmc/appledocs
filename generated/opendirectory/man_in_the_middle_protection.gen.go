// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class manInTheMiddleProtection */


/* debug [class_header]: Header for manInTheMiddleProtection */
// The class instance for the [manInTheMiddleProtection] class.
var (
	ManInTheMiddleProtectionClass     _manInTheMiddleProtectionClass
	ManInTheMiddleProtectionClassOnce sync.Once
)

func getmanInTheMiddleProtectionClass() _manInTheMiddleProtectionClass {
	ManInTheMiddleProtectionClassOnce.Do(func() {
		ManInTheMiddleProtectionClass = _manInTheMiddleProtectionClass{objc.GetClass("manInTheMiddleProtection")}
	})
	return ManInTheMiddleProtectionClass
}

type _manInTheMiddleProtectionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for manInTheMiddleProtection */
// An interface definition for the [manInTheMiddleProtection] class.
type ImanInTheMiddleProtection interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for manInTheMiddleProtection */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for manInTheMiddleProtection */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for manInTheMiddleProtection */
// Alloc allocates a new instance without initialization.
func (mc _manInTheMiddleProtectionClass) Alloc() manInTheMiddleProtection {
	rv := objc.Send[manInTheMiddleProtection](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _manInTheMiddleProtectionClass) New() manInTheMiddleProtection {
	rv := objc.Send[manInTheMiddleProtection](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ manInTheMiddleProtection) Init() manInTheMiddleProtection {
	rv := objc.Send[manInTheMiddleProtection](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ manInTheMiddleProtection) Autorelease() manInTheMiddleProtection {
	rv := objc.Send[manInTheMiddleProtection](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmanInTheMiddleProtection creates a new manInTheMiddleProtection instance.
func NewmanInTheMiddleProtection() manInTheMiddleProtection {
	return getmanInTheMiddleProtectionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for manInTheMiddleProtection */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/manInTheMiddleProtection-c.ivar
type manInTheMiddleProtection struct {
	objectivec.Object
}

// manInTheMiddleProtectionFrom constructs a [manInTheMiddleProtection] from an unsafe.Pointer.
func manInTheMiddleProtectionFrom(ptr unsafe.Pointer) manInTheMiddleProtection {
	return manInTheMiddleProtection{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for manInTheMiddleProtection *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for manInTheMiddleProtection */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for manInTheMiddleProtection */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for manInTheMiddleProtection */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for manInTheMiddleProtection */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class manInTheMiddleProtection */



