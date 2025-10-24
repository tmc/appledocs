// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRChannelClusterCancelRecordProgramParams */


/* debug [class_header]: Header for MTRChannelClusterCancelRecordProgramParams */
// The class instance for the [MTRChannelClusterCancelRecordProgramParams] class.
var (
	MTRChannelClusterCancelRecordProgramParamsClass     _MTRChannelClusterCancelRecordProgramParamsClass
	MTRChannelClusterCancelRecordProgramParamsClassOnce sync.Once
)

func getMTRChannelClusterCancelRecordProgramParamsClass() _MTRChannelClusterCancelRecordProgramParamsClass {
	MTRChannelClusterCancelRecordProgramParamsClassOnce.Do(func() {
		MTRChannelClusterCancelRecordProgramParamsClass = _MTRChannelClusterCancelRecordProgramParamsClass{objc.GetClass("MTRChannelClusterCancelRecordProgramParams")}
	})
	return MTRChannelClusterCancelRecordProgramParamsClass
}

type _MTRChannelClusterCancelRecordProgramParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRChannelClusterCancelRecordProgramParams */
// An interface definition for the [MTRChannelClusterCancelRecordProgramParams] class.
type IMTRChannelClusterCancelRecordProgramParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRChannelClusterCancelRecordProgramParams */
	// properties:
	ShouldRecordSeries() objc.IObject /* cross-framework: NSNumber */
	SetShouldRecordSeries(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	Data() foundation.Data
	SetData(value foundation.Data)
	ProgramIdentifier() objc.IObject /* cross-framework: NSString */
	SetProgramIdentifier(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRChannelClusterCancelRecordProgramParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRChannelClusterCancelRecordProgramParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterCancelRecordProgramParamsClass) Alloc() MTRChannelClusterCancelRecordProgramParams {
	rv := objc.Send[MTRChannelClusterCancelRecordProgramParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRChannelClusterCancelRecordProgramParamsClass) New() MTRChannelClusterCancelRecordProgramParams {
	rv := objc.Send[MTRChannelClusterCancelRecordProgramParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterCancelRecordProgramParams) Init() MTRChannelClusterCancelRecordProgramParams {
	rv := objc.Send[MTRChannelClusterCancelRecordProgramParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterCancelRecordProgramParams) Autorelease() MTRChannelClusterCancelRecordProgramParams {
	rv := objc.Send[MTRChannelClusterCancelRecordProgramParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterCancelRecordProgramParams creates a new MTRChannelClusterCancelRecordProgramParams instance.
func NewMTRChannelClusterCancelRecordProgramParams() MTRChannelClusterCancelRecordProgramParams {
	return getMTRChannelClusterCancelRecordProgramParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRChannelClusterCancelRecordProgramParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams
type MTRChannelClusterCancelRecordProgramParams struct {
	objectivec.Object
}

// MTRChannelClusterCancelRecordProgramParamsFrom constructs a [MTRChannelClusterCancelRecordProgramParams] from an unsafe.Pointer.
func MTRChannelClusterCancelRecordProgramParamsFrom(ptr unsafe.Pointer) MTRChannelClusterCancelRecordProgramParams {
	return MTRChannelClusterCancelRecordProgramParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRChannelClusterCancelRecordProgramParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRChannelClusterCancelRecordProgramParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRChannelClusterCancelRecordProgramParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRChannelClusterCancelRecordProgramParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRChannelClusterCancelRecordProgramParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/shouldRecordSeries
func (m_ MTRChannelClusterCancelRecordProgramParams) ShouldRecordSeries() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("shouldRecordSeries"))
	return rv
}/* debug [instance_properties/getter]: shouldRecordSeries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/shouldRecordSeries
func (m_ MTRChannelClusterCancelRecordProgramParams) SetShouldRecordSeries(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldRecordSeries:"), value)
}/* debug [instance_properties/setter]: shouldRecordSeries */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterCancelRecordProgramParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterCancelRecordProgramParams/timedInvokeTimeoutMs
func (m_ MTRChannelClusterCancelRecordProgramParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustercancelrecordprogramparams/data
func (m_ MTRChannelClusterCancelRecordProgramParams) Data() foundation.Data {
	rv := objc.Send[foundation.Data](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustercancelrecordprogramparams/data
func (m_ MTRChannelClusterCancelRecordProgramParams) SetData(value foundation.Data) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustercancelrecordprogramparams/programidentifier
func (m_ MTRChannelClusterCancelRecordProgramParams) ProgramIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("programIdentifier"))
	return rv
}/* debug [instance_properties/getter]: programIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustercancelrecordprogramparams/programidentifier
func (m_ MTRChannelClusterCancelRecordProgramParams) SetProgramIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgramIdentifier:"), value)
}/* debug [instance_properties/setter]: programIdentifier */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustercancelrecordprogramparams/serversideprocessingtimeout
func (m_ MTRChannelClusterCancelRecordProgramParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclustercancelrecordprogramparams/serversideprocessingtimeout
func (m_ MTRChannelClusterCancelRecordProgramParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRChannelClusterCancelRecordProgramParams */



