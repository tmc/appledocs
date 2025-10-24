// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/cloudkit"
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
	// properties:
	FetchLimit() uint
	SetFetchLimit(value uint)
	IncludeAllBurstAssets() bool
	SetIncludeAllBurstAssets(value bool)
	IncludeAssetSourceTypes() PHAssetSourceType
	SetIncludeAssetSourceTypes(value PHAssetSourceType)
	IncludeHiddenAssets() bool
	SetIncludeHiddenAssets(value bool)
	Predicate() objc.IObject /* cross-framework: Predicate */
	SetPredicate(value objc.IObject /* cross-framework: Predicate */)
	SortDescriptors() []objc.IObject /* cross-framework: SortDescriptor */
	SetSortDescriptors(value []objc.IObject /* cross-framework: SortDescriptor */)
	WantsIncrementalChangeDetails() bool
	SetWantsIncrementalChangeDetails(value bool)
	BurstIdentifier() objc.IObject /* cross-framework: NSString */
	SetBurstIdentifier(value objc.IObject /* cross-framework: NSString */)
	CreationDate() objc.IObject /* cross-framework: Date */
	SetCreationDate(value objc.IObject /* cross-framework: Date */)
	Duration() float64
	SetDuration(value float64)
	IsFavorite() bool
	SetIsFavorite(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	MediaSubtypes() PHAssetMediaSubtype
	SetMediaSubtypes(value PHAssetMediaSubtype)
	MediaType() PHAssetMediaType
	SetMediaType(value PHAssetMediaType)
	ModificationDate() objc.IObject /* cross-framework: Date */
	SetModificationDate(value objc.IObject /* cross-framework: Date */)
	PixelHeight() int
	SetPixelHeight(value int)
	PixelWidth() int
	SetPixelWidth(value int)
	EndDate() objc.IObject /* cross-framework: Date */
	SetEndDate(value objc.IObject /* cross-framework: Date */)
	EstimatedAssetCount() int
	SetEstimatedAssetCount(value int)
	StartDate() objc.IObject /* cross-framework: Date */
	SetStartDate(value objc.IObject /* cross-framework: Date */)
	LocalizedTitle() objc.IObject /* cross-framework: NSString */
	SetLocalizedTitle(value objc.IObject /* cross-framework: NSString */)
	LocalIdentifier() objc.IObject /* cross-framework: NSString */
	SetLocalIdentifier(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// A set of options that affect the filtering, sorting, and management of results that Photos returns when you fetch asset or collection objects.
//
// Using class methods on the , , , and classes to fetch assets or collections produces a object containing the requested objects. The options you specify control which objects the fetch result includes, how those objects are arranged in the fetch result, and how Photos should notify your app of changes to the fetch result. Photos supports only a restricted set of keys for the and properties. The set of available keys depends on which class you’re using to fetch assets or collections. The following table lists the keys supported by each class:


// A set of options that affect the filtering, sorting, and management of results that Photos returns when you fetch asset or collection objects.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/fetchLimit
func (p_ PHFetchOptions) FetchLimit() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("fetchLimit"))
	return rv
}


// The maximum number of objects to include in the fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/fetchLimit
func (p_ PHFetchOptions) SetFetchLimit(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFetchLimit:"), value)
}


// A Boolean value that determines whether the fetch result includes all assets from burst photo sequences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAllBurstAssets
func (p_ PHFetchOptions) IncludeAllBurstAssets() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("includeAllBurstAssets"))
	return rv
}


// A Boolean value that determines whether the fetch result includes all assets from burst photo sequences.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAllBurstAssets
func (p_ PHFetchOptions) SetIncludeAllBurstAssets(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIncludeAllBurstAssets:"), value)
}


// The set of source types for which to include assets in the fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAssetSourceTypes
func (p_ PHFetchOptions) IncludeAssetSourceTypes() PHAssetSourceType {
	rv := objc.Send[PHAssetSourceType](p_.ID, objc.Sel("includeAssetSourceTypes"))
	return rv
}


// The set of source types for which to include assets in the fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeAssetSourceTypes
func (p_ PHFetchOptions) SetIncludeAssetSourceTypes(value PHAssetSourceType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIncludeAssetSourceTypes:"), value)
}


// A Boolean value that determines whether the fetch result includes assets marked as hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeHiddenAssets
func (p_ PHFetchOptions) IncludeHiddenAssets() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("includeHiddenAssets"))
	return rv
}


// A Boolean value that determines whether the fetch result includes assets marked as hidden.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/includeHiddenAssets
func (p_ PHFetchOptions) SetIncludeHiddenAssets(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIncludeHiddenAssets:"), value)
}


// A predicate that specifies which properties to select results by and that also specifies any constraints on selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/predicate
func (p_ PHFetchOptions) Predicate() objc.IObject /* cross-framework: Predicate */ {
	rv := objc.Send[foundation.Predicate](p_.ID, objc.Sel("predicate"))
	return rv
}


// A predicate that specifies which properties to select results by and that also specifies any constraints on selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/predicate
func (p_ PHFetchOptions) SetPredicate(value objc.IObject /* cross-framework: Predicate */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPredicate:"), value)
}


// A list of sort descriptors, specifying an order for the fetched objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/sortDescriptors
func (p_ PHFetchOptions) SortDescriptors() []objc.IObject /* cross-framework: SortDescriptor */ {
	rv := objc.Send[[]cloudkit.SortDescriptor](p_.ID, objc.Sel("sortDescriptors"))
	return rv
}


// A list of sort descriptors, specifying an order for the fetched objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/sortDescriptors
func (p_ PHFetchOptions) SetSortDescriptors(value []objc.IObject /* cross-framework: SortDescriptor */) {
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/wantsIncrementalChangeDetails
func (p_ PHFetchOptions) WantsIncrementalChangeDetails() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("wantsIncrementalChangeDetails"))
	return rv
}


// A Boolean value that determines whether your app receives detailed change information for the objects in the fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHFetchOptions/wantsIncrementalChangeDetails
func (p_ PHFetchOptions) SetWantsIncrementalChangeDetails(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setWantsIncrementalChangeDetails:"), value)
}


// The unique identifier shared by photo assets from the same burst sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/burstidentifier
func (p_ PHFetchOptions) BurstIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("burstIdentifier"))
	return rv
}


// The unique identifier shared by photo assets from the same burst sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/burstidentifier
func (p_ PHFetchOptions) SetBurstIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setBurstIdentifier:"), value)
}


// The date and time of the asset’s creation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/creationdate
func (p_ PHFetchOptions) CreationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("creationDate"))
	return rv
}


// The date and time of the asset’s creation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/creationdate
func (p_ PHFetchOptions) SetCreationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCreationDate:"), value)
}


// The duration, in seconds, of the video asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/duration
func (p_ PHFetchOptions) Duration() float64 {
	rv := objc.Send[float64](p_.ID, objc.Sel("duration"))
	return rv
}


// The duration, in seconds, of the video asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/duration
func (p_ PHFetchOptions) SetDuration(value float64) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDuration:"), value)
}


// A Boolean value that indicates whether the user marks the asset as a favorite.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/isfavorite
func (p_ PHFetchOptions) IsFavorite() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFavorite"))
	return rv
}


// A Boolean value that indicates whether the user marks the asset as a favorite.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/isfavorite
func (p_ PHFetchOptions) SetIsFavorite(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFavorite:"), value)
}


// A Boolean value that indicates whether the user hides the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/ishidden
func (p_ PHFetchOptions) IsHidden() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isHidden"))
	return rv
}


// A Boolean value that indicates whether the user hides the asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/ishidden
func (p_ PHFetchOptions) SetIsHidden(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsHidden:"), value)
}


// The subtypes of the asset, identifying special kinds of assets, such as panoramic photo or high-frame-rate video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/mediasubtypes
func (p_ PHFetchOptions) MediaSubtypes() PHAssetMediaSubtype {
	rv := objc.Send[PHAssetMediaSubtype](p_.ID, objc.Sel("mediaSubtypes"))
	return rv
}


// The subtypes of the asset, identifying special kinds of assets, such as panoramic photo or high-frame-rate video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/mediasubtypes
func (p_ PHFetchOptions) SetMediaSubtypes(value PHAssetMediaSubtype) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMediaSubtypes:"), value)
}


// The type of the asset, such as video or audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/mediatype
func (p_ PHFetchOptions) MediaType() PHAssetMediaType {
	rv := objc.Send[PHAssetMediaType](p_.ID, objc.Sel("mediaType"))
	return rv
}


// The type of the asset, such as video or audio.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/mediatype
func (p_ PHFetchOptions) SetMediaType(value PHAssetMediaType) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setMediaType:"), value)
}


// The date and time of the asset’s last modification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/modificationdate
func (p_ PHFetchOptions) ModificationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("modificationDate"))
	return rv
}


// The date and time of the asset’s last modification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/modificationdate
func (p_ PHFetchOptions) SetModificationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setModificationDate:"), value)
}


// The height, in pixels, of the asset’s image or video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/pixelheight
func (p_ PHFetchOptions) PixelHeight() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pixelHeight"))
	return rv
}


// The height, in pixels, of the asset’s image or video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/pixelheight
func (p_ PHFetchOptions) SetPixelHeight(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelHeight:"), value)
}


// The width, in pixels, of the asset’s image or video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/pixelwidth
func (p_ PHFetchOptions) PixelWidth() int {
	rv := objc.Send[int](p_.ID, objc.Sel("pixelWidth"))
	return rv
}


// The width, in pixels, of the asset’s image or video data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phasset/pixelwidth
func (p_ PHFetchOptions) SetPixelWidth(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPixelWidth:"), value)
}


// The latest creation date among all assets in the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/enddate
func (p_ PHFetchOptions) EndDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("endDate"))
	return rv
}


// The latest creation date among all assets in the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/enddate
func (p_ PHFetchOptions) SetEndDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEndDate:"), value)
}


// The estimated number of assets in the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/estimatedassetcount
func (p_ PHFetchOptions) EstimatedAssetCount() int {
	rv := objc.Send[int](p_.ID, objc.Sel("estimatedAssetCount"))
	return rv
}


// The estimated number of assets in the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/estimatedassetcount
func (p_ PHFetchOptions) SetEstimatedAssetCount(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setEstimatedAssetCount:"), value)
}


// The earliest creation date among all assets in the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/startdate
func (p_ PHFetchOptions) StartDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("startDate"))
	return rv
}


// The earliest creation date among all assets in the asset collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phassetcollection/startdate
func (p_ PHFetchOptions) SetStartDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setStartDate:"), value)
}


// The localized name of the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcollection/localizedtitle
func (p_ PHFetchOptions) LocalizedTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localizedTitle"))
	return rv
}


// The localized name of the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phcollection/localizedtitle
func (p_ PHFetchOptions) SetLocalizedTitle(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedTitle:"), value)
}


// A unique string that persistently identifies the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phobject/localidentifier
func (p_ PHFetchOptions) LocalIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localIdentifier"))
	return rv
}


// A unique string that persistently identifies the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/photos/phobject/localidentifier
func (p_ PHFetchOptions) SetLocalIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalIdentifier:"), value)
}



