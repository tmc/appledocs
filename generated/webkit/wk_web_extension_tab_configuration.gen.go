// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class WKWebExtensionTabConfiguration */


/* debug [class_header]: Header for WKWebExtensionTabConfiguration */
// The class instance for the [WebExtensionTabConfiguration] class.
var (
	WebExtensionTabConfigurationClass     _WebExtensionTabConfigurationClass
	WebExtensionTabConfigurationClassOnce sync.Once
)

func getWebExtensionTabConfigurationClass() _WebExtensionTabConfigurationClass {
	WebExtensionTabConfigurationClassOnce.Do(func() {
		WebExtensionTabConfigurationClass = _WebExtensionTabConfigurationClass{objc.GetClass("WKWebExtensionTabConfiguration")}
	})
	return WebExtensionTabConfigurationClass
}

type _WebExtensionTabConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WebExtensionTabConfiguration */
// An interface definition for the [WebExtensionTabConfiguration] class.
type IWebExtensionTabConfiguration interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WebExtensionTabConfiguration */
	// properties:
	Index() uint
	ParentTab() unsafe.Pointer
	ShouldAddToSelection() bool
	ShouldBeActive() bool
	ShouldBeMuted() bool
	ShouldBePinned() bool
	ShouldReaderModeBeActive() bool
	Url() objc.IObject /* cross-framework: NSURL */
	Window() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WebExtensionTabConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WebExtensionTabConfiguration */
// Alloc allocates a new instance without initialization.
func (wc _WebExtensionTabConfigurationClass) Alloc() WebExtensionTabConfiguration {
	rv := objc.Send[WebExtensionTabConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WebExtensionTabConfigurationClass) New() WebExtensionTabConfiguration {
	rv := objc.Send[WebExtensionTabConfiguration](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WebExtensionTabConfiguration) Init() WebExtensionTabConfiguration {
	rv := objc.Send[WebExtensionTabConfiguration](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WebExtensionTabConfiguration) Autorelease() WebExtensionTabConfiguration {
	rv := objc.Send[WebExtensionTabConfiguration](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWebExtensionTabConfiguration creates a new WebExtensionTabConfiguration instance.
func NewWebExtensionTabConfiguration() WebExtensionTabConfiguration {
	return getWebExtensionTabConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WebExtensionTabConfiguration */
// An object that encapsulates configuration options for a tab in an extension.
//
// This class holds various options that influence the behavior and initial state of a tab. The app retains the discretion to disregard any or all of these options, or even opt not to create a tab.


// An object that encapsulates configuration options for a tab in an extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration
type WebExtensionTabConfiguration struct {
	objectivec.Object
}

// WebExtensionTabConfigurationFrom constructs a [WebExtensionTabConfiguration] from an unsafe.Pointer.
//
// An object that encapsulates configuration options for a tab in an extension.
func WebExtensionTabConfigurationFrom(ptr unsafe.Pointer) WebExtensionTabConfiguration {
	return WebExtensionTabConfiguration{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WebExtensionTabConfiguration *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WebExtensionTabConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WebExtensionTabConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WebExtensionTabConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WebExtensionTabConfiguration */

// Indicates the position where the tab should be opened within the window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/index
func (w_ WebExtensionTabConfiguration) Index() uint {
	rv := objc.Send[uint](w_.ID, objc.Sel("index"))
	return rv
}/* debug [instance_properties/getter]: index */


// Indicates the parent tab with which the tab should be related.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/parentTab
func (w_ WebExtensionTabConfiguration) ParentTab() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("parentTab"))
	return rv
}/* debug [instance_properties/getter]: parentTab */


// Indicates whether the tab should be added to the current tab selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/shouldAddToSelection
func (w_ WebExtensionTabConfiguration) ShouldAddToSelection() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldAddToSelection"))
	return rv
}/* debug [instance_properties/getter]: shouldAddToSelection */


// Indicates whether the tab should be the active tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/shouldBeActive
func (w_ WebExtensionTabConfiguration) ShouldBeActive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldBeActive"))
	return rv
}/* debug [instance_properties/getter]: shouldBeActive */


// Indicates whether the tab should be muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/shouldBeMuted
func (w_ WebExtensionTabConfiguration) ShouldBeMuted() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldBeMuted"))
	return rv
}/* debug [instance_properties/getter]: shouldBeMuted */


// Indicates whether the tab should be pinned.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/shouldBePinned
func (w_ WebExtensionTabConfiguration) ShouldBePinned() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldBePinned"))
	return rv
}/* debug [instance_properties/getter]: shouldBePinned */


// Indicates whether reader mode in the tab should be active.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/shouldReaderModeBeActive
func (w_ WebExtensionTabConfiguration) ShouldReaderModeBeActive() bool {
	rv := objc.Send[bool](w_.ID, objc.Sel("shouldReaderModeBeActive"))
	return rv
}/* debug [instance_properties/getter]: shouldReaderModeBeActive */


// Indicates the initial URL for the tab.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/url
func (w_ WebExtensionTabConfiguration) Url() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](w_.ID, objc.Sel("url"))
	return rv
}/* debug [instance_properties/getter]: url */


// Indicates the window where the tab should be opened.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/WebKit/WKWebExtension/TabConfiguration/window
func (w_ WebExtensionTabConfiguration) Window() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](w_.ID, objc.Sel("window"))
	return rv
}/* debug [instance_properties/getter]: window */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class WKWebExtensionTabConfiguration */



