// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */


/* debug [class_header]: Header for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */
// The class instance for the [MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams] class.
var (
	MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass     _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass
	MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClassOnce sync.Once
)

func getMTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass() _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass {
	MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClassOnce.Do(func() {
		MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass = _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass{objc.GetClass("MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams")}
	})
	return MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass
}

type _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */
// An interface definition for the [MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams] class.
type IMTRElectricalMeasurementClusterGetMeasurementProfileCommandParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */
	// properties:
	AttributeId() objc.IObject /* cross-framework: NSNumber */
	SetAttributeId(value objc.IObject /* cross-framework: NSNumber */)
	NumberOfIntervals() objc.IObject /* cross-framework: NSNumber */
	SetNumberOfIntervals(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	StartTime() objc.IObject /* cross-framework: NSNumber */
	SetStartTime(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass) Alloc() MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass) New() MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) Init() MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) Autorelease() MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	rv := objc.Send[MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRElectricalMeasurementClusterGetMeasurementProfileCommandParams creates a new MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams instance.
func NewMTRElectricalMeasurementClusterGetMeasurementProfileCommandParams() MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	return getMTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams
type MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams struct {
	objectivec.Object
}

// MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsFrom constructs a [MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams] from an unsafe.Pointer.
func MTRElectricalMeasurementClusterGetMeasurementProfileCommandParamsFrom(ptr unsafe.Pointer) MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams {
	return MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams/attributeId
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) AttributeId() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("attributeId"))
	return rv
}/* debug [instance_properties/getter]: attributeId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams/attributeId
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) SetAttributeId(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAttributeId:"), value)
}/* debug [instance_properties/setter]: attributeId */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams/numberOfIntervals
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) NumberOfIntervals() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("numberOfIntervals"))
	return rv
}/* debug [instance_properties/getter]: numberOfIntervals */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams/numberOfIntervals
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) SetNumberOfIntervals(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNumberOfIntervals:"), value)
}/* debug [instance_properties/setter]: numberOfIntervals */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams/serverSideProcessingTimeout
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams/serverSideProcessingTimeout
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams/startTime
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) StartTime() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("startTime"))
	return rv
}/* debug [instance_properties/getter]: startTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams/startTime
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) SetStartTime(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setStartTime:"), value)
}/* debug [instance_properties/setter]: startTime */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams/timedInvokeTimeoutMs
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams/timedInvokeTimeoutMs
func (m_ MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRElectricalMeasurementClusterGetMeasurementProfileCommandParams */



