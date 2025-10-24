// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class QLPreviewSceneActivationConfiguration */


/* debug [class_header]: Header for QLPreviewSceneActivationConfiguration */
// The class instance for the [PreviewSceneActivationConfiguration] class.
var (
	PreviewSceneActivationConfigurationClass     _PreviewSceneActivationConfigurationClass
	PreviewSceneActivationConfigurationClassOnce sync.Once
)

func getPreviewSceneActivationConfigurationClass() _PreviewSceneActivationConfigurationClass {
	PreviewSceneActivationConfigurationClassOnce.Do(func() {
		PreviewSceneActivationConfigurationClass = _PreviewSceneActivationConfigurationClass{objc.GetClass("QLPreviewSceneActivationConfiguration")}
	})
	return PreviewSceneActivationConfigurationClass
}

type _PreviewSceneActivationConfigurationClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PreviewSceneActivationConfiguration */
// An interface definition for the [PreviewSceneActivationConfiguration] class.
type IPreviewSceneActivationConfiguration interface {
	IWindowSceneActivationConfiguration
	
/* debug [class_interface_properties]: Properties for PreviewSceneActivationConfiguration */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PreviewSceneActivationConfiguration */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PreviewSceneActivationConfiguration */
// Alloc allocates a new instance without initialization.
func (pc _PreviewSceneActivationConfigurationClass) Alloc() PreviewSceneActivationConfiguration {
	rv := objc.Send[PreviewSceneActivationConfiguration](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PreviewSceneActivationConfigurationClass) New() PreviewSceneActivationConfiguration {
	rv := objc.Send[PreviewSceneActivationConfiguration](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PreviewSceneActivationConfiguration) Init() PreviewSceneActivationConfiguration {
	rv := objc.Send[PreviewSceneActivationConfiguration](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PreviewSceneActivationConfiguration) Autorelease() PreviewSceneActivationConfiguration {
	rv := objc.Send[PreviewSceneActivationConfiguration](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPreviewSceneActivationConfiguration creates a new PreviewSceneActivationConfiguration instance.
func NewPreviewSceneActivationConfiguration() PreviewSceneActivationConfiguration {
	return getPreviewSceneActivationConfigurationClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PreviewSceneActivationConfiguration */
// A scene configuration to preview items at the specified URLs.
//
// This class provides the configuration for a prominent scene presentation of a preview, either from a swipe gesture or a menu action. The user can detach the prominent Quick Look window and display it independently. To provide a preview from a swipe gesture, use an instance of this class with . To provide a preview from a menu action, use an instance of this class with .


// A scene configuration to preview items at the specified URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewSceneActivationConfiguration
type PreviewSceneActivationConfiguration struct {
	WindowSceneActivationConfiguration
}

// PreviewSceneActivationConfigurationFrom constructs a [PreviewSceneActivationConfiguration] from an unsafe.Pointer.
//
// A scene configuration to preview items at the specified URLs.
func PreviewSceneActivationConfigurationFrom(ptr unsafe.Pointer) PreviewSceneActivationConfiguration {
	return PreviewSceneActivationConfiguration{
		WindowSceneActivationConfiguration: WindowSceneActivationConfigurationFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PreviewSceneActivationConfiguration */

// Creates a preview scene configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuickLook/QLPreviewSceneActivationConfiguration/init(itemsAt:options:)
func NewPreviewSceneActivationConfigurationWithItemsAtURLsOptions(urls []foundation.URL, options IQLPreviewSceneOptions) PreviewSceneActivationConfiguration {
	instance := getPreviewSceneActivationConfigurationClass().Alloc()
	rv := objc.Send[PreviewSceneActivationConfiguration](instance.ID, objc.Sel("initWithItemsAtURLs:options:"), urls, options)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPreviewSceneActivationConfigurationWithItemsAtURLsOptions */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PreviewSceneActivationConfiguration */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PreviewSceneActivationConfiguration */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PreviewSceneActivationConfiguration */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PreviewSceneActivationConfiguration */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class QLPreviewSceneActivationConfiguration */


