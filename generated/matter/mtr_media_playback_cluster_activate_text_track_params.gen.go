// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMediaPlaybackClusterActivateTextTrackParams */


/* debug [class_header]: Header for MTRMediaPlaybackClusterActivateTextTrackParams */
// The class instance for the [MTRMediaPlaybackClusterActivateTextTrackParams] class.
var (
	MTRMediaPlaybackClusterActivateTextTrackParamsClass     _MTRMediaPlaybackClusterActivateTextTrackParamsClass
	MTRMediaPlaybackClusterActivateTextTrackParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterActivateTextTrackParamsClass() _MTRMediaPlaybackClusterActivateTextTrackParamsClass {
	MTRMediaPlaybackClusterActivateTextTrackParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterActivateTextTrackParamsClass = _MTRMediaPlaybackClusterActivateTextTrackParamsClass{objc.GetClass("MTRMediaPlaybackClusterActivateTextTrackParams")}
	})
	return MTRMediaPlaybackClusterActivateTextTrackParamsClass
}

type _MTRMediaPlaybackClusterActivateTextTrackParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMediaPlaybackClusterActivateTextTrackParams */
// An interface definition for the [MTRMediaPlaybackClusterActivateTextTrackParams] class.
type IMTRMediaPlaybackClusterActivateTextTrackParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMediaPlaybackClusterActivateTextTrackParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	TrackID() objc.IObject /* cross-framework: NSString */
	SetTrackID(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMediaPlaybackClusterActivateTextTrackParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMediaPlaybackClusterActivateTextTrackParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterActivateTextTrackParamsClass) Alloc() MTRMediaPlaybackClusterActivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateTextTrackParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMediaPlaybackClusterActivateTextTrackParamsClass) New() MTRMediaPlaybackClusterActivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateTextTrackParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) Init() MTRMediaPlaybackClusterActivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateTextTrackParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) Autorelease() MTRMediaPlaybackClusterActivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateTextTrackParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterActivateTextTrackParams creates a new MTRMediaPlaybackClusterActivateTextTrackParams instance.
func NewMTRMediaPlaybackClusterActivateTextTrackParams() MTRMediaPlaybackClusterActivateTextTrackParams {
	return getMTRMediaPlaybackClusterActivateTextTrackParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMediaPlaybackClusterActivateTextTrackParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams
type MTRMediaPlaybackClusterActivateTextTrackParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterActivateTextTrackParamsFrom constructs a [MTRMediaPlaybackClusterActivateTextTrackParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterActivateTextTrackParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterActivateTextTrackParams {
	return MTRMediaPlaybackClusterActivateTextTrackParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMediaPlaybackClusterActivateTextTrackParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMediaPlaybackClusterActivateTextTrackParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMediaPlaybackClusterActivateTextTrackParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMediaPlaybackClusterActivateTextTrackParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMediaPlaybackClusterActivateTextTrackParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateTextTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusteractivatetexttrackparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusteractivatetexttrackparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusteractivatetexttrackparams/trackid
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) TrackID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("trackID"))
	return rv
}/* debug [instance_properties/getter]: trackID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusteractivatetexttrackparams/trackid
func (m_ MTRMediaPlaybackClusterActivateTextTrackParams) SetTrackID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}/* debug [instance_properties/setter]: trackID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMediaPlaybackClusterActivateTextTrackParams */



