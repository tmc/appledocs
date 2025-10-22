// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	ContentEditingOutput() PHContentEditingOutput
	SetContentEditingOutput(value IPHContentEditingOutput)
	PlaceholderForCreatedAsset() PHObjectPlaceholder
	CreationDate() foundation.Date
	SetCreationDate(value foundation.IDate)
	IsFavorite() bool
	SetIsFavorite(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	Location() corelocation.Location
	SetLocation(value corelocation.ILocation)
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
func NewPHAssetChangeRequestForAsset(asset IPHAsset) PHAssetChangeRequest {
	rv := objc.Send[PHAssetChangeRequest](objc.ID(getPHAssetChangeRequestClass().class), objc.Sel("changeRequestForAsset:"), asset)
	return rv
}


// Requests that the specified assets be deleted.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/deleteAssets(_:)
func (pc _PHAssetChangeRequestClass) DeleteAssets(assets objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("deleteAssets:"), assets)
}

// Creates a request for modifying the specified asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/init(for:)
func (pc _PHAssetChangeRequestClass) ChangeRequestForAsset(asset IPHAsset) unsafe.Pointer {
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
func (p_ PHAssetChangeRequest) ContentEditingOutput() PHContentEditingOutput {
	rv := objc.Send[PHContentEditingOutput](p_.ID, objc.Sel("contentEditingOutput"))
	return rv
}


// SetContentEditingOutput sets the value of the contentEditingOutput property.
// The output of an asset content editing session.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/contentEditingOutput
func (p_ PHAssetChangeRequest) SetContentEditingOutput(value IPHContentEditingOutput) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentEditingOutput:"), value)
}

// A placeholder object for the asset that the change request creates.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetChangeRequest/placeholderForCreatedAsset
func (p_ PHAssetChangeRequest) PlaceholderForCreatedAsset() PHObjectPlaceholder {
	rv := objc.Send[PHObjectPlaceholder](p_.ID, objc.Sel("placeholderForCreatedAsset"))
	return rv
}

// The date and time at which the asset claims to have been originally created.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/creationdate
func (p_ PHAssetChangeRequest) CreationDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("creationDate"))
	return rv
}


// SetCreationDate sets the value of the creationDate property.
// The date and time at which the asset claims to have been originally created.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/creationdate
func (p_ PHAssetChangeRequest) SetCreationDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCreationDate:"), value)
}

// A Boolean value that indicates whether the asset is marked as one of the user’s favorites.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/isfavorite
func (p_ PHAssetChangeRequest) IsFavorite() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFavorite"))
	return rv
}


// SetIsFavorite sets the value of the isFavorite property.
// A Boolean value that indicates whether the asset is marked as one of the user’s favorites.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/isfavorite
func (p_ PHAssetChangeRequest) SetIsFavorite(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFavorite:"), value)
}

// A Boolean value that indicates whether the asset is hidden in collections.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/ishidden
func (p_ PHAssetChangeRequest) IsHidden() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isHidden"))
	return rv
}


// SetIsHidden sets the value of the isHidden property.
// A Boolean value that indicates whether the asset is hidden in collections.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/ishidden
func (p_ PHAssetChangeRequest) SetIsHidden(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsHidden:"), value)
}

// The location information saved with the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/location
func (p_ PHAssetChangeRequest) Location() corelocation.Location {
	rv := objc.Send[corelocation.Location](p_.ID, objc.Sel("location"))
	return rv
}


// SetLocation sets the value of the location property.
// The location information saved with the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetchangerequest/location
func (p_ PHAssetChangeRequest) SetLocation(value corelocation.ILocation) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocation:"), value)
}


