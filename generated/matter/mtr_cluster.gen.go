// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRCluster */


/* debug [class_header]: Header for MTRCluster */
// The class instance for the [MTRCluster] class.
var (
	MTRClusterClass     _MTRClusterClass
	MTRClusterClassOnce sync.Once
)

func getMTRClusterClass() _MTRClusterClass {
	MTRClusterClassOnce.Do(func() {
		MTRClusterClass = _MTRClusterClass{objc.GetClass("MTRCluster")}
	})
	return MTRClusterClass
}

type _MTRClusterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRCluster */
// An interface definition for the [MTRCluster] class.
type IMTRCluster interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRCluster */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRCluster */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRCluster */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterClass) Alloc() MTRCluster {
	rv := objc.Send[MTRCluster](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterClass) New() MTRCluster {
	rv := objc.Send[MTRCluster](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRCluster) Init() MTRCluster {
	rv := objc.Send[MTRCluster](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRCluster) Autorelease() MTRCluster {
	rv := objc.Send[MTRCluster](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRCluster creates a new MTRCluster instance.
func NewMTRCluster() MTRCluster {
	return getMTRClusterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRCluster */
// A parent class referenced by other Matter classes.


// A parent class referenced by other Matter classes. [Full Topic]
type MTRCluster struct {
	objectivec.Object
}

// MTRClusterFrom constructs a [MTRCluster] from an unsafe.Pointer.
//
// A parent class referenced by other Matter classes.
func MTRClusterFrom(ptr unsafe.Pointer) MTRCluster {
	return MTRCluster{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRCluster *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRCluster */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRCluster */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRCluster */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRCluster */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRCluster */



