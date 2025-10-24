// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class NSBox */


/* debug [class_header]: Header for NSBox */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Box */
// An interface definition for the [Box] class.
type IBox interface {
	IView
	
/* debug [class_interface_properties]: Properties for Box */
	// properties:
	BorderColor() IColor
	SetBorderColor(value IColor)
	BorderRect() Rect /* not a class type */
	BorderType() BorderType
	SetBorderType(value BorderType)
	BorderWidth() float64
	SetBorderWidth(value float64)
	BoxType() BoxType
	SetBoxType(value BoxType)
	ContentView() IView
	SetContentView(value IView)
	ContentViewMargins() Size /* not a class type */
	SetContentViewMargins(value Size /* not a class type */)
	CornerRadius() float64
	SetCornerRadius(value float64)
	FillColor() IColor
	SetFillColor(value IColor)
	Transparent() bool
	SetTransparent(value bool)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	TitleCell() objc.ID
	TitleFont() IFont
	SetTitleFont(value IFont)
	TitlePosition() TitlePosition
	SetTitlePosition(value TitlePosition)
	TitleRect() Rect /* not a class type */
	IsTransparent() bool
	SetIsTransparent(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Box */
	// methods:
	SetFrameFromContentFrame(contentFrame Rect /* not a class type */)
	SizeToFit()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Box */
// Alloc allocates a new instance without initialization.
func (bc _BoxClass) Alloc() Box {
	rv := objc.Send[Box](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Box */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Box *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Box */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Box */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Box */

// Places the receiver so its content view lies on the specified frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/setFrameFromContentFrame(_:)
func (b_ Box) SetFrameFromContentFrame(contentFrame Rect /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFrameFromContentFrame:"), contentFrame)
}/* debug [instance_methods/method]: SetFrameFromContentFrame */


// Resizes and moves the receiver’s content view so it just encloses its subviews.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/sizeToFit()
func (b_ Box) SizeToFit() {
	objc.Send[objc.ID](b_.ID, objc.Sel("sizeToFit"))
}/* debug [instance_methods/method]: SizeToFit */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Box */

// The color of the receiver’s border when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderColor
func (b_ Box) BorderColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("borderColor"))
	return rv
}/* debug [instance_properties/getter]: borderColor */


// The color of the receiver’s border when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderColor
func (b_ Box) SetBorderColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderColor:"), value)
}/* debug [instance_properties/setter]: borderColor */


// The rectangle in which the receiver’s border is drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderRect
func (b_ Box) BorderRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](b_.ID, objc.Sel("borderRect"))
	return rv
}/* debug [instance_properties/getter]: borderRect */


// The receiver’s border type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderType
func (b_ Box) BorderType() BorderType {
	rv := objc.Send[BorderType](b_.ID, objc.Sel("borderType"))
	return rv
}/* debug [instance_properties/getter]: borderType */


// The receiver’s border type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderType
func (b_ Box) SetBorderType(value BorderType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderType:"), value)
}/* debug [instance_properties/setter]: borderType */


// The width of the receiver’s border when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderWidth
func (b_ Box) BorderWidth() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("borderWidth"))
	return rv
}/* debug [instance_properties/getter]: borderWidth */


// The width of the receiver’s border when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/borderWidth
func (b_ Box) SetBorderWidth(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBorderWidth:"), value)
}/* debug [instance_properties/setter]: borderWidth */


// The receiver’s box type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/boxType-swift.property
func (b_ Box) BoxType() BoxType {
	rv := objc.Send[BoxType](b_.ID, objc.Sel("boxType"))
	return rv
}/* debug [instance_properties/getter]: boxType */


// The receiver’s box type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/boxType-swift.property
func (b_ Box) SetBoxType(value BoxType) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBoxType:"), value)
}/* debug [instance_properties/setter]: boxType */


// The receiver’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentView
func (b_ Box) ContentView() IView {
	rv := objc.Send[View](b_.ID, objc.Sel("contentView"))
	return rv
}/* debug [instance_properties/getter]: contentView */


// The receiver’s content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentView
func (b_ Box) SetContentView(value IView) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentView:"), value)
}/* debug [instance_properties/setter]: contentView */


// The distances between the border and the content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentViewMargins
func (b_ Box) ContentViewMargins() Size /* not a class type */ {
	rv := objc.Send[Size](b_.ID, objc.Sel("contentViewMargins"))
	return rv
}/* debug [instance_properties/getter]: contentViewMargins */


// The distances between the border and the content view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/contentViewMargins
func (b_ Box) SetContentViewMargins(value Size /* not a class type */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setContentViewMargins:"), value)
}/* debug [instance_properties/setter]: contentViewMargins */


// The radius of the receiver’s corners when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/cornerRadius
func (b_ Box) CornerRadius() float64 {
	rv := objc.Send[float64](b_.ID, objc.Sel("cornerRadius"))
	return rv
}/* debug [instance_properties/getter]: cornerRadius */


// The radius of the receiver’s corners when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/cornerRadius
func (b_ Box) SetCornerRadius(value float64) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCornerRadius:"), value)
}/* debug [instance_properties/setter]: cornerRadius */


// The color of the receiver’s background when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/fillColor
func (b_ Box) FillColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("fillColor"))
	return rv
}/* debug [instance_properties/getter]: fillColor */


// The color of the receiver’s background when the receiver is a custom box with a simple line border.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/fillColor
func (b_ Box) SetFillColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFillColor:"), value)
}/* debug [instance_properties/setter]: fillColor */


// A Boolean value that indicates whether the receiver is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/isTransparent
func (b_ Box) Transparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("transparent"))
	return rv
}/* debug [instance_properties/getter]: transparent */


// A Boolean value that indicates whether the receiver is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/isTransparent
func (b_ Box) SetTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTransparent:"), value)
}/* debug [instance_properties/setter]: transparent */


// The receiver’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/title
func (b_ Box) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// The receiver’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/title
func (b_ Box) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// The cell used to display the receiver’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleCell
func (b_ Box) TitleCell() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("titleCell"))
	return rv
}/* debug [instance_properties/getter]: titleCell */


// The font object used to draw the receiver’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleFont
func (b_ Box) TitleFont() IFont {
	rv := objc.Send[Font](b_.ID, objc.Sel("titleFont"))
	return rv
}/* debug [instance_properties/getter]: titleFont */


// The font object used to draw the receiver’s title.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleFont
func (b_ Box) SetTitleFont(value IFont) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitleFont:"), value)
}/* debug [instance_properties/setter]: titleFont */


// A constant representing the title position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titlePosition-swift.property
func (b_ Box) TitlePosition() TitlePosition {
	rv := objc.Send[TitlePosition](b_.ID, objc.Sel("titlePosition"))
	return rv
}/* debug [instance_properties/getter]: titlePosition */


// A constant representing the title position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titlePosition-swift.property
func (b_ Box) SetTitlePosition(value TitlePosition) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitlePosition:"), value)
}/* debug [instance_properties/setter]: titlePosition */


// The rectangle in which the receiver’s title is drawn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/titleRect
func (b_ Box) TitleRect() Rect /* not a class type */ {
	rv := objc.Send[Rect](b_.ID, objc.Sel("titleRect"))
	return rv
}/* debug [instance_properties/getter]: titleRect */


// A Boolean value that indicates whether the receiver is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/istransparent
func (b_ Box) IsTransparent() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isTransparent"))
	return rv
}/* debug [instance_properties/getter]: isTransparent */


// A Boolean value that indicates whether the receiver is transparent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbox/istransparent
func (b_ Box) SetIsTransparent(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsTransparent:"), value)
}/* debug [instance_properties/setter]: isTransparent */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSBox */



