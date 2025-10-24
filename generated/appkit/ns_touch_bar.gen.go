// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TouchBar] class.
var (
	TouchBarClass     _TouchBarClass
	TouchBarClassOnce sync.Once
)

func getTouchBarClass() _TouchBarClass {
	TouchBarClassOnce.Do(func() {
		TouchBarClass = _TouchBarClass{objc.GetClass("NSTouchBar")}
	})
	return TouchBarClass
}

type _TouchBarClass struct {
	class objc.Class
}

// An interface definition for the [TouchBar] class.
type ITouchBar interface {
	objectivec.IObject
	// properties:
	IsAutomaticCustomizeTouchBarMenuItemEnabled() bool
	SetIsAutomaticCustomizeTouchBarMenuItemEnabled(value bool)
	BezelColor() objc.IObject /* cross-framework: Color */
	SetBezelColor(value objc.IObject /* cross-framework: Color */)
	AllowedTouchTypes() unsafe.Pointer
	SetAllowedTouchTypes(value unsafe.Pointer)
	GroupTouchBar() ITouchBar
	SetGroupTouchBar(value ITouchBar)
	PopoverTouchBar() ITouchBar
	SetPopoverTouchBar(value ITouchBar)
	PressAndHoldTouchBar() ITouchBar
	SetPressAndHoldTouchBar(value ITouchBar)
	SelectedSegmentBezelColor() objc.IObject /* cross-framework: Color */
	SetSelectedSegmentBezelColor(value objc.IObject /* cross-framework: Color */)
	TrackFillColor() objc.IObject /* cross-framework: Color */
	SetTrackFillColor(value objc.IObject /* cross-framework: Color */)
	CustomizationAllowedItemIdentifiers() unsafe.Pointer
	SetCustomizationAllowedItemIdentifiers(value unsafe.Pointer)
	CustomizationIdentifier() unsafe.Pointer
	SetCustomizationIdentifier(value unsafe.Pointer)
	CustomizationRequiredItemIdentifiers() unsafe.Pointer
	SetCustomizationRequiredItemIdentifiers(value unsafe.Pointer)
	DefaultItemIdentifiers() unsafe.Pointer
	SetDefaultItemIdentifiers(value unsafe.Pointer)
	Delegate() TouchBarDelegate /* not a class type */
	SetDelegate(value TouchBarDelegate /* not a class type */)
	EscapeKeyReplacementItemIdentifier() unsafe.Pointer
	SetEscapeKeyReplacementItemIdentifier(value unsafe.Pointer)
	IsVisible() bool
	SetIsVisible(value bool)
	ItemIdentifiers() unsafe.Pointer
	SetItemIdentifiers(value unsafe.Pointer)
	PrincipalItemIdentifier() unsafe.Pointer
	SetPrincipalItemIdentifier(value unsafe.Pointer)
	TemplateItems() ITouchBarItem
	SetTemplateItems(value ITouchBarItem)
	CustomizationLabel() objc.IObject /* cross-framework: NSString */
	SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */)
	TouchBar() ITouchBar
	SetTouchBar(value ITouchBar)
	AcceptsTouchEvents() bool
	SetAcceptsTouchEvents(value bool)
	// methods:
}

// An object that provides dynamic contextual controls in the Touch Bar of supported models of MacBook Pro.
//
// On supported MacBook Pro models, the Touch Bar, above the keyboard, shows instances of the class from the front-most app. Such an instance is called a . You define a bar to provide controls relevant to the user’s context. Each such control is an instance of the class, called an . You can provide many bars within your app, one for each responder instance; macOS frameworks can provide bars, as well, that can appear alongside your app’s bars. The system determines which bars to show at any given time. For example, an app that uses standard AppKit objects, such as text fields (instances of the class), obtains appropriate bars along with relevant items automatically. Refer to the following sample code projects, which demonstrate how to use and related classes, including the class, with its rich API that lets you build a highly customized picker control: To use the Touch Bar, define bars in objects in your app’s responder chain. At run time, the system traverses up the responder chain to discover, combine, and show bars from your app and from frameworks you link against. You can configure a bar to support dynamic composition, in which the system shows the bar in an expanded form that contains items from bars lower in the responder chain (from closer to the first responder). Because of dynamic composition and placement of items shown on the Touch Bar, always ensure that your bars appear as you expect them to, testing on the versions of macOS that you support. Instances of the class employ gesture recognizers and take advantage of macOS 10.12.1 event enhancements. Because of the physical geometry of the Touch Bar, touch events passed to gesture recognizers have only a meaningful , or horizontal, component. There’s no need, and no API, for your app to know whether or not there’s a Touch Bar available. Whether your app is running on a machine that supports the Touch Bar or not, your app’s onscreen user interface (UI) appears and behaves the same way. The Touch Bar is a Retina display, like the screen of a MacBook Pro. To perform custom drawing or animation within the Touch Bar, follow the same best practices that you would on the screen. On the right side of the Touch Bar, the system supplies the always-available . The Control Strip gives the user access to standard controls for display brightness, sound volume, Siri, and so on. Your app’s bars appear to the left of the Control Strip. The user can choose to hide the Control Strip, which gives the frontmost app the entire Touch Bar width. To the right of the Control Strip is a Touch ID sensor. To use Touch ID on supported MacBook Pro models, use methods from the framework. The Touch Bar dims automatically and wakes when the user touches it. Don’t show alerts in the Touch Bar, and don’t use the Touch Bar for widgets. For Touch Bar design guidance, read .


// An object that provides dynamic contextual controls in the Touch Bar of supported models of MacBook Pro.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar
type TouchBar struct {
	objectivec.Object
}

// TouchBarFrom constructs a [TouchBar] from an unsafe.Pointer.
//
// An object that provides dynamic contextual controls in the Touch Bar of supported models of MacBook Pro.
func TouchBarFrom(ptr unsafe.Pointer) TouchBar {
	return TouchBar{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _TouchBarClass) Alloc() TouchBar {
	rv := objc.Send[TouchBar](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TouchBarClass) New() TouchBar {
	rv := objc.Send[TouchBar](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TouchBar) Init() TouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TouchBar) Autorelease() TouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTouchBar creates a new TouchBar instance.
func NewTouchBar() TouchBar {
	return getTouchBarClass().New()
}



// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isautomaticcustomizetouchbarmenuitemenabled
func (t_ TouchBar) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}


// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isautomaticcustomizetouchbarmenuitemenabled
func (t_ TouchBar) SetIsAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}


// The color of the button’s bezel, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/bezelcolor
func (t_ TouchBar) BezelColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](t_.ID, objc.Sel("bezelColor"))
	return rv
}


// The color of the button’s bezel, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/bezelcolor
func (t_ TouchBar) SetBezelColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBezelColor:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/allowedtouchtypes
func (t_ TouchBar) AllowedTouchTypes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("allowedTouchTypes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/allowedtouchtypes
func (t_ TouchBar) SetAllowedTouchTypes(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedTouchTypes:"), value)
}


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBar) GroupTouchBar() ITouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("groupTouchBar"))
	return rv
}


// A bar that holds this group’s items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBar) SetGroupTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGroupTouchBar:"), value)
}


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBar) PopoverTouchBar() ITouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("popoverTouchBar"))
	return rv
}


// The bar displayed when this item is “popped.”
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBar) SetPopoverTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPopoverTouchBar:"), value)
}


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBar) PressAndHoldTouchBar() ITouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("pressAndHoldTouchBar"))
	return rv
}


// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBar) SetPressAndHoldTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}


// The color of the selected segment’s bezel, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/selectedsegmentbezelcolor
func (t_ TouchBar) SelectedSegmentBezelColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](t_.ID, objc.Sel("selectedSegmentBezelColor"))
	return rv
}


// The color of the selected segment’s bezel, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/selectedsegmentbezelcolor
func (t_ TouchBar) SetSelectedSegmentBezelColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedSegmentBezelColor:"), value)
}


// The color of the filled portion of the slider track, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/trackfillcolor
func (t_ TouchBar) TrackFillColor() objc.IObject /* cross-framework: Color */ {
	rv := objc.Send[Color](t_.ID, objc.Sel("trackFillColor"))
	return rv
}


// The color of the filled portion of the slider track, in appearances that support it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/trackfillcolor
func (t_ TouchBar) SetTrackFillColor(value objc.IObject /* cross-framework: Color */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTrackFillColor:"), value)
}


// A list of identifiers for items to show in the Touch Bar’s customization UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/customizationalloweditemidentifiers
func (t_ TouchBar) CustomizationAllowedItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("customizationAllowedItemIdentifiers"))
	return rv
}


// A list of identifiers for items to show in the Touch Bar’s customization UI.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/customizationalloweditemidentifiers
func (t_ TouchBar) SetCustomizationAllowedItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCustomizationAllowedItemIdentifiers:"), value)
}


// A globally unique string that makes the Touch Bar eligible for user customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/customizationidentifier-swift.property
func (t_ TouchBar) CustomizationIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("customizationIdentifier"))
	return rv
}


// A globally unique string that makes the Touch Bar eligible for user customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/customizationidentifier-swift.property
func (t_ TouchBar) SetCustomizationIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCustomizationIdentifier:"), value)
}


// An optional list of identifiers for items you want to always appear in the Touch Bar and which the user can’t remove during customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/customizationrequireditemidentifiers
func (t_ TouchBar) CustomizationRequiredItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("customizationRequiredItemIdentifiers"))
	return rv
}


// An optional list of identifiers for items you want to always appear in the Touch Bar and which the user can’t remove during customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/customizationrequireditemidentifiers
func (t_ TouchBar) SetCustomizationRequiredItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCustomizationRequiredItemIdentifiers:"), value)
}


// A required list of identifiers for items that you want to appear in the Touch Bar after instantiating it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/defaultitemidentifiers
func (t_ TouchBar) DefaultItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("defaultItemIdentifiers"))
	return rv
}


// A required list of identifiers for items that you want to appear in the Touch Bar after instantiating it.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/defaultitemidentifiers
func (t_ TouchBar) SetDefaultItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultItemIdentifiers:"), value)
}


// The delegate that provides items to the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/delegate
func (t_ TouchBar) Delegate() TouchBarDelegate /* not a class type */ {
	rv := objc.Send[TouchBarDelegate](t_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate that provides items to the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/delegate
func (t_ TouchBar) SetDelegate(value TouchBarDelegate /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}


// The identifier of an item that replaces the system-provided button in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/escapekeyreplacementitemidentifier
func (t_ TouchBar) EscapeKeyReplacementItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("escapeKeyReplacementItemIdentifier"))
	return rv
}


// The identifier of an item that replaces the system-provided button in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/escapekeyreplacementitemidentifier
func (t_ TouchBar) SetEscapeKeyReplacementItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEscapeKeyReplacementItemIdentifier:"), value)
}


// A Boolean value that Indicates whether the Touch Bar is eligible for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/isvisible
func (t_ TouchBar) IsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVisible"))
	return rv
}


// A Boolean value that Indicates whether the Touch Bar is eligible for display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/isvisible
func (t_ TouchBar) SetIsVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsVisible:"), value)
}


// The list of identifiers for the current items in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/itemidentifiers
func (t_ TouchBar) ItemIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("itemIdentifiers"))
	return rv
}


// The list of identifiers for the current items in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/itemidentifiers
func (t_ TouchBar) SetItemIdentifiers(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setItemIdentifiers:"), value)
}


// The identifier of an item you want the system to center in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t_ TouchBar) PrincipalItemIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("principalItemIdentifier"))
	return rv
}


// The identifier of an item you want the system to center in the Touch Bar.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/principalitemidentifier
func (t_ TouchBar) SetPrincipalItemIdentifier(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrincipalItemIdentifier:"), value)
}


// The primary source of items that the Touch Bar uses to fill its private items array, unless you provide items using a delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/templateitems
func (t_ TouchBar) TemplateItems() ITouchBarItem {
	rv := objc.Send[TouchBarItem](t_.ID, objc.Sel("templateItems"))
	return rv
}


// The primary source of items that the Touch Bar uses to fill its private items array, unless you provide items using a delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/templateitems
func (t_ TouchBar) SetTemplateItems(value ITouchBarItem) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTemplateItems:"), value)
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/customizationlabel
func (t_ TouchBar) CustomizationLabel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("customizationLabel"))
	return rv
}


// The user-visible string identifying this item during bar customization.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/customizationlabel
func (t_ TouchBar) SetCustomizationLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCustomizationLabel:"), value)
}


// The property you implement to provide a Touch Bar object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbarprovider/touchbar
func (t_ TouchBar) TouchBar() ITouchBar {
	rv := objc.Send[TouchBar](t_.ID, objc.Sel("touchBar"))
	return rv
}


// The property you implement to provide a Touch Bar object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbarprovider/touchbar
func (t_ TouchBar) SetTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTouchBar:"), value)
}


// A Boolean value indicating whether the view accepts touch events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/acceptstouchevents
func (t_ TouchBar) AcceptsTouchEvents() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsTouchEvents"))
	return rv
}


// A Boolean value indicating whether the view accepts touch events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/acceptstouchevents
func (t_ TouchBar) SetAcceptsTouchEvents(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptsTouchEvents:"), value)
}



