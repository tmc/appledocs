// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AssetLocalIdentifier() objc.IObject /* cross-framework: NSString */
	SetAssetLocalIdentifier(value objc.IObject /* cross-framework: NSString */)
	ContentType() objc.IObject /* cross-framework: UTType */
	SetContentType(value objc.IObject /* cross-framework: UTType */)
	OriginalFilename() objc.IObject /* cross-framework: NSString */
	SetOriginalFilename(value objc.IObject /* cross-framework: NSString */)
	PixelHeight() int
	SetPixelHeight(value int)
	PixelWidth() int
	SetPixelWidth(value int)
	Type() unsafe.Pointer
	SetType(value unsafe.Pointer)
	UniformTypeIdentifier() objc.IObject /* cross-framework: NSString */
	SetUniformTypeIdentifier(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An underlying data resource associated with a photo, video, or Live Photo asset in the Photos library.
//
// Each object references one or more resources. Use these objects to work with those resources directly, like when backing up or restoring assets. A photo asset can contain both JPEG and RAW files representing the same photo. A Live Photo asset contains both still photo and video resources. An edited asset contains resources representing asset content before and after the edit, as well as a resource corresponding to the object that describes the edit. To work with the data contained in an asset resource, fetch it using the class.

// An underlying data resource associated with a photo, video, or Live Photo asset in the Photos library.
//
// [Full Topic]
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

// The unique identifier the system associates for a local asset object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/assetlocalidentifier
func (p_ PHAssetResource) AssetLocalIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("assetLocalIdentifier"))
	return rv
}

// The unique identifier the system associates for a local asset object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/assetlocalidentifier
func (p_ PHAssetResource) SetAssetLocalIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAssetLocalIdentifier:"), value)
}

// The type of data associated with this asset resource (the data can be retrieved via PHAssetResourceManager)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/contenttype
func (p_ PHAssetResource) ContentType() objc.IObject /* cross-framework: UTType */ {
	rv := objc.Send[uniformtypeidentifiers.UTType](p_.ID, objc.Sel("contentType"))
	return rv
}

// The type of data associated with this asset resource (the data can be retrieved via PHAssetResourceManager)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/contenttype
func (p_ PHAssetResource) SetContentType(value objc.IObject /* cross-framework: UTType */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentType:"), value)
}

// The original filename of the asset resource from when it was created or imported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/originalfilename
func (p_ PHAssetResource) OriginalFilename() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("originalFilename"))
	return rv
}

// The original filename of the asset resource from when it was created or imported.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/originalfilename
func (p_ PHAssetResource) SetOriginalFilename(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOriginalFilename:"), value)
}

// The height of the resource, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/pixelheight
func (p_ PHAssetResource) PixelHeight() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pixelHeight"))
	return rv
}

// The height of the resource, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/pixelheight
func (p_ PHAssetResource) SetPixelHeight(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelHeight:"), value)
}

// The width of the resource, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/pixelwidth
func (p_ PHAssetResource) PixelWidth() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pixelWidth"))
	return rv
}

// The width of the resource, in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/pixelwidth
func (p_ PHAssetResource) SetPixelWidth(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelWidth:"), value)
}

// The relationship of an asset resource to its owning asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/type
func (p_ PHAssetResource) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("type"))
	return rv
}

// The relationship of an asset resource to its owning asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/type
func (p_ PHAssetResource) SetType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setType:"), value)
}

// The uniform type identifier for the asset resource’s image or video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/uniformtypeidentifier
func (p_ PHAssetResource) UniformTypeIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("uniformTypeIdentifier"))
	return rv
}

// The uniform type identifier for the asset resource’s image or video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/uniformtypeidentifier
func (p_ PHAssetResource) SetUniformTypeIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUniformTypeIdentifier:"), value)
}
