// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRGenericCluster */


/* debug [class_header]: Header for MTRGenericCluster */
// The class instance for the [MTRGenericCluster] class.
var (
	MTRGenericClusterClass     _MTRGenericClusterClass
	MTRGenericClusterClassOnce sync.Once
)

func getMTRGenericClusterClass() _MTRGenericClusterClass {
	MTRGenericClusterClassOnce.Do(func() {
		MTRGenericClusterClass = _MTRGenericClusterClass{objc.GetClass("MTRGenericCluster")}
	})
	return MTRGenericClusterClass
}

type _MTRGenericClusterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRGenericCluster */
// An interface definition for the [MTRGenericCluster] class.
type IMTRGenericCluster interface {
	IMTRCluster
	
/* debug [class_interface_properties]: Properties for MTRGenericCluster */
	// properties:
	Device() IMTRDevice
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRGenericCluster */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRGenericCluster */
// Alloc allocates a new instance without initialization.
func (mc _MTRGenericClusterClass) Alloc() MTRGenericCluster {
	rv := objc.Send[MTRGenericCluster](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRGenericClusterClass) New() MTRGenericCluster {
	rv := objc.Send[MTRGenericCluster](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGenericCluster) Init() MTRGenericCluster {
	rv := objc.Send[MTRGenericCluster](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGenericCluster) Autorelease() MTRGenericCluster {
	rv := objc.Send[MTRGenericCluster](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGenericCluster creates a new MTRGenericCluster instance.
func NewMTRGenericCluster() MTRGenericCluster {
	return getMTRGenericClusterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRGenericCluster */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGenericCluster
type MTRGenericCluster struct {
	MTRCluster
}

// MTRGenericClusterFrom constructs a [MTRGenericCluster] from an unsafe.Pointer.
func MTRGenericClusterFrom(ptr unsafe.Pointer) MTRGenericCluster {
	return MTRGenericCluster{
		MTRCluster: MTRClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRGenericCluster *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRGenericCluster */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRGenericCluster */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRGenericCluster */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRGenericCluster */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGenericCluster/device
func (m_ MTRGenericCluster) Device() IMTRDevice {
	rv := objc.Send[MTRDevice](m_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRGenericCluster */



