// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterServiceArea */


/* debug [class_header]: Header for MTRBaseClusterServiceArea */
// The class instance for the [MTRBaseClusterServiceArea] class.
var (
	MTRBaseClusterServiceAreaClass     _MTRBaseClusterServiceAreaClass
	MTRBaseClusterServiceAreaClassOnce sync.Once
)

func getMTRBaseClusterServiceAreaClass() _MTRBaseClusterServiceAreaClass {
	MTRBaseClusterServiceAreaClassOnce.Do(func() {
		MTRBaseClusterServiceAreaClass = _MTRBaseClusterServiceAreaClass{objc.GetClass("MTRBaseClusterServiceArea")}
	})
	return MTRBaseClusterServiceAreaClass
}

type _MTRBaseClusterServiceAreaClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterServiceArea */
// An interface definition for the [MTRBaseClusterServiceArea] class.
type IMTRBaseClusterServiceArea interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterServiceArea */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterServiceArea */
	// methods:
	SubscribeAttributeEstimatedEndTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterServiceArea */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterServiceAreaClass) Alloc() MTRBaseClusterServiceArea {
	rv := objc.Send[MTRBaseClusterServiceArea](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterServiceAreaClass) New() MTRBaseClusterServiceArea {
	rv := objc.Send[MTRBaseClusterServiceArea](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterServiceArea) Init() MTRBaseClusterServiceArea {
	rv := objc.Send[MTRBaseClusterServiceArea](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterServiceArea) Autorelease() MTRBaseClusterServiceArea {
	rv := objc.Send[MTRBaseClusterServiceArea](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterServiceArea creates a new MTRBaseClusterServiceArea instance.
func NewMTRBaseClusterServiceArea() MTRBaseClusterServiceArea {
	return getMTRBaseClusterServiceAreaClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterServiceArea */
// Cluster Service Area
//
// The Service Area cluster provides an interface for controlling the areas where a device should operate, and for querying the current area being serviced.


// Cluster Service Area
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea
type MTRBaseClusterServiceArea struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterServiceAreaFrom constructs a [MTRBaseClusterServiceArea] from an unsafe.Pointer.
//
// Cluster Service Area
func MTRBaseClusterServiceAreaFrom(ptr unsafe.Pointer) MTRBaseClusterServiceArea {
	return MTRBaseClusterServiceArea{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterServiceArea *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterServiceArea */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterServiceArea */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterServiceArea */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterServiceArea/subscribeAttributeEstimatedEndTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterServiceArea) SubscribeAttributeEstimatedEndTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeEstimatedEndTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeEstimatedEndTimeWithParamsSubscriptionEstablishedReportHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterServiceArea */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterServiceArea */



