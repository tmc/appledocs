// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMediaExtensionProperties */


/* debug [class_header]: Header for AVMediaExtensionProperties */
// The class instance for the [MediaExtensionProperties] class.
var (
	MediaExtensionPropertiesClass     _MediaExtensionPropertiesClass
	MediaExtensionPropertiesClassOnce sync.Once
)

func getMediaExtensionPropertiesClass() _MediaExtensionPropertiesClass {
	MediaExtensionPropertiesClassOnce.Do(func() {
		MediaExtensionPropertiesClass = _MediaExtensionPropertiesClass{objc.GetClass("AVMediaExtensionProperties")}
	})
	return MediaExtensionPropertiesClass
}

type _MediaExtensionPropertiesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MediaExtensionProperties */
// An interface definition for the [MediaExtensionProperties] class.
type IMediaExtensionProperties interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MediaExtensionProperties */
	// properties:
	ContainingBundleName() objc.IObject /* cross-framework: NSString */
	ContainingBundleURL() objc.IObject /* cross-framework: NSURL */
	ExtensionIdentifier() objc.IObject /* cross-framework: NSString */
	ExtensionName() objc.IObject /* cross-framework: NSString */
	ExtensionURL() objc.IObject /* cross-framework: NSURL */
	MediaExtensionProperties() IAVMediaExtensionProperties
	SetMediaExtensionProperties(value IAVMediaExtensionProperties)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MediaExtensionProperties */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MediaExtensionProperties */
// Alloc allocates a new instance without initialization.
func (mc _MediaExtensionPropertiesClass) Alloc() MediaExtensionProperties {
	rv := objc.Send[MediaExtensionProperties](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MediaExtensionPropertiesClass) New() MediaExtensionProperties {
	rv := objc.Send[MediaExtensionProperties](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MediaExtensionProperties) Init() MediaExtensionProperties {
	rv := objc.Send[MediaExtensionProperties](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MediaExtensionProperties) Autorelease() MediaExtensionProperties {
	rv := objc.Send[MediaExtensionProperties](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaExtensionProperties creates a new MediaExtensionProperties instance.
func NewMediaExtensionProperties() MediaExtensionProperties {
	return getMediaExtensionPropertiesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MediaExtensionProperties */
// An object that describes a Media Extension.


// An object that describes a Media Extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties
type MediaExtensionProperties struct {
	objectivec.Object
}

// MediaExtensionPropertiesFrom constructs a [MediaExtensionProperties] from an unsafe.Pointer.
//
// An object that describes a Media Extension.
func MediaExtensionPropertiesFrom(ptr unsafe.Pointer) MediaExtensionProperties {
	return MediaExtensionProperties{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MediaExtensionProperties *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MediaExtensionProperties */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MediaExtensionProperties */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MediaExtensionProperties */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MediaExtensionProperties */

// The name of the containing app bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties/containingBundleName
func (m_ MediaExtensionProperties) ContainingBundleName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("containingBundleName"))
	return rv
}/* debug [instance_properties/getter]: containingBundleName */


// The file URL of the host application for the Media Extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties/containingBundleURL
func (m_ MediaExtensionProperties) ContainingBundleURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("containingBundleURL"))
	return rv
}/* debug [instance_properties/getter]: containingBundleURL */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties/extensionIdentifier
func (m_ MediaExtensionProperties) ExtensionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extensionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: extensionIdentifier */


// The name of the Media Extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties/extensionName
func (m_ MediaExtensionProperties) ExtensionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extensionName"))
	return rv
}/* debug [instance_properties/getter]: extensionName */


// The file URL of the Media Extension bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMediaExtensionProperties/extensionURL
func (m_ MediaExtensionProperties) ExtensionURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](m_.ID, objc.Sel("extensionURL"))
	return rv
}/* debug [instance_properties/getter]: extensionURL */


// The properties of the media extension format reader that decodes the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/mediaextensionproperties
func (m_ MediaExtensionProperties) MediaExtensionProperties() IAVMediaExtensionProperties {
	rv := objc.Send[MediaExtensionProperties](m_.ID, objc.Sel("mediaExtensionProperties"))
	return rv
}/* debug [instance_properties/getter]: mediaExtensionProperties */


// The properties of the media extension format reader that decodes the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/mediaextensionproperties
func (m_ MediaExtensionProperties) SetMediaExtensionProperties(value IAVMediaExtensionProperties) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaExtensionProperties:"), value)
}/* debug [instance_properties/setter]: mediaExtensionProperties */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMediaExtensionProperties */



