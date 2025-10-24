// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMediaPlaybackClusterDeactivateTextTrackParams */


/* debug [class_header]: Header for MTRMediaPlaybackClusterDeactivateTextTrackParams */
// The class instance for the [MTRMediaPlaybackClusterDeactivateTextTrackParams] class.
var (
	MTRMediaPlaybackClusterDeactivateTextTrackParamsClass     _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass
	MTRMediaPlaybackClusterDeactivateTextTrackParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterDeactivateTextTrackParamsClass() _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass {
	MTRMediaPlaybackClusterDeactivateTextTrackParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterDeactivateTextTrackParamsClass = _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass{objc.GetClass("MTRMediaPlaybackClusterDeactivateTextTrackParams")}
	})
	return MTRMediaPlaybackClusterDeactivateTextTrackParamsClass
}

type _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMediaPlaybackClusterDeactivateTextTrackParams */
// An interface definition for the [MTRMediaPlaybackClusterDeactivateTextTrackParams] class.
type IMTRMediaPlaybackClusterDeactivateTextTrackParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMediaPlaybackClusterDeactivateTextTrackParams */
	// properties:
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMediaPlaybackClusterDeactivateTextTrackParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMediaPlaybackClusterDeactivateTextTrackParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass) Alloc() MTRMediaPlaybackClusterDeactivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterDeactivateTextTrackParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMediaPlaybackClusterDeactivateTextTrackParamsClass) New() MTRMediaPlaybackClusterDeactivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterDeactivateTextTrackParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) Init() MTRMediaPlaybackClusterDeactivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterDeactivateTextTrackParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) Autorelease() MTRMediaPlaybackClusterDeactivateTextTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterDeactivateTextTrackParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterDeactivateTextTrackParams creates a new MTRMediaPlaybackClusterDeactivateTextTrackParams instance.
func NewMTRMediaPlaybackClusterDeactivateTextTrackParams() MTRMediaPlaybackClusterDeactivateTextTrackParams {
	return getMTRMediaPlaybackClusterDeactivateTextTrackParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMediaPlaybackClusterDeactivateTextTrackParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterDeactivateTextTrackParams
type MTRMediaPlaybackClusterDeactivateTextTrackParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterDeactivateTextTrackParamsFrom constructs a [MTRMediaPlaybackClusterDeactivateTextTrackParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterDeactivateTextTrackParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterDeactivateTextTrackParams {
	return MTRMediaPlaybackClusterDeactivateTextTrackParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMediaPlaybackClusterDeactivateTextTrackParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMediaPlaybackClusterDeactivateTextTrackParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMediaPlaybackClusterDeactivateTextTrackParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMediaPlaybackClusterDeactivateTextTrackParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMediaPlaybackClusterDeactivateTextTrackParams */

// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterDeactivateTextTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterDeactivateTextTrackParams/serverSideProcessingTimeout
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterdeactivatetexttrackparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusterdeactivatetexttrackparams/timedinvoketimeoutms
func (m_ MTRMediaPlaybackClusterDeactivateTextTrackParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMediaPlaybackClusterDeactivateTextTrackParams */



