// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKPreferences */


/* debug [class_header]: Header for WKPreferences */
// The class instance for the [Preferences] class.
var (
	PreferencesClass     _PreferencesClass
	PreferencesClassOnce sync.Once
)

func getPreferencesClass() _PreferencesClass {
	PreferencesClassOnce.Do(func() {
		PreferencesClass = _PreferencesClass{objc.GetClass("WKPreferences")}
	})
	return PreferencesClass
}

type _PreferencesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Preferences */
// An interface definition for the [Preferences] class.
type IPreferences interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Preferences */
	// properties:
	InactiveSchedulingPolicy() InactiveSchedulingPolicy
	SetInactiveSchedulingPolicy(value InactiveSchedulingPolicy)
	ElementFullscreenEnabled() bool
	SetElementFullscreenEnabled(value bool)
	FraudulentWebsiteWarningEnabled() bool
	SetFraudulentWebsiteWarningEnabled(value bool)
	SiteSpecificQuirksModeEnabled() bool
	SetSiteSpecificQuirksModeEnabled(value bool)
	TextInteractionEnabled() bool
	SetTextInteractionEnabled(value bool)
	JavaEnabled() bool
	SetJavaEnabled(value bool)
	JavaScriptCanOpenWindowsAutomatically() bool
	SetJavaScriptCanOpenWindowsAutomatically(value bool)
	JavaScriptEnabled() bool
	SetJavaScriptEnabled(value bool)
	MinimumFontSize() float64
	SetMinimumFontSize(value float64)
	PlugInsEnabled() bool
	SetPlugInsEnabled(value bool)
	ShouldPrintBackgrounds() bool
	SetShouldPrintBackgrounds(value bool)
	TabFocusesLinks() bool
	SetTabFocusesLinks(value bool)
	IsElementFullscreenEnabled() bool
	SetIsElementFullscreenEnabled(value bool)
	IsFraudulentWebsiteWarningEnabled() bool
	SetIsFraudulentWebsiteWarningEnabled(value bool)
	IsSiteSpecificQuirksModeEnabled() bool
	SetIsSiteSpecificQuirksModeEnabled(value bool)
	IsTextInteractionEnabled() bool
	SetIsTextInteractionEnabled(value bool)
	Preferences() IWKPreferences
	SetPreferences(value IWKPreferences)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Preferences */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Preferences */
// Alloc allocates a new instance without initialization.
func (pc _PreferencesClass) Alloc() Preferences {
	rv := objc.Send[Preferences](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreferencesClass) New() Preferences {
	rv := objc.Send[Preferences](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Preferences) Init() Preferences {
	rv := objc.Send[Preferences](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Preferences) Autorelease() Preferences {
	rv := objc.Send[Preferences](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreferences creates a new Preferences instance.
func NewPreferences() Preferences {
	return getPreferencesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Preferences */
// An object that encapsulates the standard behaviors to apply to websites.
//
// Use a object to specify the preferences for your website, including the minimum font size, the JavaScript behavior, and the behavior for handling fraudulent websites. Create this object and assign it to the property of the object you use to create your web view.


// An object that encapsulates the standard behaviors to apply to websites.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences
type Preferences struct {
	objectivec.Object
}

// PreferencesFrom constructs a [Preferences] from an unsafe.Pointer.
//
// An object that encapsulates the standard behaviors to apply to websites.
func PreferencesFrom(ptr unsafe.Pointer) Preferences {
	return Preferences{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Preferences *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Preferences */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Preferences */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Preferences */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Preferences */

// A policy you set to specify how a web view that’s not in a window handles tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/inactiveSchedulingPolicy-swift.property
func (p_ Preferences) InactiveSchedulingPolicy() InactiveSchedulingPolicy {
	rv := objc.Send[InactiveSchedulingPolicy](p_.ID, objc.Sel("inactiveSchedulingPolicy"))
	return rv
}/* debug [instance_properties/getter]: inactiveSchedulingPolicy */


// A policy you set to specify how a web view that’s not in a window handles tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/inactiveSchedulingPolicy-swift.property
func (p_ Preferences) SetInactiveSchedulingPolicy(value InactiveSchedulingPolicy) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInactiveSchedulingPolicy:"), value)
}/* debug [instance_properties/setter]: inactiveSchedulingPolicy */


// A Boolean value that indicates whether a web view can display content full screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isElementFullscreenEnabled
func (p_ Preferences) ElementFullscreenEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("elementFullscreenEnabled"))
	return rv
}/* debug [instance_properties/getter]: elementFullscreenEnabled */


// A Boolean value that indicates whether a web view can display content full screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isElementFullscreenEnabled
func (p_ Preferences) SetElementFullscreenEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setElementFullscreenEnabled:"), value)
}/* debug [instance_properties/setter]: elementFullscreenEnabled */


// A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isFraudulentWebsiteWarningEnabled
func (p_ Preferences) FraudulentWebsiteWarningEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("fraudulentWebsiteWarningEnabled"))
	return rv
}/* debug [instance_properties/getter]: fraudulentWebsiteWarningEnabled */


// A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isFraudulentWebsiteWarningEnabled
func (p_ Preferences) SetFraudulentWebsiteWarningEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFraudulentWebsiteWarningEnabled:"), value)
}/* debug [instance_properties/setter]: fraudulentWebsiteWarningEnabled */


// A Boolean that indicates whether to apply site-specific compatibility workarounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isSiteSpecificQuirksModeEnabled
func (p_ Preferences) SiteSpecificQuirksModeEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("siteSpecificQuirksModeEnabled"))
	return rv
}/* debug [instance_properties/getter]: siteSpecificQuirksModeEnabled */


// A Boolean that indicates whether to apply site-specific compatibility workarounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isSiteSpecificQuirksModeEnabled
func (p_ Preferences) SetSiteSpecificQuirksModeEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSiteSpecificQuirksModeEnabled:"), value)
}/* debug [instance_properties/setter]: siteSpecificQuirksModeEnabled */


// A Boolean value that indicates whether to allow people to select or otherwise interact with text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isTextInteractionEnabled
func (p_ Preferences) TextInteractionEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("textInteractionEnabled"))
	return rv
}/* debug [instance_properties/getter]: textInteractionEnabled */


// A Boolean value that indicates whether to allow people to select or otherwise interact with text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isTextInteractionEnabled
func (p_ Preferences) SetTextInteractionEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTextInteractionEnabled:"), value)
}/* debug [instance_properties/setter]: textInteractionEnabled */


// A Boolean value that indicates whether Java is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/javaEnabled
func (p_ Preferences) JavaEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("javaEnabled"))
	return rv
}/* debug [instance_properties/getter]: javaEnabled */


// A Boolean value that indicates whether Java is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/javaEnabled
func (p_ Preferences) SetJavaEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJavaEnabled:"), value)
}/* debug [instance_properties/setter]: javaEnabled */


// A Boolean value that indicates whether JavaScript can open windows without user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/javaScriptCanOpenWindowsAutomatically
func (p_ Preferences) JavaScriptCanOpenWindowsAutomatically() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("javaScriptCanOpenWindowsAutomatically"))
	return rv
}/* debug [instance_properties/getter]: javaScriptCanOpenWindowsAutomatically */


// A Boolean value that indicates whether JavaScript can open windows without user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/javaScriptCanOpenWindowsAutomatically
func (p_ Preferences) SetJavaScriptCanOpenWindowsAutomatically(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJavaScriptCanOpenWindowsAutomatically:"), value)
}/* debug [instance_properties/setter]: javaScriptCanOpenWindowsAutomatically */


// A Boolean value that indicates whether JavaScript is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/javaScriptEnabled
func (p_ Preferences) JavaScriptEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("javaScriptEnabled"))
	return rv
}/* debug [instance_properties/getter]: javaScriptEnabled */


// A Boolean value that indicates whether JavaScript is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/javaScriptEnabled
func (p_ Preferences) SetJavaScriptEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJavaScriptEnabled:"), value)
}/* debug [instance_properties/setter]: javaScriptEnabled */


// The minimum font size, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/minimumFontSize
func (p_ Preferences) MinimumFontSize() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minimumFontSize"))
	return rv
}/* debug [instance_properties/getter]: minimumFontSize */


// The minimum font size, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/minimumFontSize
func (p_ Preferences) SetMinimumFontSize(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMinimumFontSize:"), value)
}/* debug [instance_properties/setter]: minimumFontSize */


