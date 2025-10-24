// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRMediaPlaybackClusterActivateAudioTrackParams */


/* debug [class_header]: Header for MTRMediaPlaybackClusterActivateAudioTrackParams */
// The class instance for the [MTRMediaPlaybackClusterActivateAudioTrackParams] class.
var (
	MTRMediaPlaybackClusterActivateAudioTrackParamsClass     _MTRMediaPlaybackClusterActivateAudioTrackParamsClass
	MTRMediaPlaybackClusterActivateAudioTrackParamsClassOnce sync.Once
)

func getMTRMediaPlaybackClusterActivateAudioTrackParamsClass() _MTRMediaPlaybackClusterActivateAudioTrackParamsClass {
	MTRMediaPlaybackClusterActivateAudioTrackParamsClassOnce.Do(func() {
		MTRMediaPlaybackClusterActivateAudioTrackParamsClass = _MTRMediaPlaybackClusterActivateAudioTrackParamsClass{objc.GetClass("MTRMediaPlaybackClusterActivateAudioTrackParams")}
	})
	return MTRMediaPlaybackClusterActivateAudioTrackParamsClass
}

type _MTRMediaPlaybackClusterActivateAudioTrackParamsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRMediaPlaybackClusterActivateAudioTrackParams */
// An interface definition for the [MTRMediaPlaybackClusterActivateAudioTrackParams] class.
type IMTRMediaPlaybackClusterActivateAudioTrackParams interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRMediaPlaybackClusterActivateAudioTrackParams */
	// properties:
	TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */
	SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */)
	AudioOutputIndex() objc.IObject /* cross-framework: NSNumber */
	SetAudioOutputIndex(value objc.IObject /* cross-framework: NSNumber */)
	ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */
	SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */)
	TrackID() objc.IObject /* cross-framework: NSString */
	SetTrackID(value objc.IObject /* cross-framework: NSString */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRMediaPlaybackClusterActivateAudioTrackParams */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRMediaPlaybackClusterActivateAudioTrackParams */
// Alloc allocates a new instance without initialization.
func (mc _MTRMediaPlaybackClusterActivateAudioTrackParamsClass) Alloc() MTRMediaPlaybackClusterActivateAudioTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateAudioTrackParams](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRMediaPlaybackClusterActivateAudioTrackParamsClass) New() MTRMediaPlaybackClusterActivateAudioTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateAudioTrackParams](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) Init() MTRMediaPlaybackClusterActivateAudioTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateAudioTrackParams](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) Autorelease() MTRMediaPlaybackClusterActivateAudioTrackParams {
	rv := objc.Send[MTRMediaPlaybackClusterActivateAudioTrackParams](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRMediaPlaybackClusterActivateAudioTrackParams creates a new MTRMediaPlaybackClusterActivateAudioTrackParams instance.
func NewMTRMediaPlaybackClusterActivateAudioTrackParams() MTRMediaPlaybackClusterActivateAudioTrackParams {
	return getMTRMediaPlaybackClusterActivateAudioTrackParamsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRMediaPlaybackClusterActivateAudioTrackParams */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams
type MTRMediaPlaybackClusterActivateAudioTrackParams struct {
	objectivec.Object
}

// MTRMediaPlaybackClusterActivateAudioTrackParamsFrom constructs a [MTRMediaPlaybackClusterActivateAudioTrackParams] from an unsafe.Pointer.
func MTRMediaPlaybackClusterActivateAudioTrackParamsFrom(ptr unsafe.Pointer) MTRMediaPlaybackClusterActivateAudioTrackParams {
	return MTRMediaPlaybackClusterActivateAudioTrackParams{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRMediaPlaybackClusterActivateAudioTrackParams *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRMediaPlaybackClusterActivateAudioTrackParams */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRMediaPlaybackClusterActivateAudioTrackParams */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRMediaPlaybackClusterActivateAudioTrackParams */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRMediaPlaybackClusterActivateAudioTrackParams */

// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) TimedInvokeTimeoutMs() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("timedInvokeTimeoutMs"))
	return rv
}/* debug [instance_properties/getter]: timedInvokeTimeoutMs */


// Controls whether the command is a timed command (using Timed Invoke).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRMediaPlaybackClusterActivateAudioTrackParams/timedInvokeTimeoutMs
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetTimedInvokeTimeoutMs(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTimedInvokeTimeoutMs:"), value)
}/* debug [instance_properties/setter]: timedInvokeTimeoutMs */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusteractivateaudiotrackparams/audiooutputindex
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) AudioOutputIndex() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("audioOutputIndex"))
	return rv
}/* debug [instance_properties/getter]: audioOutputIndex */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusteractivateaudiotrackparams/audiooutputindex
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetAudioOutputIndex(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setAudioOutputIndex:"), value)
}/* debug [instance_properties/setter]: audioOutputIndex */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusteractivateaudiotrackparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) ServerSideProcessingTimeout() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("serverSideProcessingTimeout"))
	return rv
}/* debug [instance_properties/getter]: serverSideProcessingTimeout */


// Controls how much time, in seconds, we will allow for the server to process the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusteractivateaudiotrackparams/serversideprocessingtimeout
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetServerSideProcessingTimeout(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setServerSideProcessingTimeout:"), value)
}/* debug [instance_properties/setter]: serverSideProcessingTimeout */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusteractivateaudiotrackparams/trackid
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) TrackID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("trackID"))
	return rv
}/* debug [instance_properties/getter]: trackID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrmediaplaybackclusteractivateaudiotrackparams/trackid
func (m_ MTRMediaPlaybackClusterActivateAudioTrackParams) SetTrackID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTrackID:"), value)
}/* debug [instance_properties/setter]: trackID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRMediaPlaybackClusterActivateAudioTrackParams */



