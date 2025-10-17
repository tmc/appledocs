// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ScrollView] class.
var scrollViewClass = _ScrollViewClass{objc.GetClass("NSScrollView")}

type _ScrollViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrollView] class.
type IScrollView interface {
	IView
	AddFloatingSubviewForAxis(view unsafe.Pointer, axis unsafe.Pointer)
	FlashScrollers()
	MagnifyToFitRect(rect unsafe.Pointer)
	ReflectScrolledClipView(cView unsafe.Pointer)
	ScrollWheel(event unsafe.Pointer)
	SetMagnificationCenteredAtPoint(magnification float64, point unsafe.Pointer)
	Tile()
}

// A view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView

type ScrollView struct {
	View
}

// ScrollViewFrom constructs a [ScrollView] from an unsafe.Pointer.
//
// A view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view.
func ScrollViewFrom(ptr unsafe.Pointer) ScrollView {
	return ScrollView{
		View: ViewFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (sc _ScrollViewClass) Alloc() ScrollView {
	rv := objc.Send[ScrollView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (sc _ScrollViewClass) New() ScrollView {
	rv := objc.Send[ScrollView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrollView) Init() ScrollView {
	rv := objc.Send[ScrollView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrollView) Autorelease() ScrollView {
	rv := objc.Send[ScrollView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrollView creates a new ScrollView instance.
func NewScrollView() ScrollView {
	return scrollViewClass.New()
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/init(coder:)
func NewScrollViewWithCoder(coder unsafe.Pointer) ScrollView {
	instance := scrollViewClass.Alloc()
	rv := objc.Send[ScrollView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/init(frame:)
func NewScrollViewWithFrame(frameRect unsafe.Pointer) ScrollView {
	instance := scrollViewClass.Alloc()
	rv := objc.Send[ScrollView](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}


// Returns the content size calculated from the frame size and the specified specifications. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentSize(forFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (sc _ScrollViewClass) ContentSizeForFrameSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(fSize unsafe.Pointer, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ unsafe.Pointer, controlSize unsafe.Pointer, scrollerStyle unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("contentSizeForFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), fSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return rv
}
// Returns the content size calculated from the frame size and the specified specifications. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:
func (sc _ScrollViewClass) ContentSizeForFrameSizeHasHorizontalScrollerHasVerticalScrollerBorderType(fSize unsafe.Pointer, hFlag bool, vFlag bool, type_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:"), fSize, hFlag, vFlag, type_)
	return rv
}
// Returns the frame size of a scroll view that contains a content view with the specified size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/frameSize(forContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (sc _ScrollViewClass) FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize unsafe.Pointer, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ unsafe.Pointer, controlSize unsafe.Pointer, scrollerStyle unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("frameSizeForContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), cSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return rv
}
// Returns the frame size of an scroll view that contains a content view with the specified size. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:
func (sc _ScrollViewClass) FrameSizeForContentSizeHasHorizontalScrollerHasVerticalScrollerBorderType(cSize unsafe.Pointer, hFlag bool, vFlag bool, type_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:"), cSize, hFlag, vFlag, type_)
	return rv
}
// Adds a floating subview to the document view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/addFloatingSubview(_:for:)
func (s_ ScrollView) AddFloatingSubviewForAxis(view unsafe.Pointer, axis unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addFloatingSubview:forAxis:"), view, axis)
}
// Flash the overlay scroll bars. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/flashScrollers()
func (s_ ScrollView) FlashScrollers() {
	objc.Send[objc.ID](s_.ID, objc.Sel("flashScrollers"))
}
// Magnifies the content view proportionally such that the given rectangle fits centered in the scroll view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/magnify(toFit:)
func (s_ ScrollView) MagnifyToFitRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("magnifyToFitRect:"), rect)
}
// Adjusts the receiver’s scrollers to reflect the size and positioning of its content view. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/reflectScrolledClipView(_:)
func (s_ ScrollView) ReflectScrolledClipView(cView unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("reflectScrolledClipView:"), cView)
}
// Scrolls the receiver up or down, in response to the user moving the mouse’s scroll wheel specified by . [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollWheel(with:)
func (s_ ScrollView) ScrollWheel(event unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("scrollWheel:"), event)
}
// Magnify the content by the given amount and center the result on the given point. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/setMagnification(_:centeredAt:)
func (s_ ScrollView) SetMagnificationCenteredAtPoint(magnification float64, point unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}
// Lays out the components of the receiver: the content view, the scrollers, and the ruler views. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/tile()
func (s_ ScrollView) Tile() {
	objc.Send[objc.ID](s_.ID, objc.Sel("tile"))
}

