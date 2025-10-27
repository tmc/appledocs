// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)





// The class instance for the [Scroller] class.
var (
	ScrollerClass     _ScrollerClass
	ScrollerClassOnce sync.Once
)

func getScrollerClass() _ScrollerClass {
	ScrollerClassOnce.Do(func() {
		ScrollerClass = _ScrollerClass{objc.GetClass("NSScroller")}
	})
	return ScrollerClass
}

type _ScrollerClass struct {
	class objc.Class
}





// An interface definition for the [Scroller] class.
type IScroller interface {
	IControl
	

	// properties:
	ArrowsPosition() ScrollArrowPosition
	SetArrowsPosition(value ScrollArrowPosition)
	ControlSize() ControlSize
	SetControlSize(value ControlSize)
	ControlTint() ControlTint
	SetControlTint(value ControlTint)
	HitPart() ScrollerPart
	KnobProportion() float64
	SetKnobProportion(value float64)
	KnobStyle() ScrollerKnobStyle
	SetKnobStyle(value ScrollerKnobStyle)
	ScrollerStyle() ScrollerStyle
	SetScrollerStyle(value ScrollerStyle)
	UsableParts() UsableScrollerParts


	

	// methods:
	CheckSpaceForParts()
	DrawKnob()
	DrawKnobSlotInRectHighlight(slotRect corefoundation.CGRect, flag bool)
	RectForPart(partCode ScrollerPart) corefoundation.CGRect
	TestPart(point corefoundation.CGPoint) ScrollerPart
	TrackKnob(event IEvent)


}





// Alloc allocates a new instance without initialization.
func (sc _ScrollerClass) Alloc() Scroller {
	rv := objc.Send[Scroller](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ScrollerClass) New() Scroller {
	rv := objc.Send[Scroller](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Scroller) Init() Scroller {
	rv := objc.Send[Scroller](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Scroller) Autorelease() Scroller {
	rv := objc.Send[Scroller](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScroller creates a new Scroller instance.
func NewScroller() Scroller {
	return getScrollerClass().New()
}





// An object that controls scrolling of a document view within a scroll view or other type of container view.
//
// A scroller displays a slot containing a knob that the user can drag directly to the desired location. The knob indicates both the position within the document view and—by varying in size within the slot—the amount visible relative to the size of the document view. Typically, you don’t need to program with scrollers; instead, you configure them with an object in a . Don’t use an scroller when a slider would be more appropriate. An object represents a range of values for something in the application and lets the user choose a setting. A scroller represents the relative position of the visible portion of a view and lets the user choose which portion to view.


// An object that controls scrolling of a document view within a scroll view or other type of container view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller
type Scroller struct {
	Control
}

// ScrollerFrom constructs a [Scroller] from an unsafe.Pointer.
//
// An object that controls scrolling of a document view within a scroll view or other type of container view.
func ScrollerFrom(ptr unsafe.Pointer) Scroller {
	return Scroller{
		Control: ControlFrom(ptr),
	}
}










// Returns the width for scrollers of the receiving class, assuming a control size , and a scroller style of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/scrollerWidth
func (sc _ScrollerClass) ScrollerWidth() float64 {
	rv := objc.Send[float64](objc.ID(sc.class), objc.Sel("scrollerWidth"))
	return rv
}


// Returns the width of the scroller based on and assuming a scroller style of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/scrollerWidthForControlSize:
func (sc _ScrollerClass) ScrollerWidthForControlSize(controlSize ControlSize) float64 {
	rv := objc.Send[float64](objc.ID(sc.class), objc.Sel("scrollerWidthForControlSize:"), controlSize)
	return rv
}


// Returns the width for scrollers of the receiving class for a given control size and scroller style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/scrollerWidth(for:scrollerStyle:)
func (sc _ScrollerClass) ScrollerWidthForControlSizeScrollerStyle(controlSize ControlSize, scrollerStyle ScrollerStyle) float64 {
	rv := objc.Send[float64](objc.ID(sc.class), objc.Sel("scrollerWidthForControlSize:scrollerStyle:"), controlSize, scrollerStyle)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/isCompatibleWithOverlayScrollers
func (sc _ScrollerClass) CompatibleWithOverlayScrollers() bool {
	rv := objc.Send[bool](objc.ID(sc.class), objc.Sel("compatibleWithOverlayScrollers"))
	return rv
}

// Returns the style of scrollers that applications should use wherever possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/preferredScrollerStyle
func (sc _ScrollerClass) PreferredScrollerStyle() ScrollerStyle {
	rv := objc.Send[ScrollerStyle](objc.ID(sc.class), objc.Sel("preferredScrollerStyle"))
	return rv
}






// Checks to see if there is enough room in the receiver to display the knob and buttons.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/checkSpaceForParts()
func (s_ Scroller) CheckSpaceForParts() {
	objc.Send[objc.ID](s_.ID, objc.Sel("checkSpaceForParts"))
}


// Draws the knob.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/drawKnob()
func (s_ Scroller) DrawKnob() {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawKnob"))
}


// Draws the portion of the scroller’s track, possibly including the line increment and decrement arrow buttons, that falls in the given rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/drawKnobSlot(in:highlight:)
func (s_ Scroller) DrawKnobSlotInRectHighlight(slotRect corefoundation.CGRect, flag bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("drawKnobSlotInRect:highlight:"), slotRect, flag)
}


// Returns the rectangle occupied by , which for this method is interpreted literally rather than as an indicator of scrolling direction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/rect(for:)
func (s_ Scroller) RectForPart(partCode ScrollerPart) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s_.ID, objc.Sel("rectForPart:"), partCode)
	return rv
}


// Returns the part that would be hit by a mouse-down event at (expressed in the window’s coordinate system).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/testPart(_:)
func (s_ Scroller) TestPart(point corefoundation.CGPoint) ScrollerPart {
	rv := objc.Send[ScrollerPart](s_.ID, objc.Sel("testPart:"), point)
	return rv
}


// Tracks the knob and sends action messages to the receiver’s target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/trackKnob(with:)
func (s_ Scroller) TrackKnob(event IEvent) {
	objc.Send[objc.ID](s_.ID, objc.Sel("trackKnob:"), event)
}







// The location of the scroll buttons within the scroller, as described in .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/arrowsPosition
func (s_ Scroller) ArrowsPosition() ScrollArrowPosition {
	rv := objc.Send[ScrollArrowPosition](s_.ID, objc.Sel("arrowsPosition"))
	return rv
}


// The location of the scroll buttons within the scroller, as described in .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/arrowsPosition
func (s_ Scroller) SetArrowsPosition(value ScrollArrowPosition) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setArrowsPosition:"), value)
}


// The size of the scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/controlSize
func (s_ Scroller) ControlSize() ControlSize {
	rv := objc.Send[ControlSize](s_.ID, objc.Sel("controlSize"))
	return rv
}


// The size of the scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/controlSize
func (s_ Scroller) SetControlSize(value ControlSize) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setControlSize:"), value)
}


// The scroller’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/controlTint
func (s_ Scroller) ControlTint() ControlTint {
	rv := objc.Send[ControlTint](s_.ID, objc.Sel("controlTint"))
	return rv
}


// The scroller’s control tint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/controlTint
func (s_ Scroller) SetControlTint(value ControlTint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setControlTint:"), value)
}


// A part code indicating the manner in which the scrolling should be performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/hitPart
func (s_ Scroller) HitPart() ScrollerPart {
	rv := objc.Send[ScrollerPart](s_.ID, objc.Sel("hitPart"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/isCompatibleWithOverlayScrollers
func (s_ Scroller) CompatibleWithOverlayScrollers() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("compatibleWithOverlayScrollers"))
	return rv
}


// The proportion of the knob slot that the knob should fill.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/knobProportion
func (s_ Scroller) KnobProportion() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("knobProportion"))
	return rv
}


