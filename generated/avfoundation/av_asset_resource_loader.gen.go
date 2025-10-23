// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	SendsCommonMediaClientDataAsHTTPHeaders() bool /* primitive/slice/pointer */
	SetSendsCommonMediaClientDataAsHTTPHeaders(value bool /* primitive/slice/pointer */)
	Delegate() AssetResourceLoaderDelegate /* not a class type */
	SetDelegate(value AssetResourceLoaderDelegate /* not a class type */)
	DelegateQueue() unsafe.Pointer
	SetDelegateQueue(value unsafe.Pointer)
	PreloadsEligibleContentKeys() bool /* primitive/slice/pointer */
	SetPreloadsEligibleContentKeys(value bool /* primitive/slice/pointer */)
	ResourceLoader() IAVAssetResourceLoader
	SetResourceLoader(value IAVAssetResourceLoader)
	// methods:
}

// An object that mediates resource requests from a URL asset.
//
// You do not create resource loader objects yourself. Instead, you retrieve a resource loader from the property of an object and use it to assign your custom delegate object. The delegate you associate with this object must adopt the protocol. For more information, see .


// An object that mediates resource requests from a URL asset.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/sendsCommonMediaClientDataAsHTTPHeaders
func (a_ AssetResourceLoader) SendsCommonMediaClientDataAsHTTPHeaders() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("sendsCommonMediaClientDataAsHTTPHeaders"))
	return rv
}


// A Boolean value that indicates whether to enable attaching Common Media Client Data as HTTP request headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/sendsCommonMediaClientDataAsHTTPHeaders
func (a_ AssetResourceLoader) SetSendsCommonMediaClientDataAsHTTPHeaders(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSendsCommonMediaClientDataAsHTTPHeaders:"), value)
}


// The delegate object to use when handling resource requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/delegate
func (a_ AssetResourceLoader) Delegate() AssetResourceLoaderDelegate /* not a class type */ {
	rv := objc.Send[AssetResourceLoaderDelegate](a_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate object to use when handling resource requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/delegate
func (a_ AssetResourceLoader) SetDelegate(value AssetResourceLoaderDelegate /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}


// The dispatch queue to use when handling resource requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/delegatequeue
func (a_ AssetResourceLoader) DelegateQueue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegateQueue"))
	return rv
}


// The dispatch queue to use when handling resource requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/delegatequeue
func (a_ AssetResourceLoader) SetDelegateQueue(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegateQueue:"), value)
}


// A Boolean value that indicates whether content keys will be loaded as quickly as possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/preloadseligiblecontentkeys
func (a_ AssetResourceLoader) PreloadsEligibleContentKeys() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("preloadsEligibleContentKeys"))
	return rv
}


// A Boolean value that indicates whether content keys will be loaded as quickly as possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetresourceloader/preloadseligiblecontentkeys
func (a_ AssetResourceLoader) SetPreloadsEligibleContentKeys(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreloadsEligibleContentKeys:"), value)
}


// The resource loader for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/resourceloader
func (a_ AssetResourceLoader) ResourceLoader() IAVAssetResourceLoader {
	rv := objc.Send[AssetResourceLoader](a_.ID, objc.Sel("resourceLoader"))
	return rv
}


// The resource loader for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/resourceloader
func (a_ AssetResourceLoader) SetResourceLoader(value IAVAssetResourceLoader) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setResourceLoader:"), value)
}



