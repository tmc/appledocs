// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	

	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	BezelColor() IColor
	SetBezelColor(value IColor)
	CustomizationLabel() foundation.foundation.INSString
	SetCustomizationLabel(value foundation.foundation.INSString)
	Image() IImage
	SetImage(value IImage)
	Enabled() bool
	SetEnabled(value bool)
	Target() objc.ID
	SetTarget(value objc.ID)
	Title() foundation.foundation.INSString
	SetTitle(value foundation.foundation.INSString)
	IsEnabled() bool
	SetIsEnabled(value bool)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (bc _ButtonTouchBarItemClass) Alloc() ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





// A bar item that provides a button.


// A bar item that provides a button.
//
// [Full Topic]
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






// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:image:target:action:)
func NewButtonTouchBarItemWithIdentifierImageTargetAction(identifier TouchBarItemIdentifier, image IImage, target objectivec.IObject, action objc.SEL) ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(getButtonTouchBarItemClass().class), objc.Sel("buttonTouchBarItemWithIdentifier:image:target:action:"), identifier, image, target, action)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:title:image:target:action:)
func NewButtonTouchBarItemWithIdentifierTitleImageTargetAction(identifier TouchBarItemIdentifier, title foundation.foundation.INSString, image IImage, target objectivec.IObject, action objc.SEL) ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(getButtonTouchBarItemClass().class), objc.Sel("buttonTouchBarItemWithIdentifier:title:image:target:action:"), identifier, title, image, target, action)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:title:target:action:)
func NewButtonTouchBarItemWithIdentifierTitleTargetAction(identifier TouchBarItemIdentifier, title foundation.foundation.INSString, target objectivec.IObject, action objc.SEL) ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(getButtonTouchBarItemClass().class), objc.Sel("buttonTouchBarItemWithIdentifier:title:target:action:"), identifier, title, target, action)
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:image:target:action:)
func (bc _ButtonTouchBarItemClass) ButtonTouchBarItemWithIdentifierImageTargetAction(identifier TouchBarItemIdentifier, image IImage, target objectivec.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("buttonTouchBarItemWithIdentifier:image:target:action:"), identifier, image, target, action)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:title:image:target:action:)
func (bc _ButtonTouchBarItemClass) ButtonTouchBarItemWithIdentifierTitleImageTargetAction(identifier TouchBarItemIdentifier, title foundation.foundation.INSString, image IImage, target objectivec.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("buttonTouchBarItemWithIdentifier:title:image:target:action:"), identifier, title, image, target, action)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:title:target:action:)
func (bc _ButtonTouchBarItemClass) ButtonTouchBarItemWithIdentifierTitleTargetAction(identifier TouchBarItemIdentifier, title foundation.foundation.INSString, target objectivec.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("buttonTouchBarItemWithIdentifier:title:target:action:"), identifier, title, target, action)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/action
func (b_ ButtonTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](b_.ID, objc.Sel("action"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/action
func (b_ ButtonTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/bezelColor
func (b_ ButtonTouchBarItem) BezelColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("bezelColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/bezelColor
func (b_ ButtonTouchBarItem) SetBezelColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/customizationLabel
func (b_ ButtonTouchBarItem) CustomizationLabel() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("customizationLabel"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/customizationLabel
func (b_ ButtonTouchBarItem) SetCustomizationLabel(value foundation.foundation.INSString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCustomizationLabel:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/image
func (b_ ButtonTouchBarItem) Image() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("image"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/image
func (b_ ButtonTouchBarItem) SetImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/isEnabled
func (b_ ButtonTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("enabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/isEnabled
func (b_ ButtonTouchBarItem) SetEnabled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/target
func (b_ ButtonTouchBarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("target"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/target
func (b_ ButtonTouchBarItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTarget:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/title
func (b_ ButtonTouchBarItem) Title() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("title"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/title
func (b_ ButtonTouchBarItem) SetTitle(value foundation.foundation.INSString) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/isenabled
func (b_ ButtonTouchBarItem) IsEnabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/isenabled
func (b_ ButtonTouchBarItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsEnabled:"), value)
}







