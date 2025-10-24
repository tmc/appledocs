// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PHCollectionList] class.
var (
	PHCollectionListClass     _PHCollectionListClass
	PHCollectionListClassOnce sync.Once
)

func getPHCollectionListClass() _PHCollectionListClass {
	PHCollectionListClassOnce.Do(func() {
		PHCollectionListClass = _PHCollectionListClass{objc.GetClass("PHCollectionList")}
	})
	return PHCollectionListClass
}

type _PHCollectionListClass struct {
	class objc.Class
}

// An interface definition for the [PHCollectionList] class.
type IPHCollectionList interface {
	IPHCollection
	// properties:
	CollectionListSubtype() PHCollectionListSubtype
	CollectionListType() PHCollectionListType
	EndDate() objc.IObject /* cross-framework: NSDate */
	LocalizedLocationNames() []string
	StartDate() objc.IObject /* cross-framework: NSDate */
	// methods:
}

// A group containing Photos asset collections, such as Moments, Years, or folders of user-created albums.
//
// In the Photos framework, collection objects (including asset collections) do not directly reference their member objects, and there are no other objects that directly reference collection objects. To retrieve the members of a collection list, fetch them with a class method such as . To find objects at the root of the collection list hierarchy (such as album folders with no parent folders), use the method. Like assets and asset collections, collection lists are immutable. To create, rename, or delete collection lists, or to add, remove, or rearrange members in a collection list, create a object within a photo library change block. For details on using change requests and change blocks to update the photo library, see .


// A group containing Photos asset collections, such as Moments, Years, or folders of user-created albums.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList
type PHCollectionList struct {
	PHCollection
}

// PHCollectionListFrom constructs a [PHCollectionList] from an unsafe.Pointer.
//
// A group containing Photos asset collections, such as Moments, Years, or folders of user-created albums.
func PHCollectionListFrom(ptr unsafe.Pointer) PHCollectionList {
	return PHCollectionList{
		PHCollection: PHCollectionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHCollectionListClass) Alloc() PHCollectionList {
	rv := objc.Send[PHCollectionList](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHCollectionListClass) New() PHCollectionList {
	rv := objc.Send[PHCollectionList](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHCollectionList) Init() PHCollectionList {
	rv := objc.Send[PHCollectionList](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHCollectionList) Autorelease() PHCollectionList {
	rv := objc.Send[PHCollectionList](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHCollectionList creates a new PHCollectionList instance.
func NewPHCollectionList() PHCollectionList {
	return getPHCollectionListClass().New()
}



// Retrieves collection lists of the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/fetchCollectionLists(with:subtype:options:)
func (pc _PHCollectionListClass) FetchCollectionListsWithTypeSubtypeOptions(collectionListType PHCollectionListType, subtype PHCollectionListSubtype, options IPHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchCollectionListsWithType:subtype:options:"), collectionListType, subtype, options)
	return rv
}


// Retrieves collection lists with the specified local-device-specific unique identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/fetchCollectionLists(withLocalIdentifiers:options:)
func (pc _PHCollectionListClass) FetchCollectionListsWithLocalIdentifiersOptions(identifiers []string, options IPHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchCollectionListsWithLocalIdentifiers:options:"), identifiers, options)
	return rv
}


// Retrieves collection lists that contain the specified collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/fetchCollectionListsContaining(_:options:)
func (pc _PHCollectionListClass) FetchCollectionListsContainingCollectionOptions(collection IPHCollection, options IPHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchCollectionListsContainingCollection:options:"), collection, options)
	return rv
}


// Retrieves collection lists of the specified moment list type containing the specified moment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/fetchMomentLists(with:containingMoment:options:)
func (pc _PHCollectionListClass) FetchMomentListsWithSubtypeContainingMomentOptions(momentListSubtype PHCollectionListSubtype, moment IPHAssetCollection, options IPHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchMomentListsWithSubtype:containingMoment:options:"), momentListSubtype, moment, options)
	return rv
}


// Retrieves collection lists of the specified moment list type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/fetchMomentLists(with:options:)
func (pc _PHCollectionListClass) FetchMomentListsWithSubtypeOptions(momentListSubtype PHCollectionListSubtype, options IPHFetchOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchMomentListsWithSubtype:options:"), momentListSubtype, options)
	return rv
}


// Creates a temporary collection list that contains the specified asset collections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/transientCollectionList(with:title:)
func (pc _PHCollectionListClass) TransientCollectionListWithCollectionsTitle(collections []IPHCollection, title objc.IObject /* cross-framework: NSString */) PHCollectionList {
	rv := objc.Send[PHCollectionList](objc.ID(pc.class), objc.Sel("transientCollectionListWithCollections:title:"), collections, title)
	return rv
}


// Creates a temporary collection list containing the asset collections in the specified fetch result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/transientCollectionList(withCollectionsFetchResult:title:)
func (pc _PHCollectionListClass) TransientCollectionListWithCollectionsFetchResultTitle(fetchResult unsafe.Pointer, title objc.IObject /* cross-framework: NSString */) PHCollectionList {
	rv := objc.Send[PHCollectionList](objc.ID(pc.class), objc.Sel("transientCollectionListWithCollectionsFetchResult:title:"), fetchResult, title)
	return rv
}


// The type of asset collection grouping the collection list represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/collectionListSubtype
func (p_ PHCollectionList) CollectionListSubtype() PHCollectionListSubtype {
	rv := objc.Send[PHCollectionListSubtype](p_.ID, objc.Sel("collectionListSubtype"))
	return rv
}


// The type of asset collection group that the collection list represents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/collectionListType
func (p_ PHCollectionList) CollectionListType() PHCollectionListType {
	rv := objc.Send[PHCollectionListType](p_.ID, objc.Sel("collectionListType"))
	return rv
}


// The latest creation date among all assets in the collection list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/endDate
func (p_ PHCollectionList) EndDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("endDate"))
	return rv
}


// The names of locations grouped by the collection (an array of objects).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/localizedLocationNames
func (p_ PHCollectionList) LocalizedLocationNames() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("localizedLocationNames"))
	return rv
}


// The earliest creation date among all assets in the collection list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionList/startDate
func (p_ PHCollectionList) StartDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("startDate"))
	return rv
}


