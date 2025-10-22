// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHAssetCollection] class.
var (
	PHAssetCollectionClass     _PHAssetCollectionClass
	PHAssetCollectionClassOnce sync.Once
)

func getPHAssetCollectionClass() _PHAssetCollectionClass {
	PHAssetCollectionClassOnce.Do(func() {
		PHAssetCollectionClass = _PHAssetCollectionClass{objc.GetClass("PHAssetCollection")}
	})
	return PHAssetCollectionClass
}

type _PHAssetCollectionClass struct {
	class objc.Class
}

// An interface definition for the [PHAssetCollection] class.
type IPHAssetCollection interface {
	IPHCollection
	ApproximateLocation() corelocation.Location
	AssetCollectionSubtype() PHAssetCollectionSubtype
	AssetCollectionType() PHAssetCollectionType
	EndDate() foundation.NSDate
	EstimatedAssetCount() uint
	LocalizedLocationNames() []string
	StartDate() foundation.NSDate
}

// A representation of a Photos asset grouping, such as a moment, user-created album, or smart album.
//
// In the Photos framework, collection objects (including asset collections) do not directly reference their member objects, and there are no other objects that directly reference collection objects. To retrieve the members of an asset collection, fetch them with a class method such as . To find asset collections, use one of the methods listed in the Fetching Asset Collections group below. Like assets and collection lists, asset collections are immutable. To create, rename, or delete asset collections, or to add, remove, or rearrange members in an asset collection, create a object within a photo library change block. For details on using change requests and change blocks to update the photo library, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection
type PHAssetCollection struct {
	PHCollection
}

// PHAssetCollectionFrom constructs a [PHAssetCollection] from an unsafe.Pointer.
//
// A representation of a Photos asset grouping, such as a moment, user-created album, or smart album.
func PHAssetCollectionFrom(ptr unsafe.Pointer) PHAssetCollection {
	return PHAssetCollection{
		PHCollection: PHCollectionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHAssetCollectionClass) Alloc() PHAssetCollection {
	rv := objc.Send[PHAssetCollection](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHAssetCollectionClass) New() PHAssetCollection {
	rv := objc.Send[PHAssetCollection](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHAssetCollection) Init() PHAssetCollection {
	rv := objc.Send[PHAssetCollection](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHAssetCollection) Autorelease() PHAssetCollection {
	rv := objc.Send[PHAssetCollection](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHAssetCollection creates a new PHAssetCollection instance.
func NewPHAssetCollection() PHAssetCollection {
	return getPHAssetCollectionClass().New()
}


// Retrieves asset collections of the specified type and subtype.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/fetchAssetCollections(with:subtype:options:)
func (pc _PHAssetCollectionClass) FetchAssetCollectionsWithTypeSubtypeOptions(type_ PHAssetCollectionType, subtype IPHAssetCollectionSubtype, options PHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchAssetCollectionsWithType:subtype:options:"), type_, subtype, options)
	return rv
}

// Retrieves asset collections using URLs provided by the Assets Library framework.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/fetchAssetCollections(withALAssetGroupURLs:options:)
func (pc _PHAssetCollectionClass) FetchAssetCollectionsWithALAssetGroupURLsOptions(assetGroupURLs []foundation.IURL, options PHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchAssetCollectionsWithALAssetGroupURLs:options:"), assetGroupURLs, options)
	return rv
}

// Retrieves asset collections with the specified unique identifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/fetchAssetCollections(withLocalIdentifiers:options:)
func (pc _PHAssetCollectionClass) FetchAssetCollectionsWithLocalIdentifiersOptions(identifiers []string, options PHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchAssetCollectionsWithLocalIdentifiers:options:"), identifiers, options)
	return rv
}

// Retrieves asset collections of the specified type containing the specified asset.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/fetchAssetCollectionsContaining(_:with:options:)
func (pc _PHAssetCollectionClass) FetchAssetCollectionsContainingAssetWithTypeOptions(asset IPHAsset, type_ PHAssetCollectionType, options PHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchAssetCollectionsContainingAsset:withType:options:"), asset, type_, options)
	return rv
}

// Retrieves asset collections in the specified moment list collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/fetchMoments(inMomentList:options:)
func (pc _PHAssetCollectionClass) FetchMomentsInMomentListOptions(momentList IPHCollectionList, options PHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchMomentsInMomentList:options:"), momentList, options)
	return rv
}

// Retrieves asset collections corresponding to moments seen in the Photos app.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/fetchMoments(with:)
func (pc _PHAssetCollectionClass) FetchMomentsWithOptions(options PHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchMomentsWithOptions:"), options)
	return rv
}

// Creates a temporary asset collection containing the specified assets.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/transientAssetCollection(with:title:)
func (pc _PHAssetCollectionClass) TransientAssetCollectionWithAssetsTitle(assets []PHAsset, title string) PHAssetCollection {
	rv := objc.Send[PHAssetCollection](objc.ID(pc.class), objc.Sel("transientAssetCollectionWithAssets:title:"), assets, objc.String(title))
	return rv
}

// Creates a temporary asset collection containing the assets from the specified fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/transientAssetCollection(withAssetFetchResult:title:)
func (pc _PHAssetCollectionClass) TransientAssetCollectionWithAssetFetchResultTitle(fetchResult unsafe.Pointer, title string) PHAssetCollection {
	rv := objc.Send[PHAssetCollection](objc.ID(pc.class), objc.Sel("transientAssetCollectionWithAssetFetchResult:title:"), fetchResult, objc.String(title))
	return rv
}

// A location representing those of all assets in the collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/approximateLocation
func (p_ PHAssetCollection) ApproximateLocation() corelocation.Location {
	rv := objc.Send[corelocation.Location](p_.ID, objc.Sel("approximateLocation"))
	return rv
}

// The subtype of the asset collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/assetCollectionSubtype
func (p_ PHAssetCollection) AssetCollectionSubtype() PHAssetCollectionSubtype {
	rv := objc.Send[PHAssetCollectionSubtype](p_.ID, objc.Sel("assetCollectionSubtype"))
	return rv
}

// The type of the asset collection, such as an album or a moment.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/assetCollectionType
func (p_ PHAssetCollection) AssetCollectionType() PHAssetCollectionType {
	rv := objc.Send[PHAssetCollectionType](p_.ID, objc.Sel("assetCollectionType"))
	return rv
}

// The latest creation date among all assets in the asset collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/endDate
func (p_ PHAssetCollection) EndDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("endDate"))
	return rv
}

// The estimated number of assets in the asset collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/estimatedAssetCount
func (p_ PHAssetCollection) EstimatedAssetCount() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("estimatedAssetCount"))
	return rv
}

// The names of locations grouped by the collection (an array of objects).
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/localizedLocationNames
func (p_ PHAssetCollection) LocalizedLocationNames() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("localizedLocationNames"))
	return rv
}

// The earliest creation date among all assets in the asset collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/startDate
func (p_ PHAssetCollection) StartDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("startDate"))
	return rv
}



