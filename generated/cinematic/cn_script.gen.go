// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNScript */


/* debug [class_header]: Header for CNScript */
// The class instance for the [CNScript] class.
var (
	CNScriptClass     _CNScriptClass
	CNScriptClassOnce sync.Once
)

func getCNScriptClass() _CNScriptClass {
	CNScriptClassOnce.Do(func() {
		CNScriptClass = _CNScriptClass{objc.GetClass("CNScript")}
	})
	return CNScriptClass
}

type _CNScriptClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNScript */
// An interface definition for the [CNScript] class.
type ICNScript interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNScript */
	// properties:
	AddedDetectionTracks() []CNDetectionTrack
	FNumber() float32
	SetFNumber(value float32)
	TimeRange() TimeRange /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNScript */
	// methods:
	AddDetectionTrack(detectionTrack ICNDetectionTrack) CNDetectionID /* typedef */
	AddUserDecision(decision ICNDecision) bool
	BaseDecisionsInTimeRange(timeRange TimeRange /* not a class type */) []CNDecision
	Changes() ICNScriptChanges
	ChangesTrimmedByTimeRange(timeRange TimeRange /* not a class type */) ICNScriptChanges
	DecisionAfterTime(time objc.IObject /* cross-framework: Time */) ICNDecision
	DecisionAtTimeTolerance(time objc.IObject /* cross-framework: Time */, tolerance objc.IObject /* cross-framework: Time */) ICNDecision
	DecisionBeforeTime(time objc.IObject /* cross-framework: Time */) ICNDecision
	DecisionsInTimeRange(timeRange TimeRange /* not a class type */) []CNDecision
	DetectionTrackForDecision(decision ICNDecision) ICNDetectionTrack
	DetectionTrackForID(detectionID CNDetectionID /* typedef */) ICNDetectionTrack
	FrameAtTimeTolerance(time objc.IObject /* cross-framework: Time */, tolerance objc.IObject /* cross-framework: Time */) ICNScriptFrame
	FramesInTimeRange(timeRange TimeRange /* not a class type */) []CNScriptFrame
	PrimaryDecisionAtTime(time objc.IObject /* cross-framework: Time */) ICNDecision
	ReloadWithChanges(changes ICNScriptChanges)
	RemoveAllUserDecisions()
	RemoveDetectionTrack(detectionTrack ICNDetectionTrack) bool
	RemoveUserDecision(decision ICNDecision) bool
	SecondaryDecisionAtTime(time objc.IObject /* cross-framework: Time */) ICNDecision
	TimeRangeOfTransitionAfterDecision(decision ICNDecision) TimeRange /* not a class type */
	TimeRangeOfTransitionBeforeDecision(decision ICNDecision) TimeRange /* not a class type */
	UserDecisionsInTimeRange(timeRange TimeRange /* not a class type */) []CNDecision
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNScript */
// Alloc allocates a new instance without initialization.
func (cc _CNScriptClass) Alloc() CNScript {
	rv := objc.Send[CNScript](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNScriptClass) New() CNScript {
	rv := objc.Send[CNScript](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNScript) Init() CNScript {
	rv := objc.Send[CNScript](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNScript) Autorelease() CNScript {
	rv := objc.Send[CNScript](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNScript creates a new CNScript instance.
func NewCNScript() CNScript {
	return getCNScriptClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNScript */
// A collection of focus decisions, focus transitions, detections, and detection tracks associated with a movie captured in Cinematic mode and methods to change them.
//
// The Cinematic script provides thread-safe access to information about the focus decisions made in the original recorded Cinematic movie. The script supports changing those decisions and obtaining updated information about where to focus each frame. You can snapshot changes to a script and later reload them.


// A collection of focus decisions, focus transitions, detections, and detection tracks associated with a movie captured in Cinematic mode and methods to change them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn
type CNScript struct {
	objectivec.Object
}

// CNScriptFrom constructs a [CNScript] from an unsafe.Pointer.
//
// A collection of focus decisions, focus transitions, detections, and detection tracks associated with a movie captured in Cinematic mode and methods to change them.
func CNScriptFrom(ptr unsafe.Pointer) CNScript {
	return CNScript{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNScript *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNScript */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/loadFromAsset:changes:progress:completionHandler:
func (cc _CNScriptClass) LoadFromAssetChangesProgressCompletionHandler(asset avfoundation.Asset, changes ICNScriptChanges, progress foundation.Progress, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadFromAsset:changes:progress:completionHandler:"), asset, changes, progress, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadFromAssetChangesProgressCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNScript */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNScript */

// Adds a user-created detection track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/addDetectionTrack:
func (c_ CNScript) AddDetectionTrack(detectionTrack ICNDetectionTrack) CNDetectionID /* typedef */ {
	rv := objc.Send[int64](c_.ID, objc.Sel("addDetectionTrack:"), detectionTrack)
	return rv
}/* debug [instance_methods/method]: AddDetectionTrack */


// Adds a new user decision, and replaces an existing user decision if the times are identical.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/addUserDecision:
func (c_ CNScript) AddUserDecision(decision ICNDecision) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("addUserDecision:"), decision)
	return rv
}/* debug [instance_methods/method]: AddUserDecision */


// All base decisions made automatically during recording in the given time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/baseDecisionsInTimeRange:
func (c_ CNScript) BaseDecisionsInTimeRange(timeRange TimeRange /* not a class type */) []CNDecision {
	rv := objc.Send[[]CNDecision](c_.ID, objc.Sel("baseDecisionsInTimeRange:"), timeRange)
	return rv
}/* debug [instance_methods/method]: BaseDecisionsInTimeRange */


// Changes made since recording the Cinematic asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/changes
func (c_ CNScript) Changes() ICNScriptChanges {
	rv := objc.Send[CNScriptChanges](c_.ID, objc.Sel("changes"))
	return rv
}/* debug [instance_methods/method]: Changes */


// Changes trimmed and time range shifted to start at zero.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/changesTrimmedByTimeRange:
func (c_ CNScript) ChangesTrimmedByTimeRange(timeRange TimeRange /* not a class type */) ICNScriptChanges {
	rv := objc.Send[CNScriptChanges](c_.ID, objc.Sel("changesTrimmedByTimeRange:"), timeRange)
	return rv
}/* debug [instance_methods/method]: ChangesTrimmedByTimeRange */


// The decision that occurs after the given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/decisionAfterTime:
func (c_ CNScript) DecisionAfterTime(time objc.IObject /* cross-framework: Time */) ICNDecision {
	rv := objc.Send[CNDecision](c_.ID, objc.Sel("decisionAfterTime:"), time)
	return rv
}/* debug [instance_methods/method]: DecisionAfterTime */


// The closest frame to the given time within the given tolerance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/decisionAtTime:tolerance:
func (c_ CNScript) DecisionAtTimeTolerance(time objc.IObject /* cross-framework: Time */, tolerance objc.IObject /* cross-framework: Time */) ICNDecision {
	rv := objc.Send[CNDecision](c_.ID, objc.Sel("decisionAtTime:tolerance:"), time, tolerance)
	return rv
}/* debug [instance_methods/method]: DecisionAtTimeTolerance */


// The decision that occurs before the given time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/decisionBeforeTime:
func (c_ CNScript) DecisionBeforeTime(time objc.IObject /* cross-framework: Time */) ICNDecision {
	rv := objc.Send[CNDecision](c_.ID, objc.Sel("decisionBeforeTime:"), time)
	return rv
}/* debug [instance_methods/method]: DecisionBeforeTime */


// All decisions within the given time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/decisionsInTimeRange:
func (c_ CNScript) DecisionsInTimeRange(timeRange TimeRange /* not a class type */) []CNDecision {
	rv := objc.Send[[]CNDecision](c_.ID, objc.Sel("decisionsInTimeRange:"), timeRange)
	return rv
}/* debug [instance_methods/method]: DecisionsInTimeRange */


// A detection track representing all detections selected by a given decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/detectionTrackForDecision:
func (c_ CNScript) DetectionTrackForDecision(decision ICNDecision) ICNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](c_.ID, objc.Sel("detectionTrackForDecision:"), decision)
	return rv
}/* debug [instance_methods/method]: DetectionTrackForDecision */


// A detection track representing all detections with the given detection ID, over the entire Cinematic script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/detectionTrackForID:
func (c_ CNScript) DetectionTrackForID(detectionID CNDetectionID /* typedef */) ICNDetectionTrack {
	rv := objc.Send[CNDetectionTrack](c_.ID, objc.Sel("detectionTrackForID:"), detectionID)
	return rv
}/* debug [instance_methods/method]: DetectionTrackForID */


// The closest frame to the given time within the given tolerance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/frameAtTime:tolerance:
func (c_ CNScript) FrameAtTimeTolerance(time objc.IObject /* cross-framework: Time */, tolerance objc.IObject /* cross-framework: Time */) ICNScriptFrame {
	rv := objc.Send[CNScriptFrame](c_.ID, objc.Sel("frameAtTime:tolerance:"), time, tolerance)
	return rv
}/* debug [instance_methods/method]: FrameAtTimeTolerance */


// All frames within the given time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/framesInTimeRange:
func (c_ CNScript) FramesInTimeRange(timeRange TimeRange /* not a class type */) []CNScriptFrame {
	rv := objc.Send[[]CNScriptFrame](c_.ID, objc.Sel("framesInTimeRange:"), timeRange)
	return rv
}/* debug [instance_methods/method]: FramesInTimeRange */


// The primary decision that’s in effect at the specified time, unless it’s outside the time range of the Cinematic script.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/primaryDecisionAtTime:
func (c_ CNScript) PrimaryDecisionAtTime(time objc.IObject /* cross-framework: Time */) ICNDecision {
	rv := objc.Send[CNDecision](c_.ID, objc.Sel("primaryDecisionAtTime:"), time)
	return rv
}/* debug [instance_methods/method]: PrimaryDecisionAtTime */


// Reloads the Cinematic script with optional changes applied, removing any previous changes made.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/reloadWithChanges:
func (c_ CNScript) ReloadWithChanges(changes ICNScriptChanges) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadWithChanges:"), changes)
}/* debug [instance_methods/method]: ReloadWithChanges */


// Removes all user decisions and reverts to base decisions only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/removeAllUserDecisions
func (c_ CNScript) RemoveAllUserDecisions() {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeAllUserDecisions"))
}/* debug [instance_methods/method]: RemoveAllUserDecisions */


// Removes the user-created detection track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/removeDetectionTrack:
func (c_ CNScript) RemoveDetectionTrack(detectionTrack ICNDetectionTrack) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("removeDetectionTrack:"), detectionTrack)
	return rv
}/* debug [instance_methods/method]: RemoveDetectionTrack */


// Removes an existing user decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/removeUserDecision:
func (c_ CNScript) RemoveUserDecision(decision ICNDecision) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("removeUserDecision:"), decision)
	return rv
}/* debug [instance_methods/method]: RemoveUserDecision */


// If a given time is during a focus transition, the system transitions toward a secondary decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/secondaryDecisionAtTime:
func (c_ CNScript) SecondaryDecisionAtTime(time objc.IObject /* cross-framework: Time */) ICNDecision {
	rv := objc.Send[CNDecision](c_.ID, objc.Sel("secondaryDecisionAtTime:"), time)
	return rv
}/* debug [instance_methods/method]: SecondaryDecisionAtTime */


// The time range during which the focus transitioned away from the given decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/timeRangeOfTransitionAfterDecision:
func (c_ CNScript) TimeRangeOfTransitionAfterDecision(decision ICNDecision) TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRangeOfTransitionAfterDecision:"), decision)
	return rv
}/* debug [instance_methods/method]: TimeRangeOfTransitionAfterDecision */


// The time range during which the focus transitioned toward the given decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/timeRangeOfTransitionBeforeDecision:
func (c_ CNScript) TimeRangeOfTransitionBeforeDecision(decision ICNDecision) TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRangeOfTransitionBeforeDecision:"), decision)
	return rv
}/* debug [instance_methods/method]: TimeRangeOfTransitionBeforeDecision */


// All user decisions in the given time range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/userDecisionsInTimeRange:
func (c_ CNScript) UserDecisionsInTimeRange(timeRange TimeRange /* not a class type */) []CNDecision {
	rv := objc.Send[[]CNDecision](c_.ID, objc.Sel("userDecisionsInTimeRange:"), timeRange)
	return rv
}/* debug [instance_methods/method]: UserDecisionsInTimeRange */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNScript */

// An array of the detection tracks added since recording the original Cinematic movie.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/addedDetectionTracks
func (c_ CNScript) AddedDetectionTracks() []CNDetectionTrack {
	rv := objc.Send[[]CNDetectionTrack](c_.ID, objc.Sel("addedDetectionTracks"))
	return rv
}/* debug [instance_properties/getter]: addedDetectionTracks */


// The f-stop value that inversely affects the aperture used to render the Cinematic image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/fNumber
func (c_ CNScript) FNumber() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("fNumber"))
	return rv
}/* debug [instance_properties/getter]: fNumber */


// The f-stop value that inversely affects the aperture used to render the Cinematic image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/fNumber
func (c_ CNScript) SetFNumber(value float32) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setFNumber:"), value)
}/* debug [instance_properties/setter]: fNumber */


// The time range of the Cinematic asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScript-9e1zn/timeRange
func (c_ CNScript) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNScript */



