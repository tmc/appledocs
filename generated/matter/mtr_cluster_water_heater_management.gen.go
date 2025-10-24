// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterWaterHeaterManagement */


/* debug [class_header]: Header for MTRClusterWaterHeaterManagement */
// The class instance for the [MTRClusterWaterHeaterManagement] class.
var (
	MTRClusterWaterHeaterManagementClass     _MTRClusterWaterHeaterManagementClass
	MTRClusterWaterHeaterManagementClassOnce sync.Once
)

func getMTRClusterWaterHeaterManagementClass() _MTRClusterWaterHeaterManagementClass {
	MTRClusterWaterHeaterManagementClassOnce.Do(func() {
		MTRClusterWaterHeaterManagementClass = _MTRClusterWaterHeaterManagementClass{objc.GetClass("MTRClusterWaterHeaterManagement")}
	})
	return MTRClusterWaterHeaterManagementClass
}

type _MTRClusterWaterHeaterManagementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterWaterHeaterManagement */
// An interface definition for the [MTRClusterWaterHeaterManagement] class.
type IMTRClusterWaterHeaterManagement interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterWaterHeaterManagement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterWaterHeaterManagement */
	// methods:
	ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterWaterHeaterManagement */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterWaterHeaterManagementClass) Alloc() MTRClusterWaterHeaterManagement {
	rv := objc.Send[MTRClusterWaterHeaterManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterWaterHeaterManagementClass) New() MTRClusterWaterHeaterManagement {
	rv := objc.Send[MTRClusterWaterHeaterManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterWaterHeaterManagement) Init() MTRClusterWaterHeaterManagement {
	rv := objc.Send[MTRClusterWaterHeaterManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterWaterHeaterManagement) Autorelease() MTRClusterWaterHeaterManagement {
	rv := objc.Send[MTRClusterWaterHeaterManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterWaterHeaterManagement creates a new MTRClusterWaterHeaterManagement instance.
func NewMTRClusterWaterHeaterManagement() MTRClusterWaterHeaterManagement {
	return getMTRClusterWaterHeaterManagementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterWaterHeaterManagement */
// Cluster Water Heater Management This cluster is used to allow clients to control the operation of a hot water heating appliance so that it can be used with energy management.


// Cluster Water Heater Management This cluster is used to allow clients to control the operation of a hot water heating appliance so that it can be used with energy management.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement
type MTRClusterWaterHeaterManagement struct {
	MTRGenericCluster
}

// MTRClusterWaterHeaterManagementFrom constructs a [MTRClusterWaterHeaterManagement] from an unsafe.Pointer.
//
// Cluster Water Heater Management This cluster is used to allow clients to control the operation of a hot water heating appliance so that it can be used with energy management.
func MTRClusterWaterHeaterManagementFrom(ptr unsafe.Pointer) MTRClusterWaterHeaterManagement {
	return MTRClusterWaterHeaterManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterWaterHeaterManagement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterWaterHeaterManagement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterWaterHeaterManagement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterWaterHeaterManagement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterWaterHeaterManagement/readAttributeGeneratedCommandList(with:)
func (m_ MTRClusterWaterHeaterManagement) ReadAttributeGeneratedCommandListWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeGeneratedCommandListWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeGeneratedCommandListWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterWaterHeaterManagement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterWaterHeaterManagement */



