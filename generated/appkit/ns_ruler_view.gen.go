// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSRulerView */


/* debug [class_header]: Header for NSRulerView */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RulerView */
// An interface definition for the [RulerView] class.
type IRulerView interface {
	IView
	
/* debug [class_interface_properties]: Properties for RulerView */
	// properties:
	AccessoryView() IView
	SetAccessoryView(value IView)
	BaselineLocation() float64
	ClientView() IView
	SetClientView(value IView)
	Flipped() bool
	Markers() []RulerMarker
	SetMarkers(value []RulerMarker)
	MeasurementUnits() RulerViewUnitName /* typedef */
	SetMeasurementUnits(value RulerViewUnitName /* typedef */)
	Orientation() RulerOrientation
	SetOrientation(value RulerOrientation)
	OriginOffset() float64
	SetOriginOffset(value float64)
	RequiredThickness() float64
	ReservedThicknessForAccessoryView() float64
	SetReservedThicknessForAccessoryView(value float64)
	ReservedThicknessForMarkers() float64
	SetReservedThicknessForMarkers(value float64)
	RuleThickness() float64
	SetRuleThickness(value float64)
	ScrollView() IScrollView
	SetScrollView(value IScrollView)
	IsFlipped() bool
	SetIsFlipped(value bool)
	HasHorizontalRuler() bool
	SetHasHorizontalRuler(value bool)
	HasVerticalRuler() bool
	SetHasVerticalRuler(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RulerView */
	// methods:
	AddMarker(marker IRulerMarker)
	DrawHashMarksAndLabelsInRect(rect Rect /* not a class type */)
	DrawMarkersInRect(rect Rect /* not a class type */)
	InvalidateHashMarks()
	MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64)
	RemoveMarker(marker IRulerMarker)
	TrackMarkerWithMouseEvent(marker IRulerMarker, event IEvent) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RulerView */
// Alloc allocates a new instance without initialization.
func (rc _RulerViewClass) Alloc() RulerView {
	rv := objc.Send[RulerView](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RulerView */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RulerView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/init(coder:)
func NewRulerViewWithCoder(coder foundation.Coder) RulerView {
	instance := getRulerViewClass().Alloc()
	rv := objc.Send[RulerView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRulerViewWithCoder */


// Initializes a newly allocated NSRulerView to have ( or ) within .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/init(scrollView:orientation:)
func NewRulerViewWithScrollViewOrientation(scrollView IScrollView, orientation RulerOrientation) RulerView {
	instance := getRulerViewClass().Alloc()
	rv := objc.Send[RulerView](instance.ID, objc.Sel("initWithScrollView:orientation:"), scrollView, orientation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRulerViewWithScrollViewOrientation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RulerView */

// Registers a new unit of measurement with the NSRulerView class, making it available to all instances of NSRulerView.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/registerUnit(withName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:)
func (rc _RulerViewClass) RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName RulerViewUnitName /* typedef */, abbreviation objc.IObject /* cross-framework: NSString */, conversionFactor float64, stepUpCycle []foundation.Number, stepDownCycle []foundation.Number) {
	objc.Send[objc.ID](objc.ID(rc.class), objc.Sel("registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:"), unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RulerView */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RulerView */

// Adds to the receiver, without consulting the client view for approval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/addMarker(_:)
func (r_ RulerView) AddMarker(marker IRulerMarker) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addMarker:"), marker)
}/* debug [instance_methods/method]: AddMarker */


// Draws the receiver’s hash marks and labels in , which is expressed in the receiver’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/drawHashMarksAndLabels(in:)
func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("drawHashMarksAndLabelsInRect:"), rect)
}/* debug [instance_methods/method]: DrawHashMarksAndLabelsInRect */


// Draws the receiver’s markers in , which is expressed in the receiver’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/drawMarkers(in:)
func (r_ RulerView) DrawMarkersInRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("drawMarkersInRect:"), rect)
}/* debug [instance_methods/method]: DrawMarkersInRect */


// Forces recalculation of the hash mark spacing for the next time the receiver is displayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/invalidateHashMarks()
func (r_ RulerView) InvalidateHashMarks() {
	objc.Send[objc.ID](r_.ID, objc.Sel("invalidateHashMarks"))
}/* debug [instance_methods/method]: InvalidateHashMarks */


// Draws temporary lines in the ruler area.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/moveRulerline(fromLocation:toLocation:)
func (r_ RulerView) MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("moveRulerlineFromLocation:toLocation:"), oldLocation, newLocation)
}/* debug [instance_methods/method]: MoveRulerlineFromLocationToLocation */


// Removes from the receiver, without consulting the client view for approval.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/removeMarker(_:)
func (r_ RulerView) RemoveMarker(marker IRulerMarker) {
	objc.Send[objc.ID](r_.ID, objc.Sel("removeMarker:"), marker)
}/* debug [instance_methods/method]: RemoveMarker */


// Tracks the mouse to add based on the initial mouse-down or mouse-dragged event .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/trackMarker(_:withMouseEvent:)
func (r_ RulerView) TrackMarkerWithMouseEvent(marker IRulerMarker, event IEvent) bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("trackMarker:withMouseEvent:"), marker, event)
	return rv
}/* debug [instance_methods/method]: TrackMarkerWithMouseEvent */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RulerView */

