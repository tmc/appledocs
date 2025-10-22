// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ComboButton] class.
var (
	ComboButtonClass     _ComboButtonClass
	ComboButtonClassOnce sync.Once
)

func getComboButtonClass() _ComboButtonClass {
	ComboButtonClassOnce.Do(func() {
		ComboButtonClass = _ComboButtonClass{objc.GetClass("NSComboButton")}
	})
	return ComboButtonClass
}

type _ComboButtonClass struct {
	class objc.Class
}

// An interface definition for the [ComboButton] class.
type IComboButton interface {
	IControl
	Image() Image
	SetImage(value IImage)
	ImageScaling() ImageScaling
	SetImageScaling(value ImageScaling)
	Menu() NSMenu
	SetMenu(value IMenu)
	Style() unsafe.Pointer
	SetStyle(value unsafe.Pointer)
	Title() string
	SetTitle(value string)
	Action() unsafe.Pointer
	SetAction(value unsafe.Pointer)
	Target() unsafe.Pointer
	SetTarget(value unsafe.Pointer)
}

// A button with a pull-down menu and a default action.
//
// An object is a button that displays a title string, image, and an optional control for displaying a menu. Use this control in places where you want to offer a button with a default action and one or more alternative actions. Clicking the title or image executes the default action you provide, and clicking the menu control displays a menu for selecting a different action. If you configure the button to hide the menu control, a long-press gesture displays the menu. After you create a combo button programmatically or in Interface Builder, choose the button you want and add a title or image for your content. A combo button has a default action, which you specify at creation time. You can also change that action later using the inherited and properties. To specify one or more alternative actions, configure a menu with those actions and assign it to the button’s property. This control doesn’t use an object for its underlying implementation. It also doesn’t support the addition of a contextual menu.


// A button with a pull-down menu and a default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboButton

type ComboButton struct {
	Control
}

// ComboButtonFrom constructs a [ComboButton] from an unsafe.Pointer.
//
// A button with a pull-down menu and a default action.
func ComboButtonFrom(ptr unsafe.Pointer) ComboButton {
	return ComboButton{
		Control: ControlFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _ComboButtonClass) Alloc() ComboButton {
	rv := objc.Send[ComboButton](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ComboButtonClass) New() ComboButton {
	rv := objc.Send[ComboButton](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ComboButton) Init() ComboButton {
	rv := objc.Send[ComboButton](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ComboButton) Autorelease() ComboButton {
	rv := objc.Send[ComboButton](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewComboButton creates a new ComboButton instance.
func NewComboButton() ComboButton {
	return getComboButtonClass().New()
}



// The image that the button displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobutton/image

func (c_ ComboButton) Image() Image {
	rv := objc.Send[Image](c_.ID, objc.Sel("image"))
	return rv
}


// The image that the button displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobutton/image

func (c_ ComboButton) SetImage(value IImage) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImage:"), value)
}


// The scaling behavior to apply to the button’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobutton/imagescaling

func (c_ ComboButton) ImageScaling() ImageScaling {
	rv := objc.Send[ImageScaling](c_.ID, objc.Sel("imageScaling"))
	return rv
}


// The scaling behavior to apply to the button’s image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobutton/imagescaling

func (c_ ComboButton) SetImageScaling(value ImageScaling) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setImageScaling:"), value)
}


// The menu that contains the button’s alternate actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobutton/menu

func (c_ ComboButton) Menu() NSMenu {
	rv := objc.Send[NSMenu](c_.ID, objc.Sel("menu"))
	return rv
}


// The menu that contains the button’s alternate actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobutton/menu

func (c_ ComboButton) SetMenu(value IMenu) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMenu:"), value)
}


// The appearance setting that determines how the button presents its menu .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobutton/style-swift.property

func (c_ ComboButton) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("style"))
	return rv
}


// The appearance setting that determines how the button presents its menu .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobutton/style-swift.property

func (c_ ComboButton) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setStyle:"), value)
}


// The localized string that the button displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobutton/title

func (c_ ComboButton) Title() string {
	rv := objc.Send[string](c_.ID, objc.Sel("title"))
	return rv
}


// The localized string that the button displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscombobutton/title

func (c_ ComboButton) SetTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTitle:"), objc.String(value))
}


// The default action-message selector associated with the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/action

func (c_ ComboButton) Action() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("action"))
	return rv
}


// The default action-message selector associated with the control.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/action

func (c_ ComboButton) SetAction(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAction:"), value)
}


// The target object that receives action messages from the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/target

func (c_ ComboButton) Target() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("target"))
	return rv
}


// The target object that receives action messages from the cell.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscontrol/target

func (c_ ComboButton) SetTarget(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTarget:"), value)
}



