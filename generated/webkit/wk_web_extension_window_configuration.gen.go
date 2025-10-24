// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtensionWindowConfiguration */


/* debug [class_header]: Header for WKWebExtensionWindowConfiguration */
// The class instance for the [WebExtensionWindowConfiguration] class.
var (
	WebExtensionWindowConfigurationClass     _WebExtensionWindowConfigurationClass
	WebExtensionWindowConfigurationClassOnce sync.Once
)

func getWebExtensionWindowConfigurationClass() _WebExtensionWindowConfigurationClass {
	WebExtensionWindowConfigurationClassOnce.Do(func() {
		WebExtensionWindowConfigurationClass = _WebExtensionWindowConfigurationClass{objc.GetClass("WKWebExtensionWindowConfiguration")}
	})
	return WebExtensionWindowConfigurationClass
}

type _WebExtensionWindowConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebExtensionWindowConfiguration */
// An interface definition for the [WebExtensionWindowConfiguration] class.
type IWebExtensionWindowConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebExtensionWindowConfiguration */
	// properties:
	Frame() corefoundation.CGRect
	ShouldBeFocused() bool
	ShouldBePrivate() bool
	Tabs() []objc.ID
	TabURLs() []foundation.URL
	WindowState() WebExtensionWindowState
	WindowType() WebExtensionWindowType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebExtensionWindowConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebExtensionWindowConfiguration */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionWindowConfigurationClass) Alloc() WebExtensionWindowConfiguration {
	rv := objc.Send[WebExtensionWindowConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebExtensionWindowConfigurationClass) New() WebExtensionWindowConfiguration {
	rv := objc.Send[WebExtensionWindowConfiguration](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionWindowConfiguration) Init() WebExtensionWindowConfiguration {
	rv := objc.Send[WebExtensionWindowConfiguration](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionWindowConfiguration) Autorelease() WebExtensionWindowConfiguration {
	rv := objc.Send[WebExtensionWindowConfiguration](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionWindowConfiguration creates a new WebExtensionWindowConfiguration instance.
func NewWebExtensionWindowConfiguration() WebExtensionWindowConfiguration {
	return getWebExtensionWindowConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebExtensionWindowConfiguration */
// An object that encapsulates configuration options for a window in an extension.
//
// This class holds various options that influence the behavior and initial state of a window. The app retains the discretion to disregard any or all of these options, or even opt not to create a window.


// An object that encapsulates configuration options for a window in an extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration
type WebExtensionWindowConfiguration struct {
	objectivec.Object
}

// WebExtensionWindowConfigurationFrom constructs a [WebExtensionWindowConfiguration] from an unsafe.Pointer.
//
// An object that encapsulates configuration options for a window in an extension.
func WebExtensionWindowConfigurationFrom(ptr unsafe.Pointer) WebExtensionWindowConfiguration {
	return WebExtensionWindowConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebExtensionWindowConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebExtensionWindowConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebExtensionWindowConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebExtensionWindowConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebExtensionWindowConfiguration */

// Indicates the frame where the window should be positioned on the main screen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/frame
func (w_ WebExtensionWindowConfiguration) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](w_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// Indicates whether the window should be focused.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/shouldBeFocused
func (w_ WebExtensionWindowConfiguration) ShouldBeFocused() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldBeFocused"))
	return rv
}/* debug [instance_properties/getter]: shouldBeFocused */


// Indicates whether the window should be private.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/shouldBePrivate
func (w_ WebExtensionWindowConfiguration) ShouldBePrivate() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldBePrivate"))
	return rv
}/* debug [instance_properties/getter]: shouldBePrivate */


// Indicates the existing tabs that should be moved to the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/tabs
func (w_ WebExtensionWindowConfiguration) Tabs() []objc.ID {
	rv := objc.Send[[]objc.ID](w_.ID, objc.Sel("tabs"))
	return rv
}/* debug [instance_properties/getter]: tabs */


// Indicates the URLs that the window should initially load as tabs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/tabURLs
func (w_ WebExtensionWindowConfiguration) TabURLs() []foundation.URL {
	rv := objc.Send[[]foundation.URL](w_.ID, objc.Sel("tabURLs"))
	return rv
}/* debug [instance_properties/getter]: tabURLs */


// Indicates the window state for the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/windowState
func (w_ WebExtensionWindowConfiguration) WindowState() WebExtensionWindowState {
	rv := objc.Send[WebExtensionWindowState](w_.ID, objc.Sel("windowState"))
	return rv
}/* debug [instance_properties/getter]: windowState */


// Indicates the window type for the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/WindowConfiguration/windowType
func (w_ WebExtensionWindowConfiguration) WindowType() WebExtensionWindowType {
	rv := objc.Send[WebExtensionWindowType](w_.ID, objc.Sel("windowType"))
	return rv
}/* debug [instance_properties/getter]: windowType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebExtensionWindowConfiguration */



