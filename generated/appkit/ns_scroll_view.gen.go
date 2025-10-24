// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class NSScrollView */


/* debug [class_header]: Header for NSScrollView */
// The class instance for the [ScrollView] class.
var (
	ScrollViewClass     _ScrollViewClass
	ScrollViewClassOnce sync.Once
)

func getScrollViewClass() _ScrollViewClass {
	ScrollViewClassOnce.Do(func() {
		ScrollViewClass = _ScrollViewClass{objc.GetClass("NSScrollView")}
	})
	return ScrollViewClass
}

type _ScrollViewClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ScrollView */
// An interface definition for the [ScrollView] class.
type IScrollView interface {
	IView
	
/* debug [class_interface_properties]: Properties for ScrollView */
	// properties:
	AllowsMagnification() bool
	SetAllowsMagnification(value bool)
	AutohidesScrollers() bool
	SetAutohidesScrollers(value bool)
	AutomaticallyAdjustsContentInsets() bool
	SetAutomaticallyAdjustsContentInsets(value bool)
	BackgroundColor() IColor
	SetBackgroundColor(value IColor)
	BorderType() BorderType
	SetBorderType(value BorderType)
	ContentInsets() foundation.EdgeInsets
	SetContentInsets(value foundation.EdgeInsets)
	ContentSize() Size /* not a class type */
	ContentView() IClipView
	SetContentView(value IClipView)
	DocumentCursor() ICursor
	SetDocumentCursor(value ICursor)
	DocumentView() IView
	SetDocumentView(value IView)
	DocumentVisibleRect() Rect /* not a class type */
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	FindBarPosition() ScrollViewFindBarPosition
	SetFindBarPosition(value ScrollViewFindBarPosition)
	HasHorizontalRuler() bool
	SetHasHorizontalRuler(value bool)
	HasHorizontalScroller() bool
	SetHasHorizontalScroller(value bool)
	HasVerticalRuler() bool
	SetHasVerticalRuler(value bool)
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	HorizontalLineScroll() float64
	SetHorizontalLineScroll(value float64)
	HorizontalPageScroll() float64
	SetHorizontalPageScroll(value float64)
	HorizontalRulerView() IRulerView
	SetHorizontalRulerView(value IRulerView)
	HorizontalScrollElasticity() ScrollElasticity
	SetHorizontalScrollElasticity(value ScrollElasticity)
	HorizontalScroller() IScroller
	SetHorizontalScroller(value IScroller)
	LineScroll() float64
	SetLineScroll(value float64)
	Magnification() float64
	SetMagnification(value float64)
	MaxMagnification() float64
	SetMaxMagnification(value float64)
	MinMagnification() float64
	SetMinMagnification(value float64)
	PageScroll() float64
	SetPageScroll(value float64)
	RulersVisible() bool
	SetRulersVisible(value bool)
	ScrollerInsets() foundation.EdgeInsets
	SetScrollerInsets(value foundation.EdgeInsets)
	ScrollerKnobStyle() ScrollerKnobStyle
	SetScrollerKnobStyle(value ScrollerKnobStyle)
	ScrollerStyle() ScrollerStyle
	SetScrollerStyle(value ScrollerStyle)
	ScrollsDynamically() bool
	SetScrollsDynamically(value bool)
	UsesPredominantAxisScrolling() bool
	SetUsesPredominantAxisScrolling(value bool)
	VerticalLineScroll() float64
	SetVerticalLineScroll(value float64)
	VerticalPageScroll() float64
	SetVerticalPageScroll(value float64)
	VerticalRulerView() IRulerView
	SetVerticalRulerView(value IRulerView)
	VerticalScrollElasticity() ScrollElasticity
	SetVerticalScrollElasticity(value ScrollElasticity)
	VerticalScroller() IScroller
	SetVerticalScroller(value IScroller)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ScrollView */
	// methods:
	AddFloatingSubviewForAxis(view IView, axis EventGestureAxis)
	FlashScrollers()
	MagnifyToFitRect(rect Rect /* not a class type */)
	ReflectScrolledClipView(cView IClipView)
	ScrollWheel(event IEvent)
	SetMagnificationCenteredAtPoint(magnification float64, point vision.Point)
	Tile()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ScrollView */
// Alloc allocates a new instance without initialization.
func (sc _ScrollViewClass) Alloc() ScrollView {
	rv := objc.Send[ScrollView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ScrollView */
// A view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view.
//
// The class is the central coordinator for AppKit’s scrolling machinery, which is composed of this class, and the and classes. When using an object within a scroll view (the usual configuration), you should issue messages that control background drawing state to the scroll view directly, rather than messaging the clip view.


// A view that displays a portion of a document view and provides scroll bars that allow the user to move the document view within the scroll view.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ScrollView */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/init(coder:)
func NewScrollViewWithCoder(coder foundation.Coder) ScrollView {
	instance := getScrollViewClass().Alloc()
	rv := objc.Send[ScrollView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScrollViewWithCoder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/init(frame:)
func NewScrollViewWithFrame(frameRect Rect /* not a class type */) ScrollView {
	instance := getScrollViewClass().Alloc()
	rv := objc.Send[ScrollView](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewScrollViewWithFrame */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ScrollView */

// Returns the content size calculated from the frame size and the specified specifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentSize(forFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (sc _ScrollViewClass) ContentSizeForFrameSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(fSize Size /* not a class type */, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ BorderType, controlSize ControlSize, scrollerStyle ScrollerStyle) Size /* not a class type */ {
	rv := objc.Send[Size](objc.ID(sc.class), objc.Sel("contentSizeForFrameSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), fSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContentSizeForFrameSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle) */


// Returns the content size calculated from the frame size and the specified specifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:
func (sc _ScrollViewClass) ContentSizeForFrameSizeHasHorizontalScrollerHasVerticalScrollerBorderType(fSize Size /* not a class type */, hFlag bool, vFlag bool, type_ BorderType) Size /* not a class type */ {
	rv := objc.Send[Size](objc.ID(sc.class), objc.Sel("contentSizeForFrameSize:hasHorizontalScroller:hasVerticalScroller:borderType:"), fSize, hFlag, vFlag, type_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ContentSizeForFrameSizeHasHorizontalScrollerHasVerticalScrollerBorderType) */


// Returns the frame size of a scroll view that contains a content view with the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/frameSize(forContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:)
func (sc _ScrollViewClass) FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle(cSize Size /* not a class type */, horizontalScrollerClass objc.Class, verticalScrollerClass objc.Class, type_ BorderType, controlSize ControlSize, scrollerStyle ScrollerStyle) Size /* not a class type */ {
	rv := objc.Send[Size](objc.ID(sc.class), objc.Sel("frameSizeForContentSize:horizontalScrollerClass:verticalScrollerClass:borderType:controlSize:scrollerStyle:"), cSize, horizontalScrollerClass, verticalScrollerClass, type_, controlSize, scrollerStyle)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FrameSizeForContentSizeHorizontalScrollerClassVerticalScrollerClassBorderTypeControlSizeScrollerStyle) */


// Returns the frame size of an scroll view that contains a content view with the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:
func (sc _ScrollViewClass) FrameSizeForContentSizeHasHorizontalScrollerHasVerticalScrollerBorderType(cSize Size /* not a class type */, hFlag bool, vFlag bool, type_ BorderType) Size /* not a class type */ {
	rv := objc.Send[Size](objc.ID(sc.class), objc.Sel("frameSizeForContentSize:hasHorizontalScroller:hasVerticalScroller:borderType:"), cSize, hFlag, vFlag, type_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FrameSizeForContentSizeHasHorizontalScrollerHasVerticalScrollerBorderType) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ScrollView */

// Returns the default class to be used for ruler objects in NSScrollViews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/rulerViewClass
func (sc _ScrollViewClass) RulerViewClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(sc.class), objc.Sel("rulerViewClass"))
	return rv
}/* debug [class_properties_class/property]: rulerViewClass */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ScrollView */

// Adds a floating subview to the document view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/addFloatingSubview(_:for:)
func (s_ ScrollView) AddFloatingSubviewForAxis(view IView, axis EventGestureAxis) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addFloatingSubview:forAxis:"), view, axis)
}/* debug [instance_methods/method]: AddFloatingSubviewForAxis */


// Flash the overlay scroll bars.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/flashScrollers()
func (s_ ScrollView) FlashScrollers() {
	objc.Send[objc.ID](s_.ID, objc.Sel("flashScrollers"))
}/* debug [instance_methods/method]: FlashScrollers */


// Magnifies the content view proportionally such that the given rectangle fits centered in the scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/magnify(toFit:)
func (s_ ScrollView) MagnifyToFitRect(rect Rect /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("magnifyToFitRect:"), rect)
}/* debug [instance_methods/method]: MagnifyToFitRect */


// Adjusts the receiver’s scrollers to reflect the size and positioning of its content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/reflectScrolledClipView(_:)
func (s_ ScrollView) ReflectScrolledClipView(cView IClipView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("reflectScrolledClipView:"), cView)
}/* debug [instance_methods/method]: ReflectScrolledClipView */


// Scrolls the receiver up or down, in response to the user moving the mouse’s scroll wheel specified by .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollWheel(with:)
func (s_ ScrollView) ScrollWheel(event IEvent) {
	objc.Send[objc.ID](s_.ID, objc.Sel("scrollWheel:"), event)
}/* debug [instance_methods/method]: ScrollWheel */


// Magnify the content by the given amount and center the result on the given point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/setMagnification(_:centeredAt:)
func (s_ ScrollView) SetMagnificationCenteredAtPoint(magnification float64, point vision.Point) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMagnification:centeredAtPoint:"), magnification, point)
}/* debug [instance_methods/method]: SetMagnificationCenteredAtPoint */


// Lays out the components of the receiver: the content view, the scrollers, and the ruler views.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/tile()
func (s_ ScrollView) Tile() {
	objc.Send[objc.ID](s_.ID, objc.Sel("tile"))
}/* debug [instance_methods/method]: Tile */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ScrollView */

// Allows the user to magnify the scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/allowsMagnification
func (s_ ScrollView) AllowsMagnification() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsMagnification"))
	return rv
}/* debug [instance_properties/getter]: allowsMagnification */


// Allows the user to magnify the scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/allowsMagnification
func (s_ ScrollView) SetAllowsMagnification(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsMagnification:"), value)
}/* debug [instance_properties/setter]: allowsMagnification */


// A Boolean that indicates whether the scroll view automatically hides its scroll bars when they are not needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/autohidesScrollers
func (s_ ScrollView) AutohidesScrollers() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("autohidesScrollers"))
	return rv
}/* debug [instance_properties/getter]: autohidesScrollers */


// A Boolean that indicates whether the scroll view automatically hides its scroll bars when they are not needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/autohidesScrollers
func (s_ ScrollView) SetAutohidesScrollers(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutohidesScrollers:"), value)
}/* debug [instance_properties/setter]: autohidesScrollers */


// A Boolean that indicates whether the scroll view automatically adjusts its content insets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/automaticallyAdjustsContentInsets
func (s_ ScrollView) AutomaticallyAdjustsContentInsets() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("automaticallyAdjustsContentInsets"))
	return rv
}/* debug [instance_properties/getter]: automaticallyAdjustsContentInsets */


// A Boolean that indicates whether the scroll view automatically adjusts its content insets.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/automaticallyAdjustsContentInsets
func (s_ ScrollView) SetAutomaticallyAdjustsContentInsets(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAutomaticallyAdjustsContentInsets:"), value)
}/* debug [instance_properties/setter]: automaticallyAdjustsContentInsets */


// The color of the content view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/backgroundColor
func (s_ ScrollView) BackgroundColor() IColor {
	rv := objc.Send[Color](s_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The color of the content view’s background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/backgroundColor
func (s_ ScrollView) SetBackgroundColor(value IColor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// A value that specifies the appearance of the scroll view’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/borderType
func (s_ ScrollView) BorderType() BorderType {
	rv := objc.Send[BorderType](s_.ID, objc.Sel("borderType"))
	return rv
}/* debug [instance_properties/getter]: borderType */


// A value that specifies the appearance of the scroll view’s border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/borderType
func (s_ ScrollView) SetBorderType(value BorderType) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setBorderType:"), value)
}/* debug [instance_properties/setter]: borderType */


// The distance that the scroll view’s subviews are inset from the enclosing scroll view during tiling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentInsets
func (s_ ScrollView) ContentInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](s_.ID, objc.Sel("contentInsets"))
	return rv
}/* debug [instance_properties/getter]: contentInsets */


// The distance that the scroll view’s subviews are inset from the enclosing scroll view during tiling.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentInsets
func (s_ ScrollView) SetContentInsets(value foundation.EdgeInsets) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContentInsets:"), value)
}/* debug [instance_properties/setter]: contentInsets */


// The size of the scroll view’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentSize
func (s_ ScrollView) ContentSize() Size /* not a class type */ {
	rv := objc.Send[Size](s_.ID, objc.Sel("contentSize"))
	return rv
}/* debug [instance_properties/getter]: contentSize */


// The scroll view’s content view, the view that clips the document view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentView
func (s_ ScrollView) ContentView() IClipView {
	rv := objc.Send[ClipView](s_.ID, objc.Sel("contentView"))
	return rv
}/* debug [instance_properties/getter]: contentView */


// The scroll view’s content view, the view that clips the document view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/contentView
func (s_ ScrollView) SetContentView(value IClipView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setContentView:"), value)
}/* debug [instance_properties/setter]: contentView */


// The content view’s document cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/documentCursor
func (s_ ScrollView) DocumentCursor() ICursor {
	rv := objc.Send[Cursor](s_.ID, objc.Sel("documentCursor"))
	return rv
}/* debug [instance_properties/getter]: documentCursor */


// The content view’s document cursor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/documentCursor
func (s_ ScrollView) SetDocumentCursor(value ICursor) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDocumentCursor:"), value)
}/* debug [instance_properties/setter]: documentCursor */


// The view the scroll view scrolls within its content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/documentView
func (s_ ScrollView) DocumentView() IView {
	rv := objc.Send[View](s_.ID, objc.Sel("documentView"))
	return rv
}/* debug [instance_properties/getter]: documentView */


// The view the scroll view scrolls within its content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/documentView
func (s_ ScrollView) SetDocumentView(value IView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDocumentView:"), value)
}/* debug [instance_properties/setter]: documentView */


// The portion of the document view, in its own coordinate system, visible through the scroll view’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/documentVisibleRect
func (s_ ScrollView) DocumentVisibleRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](s_.ID, objc.Sel("documentVisibleRect"))
	return rv
}/* debug [instance_properties/getter]: documentVisibleRect */


// A Boolean that indicates whether the scroll view draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/drawsBackground
func (s_ ScrollView) DrawsBackground() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("drawsBackground"))
	return rv
}/* debug [instance_properties/getter]: drawsBackground */


// A Boolean that indicates whether the scroll view draws its background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/drawsBackground
func (s_ ScrollView) SetDrawsBackground(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDrawsBackground:"), value)
}/* debug [instance_properties/setter]: drawsBackground */


// The position of the find bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/findBarPosition-swift.property
func (s_ ScrollView) FindBarPosition() ScrollViewFindBarPosition {
	rv := objc.Send[ScrollViewFindBarPosition](s_.ID, objc.Sel("findBarPosition"))
	return rv
}/* debug [instance_properties/getter]: findBarPosition */


// The position of the find bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/findBarPosition-swift.property
func (s_ ScrollView) SetFindBarPosition(value ScrollViewFindBarPosition) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setFindBarPosition:"), value)
}/* debug [instance_properties/setter]: findBarPosition */


// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasHorizontalRuler
func (s_ ScrollView) HasHorizontalRuler() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasHorizontalRuler"))
	return rv
}/* debug [instance_properties/getter]: hasHorizontalRuler */


// A Boolean that indicates whether the scroll view keeps a horizontal ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasHorizontalRuler
func (s_ ScrollView) SetHasHorizontalRuler(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHasHorizontalRuler:"), value)
}/* debug [instance_properties/setter]: hasHorizontalRuler */


// A Boolean that indicates whether the scroll view has a horizontal scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasHorizontalScroller
func (s_ ScrollView) HasHorizontalScroller() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasHorizontalScroller"))
	return rv
}/* debug [instance_properties/getter]: hasHorizontalScroller */


