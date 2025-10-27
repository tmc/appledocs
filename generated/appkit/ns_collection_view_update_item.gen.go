// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CollectionViewUpdateItem] class.
var (
	CollectionViewUpdateItemClass     _CollectionViewUpdateItemClass
	CollectionViewUpdateItemClassOnce sync.Once
)

func getCollectionViewUpdateItemClass() _CollectionViewUpdateItemClass {
	CollectionViewUpdateItemClassOnce.Do(func() {
		CollectionViewUpdateItemClass = _CollectionViewUpdateItemClass{objc.GetClass("NSCollectionViewUpdateItem")}
	})
	return CollectionViewUpdateItemClass
}

type _CollectionViewUpdateItemClass struct {
	class objc.Class
}





// An interface definition for the [CollectionViewUpdateItem] class.
type ICollectionViewUpdateItem interface {
	objectivec.IObject
	

	// properties:
	IndexPathAfterUpdate() foundation.foundation.INSIndexPath
	IndexPathBeforeUpdate() foundation.foundation.INSIndexPath
	UpdateAction() CollectionUpdateAction


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CollectionViewUpdateItemClass) Alloc() CollectionViewUpdateItem {
	rv := objc.Send[CollectionViewUpdateItem](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionViewUpdateItemClass) New() CollectionViewUpdateItem {
	rv := objc.Send[CollectionViewUpdateItem](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewUpdateItem) Init() CollectionViewUpdateItem {
	rv := objc.Send[CollectionViewUpdateItem](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewUpdateItem) Autorelease() CollectionViewUpdateItem {
	rv := objc.Send[CollectionViewUpdateItem](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewUpdateItem creates a new CollectionViewUpdateItem instance.
func NewCollectionViewUpdateItem() CollectionViewUpdateItem {
	return getCollectionViewUpdateItemClass().New()
}





// A description of a single change to make to an item in a collection view.
//
// You do not create instances of this class directly. When updating its content, the collection view object creates them and passes them to the layout object’s method, which can use them to prepare for the upcoming changes.


// A description of a single change to make to an item in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem
type CollectionViewUpdateItem struct {
	objectivec.Object
}

// CollectionViewUpdateItemFrom constructs a [CollectionViewUpdateItem] from an unsafe.Pointer.
//
// A description of a single change to make to an item in a collection view.
func CollectionViewUpdateItemFrom(ptr unsafe.Pointer) CollectionViewUpdateItem {
	return CollectionViewUpdateItem{objectivec.Object{objc.ID(ptr)}}
}

























// The index path of the item after the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem/indexPathAfterUpdate
func (c_ CollectionViewUpdateItem) IndexPathAfterUpdate() foundation.foundation.INSIndexPath {
	rv := objc.Send[foundation.NSIndexPath](c_.ID, objc.Sel("indexPathAfterUpdate"))
	return rv
}


// The index path of the item before the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem/indexPathBeforeUpdate
func (c_ CollectionViewUpdateItem) IndexPathBeforeUpdate() foundation.foundation.INSIndexPath {
	rv := objc.Send[foundation.NSIndexPath](c_.ID, objc.Sel("indexPathBeforeUpdate"))
	return rv
}


// The action being performed on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem/updateAction
func (c_ CollectionViewUpdateItem) UpdateAction() CollectionUpdateAction {
	rv := objc.Send[CollectionUpdateAction](c_.ID, objc.Sel("updateAction"))
	return rv
}








