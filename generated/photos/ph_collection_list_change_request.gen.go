// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHCollectionListChangeRequest] class.
var (
	PHCollectionListChangeRequestClass     _PHCollectionListChangeRequestClass
	PHCollectionListChangeRequestClassOnce sync.Once
)

func getPHCollectionListChangeRequestClass() _PHCollectionListChangeRequestClass {
	PHCollectionListChangeRequestClassOnce.Do(func() {
		PHCollectionListChangeRequestClass = _PHCollectionListChangeRequestClass{objc.GetClass("PHCollectionListChangeRequest")}
	})
	return PHCollectionListChangeRequestClass
}

type _PHCollectionListChangeRequestClass struct {
	class objc.Class
}

// An interface definition for the [PHCollectionListChangeRequest] class.
type IPHCollectionListChangeRequest interface {
	IPHChangeRequest
	AddChildCollections(collections objectivec.IObject)
	InsertChildCollectionsAtIndexes(collections objectivec.IObject, indexes foundation.IIndexSet)
	ReplaceChildCollectionsAtIndexesWithChildCollections(indexes foundation.IIndexSet, collections objectivec.IObject)
}

// A request to create, delete, or modify a Photos collection list, for use in a photo library change block.
//
// You use the class to request changes for objects. To make changes to collection lists (such as folders containing user-created albums) in the Photos library, create a change request using the appropriate class method for the change you want to perform. Call the method to create a new asset collection. Call the method to delete existing asset collections. Call the or method to modify a collection’s metadata or its list of child collections. Before creating a change request, use the method to verify that the collection allows the edit operation you’re requesting. If you attempt to perform an unsupported edit operation, Photos throws an exception. A change request for creating or modifying a collection list works like a mutable version of the collection list object. Use the change request’s properties and instance methods to request changes to the collection list itself. For example, the following code removes an album from a folder. After Photos runs the change block and calls your completion handler, the collection list’s state reflects the changes you requested in the block. If you create or use a change request object outside a photo library change block, Photos raises an Objective-C exception. For details on change blocks, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest
type PHCollectionListChangeRequest struct {
	PHChangeRequest
}

// PHCollectionListChangeRequestFrom constructs a [PHCollectionListChangeRequest] from an unsafe.Pointer.
//
// A request to create, delete, or modify a Photos collection list, for use in a photo library change block.
func PHCollectionListChangeRequestFrom(ptr unsafe.Pointer) PHCollectionListChangeRequest {
	return PHCollectionListChangeRequest{
		PHChangeRequest: PHChangeRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHCollectionListChangeRequestClass) Alloc() PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHCollectionListChangeRequestClass) New() PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHCollectionListChangeRequest) Init() PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHCollectionListChangeRequest) Autorelease() PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHCollectionListChangeRequest creates a new PHCollectionListChangeRequest instance.
func NewPHCollectionListChangeRequest() PHCollectionListChangeRequest {
	return getPHCollectionListChangeRequestClass().New()
}




// Creates a request for modifying the specified collection list.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/init(for:)
func NewPHCollectionListChangeRequestForCollectionList(collectionList IPHCollectionList) PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](objc.ID(getPHCollectionListChangeRequestClass().class), objc.Sel("changeRequestForCollectionList:"), collectionList)
	return rv
}



// Creates a request for modifying the specified collection list, with a fetch result for tracking changes.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/init(for:childCollections:)
func NewPHCollectionListChangeRequestForCollectionListChildCollections(collectionList IPHCollectionList, childCollections unsafe.Pointer) PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](objc.ID(getPHCollectionListChangeRequestClass().class), objc.Sel("changeRequestForCollectionList:childCollections:"), collectionList, childCollections)
	return rv
}



// Creates a request to add, remove, or rearrange child collections in the top-level collection list.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/init(forTopLevelCollectionListUserCollections:)
func NewPHCollectionListChangeRequestForTopLevelCollectionListUserCollections(childCollections unsafe.Pointer) PHCollectionListChangeRequest {
	rv := objc.Send[PHCollectionListChangeRequest](objc.ID(getPHCollectionListChangeRequestClass().class), objc.Sel("changeRequestForTopLevelCollectionListUserCollections:"), childCollections)
	return rv
}


// Creates a request for adding a new collection list to the Photos library.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/creationRequestForCollectionList(withTitle:)
func (pc _PHCollectionListChangeRequestClass) CreationRequestForCollectionListWithTitle(title appkit.string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("creationRequestForCollectionListWithTitle:"), title)
	return rv
}

// Requests to delete the specified asset collections.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/deleteCollectionLists(_:)
func (pc _PHCollectionListChangeRequestClass) DeleteCollectionLists(collectionLists objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("deleteCollectionLists:"), collectionLists)
}

// Creates a request for modifying the specified collection list.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/init(for:)
func (pc _PHCollectionListChangeRequestClass) ChangeRequestForCollectionList(collectionList IPHCollectionList) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("changeRequestForCollectionList:"), collectionList)
	return rv
}

// Creates a request for modifying the specified collection list, with a fetch result for tracking changes.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/init(for:childCollections:)
func (pc _PHCollectionListChangeRequestClass) ChangeRequestForCollectionListChildCollections(collectionList IPHCollectionList, childCollections unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("changeRequestForCollectionList:childCollections:"), collectionList, childCollections)
	return rv
}

// Creates a request to add, remove, or rearrange child collections in the top-level collection list.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/init(forTopLevelCollectionListUserCollections:)
func (pc _PHCollectionListChangeRequestClass) ChangeRequestForTopLevelCollectionListUserCollections(childCollections unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("changeRequestForTopLevelCollectionListUserCollections:"), childCollections)
	return rv
}

// Adds the specified collections as children of the collection list.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/addChildCollections(_:)
func (p_ PHCollectionListChangeRequest) AddChildCollections(collections objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addChildCollections:"), collections)
}

// Inserts the specified collections into the collection list at the specified indexes.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/insertChildCollections(_:at:)
func (p_ PHCollectionListChangeRequest) InsertChildCollectionsAtIndexes(collections objectivec.IObject, indexes foundation.IIndexSet) {
	objc.Send[objc.ID](p_.ID, objc.Sel("insertChildCollections:atIndexes:"), collections, indexes)
}

// Replaces the child collections at the specified indexes in the collection list with the specified collections.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/replaceChildCollections(at:withChildCollections:)
func (p_ PHCollectionListChangeRequest) ReplaceChildCollectionsAtIndexesWithChildCollections(indexes foundation.IIndexSet, collections objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("replaceChildCollectionsAtIndexes:withChildCollections:"), indexes, collections)
}

// A placeholder object for the collection list that the change request creates.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollectionListChangeRequest/placeholderForCreatedCollectionList
func (p_ PHCollectionListChangeRequest) PlaceholderForCreatedCollectionList() PHObjectPlaceholder {
	rv := objc.Send[PHObjectPlaceholder](p_.ID, objc.Sel("placeholderForCreatedCollectionList"))
	return rv
}

// The displayed name of the collection list.
//
// [Full Topic]: https://developer.apple.com/documentation/photos/phcollectionlistchangerequest/title
func (p_ PHCollectionListChangeRequest) Title() appkit.string {
	rv := objc.Send[appkit.string](p_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The displayed name of the collection list.

//
// [Full Topic]: https://developer.apple.com/documentation/photos/phcollectionlistchangerequest/title
func (p_ PHCollectionListChangeRequest) SetTitle(value appkit.string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTitle:"), value)
}