// A Boolean that indicates whether the scroll view has a horizontal scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasHorizontalScroller
func (s_ ScrollView) SetHasHorizontalScroller(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHasHorizontalScroller:"), value)
}/* debug [instance_properties/setter]: hasHorizontalScroller */


// A Boolean that indicates whether the scroll view keeps a vertical ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasVerticalRuler
func (s_ ScrollView) HasVerticalRuler() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasVerticalRuler"))
	return rv
}/* debug [instance_properties/getter]: hasVerticalRuler */


// A Boolean that indicates whether the scroll view keeps a vertical ruler object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasVerticalRuler
func (s_ ScrollView) SetHasVerticalRuler(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHasVerticalRuler:"), value)
}/* debug [instance_properties/setter]: hasVerticalRuler */


// A Boolean that indicates whether the scroll view has a vertical scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasVerticalScroller
func (s_ ScrollView) HasVerticalScroller() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("hasVerticalScroller"))
	return rv
}/* debug [instance_properties/getter]: hasVerticalScroller */


// A Boolean that indicates whether the scroll view has a vertical scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/hasVerticalScroller
func (s_ ScrollView) SetHasVerticalScroller(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHasVerticalScroller:"), value)
}/* debug [instance_properties/setter]: hasVerticalScroller */


// The scroll view’s horizontal line by line scroll amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalLineScroll
func (s_ ScrollView) HorizontalLineScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("horizontalLineScroll"))
	return rv
}/* debug [instance_properties/getter]: horizontalLineScroll */


// The scroll view’s horizontal line by line scroll amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalLineScroll
func (s_ ScrollView) SetHorizontalLineScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalLineScroll:"), value)
}/* debug [instance_properties/setter]: horizontalLineScroll */


// The amount of the document view kept visible when scrolling horizontally page by page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalPageScroll
func (s_ ScrollView) HorizontalPageScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("horizontalPageScroll"))
	return rv
}/* debug [instance_properties/getter]: horizontalPageScroll */


// The amount of the document view kept visible when scrolling horizontally page by page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalPageScroll
func (s_ ScrollView) SetHorizontalPageScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalPageScroll:"), value)
}/* debug [instance_properties/setter]: horizontalPageScroll */


// The scroll view’s horizontal ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalRulerView
func (s_ ScrollView) HorizontalRulerView() IRulerView {
	rv := objc.Send[RulerView](s_.ID, objc.Sel("horizontalRulerView"))
	return rv
}/* debug [instance_properties/getter]: horizontalRulerView */


// The scroll view’s horizontal ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalRulerView
func (s_ ScrollView) SetHorizontalRulerView(value IRulerView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalRulerView:"), value)
}/* debug [instance_properties/setter]: horizontalRulerView */


// The scroll view’s horizontal scrolling elasticity mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalScrollElasticity
func (s_ ScrollView) HorizontalScrollElasticity() ScrollElasticity {
	rv := objc.Send[ScrollElasticity](s_.ID, objc.Sel("horizontalScrollElasticity"))
	return rv
}/* debug [instance_properties/getter]: horizontalScrollElasticity */


// The scroll view’s horizontal scrolling elasticity mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalScrollElasticity
func (s_ ScrollView) SetHorizontalScrollElasticity(value ScrollElasticity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalScrollElasticity:"), value)
}/* debug [instance_properties/setter]: horizontalScrollElasticity */


// The scroll view’s horizontal scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalScroller
func (s_ ScrollView) HorizontalScroller() IScroller {
	rv := objc.Send[Scroller](s_.ID, objc.Sel("horizontalScroller"))
	return rv
}/* debug [instance_properties/getter]: horizontalScroller */


// The scroll view’s horizontal scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/horizontalScroller
func (s_ ScrollView) SetHorizontalScroller(value IScroller) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setHorizontalScroller:"), value)
}/* debug [instance_properties/setter]: horizontalScroller */


// The scroll view’s line by line scroll amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/lineScroll
func (s_ ScrollView) LineScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("lineScroll"))
	return rv
}/* debug [instance_properties/getter]: lineScroll */


