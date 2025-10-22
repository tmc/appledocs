// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHFetchOptions] class.
var (
	PHFetchOptionsClass     _PHFetchOptionsClass
	PHFetchOptionsClassOnce sync.Once
)

func getPHFetchOptionsClass() _PHFetchOptionsClass {
	PHFetchOptionsClassOnce.Do(func() {
		PHFetchOptionsClass = _PHFetchOptionsClass{objc.GetClass("PHFetchOptions")}
	})
	return PHFetchOptionsClass
}

type _PHFetchOptionsClass struct {
	class objc.Class
}

// An interface definition for the [PHFetchOptions] class.
type IPHFetchOptions interface {
	objectivec.IObject
	FetchLimit() uint
	SetFetchLimit(value uint)
	IncludeAllBurstAssets() bool
	SetIncludeAllBurstAssets(value bool)
	IncludeAssetSourceTypes() PHAssetSourceType
	SetIncludeAssetSourceTypes(value PHAssetSourceType)
	IncludeHiddenAssets() bool
	SetIncludeHiddenAssets(value bool)
	Predicate() foundation.Predicate
	SetPredicate(value foundation.IPredicate)
	SortDescriptors() []foundation.SortDescriptor
	SetSortDescriptors(value []foundation.ISortDescriptor)
	WantsIncrementalChangeDetails() bool
	SetWantsIncrementalChangeDetails(value bool)
	BurstIdentifier() string
	SetBurstIdentifier(value string)
	CreationDate() foundation.Date
	SetCreationDate(value foundation.IDate)
	Duration() unsafe.Pointer
	SetDuration(value unsafe.Pointer)
	IsFavorite() bool
	SetIsFavorite(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	MediaSubtypes() PHAssetMediaSubtype
	SetMediaSubtypes(value IPHAssetMediaSubtype)
	MediaType() PHAssetMediaType
	SetMediaType(value PHAssetMediaType)
	ModificationDate() foundation.Date
	SetModificationDate(value foundation.IDate)
	PixelHeight() int
	SetPixelHeight(value int)
	PixelWidth() int
	SetPixelWidth(value int)
	EndDate() foundation.Date
	SetEndDate(value foundation.IDate)
	EstimatedAssetCount() int
	SetEstimatedAssetCount(value int)
	StartDate() foundation.Date
	SetStartDate(value foundation.IDate)
	LocalizedTitle() string
	SetLocalizedTitle(value string)
	LocalIdentifier() string
	SetLocalIdentifier(value string)
}

// A set of options that affect the filtering, sorting, and management of results that Photos returns when you fetch asset or collection objects.
//
// Using class methods on the , , , and classes to fetch assets or collections produces a object containing the requested objects. The options you specify control which objects the fetch result includes, how those objects are arranged in the fetch result, and how Photos should notify your app of changes to the fetch result. Photos supports only a restricted set of keys for the and properties. The set of available keys depends on which class you’re using to fetch assets or collections. The following table lists the keys supported by each class:
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions
type PHFetchOptions struct {
	objectivec.Object
}

// PHFetchOptionsFrom constructs a [PHFetchOptions] from an unsafe.Pointer.
//
// A set of options that affect the filtering, sorting, and management of results that Photos returns when you fetch asset or collection objects.
func PHFetchOptionsFrom(ptr unsafe.Pointer) PHFetchOptions {
	return PHFetchOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHFetchOptionsClass) Alloc() PHFetchOptions {
	rv := objc.Send[PHFetchOptions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHFetchOptionsClass) New() PHFetchOptions {
	rv := objc.Send[PHFetchOptions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHFetchOptions) Init() PHFetchOptions {
	rv := objc.Send[PHFetchOptions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHFetchOptions) Autorelease() PHFetchOptions {
	rv := objc.Send[PHFetchOptions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHFetchOptions creates a new PHFetchOptions instance.
func NewPHFetchOptions() PHFetchOptions {
	return getPHFetchOptionsClass().New()
}


// The maximum number of objects to include in the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/fetchLimit
func (p_ PHFetchOptions) FetchLimit() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("fetchLimit"))
	return rv
}


// SetFetchLimit sets the value of the fetchLimit property.
// The maximum number of objects to include in the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/fetchLimit
func (p_ PHFetchOptions) SetFetchLimit(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFetchLimit:"), value)
}

// A Boolean value that determines whether the fetch result includes all assets from burst photo sequences.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAllBurstAssets
func (p_ PHFetchOptions) IncludeAllBurstAssets() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("includeAllBurstAssets"))
	return rv
}


// SetIncludeAllBurstAssets sets the value of the includeAllBurstAssets property.
// A Boolean value that determines whether the fetch result includes all assets from burst photo sequences.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAllBurstAssets
func (p_ PHFetchOptions) SetIncludeAllBurstAssets(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIncludeAllBurstAssets:"), value)
}

// The set of source types for which to include assets in the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAssetSourceTypes
func (p_ PHFetchOptions) IncludeAssetSourceTypes() PHAssetSourceType {
	rv := objc.Send[PHAssetSourceType](p_.ID, objc.Sel("includeAssetSourceTypes"))
	return rv
}


// SetIncludeAssetSourceTypes sets the value of the includeAssetSourceTypes property.
// The set of source types for which to include assets in the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAssetSourceTypes
func (p_ PHFetchOptions) SetIncludeAssetSourceTypes(value PHAssetSourceType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIncludeAssetSourceTypes:"), value)
}

// A Boolean value that determines whether the fetch result includes assets marked as hidden.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeHiddenAssets
func (p_ PHFetchOptions) IncludeHiddenAssets() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("includeHiddenAssets"))
	return rv
}


// SetIncludeHiddenAssets sets the value of the includeHiddenAssets property.
// A Boolean value that determines whether the fetch result includes assets marked as hidden.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeHiddenAssets
func (p_ PHFetchOptions) SetIncludeHiddenAssets(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIncludeHiddenAssets:"), value)
}

// A predicate that specifies which properties to select results by and that also specifies any constraints on selection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/predicate
func (p_ PHFetchOptions) Predicate() foundation.Predicate {
	rv := objc.Send[foundation.Predicate](p_.ID, objc.Sel("predicate"))
	return rv
}


// SetPredicate sets the value of the predicate property.
// A predicate that specifies which properties to select results by and that also specifies any constraints on selection.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/predicate
func (p_ PHFetchOptions) SetPredicate(value foundation.IPredicate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicate:"), value)
}

// A list of sort descriptors, specifying an order for the fetched objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/sortDescriptors
func (p_ PHFetchOptions) SortDescriptors() []foundation.SortDescriptor {
	rv := objc.Send[[]foundation.SortDescriptor](p_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// SetSortDescriptors sets the value of the sortDescriptors property.
// A list of sort descriptors, specifying an order for the fetched objects.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/sortDescriptors
func (p_ PHFetchOptions) SetSortDescriptors(value []foundation.ISortDescriptor) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setSortDescriptors:"), nsArray)
}

// A Boolean value that determines whether your app receives detailed change information for the objects in the fetch result.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/wantsIncrementalChangeDetails
func (p_ PHFetchOptions) WantsIncrementalChangeDetails() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("wantsIncrementalChangeDetails"))
	return rv
}


// SetWantsIncrementalChangeDetails sets the value of the wantsIncrementalChangeDetails property.
// A Boolean value that determines whether your app receives detailed change information for the objects in the fetch result.

//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/wantsIncrementalChangeDetails
func (p_ PHFetchOptions) SetWantsIncrementalChangeDetails(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWantsIncrementalChangeDetails:"), value)
}

// The unique identifier shared by photo assets from the same burst sequence.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/burstidentifier
func (p_ PHFetchOptions) BurstIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("burstIdentifier"))
	return rv
}


// SetBurstIdentifier sets the value of the burstIdentifier property.
// The unique identifier shared by photo assets from the same burst sequence.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/burstidentifier
func (p_ PHFetchOptions) SetBurstIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBurstIdentifier:"), objc.String(value))
}

// The date and time of the asset’s creation.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/creationdate
func (p_ PHFetchOptions) CreationDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("creationDate"))
	return rv
}


// SetCreationDate sets the value of the creationDate property.
// The date and time of the asset’s creation.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/creationdate
func (p_ PHFetchOptions) SetCreationDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCreationDate:"), value)
}

// The duration, in seconds, of the video asset.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/duration
func (p_ PHFetchOptions) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("duration"))
	return rv
}


// SetDuration sets the value of the duration property.
// The duration, in seconds, of the video asset.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/duration
func (p_ PHFetchOptions) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDuration:"), value)
}

// A Boolean value that indicates whether the user marks the asset as a favorite.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/isfavorite
func (p_ PHFetchOptions) IsFavorite() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFavorite"))
	return rv
}


// SetIsFavorite sets the value of the isFavorite property.
// A Boolean value that indicates whether the user marks the asset as a favorite.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/isfavorite
func (p_ PHFetchOptions) SetIsFavorite(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFavorite:"), value)
}

// A Boolean value that indicates whether the user hides the asset.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/ishidden
func (p_ PHFetchOptions) IsHidden() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isHidden"))
	return rv
}


// SetIsHidden sets the value of the isHidden property.
// A Boolean value that indicates whether the user hides the asset.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/ishidden
func (p_ PHFetchOptions) SetIsHidden(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsHidden:"), value)
}

// The subtypes of the asset, identifying special kinds of assets, such as panoramic photo or high-frame-rate video.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/mediasubtypes
func (p_ PHFetchOptions) MediaSubtypes() PHAssetMediaSubtype {
	rv := objc.Send[PHAssetMediaSubtype](p_.ID, objc.Sel("mediaSubtypes"))
	return rv
}


// SetMediaSubtypes sets the value of the mediaSubtypes property.
// The subtypes of the asset, identifying special kinds of assets, such as panoramic photo or high-frame-rate video.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/mediasubtypes
func (p_ PHFetchOptions) SetMediaSubtypes(value IPHAssetMediaSubtype) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMediaSubtypes:"), value)
}

// The type of the asset, such as video or audio.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/mediatype
func (p_ PHFetchOptions) MediaType() PHAssetMediaType {
	rv := objc.Send[PHAssetMediaType](p_.ID, objc.Sel("mediaType"))
	return rv
}


// SetMediaType sets the value of the mediaType property.
// The type of the asset, such as video or audio.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/mediatype
func (p_ PHFetchOptions) SetMediaType(value PHAssetMediaType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMediaType:"), value)
}

// The date and time of the asset’s last modification.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/modificationdate
func (p_ PHFetchOptions) ModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("modificationDate"))
	return rv
}


// SetModificationDate sets the value of the modificationDate property.
// The date and time of the asset’s last modification.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/modificationdate
func (p_ PHFetchOptions) SetModificationDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModificationDate:"), value)
}

// The height, in pixels, of the asset’s image or video data.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/pixelheight
func (p_ PHFetchOptions) PixelHeight() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pixelHeight"))
	return rv
}


// SetPixelHeight sets the value of the pixelHeight property.
// The height, in pixels, of the asset’s image or video data.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/pixelheight
func (p_ PHFetchOptions) SetPixelHeight(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelHeight:"), value)
}

// The width, in pixels, of the asset’s image or video data.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/pixelwidth
func (p_ PHFetchOptions) PixelWidth() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pixelWidth"))
	return rv
}


// SetPixelWidth sets the value of the pixelWidth property.
// The width, in pixels, of the asset’s image or video data.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/pixelwidth
func (p_ PHFetchOptions) SetPixelWidth(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelWidth:"), value)
}

// The latest creation date among all assets in the asset collection.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/enddate
func (p_ PHFetchOptions) EndDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("endDate"))
	return rv
}


// SetEndDate sets the value of the endDate property.
// The latest creation date among all assets in the asset collection.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/enddate
func (p_ PHFetchOptions) SetEndDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEndDate:"), value)
}

// The estimated number of assets in the asset collection.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/estimatedassetcount
func (p_ PHFetchOptions) EstimatedAssetCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("estimatedAssetCount"))
	return rv
}


// SetEstimatedAssetCount sets the value of the estimatedAssetCount property.
// The estimated number of assets in the asset collection.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/estimatedassetcount
func (p_ PHFetchOptions) SetEstimatedAssetCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEstimatedAssetCount:"), value)
}

// The earliest creation date among all assets in the asset collection.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/startdate
func (p_ PHFetchOptions) StartDate() foundation.Date {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("startDate"))
	return rv
}


// SetStartDate sets the value of the startDate property.
// The earliest creation date among all assets in the asset collection.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/startdate
func (p_ PHFetchOptions) SetStartDate(value foundation.IDate) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStartDate:"), value)
}

// The localized name of the collection.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phcollection/localizedtitle
func (p_ PHFetchOptions) LocalizedTitle() string {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedTitle"))
	return rv
}


// SetLocalizedTitle sets the value of the localizedTitle property.
// The localized name of the collection.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phcollection/localizedtitle
func (p_ PHFetchOptions) SetLocalizedTitle(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedTitle:"), objc.String(value))
}

// A unique string that persistently identifies the object.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobject/localidentifier
func (p_ PHFetchOptions) LocalIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("localIdentifier"))
	return rv
}


// SetLocalIdentifier sets the value of the localIdentifier property.
// A unique string that persistently identifies the object.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phobject/localidentifier
func (p_ PHFetchOptions) SetLocalIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalIdentifier:"), objc.String(value))
}