// The proportion of the knob slot that the knob should fill.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/knobProportion
func (s_ Scroller) SetKnobProportion(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKnobProportion:"), value)
}


// The scroller’s knob style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/knobStyle-swift.property
func (s_ Scroller) KnobStyle() ScrollerKnobStyle {
	rv := objc.Send[ScrollerKnobStyle](s_.ID, objc.Sel("knobStyle"))
	return rv
}


// The scroller’s knob style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/knobStyle-swift.property
func (s_ Scroller) SetKnobStyle(value ScrollerKnobStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKnobStyle:"), value)
}


// Returns the style of scrollers that applications should use wherever possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/preferredScrollerStyle
func (s_ Scroller) PreferredScrollerStyle() ScrollerStyle {
	rv := objc.Send[ScrollerStyle](s_.ID, objc.Sel("preferredScrollerStyle"))
	return rv
}


// The scroller style for this scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/scrollerStyle
func (s_ Scroller) ScrollerStyle() ScrollerStyle {
	rv := objc.Send[ScrollerStyle](s_.ID, objc.Sel("scrollerStyle"))
	return rv
}


// The scroller style for this scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/scrollerStyle
func (s_ Scroller) SetScrollerStyle(value ScrollerStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollerStyle:"), value)
}


// A value that indicates which parts of the receiver are displayed and usable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/usableParts-swift.property
func (s_ Scroller) UsableParts() UsableScrollerParts {
	rv := objc.Send[UsableScrollerParts](s_.ID, objc.Sel("usableParts"))
	return rv
}








