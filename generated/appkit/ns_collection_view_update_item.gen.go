// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionViewUpdateItem */


/* debug [class_header]: Header for NSCollectionViewUpdateItem */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionViewUpdateItem */
// An interface definition for the [CollectionViewUpdateItem] class.
type ICollectionViewUpdateItem interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionViewUpdateItem */
	// properties:
	IndexPathAfterUpdate() foundation.IndexPath
	IndexPathBeforeUpdate() foundation.IndexPath
	UpdateAction() CollectionUpdateAction
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionViewUpdateItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionViewUpdateItem */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionViewUpdateItem */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionViewUpdateItem *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionViewUpdateItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionViewUpdateItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionViewUpdateItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionViewUpdateItem */

// The index path of the item after the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem/indexPathAfterUpdate
func (c_ CollectionViewUpdateItem) IndexPathAfterUpdate() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](c_.ID, objc.Sel("indexPathAfterUpdate"))
	return rv
}/* debug [instance_properties/getter]: indexPathAfterUpdate */


// The index path of the item before the update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem/indexPathBeforeUpdate
func (c_ CollectionViewUpdateItem) IndexPathBeforeUpdate() foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](c_.ID, objc.Sel("indexPathBeforeUpdate"))
	return rv
}/* debug [instance_properties/getter]: indexPathBeforeUpdate */


// The action being performed on the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewUpdateItem/updateAction
func (c_ CollectionViewUpdateItem) UpdateAction() CollectionUpdateAction {
	rv := objc.Send[CollectionUpdateAction](c_.ID, objc.Sel("updateAction"))
	return rv
}/* debug [instance_properties/getter]: updateAction */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionViewUpdateItem */



