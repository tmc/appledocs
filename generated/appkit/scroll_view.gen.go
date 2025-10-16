
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ScrollView] class.
var ScrollViewClass _ScrollViewClass

func init() {
	ScrollViewClass = _ScrollViewClass{objc.GetClass("NSScrollView")}
}

type _ScrollViewClass struct {
	objc.Class
}

// An interface definition for the [ScrollView] class.
type IScrollView interface {
	ID() objc.ID
	AddFloatingSubviewForAxis(view unsafe.Pointer, axis unsafe.Pointer)
	FlashScrollers()
	InitWithCoder(coder unsafe.Pointer) unsafe.Pointer
	InitWithFrame(frameRect unsafe.Pointer) unsafe.Pointer
	MagnifyToFitRect(rect unsafe.Pointer)
	ReflectScrolledClipView(cView unsafe.Pointer)
	ScrollWheel(event unsafe.Pointer)
	SetMagnificationCenteredAtPoint(magnification float64, point unsafe.Pointer)
	Tile()
}

type ScrollView struct {
	id objc.ID
}

func ScrollViewFrom(ptr unsafe.Pointer) ScrollView {
	return ScrollView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ ScrollView) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _ScrollViewClass) Alloc() ScrollView {
	rv := objc.Send[ScrollView](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _ScrollViewClass) New() ScrollView {
	rv := objc.Send[ScrollView](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewScrollView creates and returns a new initialized instance.
func NewScrollView() ScrollView {
	return ScrollViewClass.New()
}

// Init initializes the instance.
func (s_ ScrollView) Init() ScrollView {
	rv := objc.Send[ScrollView](s_.ID(), selInit)
	return rv
}
// Returns the content size calculated from the frame size and the specified specifications. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/contentSize(forFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (sc _ScrollViewClass) ContentSizeForFrameSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(fSize unsafe.Pointer, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ unsafe.Pointer, controlSize unsafe.Pointer, scrollerStyle unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.Class), objc.RegisterName("contentSizeForFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), fSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return rv
}
// Returns the content size calculated from the frame size and the specified specifications. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:
func (sc _ScrollViewClass) ContentSizeForFrameSizeHasHorizontalScrollerHasVerticalScrollerBorderType(fSize unsafe.Pointer, hFlag bool, vFlag bool, type_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.Class), objc.RegisterName("contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:"), fSize, hFlag, vFlag, type_)
	return rv
}
// Returns the frame size of a scroll view that contains a content view with the specified size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/frameSize(forContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (sc _ScrollViewClass) FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize unsafe.Pointer, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ unsafe.Pointer, controlSize unsafe.Pointer, scrollerStyle unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.Class), objc.RegisterName("frameSizeForContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), cSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return rv
}
// Returns the frame size of an scroll view that contains a content view with the specified size. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:
func (sc _ScrollViewClass) FrameSizeForContentSizeHasHorizontalScrollerHasVerticalScrollerBorderType(cSize unsafe.Pointer, hFlag bool, vFlag bool, type_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.Class), objc.RegisterName("frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:"), cSize, hFlag, vFlag, type_)
	return rv
}
// Adds a floating subview to the document view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/addFloatingSubview(_:for:)
func (s_ ScrollView) AddFloatingSubviewForAxis(view unsafe.Pointer, axis unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("addFloatingSubview:forAxis:"), view, axis)
}
// Flash the overlay scroll bars. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/flashScrollers()
func (s_ ScrollView) FlashScrollers() {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("flashScrollers"))
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/init(coder:)
func (s_ ScrollView) InitWithCoder(coder unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("initWithCoder:"), coder)
	return rv
}
//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/init(frame:)
func (s_ ScrollView) InitWithFrame(frameRect unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("initWithFrame:"), frameRect)
	return rv
}
// Magnifies the content view proportionally such that the given rectangle fits centered in the scroll view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/magnify(toFit:)
func (s_ ScrollView) MagnifyToFitRect(rect unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("magnifyToFitRect:"), rect)
}
// Adjusts the receiver’s scrollers to reflect the size and positioning of its content view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/reflectScrolledClipView(_:)
func (s_ ScrollView) ReflectScrolledClipView(cView unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("reflectScrolledClipView:"), cView)
}
// Scrolls the receiver up or down, in response to the user moving the mouse’s scroll wheel specified by  . [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/scrollWheel(with:)
func (s_ ScrollView) ScrollWheel(event unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("scrollWheel:"), event)
}
// Magnify the content by the given amount and center the result on the given point. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/setMagnification(_:centeredAt:)
func (s_ ScrollView) SetMagnificationCenteredAtPoint(magnification float64, point unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setMagnification:centeredAtPoint:"), magnification, point)
}
// Lays out the components of the receiver: the content view, the scrollers, and the ruler views. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/tile()
func (s_ ScrollView) Tile() {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("tile"))
}
// Allows the user to magnify the scroll view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/allowsMagnification
func (s_ ScrollView) AllowsMagnification() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("allowsMagnification"))
	return rv
}
// SetAllowsMagnification sets the value of the allowsMagnification property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/allowsMagnification
func (s_ ScrollView) SetAllowsMagnification(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setAllowsMagnification:"), value)
}
// A Boolean that indicates whether the scroll view automatically hides its scroll bars when they are not needed. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/autohidesScrollers
func (s_ ScrollView) AutohidesScrollers() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("autohidesScrollers"))
	return rv
}
// SetAutohidesScrollers sets the value of the autohidesScrollers property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/autohidesScrollers
func (s_ ScrollView) SetAutohidesScrollers(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setAutohidesScrollers:"), value)
}
// A Boolean that indicates whether the scroll view automatically adjusts its content insets. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/automaticallyAdjustsContentInsets
func (s_ ScrollView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("automaticallyAdjustsContentInsets"))
	return rv
}
// SetAutomaticallyAdjustsContentInsets sets the value of the automaticallyAdjustsContentInsets property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/automaticallyAdjustsContentInsets
func (s_ ScrollView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setAutomaticallyAdjustsContentInsets:"), value)
}
// The color of the content view’s background. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/backgroundColor
func (s_ ScrollView) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("backgroundColor"))
	return rv
}
// SetBackgroundColor sets the value of the backgroundColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/backgroundColor
func (s_ ScrollView) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setBackgroundColor:"), value)
}
// A value that specifies the appearance of the scroll view’s border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/borderType
func (s_ ScrollView) BorderType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("borderType"))
	return rv
}
// SetBorderType sets the value of the borderType property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/borderType
func (s_ ScrollView) SetBorderType(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setBorderType:"), value)
}
// The distance that the scroll view’s subviews are inset from the enclosing scroll view during tiling. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/contentInsets
func (s_ ScrollView) ContentInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("contentInsets"))
	return rv
}
// SetContentInsets sets the value of the contentInsets property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/contentInsets
func (s_ ScrollView) SetContentInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setContentInsets:"), value)
}
// The size of the scroll view’s content view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/contentSize
func (s_ ScrollView) ContentSize() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("contentSize"))
	return rv
}
// The scroll view’s content view, the view that clips the document view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/contentView
func (s_ ScrollView) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("contentView"))
	return rv
}
// SetContentView sets the value of the contentView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/contentView
func (s_ ScrollView) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setContentView:"), value)
}
// The content view’s document cursor. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/documentCursor
func (s_ ScrollView) DocumentCursor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("documentCursor"))
	return rv
}
// SetDocumentCursor sets the value of the documentCursor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/documentCursor
func (s_ ScrollView) SetDocumentCursor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDocumentCursor:"), value)
}
// The view the scroll view scrolls within its content view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/documentView
func (s_ ScrollView) DocumentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("documentView"))
	return rv
}
// SetDocumentView sets the value of the documentView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/documentView
func (s_ ScrollView) SetDocumentView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDocumentView:"), value)
}
// The portion of the document view, in its own coordinate system, visible through the scroll view’s content view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/documentVisibleRect
func (s_ ScrollView) DocumentVisibleRect() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("documentVisibleRect"))
	return rv
}
// A Boolean that indicates whether the scroll view draws its background. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/drawsBackground
func (s_ ScrollView) DrawsBackground() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("drawsBackground"))
	return rv
}
// SetDrawsBackground sets the value of the drawsBackground property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/drawsBackground
func (s_ ScrollView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setDrawsBackground:"), value)
}
// The position of the find bar. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/findBarPosition-swift.property
func (s_ ScrollView) FindBarPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("findBarPosition"))
	return rv
}
// SetFindBarPosition sets the value of the findBarPosition property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/findBarPosition-swift.property
func (s_ ScrollView) SetFindBarPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setFindBarPosition:"), value)
}
// A Boolean that indicates whether the scroll view keeps a horizontal ruler object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/hasHorizontalRuler
func (s_ ScrollView) HasHorizontalRuler() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("hasHorizontalRuler"))
	return rv
}
// SetHasHorizontalRuler sets the value of the hasHorizontalRuler property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/hasHorizontalRuler
func (s_ ScrollView) SetHasHorizontalRuler(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHasHorizontalRuler:"), value)
}
// A Boolean that indicates whether the scroll view has a horizontal scroller. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/hasHorizontalScroller
func (s_ ScrollView) HasHorizontalScroller() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("hasHorizontalScroller"))
	return rv
}
// SetHasHorizontalScroller sets the value of the hasHorizontalScroller property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/hasHorizontalScroller
func (s_ ScrollView) SetHasHorizontalScroller(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHasHorizontalScroller:"), value)
}
// A Boolean that indicates whether the scroll view keeps a vertical ruler object. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/hasVerticalRuler
func (s_ ScrollView) HasVerticalRuler() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("hasVerticalRuler"))
	return rv
}
// SetHasVerticalRuler sets the value of the hasVerticalRuler property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/hasVerticalRuler
func (s_ ScrollView) SetHasVerticalRuler(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHasVerticalRuler:"), value)
}
// A Boolean that indicates whether the scroll view has a vertical scroller. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/hasVerticalScroller
func (s_ ScrollView) HasVerticalScroller() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("hasVerticalScroller"))
	return rv
}
// SetHasVerticalScroller sets the value of the hasVerticalScroller property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/hasVerticalScroller
func (s_ ScrollView) SetHasVerticalScroller(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHasVerticalScroller:"), value)
}
// The scroll view’s horizontal line by line scroll amount. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/horizontalLineScroll
func (s_ ScrollView) HorizontalLineScroll() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("horizontalLineScroll"))
	return rv
}
// SetHorizontalLineScroll sets the value of the horizontalLineScroll property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/horizontalLineScroll
func (s_ ScrollView) SetHorizontalLineScroll(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHorizontalLineScroll:"), value)
}
// The amount of the document view kept visible when scrolling horizontally page by page. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/horizontalPageScroll
func (s_ ScrollView) HorizontalPageScroll() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("horizontalPageScroll"))
	return rv
}
// SetHorizontalPageScroll sets the value of the horizontalPageScroll property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/horizontalPageScroll
func (s_ ScrollView) SetHorizontalPageScroll(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHorizontalPageScroll:"), value)
}
// The scroll view’s horizontal ruler view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/horizontalRulerView
func (s_ ScrollView) HorizontalRulerView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("horizontalRulerView"))
	return rv
}
// SetHorizontalRulerView sets the value of the horizontalRulerView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/horizontalRulerView
func (s_ ScrollView) SetHorizontalRulerView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHorizontalRulerView:"), value)
}
// The scroll view’s horizontal scrolling elasticity mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/horizontalScrollElasticity
func (s_ ScrollView) HorizontalScrollElasticity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("horizontalScrollElasticity"))
	return rv
}
// SetHorizontalScrollElasticity sets the value of the horizontalScrollElasticity property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/horizontalScrollElasticity
func (s_ ScrollView) SetHorizontalScrollElasticity(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHorizontalScrollElasticity:"), value)
}
// The scroll view’s horizontal scroller. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/horizontalScroller
func (s_ ScrollView) HorizontalScroller() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("horizontalScroller"))
	return rv
}
// SetHorizontalScroller sets the value of the horizontalScroller property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/horizontalScroller
func (s_ ScrollView) SetHorizontalScroller(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setHorizontalScroller:"), value)
}
// The scroll view’s line by line scroll amount. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/lineScroll
func (s_ ScrollView) LineScroll() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("lineScroll"))
	return rv
}
// SetLineScroll sets the value of the lineScroll property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/lineScroll
func (s_ ScrollView) SetLineScroll(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setLineScroll:"), value)
}
// The amount by which the content is currently scaled. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/magnification
func (s_ ScrollView) Magnification() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("magnification"))
	return rv
}
// SetMagnification sets the value of the magnification property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/magnification
func (s_ ScrollView) SetMagnification(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setMagnification:"), value)
}
// The maximum value to which the content can be magnified. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/maxMagnification
func (s_ ScrollView) MaxMagnification() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("maxMagnification"))
	return rv
}
// SetMaxMagnification sets the value of the maxMagnification property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/maxMagnification
func (s_ ScrollView) SetMaxMagnification(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setMaxMagnification:"), value)
}
// The minimum value to which the content can be magnified. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/minMagnification
func (s_ ScrollView) MinMagnification() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("minMagnification"))
	return rv
}
// SetMinMagnification sets the value of the minMagnification property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/minMagnification
func (s_ ScrollView) SetMinMagnification(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setMinMagnification:"), value)
}
// The amount of the document view kept visible when scrolling page by page. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/pageScroll
func (s_ ScrollView) PageScroll() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("pageScroll"))
	return rv
}
// SetPageScroll sets the value of the pageScroll property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/pageScroll
func (s_ ScrollView) SetPageScroll(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setPageScroll:"), value)
}
// A Boolean that indicates whether the scroll view displays its rulers. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/rulersVisible
func (s_ ScrollView) RulersVisible() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("rulersVisible"))
	return rv
}
// SetRulersVisible sets the value of the rulersVisible property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/rulersVisible
func (s_ ScrollView) SetRulersVisible(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setRulersVisible:"), value)
}
// The distance the scrollers are inset from the edge of the scroll view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/scrollerInsets
func (s_ ScrollView) ScrollerInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("scrollerInsets"))
	return rv
}
// SetScrollerInsets sets the value of the scrollerInsets property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/scrollerInsets
func (s_ ScrollView) SetScrollerInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setScrollerInsets:"), value)
}
// The knob style of scroll views that use the overlay scroller style. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/scrollerKnobStyle
func (s_ ScrollView) ScrollerKnobStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("scrollerKnobStyle"))
	return rv
}
// SetScrollerKnobStyle sets the value of the scrollerKnobStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/scrollerKnobStyle
func (s_ ScrollView) SetScrollerKnobStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setScrollerKnobStyle:"), value)
}
// The scroller style used by the scroll view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/scrollerStyle
func (s_ ScrollView) ScrollerStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("scrollerStyle"))
	return rv
}
// SetScrollerStyle sets the value of the scrollerStyle property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/scrollerStyle
func (s_ ScrollView) SetScrollerStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setScrollerStyle:"), value)
}
// A Boolean that indicates whether the scroll view redraws its document view while scrolling continuously. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/scrollsDynamically
func (s_ ScrollView) ScrollsDynamically() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("scrollsDynamically"))
	return rv
}
// SetScrollsDynamically sets the value of the scrollsDynamically property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/scrollsDynamically
func (s_ ScrollView) SetScrollsDynamically(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setScrollsDynamically:"), value)
}
// A Boolean that indicates whether the scroll view uses a predominant scrolling axis for content. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/usesPredominantAxisScrolling
func (s_ ScrollView) UsesPredominantAxisScrolling() bool {
	rv := objc.Send[bool](s_.ID(), objc.RegisterName("usesPredominantAxisScrolling"))
	return rv
}
// SetUsesPredominantAxisScrolling sets the value of the usesPredominantAxisScrolling property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/usesPredominantAxisScrolling
func (s_ ScrollView) SetUsesPredominantAxisScrolling(value bool) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setUsesPredominantAxisScrolling:"), value)
}
// The scroll view’s vertical line by line scroll amount. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/verticalLineScroll
func (s_ ScrollView) VerticalLineScroll() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("verticalLineScroll"))
	return rv
}
// SetVerticalLineScroll sets the value of the verticalLineScroll property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/verticalLineScroll
func (s_ ScrollView) SetVerticalLineScroll(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setVerticalLineScroll:"), value)
}
// The amount of the document view kept visible when scrolling vertically page by page. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/verticalPageScroll
func (s_ ScrollView) VerticalPageScroll() float64 {
	rv := objc.Send[float64](s_.ID(), objc.RegisterName("verticalPageScroll"))
	return rv
}
// SetVerticalPageScroll sets the value of the verticalPageScroll property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/verticalPageScroll
func (s_ ScrollView) SetVerticalPageScroll(value float64) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setVerticalPageScroll:"), value)
}
// The scroll view’s vertical ruler view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/verticalRulerView
func (s_ ScrollView) VerticalRulerView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("verticalRulerView"))
	return rv
}
// SetVerticalRulerView sets the value of the verticalRulerView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/verticalRulerView
func (s_ ScrollView) SetVerticalRulerView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setVerticalRulerView:"), value)
}
// The scroll view’s vertical scrolling elasticity mode. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/verticalScrollElasticity
func (s_ ScrollView) VerticalScrollElasticity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("verticalScrollElasticity"))
	return rv
}
// SetVerticalScrollElasticity sets the value of the verticalScrollElasticity property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/verticalScrollElasticity
func (s_ ScrollView) SetVerticalScrollElasticity(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setVerticalScrollElasticity:"), value)
}
// The scroll view’s vertical scroller. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/verticalScroller
func (s_ ScrollView) VerticalScroller() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID(), objc.RegisterName("verticalScroller"))
	return rv
}
// SetVerticalScroller sets the value of the verticalScroller property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSScrollView/verticalScroller
func (s_ ScrollView) SetVerticalScroller(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID(), objc.RegisterName("setVerticalScroller:"), value)
}