// The scroll view’s line by line scroll amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/lineScroll
func (s_ ScrollView) SetLineScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLineScroll:"), value)
}/* debug [instance_properties/setter]: lineScroll */


// The amount by which the content is currently scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/magnification
func (s_ ScrollView) Magnification() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("magnification"))
	return rv
}/* debug [instance_properties/getter]: magnification */


// The amount by which the content is currently scaled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/magnification
func (s_ ScrollView) SetMagnification(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMagnification:"), value)
}/* debug [instance_properties/setter]: magnification */


// The maximum value to which the content can be magnified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/maxMagnification
func (s_ ScrollView) MaxMagnification() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxMagnification"))
	return rv
}/* debug [instance_properties/getter]: maxMagnification */


// The maximum value to which the content can be magnified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/maxMagnification
func (s_ ScrollView) SetMaxMagnification(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxMagnification:"), value)
}/* debug [instance_properties/setter]: maxMagnification */


// The minimum value to which the content can be magnified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/minMagnification
func (s_ ScrollView) MinMagnification() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("minMagnification"))
	return rv
}/* debug [instance_properties/getter]: minMagnification */


// The minimum value to which the content can be magnified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/minMagnification
func (s_ ScrollView) SetMinMagnification(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMinMagnification:"), value)
}/* debug [instance_properties/setter]: minMagnification */


// The amount of the document view kept visible when scrolling page by page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/pageScroll
func (s_ ScrollView) PageScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("pageScroll"))
	return rv
}/* debug [instance_properties/getter]: pageScroll */


// The amount of the document view kept visible when scrolling page by page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/pageScroll
func (s_ ScrollView) SetPageScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPageScroll:"), value)
}/* debug [instance_properties/setter]: pageScroll */


// Returns the default class to be used for ruler objects in NSScrollViews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/rulerViewClass
func (s_ ScrollView) RulerViewClass() objc.Class {
	rv := objc.Send[objc.Class](s_.ID, objc.Sel("rulerViewClass"))
	return rv
}/* debug [instance_properties/getter]: rulerViewClass */


// Returns the default class to be used for ruler objects in NSScrollViews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/rulerViewClass
func (s_ ScrollView) SetRulerViewClass(value objc.Class) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRulerViewClass:"), value)
}/* debug [instance_properties/setter]: rulerViewClass */


// A Boolean that indicates whether the scroll view displays its rulers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/rulersVisible
func (s_ ScrollView) RulersVisible() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("rulersVisible"))
	return rv
}/* debug [instance_properties/getter]: rulersVisible */


// A Boolean that indicates whether the scroll view displays its rulers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/rulersVisible
func (s_ ScrollView) SetRulersVisible(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setRulersVisible:"), value)
}/* debug [instance_properties/setter]: rulersVisible */


// The distance the scrollers are inset from the edge of the scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerInsets
func (s_ ScrollView) ScrollerInsets() foundation.EdgeInsets {
	rv := objc.Send[foundation.EdgeInsets](s_.ID, objc.Sel("scrollerInsets"))
	return rv
}/* debug [instance_properties/getter]: scrollerInsets */


// The distance the scrollers are inset from the edge of the scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerInsets
func (s_ ScrollView) SetScrollerInsets(value foundation.EdgeInsets) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollerInsets:"), value)
}/* debug [instance_properties/setter]: scrollerInsets */


// The knob style of scroll views that use the overlay scroller style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerKnobStyle
func (s_ ScrollView) ScrollerKnobStyle() ScrollerKnobStyle {
	rv := objc.Send[ScrollerKnobStyle](s_.ID, objc.Sel("scrollerKnobStyle"))
	return rv
}/* debug [instance_properties/getter]: scrollerKnobStyle */


// The knob style of scroll views that use the overlay scroller style.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerKnobStyle
func (s_ ScrollView) SetScrollerKnobStyle(value ScrollerKnobStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollerKnobStyle:"), value)
}/* debug [instance_properties/setter]: scrollerKnobStyle */


// The scroller style used by the scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerStyle
func (s_ ScrollView) ScrollerStyle() ScrollerStyle {
	rv := objc.Send[ScrollerStyle](s_.ID, objc.Sel("scrollerStyle"))
	return rv
}/* debug [instance_properties/getter]: scrollerStyle */


// The scroller style used by the scroll view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollerStyle
func (s_ ScrollView) SetScrollerStyle(value ScrollerStyle) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollerStyle:"), value)
}/* debug [instance_properties/setter]: scrollerStyle */


// A Boolean that indicates whether the scroll view redraws its document view while scrolling continuously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollsDynamically
func (s_ ScrollView) ScrollsDynamically() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("scrollsDynamically"))
	return rv
}/* debug [instance_properties/getter]: scrollsDynamically */


// A Boolean that indicates whether the scroll view redraws its document view while scrolling continuously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/scrollsDynamically
func (s_ ScrollView) SetScrollsDynamically(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScrollsDynamically:"), value)
}/* debug [instance_properties/setter]: scrollsDynamically */


// A Boolean that indicates whether the scroll view uses a predominant scrolling axis for content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/usesPredominantAxisScrolling
func (s_ ScrollView) UsesPredominantAxisScrolling() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("usesPredominantAxisScrolling"))
	return rv
}/* debug [instance_properties/getter]: usesPredominantAxisScrolling */


// A Boolean that indicates whether the scroll view uses a predominant scrolling axis for content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/usesPredominantAxisScrolling
func (s_ ScrollView) SetUsesPredominantAxisScrolling(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUsesPredominantAxisScrolling:"), value)
}/* debug [instance_properties/setter]: usesPredominantAxisScrolling */


// The scroll view’s vertical line by line scroll amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalLineScroll
func (s_ ScrollView) VerticalLineScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("verticalLineScroll"))
	return rv
}/* debug [instance_properties/getter]: verticalLineScroll */


// The scroll view’s vertical line by line scroll amount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalLineScroll
func (s_ ScrollView) SetVerticalLineScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVerticalLineScroll:"), value)
}/* debug [instance_properties/setter]: verticalLineScroll */


// The amount of the document view kept visible when scrolling vertically page by page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalPageScroll
func (s_ ScrollView) VerticalPageScroll() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("verticalPageScroll"))
	return rv
}/* debug [instance_properties/getter]: verticalPageScroll */


// The amount of the document view kept visible when scrolling vertically page by page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalPageScroll
func (s_ ScrollView) SetVerticalPageScroll(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVerticalPageScroll:"), value)
}/* debug [instance_properties/setter]: verticalPageScroll */


// The scroll view’s vertical ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalRulerView
func (s_ ScrollView) VerticalRulerView() IRulerView {
	rv := objc.Send[RulerView](s_.ID, objc.Sel("verticalRulerView"))
	return rv
}/* debug [instance_properties/getter]: verticalRulerView */


// The scroll view’s vertical ruler view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalRulerView
func (s_ ScrollView) SetVerticalRulerView(value IRulerView) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVerticalRulerView:"), value)
}/* debug [instance_properties/setter]: verticalRulerView */


// The scroll view’s vertical scrolling elasticity mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalScrollElasticity
func (s_ ScrollView) VerticalScrollElasticity() ScrollElasticity {
	rv := objc.Send[ScrollElasticity](s_.ID, objc.Sel("verticalScrollElasticity"))
	return rv
}/* debug [instance_properties/getter]: verticalScrollElasticity */


// The scroll view’s vertical scrolling elasticity mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalScrollElasticity
func (s_ ScrollView) SetVerticalScrollElasticity(value ScrollElasticity) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVerticalScrollElasticity:"), value)
}/* debug [instance_properties/setter]: verticalScrollElasticity */


// The scroll view’s vertical scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalScroller
func (s_ ScrollView) VerticalScroller() IScroller {
	rv := objc.Send[Scroller](s_.ID, objc.Sel("verticalScroller"))
	return rv
}/* debug [instance_properties/getter]: verticalScroller */


// The scroll view’s vertical scroller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/verticalScroller
func (s_ ScrollView) SetVerticalScroller(value IScroller) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setVerticalScroller:"), value)
}/* debug [instance_properties/setter]: verticalScroller */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSScrollView */


