// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRChannelClusterRecordProgramParams */


/* debug [class_header]: Header for MTRChannelClusterRecordProgramParams */
// The class instance for the [MTRChannelClusterRecordProgramParams] class.
var (
	MTRChannelClusterRecordProgramParamsClass     _MTRChannelClusterRecordProgramParamsClass
	MTRChannelClusterRecordProgramParamsClassOnce sync.Once
)

func getMTRChannelClusterRecordProgramParamsClass() _MTRChannelClusterRecordProgramParamsClass {
	MTRChannelClusterRecordProgramParamsClassOnce.Do(func() {
		MTRChannelClusterRecordProgramParamsClass = _MTRChannelClusterRecordProgramParamsClass{objc.GetClass("MTRChannelClusterRecordProgramParams")}
	})
	return MTRChannelClusterRecordProgramParamsClass
}

type _MTRChannelClusterRecordProgramParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRChannelClusterRecordProgramParams */
// An interface definition for the [MTRChannelClusterRecordProgramParams] class.
type IMTRChannelClusterRecordProgramParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRChannelClusterRecordProgramParams */
	// properties:
	Data() objc.IObject /* cross-framework: NSData */
	SetData(value objc.IObject /* cross-framework: NSData */)
	ProgramIdentifier() objc.IObject /* cross-framework: NSString */
	SetProgramIdentifier(value objc.IObject /* cross-framework: NSString */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	ShouldRecordSeries() objc.IObject /* cross-framework: NSNumber */
	SetShouldRecordSeries(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRChannelClusterRecordProgramParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRChannelClusterRecordProgramParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRChannelClusterRecordProgramParamsClass) Alloc() MTRChannelClusterRecordProgramParams {
	rv := objc.Send[MTRChannelClusterRecordProgramParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRChannelClusterRecordProgramParamsClass) New() MTRChannelClusterRecordProgramParams {
	rv := objc.Send[MTRChannelClusterRecordProgramParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRChannelClusterRecordProgramParams) Init() MTRChannelClusterRecordProgramParams {
	rv := objc.Send[MTRChannelClusterRecordProgramParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRChannelClusterRecordProgramParams) Autorelease() MTRChannelClusterRecordProgramParams {
	rv := objc.Send[MTRChannelClusterRecordProgramParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRChannelClusterRecordProgramParams creates a new MTRChannelClusterRecordProgramParams instance.
func NewMTRChannelClusterRecordProgramParams() MTRChannelClusterRecordProgramParams {
	return getMTRChannelClusterRecordProgramParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRChannelClusterRecordProgramParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams
type MTRChannelClusterRecordProgramParams struct {
	objectivec.Object
}

// MTRChannelClusterRecordProgramParamsFrom constructs a [MTRChannelClusterRecordProgramParams] from an unsafe.Pointer.
func MTRChannelClusterRecordProgramParamsFrom(ptr unsafe.Pointer) MTRChannelClusterRecordProgramParams {
	return MTRChannelClusterRecordProgramParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRChannelClusterRecordProgramParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRChannelClusterRecordProgramParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRChannelClusterRecordProgramParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRChannelClusterRecordProgramParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRChannelClusterRecordProgramParams */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/data
func (m_ MTRChannelClusterRecordProgramParams) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRChannelClusterRecordProgramParams/data
func (m_ MTRChannelClusterRecordProgramParams) SetData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterrecordprogramparams/programidentifier
func (m_ MTRChannelClusterRecordProgramParams) ProgramIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("programIdentifier"))
	return rv
}/* debug [instance_properties/getter]: programIdentifier */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterrecordprogramparams/programidentifier
func (m_ MTRChannelClusterRecordProgramParams) SetProgramIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProgramIdentifier:"), value)
}/* debug [instance_properties/setter]: programIdentifier */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterrecordprogramparams/serversideprocessingtimeout
func (m_ MTRChannelClusterRecordProgramParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterrecordprogramparams/serversideprocessingtimeout
func (m_ MTRChannelClusterRecordProgramParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterrecordprogramparams/shouldrecordseries
func (m_ MTRChannelClusterRecordProgramParams) ShouldRecordSeries() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("shouldRecordSeries"))
	return rv
}/* debug [instance_properties/getter]: shouldRecordSeries */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterrecordprogramparams/shouldrecordseries
func (m_ MTRChannelClusterRecordProgramParams) SetShouldRecordSeries(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setShouldRecordSeries:"), value)
}/* debug [instance_properties/setter]: shouldRecordSeries */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterrecordprogramparams/timedinvoketimeoutms
func (m_ MTRChannelClusterRecordProgramParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrchannelclusterrecordprogramparams/timedinvoketimeoutms
func (m_ MTRChannelClusterRecordProgramParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRChannelClusterRecordProgramParams */



