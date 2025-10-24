// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Preferences] class.
type IPreferences interface {
	objectivec.IObject
	// properties:
	SiteSpecificQuirksModeEnabled() bool
	SetSiteSpecificQuirksModeEnabled(value bool)
	InactiveSchedulingPolicy() unsafe.Pointer
	SetInactiveSchedulingPolicy(value unsafe.Pointer)
	IsElementFullscreenEnabled() bool
	SetIsElementFullscreenEnabled(value bool)
	IsFraudulentWebsiteWarningEnabled() bool
	SetIsFraudulentWebsiteWarningEnabled(value bool)
	IsLookToScrollEnabled() bool
	SetIsLookToScrollEnabled(value bool)
	IsSiteSpecificQuirksModeEnabled() bool
	SetIsSiteSpecificQuirksModeEnabled(value bool)
	IsTextInteractionEnabled() bool
	SetIsTextInteractionEnabled(value bool)
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
	Preferences() IWKPreferences
	SetPreferences(value IWKPreferences)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (pc _PreferencesClass) Alloc() Preferences {
	rv := objc.Send[Preferences](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A Boolean that indicates whether to apply site-specific compatibility workarounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isSiteSpecificQuirksModeEnabled
func (p_ Preferences) SiteSpecificQuirksModeEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("siteSpecificQuirksModeEnabled"))
	return rv
}


// A Boolean that indicates whether to apply site-specific compatibility workarounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKPreferences/isSiteSpecificQuirksModeEnabled
func (p_ Preferences) SetSiteSpecificQuirksModeEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSiteSpecificQuirksModeEnabled:"), value)
}


// A policy you set to specify how a web view that’s not in a window handles tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/inactiveschedulingpolicy-swift.property
func (p_ Preferences) InactiveSchedulingPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("inactiveSchedulingPolicy"))
	return rv
}


// A policy you set to specify how a web view that’s not in a window handles tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/inactiveschedulingpolicy-swift.property
func (p_ Preferences) SetInactiveSchedulingPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInactiveSchedulingPolicy:"), value)
}


// A Boolean value that indicates whether a web view can display content full screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/iselementfullscreenenabled
func (p_ Preferences) IsElementFullscreenEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isElementFullscreenEnabled"))
	return rv
}


// A Boolean value that indicates whether a web view can display content full screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/iselementfullscreenenabled
func (p_ Preferences) SetIsElementFullscreenEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsElementFullscreenEnabled:"), value)
}


// A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/isfraudulentwebsitewarningenabled
func (p_ Preferences) IsFraudulentWebsiteWarningEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFraudulentWebsiteWarningEnabled"))
	return rv
}


// A Boolean value that indicates whether the web view shows warnings for suspected fraudulent content, such as malware or phishing attemps.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/isfraudulentwebsitewarningenabled
func (p_ Preferences) SetIsFraudulentWebsiteWarningEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFraudulentWebsiteWarningEnabled:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/islooktoscrollenabled
func (p_ Preferences) IsLookToScrollEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isLookToScrollEnabled"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/islooktoscrollenabled
func (p_ Preferences) SetIsLookToScrollEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsLookToScrollEnabled:"), value)
}


// A Boolean that indicates whether to apply site-specific compatibility workarounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/issitespecificquirksmodeenabled
func (p_ Preferences) IsSiteSpecificQuirksModeEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isSiteSpecificQuirksModeEnabled"))
	return rv
}


// A Boolean that indicates whether to apply site-specific compatibility workarounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/issitespecificquirksmodeenabled
func (p_ Preferences) SetIsSiteSpecificQuirksModeEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsSiteSpecificQuirksModeEnabled:"), value)
}


// A Boolean value that indicates whether to allow people to select or otherwise interact with text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/istextinteractionenabled
func (p_ Preferences) IsTextInteractionEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isTextInteractionEnabled"))
	return rv
}


// A Boolean value that indicates whether to allow people to select or otherwise interact with text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/istextinteractionenabled
func (p_ Preferences) SetIsTextInteractionEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsTextInteractionEnabled:"), value)
}


// A Boolean value that indicates whether Java is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/javaenabled
func (p_ Preferences) JavaEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("javaEnabled"))
	return rv
}


// A Boolean value that indicates whether Java is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/javaenabled
func (p_ Preferences) SetJavaEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJavaEnabled:"), value)
}


// A Boolean value that indicates whether JavaScript can open windows without user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/javascriptcanopenwindowsautomatically
func (p_ Preferences) JavaScriptCanOpenWindowsAutomatically() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("javaScriptCanOpenWindowsAutomatically"))
	return rv
}


// A Boolean value that indicates whether JavaScript can open windows without user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/javascriptcanopenwindowsautomatically
func (p_ Preferences) SetJavaScriptCanOpenWindowsAutomatically(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJavaScriptCanOpenWindowsAutomatically:"), value)
}


// A Boolean value that indicates whether JavaScript is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/javascriptenabled
func (p_ Preferences) JavaScriptEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("javaScriptEnabled"))
	return rv
}


// A Boolean value that indicates whether JavaScript is enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/javascriptenabled
func (p_ Preferences) SetJavaScriptEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJavaScriptEnabled:"), value)
}


// The minimum font size, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/minimumfontsize
func (p_ Preferences) MinimumFontSize() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("minimumFontSize"))
	return rv
}


// The minimum font size, in points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/minimumfontsize
func (p_ Preferences) SetMinimumFontSize(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMinimumFontSize:"), value)
}


// A Boolean value that indicates whether plug-ins are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/pluginsenabled
func (p_ Preferences) PlugInsEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("plugInsEnabled"))
	return rv
}


// A Boolean value that indicates whether plug-ins are enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/pluginsenabled
func (p_ Preferences) SetPlugInsEnabled(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlugInsEnabled:"), value)
}


// A Boolean value that indicates whether to include any background color or graphics when printing content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/shouldprintbackgrounds
func (p_ Preferences) ShouldPrintBackgrounds() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("shouldPrintBackgrounds"))
	return rv
}


// A Boolean value that indicates whether to include any background color or graphics when printing content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/shouldprintbackgrounds
func (p_ Preferences) SetShouldPrintBackgrounds(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setShouldPrintBackgrounds:"), value)
}


// A Boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/tabfocuseslinks
func (p_ Preferences) TabFocusesLinks() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("tabFocusesLinks"))
	return rv
}


// A Boolean value that indicates whether pressing the tab key changes the focus to links and form controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkpreferences/tabfocuseslinks
func (p_ Preferences) SetTabFocusesLinks(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTabFocusesLinks:"), value)
}


// The object that manages the preference-related settings for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/preferences
func (p_ Preferences) Preferences() IWKPreferences {
	rv := objc.Send[Preferences](p_.ID, objc.Sel("preferences"))
	return rv
}


// The object that manages the preference-related settings for the web view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebviewconfiguration/preferences
func (p_ Preferences) SetPreferences(value IWKPreferences) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPreferences:"), value)
}



