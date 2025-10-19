// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [RulerView] class.
var rulerViewClass = _RulerViewClass{objc.GetClass("NSRulerView")}

type _RulerViewClass struct {
	class objc.Class
}

// An interface definition for the [RulerView] class.
type IRulerView interface {
	IView
	AddMarker(marker unsafe.Pointer)
	DrawHashMarksAndLabelsInRect(rect unsafe.Pointer)
	DrawMarkersInRect(rect unsafe.Pointer)
	InvalidateHashMarks()
	MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64)
	RemoveMarker(marker unsafe.Pointer)
	TrackMarkerWithMouseEvent(marker unsafe.Pointer, event unsafe.Pointer) bool
}

// A ruler and the markers above or to the side of a scroll view’s document view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView

type RulerView struct {
	View
}

// RulerViewFrom constructs a [RulerView] from an unsafe.Pointer.
//
// A ruler and the markers above or to the side of a scroll view’s document view.
func RulerViewFrom(ptr unsafe.Pointer) RulerView {
	return RulerView{
		View: ViewFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (rc _RulerViewClass) Alloc() RulerView {
	rv := objc.Send[RulerView](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (rc _RulerViewClass) New() RulerView {
	rv := objc.Send[RulerView](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RulerView) Init() RulerView {
	rv := objc.Send[RulerView](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RulerView) Autorelease() RulerView {
	rv := objc.Send[RulerView](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRulerView creates a new RulerView instance.
func NewRulerView() RulerView {
	return rulerViewClass.New()
}
// Initializes a newly allocated NSRulerView to have ( or ) within . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/init(scrollView:orientation:)
func NewRulerViewWithScrollViewOrientation(scrollView unsafe.Pointer, orientation unsafe.Pointer) RulerView {
	instance := rulerViewClass.Alloc()
	rv := objc.Send[RulerView](instance.ID, objc.Sel("initWithScrollView:orientation:"), scrollView, orientation)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/init(coder:)
func NewRulerViewWithCoder(coder unsafe.Pointer) RulerView {
	instance := rulerViewClass.Alloc()
	rv := objc.Send[RulerView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Registers a new unit of measurement with the NSRulerView class, making it available to all instances of NSRulerView. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/registerUnit(withName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:)
func (rc _RulerViewClass) RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName unsafe.Pointer, abbreviation string, conversionFactor float64, stepUpCycle unsafe.Pointer, stepDownCycle unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:"), unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}
// Adds to the receiver, without consulting the client view for approval. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/addMarker(_:)
func (r_ RulerView) AddMarker(marker unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addMarker:"), marker)
}
// Draws the receiver’s hash marks and labels in , which is expressed in the receiver’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/drawHashMarksAndLabels(in:)
func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("drawHashMarksAndLabelsInRect:"), rect)
}
// Draws the receiver’s markers in , which is expressed in the receiver’s coordinate system. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/drawMarkers(in:)
func (r_ RulerView) DrawMarkersInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("drawMarkersInRect:"), rect)
}
// Forces recalculation of the hash mark spacing for the next time the receiver is displayed. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/invalidateHashMarks()
func (r_ RulerView) InvalidateHashMarks() {
	objc.Send[objc.ID](r_.ID, objc.Sel("invalidateHashMarks"))
}
// Draws temporary lines in the ruler area. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/moveRulerline(fromLocation:toLocation:)
func (r_ RulerView) MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("moveRulerlineFromLocation:toLocation:"), oldLocation, newLocation)
}
// Removes from the receiver, without consulting the client view for approval. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/removeMarker(_:)
func (r_ RulerView) RemoveMarker(marker unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeMarker:"), marker)
}
// Tracks the mouse to add based on the initial mouse-down or mouse-dragged event . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/trackMarker(_:withMouseEvent:)
func (r_ RulerView) TrackMarkerWithMouseEvent(marker unsafe.Pointer, event unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("trackMarker:withMouseEvent:"), marker, event)
	return rv
}

