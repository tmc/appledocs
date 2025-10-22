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
	ItemForIdentifier(identifier ITouchBarItemIdentifier) TouchBarItem
	CustomizationAllowedItemIdentifiers() []string
	SetCustomizationAllowedItemIdentifiers(value []string)
	CustomizationIdentifier() TouchBarCustomizationIdentifier
	SetCustomizationIdentifier(value ITouchBarCustomizationIdentifier)
	CustomizationRequiredItemIdentifiers() []string
	SetCustomizationRequiredItemIdentifiers(value []string)
	DefaultItemIdentifiers() []string
	SetDefaultItemIdentifiers(value []string)
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier
	SetEscapeKeyReplacementItemIdentifier(value ITouchBarItemIdentifier)
	Visible() bool
	ItemIdentifiers() []string
	PrincipalItemIdentifier() TouchBarItemIdentifier
	SetPrincipalItemIdentifier(value ITouchBarItemIdentifier)
	TemplateItems() unsafe.Pointer
	SetTemplateItems(value unsafe.Pointer)
	IsAutomaticCustomizeTouchBarMenuItemEnabled() bool
	SetIsAutomaticCustomizeTouchBarMenuItemEnabled(value bool)
	BezelColor() NSColor
	SetBezelColor(value IColor)
	AllowedTouchTypes() TouchTypeMask
	SetAllowedTouchTypes(value TouchTypeMask)
	GroupTouchBar() NSTouchBar
	SetGroupTouchBar(value ITouchBar)
	PopoverTouchBar() NSTouchBar
	SetPopoverTouchBar(value ITouchBar)
	PressAndHoldTouchBar() NSTouchBar
	SetPressAndHoldTouchBar(value ITouchBar)
	SelectedSegmentBezelColor() NSColor
	SetSelectedSegmentBezelColor(value IColor)
	TrackFillColor() NSColor
	SetTrackFillColor(value IColor)
	IsVisible() bool
	SetIsVisible(value bool)
	CustomizationLabel() string
	SetCustomizationLabel(value string)
	TouchBar() NSTouchBar
	SetTouchBar(value ITouchBar)
	AcceptsTouchEvents() bool
	SetAcceptsTouchEvents(value bool)
}

// An object that provides dynamic contextual controls in the Touch Bar of supported models of MacBook Pro.
//
// On supported MacBook Pro models, the Touch Bar, above the keyboard, shows instances of the class from the front-most app. Such an instance is called a . You define a bar to provide controls relevant to the user’s context. Each such control is an instance of the class, called an . You can provide many bars within your app, one for each responder instance; macOS frameworks can provide bars, as well, that can appear alongside your app’s bars. The system determines which bars to show at any given time. For example, an app that uses standard AppKit objects, such as text fields (instances of the class), obtains appropriate bars along with relevant items automatically. Refer to the following sample code projects, which demonstrate how to use and related classes, including the class, with its rich API that lets you build a highly customized picker control: To use the Touch Bar, define bars in objects in your app’s responder chain. At run time, the system traverses up the responder chain to discover, combine, and show bars from your app and from frameworks you link against. You can configure a bar to support dynamic composition, in which the system shows the bar in an expanded form that contains items from bars lower in the responder chain (from closer to the first responder). Because of dynamic composition and placement of items shown on the Touch Bar, always ensure that your bars appear as you expect them to, testing on the versions of macOS that you support. Instances of the class employ gesture recognizers and take advantage of macOS 10.12.1 event enhancements. Because of the physical geometry of the Touch Bar, touch events passed to gesture recognizers have only a meaningful , or horizontal, component. There’s no need, and no API, for your app to know whether or not there’s a Touch Bar available. Whether your app is running on a machine that supports the Touch Bar or not, your app’s onscreen user interface (UI) appears and behaves the same way. The Touch Bar is a Retina display, like the screen of a MacBook Pro. To perform custom drawing or animation within the Touch Bar, follow the same best practices that you would on the screen. On the right side of the Touch Bar, the system supplies the always-available . The Control Strip gives the user access to standard controls for display brightness, sound volume, Siri, and so on. Your app’s bars appear to the left of the Control Strip. The user can choose to hide the Control Strip, which gives the frontmost app the entire Touch Bar width. To the right of the Control Strip is a Touch ID sensor. To use Touch ID on supported MacBook Pro models, use methods from the framework. The Touch Bar dims automatically and wakes when the user touches it. Don’t show alerts in the Touch Bar, and don’t use the Touch Bar for widgets. For Touch Bar design guidance, read .
//
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




// Creates a Touch Bar object from a coder object provided by a storyboard or NIB file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/init(coder:)
func NewTouchBarWithCoder(coder foundation.ICoder) TouchBar {
	instance := getTouchBarClass().Alloc()
	rv := objc.Send[TouchBar](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/isAutomaticCustomizeTouchBarMenuItemEnabled
func (tc _TouchBarClass) AutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.Send[bool](objc.ID(tc.class), objc.Sel("automaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}
// Returns the Touch Bar item that corresponds to a given identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/item(forIdentifier:)
func (t_ TouchBar) ItemForIdentifier(identifier ITouchBarItemIdentifier) TouchBarItem {
	rv := objc.Send[TouchBarItem](t_.ID, objc.Sel("itemForIdentifier:"), identifier)
	return rv
}

// A list of identifiers for items to show in the Touch Bar’s customization UI.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/customizationAllowedItemIdentifiers
func (t_ TouchBar) CustomizationAllowedItemIdentifiers() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("customizationAllowedItemIdentifiers"))
	return rv
}


// SetCustomizationAllowedItemIdentifiers sets the value of the customizationAllowedItemIdentifiers property.
// A list of identifiers for items to show in the Touch Bar’s customization UI.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/customizationAllowedItemIdentifiers
func (t_ TouchBar) SetCustomizationAllowedItemIdentifiers(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setCustomizationAllowedItemIdentifiers:"), nsArray)
}

// A globally unique string that makes the Touch Bar eligible for user customization.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/customizationIdentifier-swift.property
func (t_ TouchBar) CustomizationIdentifier() TouchBarCustomizationIdentifier {
	rv := objc.Send[TouchBarCustomizationIdentifier](t_.ID, objc.Sel("customizationIdentifier"))
	return rv
}


// SetCustomizationIdentifier sets the value of the customizationIdentifier property.
// A globally unique string that makes the Touch Bar eligible for user customization.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/customizationIdentifier-swift.property
func (t_ TouchBar) SetCustomizationIdentifier(value ITouchBarCustomizationIdentifier) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCustomizationIdentifier:"), value)
}

// An optional list of identifiers for items you want to always appear in the Touch Bar and which the user can’t remove during customization.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/customizationRequiredItemIdentifiers
func (t_ TouchBar) CustomizationRequiredItemIdentifiers() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("customizationRequiredItemIdentifiers"))
	return rv
}


// SetCustomizationRequiredItemIdentifiers sets the value of the customizationRequiredItemIdentifiers property.
// An optional list of identifiers for items you want to always appear in the Touch Bar and which the user can’t remove during customization.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/customizationRequiredItemIdentifiers
func (t_ TouchBar) SetCustomizationRequiredItemIdentifiers(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setCustomizationRequiredItemIdentifiers:"), nsArray)
}

// A required list of identifiers for items that you want to appear in the Touch Bar after instantiating it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/defaultItemIdentifiers
func (t_ TouchBar) DefaultItemIdentifiers() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("defaultItemIdentifiers"))
	return rv
}


// SetDefaultItemIdentifiers sets the value of the defaultItemIdentifiers property.
// A required list of identifiers for items that you want to appear in the Touch Bar after instantiating it.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/defaultItemIdentifiers
func (t_ TouchBar) SetDefaultItemIdentifiers(value []string) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](t_.ID, objc.Sel("setDefaultItemIdentifiers:"), nsArray)
}

// The delegate that provides items to the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/delegate
func (t_ TouchBar) Delegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate that provides items to the Touch Bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/delegate
func (t_ TouchBar) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}

// The identifier of an item that replaces the system-provided button in the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/escapeKeyReplacementItemIdentifier
func (t_ TouchBar) EscapeKeyReplacementItemIdentifier() TouchBarItemIdentifier {
	rv := objc.Send[TouchBarItemIdentifier](t_.ID, objc.Sel("escapeKeyReplacementItemIdentifier"))
	return rv
}


// SetEscapeKeyReplacementItemIdentifier sets the value of the escapeKeyReplacementItemIdentifier property.
// The identifier of an item that replaces the system-provided button in the Touch Bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/escapeKeyReplacementItemIdentifier
func (t_ TouchBar) SetEscapeKeyReplacementItemIdentifier(value ITouchBarItemIdentifier) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setEscapeKeyReplacementItemIdentifier:"), value)
}

// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/isAutomaticCustomizeTouchBarMenuItemEnabled
func (t_ TouchBar) AutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("automaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}


// SetAutomaticCustomizeTouchBarMenuItemEnabled sets the value of the automaticCustomizeTouchBarMenuItemEnabled property.
// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/isAutomaticCustomizeTouchBarMenuItemEnabled
func (t_ TouchBar) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}

// A Boolean value that Indicates whether the Touch Bar is eligible for display.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/isVisible
func (t_ TouchBar) Visible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("visible"))
	return rv
}

// The list of identifiers for the current items in the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/itemIdentifiers
func (t_ TouchBar) ItemIdentifiers() []string {
	rv := objc.Send[[]string](t_.ID, objc.Sel("itemIdentifiers"))
	return rv
}

// The identifier of an item you want the system to center in the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/principalItemIdentifier
func (t_ TouchBar) PrincipalItemIdentifier() TouchBarItemIdentifier {
	rv := objc.Send[TouchBarItemIdentifier](t_.ID, objc.Sel("principalItemIdentifier"))
	return rv
}


// SetPrincipalItemIdentifier sets the value of the principalItemIdentifier property.
// The identifier of an item you want the system to center in the Touch Bar.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/principalItemIdentifier
func (t_ TouchBar) SetPrincipalItemIdentifier(value ITouchBarItemIdentifier) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPrincipalItemIdentifier:"), value)
}

// The primary source of items that the Touch Bar uses to fill its private items array, unless you provide items using a delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/templateItems
func (t_ TouchBar) TemplateItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("templateItems"))
	return rv
}


// SetTemplateItems sets the value of the templateItems property.
// The primary source of items that the Touch Bar uses to fill its private items array, unless you provide items using a delegate.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/templateItems
func (t_ TouchBar) SetTemplateItems(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTemplateItems:"), value)
}

// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isautomaticcustomizetouchbarmenuitemenabled
func (t_ TouchBar) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isAutomaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}


// SetIsAutomaticCustomizeTouchBarMenuItemEnabled sets the value of the isAutomaticCustomizeTouchBarMenuItemEnabled property.
// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsapplication/isautomaticcustomizetouchbarmenuitemenabled
func (t_ TouchBar) SetIsAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}

// The color of the button’s bezel, in appearances that support it.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/bezelcolor
func (t_ TouchBar) BezelColor() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("bezelColor"))
	return rv
}


// SetBezelColor sets the value of the bezelColor property.
// The color of the button’s bezel, in appearances that support it.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsbutton/bezelcolor
func (t_ TouchBar) SetBezelColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setBezelColor:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/allowedtouchtypes
func (t_ TouchBar) AllowedTouchTypes() TouchTypeMask {
	rv := objc.Send[TouchTypeMask](t_.ID, objc.Sel("allowedTouchTypes"))
	return rv
}


// SetAllowedTouchTypes sets the value of the allowedTouchTypes property.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgesturerecognizer/allowedtouchtypes
func (t_ TouchBar) SetAllowedTouchTypes(value TouchTypeMask) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAllowedTouchTypes:"), value)
}

// A bar that holds this group’s items.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBar) GroupTouchBar() NSTouchBar {
	rv := objc.Send[NSTouchBar](t_.ID, objc.Sel("groupTouchBar"))
	return rv
}


// SetGroupTouchBar sets the value of the groupTouchBar property.
// A bar that holds this group’s items.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsgrouptouchbaritem/grouptouchbar
func (t_ TouchBar) SetGroupTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setGroupTouchBar:"), value)
}

// The bar displayed when this item is “popped.”
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBar) PopoverTouchBar() NSTouchBar {
	rv := objc.Send[NSTouchBar](t_.ID, objc.Sel("popoverTouchBar"))
	return rv
}


// SetPopoverTouchBar sets the value of the popoverTouchBar property.
// The bar displayed when this item is “popped.”

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/popovertouchbar
func (t_ TouchBar) SetPopoverTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPopoverTouchBar:"), value)
}

// The bar that is displayed when a user press-and-holds on the popover item.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBar) PressAndHoldTouchBar() NSTouchBar {
	rv := objc.Send[NSTouchBar](t_.ID, objc.Sel("pressAndHoldTouchBar"))
	return rv
}


// SetPressAndHoldTouchBar sets the value of the pressAndHoldTouchBar property.
// The bar that is displayed when a user press-and-holds on the popover item.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspopovertouchbaritem/pressandholdtouchbar
func (t_ TouchBar) SetPressAndHoldTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setPressAndHoldTouchBar:"), value)
}

// The color of the selected segment’s bezel, in appearances that support it.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/selectedsegmentbezelcolor
func (t_ TouchBar) SelectedSegmentBezelColor() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("selectedSegmentBezelColor"))
	return rv
}


// SetSelectedSegmentBezelColor sets the value of the selectedSegmentBezelColor property.
// The color of the selected segment’s bezel, in appearances that support it.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nssegmentedcontrol/selectedsegmentbezelcolor
func (t_ TouchBar) SetSelectedSegmentBezelColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setSelectedSegmentBezelColor:"), value)
}

// The color of the filled portion of the slider track, in appearances that support it.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/trackfillcolor
func (t_ TouchBar) TrackFillColor() NSColor {
	rv := objc.Send[NSColor](t_.ID, objc.Sel("trackFillColor"))
	return rv
}


// SetTrackFillColor sets the value of the trackFillColor property.
// The color of the filled portion of the slider track, in appearances that support it.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsslider/trackfillcolor
func (t_ TouchBar) SetTrackFillColor(value IColor) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTrackFillColor:"), value)
}

// A Boolean value that Indicates whether the Touch Bar is eligible for display.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/isvisible
func (t_ TouchBar) IsVisible() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("isVisible"))
	return rv
}


// SetIsVisible sets the value of the isVisible property.
// A Boolean value that Indicates whether the Touch Bar is eligible for display.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbar/isvisible
func (t_ TouchBar) SetIsVisible(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setIsVisible:"), value)
}

// The user-visible string identifying this item during bar customization.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/customizationlabel
func (t_ TouchBar) CustomizationLabel() string {
	rv := objc.Send[string](t_.ID, objc.Sel("customizationLabel"))
	return rv
}


// SetCustomizationLabel sets the value of the customizationLabel property.
// The user-visible string identifying this item during bar customization.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbaritem/customizationlabel
func (t_ TouchBar) SetCustomizationLabel(value string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCustomizationLabel:"), objc.String(value))
}

// The property you implement to provide a Touch Bar object.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbarprovider/touchbar
func (t_ TouchBar) TouchBar() NSTouchBar {
	rv := objc.Send[NSTouchBar](t_.ID, objc.Sel("touchBar"))
	return rv
}


// SetTouchBar sets the value of the touchBar property.
// The property you implement to provide a Touch Bar object.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nstouchbarprovider/touchbar
func (t_ TouchBar) SetTouchBar(value ITouchBar) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTouchBar:"), value)
}

// A Boolean value indicating whether the view accepts touch events.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/acceptstouchevents
func (t_ TouchBar) AcceptsTouchEvents() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("acceptsTouchEvents"))
	return rv
}


// SetAcceptsTouchEvents sets the value of the acceptsTouchEvents property.
// A Boolean value indicating whether the view accepts touch events.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/acceptstouchevents
func (t_ TouchBar) SetAcceptsTouchEvents(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setAcceptsTouchEvents:"), value)
}


