// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRGenericBaseCluster */


/* debug [class_header]: Header for MTRGenericBaseCluster */
// The class instance for the [MTRGenericBaseCluster] class.
var (
	MTRGenericBaseClusterClass     _MTRGenericBaseClusterClass
	MTRGenericBaseClusterClassOnce sync.Once
)

func getMTRGenericBaseClusterClass() _MTRGenericBaseClusterClass {
	MTRGenericBaseClusterClassOnce.Do(func() {
		MTRGenericBaseClusterClass = _MTRGenericBaseClusterClass{objc.GetClass("MTRGenericBaseCluster")}
	})
	return MTRGenericBaseClusterClass
}

type _MTRGenericBaseClusterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRGenericBaseCluster */
// An interface definition for the [MTRGenericBaseCluster] class.
type IMTRGenericBaseCluster interface {
	IMTRCluster
	
/* debug [class_interface_properties]: Properties for MTRGenericBaseCluster */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRGenericBaseCluster */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRGenericBaseCluster */
// Alloc allocates a new instance without initialization.
func (mc _MTRGenericBaseClusterClass) Alloc() MTRGenericBaseCluster {
	rv := objc.Send[MTRGenericBaseCluster](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRGenericBaseClusterClass) New() MTRGenericBaseCluster {
	rv := objc.Send[MTRGenericBaseCluster](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRGenericBaseCluster) Init() MTRGenericBaseCluster {
	rv := objc.Send[MTRGenericBaseCluster](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRGenericBaseCluster) Autorelease() MTRGenericBaseCluster {
	rv := objc.Send[MTRGenericBaseCluster](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRGenericBaseCluster creates a new MTRGenericBaseCluster instance.
func NewMTRGenericBaseCluster() MTRGenericBaseCluster {
	return getMTRGenericBaseClusterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRGenericBaseCluster */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRGenericBaseCluster
type MTRGenericBaseCluster struct {
	MTRCluster
}

// MTRGenericBaseClusterFrom constructs a [MTRGenericBaseCluster] from an unsafe.Pointer.
func MTRGenericBaseClusterFrom(ptr unsafe.Pointer) MTRGenericBaseCluster {
	return MTRGenericBaseCluster{
		MTRCluster: MTRClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRGenericBaseCluster *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRGenericBaseCluster */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRGenericBaseCluster */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRGenericBaseCluster */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRGenericBaseCluster */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRGenericBaseCluster */



