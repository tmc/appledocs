// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ButtonTouchBarItem] class.
var (
	ButtonTouchBarItemClass     _ButtonTouchBarItemClass
	ButtonTouchBarItemClassOnce sync.Once
)

func getButtonTouchBarItemClass() _ButtonTouchBarItemClass {
	ButtonTouchBarItemClassOnce.Do(func() {
		ButtonTouchBarItemClass = _ButtonTouchBarItemClass{objc.GetClass("NSButtonTouchBarItem")}
	})
	return ButtonTouchBarItemClass
}

type _ButtonTouchBarItemClass struct {
	class objc.Class
}

// An interface definition for the [ButtonTouchBarItem] class.
type IButtonTouchBarItem interface {
	ITouchBarItem
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	BezelColor() Color
	SetBezelColor(value IColor)
	CustomizationLabel() string
	SetCustomizationLabel(value string)
	Image() Image
	SetImage(value IImage)
	IsEnabled() bool
	SetIsEnabled(value bool)
	Target() unsafe.Pointer
	SetTarget(value unsafe.Pointer)
	Title() string
	SetTitle(value string)
}

// A bar item that provides a button.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem
type ButtonTouchBarItem struct {
	TouchBarItem
}

// ButtonTouchBarItemFrom constructs a [ButtonTouchBarItem] from an unsafe.Pointer.
//
// A bar item that provides a button.
func ButtonTouchBarItemFrom(ptr unsafe.Pointer) ButtonTouchBarItem {
	return ButtonTouchBarItem{
		TouchBarItem: TouchBarItemFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _ButtonTouchBarItemClass) Alloc() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _ButtonTouchBarItemClass) New() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ ButtonTouchBarItem) Init() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ ButtonTouchBarItem) Autorelease() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewButtonTouchBarItem creates a new ButtonTouchBarItem instance.
func NewButtonTouchBarItem() ButtonTouchBarItem {
	return getButtonTouchBarItemClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/action
func (b_ ButtonTouchBarItem) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("action"))
	return rv
}


// SetAction sets the value of the action property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/action
func (b_ ButtonTouchBarItem) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAction:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/bezelcolor
func (b_ ButtonTouchBarItem) BezelColor() Color {
	rv := objc.Send[Color](b_.ID, objc.Sel("bezelColor"))
	return rv
}


// SetBezelColor sets the value of the bezelColor property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/bezelcolor
func (b_ ButtonTouchBarItem) SetBezelColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelColor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/customizationlabel
func (b_ ButtonTouchBarItem) CustomizationLabel() string {
	rv := objc.Send[string](b_.ID, objc.Sel("customizationLabel"))
	return rv
}


// SetCustomizationLabel sets the value of the customizationLabel property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/customizationlabel
func (b_ ButtonTouchBarItem) SetCustomizationLabel(value string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCustomizationLabel:"), objc.String(value))
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/image
func (b_ ButtonTouchBarItem) Image() Image {
	rv := objc.Send[Image](b_.ID, objc.Sel("image"))
	return rv
}


// SetImage sets the value of the image property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/image
func (b_ ButtonTouchBarItem) SetImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImage:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/isenabled
func (b_ ButtonTouchBarItem) IsEnabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isEnabled"))
	return rv
}


// SetIsEnabled sets the value of the isEnabled property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/isenabled
func (b_ ButtonTouchBarItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsEnabled:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/target
func (b_ ButtonTouchBarItem) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b_.ID, objc.Sel("target"))
	return rv
}


// SetTarget sets the value of the target property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/target
func (b_ ButtonTouchBarItem) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTarget:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/title
func (b_ ButtonTouchBarItem) Title() string {
	rv := objc.Send[string](b_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/title
func (b_ ButtonTouchBarItem) SetTitle(value string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), objc.String(value))
}



