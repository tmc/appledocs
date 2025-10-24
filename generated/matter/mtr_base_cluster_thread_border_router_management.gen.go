// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterThreadBorderRouterManagement */


/* debug [class_header]: Header for MTRBaseClusterThreadBorderRouterManagement */
// The class instance for the [MTRBaseClusterThreadBorderRouterManagement] class.
var (
	MTRBaseClusterThreadBorderRouterManagementClass     _MTRBaseClusterThreadBorderRouterManagementClass
	MTRBaseClusterThreadBorderRouterManagementClassOnce sync.Once
)

func getMTRBaseClusterThreadBorderRouterManagementClass() _MTRBaseClusterThreadBorderRouterManagementClass {
	MTRBaseClusterThreadBorderRouterManagementClassOnce.Do(func() {
		MTRBaseClusterThreadBorderRouterManagementClass = _MTRBaseClusterThreadBorderRouterManagementClass{objc.GetClass("MTRBaseClusterThreadBorderRouterManagement")}
	})
	return MTRBaseClusterThreadBorderRouterManagementClass
}

type _MTRBaseClusterThreadBorderRouterManagementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterThreadBorderRouterManagement */
// An interface definition for the [MTRBaseClusterThreadBorderRouterManagement] class.
type IMTRBaseClusterThreadBorderRouterManagement interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterThreadBorderRouterManagement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterThreadBorderRouterManagement */
	// methods:
	SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterThreadBorderRouterManagement */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) Alloc() MTRBaseClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRBaseClusterThreadBorderRouterManagement](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterThreadBorderRouterManagementClass) New() MTRBaseClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRBaseClusterThreadBorderRouterManagement](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterThreadBorderRouterManagement) Init() MTRBaseClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRBaseClusterThreadBorderRouterManagement](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterThreadBorderRouterManagement) Autorelease() MTRBaseClusterThreadBorderRouterManagement {
	rv := objc.Send[MTRBaseClusterThreadBorderRouterManagement](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterThreadBorderRouterManagement creates a new MTRBaseClusterThreadBorderRouterManagement instance.
func NewMTRBaseClusterThreadBorderRouterManagement() MTRBaseClusterThreadBorderRouterManagement {
	return getMTRBaseClusterThreadBorderRouterManagementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterThreadBorderRouterManagement */
// Cluster Thread Border Router Management
//
// Manage the Thread network of Thread Border Router


// Cluster Thread Border Router Management
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement
type MTRBaseClusterThreadBorderRouterManagement struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterThreadBorderRouterManagementFrom constructs a [MTRBaseClusterThreadBorderRouterManagement] from an unsafe.Pointer.
//
// Cluster Thread Border Router Management
func MTRBaseClusterThreadBorderRouterManagementFrom(ptr unsafe.Pointer) MTRBaseClusterThreadBorderRouterManagement {
	return MTRBaseClusterThreadBorderRouterManagement{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterThreadBorderRouterManagement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterThreadBorderRouterManagement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterThreadBorderRouterManagement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterThreadBorderRouterManagement */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterThreadBorderRouterManagement/subscribeAttributeClusterRevision(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterThreadBorderRouterManagement) SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeClusterRevisionWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeClusterRevisionWithParamsSubscriptionEstablishedReportHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterThreadBorderRouterManagement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterThreadBorderRouterManagement */



