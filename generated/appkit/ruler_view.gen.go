// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RulerView] class.
var RulerViewClass objc.Class

func init() {
	RulerViewClass = objc.GetClass("NSRulerView")
}

type RulerView struct {
	objc.ID
}

func RulerViewFrom(ptr unsafe.Pointer) RulerView {
	return RulerView{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (rc RulerView) Alloc() RulerView {
	ret := objc.ID(RulerViewClass).Send(objc.RegisterName("alloc"))
	return RulerView{ret}
}

// New creates and returns a new initialized instance.
func (rc RulerView) New() RulerView {
	ret := objc.ID(RulerViewClass).Send(objc.RegisterName("new"))
	return RulerView{ret}
}

// NewRulerView creates and returns a new initialized instance.
func NewRulerView() RulerView {
	ret := objc.ID(RulerViewClass).Send(objc.RegisterName("new"))
	return RulerView{ret}
}

// Init initializes the instance.
func (r_ RulerView) Init() RulerView {
	ret := r_.ID.Send(objc.RegisterName("init"))
	return RulerView{ret}
}
// Registers a new unit of measurement with the NSRulerView class, making it available to all instances of NSRulerView. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/registerUnit(withName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:)
func (rc RulerView) RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName unsafe.Pointer, abbreviation unsafe.Pointer, conversionFactor float64, stepUpCycle unsafe.Pointer, stepDownCycle unsafe.Pointer) {
	sel := objc.RegisterName("registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:")
	objc.ID(RulerViewClass).Send(sel, unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}
// Adds   to the receiver, without consulting the client view for approval. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/addMarker(_:)
func (r_ RulerView) AddMarker(marker unsafe.Pointer) {
	sel := objc.RegisterName("addMarker:")
	r_.ID.Send(sel, marker)
}
// Draws the receiver’s hash marks and labels in  , which is expressed in the receiver’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/drawHashMarksAndLabels(in:)
func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect unsafe.Pointer) {
	sel := objc.RegisterName("drawHashMarksAndLabelsInRect:")
	r_.ID.Send(sel, rect)
}
// Draws the receiver’s markers in  , which is expressed in the receiver’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/drawMarkers(in:)
func (r_ RulerView) DrawMarkersInRect(rect unsafe.Pointer) {
	sel := objc.RegisterName("drawMarkersInRect:")
	r_.ID.Send(sel, rect)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/init(coder:)
func (r_ RulerView) InitWithCoder(coder unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("initWithCoder:")
	ret := r_.ID.Send(sel, coder)
	return unsafe.Pointer(ret)
}
// Initializes a newly allocated NSRulerView to have   (  or  ) within  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/init(scrollView:orientation:)
func (r_ RulerView) InitWithScrollViewOrientation(scrollView unsafe.Pointer, orientation unsafe.Pointer) unsafe.Pointer {
	sel := objc.RegisterName("initWithScrollView:orientation:")
	ret := r_.ID.Send(sel, scrollView, orientation)
	return unsafe.Pointer(ret)
}
// Forces recalculation of the hash mark spacing for the next time the receiver is displayed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/invalidateHashMarks()
func (r_ RulerView) InvalidateHashMarks() {
	sel := objc.RegisterName("invalidateHashMarks")
	r_.ID.Send(sel)
}
// Draws temporary lines in the ruler area. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/moveRulerline(fromLocation:toLocation:)
func (r_ RulerView) MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64) {
	sel := objc.RegisterName("moveRulerlineFromLocation:toLocation:")
	r_.ID.Send(sel, oldLocation, newLocation)
}
// Removes   from the receiver, without consulting the client view for approval. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/removeMarker(_:)
func (r_ RulerView) RemoveMarker(marker unsafe.Pointer) {
	sel := objc.RegisterName("removeMarker:")
	r_.ID.Send(sel, marker)
}
// Tracks the mouse to add   based on the initial mouse-down or mouse-dragged event  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/trackMarker(_:withMouseEvent:)
func (r_ RulerView) TrackMarkerWithMouseEvent(marker unsafe.Pointer, event unsafe.Pointer) bool {
	sel := objc.RegisterName("trackMarker:withMouseEvent:")
	ret := r_.ID.Send(sel, marker, event)
	return ret != 0
}

