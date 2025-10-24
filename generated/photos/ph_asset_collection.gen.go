// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/corelocation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	EndDate() objc.IObject /* cross-framework: NSDate */
	EstimatedAssetCount() uint
	StartDate() objc.IObject           /* cross-framework: NSDate */
	ApproximateLocation() objc.IObject /* cross-framework: Location */
	SetApproximateLocation(value objc.IObject /* cross-framework: Location */)
	AssetCollectionSubtype() unsafe.Pointer
	SetAssetCollectionSubtype(value unsafe.Pointer)
	AssetCollectionType() unsafe.Pointer
	SetAssetCollectionType(value unsafe.Pointer)
	LocalizedLocationNames() objc.IObject /* cross-framework: NSString */
	SetLocalizedLocationNames(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A representation of a Photos asset grouping, such as a moment, user-created album, or smart album.
//
// In the Photos framework, collection objects (including asset collections) do not directly reference their member objects, and there are no other objects that directly reference collection objects. To retrieve the members of an asset collection, fetch them with a class method such as . To find asset collections, use one of the methods listed in the Fetching Asset Collections group below. Like assets and collection lists, asset collections are immutable. To create, rename, or delete asset collections, or to add, remove, or rearrange members in an asset collection, create a object within a photo library change block. For details on using change requests and change blocks to update the photo library, see .

// A representation of a Photos asset grouping, such as a moment, user-created album, or smart album.
//
// [Full Topic]
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

// The latest creation date among all assets in the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/endDate
func (p_ PHAssetCollection) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("endDate"))
	return rv
}

// The estimated number of assets in the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/estimatedAssetCount
func (p_ PHAssetCollection) EstimatedAssetCount() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("estimatedAssetCount"))
	return rv
}

// The earliest creation date among all assets in the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHAssetCollection/startDate
func (p_ PHAssetCollection) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("startDate"))
	return rv
}

// A location representing those of all assets in the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/approximatelocation
func (p_ PHAssetCollection) ApproximateLocation() objc.IObject /* cross-framework: Location */ {
	rv := objc.Send[corelocation.Location](p_.ID, objc.Sel("approximateLocation"))
	return rv
}

// A location representing those of all assets in the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/approximatelocation
func (p_ PHAssetCollection) SetApproximateLocation(value objc.IObject /* cross-framework: Location */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setApproximateLocation:"), value)
}

// The subtype of the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/assetcollectionsubtype
func (p_ PHAssetCollection) AssetCollectionSubtype() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("assetCollectionSubtype"))
	return rv
}

// The subtype of the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/assetcollectionsubtype
func (p_ PHAssetCollection) SetAssetCollectionSubtype(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAssetCollectionSubtype:"), value)
}

// The type of the asset collection, such as an album or a moment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/assetcollectiontype
func (p_ PHAssetCollection) AssetCollectionType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("assetCollectionType"))
	return rv
}

// The type of the asset collection, such as an album or a moment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/assetcollectiontype
func (p_ PHAssetCollection) SetAssetCollectionType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAssetCollectionType:"), value)
}

// The names of locations grouped by the collection (an array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/localizedlocationnames
func (p_ PHAssetCollection) LocalizedLocationNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localizedLocationNames"))
	return rv
}

// The names of locations grouped by the collection (an array of
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/localizedlocationnames
func (p_ PHAssetCollection) SetLocalizedLocationNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedLocationNames:"), value)
}
