// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSButtonTouchBarItem */


/* debug [class_header]: Header for NSButtonTouchBarItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ButtonTouchBarItem */
// An interface definition for the [ButtonTouchBarItem] class.
type IButtonTouchBarItem interface {
	ITouchBarItem
	
/* debug [class_interface_properties]: Properties for ButtonTouchBarItem */
	// properties:
	Action() objc.SEL
	SetAction(value objc.SEL)
	BezelColor() IColor
	SetBezelColor(value IColor)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	Image() IImage
	SetImage(value IImage)
	Enabled() bool
	SetEnabled(value bool)
	Target() objc.ID
	SetTarget(value objc.ID)
	Title() objc.IObject /* cross-framework: NSString */
	SetTitle(value objc.IObject /* cross-framework: NSString */)
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ButtonTouchBarItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ButtonTouchBarItem */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ButtonTouchBarItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ButtonTouchBarItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:image:target:action:)
func NewButtonTouchBarItemWithIdentifierImageTargetAction(identifier TouchBarItemIdentifier /* typedef */, image IImage, target objc.IObject, action objc.SEL) ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(getButtonTouchBarItemClass().class), objc.Sel("buttonTouchBarItemWithIdentifier:image:target:action:"), identifier, image, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewButtonTouchBarItemWithIdentifierImageTargetAction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:title:image:target:action:)
func NewButtonTouchBarItemWithIdentifierTitleImageTargetAction(identifier TouchBarItemIdentifier /* typedef */, title objc.IObject /* cross-framework: NSString */, image IImage, target objc.IObject, action objc.SEL) ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(getButtonTouchBarItemClass().class), objc.Sel("buttonTouchBarItemWithIdentifier:title:image:target:action:"), identifier, title, image, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewButtonTouchBarItemWithIdentifierTitleImageTargetAction */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:title:target:action:)
func NewButtonTouchBarItemWithIdentifierTitleTargetAction(identifier TouchBarItemIdentifier /* typedef */, title objc.IObject /* cross-framework: NSString */, target objc.IObject, action objc.SEL) ButtonTouchBarItem {
	rv := objc.Send[ButtonTouchBarItem](objc.ID(getButtonTouchBarItemClass().class), objc.Sel("buttonTouchBarItemWithIdentifier:title:target:action:"), identifier, title, target, action)
	return rv
}/* debug [class_init_methods/constructor]: NewButtonTouchBarItemWithIdentifierTitleTargetAction */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ButtonTouchBarItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:image:target:action:)
func (bc _ButtonTouchBarItemClass) ButtonTouchBarItemWithIdentifierImageTargetAction(identifier TouchBarItemIdentifier /* typedef */, image IImage, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("buttonTouchBarItemWithIdentifier:image:target:action:"), identifier, image, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ButtonTouchBarItemWithIdentifierImageTargetAction) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:title:image:target:action:)
func (bc _ButtonTouchBarItemClass) ButtonTouchBarItemWithIdentifierTitleImageTargetAction(identifier TouchBarItemIdentifier /* typedef */, title objc.IObject /* cross-framework: NSString */, image IImage, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("buttonTouchBarItemWithIdentifier:title:image:target:action:"), identifier, title, image, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ButtonTouchBarItemWithIdentifierTitleImageTargetAction) */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/init(identifier:title:target:action:)
func (bc _ButtonTouchBarItemClass) ButtonTouchBarItemWithIdentifierTitleTargetAction(identifier TouchBarItemIdentifier /* typedef */, title objc.IObject /* cross-framework: NSString */, target objc.IObject, action objc.SEL) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(bc.class), objc.Sel("buttonTouchBarItemWithIdentifier:title:target:action:"), identifier, title, target, action)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ButtonTouchBarItemWithIdentifierTitleTargetAction) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ButtonTouchBarItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ButtonTouchBarItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ButtonTouchBarItem */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/action
func (b_ ButtonTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](b_.ID, objc.Sel("action"))
	return rv
}/* debug [instance_properties/getter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/action
func (b_ ButtonTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setAction:"), value)
}/* debug [instance_properties/setter]: action */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/bezelColor
func (b_ ButtonTouchBarItem) BezelColor() IColor {
	rv := objc.Send[Color](b_.ID, objc.Sel("bezelColor"))
	return rv
}/* debug [instance_properties/getter]: bezelColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/bezelColor
func (b_ ButtonTouchBarItem) SetBezelColor(value IColor) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setBezelColor:"), value)
}/* debug [instance_properties/setter]: bezelColor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/customizationLabel
func (b_ ButtonTouchBarItem) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("customizationLabel"))
	return rv
}/* debug [instance_properties/getter]: customizationLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/customizationLabel
func (b_ ButtonTouchBarItem) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setCustomizationLabel:"), value)
}/* debug [instance_properties/setter]: customizationLabel */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/image
func (b_ ButtonTouchBarItem) Image() IImage {
	rv := objc.Send[Image](b_.ID, objc.Sel("image"))
	return rv
}/* debug [instance_properties/getter]: image */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/image
func (b_ ButtonTouchBarItem) SetImage(value IImage) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setImage:"), value)
}/* debug [instance_properties/setter]: image */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/isEnabled
func (b_ ButtonTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/isEnabled
func (b_ ButtonTouchBarItem) SetEnabled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setEnabled:"), value)
}/* debug [instance_properties/setter]: enabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/target
func (b_ ButtonTouchBarItem) Target() objc.ID {
	rv := objc.Send[objc.ID](b_.ID, objc.Sel("target"))
	return rv
}/* debug [instance_properties/getter]: target */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/target
func (b_ ButtonTouchBarItem) SetTarget(value objc.ID) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTarget:"), value)
}/* debug [instance_properties/setter]: target */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/title
func (b_ ButtonTouchBarItem) Title() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](b_.ID, objc.Sel("title"))
	return rv
}/* debug [instance_properties/getter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButtonTouchBarItem/title
func (b_ ButtonTouchBarItem) SetTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitle:"), value)
}/* debug [instance_properties/setter]: title */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/isenabled
func (b_ ButtonTouchBarItem) IsEnabled() bool {
	rv := objc.Send[bool](b_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbuttontouchbaritem/isenabled
func (b_ ButtonTouchBarItem) SetIsEnabled(value bool) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSButtonTouchBarItem */


