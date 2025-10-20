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
	SetFrameFromContentFrame(contentFrame coregraphics.CGRect)
	SetTitleWithMnemonic(stringWithAmpersand string)
	SizeToFit()
}

// A stylized rectangular box with an optional title.
//
// Use box objects to visually group the contents of your window. For example, you might use boxes to group related views. Use an object to configure the appearance of the box.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/setFrameFromContentFrame(_:)
func (b_ Box) SetFrameFromContentFrame(contentFrame coregraphics.CGRect) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFrameFromContentFrame:"), contentFrame)
}

// Sets the title of the receiver with a character denoted as an access key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/setTitleWithMnemonic:
func (b_ Box) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitleWithMnemonic:"), objc.String(stringWithAmpersand))
}

// Resizes and moves the receiver’s content view so it just encloses its subviews.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/sizeToFit()
func (b_ Box) SizeToFit() {
	objc.Send[objc.ID](b_.ID, objc.Sel("sizeToFit"))
}

// The color of the receiver’s border when the receiver is a custom box with a simple line border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderColor
func (b_ Box) BorderColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("borderColor"))
	return rv
}


// SetBorderColor sets the value of the borderColor property.
// The color of the receiver’s border when the receiver is a custom box with a simple line border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderColor
func (b_ Box) SetBorderColor(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderColor:"), value)
}
// The rectangle in which the receiver’s border is drawn.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderRect
func (b_ Box) BorderRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](b_.ID, objc.Sel("borderRect"))
	return rv
}

// The receiver’s border type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderType
func (b_ Box) BorderType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("borderType"))
	return rv
}


// SetBorderType sets the value of the borderType property.
// The receiver’s border type.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderType
func (b_ Box) SetBorderType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderType:"), value)
}
// The width of the receiver’s border when the receiver is a custom box with a simple line border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderWidth
func (b_ Box) BorderWidth() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("borderWidth"))
	return rv
}


// SetBorderWidth sets the value of the borderWidth property.
// The width of the receiver’s border when the receiver is a custom box with a simple line border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderWidth
func (b_ Box) SetBorderWidth(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderWidth:"), value)
}
// The receiver’s box type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/boxType-swift.property
func (b_ Box) BoxType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("boxType"))
	return rv
}


// SetBoxType sets the value of the boxType property.
// The receiver’s box type.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/boxType-swift.property
func (b_ Box) SetBoxType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBoxType:"), value)
}
// The receiver’s content view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentView
func (b_ Box) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("contentView"))
	return rv
}


// SetContentView sets the value of the contentView property.
// The receiver’s content view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentView
func (b_ Box) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentView:"), value)
}
// The distances between the border and the content view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentViewMargins
func (b_ Box) ContentViewMargins() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](b_.ID, objc.Sel("contentViewMargins"))
	return rv
}


// SetContentViewMargins sets the value of the contentViewMargins property.
// The distances between the border and the content view.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentViewMargins
func (b_ Box) SetContentViewMargins(value coregraphics.CGSize) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentViewMargins:"), value)
}
// The radius of the receiver’s corners when the receiver is a custom box with a simple line border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/cornerRadius
func (b_ Box) CornerRadius() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("cornerRadius"))
	return rv
}


// SetCornerRadius sets the value of the cornerRadius property.
// The radius of the receiver’s corners when the receiver is a custom box with a simple line border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/cornerRadius
func (b_ Box) SetCornerRadius(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCornerRadius:"), value)
}
// The color of the receiver’s background when the receiver is a custom box with a simple line border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/fillColor
func (b_ Box) FillColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("fillColor"))
	return rv
}


// SetFillColor sets the value of the fillColor property.
// The color of the receiver’s background when the receiver is a custom box with a simple line border.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/fillColor
func (b_ Box) SetFillColor(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFillColor:"), value)
}
// A Boolean value that indicates whether the receiver is transparent.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/isTransparent
func (b_ Box) Transparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("transparent"))
	return rv
}


// SetTransparent sets the value of the transparent property.
// A Boolean value that indicates whether the receiver is transparent.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/isTransparent
func (b_ Box) SetTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTransparent:"), value)
}
// The receiver’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/title
func (b_ Box) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The receiver’s title.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/title
func (b_ Box) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), value)
}
// The cell used to display the receiver’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleCell
func (b_ Box) TitleCell() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("titleCell"))
	return rv
}

// The font object used to draw the receiver’s title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleFont
func (b_ Box) TitleFont() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("titleFont"))
	return rv
}


// SetTitleFont sets the value of the titleFont property.
// The font object used to draw the receiver’s title.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleFont
func (b_ Box) SetTitleFont(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitleFont:"), value)
}
// A constant representing the title position.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titlePosition-swift.property
func (b_ Box) TitlePosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("titlePosition"))
	return rv
}


// SetTitlePosition sets the value of the titlePosition property.
// A constant representing the title position.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titlePosition-swift.property
func (b_ Box) SetTitlePosition(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitlePosition:"), value)
}
// The rectangle in which the receiver’s title is drawn.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleRect
func (b_ Box) TitleRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](b_.ID, objc.Sel("titleRect"))
	return rv
}



