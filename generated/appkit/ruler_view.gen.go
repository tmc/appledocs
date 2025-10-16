
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [RulerView] class.
var RulerViewClass _RulerViewClass

func init() {
	RulerViewClass = _RulerViewClass{objc.GetClass("NSRulerView")}
}

type _RulerViewClass struct {
	objc.Class
}

// An interface definition for the [RulerView] class.
type IRulerView interface {
	ID() objc.ID
	AddMarker(marker unsafe.Pointer)
	DrawHashMarksAndLabelsInRect(rect unsafe.Pointer)
	DrawMarkersInRect(rect unsafe.Pointer)
	InitWithCoder(coder unsafe.Pointer) unsafe.Pointer
	InitWithScrollViewOrientation(scrollView unsafe.Pointer, orientation unsafe.Pointer) unsafe.Pointer
	InvalidateHashMarks()
	MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64)
	RemoveMarker(marker unsafe.Pointer)
	TrackMarkerWithMouseEvent(marker unsafe.Pointer, event unsafe.Pointer) bool
}

type RulerView struct {
	id objc.ID
}

func RulerViewFrom(ptr unsafe.Pointer) RulerView {
	return RulerView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ RulerView) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _RulerViewClass) Alloc() RulerView {
	rv := objc.Send[RulerView](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _RulerViewClass) New() RulerView {
	rv := objc.Send[RulerView](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewRulerView creates and returns a new initialized instance.
func NewRulerView() RulerView {
	return RulerViewClass.New()
}

// Init initializes the instance.
func (r_ RulerView) Init() RulerView {
	rv := objc.Send[RulerView](r_.ID(), selInit)
	return rv
}
// Registers a new unit of measurement with the NSRulerView class, making it available to all instances of NSRulerView. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/registerUnit(withName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:)
func (rc _RulerViewClass) RegisterUnitWithNameAbbreviationUnitToPointsConversionFactorStepUpCycleStepDownCycle(unitName unsafe.Pointer, abbreviation unsafe.Pointer, conversionFactor float64, stepUpCycle unsafe.Pointer, stepDownCycle unsafe.Pointer)  {
	objc.Send[objc.ID](objc.ID(rc.Class), objc.RegisterName("registerUnitWithName:abbreviation:unitToPointsConversionFactor:stepUpCycle:stepDownCycle:"), unitName, abbreviation, conversionFactor, stepUpCycle, stepDownCycle)
}
// Adds   to the receiver, without consulting the client view for approval. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/addMarker(_:)
func (r_ RulerView) AddMarker(marker unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("addMarker:"), marker)
}
// Draws the receiver’s hash marks and labels in  , which is expressed in the receiver’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/drawHashMarksAndLabels(in:)
func (r_ RulerView) DrawHashMarksAndLabelsInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("drawHashMarksAndLabelsInRect:"), rect)
}
// Draws the receiver’s markers in  , which is expressed in the receiver’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/drawMarkers(in:)
func (r_ RulerView) DrawMarkersInRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("drawMarkersInRect:"), rect)
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/init(coder:)
func (r_ RulerView) InitWithCoder(coder unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("initWithCoder:"), coder)
	return rv
}
// Initializes a newly allocated NSRulerView to have   (  or  ) within  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/init(scrollView:orientation:)
func (r_ RulerView) InitWithScrollViewOrientation(scrollView unsafe.Pointer, orientation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("initWithScrollView:orientation:"), scrollView, orientation)
	return rv
}
// Forces recalculation of the hash mark spacing for the next time the receiver is displayed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/invalidateHashMarks()
func (r_ RulerView) InvalidateHashMarks() {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("invalidateHashMarks"))
}
// Draws temporary lines in the ruler area. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/moveRulerline(fromLocation:toLocation:)
func (r_ RulerView) MoveRulerlineFromLocationToLocation(oldLocation float64, newLocation float64) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("moveRulerlineFromLocation:toLocation:"), oldLocation, newLocation)
}
// Removes   from the receiver, without consulting the client view for approval. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/removeMarker(_:)
func (r_ RulerView) RemoveMarker(marker unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("removeMarker:"), marker)
}
// Tracks the mouse to add   based on the initial mouse-down or mouse-dragged event  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/trackMarker(_:withMouseEvent:)
func (r_ RulerView) TrackMarkerWithMouseEvent(marker unsafe.Pointer, event unsafe.Pointer) bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("trackMarker:withMouseEvent:"), marker, event)
	return rv
}
// The receiver’s accessory view to  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/accessoryView
func (r_ RulerView) AccessoryView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("accessoryView"))
	return rv
}
// SetAccessoryView sets the value of the accessoryView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/accessoryView
func (r_ RulerView) SetAccessoryView(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setAccessoryView:"), value)
}
// The location of the receiver’s baseline, in its own coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/baselineLocation
func (r_ RulerView) BaselineLocation() float64 {
	rv := objc.Send[float64](r_.ID(), objc.RegisterName("baselineLocation"))
	return rv
}
// The receiver’s client view, if it has one. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/clientView
func (r_ RulerView) ClientView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("clientView"))
	return rv
}
// SetClientView sets the value of the clientView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/clientView
func (r_ RulerView) SetClientView(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setClientView:"), value)
}
// A Boolean that indicates if the ruler view’s coordinate system is flipped. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/isFlipped
func (r_ RulerView) Flipped() bool {
	rv := objc.Send[bool](r_.ID(), objc.RegisterName("flipped"))
	return rv
}
// The receiver’s ruler markers to  , removing any existing ruler markers and not consulting with the client view about the new markers. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/markers
func (r_ RulerView) Markers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("markers"))
	return rv
}
// SetMarkers sets the value of the markers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/markers
func (r_ RulerView) SetMarkers(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setMarkers:"), value)
}
// The measurement units used by the ruler to  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/measurementUnits
func (r_ RulerView) MeasurementUnits() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("measurementUnits"))
	return rv
}
// SetMeasurementUnits sets the value of the measurementUnits property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/measurementUnits
func (r_ RulerView) SetMeasurementUnits(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setMeasurementUnits:"), value)
}
// The orientation of the receiver to  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/orientation-swift.property
func (r_ RulerView) Orientation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("orientation"))
	return rv
}
// SetOrientation sets the value of the orientation property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/orientation-swift.property
func (r_ RulerView) SetOrientation(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setOrientation:"), value)
}
// The distance to the zero hash mark from the bounds origin of the NSScrollView’s document view (not of the receiver’s client view), in the document view’s coordinate system. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/originOffset
func (r_ RulerView) OriginOffset() float64 {
	rv := objc.Send[float64](r_.ID(), objc.RegisterName("originOffset"))
	return rv
}
// SetOriginOffset sets the value of the originOffset property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/originOffset
func (r_ RulerView) SetOriginOffset(value float64) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setOriginOffset:"), value)
}
// The thickness needed for proper tiling of the receiver within an NSScrollView. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/requiredThickness
func (r_ RulerView) RequiredThickness() float64 {
	rv := objc.Send[float64](r_.ID(), objc.RegisterName("requiredThickness"))
	return rv
}
// The room available for the receiver’s accessory view to  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/reservedThicknessForAccessoryView
func (r_ RulerView) ReservedThicknessForAccessoryView() float64 {
	rv := objc.Send[float64](r_.ID(), objc.RegisterName("reservedThicknessForAccessoryView"))
	return rv
}
// SetReservedThicknessForAccessoryView sets the value of the reservedThicknessForAccessoryView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/reservedThicknessForAccessoryView
func (r_ RulerView) SetReservedThicknessForAccessoryView(value float64) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setReservedThicknessForAccessoryView:"), value)
}
// The room available for ruler markers to  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/reservedThicknessForMarkers
func (r_ RulerView) ReservedThicknessForMarkers() float64 {
	rv := objc.Send[float64](r_.ID(), objc.RegisterName("reservedThicknessForMarkers"))
	return rv
}
// SetReservedThicknessForMarkers sets the value of the reservedThicknessForMarkers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/reservedThicknessForMarkers
func (r_ RulerView) SetReservedThicknessForMarkers(value float64) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setReservedThicknessForMarkers:"), value)
}
// The thickness of the area where ruler hash marks and labels are drawn. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/ruleThickness
func (r_ RulerView) RuleThickness() float64 {
	rv := objc.Send[float64](r_.ID(), objc.RegisterName("ruleThickness"))
	return rv
}
// SetRuleThickness sets the value of the ruleThickness property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/ruleThickness
func (r_ RulerView) SetRuleThickness(value float64) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setRuleThickness:"), value)
}
// The NSScrollView that owns the receiver to  , without retaining it. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/scrollView
func (r_ RulerView) ScrollView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID(), objc.RegisterName("scrollView"))
	return rv
}
// SetScrollView sets the value of the scrollView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSRulerView/scrollView
func (r_ RulerView) SetScrollView(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID(), objc.RegisterName("setScrollView:"), value)
}
