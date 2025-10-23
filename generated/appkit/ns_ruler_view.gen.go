// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [RulerView] class.
var (
	RulerViewClass     _RulerViewClass
	RulerViewClassOnce sync.Once
)

func getRulerViewClass() _RulerViewClass {
	RulerViewClassOnce.Do(func() {
		RulerViewClass = _RulerViewClass{objc.GetClass("NSRulerView")}
	})
	return RulerViewClass
}

type _RulerViewClass struct {
	class objc.Class
}

// An interface definition for the [RulerView] class.
type IRulerView interface {
	IView
	// properties:
	AccessoryView() IView
	SetAccessoryView(value IView)
	BaselineLocation() float64 /* primitive/slice/pointer. */
	ClientView() IView
	SetClientView(value IView)
	Flipped() bool /* primitive/slice/pointer. */
	Markers() []RulerMarker /* primitive/slice/pointer. */
	SetMarkers(value []RulerMarker /* primitive/slice/pointer. */)
	MeasurementUnits() objc.IObject /* cross-framework: RulerViewUnitName */
	SetMeasurementUnits(value objc.IObject /* cross-framework: RulerViewUnitName */)
	Orientation() RulerOrientation
	SetOrientation(value RulerOrientation)
	OriginOffset() float64 /* primitive/slice/pointer. */
	SetOriginOffset(value float64 /* primitive/slice/pointer. */)
	RequiredThickness() float64 /* primitive/slice/pointer. */
	ReservedThicknessForAccessoryView() float64 /* primitive/slice/pointer. */
	SetReservedThicknessForAccessoryView(value float64 /* primitive/slice/pointer. */)
	ReservedThicknessForMarkers() float64 /* primitive/slice/pointer. */
	SetReservedThicknessForMarkers(value float64 /* primitive/slice/pointer. */)
	RuleThickness() float64 /* primitive/slice/pointer. */
	SetRuleThickness(value float64 /* primitive/slice/pointer. */)
	ScrollView() IScrollView
	SetScrollView(value IScrollView)
	IsFlipped() bool /* primitive/slice/pointer. */
	SetIsFlipped(value bool /* primitive/slice/pointer. */)
	HasHorizontalRuler() bool /* primitive/slice/pointer. */
	SetHasHorizontalRuler(value bool /* primitive/slice/pointer. */)
	HasVerticalRuler() bool /* primitive/slice/pointer. */
	SetHasVerticalRuler(value bool /* primitive/slice/pointer. */)
	// methods:
	AddMarker(marker IRulerMarker)
	DrawHashMarksAndLabelsInRect(rect objc.IObject /* cross-framework Rect */)
	DrawMarkersInRect(rect objc.IObject /* cross-framework Rect */)
	InvalidateHashMarks()
	MoveRulerlineFromLocationToLocation(oldLocation float64 /* primitive/slice/pointer. */, newLocation float64 /* primitive/slice/pointer. */)
	RemoveMarker(marker IRulerMarker)
	TrackMarkerWithMouseEvent(marker IRulerMarker, event IEvent) bool /* primitive/slice/pointer. */
}

// A ruler and the markers above or to the side of a scroll view’s document view.
//
// Views within the scroll view can become clients of the ruler view, having it display markers for their elements, and receiving messages from the ruler view when the user manipulates the markers.


// A ruler and the markers above or to the side of a scroll view’s document view.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getRulerViewClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/init(coder:)
func NewRulerViewWithCoder(coder objc.IObject /* cross-framework Coder */) RulerView {
	instance := getRulerViewClass().Alloc()
	rv := objc.Send[RulerView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Initializes a newly allocated NSRulerView to have ( or ) within .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/init(scrollView:orientation:)
func NewRulerViewWithScrollViewOrientation(scrollView IScrollView, orientation RulerOrientation) RulerView {
	instance := getRulerViewClass().Alloc()
	rv := objc.Send[RulerView](instance.ID, objc.Sel("initWithScrollView:orientation:"), scrollView, orientation)
	rv.Autorelease()
	return rv
}



// Registers a new unit of measurement with the NSRulerView class, making it available to all instances of NSRulerView.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/registerUnit(withName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:)
func (rc _RulerViewClass) RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName objc.IObject /* cross-framework RulerViewUnitName */, abbreviation objc.IObject /* cross-framework NSString */, conversionFactor float64 /* primitive/slice/pointer. */, stepUpCycle objc.IObject /* cross-framework Number */, stepDownCycle objc.IObject /* cross-framework Number */) {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:"), unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}


// Adds to the receiver, without consulting the client view for approval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/addMarker(_:)
func (r_ RulerView) AddMarker(marker IRulerMarker) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addMarker:"), marker)
}


// Draws the receiver’s hash marks and labels in , which is expressed in the receiver’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/drawHashMarksAndLabels(in:)
func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect objc.IObject /* cross-framework Rect */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("drawHashMarksAndLabelsInRect:"), rect)
}


// Draws the receiver’s markers in , which is expressed in the receiver’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/drawMarkers(in:)
func (r_ RulerView) DrawMarkersInRect(rect objc.IObject /* cross-framework Rect */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("drawMarkersInRect:"), rect)
}


// Forces recalculation of the hash mark spacing for the next time the receiver is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/invalidateHashMarks()
func (r_ RulerView) InvalidateHashMarks() {
	objc.Send[objc.ID](r_.ID, objc.Sel("invalidateHashMarks"))
}


// Draws temporary lines in the ruler area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/moveRulerline(fromLocation:toLocation:)
func (r_ RulerView) MoveRulerlineFromLocationToLocation(oldLocation float64 /* primitive/slice/pointer. */, newLocation float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("moveRulerlineFromLocation:toLocation:"), oldLocation, newLocation)
}


// Removes from the receiver, without consulting the client view for approval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/removeMarker(_:)
func (r_ RulerView) RemoveMarker(marker IRulerMarker) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeMarker:"), marker)
}


// Tracks the mouse to add based on the initial mouse-down or mouse-dragged event .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/trackMarker(_:withMouseEvent:)
func (r_ RulerView) TrackMarkerWithMouseEvent(marker IRulerMarker, event IEvent) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("trackMarker:withMouseEvent:"), marker, event)
	return rv
}


// The receiver’s accessory view to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/accessoryView
func (r_ RulerView) AccessoryView() IView {
	rv := objc.Send[View](r_.ID, objc.Sel("accessoryView"))
	return rv
}


// The receiver’s accessory view to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/accessoryView
func (r_ RulerView) SetAccessoryView(value IView) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAccessoryView:"), value)
}


// The location of the receiver’s baseline, in its own coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/baselineLocation
func (r_ RulerView) BaselineLocation() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](r_.ID, objc.Sel("baselineLocation"))
	return rv
}


// The receiver’s client view, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/clientView
func (r_ RulerView) ClientView() IView {
	rv := objc.Send[View](r_.ID, objc.Sel("clientView"))
	return rv
}


// The receiver’s client view, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/clientView
func (r_ RulerView) SetClientView(value IView) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setClientView:"), value)
}


// A Boolean that indicates if the ruler view’s coordinate system is flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/isFlipped
func (r_ RulerView) Flipped() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("flipped"))
	return rv
}


// The receiver’s ruler markers to , removing any existing ruler markers and not consulting with the client view about the new markers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/markers
func (r_ RulerView) Markers() []RulerMarker /* primitive/slice/pointer. */ {
	rv := objc.Send[[]RulerMarker](r_.ID, objc.Sel("markers"))
	return rv
}


// The receiver’s ruler markers to , removing any existing ruler markers and not consulting with the client view about the new markers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/markers
func (r_ RulerView) SetMarkers(value []RulerMarker /* primitive/slice/pointer. */) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](r_.ID, objc.Sel("setMarkers:"), nsArray)
}


// The measurement units used by the ruler to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/measurementUnits
func (r_ RulerView) MeasurementUnits() objc.IObject /* cross-framework: RulerViewUnitName */ {
	rv := objc.Send[RulerViewUnitName](r_.ID, objc.Sel("measurementUnits"))
	return rv
}


// The measurement units used by the ruler to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/measurementUnits
func (r_ RulerView) SetMeasurementUnits(value objc.IObject /* cross-framework: RulerViewUnitName */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMeasurementUnits:"), value)
}


// The orientation of the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/orientation-swift.property
func (r_ RulerView) Orientation() RulerOrientation {
	rv := objc.Send[RulerOrientation](r_.ID, objc.Sel("orientation"))
	return rv
}


// The orientation of the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/orientation-swift.property
func (r_ RulerView) SetOrientation(value RulerOrientation) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOrientation:"), value)
}


// The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/originOffset
func (r_ RulerView) OriginOffset() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](r_.ID, objc.Sel("originOffset"))
	return rv
}


// The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/originOffset
func (r_ RulerView) SetOriginOffset(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOriginOffset:"), value)
}


// The thickness needed for proper tiling of the receiver within an NSScrollView.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/requiredThickness
func (r_ RulerView) RequiredThickness() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](r_.ID, objc.Sel("requiredThickness"))
	return rv
}


// The room available for the receiver’s accessory view to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForAccessoryView
func (r_ RulerView) ReservedThicknessForAccessoryView() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](r_.ID, objc.Sel("reservedThicknessForAccessoryView"))
	return rv
}


// The room available for the receiver’s accessory view to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForAccessoryView
func (r_ RulerView) SetReservedThicknessForAccessoryView(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setReservedThicknessForAccessoryView:"), value)
}


// The room available for ruler markers to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForMarkers
func (r_ RulerView) ReservedThicknessForMarkers() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](r_.ID, objc.Sel("reservedThicknessForMarkers"))
	return rv
}


// The room available for ruler markers to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForMarkers
func (r_ RulerView) SetReservedThicknessForMarkers(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setReservedThicknessForMarkers:"), value)
}


// The thickness of the area where ruler hash marks and labels are drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/ruleThickness
func (r_ RulerView) RuleThickness() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](r_.ID, objc.Sel("ruleThickness"))
	return rv
}


// The thickness of the area where ruler hash marks and labels are drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/ruleThickness
func (r_ RulerView) SetRuleThickness(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRuleThickness:"), value)
}


// The NSScrollView that owns the receiver to , without retaining it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/scrollView
func (r_ RulerView) ScrollView() IScrollView {
	rv := objc.Send[ScrollView](r_.ID, objc.Sel("scrollView"))
	return rv
}


// The NSScrollView that owns the receiver to , without retaining it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/scrollView
func (r_ RulerView) SetScrollView(value IScrollView) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setScrollView:"), value)
}


// A Boolean that indicates if the ruler view’s coordinate system is flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/isflipped
func (r_ RulerView) IsFlipped() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("isFlipped"))
	return rv
}


// A Boolean that indicates if the ruler view’s coordinate system is flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/isflipped
func (r_ RulerView) SetIsFlipped(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsFlipped:"), value)
}


// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hashorizontalruler
func (r_ RulerView) HasHorizontalRuler() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("hasHorizontalRuler"))
	return rv
}


// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hashorizontalruler
func (r_ RulerView) SetHasHorizontalRuler(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHasHorizontalRuler:"), value)
}


// A Boolean that indicates whether the scroll view keeps a vertical ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hasverticalruler
func (r_ RulerView) HasVerticalRuler() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](r_.ID, objc.Sel("hasVerticalRuler"))
	return rv
}


// A Boolean that indicates whether the scroll view keeps a vertical ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hasverticalruler
func (r_ RulerView) SetHasVerticalRuler(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHasVerticalRuler:"), value)
}


