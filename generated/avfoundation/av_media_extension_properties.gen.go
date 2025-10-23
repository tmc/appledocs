// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [MediaExtensionProperties] class.
type IMediaExtensionProperties interface {
	objectivec.IObject
	// properties:
	ContainingBundleName() objc.IObject /* cross-framework: NSString */
	SetContainingBundleName(value objc.IObject /* cross-framework: NSString */)
	ContainingBundleURL() objc.IObject /* cross-framework: URL */
	SetContainingBundleURL(value objc.IObject /* cross-framework: URL */)
	ExtensionIdentifier() objc.IObject /* cross-framework: NSString */
	SetExtensionIdentifier(value objc.IObject /* cross-framework: NSString */)
	ExtensionName() objc.IObject /* cross-framework: NSString */
	SetExtensionName(value objc.IObject /* cross-framework: NSString */)
	ExtensionURL() objc.IObject /* cross-framework: URL */
	SetExtensionURL(value objc.IObject /* cross-framework: URL */)
	MediaExtensionProperties() IAVMediaExtensionProperties
	SetMediaExtensionProperties(value IAVMediaExtensionProperties)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (mc _MediaExtensionPropertiesClass) Alloc() MediaExtensionProperties {
	rv := objc.Send[MediaExtensionProperties](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The name of the containing app bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaextensionproperties/containingbundlename
func (m_ MediaExtensionProperties) ContainingBundleName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("containingBundleName"))
	return rv
}


// The name of the containing app bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaextensionproperties/containingbundlename
func (m_ MediaExtensionProperties) SetContainingBundleName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContainingBundleName:"), value)
}


// The file URL of the host application for the Media Extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaextensionproperties/containingbundleurl
func (m_ MediaExtensionProperties) ContainingBundleURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("containingBundleURL"))
	return rv
}


// The file URL of the host application for the Media Extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaextensionproperties/containingbundleurl
func (m_ MediaExtensionProperties) SetContainingBundleURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setContainingBundleURL:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaextensionproperties/extensionidentifier
func (m_ MediaExtensionProperties) ExtensionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extensionIdentifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaextensionproperties/extensionidentifier
func (m_ MediaExtensionProperties) SetExtensionIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtensionIdentifier:"), value)
}


// The name of the Media Extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaextensionproperties/extensionname
func (m_ MediaExtensionProperties) ExtensionName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("extensionName"))
	return rv
}


// The name of the Media Extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaextensionproperties/extensionname
func (m_ MediaExtensionProperties) SetExtensionName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtensionName:"), value)
}


// The file URL of the Media Extension bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaextensionproperties/extensionurl
func (m_ MediaExtensionProperties) ExtensionURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](m_.ID, objc.Sel("extensionURL"))
	return rv
}


// The file URL of the Media Extension bundle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avmediaextensionproperties/extensionurl
func (m_ MediaExtensionProperties) SetExtensionURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExtensionURL:"), value)
}


// The properties of the media extension format reader that decodes the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/mediaextensionproperties
func (m_ MediaExtensionProperties) MediaExtensionProperties() IAVMediaExtensionProperties {
	rv := objc.Send[MediaExtensionProperties](m_.ID, objc.Sel("mediaExtensionProperties"))
	return rv
}


// The properties of the media extension format reader that decodes the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/mediaextensionproperties
func (m_ MediaExtensionProperties) SetMediaExtensionProperties(value IAVMediaExtensionProperties) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMediaExtensionProperties:"), value)
}