// A Boolean value that indicates whether plug-ins are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/plugInsEnabled
func (p_ Preferences) PlugInsEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("plugInsEnabled"))
	return rv
}/* debug [instance_properties/getter]: plugInsEnabled */


// A Boolean value that indicates whether plug-ins are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/plugInsEnabled
func (p_ Preferences) SetPlugInsEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlugInsEnabled:"), value)
}/* debug [instance_properties/setter]: plugInsEnabled */


// A Boolean value that indicates whether to include any background color or graphics when printing content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/shouldPrintBackgrounds
func (p_ Preferences) ShouldPrintBackgrounds() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldPrintBackgrounds"))
	return rv
}/* debug [instance_properties/getter]: shouldPrintBackgrounds */


// A Boolean value that indicates whether to include any background color or graphics when printing content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/shouldPrintBackgrounds
func (p_ Preferences) SetShouldPrintBackgrounds(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldPrintBackgrounds:"), value)
}/* debug [instance_properties/setter]: shouldPrintBackgrounds */


// A Boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/tabFocusesLinks
func (p_ Preferences) TabFocusesLinks() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("tabFocusesLinks"))
	return rv
}/* debug [instance_properties/getter]: tabFocusesLinks */


// A Boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/tabFocusesLinks
func (p_ Preferences) SetTabFocusesLinks(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTabFocusesLinks:"), value)
}/* debug [instance_properties/setter]: tabFocusesLinks */


// A Boolean value that indicates whether a web view can display content full screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/iselementfullscreenenabled
func (p_ Preferences) IsElementFullscreenEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isElementFullscreenEnabled"))
	return rv
}/* debug [instance_properties/getter]: isElementFullscreenEnabled */


// A Boolean value that indicates whether a web view can display content full screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/iselementfullscreenenabled
func (p_ Preferences) SetIsElementFullscreenEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsElementFullscreenEnabled:"), value)
}/* debug [instance_properties/setter]: isElementFullscreenEnabled */


// A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/isfraudulentwebsitewarningenabled
func (p_ Preferences) IsFraudulentWebsiteWarningEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFraudulentWebsiteWarningEnabled"))
	return rv
}/* debug [instance_properties/getter]: isFraudulentWebsiteWarningEnabled */


// A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/isfraudulentwebsitewarningenabled
func (p_ Preferences) SetIsFraudulentWebsiteWarningEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFraudulentWebsiteWarningEnabled:"), value)
}/* debug [instance_properties/setter]: isFraudulentWebsiteWarningEnabled */


// A Boolean that indicates whether to apply site-specific compatibility workarounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/issitespecificquirksmodeenabled
func (p_ Preferences) IsSiteSpecificQuirksModeEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSiteSpecificQuirksModeEnabled"))
	return rv
}/* debug [instance_properties/getter]: isSiteSpecificQuirksModeEnabled */


// A Boolean that indicates whether to apply site-specific compatibility workarounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/issitespecificquirksmodeenabled
func (p_ Preferences) SetIsSiteSpecificQuirksModeEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSiteSpecificQuirksModeEnabled:"), value)
}/* debug [instance_properties/setter]: isSiteSpecificQuirksModeEnabled */


// A Boolean value that indicates whether to allow people to select or otherwise interact with text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/istextinteractionenabled
func (p_ Preferences) IsTextInteractionEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isTextInteractionEnabled"))
	return rv
}/* debug [instance_properties/getter]: isTextInteractionEnabled */


// A Boolean value that indicates whether to allow people to select or otherwise interact with text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/istextinteractionenabled
func (p_ Preferences) SetIsTextInteractionEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsTextInteractionEnabled:"), value)
}/* debug [instance_properties/setter]: isTextInteractionEnabled */


// The object that manages the preference-related settings for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/preferences
func (p_ Preferences) Preferences() IWKPreferences {
	rv := objc.Send[Preferences](p_.ID, objc.Sel("preferences"))
	return rv
}/* debug [instance_properties/getter]: preferences */


// The object that manages the preference-related settings for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/preferences
func (p_ Preferences) SetPreferences(value IWKPreferences) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferences:"), value)
}/* debug [instance_properties/setter]: preferences */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKPreferences */


