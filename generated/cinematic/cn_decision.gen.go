// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNDecision */


/* debug [class_header]: Header for CNDecision */
// The class instance for the [CNDecision] class.
var (
	CNDecisionClass     _CNDecisionClass
	CNDecisionClassOnce sync.Once
)

func getCNDecisionClass() _CNDecisionClass {
	CNDecisionClassOnce.Do(func() {
		CNDecisionClass = _CNDecisionClass{objc.GetClass("CNDecision")}
	})
	return CNDecisionClass
}

type _CNDecisionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNDecision */
// An interface definition for the [CNDecision] class.
type ICNDecision interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNDecision */
	// properties:
	DetectionGroupID() CNDetectionGroupID /* typedef */
	DetectionID() CNDetectionID /* typedef */
	GroupDecision() bool
	StrongDecision() bool
	Time() objc.IObject /* cross-framework: Time */
	UserDecision() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNDecision */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNDecision */
// Alloc allocates a new instance without initialization.
func (cc _CNDecisionClass) Alloc() CNDecision {
	rv := objc.Send[CNDecision](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNDecisionClass) New() CNDecision {
	rv := objc.Send[CNDecision](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNDecision) Init() CNDecision {
	rv := objc.Send[CNDecision](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNDecision) Autorelease() CNDecision {
	rv := objc.Send[CNDecision](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNDecision creates a new CNDecision instance.
func NewCNDecision() CNDecision {
	return getCNDecisionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNDecision */
// An object that represents a decision to focus on a particular detection, or group of detections, at a particular time.


// An object that represents a decision to focus on a particular detection, or group of detections, at a particular time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class
type CNDecision struct {
	objectivec.Object
}

// CNDecisionFrom constructs a [CNDecision] from an unsafe.Pointer.
//
// An object that represents a decision to focus on a particular detection, or group of detections, at a particular time.
func CNDecisionFrom(ptr unsafe.Pointer) CNDecision {
	return CNDecision{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNDecision */

// Makes a decision to focus on the detection with the given unique detection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class/initWithTime:detectionGroupID:strong:
func NewCNDecisionWithTimeDetectionGroupIDStrong(time objc.IObject /* cross-framework: Time */, detectionGroupID CNDetectionGroupID /* typedef */, isStrong bool) CNDecision {
	instance := getCNDecisionClass().Alloc()
	rv := objc.Send[CNDecision](instance.ID, objc.Sel("initWithTime:detectionGroupID:strong:"), time, detectionGroupID, isStrong)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNDecisionWithTimeDetectionGroupIDStrong */


// Makes a decision to focus on the best among those detections with the same detection group ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class/initWithTime:detectionID:strong:
func NewCNDecisionWithTimeDetectionIDStrong(time objc.IObject /* cross-framework: Time */, detectionID CNDetectionID /* typedef */, isStrong bool) CNDecision {
	instance := getCNDecisionClass().Alloc()
	rv := objc.Send[CNDecision](instance.ID, objc.Sel("initWithTime:detectionID:strong:"), time, detectionID, isStrong)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCNDecisionWithTimeDetectionIDStrong */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNDecision */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNDecision */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNDecision */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNDecision */

// A unique number representing the detection to focus on if this is a group decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class/detectionGroupID
func (c_ CNDecision) DetectionGroupID() CNDetectionGroupID /* typedef */ {
	rv := objc.Send[int64](c_.ID, objc.Sel("detectionGroupID"))
	return rv
}/* debug [instance_properties/getter]: detectionGroupID */


// The unique ID representing the detection to focus on if this isn’t a group decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class/detectionID
func (c_ CNDecision) DetectionID() CNDetectionID /* typedef */ {
	rv := objc.Send[int64](c_.ID, objc.Sel("detectionID"))
	return rv
}/* debug [instance_properties/getter]: detectionID */


// A flag representing whether this is a group decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class/groupDecision
func (c_ CNDecision) GroupDecision() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("groupDecision"))
	return rv
}/* debug [instance_properties/getter]: groupDecision */


// A flag representing whether this is a strong decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class/strongDecision
func (c_ CNDecision) StrongDecision() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("strongDecision"))
	return rv
}/* debug [instance_properties/getter]: strongDecision */


// The first presentation time that the subject should be in focus.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class/time
func (c_ CNDecision) Time() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](c_.ID, objc.Sel("time"))
	return rv
}/* debug [instance_properties/getter]: time */


// A flag representing whether this is a user-created decision or a base decision.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNDecision-c.class/userDecision
func (c_ CNDecision) UserDecision() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("userDecision"))
	return rv
}/* debug [instance_properties/getter]: userDecision */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNDecision */


