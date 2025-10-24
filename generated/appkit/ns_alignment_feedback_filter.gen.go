// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSAlignmentFeedbackFilter */


/* debug [class_header]: Header for NSAlignmentFeedbackFilter */
// The class instance for the [AlignmentFeedbackFilter] class.
var (
	AlignmentFeedbackFilterClass     _AlignmentFeedbackFilterClass
	AlignmentFeedbackFilterClassOnce sync.Once
)

func getAlignmentFeedbackFilterClass() _AlignmentFeedbackFilterClass {
	AlignmentFeedbackFilterClassOnce.Do(func() {
		AlignmentFeedbackFilterClass = _AlignmentFeedbackFilterClass{objc.GetClass("NSAlignmentFeedbackFilter")}
	})
	return AlignmentFeedbackFilterClass
}

type _AlignmentFeedbackFilterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AlignmentFeedbackFilter */
// An interface definition for the [AlignmentFeedbackFilter] class.
type IAlignmentFeedbackFilter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AlignmentFeedbackFilter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AlignmentFeedbackFilter */
	// methods:
	AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX(view IView, previousX float64, alignedX float64, defaultX float64) unsafe.Pointer
	AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint(view IView, previousPoint vision.Point, alignedPoint vision.Point, defaultPoint vision.Point) unsafe.Pointer
	AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY(view IView, previousY float64, alignedY float64, defaultY float64) unsafe.Pointer
	PerformFeedbackPerformanceTime(alignmentFeedbackTokens []objc.ID, performanceTime HapticFeedbackPerformanceTime)
	UpdateWithEvent(event IEvent)
	UpdateWithPanRecognizer(panRecognizer IPanGestureRecognizer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AlignmentFeedbackFilter */
// Alloc allocates a new instance without initialization.
func (ac _AlignmentFeedbackFilterClass) Alloc() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AlignmentFeedbackFilterClass) New() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AlignmentFeedbackFilter) Init() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AlignmentFeedbackFilter) Autorelease() AlignmentFeedbackFilter {
	rv := objc.Send[AlignmentFeedbackFilter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAlignmentFeedbackFilter creates a new AlignmentFeedbackFilter instance.
func NewAlignmentFeedbackFilter() AlignmentFeedbackFilter {
	return getAlignmentFeedbackFilterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AlignmentFeedbackFilter */
// An object that can filter the movement of an object and provides haptic feedback when alignment occurs.
//
// With a Force Touch trackpad, apps can produce tactile feedback to complement user actions. If your app implements alignment features, you can use the class to filter object movements and provide haptic feedback to the user at appropriate times. As the user drags objects into alignment with a guide or another object, the user actually feels a physical bump as the object snaps into place.


// An object that can filter the movement of an object and provides haptic feedback when alignment occurs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter
type AlignmentFeedbackFilter struct {
	objectivec.Object
}

// AlignmentFeedbackFilterFrom constructs a [AlignmentFeedbackFilter] from an unsafe.Pointer.
//
// An object that can filter the movement of an object and provides haptic feedback when alignment occurs.
func AlignmentFeedbackFilterFrom(ptr unsafe.Pointer) AlignmentFeedbackFilter {
	return AlignmentFeedbackFilter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AlignmentFeedbackFilter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AlignmentFeedbackFilter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AlignmentFeedbackFilter */

// Retrieves the event types the filter accepts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/inputEventMask
func (ac _AlignmentFeedbackFilterClass) InputEventMask() EventMask {
	rv := objc.Send[EventMask](objc.ID(ac.class), objc.Sel("inputEventMask"))
	return rv
}/* debug [class_properties_class/property]: inputEventMask */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AlignmentFeedbackFilter */

// Requests a feedback token for the alignment of an object requiring horizontal movement only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/alignmentFeedbackTokenForHorizontalMovement(in:previousX:alignedX:defaultX:)
func (a_ AlignmentFeedbackFilter) AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX(view IView, previousX float64, alignedX float64, defaultX float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("alignmentFeedbackTokenForHorizontalMovementInView:previousX:alignedX:defaultX:"), view, previousX, alignedX, defaultX)
	return rv
}/* debug [instance_methods/method]: AlignmentFeedbackTokenForHorizontalMovementInViewPreviousXAlignedXDefaultX */


// Requests a feedback token for the alignment of an object requiring horizontal and vertical movement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/alignmentFeedbackTokenForMovement(in:previousPoint:alignedPoint:defaultPoint:)
func (a_ AlignmentFeedbackFilter) AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint(view IView, previousPoint vision.Point, alignedPoint vision.Point, defaultPoint vision.Point) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("alignmentFeedbackTokenForMovementInView:previousPoint:alignedPoint:defaultPoint:"), view, previousPoint, alignedPoint, defaultPoint)
	return rv
}/* debug [instance_methods/method]: AlignmentFeedbackTokenForMovementInViewPreviousPointAlignedPointDefaultPoint */


// Requests a feedback token for the alignment of an object requiring vertical movement only.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/alignmentFeedbackTokenForVerticalMovement(in:previousY:alignedY:defaultY:)
func (a_ AlignmentFeedbackFilter) AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY(view IView, previousY float64, alignedY float64, defaultY float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("alignmentFeedbackTokenForVerticalMovementInView:previousY:alignedY:defaultY:"), view, previousY, alignedY, defaultY)
	return rv
}/* debug [instance_methods/method]: AlignmentFeedbackTokenForVerticalMovementInViewPreviousYAlignedYDefaultY */


// Performs the haptic feedback described by one or more alignment feedback tokens.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/performFeedback(_:performanceTime:)
func (a_ AlignmentFeedbackFilter) PerformFeedbackPerformanceTime(alignmentFeedbackTokens []objc.ID, performanceTime HapticFeedbackPerformanceTime) {
	objc.Send[objc.ID](a_.ID, objc.Sel("performFeedback:performanceTime:"), alignmentFeedbackTokens, performanceTime)
}/* debug [instance_methods/method]: PerformFeedbackPerformanceTime */


// Informs the feedback filter about a new event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/update(with:)
func (a_ AlignmentFeedbackFilter) UpdateWithEvent(event IEvent) {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateWithEvent:"), event)
}/* debug [instance_methods/method]: UpdateWithEvent */


// Informs the feedback filter about a new pan (drag) gesture recognizer event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/update(withPanRecognizer:)
func (a_ AlignmentFeedbackFilter) UpdateWithPanRecognizer(panRecognizer IPanGestureRecognizer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("updateWithPanRecognizer:"), panRecognizer)
}/* debug [instance_methods/method]: UpdateWithPanRecognizer */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AlignmentFeedbackFilter */

// Retrieves the event types the filter accepts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackFilter/inputEventMask
func (a_ AlignmentFeedbackFilter) InputEventMask() EventMask {
	rv := objc.Send[EventMask](a_.ID, objc.Sel("inputEventMask"))
	return rv
}/* debug [instance_properties/getter]: inputEventMask */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAlignmentFeedbackFilter */



