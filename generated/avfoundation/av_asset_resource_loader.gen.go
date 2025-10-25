// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetResourceLoader */


/* debug [class_header]: Header for AVAssetResourceLoader */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetResourceLoader */
// An interface definition for the [AssetResourceLoader] class.
type IAssetResourceLoader interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetResourceLoader */
	// properties:
	Delegate() unsafe.Pointer
	DelegateQueue() objectivec.IObject
	PreloadsEligibleContentKeys() bool
	SetPreloadsEligibleContentKeys(value bool)
	SendsCommonMediaClientDataAsHTTPHeaders() bool
	SetSendsCommonMediaClientDataAsHTTPHeaders(value bool)
	ResourceLoader() IAVAssetResourceLoader
	SetResourceLoader(value IAVAssetResourceLoader)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetResourceLoader */
	// methods:
	SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetResourceLoader */
// Alloc allocates a new instance without initialization.
func (ac _AssetResourceLoaderClass) Alloc() AssetResourceLoader {
	rv := objc.Send[AssetResourceLoader](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetResourceLoader */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetResourceLoader *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetResourceLoader */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetResourceLoader */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetResourceLoader */

// Sets the delegate and dispatch queue to use with the resource loader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/setDelegate(_:queue:)
func (a_ AssetResourceLoader) SetDelegateQueue(delegate unsafe.Pointer, delegateQueue objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:queue:"), delegate, delegateQueue)
}/* debug [instance_methods/method]: SetDelegateQueue */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetResourceLoader */

// The delegate object to use when handling resource requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/delegate
func (a_ AssetResourceLoader) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The dispatch queue to use when handling resource requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/delegateQueue
func (a_ AssetResourceLoader) DelegateQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("delegateQueue"))
	return rv
}/* debug [instance_properties/getter]: delegateQueue */


// A Boolean value that indicates whether content keys will be loaded as quickly as possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/preloadsEligibleContentKeys
func (a_ AssetResourceLoader) PreloadsEligibleContentKeys() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("preloadsEligibleContentKeys"))
	return rv
}/* debug [instance_properties/getter]: preloadsEligibleContentKeys */


// A Boolean value that indicates whether content keys will be loaded as quickly as possible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/preloadsEligibleContentKeys
func (a_ AssetResourceLoader) SetPreloadsEligibleContentKeys(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreloadsEligibleContentKeys:"), value)
}/* debug [instance_properties/setter]: preloadsEligibleContentKeys */


// A Boolean value that indicates whether to enable attaching Common Media Client Data as HTTP request headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/sendsCommonMediaClientDataAsHTTPHeaders
func (a_ AssetResourceLoader) SendsCommonMediaClientDataAsHTTPHeaders() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("sendsCommonMediaClientDataAsHTTPHeaders"))
	return rv
}/* debug [instance_properties/getter]: sendsCommonMediaClientDataAsHTTPHeaders */


// A Boolean value that indicates whether to enable attaching Common Media Client Data as HTTP request headers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoader/sendsCommonMediaClientDataAsHTTPHeaders
func (a_ AssetResourceLoader) SetSendsCommonMediaClientDataAsHTTPHeaders(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSendsCommonMediaClientDataAsHTTPHeaders:"), value)
}/* debug [instance_properties/setter]: sendsCommonMediaClientDataAsHTTPHeaders */


// The resource loader for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/resourceloader
func (a_ AssetResourceLoader) ResourceLoader() IAVAssetResourceLoader {
	rv := objc.Send[AssetResourceLoader](a_.ID, objc.Sel("resourceLoader"))
	return rv
}/* debug [instance_properties/getter]: resourceLoader */


// The resource loader for the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avurlasset/resourceloader
func (a_ AssetResourceLoader) SetResourceLoader(value IAVAssetResourceLoader) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setResourceLoader:"), value)
}/* debug [instance_properties/setter]: resourceLoader */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetResourceLoader */



