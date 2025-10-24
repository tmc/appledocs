// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptureReactionEffectState */


/* debug [class_header]: Header for AVCaptureReactionEffectState */
// The class instance for the [CaptureReactionEffectState] class.
var (
	CaptureReactionEffectStateClass     _CaptureReactionEffectStateClass
	CaptureReactionEffectStateClassOnce sync.Once
)

func getCaptureReactionEffectStateClass() _CaptureReactionEffectStateClass {
	CaptureReactionEffectStateClassOnce.Do(func() {
		CaptureReactionEffectStateClass = _CaptureReactionEffectStateClass{objc.GetClass("AVCaptureReactionEffectState")}
	})
	return CaptureReactionEffectStateClass
}

type _CaptureReactionEffectStateClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptureReactionEffectState */
// An interface definition for the [CaptureReactionEffectState] class.
type ICaptureReactionEffectState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptureReactionEffectState */
	// properties:
	EndTime() objc.IObject /* cross-framework: Time */
	ReactionType() CaptureReactionType /* typedef */
	StartTime() objc.IObject /* cross-framework: Time */
	AvailableReactionTypes() CaptureReactionType /* typedef */
	SetAvailableReactionTypes(value CaptureReactionType /* typedef */)
	CanPerformReactionEffects() bool
	SetCanPerformReactionEffects(value bool)
	ReactionEffectsInProgress() IAVCaptureReactionEffectState
	SetReactionEffectsInProgress(value IAVCaptureReactionEffectState)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptureReactionEffectState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptureReactionEffectState */
// Alloc allocates a new instance without initialization.
func (cc _CaptureReactionEffectStateClass) Alloc() CaptureReactionEffectState {
	rv := objc.Send[CaptureReactionEffectState](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptureReactionEffectStateClass) New() CaptureReactionEffectState {
	rv := objc.Send[CaptureReactionEffectState](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureReactionEffectState) Init() CaptureReactionEffectState {
	rv := objc.Send[CaptureReactionEffectState](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureReactionEffectState) Autorelease() CaptureReactionEffectState {
	rv := objc.Send[CaptureReactionEffectState](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureReactionEffectState creates a new CaptureReactionEffectState instance.
func NewCaptureReactionEffectState() CaptureReactionEffectState {
	return getCaptureReactionEffectStateClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptureReactionEffectState */
// An object that reports the state of a reaction effect performed on a capture device.
//
// Obtain an instance of this class by querying a capture device’s property. The system adds new entries to this array when you call or by gesture detection in the capture stream when the value of is . The system renders the effect before providing frames to your app, and these status objects let you know when it performs the effect.


// An object that reports the state of a reaction effect performed on a capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionEffectState
type CaptureReactionEffectState struct {
	objectivec.Object
}

// CaptureReactionEffectStateFrom constructs a [CaptureReactionEffectState] from an unsafe.Pointer.
//
// An object that reports the state of a reaction effect performed on a capture device.
func CaptureReactionEffectStateFrom(ptr unsafe.Pointer) CaptureReactionEffectState {
	return CaptureReactionEffectState{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptureReactionEffectState *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptureReactionEffectState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptureReactionEffectState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptureReactionEffectState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptureReactionEffectState */

// The presentation time of the first frame following the end of a reaction effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionEffectState/endTime
func (c_ CaptureReactionEffectState) EndTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("endTime"))
	return rv
}/* debug [instance_properties/getter]: endTime */


// The type of reaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionEffectState/reactionType
func (c_ CaptureReactionEffectState) ReactionType() CaptureReactionType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("reactionType"))
	return rv
}/* debug [instance_properties/getter]: reactionType */


// The presentation time of the first frame where the system renders the effect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionEffectState/startTime
func (c_ CaptureReactionEffectState) StartTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("startTime"))
	return rv
}/* debug [instance_properties/getter]: startTime */


// A set of reactions types that a device supports performing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/availablereactiontypes
func (c_ CaptureReactionEffectState) AvailableReactionTypes() CaptureReactionType /* typedef */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("availableReactionTypes"))
	return rv
}/* debug [instance_properties/getter]: availableReactionTypes */


// A set of reactions types that a device supports performing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/availablereactiontypes
func (c_ CaptureReactionEffectState) SetAvailableReactionTypes(value CaptureReactionType /* typedef */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableReactionTypes:"), value)
}/* debug [instance_properties/setter]: availableReactionTypes */


// A Boolean value that indicates whether you can perform reaction effects on a capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/canperformreactioneffects
func (c_ CaptureReactionEffectState) CanPerformReactionEffects() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canPerformReactionEffects"))
	return rv
}/* debug [instance_properties/getter]: canPerformReactionEffects */


// A Boolean value that indicates whether you can perform reaction effects on a capture device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/canperformreactioneffects
func (c_ CaptureReactionEffectState) SetCanPerformReactionEffects(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCanPerformReactionEffects:"), value)
}/* debug [instance_properties/setter]: canPerformReactionEffects */


// An array of reaction effects that the device is currently performing, sorted by timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/reactioneffectsinprogress
func (c_ CaptureReactionEffectState) ReactionEffectsInProgress() IAVCaptureReactionEffectState {
	rv := objc.Send[CaptureReactionEffectState](c_.ID, objc.Sel("reactionEffectsInProgress"))
	return rv
}/* debug [instance_properties/getter]: reactionEffectsInProgress */


// An array of reaction effects that the device is currently performing, sorted by timestamp.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedevice/reactioneffectsinprogress
func (c_ CaptureReactionEffectState) SetReactionEffectsInProgress(value IAVCaptureReactionEffectState) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setReactionEffectsInProgress:"), value)
}/* debug [instance_properties/setter]: reactionEffectsInProgress */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptureReactionEffectState */



