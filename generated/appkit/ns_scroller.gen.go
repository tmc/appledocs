// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that controls scrolling of a document view within a scroll view or other type of container view.
//
// A scroller displays a slot containing a knob that the user can drag directly to the desired location. The knob indicates both the position within the document view and—by varying in size within the slot—the amount visible relative to the size of the document view. Typically, you don’t need to program with scrollers; instead, you configure them with an object in a . Don’t use an scroller when a slider would be more appropriate. An object represents a range of values for something in the application and lets the user choose a setting. A scroller represents the relative position of the visible portion of a view and lets the user choose which portion to view.
//
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

// Alloc allocates a new instance without initialization.
func (sc _ScrollerClass) Alloc() Scroller {
	rv := objc.Send[Scroller](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns the style of scrollers that applications should use wherever possible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/preferredScrollerStyle
func (sc _ScrollerClass) PreferredScrollerStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("preferredScrollerStyle"))
	return rv
}
// Returns the style of scrollers that applications should use wherever possible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/preferredScrollerStyle
func (s_ Scroller) PreferredScrollerStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("preferredScrollerStyle"))
	return rv
}

// The location of the scroll buttons within the scroller, as described in
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/arrowsposition
func (s_ Scroller) ArrowsPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("arrowsPosition"))
	return rv
}


// SetArrowsPosition sets the value of the arrowsPosition property.
// The location of the scroll buttons within the scroller, as described in

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/arrowsposition
func (s_ Scroller) SetArrowsPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setArrowsPosition:"), value)
}

// The size of the scroller.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/controlsize
func (s_ Scroller) ControlSize() ControlSize {
	rv := objc.Send[ControlSize](s_.ID, objc.Sel("controlSize"))
	return rv
}


// SetControlSize sets the value of the controlSize property.
// The size of the scroller.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/controlsize
func (s_ Scroller) SetControlSize(value IControlSize) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setControlSize:"), value)
}

// The scroller’s control tint.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/controltint
func (s_ Scroller) ControlTint() ControlTint {
	rv := objc.Send[ControlTint](s_.ID, objc.Sel("controlTint"))
	return rv
}


// SetControlTint sets the value of the controlTint property.
// The scroller’s control tint.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/controltint
func (s_ Scroller) SetControlTint(value IControlTint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setControlTint:"), value)
}

// A part code indicating the manner in which the scrolling should be performed.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/hitpart
func (s_ Scroller) HitPart() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("hitPart"))
	return rv
}


// SetHitPart sets the value of the hitPart property.
// A part code indicating the manner in which the scrolling should be performed.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/hitpart
func (s_ Scroller) SetHitPart(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHitPart:"), value)
}

// The proportion of the knob slot that the knob should fill.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/knobproportion
func (s_ Scroller) KnobProportion() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("knobProportion"))
	return rv
}


// SetKnobProportion sets the value of the knobProportion property.
// The proportion of the knob slot that the knob should fill.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/knobproportion
func (s_ Scroller) SetKnobProportion(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKnobProportion:"), value)
}

// The scroller’s knob style.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/knobstyle-swift.property
func (s_ Scroller) KnobStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("knobStyle"))
	return rv
}


// SetKnobStyle sets the value of the knobStyle property.
// The scroller’s knob style.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/knobstyle-swift.property
func (s_ Scroller) SetKnobStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKnobStyle:"), value)
}

// The scroller style for this scroller.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/scrollerstyle
func (s_ Scroller) ScrollerStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("scrollerStyle"))
	return rv
}


// SetScrollerStyle sets the value of the scrollerStyle property.
// The scroller style for this scroller.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/scrollerstyle
func (s_ Scroller) SetScrollerStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollerStyle:"), value)
}

// A value that indicates which parts of the receiver are displayed and usable.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/usableparts-swift.property
func (s_ Scroller) UsableParts() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("usableParts"))
	return rv
}


// SetUsableParts sets the value of the usableParts property.
// A value that indicates which parts of the receiver are displayed and usable.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsscroller/usableparts-swift.property
func (s_ Scroller) SetUsableParts(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUsableParts:"), value)
}



