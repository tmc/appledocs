// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
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
	AddMarker(marker unsafe.Pointer)
	DrawHashMarksAndLabelsInRect(rect coregraphics.CGRect)
	DrawMarkersInRect(rect coregraphics.CGRect)
	InvalidateHashMarks()
	MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64)
	RemoveMarker(marker unsafe.Pointer)
	TrackMarkerWithMouseEvent(marker unsafe.Pointer, event unsafe.Pointer) bool
}

// A ruler and the markers above or to the side of a scroll view’s document view.
//
// Views within the scroll view can become clients of the ruler view, having it display markers for their elements, and receiving messages from the ruler view when the user manipulates the markers.
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


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/init(coder:)
func NewRulerViewWithCoder(coder unsafe.Pointer) RulerView {
	instance := getRulerViewClass().Alloc()
	rv := objc.Send[RulerView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Initializes a newly allocated NSRulerView to have ( or ) within .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/init(scrollView:orientation:)
func NewRulerViewWithScrollViewOrientation(scrollView unsafe.Pointer, orientation unsafe.Pointer) RulerView {
	instance := getRulerViewClass().Alloc()
	rv := objc.Send[RulerView](instance.ID, objc.Sel("initWithScrollView:orientation:"), scrollView, orientation)
	rv.Autorelease()
	return rv
}


// Registers a new unit of measurement with the NSRulerView class, making it available to all instances of NSRulerView.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/registerUnit(withName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:)
func (rc _RulerViewClass) RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName unsafe.Pointer, abbreviation string, conversionFactor float64, stepUpCycle unsafe.Pointer, stepDownCycle unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:"), unitName, objc.String(abbreviation), conversionFactor, stepUpCycle, stepDownCycle)
}

// Adds to the receiver, without consulting the client view for approval.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/addMarker(_:)
func (r_ RulerView) AddMarker(marker unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addMarker:"), marker)
}

// Draws the receiver’s hash marks and labels in , which is expressed in the receiver’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/drawHashMarksAndLabels(in:)
func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](r_.ID, objc.Sel("drawHashMarksAndLabelsInRect:"), rect)
}

// Draws the receiver’s markers in , which is expressed in the receiver’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/drawMarkers(in:)
func (r_ RulerView) DrawMarkersInRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](r_.ID, objc.Sel("drawMarkersInRect:"), rect)
}

// Forces recalculation of the hash mark spacing for the next time the receiver is displayed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/invalidateHashMarks()
func (r_ RulerView) InvalidateHashMarks() {
	objc.Send[objc.ID](r_.ID, objc.Sel("invalidateHashMarks"))
}

// Draws temporary lines in the ruler area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/moveRulerline(fromLocation:toLocation:)
func (r_ RulerView) MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("moveRulerlineFromLocation:toLocation:"), oldLocation, newLocation)
}

// Removes from the receiver, without consulting the client view for approval.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/removeMarker(_:)
func (r_ RulerView) RemoveMarker(marker unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeMarker:"), marker)
}

// Tracks the mouse to add based on the initial mouse-down or mouse-dragged event .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/trackMarker(_:withMouseEvent:)
func (r_ RulerView) TrackMarkerWithMouseEvent(marker unsafe.Pointer, event unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("trackMarker:withMouseEvent:"), marker, event)
	return rv
}

// The receiver’s accessory view to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/accessoryView
func (r_ RulerView) AccessoryView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("accessoryView"))
	return rv
}


// SetAccessoryView sets the value of the accessoryView property.
// The receiver’s accessory view to .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/accessoryView
func (r_ RulerView) SetAccessoryView(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAccessoryView:"), value)
}

// The location of the receiver’s baseline, in its own coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/baselineLocation
func (r_ RulerView) BaselineLocation() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("baselineLocation"))
	return rv
}

// The receiver’s client view, if it has one.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/clientView
func (r_ RulerView) ClientView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("clientView"))
	return rv
}


// SetClientView sets the value of the clientView property.
// The receiver’s client view, if it has one.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/clientView
func (r_ RulerView) SetClientView(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setClientView:"), value)
}

// A Boolean that indicates if the ruler view’s coordinate system is flipped.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/isFlipped
func (r_ RulerView) Flipped() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("flipped"))
	return rv
}

// The receiver’s ruler markers to , removing any existing ruler markers and not consulting with the client view about the new markers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/markers
func (r_ RulerView) Markers() []RulerMarker {
	rv := objc.Send[[]RulerMarker](r_.ID, objc.Sel("markers"))
	return rv
}


// SetMarkers sets the value of the markers property.
// The receiver’s ruler markers to , removing any existing ruler markers and not consulting with the client view about the new markers.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/markers
func (r_ RulerView) SetMarkers(value []RulerMarker) {
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/measurementUnits
func (r_ RulerView) MeasurementUnits() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("measurementUnits"))
	return rv
}


// SetMeasurementUnits sets the value of the measurementUnits property.
// The measurement units used by the ruler to .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/measurementUnits
func (r_ RulerView) SetMeasurementUnits(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMeasurementUnits:"), value)
}

// The orientation of the receiver to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/orientation-swift.property
func (r_ RulerView) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("orientation"))
	return rv
}


// SetOrientation sets the value of the orientation property.
// The orientation of the receiver to .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/orientation-swift.property
func (r_ RulerView) SetOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOrientation:"), value)
}

// The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/originOffset
func (r_ RulerView) OriginOffset() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("originOffset"))
	return rv
}


// SetOriginOffset sets the value of the originOffset property.
// The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/originOffset
func (r_ RulerView) SetOriginOffset(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOriginOffset:"), value)
}

// The thickness needed for proper tiling of the receiver within an NSScrollView.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/requiredThickness
func (r_ RulerView) RequiredThickness() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("requiredThickness"))
	return rv
}

// The room available for the receiver’s accessory view to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForAccessoryView
func (r_ RulerView) ReservedThicknessForAccessoryView() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("reservedThicknessForAccessoryView"))
	return rv
}


// SetReservedThicknessForAccessoryView sets the value of the reservedThicknessForAccessoryView property.
// The room available for the receiver’s accessory view to .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForAccessoryView
func (r_ RulerView) SetReservedThicknessForAccessoryView(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setReservedThicknessForAccessoryView:"), value)
}

// The room available for ruler markers to .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForMarkers
func (r_ RulerView) ReservedThicknessForMarkers() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("reservedThicknessForMarkers"))
	return rv
}


// SetReservedThicknessForMarkers sets the value of the reservedThicknessForMarkers property.
// The room available for ruler markers to .

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForMarkers
func (r_ RulerView) SetReservedThicknessForMarkers(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setReservedThicknessForMarkers:"), value)
}

// The thickness of the area where ruler hash marks and labels are drawn.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/ruleThickness
func (r_ RulerView) RuleThickness() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("ruleThickness"))
	return rv
}


// SetRuleThickness sets the value of the ruleThickness property.
// The thickness of the area where ruler hash marks and labels are drawn.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/ruleThickness
func (r_ RulerView) SetRuleThickness(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRuleThickness:"), value)
}

// The NSScrollView that owns the receiver to , without retaining it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/scrollView
func (r_ RulerView) ScrollView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("scrollView"))
	return rv
}


// SetScrollView sets the value of the scrollView property.
// The NSScrollView that owns the receiver to , without retaining it.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/scrollView
func (r_ RulerView) SetScrollView(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setScrollView:"), value)
}

// A Boolean that indicates if the ruler view’s coordinate system is flipped.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/isflipped
func (r_ RulerView) IsFlipped() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isFlipped"))
	return rv
}


// SetIsFlipped sets the value of the isFlipped property.
// A Boolean that indicates if the ruler view’s coordinate system is flipped.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/isflipped
func (r_ RulerView) SetIsFlipped(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsFlipped:"), value)
}

// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hashorizontalruler
func (r_ RulerView) HasHorizontalRuler() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("hasHorizontalRuler"))
	return rv
}


// SetHasHorizontalRuler sets the value of the hasHorizontalRuler property.
// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hashorizontalruler
func (r_ RulerView) SetHasHorizontalRuler(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHasHorizontalRuler:"), value)
}

// A Boolean that indicates whether the scroll view keeps a vertical ruler object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hasverticalruler
func (r_ RulerView) HasVerticalRuler() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("hasVerticalRuler"))
	return rv
}


// SetHasVerticalRuler sets the value of the hasVerticalRuler property.
// A Boolean that indicates whether the scroll view keeps a vertical ruler object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hasverticalruler
func (r_ RulerView) SetHasVerticalRuler(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHasVerticalRuler:"), value)
}