// The receiver’s accessory view to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/accessoryView
func (r_ RulerView) AccessoryView() IView {
	rv := objc.Send[View](r_.ID, objc.Sel("accessoryView"))
	return rv
}/* debug [instance_properties/getter]: accessoryView */


// The receiver’s accessory view to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/accessoryView
func (r_ RulerView) SetAccessoryView(value IView) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAccessoryView:"), value)
}/* debug [instance_properties/setter]: accessoryView */


// The location of the receiver’s baseline, in its own coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/baselineLocation
func (r_ RulerView) BaselineLocation() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("baselineLocation"))
	return rv
}/* debug [instance_properties/getter]: baselineLocation */


// The receiver’s client view, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/clientView
func (r_ RulerView) ClientView() IView {
	rv := objc.Send[View](r_.ID, objc.Sel("clientView"))
	return rv
}/* debug [instance_properties/getter]: clientView */


// The receiver’s client view, if it has one.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/clientView
func (r_ RulerView) SetClientView(value IView) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setClientView:"), value)
}/* debug [instance_properties/setter]: clientView */


// A Boolean that indicates if the ruler view’s coordinate system is flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/isFlipped
func (r_ RulerView) Flipped() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("flipped"))
	return rv
}/* debug [instance_properties/getter]: flipped */


// The receiver’s ruler markers to , removing any existing ruler markers and not consulting with the client view about the new markers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/markers
func (r_ RulerView) Markers() []RulerMarker {
	rv := objc.Send[[]RulerMarker](r_.ID, objc.Sel("markers"))
	return rv
}/* debug [instance_properties/getter]: markers */


// The receiver’s ruler markers to , removing any existing ruler markers and not consulting with the client view about the new markers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/markers
func (r_ RulerView) SetMarkers(value []RulerMarker) {
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
}/* debug [instance_properties/setter]: markers */


// The measurement units used by the ruler to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/measurementUnits
func (r_ RulerView) MeasurementUnits() RulerViewUnitName /* typedef */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("measurementUnits"))
	return rv
}/* debug [instance_properties/getter]: measurementUnits */


// The measurement units used by the ruler to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/measurementUnits
func (r_ RulerView) SetMeasurementUnits(value RulerViewUnitName /* typedef */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setMeasurementUnits:"), value)
}/* debug [instance_properties/setter]: measurementUnits */


// The orientation of the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/orientation-swift.property
func (r_ RulerView) Orientation() RulerOrientation {
	rv := objc.Send[RulerOrientation](r_.ID, objc.Sel("orientation"))
	return rv
}/* debug [instance_properties/getter]: orientation */


// The orientation of the receiver to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/orientation-swift.property
func (r_ RulerView) SetOrientation(value RulerOrientation) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOrientation:"), value)
}/* debug [instance_properties/setter]: orientation */


// The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/originOffset
func (r_ RulerView) OriginOffset() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("originOffset"))
	return rv
}/* debug [instance_properties/getter]: originOffset */


// The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/originOffset
func (r_ RulerView) SetOriginOffset(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOriginOffset:"), value)
}/* debug [instance_properties/setter]: originOffset */


// The thickness needed for proper tiling of the receiver within an NSScrollView.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/requiredThickness
func (r_ RulerView) RequiredThickness() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("requiredThickness"))
	return rv
}/* debug [instance_properties/getter]: requiredThickness */


// The room available for the receiver’s accessory view to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForAccessoryView
func (r_ RulerView) ReservedThicknessForAccessoryView() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("reservedThicknessForAccessoryView"))
	return rv
}/* debug [instance_properties/getter]: reservedThicknessForAccessoryView */


// The room available for the receiver’s accessory view to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForAccessoryView
func (r_ RulerView) SetReservedThicknessForAccessoryView(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setReservedThicknessForAccessoryView:"), value)
}/* debug [instance_properties/setter]: reservedThicknessForAccessoryView */


// The room available for ruler markers to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForMarkers
func (r_ RulerView) ReservedThicknessForMarkers() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("reservedThicknessForMarkers"))
	return rv
}/* debug [instance_properties/getter]: reservedThicknessForMarkers */


// The room available for ruler markers to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/reservedThicknessForMarkers
func (r_ RulerView) SetReservedThicknessForMarkers(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setReservedThicknessForMarkers:"), value)
}/* debug [instance_properties/setter]: reservedThicknessForMarkers */


// The thickness of the area where ruler hash marks and labels are drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/ruleThickness
func (r_ RulerView) RuleThickness() float64 {
	rv := objc.Send[float64](r_.ID, objc.Sel("ruleThickness"))
	return rv
}/* debug [instance_properties/getter]: ruleThickness */


// The thickness of the area where ruler hash marks and labels are drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/ruleThickness
func (r_ RulerView) SetRuleThickness(value float64) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setRuleThickness:"), value)
}/* debug [instance_properties/setter]: ruleThickness */


// The NSScrollView that owns the receiver to , without retaining it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/scrollView
func (r_ RulerView) ScrollView() IScrollView {
	rv := objc.Send[ScrollView](r_.ID, objc.Sel("scrollView"))
	return rv
}/* debug [instance_properties/getter]: scrollView */


// The NSScrollView that owns the receiver to , without retaining it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/scrollView
func (r_ RulerView) SetScrollView(value IScrollView) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setScrollView:"), value)
}/* debug [instance_properties/setter]: scrollView */


// A Boolean that indicates if the ruler view’s coordinate system is flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/isflipped
func (r_ RulerView) IsFlipped() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("isFlipped"))
	return rv
}/* debug [instance_properties/getter]: isFlipped */


// A Boolean that indicates if the ruler view’s coordinate system is flipped.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsrulerview/isflipped
func (r_ RulerView) SetIsFlipped(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setIsFlipped:"), value)
}/* debug [instance_properties/setter]: isFlipped */


// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hashorizontalruler
func (r_ RulerView) HasHorizontalRuler() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("hasHorizontalRuler"))
	return rv
}/* debug [instance_properties/getter]: hasHorizontalRuler */


// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hashorizontalruler
func (r_ RulerView) SetHasHorizontalRuler(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHasHorizontalRuler:"), value)
}/* debug [instance_properties/setter]: hasHorizontalRuler */


// A Boolean that indicates whether the scroll view keeps a vertical ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hasverticalruler
func (r_ RulerView) HasVerticalRuler() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("hasVerticalRuler"))
	return rv
}/* debug [instance_properties/getter]: hasVerticalRuler */


// A Boolean that indicates whether the scroll view keeps a vertical ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscrollview/hasverticalruler
func (r_ RulerView) SetHasVerticalRuler(value bool) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setHasVerticalRuler:"), value)
}/* debug [instance_properties/setter]: hasVerticalRuler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSRulerView */


