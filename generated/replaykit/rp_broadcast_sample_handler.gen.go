// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class RPBroadcastSampleHandler */


/* debug [class_header]: Header for RPBroadcastSampleHandler */
// The class instance for the [RPBroadcastSampleHandler] class.
var (
	RPBroadcastSampleHandlerClass     _RPBroadcastSampleHandlerClass
	RPBroadcastSampleHandlerClassOnce sync.Once
)

func getRPBroadcastSampleHandlerClass() _RPBroadcastSampleHandlerClass {
	RPBroadcastSampleHandlerClassOnce.Do(func() {
		RPBroadcastSampleHandlerClass = _RPBroadcastSampleHandlerClass{objc.GetClass("RPBroadcastSampleHandler")}
	})
	return RPBroadcastSampleHandlerClass
}

type _RPBroadcastSampleHandlerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RPBroadcastSampleHandler */
// An interface definition for the [RPBroadcastSampleHandler] class.
type IRPBroadcastSampleHandler interface {
	IRPBroadcastHandler
	
/* debug [class_interface_properties]: Properties for RPBroadcastSampleHandler */
	// properties:
	RPApplicationInfoBundleIdentifierKey() objc.IObject /* cross-framework: NSString */
	RPVideoSampleOrientationKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RPBroadcastSampleHandler */
	// methods:
	BroadcastAnnotatedWithApplicationInfo(applicationInfo objc.IObject /* cross-framework: NSDictionary */)
	BroadcastFinished()
	BroadcastPaused()
	BroadcastResumed()
	BroadcastStartedWithSetupInfo(setupInfo foundation.IDictionary)
	FinishBroadcastWithError(error_ objc.IObject /* cross-framework: Error */)
	ProcessSampleBufferWithType(sampleBuffer SampleBufferRef /* not a class type */, sampleBufferType RPSampleBufferType)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RPBroadcastSampleHandler */
// Alloc allocates a new instance without initialization.
func (rc _RPBroadcastSampleHandlerClass) Alloc() RPBroadcastSampleHandler {
	rv := objc.Send[RPBroadcastSampleHandler](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RPBroadcastSampleHandlerClass) New() RPBroadcastSampleHandler {
	rv := objc.Send[RPBroadcastSampleHandler](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RPBroadcastSampleHandler) Init() RPBroadcastSampleHandler {
	rv := objc.Send[RPBroadcastSampleHandler](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RPBroadcastSampleHandler) Autorelease() RPBroadcastSampleHandler {
	rv := objc.Send[RPBroadcastSampleHandler](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRPBroadcastSampleHandler creates a new RPBroadcastSampleHandler instance.
func NewRPBroadcastSampleHandler() RPBroadcastSampleHandler {
	return getRPBroadcastSampleHandlerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RPBroadcastSampleHandler */
// An object that processes buffer objects as received from ReplayKit.
//
// To handle objects as captured by ReplayKit, you subclass . You enable this mode of handling by setting in the extension’s file to . In your subclass, implement the method to handle video and audio buffers, as well as the , , , and methods to handle starting and stopping the broadcast. ReplayKit invokes methods in your subclass in a serial fashion. After invoking one method, ReplayKit won’t invoke another method until the first method returns. That means it’s safe for your implementations to update their stored state without the use of locks or synchronization to provide thread safety.


// An object that processes buffer objects as received from ReplayKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler
type RPBroadcastSampleHandler struct {
	RPBroadcastHandler
}

// RPBroadcastSampleHandlerFrom constructs a [RPBroadcastSampleHandler] from an unsafe.Pointer.
//
// An object that processes buffer objects as received from ReplayKit.
func RPBroadcastSampleHandlerFrom(ptr unsafe.Pointer) RPBroadcastSampleHandler {
	return RPBroadcastSampleHandler{
		RPBroadcastHandler: RPBroadcastHandlerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RPBroadcastSampleHandler *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RPBroadcastSampleHandler */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RPBroadcastSampleHandler */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RPBroadcastSampleHandler */

// Perform any required actions after starting a live broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/broadcastAnnotated(withApplicationInfo:)
func (r_ RPBroadcastSampleHandler) BroadcastAnnotatedWithApplicationInfo(applicationInfo objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("broadcastAnnotatedWithApplicationInfo:"), applicationInfo)
}/* debug [instance_methods/method]: BroadcastAnnotatedWithApplicationInfo */


// Perform any required actions after a live broadcast is finished.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/broadcastFinished()
func (r_ RPBroadcastSampleHandler) BroadcastFinished() {
	objc.Send[objc.ID](r_.ID, objc.Sel("broadcastFinished"))
}/* debug [instance_methods/method]: BroadcastFinished */


// Perform any required actions after a live broadcast is paused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/broadcastPaused()
func (r_ RPBroadcastSampleHandler) BroadcastPaused() {
	objc.Send[objc.ID](r_.ID, objc.Sel("broadcastPaused"))
}/* debug [instance_methods/method]: BroadcastPaused */


// Perform any required actions after a live broadcast is resumed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/broadcastResumed()
func (r_ RPBroadcastSampleHandler) BroadcastResumed() {
	objc.Send[objc.ID](r_.ID, objc.Sel("broadcastResumed"))
}/* debug [instance_methods/method]: BroadcastResumed */


// Perform any required actions after starting a live broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/broadcastStarted(withSetupInfo:)
func (r_ RPBroadcastSampleHandler) BroadcastStartedWithSetupInfo(setupInfo foundation.IDictionary) {
	objc.Send[objc.ID](r_.ID, objc.Sel("broadcastStartedWithSetupInfo:"), setupInfo)
}/* debug [instance_methods/method]: BroadcastStartedWithSetupInfo */


// Stops the broadcast and passes an error back to the broadcasting app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/finishBroadcastWithError(_:)
func (r_ RPBroadcastSampleHandler) FinishBroadcastWithError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("finishBroadcastWithError:"), error_)
}/* debug [instance_methods/method]: FinishBroadcastWithError */


// Processes video and audio data as it becomes available during a live broadcast.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastSampleHandler/processSampleBuffer(_:with:)
func (r_ RPBroadcastSampleHandler) ProcessSampleBufferWithType(sampleBuffer SampleBufferRef /* not a class type */, sampleBufferType RPSampleBufferType) {
	objc.Send[objc.ID](r_.ID, objc.Sel("processSampleBuffer:withType:"), sampleBuffer, sampleBufferType)
}/* debug [instance_methods/method]: ProcessSampleBufferWithType */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RPBroadcastSampleHandler */

// The key to retrieve the app’s bundle identifier from the user-information dictionary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpapplicationinfobundleidentifierkey
func (r_ RPBroadcastSampleHandler) RPApplicationInfoBundleIdentifierKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("RPApplicationInfoBundleIdentifierKey"))
	return rv
}/* debug [instance_properties/getter]: RPApplicationInfoBundleIdentifierKey */


// The sample attachment key that describes the video orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/replaykit/rpvideosampleorientationkey
func (r_ RPBroadcastSampleHandler) RPVideoSampleOrientationKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("RPVideoSampleOrientationKey"))
	return rv
}/* debug [instance_properties/getter]: RPVideoSampleOrientationKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class RPBroadcastSampleHandler */



