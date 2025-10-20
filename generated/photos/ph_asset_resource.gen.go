// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// The class instance for the [PHAssetResource] class.
var (
	PHAssetResourceClass     _PHAssetResourceClass
	PHAssetResourceClassOnce sync.Once
)

func getPHAssetResourceClass() _PHAssetResourceClass {
	PHAssetResourceClassOnce.Do(func() {
		PHAssetResourceClass = _PHAssetResourceClass{objc.GetClass("PHAssetResource")}
	})
	return PHAssetResourceClass
}

type _PHAssetResourceClass struct {
	class objc.Class
}

// An interface definition for the [PHAssetResource] class.
type IPHAssetResource interface {
	objectivec.IObject
}

// An underlying data resource associated with a photo, video, or Live Photo asset in the Photos library.
//
// Each object references one or more resources. Use these objects to work with those resources directly, like when backing up or restoring assets. A photo asset can contain both JPEG and RAW files representing the same photo. A Live Photo asset contains both still photo and video resources. An edited asset contains resources representing asset content before and after the edit, as well as a resource corresponding to the object that describes the edit. To work with the data contained in an asset resource, fetch it using the class.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResource
type PHAssetResource struct {
	objectivec.Object
}

// PHAssetResourceFrom constructs a [PHAssetResource] from an unsafe.Pointer.
//
// An underlying data resource associated with a photo, video, or Live Photo asset in the Photos library.
func PHAssetResourceFrom(ptr unsafe.Pointer) PHAssetResource {
	return PHAssetResource{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetResourceClass) Alloc() PHAssetResource {
	rv := objc.Send[PHAssetResource](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetResourceClass) New() PHAssetResource {
	rv := objc.Send[PHAssetResource](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAssetResource) Init() PHAssetResource {
	rv := objc.Send[PHAssetResource](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAssetResource) Autorelease() PHAssetResource {
	rv := objc.Send[PHAssetResource](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAssetResource creates a new PHAssetResource instance.
func NewPHAssetResource() PHAssetResource {
	return getPHAssetResourceClass().New()
}


// Returns the list of data resources associated with an asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResource/assetResources(for:)-27o4l
func (pc _PHAssetResourceClass) AssetResourcesForAsset(asset unsafe.Pointer) []PHAssetResource {
	rv := objc.Send[[]PHAssetResource](objc.ID(pc.class), objc.Sel("assetResourcesForAsset:"), asset)
	return rv
}

// Returns the list of data resources associated with a Live Photo object.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResource/assetResources(for:)-2fedw
func (pc _PHAssetResourceClass) AssetResourcesForLivePhoto(livePhoto unsafe.Pointer) []PHAssetResource {
	rv := objc.Send[[]PHAssetResource](objc.ID(pc.class), objc.Sel("assetResourcesForLivePhoto:"), livePhoto)
	return rv
}

// The unique identifier the system associates for a local asset object.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResource/assetLocalIdentifier
func (p_ PHAssetResource) AssetLocalIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("assetLocalIdentifier"))
	return rv
}

// The type of data associated with this asset resource (the data can be retrieved via PHAssetResourceManager)
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResource/contentType
func (p_ PHAssetResource) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[uniformtypeidentifiers.UTType](p_.ID, objc.Sel("contentType"))
	return rv
}

// The original filename of the asset resource from when it was created or imported.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResource/originalFilename
func (p_ PHAssetResource) OriginalFilename() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("originalFilename"))
	return rv
}

// The height of the resource, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResource/pixelHeight
func (p_ PHAssetResource) PixelHeight() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pixelHeight"))
	return rv
}

// The width of the resource, in pixels.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResource/pixelWidth
func (p_ PHAssetResource) PixelWidth() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pixelWidth"))
	return rv
}

// The relationship of an asset resource to its owning asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResource/type
func (p_ PHAssetResource) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("type"))
	return rv
}

// The uniform type identifier for the asset resource’s image or video data.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetResource/uniformTypeIdentifier
func (p_ PHAssetResource) UniformTypeIdentifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("uniformTypeIdentifier"))
	return rv
}



