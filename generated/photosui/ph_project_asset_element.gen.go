// Code generated from Apple documentation for PhotosUI. DO NOT EDIT.

package photosui

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/photos"
)

// The class instance for the [PHProjectAssetElement] class.
var (
	PHProjectAssetElementClass     _PHProjectAssetElementClass
	PHProjectAssetElementClassOnce sync.Once
)

func getPHProjectAssetElementClass() _PHProjectAssetElementClass {
	PHProjectAssetElementClassOnce.Do(func() {
		PHProjectAssetElementClass = _PHProjectAssetElementClass{objc.GetClass("PHProjectAssetElement")}
	})
	return PHProjectAssetElementClass
}

type _PHProjectAssetElementClass struct {
	class objc.Class
}

// An interface definition for the [PHProjectAssetElement] class.
type IPHProjectAssetElement interface {
	IPHProjectElement
}

// An element that represents a media asset within project section content.
//
// Access the underlying by converting the provided to a , then calling .
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectAssetElement
type PHProjectAssetElement struct {
	PHProjectElement
}

// PHProjectAssetElementFrom constructs a [PHProjectAssetElement] from an unsafe.Pointer.
//
// An element that represents a media asset within project section content.
func PHProjectAssetElementFrom(ptr unsafe.Pointer) PHProjectAssetElement {
	return PHProjectAssetElement{
		PHProjectElement: PHProjectElementFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHProjectAssetElementClass) Alloc() PHProjectAssetElement {
	rv := objc.Send[PHProjectAssetElement](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHProjectAssetElementClass) New() PHProjectAssetElement {
	rv := objc.Send[PHProjectAssetElement](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHProjectAssetElement) Init() PHProjectAssetElement {
	rv := objc.Send[PHProjectAssetElement](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHProjectAssetElement) Autorelease() PHProjectAssetElement {
	rv := objc.Send[PHProjectAssetElement](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHProjectAssetElement creates a new PHProjectAssetElement instance.
func NewPHProjectAssetElement() PHProjectAssetElement {
	return getPHProjectAssetElementClass().New()
}


// A string annotation attached to the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectAssetElement/annotation
func (p_ PHProjectAssetElement) Annotation() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("annotation"))
	return rv
}

// The asset’s identifier in the cloud.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectAssetElement/cloudAssetIdentifier
func (p_ PHProjectAssetElement) CloudAssetIdentifier() photos.PHCloudIdentifier {
	rv := objc.Send[photos.PHCloudIdentifier](p_.ID, objc.Sel("cloudAssetIdentifier"))
	return rv
}

// A rectangle defining the cropped portion of the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectAssetElement/cropRect
func (p_ PHProjectAssetElement) CropRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](p_.ID, objc.Sel("cropRect"))
	return rv
}

// A Boolean indicating whether the asset is vertically flipped.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectAssetElement/horizontallyFlipped
func (p_ PHProjectAssetElement) HorizontallyFlipped() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("horizontallyFlipped"))
	return rv
}

// An array of regions of interest in the photo asset.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectAssetElement/regionsOfInterest
func (p_ PHProjectAssetElement) RegionsOfInterest() []PHProjectRegionOfInterest {
	rv := objc.Send[[]PHProjectRegionOfInterest](p_.ID, objc.Sel("regionsOfInterest"))
	return rv
}

// A Boolean indicating whether the asset is vertically flipped.
//
// [Full Topic]: https://developer.apple.com/documentation/PhotosUI/PHProjectAssetElement/verticallyFlipped
func (p_ PHProjectAssetElement) VerticallyFlipped() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("verticallyFlipped"))
	return rv
}

// The unique identifier the system associates for a local asset object.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/assetlocalidentifier
func (p_ PHProjectAssetElement) AssetLocalIdentifier() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("assetLocalIdentifier"))
	return rv
}


// SetAssetLocalIdentifier sets the value of the assetLocalIdentifier property.
// The unique identifier the system associates for a local asset object.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetresource/assetlocalidentifier
func (p_ PHProjectAssetElement) SetAssetLocalIdentifier(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAssetLocalIdentifier:"), value)
}

// An array containing all cloud asset identifiers referenced in the content.
//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectsectioncontent/cloudassetidentifiers
func (p_ PHProjectAssetElement) CloudAssetIdentifiers() photos.PHCloudIdentifier {
	rv := objc.Send[photos.PHCloudIdentifier](p_.ID, objc.Sel("cloudAssetIdentifiers"))
	return rv
}


// SetCloudAssetIdentifiers sets the value of the cloudAssetIdentifiers property.
// An array containing all cloud asset identifiers referenced in the content.

//
// [Full Topic]: https://developer.apple.com/documentation/photosui/phprojectsectioncontent/cloudassetidentifiers
func (p_ PHProjectAssetElement) SetCloudAssetIdentifiers(value photos.IPHCloudIdentifier) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCloudAssetIdentifiers:"), value)
}



