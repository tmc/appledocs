// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHAssetCreationRequest] class.
var (
	PHAssetCreationRequestClass     _PHAssetCreationRequestClass
	PHAssetCreationRequestClassOnce sync.Once
)

func getPHAssetCreationRequestClass() _PHAssetCreationRequestClass {
	PHAssetCreationRequestClassOnce.Do(func() {
		PHAssetCreationRequestClass = _PHAssetCreationRequestClass{objc.GetClass("PHAssetCreationRequest")}
	})
	return PHAssetCreationRequestClass
}

type _PHAssetCreationRequestClass struct {
	class objc.Class
}

// An interface definition for the [PHAssetCreationRequest] class.
type IPHAssetCreationRequest interface {
	IPHAssetChangeRequest
	AddResourceWithTypeDataOptions(type_ unsafe.Pointer, data unsafe.Pointer, options unsafe.Pointer)
	AddResourceWithTypeFileURLOptions(type_ unsafe.Pointer, fileURL unsafe.Pointer, options unsafe.Pointer)
}

// A request to create a new Photos asset from underlying data resources, for use in a photo library change block.
//
// A object, used within a photo library change block, constructs a new photo or video asset from data resources, and adds it to the Photos library. This class works in terms of the raw data resources that together form an asset, so you can use it together with the class to perform a complete copy (or backup and restore) of an asset’s underlying resources. To instead simply create a new asset from an image object, image file, or video file, see the superclass . To create a new asset from data resources, first start a change block using the shared method or . Then, within the change block: Within the change block, create a new asset creation request with the method. Add image, video, or data resources using the methods in the Providing Data Resources for the New Asset section below. (Optional.) Set metadata for the new asset using methods and properties of the superclass . After Photos runs the change block and calls your completion handler, the new asset is created in the Photos library. If you instantiate or use this class outside a photo library change block, Photos throws an exception. For details on change blocks, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCreationRequest
type PHAssetCreationRequest struct {
	PHAssetChangeRequest
}

// PHAssetCreationRequestFrom constructs a [PHAssetCreationRequest] from an unsafe.Pointer.
//
// A request to create a new Photos asset from underlying data resources, for use in a photo library change block.
func PHAssetCreationRequestFrom(ptr unsafe.Pointer) PHAssetCreationRequest {
	return PHAssetCreationRequest{
		PHAssetChangeRequest: PHAssetChangeRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetCreationRequestClass) Alloc() PHAssetCreationRequest {
	rv := objc.Send[PHAssetCreationRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetCreationRequestClass) New() PHAssetCreationRequest {
	rv := objc.Send[PHAssetCreationRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAssetCreationRequest) Init() PHAssetCreationRequest {
	rv := objc.Send[PHAssetCreationRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAssetCreationRequest) Autorelease() PHAssetCreationRequest {
	rv := objc.Send[PHAssetCreationRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAssetCreationRequest creates a new PHAssetCreationRequest instance.
func NewPHAssetCreationRequest() PHAssetCreationRequest {
	return getPHAssetCreationRequestClass().New()
}


// Creates a request for adding a new asset to the Photos library using asset resources.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCreationRequest/forAsset()
func (pc _PHAssetCreationRequestClass) CreationRequestForAsset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("creationRequestForAsset"))
	return rv
}

// Returns a Boolean value indicating whether Photos supports creating an asset with the specified combination of resource types.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCreationRequest/supportsAssetResourceTypes(_:)
func (pc _PHAssetCreationRequestClass) SupportsAssetResourceTypes(types unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("supportsAssetResourceTypes:"), types)
	return rv
}

// Adds a data resource to the asset being created, using the specified data.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCreationRequest/addResource(with:data:options:)
func (p_ PHAssetCreationRequest) AddResourceWithTypeDataOptions(type_ unsafe.Pointer, data unsafe.Pointer, options unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addResourceWithType:data:options:"), type_, data, options)
}

// Adds a data resource to the asset being created, using the file at the specified URL.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCreationRequest/addResource(with:fileURL:options:)
func (p_ PHAssetCreationRequest) AddResourceWithTypeFileURLOptions(type_ unsafe.Pointer, fileURL unsafe.Pointer, options unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addResourceWithType:fileURL:options:"), type_, fileURL, options)
}



