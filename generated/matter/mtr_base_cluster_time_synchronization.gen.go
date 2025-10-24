// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MTRBaseClusterTimeSynchronization */


/* debug [class_header]: Header for MTRBaseClusterTimeSynchronization */
// The class instance for the [MTRBaseClusterTimeSynchronization] class.
var (
	MTRBaseClusterTimeSynchronizationClass     _MTRBaseClusterTimeSynchronizationClass
	MTRBaseClusterTimeSynchronizationClassOnce sync.Once
)

func getMTRBaseClusterTimeSynchronizationClass() _MTRBaseClusterTimeSynchronizationClass {
	MTRBaseClusterTimeSynchronizationClassOnce.Do(func() {
		MTRBaseClusterTimeSynchronizationClass = _MTRBaseClusterTimeSynchronizationClass{objc.GetClass("MTRBaseClusterTimeSynchronization")}
	})
	return MTRBaseClusterTimeSynchronizationClass
}

type _MTRBaseClusterTimeSynchronizationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBaseClusterTimeSynchronization */
// An interface definition for the [MTRBaseClusterTimeSynchronization] class.
type IMTRBaseClusterTimeSynchronization interface {
	IMTRGenericBaseCluster
	
/* debug [class_interface_properties]: Properties for MTRBaseClusterTimeSynchronization */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBaseClusterTimeSynchronization */
	// methods:
	SubscribeAttributeUTCTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBaseClusterTimeSynchronization */
// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterTimeSynchronizationClass) Alloc() MTRBaseClusterTimeSynchronization {
	rv := objc.Send[MTRBaseClusterTimeSynchronization](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBaseClusterTimeSynchronizationClass) New() MTRBaseClusterTimeSynchronization {
	rv := objc.Send[MTRBaseClusterTimeSynchronization](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterTimeSynchronization) Init() MTRBaseClusterTimeSynchronization {
	rv := objc.Send[MTRBaseClusterTimeSynchronization](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterTimeSynchronization) Autorelease() MTRBaseClusterTimeSynchronization {
	rv := objc.Send[MTRBaseClusterTimeSynchronization](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterTimeSynchronization creates a new MTRBaseClusterTimeSynchronization instance.
func NewMTRBaseClusterTimeSynchronization() MTRBaseClusterTimeSynchronization {
	return getMTRBaseClusterTimeSynchronizationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBaseClusterTimeSynchronization */
// Cluster Time Synchronization
//
// Accurate time is required for a number of reasons, including scheduling, display and validating security materials.


// Cluster Time Synchronization
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization
type MTRBaseClusterTimeSynchronization struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterTimeSynchronizationFrom constructs a [MTRBaseClusterTimeSynchronization] from an unsafe.Pointer.
//
// Cluster Time Synchronization
func MTRBaseClusterTimeSynchronizationFrom(ptr unsafe.Pointer) MTRBaseClusterTimeSynchronization {
	return MTRBaseClusterTimeSynchronization{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBaseClusterTimeSynchronization *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBaseClusterTimeSynchronization */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBaseClusterTimeSynchronization */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBaseClusterTimeSynchronization */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterTimeSynchronization/subscribeAttributeUTCTime(with:subscriptionEstablished:reportHandler:)
func (m_ MTRBaseClusterTimeSynchronization) SubscribeAttributeUTCTimeWithParamsSubscriptionEstablishedReportHandler(params IMTRSubscribeParams, subscriptionEstablished unsafe.Pointer, reportHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("subscribeAttributeUTCTimeWithParams:subscriptionEstablished:reportHandler:"), params, subscriptionEstablished, reportHandler)
}/* debug [instance_methods/method]: SubscribeAttributeUTCTimeWithParamsSubscriptionEstablishedReportHandler */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBaseClusterTimeSynchronization */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBaseClusterTimeSynchronization */



