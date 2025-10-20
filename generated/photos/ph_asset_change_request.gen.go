// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHAssetChangeRequest] class.
var (
	PHAssetChangeRequestClass     _PHAssetChangeRequestClass
	PHAssetChangeRequestClassOnce sync.Once
)

func getPHAssetChangeRequestClass() _PHAssetChangeRequestClass {
	PHAssetChangeRequestClassOnce.Do(func() {
		PHAssetChangeRequestClass = _PHAssetChangeRequestClass{objc.GetClass("PHAssetChangeRequest")}
	})
	return PHAssetChangeRequestClass
}

type _PHAssetChangeRequestClass struct {
	class objc.Class
}

// An interface definition for the [PHAssetChangeRequest] class.
type IPHAssetChangeRequest interface {
	IPHChangeRequest
	RevertAssetContentToOriginal()
}

// A request to create, delete, change metadata for, or edit the content of a Photos asset, for use in a photo library change block.
//
// You use the class to request changes for objects. To make changes to assets in the Photos library, create a change request by using the appropriate class method for the change you want to perform. Call one of the methods listed in Adding New Assets to create a new asset from an image or video file. Call the method to delete existing assets. Call the method to modify an asset’s content or metadata. A change request for creating or modifying an asset works like a mutable version of the asset object. Use the change request’s properties to request changes to the corresponding properties of the asset itself. For example, the following code uses the property of a change request to mark an asset as a favorite: After Photos runs the change block and calls your completion handler, the asset’s state reflects the changes that you requested in the block. If you create or use a change request object outside a photo library change block, Photos raises an Objective-C exception. For details on change blocks, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest
type PHAssetChangeRequest struct {
	PHChangeRequest
}

// PHAssetChangeRequestFrom constructs a [PHAssetChangeRequest] from an unsafe.Pointer.
//
// A request to create, delete, change metadata for, or edit the content of a Photos asset, for use in a photo library change block.
func PHAssetChangeRequestFrom(ptr unsafe.Pointer) PHAssetChangeRequest {
	return PHAssetChangeRequest{
		PHChangeRequest: PHChangeRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetChangeRequestClass) Alloc() PHAssetChangeRequest {
	rv := objc.Send[PHAssetChangeRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetChangeRequestClass) New() PHAssetChangeRequest {
	rv := objc.Send[PHAssetChangeRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAssetChangeRequest) Init() PHAssetChangeRequest {
	rv := objc.Send[PHAssetChangeRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAssetChangeRequest) Autorelease() PHAssetChangeRequest {
	rv := objc.Send[PHAssetChangeRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAssetChangeRequest creates a new PHAssetChangeRequest instance.
func NewPHAssetChangeRequest() PHAssetChangeRequest {
	return getPHAssetChangeRequestClass().New()
}


// Creates a request for modifying the specified asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/init(for:)
func NewPHAssetChangeRequestForAsset(asset unsafe.Pointer) PHAssetChangeRequest {
	rv := objc.Send[PHAssetChangeRequest](objc.ID(getPHAssetChangeRequestClass().class), objc.Sel("changeRequestForAsset:"), asset)
	return rv
}


// Requests that the specified assets be deleted.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/deleteAssets(_:)
func (pc _PHAssetChangeRequestClass) DeleteAssets(assets objc.ID) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("deleteAssets:"), assets)
}

// Creates a request for modifying the specified asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/init(for:)
func (pc _PHAssetChangeRequestClass) ChangeRequestForAsset(asset unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("changeRequestForAsset:"), asset)
	return rv
}

// Request to revert any edits made to the asset’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/revertAssetContentToOriginal()
func (p_ PHAssetChangeRequest) RevertAssetContentToOriginal() {
	objc.Send[objc.ID](p_.ID, objc.Sel("revertAssetContentToOriginal"))
}

// The output of an asset content editing session.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/contentEditingOutput
func (p_ PHAssetChangeRequest) ContentEditingOutput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("contentEditingOutput"))
	return rv
}


// SetContentEditingOutput sets the value of the contentEditingOutput property.
// The output of an asset content editing session.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/contentEditingOutput
func (p_ PHAssetChangeRequest) SetContentEditingOutput(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentEditingOutput:"), value)
}
// A placeholder object for the asset that the change request creates.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/placeholderForCreatedAsset
func (p_ PHAssetChangeRequest) PlaceholderForCreatedAsset() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("placeholderForCreatedAsset"))
	return rv
}


