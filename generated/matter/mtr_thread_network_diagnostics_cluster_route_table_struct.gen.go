// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRThreadNetworkDiagnosticsClusterRouteTableStruct */


/* debug [class_header]: Header for MTRThreadNetworkDiagnosticsClusterRouteTableStruct */
// The class instance for the [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] class.
var (
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClass     _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClassOnce sync.Once
)

func getMTRThreadNetworkDiagnosticsClusterRouteTableStructClass() _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass {
	MTRThreadNetworkDiagnosticsClusterRouteTableStructClassOnce.Do(func() {
		MTRThreadNetworkDiagnosticsClusterRouteTableStructClass = _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass{objc.GetClass("MTRThreadNetworkDiagnosticsClusterRouteTableStruct")}
	})
	return MTRThreadNetworkDiagnosticsClusterRouteTableStructClass
}

type _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRThreadNetworkDiagnosticsClusterRouteTableStruct */
// An interface definition for the [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] class.
type IMTRThreadNetworkDiagnosticsClusterRouteTableStruct interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRThreadNetworkDiagnosticsClusterRouteTableStruct */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRThreadNetworkDiagnosticsClusterRouteTableStruct */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRThreadNetworkDiagnosticsClusterRouteTableStruct */
// Alloc allocates a new instance without initialization.
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass) Alloc() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRThreadNetworkDiagnosticsClusterRouteTableStructClass) New() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Init() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRThreadNetworkDiagnosticsClusterRouteTableStruct) Autorelease() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	rv := objc.Send[MTRThreadNetworkDiagnosticsClusterRouteTableStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRThreadNetworkDiagnosticsClusterRouteTableStruct creates a new MTRThreadNetworkDiagnosticsClusterRouteTableStruct instance.
func NewMTRThreadNetworkDiagnosticsClusterRouteTableStruct() MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	return getMTRThreadNetworkDiagnosticsClusterRouteTableStructClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRThreadNetworkDiagnosticsClusterRouteTableStruct */
// A parent class referenced by other Matter classes.


// A parent class referenced by other Matter classes. [Full Topic]
type MTRThreadNetworkDiagnosticsClusterRouteTableStruct struct {
	objectivec.Object
}

// MTRThreadNetworkDiagnosticsClusterRouteTableStructFrom constructs a [MTRThreadNetworkDiagnosticsClusterRouteTableStruct] from an unsafe.Pointer.
//
// A parent class referenced by other Matter classes.
func MTRThreadNetworkDiagnosticsClusterRouteTableStructFrom(ptr unsafe.Pointer) MTRThreadNetworkDiagnosticsClusterRouteTableStruct {
	return MTRThreadNetworkDiagnosticsClusterRouteTableStruct{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRThreadNetworkDiagnosticsClusterRouteTableStruct *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRThreadNetworkDiagnosticsClusterRouteTableStruct */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRThreadNetworkDiagnosticsClusterRouteTableStruct */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRThreadNetworkDiagnosticsClusterRouteTableStruct */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRThreadNetworkDiagnosticsClusterRouteTableStruct */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRThreadNetworkDiagnosticsClusterRouteTableStruct */



