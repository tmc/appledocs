// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [Box] class.
var (
	BoxClass     _BoxClass
	BoxClassOnce sync.Once
)

func getBoxClass() _BoxClass {
	BoxClassOnce.Do(func() {
		BoxClass = _BoxClass{objc.GetClass("NSBox")}
	})
	return BoxClass
}

type _BoxClass struct {
	class objc.Class
}

// An interface definition for the [Box] class.
type IBox interface {
	IView
	// properties:
	BorderColor() IColor
	SetBorderColor(value IColor)
	BorderRect() coregraphics.CGRect
	BorderType() BorderType
	SetBorderType(value BorderType)
	BorderWidth() float64 /* primitive/slice/pointer. */
	SetBorderWidth(value float64 /* primitive/slice/pointer. */)
	BoxType() BoxType
	SetBoxType(value BoxType)
	ContentView() IView
	SetContentView(value IView)
	ContentViewMargins() coregraphics.CGSize
	SetContentViewMargins(value coregraphics.CGSize)
	CornerRadius() float64 /* primitive/slice/pointer. */
	SetCornerRadius(value float64 /* primitive/slice/pointer. */)
	FillColor() IColor
	SetFillColor(value IColor)
	Transparent() bool /* primitive/slice/pointer. */
	SetTransparent(value bool /* primitive/slice/pointer. */)
	Title() string /* primitive/slice/pointer. */
	SetTitle(value string /* primitive/slice/pointer. */)
	TitleCell() objc.ID
	TitleFont() IFont
	SetTitleFont(value IFont)
	TitlePosition() TitlePosition
	SetTitlePosition(value TitlePosition)
	TitleRect() coregraphics.CGRect
	IsTransparent() bool /* primitive/slice/pointer. */
	SetIsTransparent(value bool /* primitive/slice/pointer. */)
	// methods:
	SetFrameFromContentFrame(contentFrame coregraphics.CGRect)
	SizeToFit()
}

// A stylized rectangular box with an optional title.
//
// Use box objects to visually group the contents of your window. For example, you might use boxes to group related views. Use an object to configure the appearance of the box.


// A stylized rectangular box with an optional title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox
type Box struct {
	View
}

// BoxFrom constructs a [Box] from an unsafe.Pointer.
//
// A stylized rectangular box with an optional title.
func BoxFrom(ptr unsafe.Pointer) Box {
	return Box{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BoxClass) Alloc() Box {
	rv := objc.Send[Box](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BoxClass) New() Box {
	rv := objc.Send[Box](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ Box) Init() Box {
	rv := objc.Send[Box](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ Box) Autorelease() Box {
	rv := objc.Send[Box](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBox creates a new Box instance.
func NewBox() Box {
	return getBoxClass().New()
}



// Places the receiver so its content view lies on the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/setFrameFromContentFrame(_:)
func (b_ Box) SetFrameFromContentFrame(contentFrame coregraphics.CGRect) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFrameFromContentFrame:"), contentFrame)
}


// Resizes and moves the receiver’s content view so it just encloses its subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/sizeToFit()
func (b_ Box) SizeToFit() {
	objc.Send[objc.ID](b_.ID, objc.Sel("sizeToFit"))
}


// The color of the receiver’s border when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderColor
func (b_ Box) BorderColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("borderColor"))
	return rv
}


// The color of the receiver’s border when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderColor
func (b_ Box) SetBorderColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderColor:"), value)
}


// The rectangle in which the receiver’s border is drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderRect
func (b_ Box) BorderRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](b_.ID, objc.Sel("borderRect"))
	return rv
}


// The receiver’s border type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderType
func (b_ Box) BorderType() BorderType {
	rv := objc.Send[BorderType](b_.ID, objc.Sel("borderType"))
	return rv
}


// The receiver’s border type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderType
func (b_ Box) SetBorderType(value BorderType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderType:"), value)
}


// The width of the receiver’s border when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderWidth
func (b_ Box) BorderWidth() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("borderWidth"))
	return rv
}


// The width of the receiver’s border when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderWidth
func (b_ Box) SetBorderWidth(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderWidth:"), value)
}


// The receiver’s box type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/boxType-swift.property
func (b_ Box) BoxType() BoxType {
	rv := objc.Send[BoxType](b_.ID, objc.Sel("boxType"))
	return rv
}


// The receiver’s box type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/boxType-swift.property
func (b_ Box) SetBoxType(value BoxType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBoxType:"), value)
}


// The receiver’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentView
func (b_ Box) ContentView() IView {
	rv := objc.Send[View](b_.ID, objc.Sel("contentView"))
	return rv
}


// The receiver’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentView
func (b_ Box) SetContentView(value IView) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentView:"), value)
}


// The distances between the border and the content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentViewMargins
func (b_ Box) ContentViewMargins() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](b_.ID, objc.Sel("contentViewMargins"))
	return rv
}


// The distances between the border and the content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentViewMargins
func (b_ Box) SetContentViewMargins(value coregraphics.CGSize) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentViewMargins:"), value)
}


// The radius of the receiver’s corners when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/cornerRadius
func (b_ Box) CornerRadius() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](b_.ID, objc.Sel("cornerRadius"))
	return rv
}


// The radius of the receiver’s corners when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/cornerRadius
func (b_ Box) SetCornerRadius(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCornerRadius:"), value)
}


// The color of the receiver’s background when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/fillColor
func (b_ Box) FillColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("fillColor"))
	return rv
}


// The color of the receiver’s background when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/fillColor
func (b_ Box) SetFillColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFillColor:"), value)
}


// A Boolean value that indicates whether the receiver is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/isTransparent
func (b_ Box) Transparent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("transparent"))
	return rv
}


// A Boolean value that indicates whether the receiver is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/isTransparent
func (b_ Box) SetTransparent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTransparent:"), value)
}


// The receiver’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/title
func (b_ Box) Title() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](b_.ID, objc.Sel("title"))
	return rv
}


// The receiver’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/title
func (b_ Box) SetTitle(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The cell used to display the receiver’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleCell
func (b_ Box) TitleCell() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("titleCell"))
	return rv
}


// The font object used to draw the receiver’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleFont
func (b_ Box) TitleFont() IFont {
	rv := objc.Send[Font](b_.ID, objc.Sel("titleFont"))
	return rv
}


// The font object used to draw the receiver’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleFont
func (b_ Box) SetTitleFont(value IFont) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitleFont:"), value)
}


// A constant representing the title position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titlePosition-swift.property
func (b_ Box) TitlePosition() TitlePosition {
	rv := objc.Send[TitlePosition](b_.ID, objc.Sel("titlePosition"))
	return rv
}


// A constant representing the title position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titlePosition-swift.property
func (b_ Box) SetTitlePosition(value TitlePosition) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitlePosition:"), value)
}


// The rectangle in which the receiver’s title is drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleRect
func (b_ Box) TitleRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](b_.ID, objc.Sel("titleRect"))
	return rv
}


// A Boolean value that indicates whether the receiver is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/istransparent
func (b_ Box) IsTransparent() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransparent"))
	return rv
}


// A Boolean value that indicates whether the receiver is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/istransparent
func (b_ Box) SetIsTransparent(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsTransparent:"), value)
}



