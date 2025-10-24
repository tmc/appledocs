// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRApplicationLauncherClusterApplication */


/* debug [class_header]: Header for MTRApplicationLauncherClusterApplication */
// The class instance for the [MTRApplicationLauncherClusterApplication] class.
var (
	MTRApplicationLauncherClusterApplicationClass     _MTRApplicationLauncherClusterApplicationClass
	MTRApplicationLauncherClusterApplicationClassOnce sync.Once
)

func getMTRApplicationLauncherClusterApplicationClass() _MTRApplicationLauncherClusterApplicationClass {
	MTRApplicationLauncherClusterApplicationClassOnce.Do(func() {
		MTRApplicationLauncherClusterApplicationClass = _MTRApplicationLauncherClusterApplicationClass{objc.GetClass("MTRApplicationLauncherClusterApplication")}
	})
	return MTRApplicationLauncherClusterApplicationClass
}

type _MTRApplicationLauncherClusterApplicationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRApplicationLauncherClusterApplication */
// An interface definition for the [MTRApplicationLauncherClusterApplication] class.
type IMTRApplicationLauncherClusterApplication interface {
	IMTRApplicationLauncherClusterApplicationStruct
	
/* debug [class_interface_properties]: Properties for MTRApplicationLauncherClusterApplication */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRApplicationLauncherClusterApplication */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRApplicationLauncherClusterApplication */
// Alloc allocates a new instance without initialization.
func (mc _MTRApplicationLauncherClusterApplicationClass) Alloc() MTRApplicationLauncherClusterApplication {
	rv := objc.Send[MTRApplicationLauncherClusterApplication](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRApplicationLauncherClusterApplicationClass) New() MTRApplicationLauncherClusterApplication {
	rv := objc.Send[MTRApplicationLauncherClusterApplication](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRApplicationLauncherClusterApplication) Init() MTRApplicationLauncherClusterApplication {
	rv := objc.Send[MTRApplicationLauncherClusterApplication](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRApplicationLauncherClusterApplication) Autorelease() MTRApplicationLauncherClusterApplication {
	rv := objc.Send[MTRApplicationLauncherClusterApplication](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRApplicationLauncherClusterApplication creates a new MTRApplicationLauncherClusterApplication instance.
func NewMTRApplicationLauncherClusterApplication() MTRApplicationLauncherClusterApplication {
	return getMTRApplicationLauncherClusterApplicationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRApplicationLauncherClusterApplication */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRApplicationLauncherClusterApplication
type MTRApplicationLauncherClusterApplication struct {
	MTRApplicationLauncherClusterApplicationStruct
}

// MTRApplicationLauncherClusterApplicationFrom constructs a [MTRApplicationLauncherClusterApplication] from an unsafe.Pointer.
func MTRApplicationLauncherClusterApplicationFrom(ptr unsafe.Pointer) MTRApplicationLauncherClusterApplication {
	return MTRApplicationLauncherClusterApplication{
		MTRApplicationLauncherClusterApplicationStruct: MTRApplicationLauncherClusterApplicationStructFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRApplicationLauncherClusterApplication *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRApplicationLauncherClusterApplication */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRApplicationLauncherClusterApplication */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRApplicationLauncherClusterApplication */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRApplicationLauncherClusterApplication */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRApplicationLauncherClusterApplication */



