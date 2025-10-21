// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHCollection] class.
var (
	PHCollectionClass     _PHCollectionClass
	PHCollectionClassOnce sync.Once
)

func getPHCollectionClass() _PHCollectionClass {
	PHCollectionClassOnce.Do(func() {
		PHCollectionClass = _PHCollectionClass{objc.GetClass("PHCollection")}
	})
	return PHCollectionClass
}

type _PHCollectionClass struct {
	class objc.Class
}

// An interface definition for the [PHCollection] class.
type IPHCollection interface {
	IPHObject
	CanPerformEditOperation(anOperation unsafe.Pointer) bool
}

// The abstract superclass for Photos asset collections and collection lists.
//
// You do not create or work with instances of this class directly. Instead, use one of its two concrete subclasses, or . A object represents a collection of photo or video assets, such as an album, moment, or Shared Photo Stream. A object represents a collection that contains other collections, such as a a folder containing albums or the set of all moments in a calendar year.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollection
type PHCollection struct {
	PHObject
}

// PHCollectionFrom constructs a [PHCollection] from an unsafe.Pointer.
//
// The abstract superclass for Photos asset collections and collection lists.
func PHCollectionFrom(ptr unsafe.Pointer) PHCollection {
	return PHCollection{
		PHObject: PHObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHCollectionClass) Alloc() PHCollection {
	rv := objc.Send[PHCollection](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHCollectionClass) New() PHCollection {
	rv := objc.Send[PHCollection](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHCollection) Init() PHCollection {
	rv := objc.Send[PHCollection](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHCollection) Autorelease() PHCollection {
	rv := objc.Send[PHCollection](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHCollection creates a new PHCollection instance.
func NewPHCollection() PHCollection {
	return getPHCollectionClass().New()
}


// Retrieves collections from the specified collection list.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollection/fetchCollections(in:options:)
func (pc _PHCollectionClass) FetchCollectionsInCollectionListOptions(collectionList unsafe.Pointer, options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchCollectionsInCollectionList:options:"), collectionList, options)
	return rv
}

// Retrieves collections from the root of the photo library’s hierarchy of user-created albums and folders.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollection/fetchTopLevelUserCollections(with:)
func (pc _PHCollectionClass) FetchTopLevelUserCollectionsWithOptions(options unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchTopLevelUserCollectionsWithOptions:"), options)
	return rv
}

// Returns whether the collection supports the specified editing operation.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollection/canPerform(_:)
func (p_ PHCollection) CanPerformEditOperation(anOperation unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canPerformEditOperation:"), anOperation)
	return rv
}

// A Boolean value indicating whether the collection can contain assets.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollection/canContainAssets
func (p_ PHCollection) CanContainAssets() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canContainAssets"))
	return rv
}

// A Boolean value indicating whether the collection can contain other collections.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollection/canContainCollections
func (p_ PHCollection) CanContainCollections() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("canContainCollections"))
	return rv
}

// The localized name of the collection.
//
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHCollection/localizedTitle
func (p_ PHCollection) LocalizedTitle() string {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedTitle"))
	return rv
}



