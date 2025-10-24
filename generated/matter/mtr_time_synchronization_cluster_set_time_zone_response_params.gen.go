// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterSetTimeZoneResponseParams */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */
// The class instance for the [MTRTimeSynchronizationClusterSetTimeZoneResponseParams] class.
var (
	MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass     _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass
	MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass() _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass {
	MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClassOnce.Do(func() {
		MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass = _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass{objc.GetClass("MTRTimeSynchronizationClusterSetTimeZoneResponseParams")}
	})
	return MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass
}

type _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */
// An interface definition for the [MTRTimeSynchronizationClusterSetTimeZoneResponseParams] class.
type IMTRTimeSynchronizationClusterSetTimeZoneResponseParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */
	// properties:
	DstOffsetRequired() objc.IObject /* cross-framework: NSNumber */
	SetDstOffsetRequired(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass) Alloc() MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneResponseParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass) New() MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneResponseParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterSetTimeZoneResponseParams) Init() MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneResponseParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterSetTimeZoneResponseParams) Autorelease() MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneResponseParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterSetTimeZoneResponseParams creates a new MTRTimeSynchronizationClusterSetTimeZoneResponseParams instance.
func NewMTRTimeSynchronizationClusterSetTimeZoneResponseParams() MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	return getMTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneResponseParams
type MTRTimeSynchronizationClusterSetTimeZoneResponseParams struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterSetTimeZoneResponseParamsFrom constructs a [MTRTimeSynchronizationClusterSetTimeZoneResponseParams] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterSetTimeZoneResponseParamsFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	return MTRTimeSynchronizationClusterSetTimeZoneResponseParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */

// Initialize an MTRTimeSynchronizationClusterSetTimeZoneResponseParams with a response-value dictionary of the sort that MTRDeviceResponseHandler would receive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterSetTimeZoneResponseParams/init(responseValue:)
func NewMTRTimeSynchronizationClusterSetTimeZoneResponseParamsWithResponseValueError(responseValue foundation.IDictionary, error_ unsafe.Pointer) MTRTimeSynchronizationClusterSetTimeZoneResponseParams {
	instance := getMTRTimeSynchronizationClusterSetTimeZoneResponseParamsClass().Alloc()
	rv := objc.Send[MTRTimeSynchronizationClusterSetTimeZoneResponseParams](instance.ID, objc.Sel("initWithResponseValue:error:"), responseValue, error_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRTimeSynchronizationClusterSetTimeZoneResponseParamsWithResponseValueError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterSetTimeZoneResponseParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersettimezoneresponseparams/dstoffsetrequired
func (m_ MTRTimeSynchronizationClusterSetTimeZoneResponseParams) DstOffsetRequired() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("dstOffsetRequired"))
	return rv
}/* debug [instance_properties/getter]: dstOffsetRequired */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustersettimezoneresponseparams/dstoffsetrequired
func (m_ MTRTimeSynchronizationClusterSetTimeZoneResponseParams) SetDstOffsetRequired(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDstOffsetRequired:"), value)
}/* debug [instance_properties/setter]: dstOffsetRequired */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterSetTimeZoneResponseParams */


