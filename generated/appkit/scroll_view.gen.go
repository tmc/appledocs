// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [ScrollView] class.
var (
	scrollViewClass     _ScrollViewClass
	scrollViewClassOnce sync.Once
)

func getScrollViewClass() _ScrollViewClass {
	scrollViewClassOnce.Do(func() {
		scrollViewClass = _ScrollViewClass{objc.GetClass("NSScrollView")}
	})
	return scrollViewClass
}

type _ScrollViewClass struct {
	class objc.Class
}

// An interface definition for the [ScrollView] class.
type IScrollView interface {
	IView
	AddFloatingSubviewForAxis(view unsafe.Pointer, axis unsafe.Pointer)
	FlashScrollers()
	MagnifyToFitRect(rect coregraphics.CGRect)
	ReflectScrolledClipView(cView unsafe.Pointer)
	ScrollWheel(event unsafe.Pointer)
	SetMagnificationCenteredAtPoint(magnification float64, point coregraphics.CGPoint)
	Tile()
}

// A view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view.
//
// The class is the central coordinator for AppKit’s scrolling machinery, which is composed of this class, and the and classes. When using an object within a scroll view (the usual configuration), you should issue messages that control background drawing state to the scroll view directly, rather than messaging the clip view.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getScrollViewClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/init(coder:)
func NewScrollViewWithCoder(coder unsafe.Pointer) ScrollView {
	instance := getScrollViewClass().Alloc()
	rv := objc.Send[ScrollView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/init(frame:)
func NewScrollViewWithFrame(frameRect coregraphics.CGRect) ScrollView {
	instance := getScrollViewClass().Alloc()
	rv := objc.Send[ScrollView](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}


// Returns the content size calculated from the frame size and the specified specifications.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentSize(forFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (sc _ScrollViewClass) ContentSizeForFrameSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(fSize coregraphics.CGSize, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ unsafe.Pointer, controlSize unsafe.Pointer, scrollerStyle unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](objc.ID(sc.class), objc.Sel("contentSizeForFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), fSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return rv
}

// Returns the content size calculated from the frame size and the specified specifications.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:
func (sc _ScrollViewClass) ContentSizeForFrameSizeHasHorizontalScrollerHasVerticalScrollerBorderType(fSize coregraphics.CGSize, hFlag bool, vFlag bool, type_ unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](objc.ID(sc.class), objc.Sel("contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:"), fSize, hFlag, vFlag, type_)
	return rv
}

// Returns the frame size of a scroll view that contains a content view with the specified size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/frameSize(forContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (sc _ScrollViewClass) FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize coregraphics.CGSize, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ unsafe.Pointer, controlSize unsafe.Pointer, scrollerStyle unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](objc.ID(sc.class), objc.Sel("frameSizeForContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), cSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return rv
}

// Returns the frame size of an scroll view that contains a content view with the specified size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:
func (sc _ScrollViewClass) FrameSizeForContentSizeHasHorizontalScrollerHasVerticalScrollerBorderType(cSize coregraphics.CGSize, hFlag bool, vFlag bool, type_ unsafe.Pointer) coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](objc.ID(sc.class), objc.Sel("frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:"), cSize, hFlag, vFlag, type_)
	return rv
}

// Adds a floating subview to the document view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/addFloatingSubview(_:for:)
func (s_ ScrollView) AddFloatingSubviewForAxis(view unsafe.Pointer, axis unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addFloatingSubview:forAxis:"), view, axis)
}

// Flash the overlay scroll bars.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/flashScrollers()
func (s_ ScrollView) FlashScrollers() {
	objc.Send[objc.ID](s_.ID, objc.Sel("flashScrollers"))
}

// Magnifies the content view proportionally such that the given rectangle fits centered in the scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/magnify(toFit:)
func (s_ ScrollView) MagnifyToFitRect(rect coregraphics.CGRect) {
	objc.Send[objc.ID](s_.ID, objc.Sel("magnifyToFitRect:"), rect)
}

// Adjusts the receiver’s scrollers to reflect the size and positioning of its content view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/reflectScrolledClipView(_:)
func (s_ ScrollView) ReflectScrolledClipView(cView unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("reflectScrolledClipView:"), cView)
}

// Scrolls the receiver up or down, in response to the user moving the mouse’s scroll wheel specified by .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollWheel(with:)
func (s_ ScrollView) ScrollWheel(event unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("scrollWheel:"), event)
}

// Magnify the content by the given amount and center the result on the given point.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/setMagnification(_:centeredAt:)
func (s_ ScrollView) SetMagnificationCenteredAtPoint(magnification float64, point coregraphics.CGPoint) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}

// Lays out the components of the receiver: the content view, the scrollers, and the ruler views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/tile()
func (s_ ScrollView) Tile() {
	objc.Send[objc.ID](s_.ID, objc.Sel("tile"))
}

// Allows the user to magnify the scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/allowsMagnification
func (s_ ScrollView) AllowsMagnification() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsMagnification"))
	return rv
}


// SetAllowsMagnification sets the value of the allowsMagnification property.
// Allows the user to magnify the scroll view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/allowsMagnification
func (s_ ScrollView) SetAllowsMagnification(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsMagnification:"), value)
}
// A Boolean that indicates whether the scroll view automatically hides its scroll bars when they are not needed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/autohidesScrollers
func (s_ ScrollView) AutohidesScrollers() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("autohidesScrollers"))
	return rv
}


// SetAutohidesScrollers sets the value of the autohidesScrollers property.
// A Boolean that indicates whether the scroll view automatically hides its scroll bars when they are not needed.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/autohidesScrollers
func (s_ ScrollView) SetAutohidesScrollers(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutohidesScrollers:"), value)
}
// A Boolean that indicates whether the scroll view automatically adjusts its content insets.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/automaticallyAdjustsContentInsets
func (s_ ScrollView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyAdjustsContentInsets"))
	return rv
}


// SetAutomaticallyAdjustsContentInsets sets the value of the automaticallyAdjustsContentInsets property.
// A Boolean that indicates whether the scroll view automatically adjusts its content insets.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/automaticallyAdjustsContentInsets
func (s_ ScrollView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticallyAdjustsContentInsets:"), value)
}
// The color of the content view’s background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/backgroundColor
func (s_ ScrollView) BackgroundColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The color of the content view’s background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/backgroundColor
func (s_ ScrollView) SetBackgroundColor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBackgroundColor:"), value)
}
// A value that specifies the appearance of the scroll view’s border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/borderType
func (s_ ScrollView) BorderType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("borderType"))
	return rv
}


// SetBorderType sets the value of the borderType property.
// A value that specifies the appearance of the scroll view’s border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/borderType
func (s_ ScrollView) SetBorderType(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBorderType:"), value)
}
// The distance that the scroll view’s subviews are inset from the enclosing scroll view during tiling.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentInsets
func (s_ ScrollView) ContentInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("contentInsets"))
	return rv
}


// SetContentInsets sets the value of the contentInsets property.
// The distance that the scroll view’s subviews are inset from the enclosing scroll view during tiling.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentInsets
func (s_ ScrollView) SetContentInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContentInsets:"), value)
}
// The size of the scroll view’s content view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentSize
func (s_ ScrollView) ContentSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](s_.ID, objc.Sel("contentSize"))
	return rv
}

// The scroll view’s content view, the view that clips the document view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentView
func (s_ ScrollView) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("contentView"))
	return rv
}


// SetContentView sets the value of the contentView property.
// The scroll view’s content view, the view that clips the document view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentView
func (s_ ScrollView) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContentView:"), value)
}
// The content view’s document cursor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/documentCursor
func (s_ ScrollView) DocumentCursor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("documentCursor"))
	return rv
}


// SetDocumentCursor sets the value of the documentCursor property.
// The content view’s document cursor.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/documentCursor
func (s_ ScrollView) SetDocumentCursor(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDocumentCursor:"), value)
}
// The view the scroll view scrolls within its content view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/documentView
func (s_ ScrollView) DocumentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("documentView"))
	return rv
}


// SetDocumentView sets the value of the documentView property.
// The view the scroll view scrolls within its content view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/documentView
func (s_ ScrollView) SetDocumentView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDocumentView:"), value)
}
// The portion of the document view, in its own coordinate system, visible through the scroll view’s content view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/documentVisibleRect
func (s_ ScrollView) DocumentVisibleRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](s_.ID, objc.Sel("documentVisibleRect"))
	return rv
}

// A Boolean that indicates whether the scroll view draws its background.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/drawsBackground
func (s_ ScrollView) DrawsBackground() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("drawsBackground"))
	return rv
}


// SetDrawsBackground sets the value of the drawsBackground property.
// A Boolean that indicates whether the scroll view draws its background.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/drawsBackground
func (s_ ScrollView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDrawsBackground:"), value)
}
// The position of the find bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/findBarPosition-swift.property
func (s_ ScrollView) FindBarPosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("findBarPosition"))
	return rv
}


// SetFindBarPosition sets the value of the findBarPosition property.
// The position of the find bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/findBarPosition-swift.property
func (s_ ScrollView) SetFindBarPosition(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFindBarPosition:"), value)
}
// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasHorizontalRuler
func (s_ ScrollView) HasHorizontalRuler() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasHorizontalRuler"))
	return rv
}


// SetHasHorizontalRuler sets the value of the hasHorizontalRuler property.
// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasHorizontalRuler
func (s_ ScrollView) SetHasHorizontalRuler(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHasHorizontalRuler:"), value)
}
// A Boolean that indicates whether the scroll view has a horizontal scroller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasHorizontalScroller
func (s_ ScrollView) HasHorizontalScroller() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasHorizontalScroller"))
	return rv
}


// SetHasHorizontalScroller sets the value of the hasHorizontalScroller property.
// A Boolean that indicates whether the scroll view has a horizontal scroller.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasHorizontalScroller
func (s_ ScrollView) SetHasHorizontalScroller(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHasHorizontalScroller:"), value)
}
// A Boolean that indicates whether the scroll view keeps a vertical ruler object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasVerticalRuler
func (s_ ScrollView) HasVerticalRuler() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasVerticalRuler"))
	return rv
}


// SetHasVerticalRuler sets the value of the hasVerticalRuler property.
// A Boolean that indicates whether the scroll view keeps a vertical ruler object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasVerticalRuler
func (s_ ScrollView) SetHasVerticalRuler(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHasVerticalRuler:"), value)
}
// A Boolean that indicates whether the scroll view has a vertical scroller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasVerticalScroller
func (s_ ScrollView) HasVerticalScroller() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasVerticalScroller"))
	return rv
}


// SetHasVerticalScroller sets the value of the hasVerticalScroller property.
// A Boolean that indicates whether the scroll view has a vertical scroller.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasVerticalScroller
func (s_ ScrollView) SetHasVerticalScroller(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHasVerticalScroller:"), value)
}
// The scroll view’s horizontal line by line scroll amount.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalLineScroll
func (s_ ScrollView) HorizontalLineScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("horizontalLineScroll"))
	return rv
}


// SetHorizontalLineScroll sets the value of the horizontalLineScroll property.
// The scroll view’s horizontal line by line scroll amount.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalLineScroll
func (s_ ScrollView) SetHorizontalLineScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalLineScroll:"), value)
}
// The amount of the document view kept visible when scrolling horizontally page by page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalPageScroll
func (s_ ScrollView) HorizontalPageScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("horizontalPageScroll"))
	return rv
}


