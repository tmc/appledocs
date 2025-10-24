// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtensionAction */


/* debug [class_header]: Header for WKWebExtensionAction */
// The class instance for the [WebExtensionAction] class.
var (
	WebExtensionActionClass     _WebExtensionActionClass
	WebExtensionActionClassOnce sync.Once
)

func getWebExtensionActionClass() _WebExtensionActionClass {
	WebExtensionActionClassOnce.Do(func() {
		WebExtensionActionClass = _WebExtensionActionClass{objc.GetClass("WKWebExtensionAction")}
	})
	return WebExtensionActionClass
}

type _WebExtensionActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebExtensionAction */
// An interface definition for the [WebExtensionAction] class.
type IWebExtensionAction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebExtensionAction */
	// properties:
	AssociatedTab() unsafe.Pointer
	BadgeText() objc.IObject /* cross-framework: NSString */
	HasUnreadBadgeText() bool
	SetHasUnreadBadgeText(value bool)
	InspectionName() objc.IObject /* cross-framework: NSString */
	SetInspectionName(value objc.IObject /* cross-framework: NSString */)
	Enabled() bool
	Label() objc.IObject /* cross-framework: NSString */
	MenuItems() []MenuElement /* not a class type */
	PopupPopover() appkit.Popover
	PopupWebView() IWKWebView
	PresentsPopup() bool
	WebExtensionContext() IWKWebExtensionContext
	IsEnabled() bool
	SetIsEnabled(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebExtensionAction */
	// methods:
	ClosePopup()
	IconForSize(size corefoundation.CGSize) appkit.Image
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebExtensionAction */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionActionClass) Alloc() WebExtensionAction {
	rv := objc.Send[WebExtensionAction](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebExtensionActionClass) New() WebExtensionAction {
	rv := objc.Send[WebExtensionAction](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionAction) Init() WebExtensionAction {
	rv := objc.Send[WebExtensionAction](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionAction) Autorelease() WebExtensionAction {
	rv := objc.Send[WebExtensionAction](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionAction creates a new WebExtensionAction instance.
func NewWebExtensionAction() WebExtensionAction {
	return getWebExtensionActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebExtensionAction */
// An object that encapsulates the properties for an individual web extension action.
//
// This class provides access to action properties, such as pop-up, icon, or title, with tab-specific values.


// An object that encapsulates the properties for an individual web extension action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action
type WebExtensionAction struct {
	objectivec.Object
}

// WebExtensionActionFrom constructs a [WebExtensionAction] from an unsafe.Pointer.
//
// An object that encapsulates the properties for an individual web extension action.
func WebExtensionActionFrom(ptr unsafe.Pointer) WebExtensionAction {
	return WebExtensionAction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebExtensionAction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebExtensionAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebExtensionAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebExtensionAction */

// Triggers the dismissal process of the pop-up.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/closePopup()
func (w_ WebExtensionAction) ClosePopup() {
	objc.Send[objc.ID](w_.ID, objc.Sel("closePopup"))
}/* debug [instance_methods/method]: ClosePopup */


// Returns the action icon for the specified size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/icon(for:)
func (w_ WebExtensionAction) IconForSize(size corefoundation.CGSize) appkit.Image {
	rv := objc.Send[appkit.Image](w_.ID, objc.Sel("iconForSize:"), size)
	return rv
}/* debug [instance_methods/method]: IconForSize */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebExtensionAction */

// The tab that this action is associated with, or if it’s the default action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/associatedTab
func (w_ WebExtensionAction) AssociatedTab() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("associatedTab"))
	return rv
}/* debug [instance_properties/getter]: associatedTab */


// The badge text for the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/badgeText
func (w_ WebExtensionAction) BadgeText() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("badgeText"))
	return rv
}/* debug [instance_properties/getter]: badgeText */


// A Boolean value indicating whether the badge text is unread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/hasUnreadBadgeText
func (w_ WebExtensionAction) HasUnreadBadgeText() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("hasUnreadBadgeText"))
	return rv
}/* debug [instance_properties/getter]: hasUnreadBadgeText */


// A Boolean value indicating whether the badge text is unread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/hasUnreadBadgeText
func (w_ WebExtensionAction) SetHasUnreadBadgeText(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setHasUnreadBadgeText:"), value)
}/* debug [instance_properties/setter]: hasUnreadBadgeText */


// The name shown when inspecting the pop-up web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/inspectionName
func (w_ WebExtensionAction) InspectionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("inspectionName"))
	return rv
}/* debug [instance_properties/getter]: inspectionName */


// The name shown when inspecting the pop-up web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/inspectionName
func (w_ WebExtensionAction) SetInspectionName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setInspectionName:"), value)
}/* debug [instance_properties/setter]: inspectionName */


// A Boolean value indicating whether the action is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/isEnabled
func (w_ WebExtensionAction) Enabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("enabled"))
	return rv
}/* debug [instance_properties/getter]: enabled */


// The localized display label for the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/label
func (w_ WebExtensionAction) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](w_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// The menu items provided by the extension for this action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/menuItems
func (w_ WebExtensionAction) MenuItems() []MenuElement /* not a class type */ {
	rv := objc.Send[[]MenuElement](w_.ID, objc.Sel("menuItems"))
	return rv
}/* debug [instance_properties/getter]: menuItems */


// A popover that presents a web view loaded with the pop-up page for this action, or if no popup is specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/popupPopover
func (w_ WebExtensionAction) PopupPopover() appkit.Popover {
	rv := objc.Send[appkit.Popover](w_.ID, objc.Sel("popupPopover"))
	return rv
}/* debug [instance_properties/getter]: popupPopover */


// A web view loaded with the pop-up page for this action, or if no pop-up is specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/popupWebView
func (w_ WebExtensionAction) PopupWebView() IWKWebView {
	rv := objc.Send[WebView](w_.ID, objc.Sel("popupWebView"))
	return rv
}/* debug [instance_properties/getter]: popupWebView */


// A Boolean value indicating whether the action has a pop-up.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/presentsPopup
func (w_ WebExtensionAction) PresentsPopup() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("presentsPopup"))
	return rv
}/* debug [instance_properties/getter]: presentsPopup */


// The extension context to which this action is related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/webExtensionContext
func (w_ WebExtensionAction) WebExtensionContext() IWKWebExtensionContext {
	rv := objc.Send[WebExtensionContext](w_.ID, objc.Sel("webExtensionContext"))
	return rv
}/* debug [instance_properties/getter]: webExtensionContext */


// A Boolean value indicating whether the action is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/action/isenabled
func (w_ WebExtensionAction) IsEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isEnabled"))
	return rv
}/* debug [instance_properties/getter]: isEnabled */


// A Boolean value indicating whether the action is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebextension/action/isenabled
func (w_ WebExtensionAction) SetIsEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsEnabled:"), value)
}/* debug [instance_properties/setter]: isEnabled */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebExtensionAction */


