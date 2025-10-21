// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [WebpagePreferences] class.
var (
	WebpagePreferencesClass     _WebpagePreferencesClass
	WebpagePreferencesClassOnce sync.Once
)

func getWebpagePreferencesClass() _WebpagePreferencesClass {
	WebpagePreferencesClassOnce.Do(func() {
		WebpagePreferencesClass = _WebpagePreferencesClass{objc.GetClass("WKWebpagePreferences")}
	})
	return WebpagePreferencesClass
}

type _WebpagePreferencesClass struct {
	class objc.Class
}

// An interface definition for the [WebpagePreferences] class.
type IWebpagePreferences interface {
	objectivec.IObject
}

// An object that specifies the behaviors to use when loading and rendering page content.
//
// Create a object when you want to change the default rendering behavior of your web view. Typically, iOS devices render web content for a mobile experience, and Mac devices render content for a desktop experience.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences
type WebpagePreferences struct {
	objectivec.Object
}

// WebpagePreferencesFrom constructs a [WebpagePreferences] from an unsafe.Pointer.
//
// An object that specifies the behaviors to use when loading and rendering page content.
func WebpagePreferencesFrom(ptr unsafe.Pointer) WebpagePreferences {
	return WebpagePreferences{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (wc _WebpagePreferencesClass) Alloc() WebpagePreferences {
	rv := objc.Send[WebpagePreferences](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (wc _WebpagePreferencesClass) New() WebpagePreferences {
	rv := objc.Send[WebpagePreferences](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebpagePreferences) Init() WebpagePreferences {
	rv := objc.Send[WebpagePreferences](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebpagePreferences) Autorelease() WebpagePreferences {
	rv := objc.Send[WebpagePreferences](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebpagePreferences creates a new WebpagePreferences instance.
func NewWebpagePreferences() WebpagePreferences {
	return getWebpagePreferencesClass().New()
}


// A Boolean value that indicates whether to use Lockdown Mode in the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebpagepreferences/islockdownmodeenabled
func (w_ WebpagePreferences) IsLockdownModeEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("isLockdownModeEnabled"))
	return rv
}


// SetIsLockdownModeEnabled sets the value of the isLockdownModeEnabled property.
// A Boolean value that indicates whether to use Lockdown Mode in the web view.

//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkwebpagepreferences/islockdownmodeenabled
func (w_ WebpagePreferences) SetIsLockdownModeEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setIsLockdownModeEnabled:"), value)
}

// A Boolean value that indicates whether JavaScript from web content is allowed to run.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/allowsContentJavaScript
func (w_ WebpagePreferences) AllowsContentJavaScript() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("allowsContentJavaScript"))
	return rv
}


// SetAllowsContentJavaScript sets the value of the allowsContentJavaScript property.
// A Boolean value that indicates whether JavaScript from web content is allowed to run.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/allowsContentJavaScript
func (w_ WebpagePreferences) SetAllowsContentJavaScript(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setAllowsContentJavaScript:"), value)
}

// A Boolean value that indicates whether to use Lockdown Mode in the web view.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/isLockdownModeEnabled
func (w_ WebpagePreferences) LockdownModeEnabled() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("lockdownModeEnabled"))
	return rv
}


// SetLockdownModeEnabled sets the value of the lockdownModeEnabled property.
// A Boolean value that indicates whether to use Lockdown Mode in the web view.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/isLockdownModeEnabled
func (w_ WebpagePreferences) SetLockdownModeEnabled(value bool) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setLockdownModeEnabled:"), value)
}

// The content mode for the web view to use when it loads and renders a webpage.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/preferredContentMode
func (w_ WebpagePreferences) PreferredContentMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("preferredContentMode"))
	return rv
}


// SetPreferredContentMode sets the value of the preferredContentMode property.
// The content mode for the web view to use when it loads and renders a webpage.

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/preferredContentMode
func (w_ WebpagePreferences) SetPreferredContentMode(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferredContentMode:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/preferredHTTPSNavigationPolicy
func (w_ WebpagePreferences) PreferredHTTPSNavigationPolicy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("preferredHTTPSNavigationPolicy"))
	return rv
}


// SetPreferredHTTPSNavigationPolicy sets the value of the preferredHTTPSNavigationPolicy property.
//
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebpagePreferences/preferredHTTPSNavigationPolicy
func (w_ WebpagePreferences) SetPreferredHTTPSNavigationPolicy(value unsafe.Pointer) {
	objc.Send[objc.ID](w_.ID, objc.Sel("setPreferredHTTPSNavigationPolicy:"), value)
}



