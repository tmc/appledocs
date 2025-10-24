// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterICDManagement */


/* debug [class_header]: Header for MTRBaseClusterICDManagement */
// The class instance for the [MTRBaseClusterICDManagement] class.
var (
	MTRBaseClusterICDManagementClass     _MTRBaseClusterICDManagementClass
	MTRBaseClusterICDManagementClassOnce sync.Once
)

func getMTRBaseClusterICDManagementClass() _MTRBaseClusterICDManagementClass {
	MTRBaseClusterICDManagementClassOnce.Do(func() {
		MTRBaseClusterICDManagementClass = _MTRBaseClusterICDManagementClass{objc.GetClass("MTRBaseClusterICDManagement")}
	})
	return MTRBaseClusterICDManagementClass
}

type _MTRBaseClusterICDManagementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterICDManagement */
// An interface definition for the [MTRBaseClusterICDManagement] class.
type IMTRBaseClusterICDManagement interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterICDManagement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterICDManagement */
	// methods:
	StayActiveRequestWithParamsCompletion(params IMTRICDManagementClusterStayActiveRequestParams, completion unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterICDManagement */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterICDManagementClass) Alloc() MTRBaseClusterICDManagement {
	rv := objc.Send[MTRBaseClusterICDManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterICDManagementClass) New() MTRBaseClusterICDManagement {
	rv := objc.Send[MTRBaseClusterICDManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterICDManagement) Init() MTRBaseClusterICDManagement {
	rv := objc.Send[MTRBaseClusterICDManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterICDManagement) Autorelease() MTRBaseClusterICDManagement {
	rv := objc.Send[MTRBaseClusterICDManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterICDManagement creates a new MTRBaseClusterICDManagement instance.
func NewMTRBaseClusterICDManagement() MTRBaseClusterICDManagement {
	return getMTRBaseClusterICDManagementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterICDManagement */
// Cluster ICD Management
//
// Allows servers to ensure that listed clients are notified when a server is available for communication.


// Cluster ICD Management
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement
type MTRBaseClusterICDManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterICDManagementFrom constructs a [MTRBaseClusterICDManagement] from an unsafe.Pointer.
//
// Cluster ICD Management
func MTRBaseClusterICDManagementFrom(ptr unsafe.Pointer) MTRBaseClusterICDManagement {
	return MTRBaseClusterICDManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterICDManagement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterICDManagement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterICDManagement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterICDManagement */

// Command StayActiveRequest
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterICDManagement/stayActiveRequest(with:completion:)
func (m_ MTRBaseClusterICDManagement) StayActiveRequestWithParamsCompletion(params IMTRICDManagementClusterStayActiveRequestParams, completion unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("stayActiveRequestWithParams:completion:"), params, completion)
}/* debug [instance_methods/method]: StayActiveRequestWithParamsCompletion */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterICDManagement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterICDManagement */



