
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
	"github.com/progrium/darwinkit/macos/foundation"
)

// The class instance for the [Box] class.
var BoxClass _BoxClass

func init() {
	BoxClass = _BoxClass{objc.GetClass("NSBox")}
}

type _BoxClass struct {
	objc.Class
}

// An interface definition for the [Box] class.
type IBox interface {
	ID() objc.ID
	SetFrameFromContentFrame(contentFrame foundation.Rect)
	SetTitleWithMnemonic(stringWithAmpersand string)
	SizeToFit()
}

type Box struct {
	id objc.ID
}

func BoxFrom(ptr unsafe.Pointer) Box {
	return Box{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (b_ Box) ID() objc.ID {
	return b_.id
}

// Alloc allocates a new instance without initialization.
func (bc _BoxClass) Alloc() Box {
	rv := objc.Send[Box](objc.ID(bc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (bc _BoxClass) New() Box {
	rv := objc.Send[Box](objc.ID(bc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewBox creates and returns a new initialized instance.
func NewBox() Box {
	return BoxClass.New()
}

// Init initializes the instance.
func (b_ Box) Init() Box {
	rv := objc.Send[Box](b_.ID(), selInit)
	return rv
}
// Places the receiver so its content view lies on the specified frame. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/setFrameFromContentFrame(_:)
func (b_ Box) SetFrameFromContentFrame(contentFrame foundation.Rect) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setFrameFromContentFrame:"), contentFrame)
}
// Sets the title of the receiver with a character denoted as an access key. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/setTitleWithMnemonic:
func (b_ Box) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setTitleWithMnemonic:"), stringWithAmpersand)
}
// Resizes and moves the receiver’s content view so it just encloses its subviews. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/sizeToFit()
func (b_ Box) SizeToFit() {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("sizeToFit"))
}
// The color of the receiver’s border when the receiver is a custom box with a simple line border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/borderColor
func (b_ Box) BorderColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("borderColor"))
	return rv
}
// SetBorderColor sets the value of the borderColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/borderColor
func (b_ Box) SetBorderColor(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setBorderColor:"), value)
}
// The rectangle in which the receiver’s border is drawn. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/borderRect
func (b_ Box) BorderRect() foundation.Rect {
	rv := objc.Send[foundation.Rect](b_.ID(), objc.RegisterName("borderRect"))
	return rv
}
// The receiver’s border type. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/borderType
func (b_ Box) BorderType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("borderType"))
	return rv
}
// SetBorderType sets the value of the borderType property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/borderType
func (b_ Box) SetBorderType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setBorderType:"), value)
}
// The width of the receiver’s border when the receiver is a custom box with a simple line border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/borderWidth
func (b_ Box) BorderWidth() float64 {
	rv := objc.Send[float64](b_.ID(), objc.RegisterName("borderWidth"))
	return rv
}
// SetBorderWidth sets the value of the borderWidth property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/borderWidth
func (b_ Box) SetBorderWidth(value float64) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setBorderWidth:"), value)
}
// The receiver’s box type. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/boxType-swift.property
func (b_ Box) BoxType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("boxType"))
	return rv
}
// SetBoxType sets the value of the boxType property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/boxType-swift.property
func (b_ Box) SetBoxType(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setBoxType:"), value)
}
// The receiver’s content view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/contentView
func (b_ Box) ContentView() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("contentView"))
	return rv
}
// SetContentView sets the value of the contentView property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/contentView
func (b_ Box) SetContentView(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setContentView:"), value)
}
// The distances between the border and the content view. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/contentViewMargins
func (b_ Box) ContentViewMargins() foundation.Size {
	rv := objc.Send[foundation.Size](b_.ID(), objc.RegisterName("contentViewMargins"))
	return rv
}
// SetContentViewMargins sets the value of the contentViewMargins property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/contentViewMargins
func (b_ Box) SetContentViewMargins(value foundation.Size) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setContentViewMargins:"), value)
}
// The radius of the receiver’s corners when the receiver is a custom box with a simple line border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/cornerRadius
func (b_ Box) CornerRadius() float64 {
	rv := objc.Send[float64](b_.ID(), objc.RegisterName("cornerRadius"))
	return rv
}
// SetCornerRadius sets the value of the cornerRadius property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/cornerRadius
func (b_ Box) SetCornerRadius(value float64) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setCornerRadius:"), value)
}
// The color of the receiver’s background when the receiver is a custom box with a simple line border. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/fillColor
func (b_ Box) FillColor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("fillColor"))
	return rv
}
// SetFillColor sets the value of the fillColor property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/fillColor
func (b_ Box) SetFillColor(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setFillColor:"), value)
}
// A Boolean value that indicates whether the receiver is transparent. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/isTransparent
func (b_ Box) Transparent() bool {
	rv := objc.Send[bool](b_.ID(), objc.RegisterName("transparent"))
	return rv
}
// SetTransparent sets the value of the transparent property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/isTransparent
func (b_ Box) SetTransparent(value bool) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setTransparent:"), value)
}
// The receiver’s title. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/title
func (b_ Box) Title() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("title"))
	return rv
}
// SetTitle sets the value of the title property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/title
func (b_ Box) SetTitle(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setTitle:"), value)
}
// The cell used to display the receiver’s title. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/titleCell
func (b_ Box) TitleCell() objc.ID {
	rv := objc.Send[objc.ID](b_.ID(), objc.RegisterName("titleCell"))
	return rv
}
// The font object used to draw the receiver’s title. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/titleFont
func (b_ Box) TitleFont() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("titleFont"))
	return rv
}
// SetTitleFont sets the value of the titleFont property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/titleFont
func (b_ Box) SetTitleFont(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setTitleFont:"), value)
}
// A constant representing the title position. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/titlePosition-swift.property
func (b_ Box) TitlePosition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID(), objc.RegisterName("titlePosition"))
	return rv
}
// SetTitlePosition sets the value of the titlePosition property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/titlePosition-swift.property
func (b_ Box) SetTitlePosition(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID(), objc.RegisterName("setTitlePosition:"), value)
}
// The rectangle in which the receiver’s title is drawn. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSBox/titleRect
func (b_ Box) TitleRect() foundation.Rect {
	rv := objc.Send[foundation.Rect](b_.ID(), objc.RegisterName("titleRect"))
	return rv
}
