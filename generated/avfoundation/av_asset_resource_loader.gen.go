// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AssetResourceLoader] class.
var (
	AssetResourceLoaderClass     _AssetResourceLoaderClass
	AssetResourceLoaderClassOnce sync.Once
)

func getAssetResourceLoaderClass() _AssetResourceLoaderClass {
	AssetResourceLoaderClassOnce.Do(func() {
		AssetResourceLoaderClass = _AssetResourceLoaderClass{objc.GetClass("AVAssetResourceLoader")}
	})
	return AssetResourceLoaderClass
}

type _AssetResourceLoaderClass struct {
	class objc.Class
}

// An interface definition for the [AssetResourceLoader] class.
type IAssetResourceLoader interface {
	objectivec.IObject
}

// An object that mediates resource requests from a URL asset.
//
// You do not create resource loader objects yourself. Instead, you retrieve a resource loader from the property of an object and use it to assign your custom delegate object. The delegate you associate with this object must adopt the protocol. For more information, see .
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader
type AssetResourceLoader struct {
	objectivec.Object
}

// AssetResourceLoaderFrom constructs a [AssetResourceLoader] from an unsafe.Pointer.
//
// An object that mediates resource requests from a URL asset.
func AssetResourceLoaderFrom(ptr unsafe.Pointer) AssetResourceLoader {
	return AssetResourceLoader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetResourceLoaderClass) Alloc() AssetResourceLoader {
	rv := objc.Send[AssetResourceLoader](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetResourceLoaderClass) New() AssetResourceLoader {
	rv := objc.Send[AssetResourceLoader](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetResourceLoader) Init() AssetResourceLoader {
	rv := objc.Send[AssetResourceLoader](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetResourceLoader) Autorelease() AssetResourceLoader {
	rv := objc.Send[AssetResourceLoader](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetResourceLoader creates a new AssetResourceLoader instance.
func NewAssetResourceLoader() AssetResourceLoader {
	return getAssetResourceLoaderClass().New()
}


// A Boolean value that indicates whether to enable attaching Common Media Client Data as HTTP request headers.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/sendsCommonMediaClientDataAsHTTPHeaders
func (a_ AssetResourceLoader) SendsCommonMediaClientDataAsHTTPHeaders() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("sendsCommonMediaClientDataAsHTTPHeaders"))
	return rv
}


// SetSendsCommonMediaClientDataAsHTTPHeaders sets the value of the sendsCommonMediaClientDataAsHTTPHeaders property.
// A Boolean value that indicates whether to enable attaching Common Media Client Data as HTTP request headers.

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/sendsCommonMediaClientDataAsHTTPHeaders
func (a_ AssetResourceLoader) SetSendsCommonMediaClientDataAsHTTPHeaders(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSendsCommonMediaClientDataAsHTTPHeaders:"), value)
}



