// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MTRClusterICDManagement */


/* debug [class_header]: Header for MTRClusterICDManagement */
// The class instance for the [MTRClusterICDManagement] class.
var (
	MTRClusterICDManagementClass     _MTRClusterICDManagementClass
	MTRClusterICDManagementClassOnce sync.Once
)

func getMTRClusterICDManagementClass() _MTRClusterICDManagementClass {
	MTRClusterICDManagementClassOnce.Do(func() {
		MTRClusterICDManagementClass = _MTRClusterICDManagementClass{objc.GetClass("MTRClusterICDManagement")}
	})
	return MTRClusterICDManagementClass
}

type _MTRClusterICDManagementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRClusterICDManagement */
// An interface definition for the [MTRClusterICDManagement] class.
type IMTRClusterICDManagement interface {
	IMTRGenericCluster
	
/* debug [class_interface_properties]: Properties for MTRClusterICDManagement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRClusterICDManagement */
	// methods:
	ReadAttributeICDCounterWithParams(params IMTRReadParams) foundation.IDictionary
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRClusterICDManagement */
// Alloc allocates a new instance without initialization.
func (mc _MTRClusterICDManagementClass) Alloc() MTRClusterICDManagement {
	rv := objc.Send[MTRClusterICDManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRClusterICDManagementClass) New() MTRClusterICDManagement {
	rv := objc.Send[MTRClusterICDManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterICDManagement) Init() MTRClusterICDManagement {
	rv := objc.Send[MTRClusterICDManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterICDManagement) Autorelease() MTRClusterICDManagement {
	rv := objc.Send[MTRClusterICDManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterICDManagement creates a new MTRClusterICDManagement instance.
func NewMTRClusterICDManagement() MTRClusterICDManagement {
	return getMTRClusterICDManagementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRClusterICDManagement */
// Cluster ICD Management Allows servers to ensure that listed clients are notified when a server is available for communication.


// Cluster ICD Management Allows servers to ensure that listed clients are notified when a server is available for communication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement
type MTRClusterICDManagement struct {
	MTRGenericCluster
}

// MTRClusterICDManagementFrom constructs a [MTRClusterICDManagement] from an unsafe.Pointer.
//
// Cluster ICD Management Allows servers to ensure that listed clients are notified when a server is available for communication.
func MTRClusterICDManagementFrom(ptr unsafe.Pointer) MTRClusterICDManagement {
	return MTRClusterICDManagement{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRClusterICDManagement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRClusterICDManagement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRClusterICDManagement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRClusterICDManagement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterICDManagement/readAttributeICDCounter(with:)
func (m_ MTRClusterICDManagement) ReadAttributeICDCounterWithParams(params IMTRReadParams) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("readAttributeICDCounterWithParams:"), params)
	return rv
}/* debug [instance_methods/method]: ReadAttributeICDCounterWithParams */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRClusterICDManagement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRClusterICDManagement */



