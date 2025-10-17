// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrollView] class.
var ScrollViewClass objc.Class

func init() {
	ScrollViewClass = objc.GetClass("NSScrollView")
}

type ScrollView struct {
	objc.ID
}

func ScrollViewFrom(ptr unsafe.Pointer) ScrollView {
	return ScrollView{
		ID: objc.ID(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc ScrollView) Alloc() ScrollView {
	ret := objc.ID(ScrollViewClass).Send(objc.RegisterName("alloc"))
	return ScrollView{ret}
}

// Init initializes the instance.
func (s_ ScrollView) Init() ScrollView {
	ret := s_.ID.Send(objc.RegisterName("init"))
	return ScrollView{ret}
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/init(coder:)
func NewScrollViewWithCoder(coder unsafe.Pointer) ScrollView {
	instance := ScrollView{}.Alloc()
	sel := objc.RegisterName("initWithCoder:")
	ret := instance.ID.Send(sel, coder)
	instance = ScrollView{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/init(frame:)
func NewScrollViewWithFrame(frameRect foundation.Rect) ScrollView {
	instance := ScrollView{}.Alloc()
	sel := objc.RegisterName("initWithFrame:")
	ret := instance.ID.Send(sel, frameRect)
	instance = ScrollView{ret}
	instance.ID = instance.ID.Send(objc.RegisterName("autorelease"))
	return instance
}


// Returns the content size calculated from the frame size and the specified specifications. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/contentSize(forFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (sc ScrollView) ContentSizeForFrameSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(fSize foundation.Size, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ unsafe.Pointer, controlSize unsafe.Pointer, scrollerStyle unsafe.Pointer) foundation.Size {
	sel := objc.RegisterName("contentSizeForFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:")
	ret := objc.ID(ScrollViewClass).Send(sel, fSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return foundation.Size(ret)
}
// Returns the content size calculated from the frame size and the specified specifications. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:
func (sc ScrollView) ContentSizeForFrameSizeHasHorizontalScrollerHasVerticalScrollerBorderType(fSize foundation.Size, hFlag bool, vFlag bool, type_ unsafe.Pointer) foundation.Size {
	sel := objc.RegisterName("contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:")
	ret := objc.ID(ScrollViewClass).Send(sel, fSize, hFlag, vFlag, type_)
	return foundation.Size(ret)
}
// Returns the frame size of a scroll view that contains a content view with the specified size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/frameSize(forContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (sc ScrollView) FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize foundation.Size, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ unsafe.Pointer, controlSize unsafe.Pointer, scrollerStyle unsafe.Pointer) foundation.Size {
	sel := objc.RegisterName("frameSizeForContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:")
	ret := objc.ID(ScrollViewClass).Send(sel, cSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return foundation.Size(ret)
}
// Returns the frame size of an scroll view that contains a content view with the specified size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:
func (sc ScrollView) FrameSizeForContentSizeHasHorizontalScrollerHasVerticalScrollerBorderType(cSize foundation.Size, hFlag bool, vFlag bool, type_ unsafe.Pointer) foundation.Size {
	sel := objc.RegisterName("frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:")
	ret := objc.ID(ScrollViewClass).Send(sel, cSize, hFlag, vFlag, type_)
	return foundation.Size(ret)
}
// Adds a floating subview to the document view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/addFloatingSubview(_:for:)
func (s_ ScrollView) AddFloatingSubviewForAxis(view unsafe.Pointer, axis unsafe.Pointer) {
	sel := objc.RegisterName("addFloatingSubview:forAxis:")
	s_.ID.Send(sel, view, axis)
}
// Flash the overlay scroll bars. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/flashScrollers()
func (s_ ScrollView) FlashScrollers() {
	sel := objc.RegisterName("flashScrollers")
	s_.ID.Send(sel)
}
// Magnifies the content view proportionally such that the given rectangle fits centered in the scroll view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/magnify(toFit:)
func (s_ ScrollView) MagnifyToFitRect(rect foundation.Rect) {
	sel := objc.RegisterName("magnifyToFitRect:")
	s_.ID.Send(sel, rect)
}
// Adjusts the receiver’s scrollers to reflect the size and positioning of its content view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/reflectScrolledClipView(_:)
func (s_ ScrollView) ReflectScrolledClipView(cView unsafe.Pointer) {
	sel := objc.RegisterName("reflectScrolledClipView:")
	s_.ID.Send(sel, cView)
}
// Scrolls the receiver up or down, in response to the user moving the mouse’s scroll wheel specified by  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/scrollWheel(with:)
func (s_ ScrollView) ScrollWheel(event unsafe.Pointer) {
	sel := objc.RegisterName("scrollWheel:")
	s_.ID.Send(sel, event)
}
// Magnify the content by the given amount and center the result on the given point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/setMagnification(_:centeredAt:)
func (s_ ScrollView) SetMagnificationCenteredAtPoint(magnification float64, point foundation.Point) {
	sel := objc.RegisterName("setMagnification:centeredAtPoint:")
	s_.ID.Send(sel, magnification, point)
}
// Lays out the components of the receiver: the content view, the scrollers, and the ruler views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/tile()
func (s_ ScrollView) Tile() {
	sel := objc.RegisterName("tile")
	s_.ID.Send(sel)
}