// SetHorizontalPageScroll sets the value of the horizontalPageScroll property.
// The amount of the document view kept visible when scrolling horizontally page by page.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalPageScroll
func (s_ ScrollView) SetHorizontalPageScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalPageScroll:"), value)
}
// The scroll view’s horizontal ruler view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalRulerView
func (s_ ScrollView) HorizontalRulerView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("horizontalRulerView"))
	return rv
}


// SetHorizontalRulerView sets the value of the horizontalRulerView property.
// The scroll view’s horizontal ruler view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalRulerView
func (s_ ScrollView) SetHorizontalRulerView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalRulerView:"), value)
}
// The scroll view’s horizontal scrolling elasticity mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalScrollElasticity
func (s_ ScrollView) HorizontalScrollElasticity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("horizontalScrollElasticity"))
	return rv
}


// SetHorizontalScrollElasticity sets the value of the horizontalScrollElasticity property.
// The scroll view’s horizontal scrolling elasticity mode.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalScrollElasticity
func (s_ ScrollView) SetHorizontalScrollElasticity(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalScrollElasticity:"), value)
}
// The scroll view’s horizontal scroller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalScroller
func (s_ ScrollView) HorizontalScroller() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("horizontalScroller"))
	return rv
}


// SetHorizontalScroller sets the value of the horizontalScroller property.
// The scroll view’s horizontal scroller.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalScroller
func (s_ ScrollView) SetHorizontalScroller(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalScroller:"), value)
}
// The scroll view’s line by line scroll amount.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/lineScroll
func (s_ ScrollView) LineScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("lineScroll"))
	return rv
}


// SetLineScroll sets the value of the lineScroll property.
// The scroll view’s line by line scroll amount.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/lineScroll
func (s_ ScrollView) SetLineScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineScroll:"), value)
}
// The amount by which the content is currently scaled.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/magnification
func (s_ ScrollView) Magnification() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("magnification"))
	return rv
}


// SetMagnification sets the value of the magnification property.
// The amount by which the content is currently scaled.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/magnification
func (s_ ScrollView) SetMagnification(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMagnification:"), value)
}
// The maximum value to which the content can be magnified.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/maxMagnification
func (s_ ScrollView) MaxMagnification() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxMagnification"))
	return rv
}


// SetMaxMagnification sets the value of the maxMagnification property.
// The maximum value to which the content can be magnified.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/maxMagnification
func (s_ ScrollView) SetMaxMagnification(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxMagnification:"), value)
}
// The minimum value to which the content can be magnified.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/minMagnification
func (s_ ScrollView) MinMagnification() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minMagnification"))
	return rv
}


// SetMinMagnification sets the value of the minMagnification property.
// The minimum value to which the content can be magnified.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/minMagnification
func (s_ ScrollView) SetMinMagnification(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinMagnification:"), value)
}
// The amount of the document view kept visible when scrolling page by page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/pageScroll
func (s_ ScrollView) PageScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("pageScroll"))
	return rv
}


// SetPageScroll sets the value of the pageScroll property.
// The amount of the document view kept visible when scrolling page by page.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/pageScroll
func (s_ ScrollView) SetPageScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPageScroll:"), value)
}
// A Boolean that indicates whether the scroll view displays its rulers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/rulersVisible
func (s_ ScrollView) RulersVisible() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("rulersVisible"))
	return rv
}


// SetRulersVisible sets the value of the rulersVisible property.
// A Boolean that indicates whether the scroll view displays its rulers.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/rulersVisible
func (s_ ScrollView) SetRulersVisible(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRulersVisible:"), value)
}
// The distance the scrollers are inset from the edge of the scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerInsets
func (s_ ScrollView) ScrollerInsets() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("scrollerInsets"))
	return rv
}


// SetScrollerInsets sets the value of the scrollerInsets property.
// The distance the scrollers are inset from the edge of the scroll view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerInsets
func (s_ ScrollView) SetScrollerInsets(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollerInsets:"), value)
}
// The knob style of scroll views that use the overlay scroller style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerKnobStyle
func (s_ ScrollView) ScrollerKnobStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("scrollerKnobStyle"))
	return rv
}


// SetScrollerKnobStyle sets the value of the scrollerKnobStyle property.
// The knob style of scroll views that use the overlay scroller style.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerKnobStyle
func (s_ ScrollView) SetScrollerKnobStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollerKnobStyle:"), value)
}
// The scroller style used by the scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerStyle
func (s_ ScrollView) ScrollerStyle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("scrollerStyle"))
	return rv
}


// SetScrollerStyle sets the value of the scrollerStyle property.
// The scroller style used by the scroll view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerStyle
func (s_ ScrollView) SetScrollerStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollerStyle:"), value)
}
// A Boolean that indicates whether the scroll view redraws its document view while scrolling continuously.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollsDynamically
func (s_ ScrollView) ScrollsDynamically() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scrollsDynamically"))
	return rv
}


// SetScrollsDynamically sets the value of the scrollsDynamically property.
// A Boolean that indicates whether the scroll view redraws its document view while scrolling continuously.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollsDynamically
func (s_ ScrollView) SetScrollsDynamically(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollsDynamically:"), value)
}
// A Boolean that indicates whether the scroll view uses a predominant scrolling axis for content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/usesPredominantAxisScrolling
func (s_ ScrollView) UsesPredominantAxisScrolling() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("usesPredominantAxisScrolling"))
	return rv
}


// SetUsesPredominantAxisScrolling sets the value of the usesPredominantAxisScrolling property.
// A Boolean that indicates whether the scroll view uses a predominant scrolling axis for content.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/usesPredominantAxisScrolling
func (s_ ScrollView) SetUsesPredominantAxisScrolling(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUsesPredominantAxisScrolling:"), value)
}
// The scroll view’s vertical line by line scroll amount.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalLineScroll
func (s_ ScrollView) VerticalLineScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("verticalLineScroll"))
	return rv
}


// SetVerticalLineScroll sets the value of the verticalLineScroll property.
// The scroll view’s vertical line by line scroll amount.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalLineScroll
func (s_ ScrollView) SetVerticalLineScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVerticalLineScroll:"), value)
}
// The amount of the document view kept visible when scrolling vertically page by page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalPageScroll
func (s_ ScrollView) VerticalPageScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("verticalPageScroll"))
	return rv
}


// SetVerticalPageScroll sets the value of the verticalPageScroll property.
// The amount of the document view kept visible when scrolling vertically page by page.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalPageScroll
func (s_ ScrollView) SetVerticalPageScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVerticalPageScroll:"), value)
}
// The scroll view’s vertical ruler view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalRulerView
func (s_ ScrollView) VerticalRulerView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("verticalRulerView"))
	return rv
}


// SetVerticalRulerView sets the value of the verticalRulerView property.
// The scroll view’s vertical ruler view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalRulerView
func (s_ ScrollView) SetVerticalRulerView(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVerticalRulerView:"), value)
}
// The scroll view’s vertical scrolling elasticity mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalScrollElasticity
func (s_ ScrollView) VerticalScrollElasticity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("verticalScrollElasticity"))
	return rv
}


// SetVerticalScrollElasticity sets the value of the verticalScrollElasticity property.
// The scroll view’s vertical scrolling elasticity mode.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalScrollElasticity
func (s_ ScrollView) SetVerticalScrollElasticity(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVerticalScrollElasticity:"), value)
}
// The scroll view’s vertical scroller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalScroller
func (s_ ScrollView) VerticalScroller() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("verticalScroller"))
	return rv
}


// SetVerticalScroller sets the value of the verticalScroller property.
// The scroll view’s vertical scroller.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalScroller
func (s_ ScrollView) SetVerticalScroller(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVerticalScroller:"), value)
}

