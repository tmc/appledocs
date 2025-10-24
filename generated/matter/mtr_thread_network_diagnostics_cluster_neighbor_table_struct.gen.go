// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */


/* debug [class_header]: Header for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */
// The class instance for the [MTRThreadNetworkDiagnosticsClusterNeighborTableStruct] class.
var (
	MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass     _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass
	MTRThreadNetworkDiagnosticsClusterNeighborTableStructClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterNeighborTableStructClass() _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass {
	MTRThreadNetworkDiagnosticsClusterNeighborTableStructClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass = _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterNeighborTableStruct")}
	})
	return MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass
}

type _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */
// An interface definition for the [MTRThreadNetworkDiagnosticsClusterNeighborTableStruct] class.
type IMTRThreadNetworkDiagnosticsClusterNeighborTableStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass) Alloc() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadNetworkDiagnosticsClusterNeighborTableStructClass) New() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) Init() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterNeighborTableStruct) Autorelease() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterNeighborTableStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterNeighborTableStruct creates a new MTRThreadNetworkDiagnosticsClusterNeighborTableStruct instance.
func NewMTRThreadNetworkDiagnosticsClusterNeighborTableStruct() MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	return getMTRThreadNetworkDiagnosticsClusterNeighborTableStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */
// A parent class referenced by other Matter classes.


// A parent class referenced by other Matter classes. [Full Topic]
type MTRThreadNetworkDiagnosticsClusterNeighborTableStruct struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterNeighborTableStructFrom constructs a [MTRThreadNetworkDiagnosticsClusterNeighborTableStruct] from an unsafe.Pointer.
//
// A parent class referenced by other Matter classes.
func MTRThreadNetworkDiagnosticsClusterNeighborTableStructFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterNeighborTableStruct {
	return MTRThreadNetworkDiagnosticsClusterNeighborTableStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDiagnosticsClusterNeighborTableStruct */



