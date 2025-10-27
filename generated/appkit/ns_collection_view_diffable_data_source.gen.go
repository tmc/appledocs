// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CollectionViewDiffableDataSource] class.
var (
	CollectionViewDiffableDataSourceClass     _CollectionViewDiffableDataSourceClass
	CollectionViewDiffableDataSourceClassOnce sync.Once
)

func getCollectionViewDiffableDataSourceClass() _CollectionViewDiffableDataSourceClass {
	CollectionViewDiffableDataSourceClassOnce.Do(func() {
		CollectionViewDiffableDataSourceClass = _CollectionViewDiffableDataSourceClass{objc.GetClass("NSCollectionViewDiffableDataSource")}
	})
	return CollectionViewDiffableDataSourceClass
}

type _CollectionViewDiffableDataSourceClass struct {
	class objc.Class
}





// An interface definition for the [CollectionViewDiffableDataSource] class.
type ICollectionViewDiffableDataSource interface {
	objectivec.IObject
	

	// properties:
	SupplementaryViewProvider() CollectionViewDiffableDataSourceSupplementaryViewProvider /* not a class type */
	SetSupplementaryViewProvider(value CollectionViewDiffableDataSourceSupplementaryViewProvider /* not a class type */)


	

	// methods:
	ApplySnapshotAnimatingDifferences(snapshot unsafe.Pointer, animatingDifferences bool)
	IndexPathForItemIdentifier(identifier objectivec.IObject) foundation.IndexPath
	ItemIdentifierForIndexPath(indexPath foundation.foundation.INSIndexPath) objectivec.IObject
	Snapshot() unsafe.Pointer


}





// Alloc allocates a new instance without initialization.
func (cc _CollectionViewDiffableDataSourceClass) Alloc() CollectionViewDiffableDataSource {
	rv := objc.Send[CollectionViewDiffableDataSource](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionViewDiffableDataSourceClass) New() CollectionViewDiffableDataSource {
	rv := objc.Send[CollectionViewDiffableDataSource](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewDiffableDataSource) Init() CollectionViewDiffableDataSource {
	rv := objc.Send[CollectionViewDiffableDataSource](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewDiffableDataSource) Autorelease() CollectionViewDiffableDataSource {
	rv := objc.Send[CollectionViewDiffableDataSource](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewDiffableDataSource creates a new CollectionViewDiffableDataSource instance.
func NewCollectionViewDiffableDataSource() CollectionViewDiffableDataSource {
	return getCollectionViewDiffableDataSourceClass().New()
}





// The object you use to manage data and provide items for a collection view.
//
// A object is a specialized type of data source that works together with your collection view object. It provides the behavior you need to manage updates to your collection view’s data and UI in a simple, efficient way. It also conforms to the protocol and provides implementations for all of the protocol’s methods. To fill a collection view with data: Connect a diffable data source to your collection view. Implement an item provider to configure your collection view’s items. Generate the current state of the data. Display the data in the UI. To connect a diffable data source to a collection view, you create the diffable data source using its initializer, passing in the collection view you want to associate with that data source. You also pass in an item provider, where you configure each of your items to determine how to display your data in the UI. Then, you generate the current state of the data and display the data in the UI by constructing and applying a snapshot. For more information, see .


// The object you use to manage data and provide items for a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference
type CollectionViewDiffableDataSource struct {
	objectivec.Object
}

// CollectionViewDiffableDataSourceFrom constructs a [CollectionViewDiffableDataSource] from an unsafe.Pointer.
//
// The object you use to manage data and provide items for a collection view.
func CollectionViewDiffableDataSourceFrom(ptr unsafe.Pointer) CollectionViewDiffableDataSource {
	return CollectionViewDiffableDataSource{objectivec.Object{objc.ID(ptr)}}
}






// Creates a diffable data source with the specified item provider, and connects it to the specified collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/init(collectionView:itemProvider:)
func NewCollectionViewDiffableDataSourceWithCollectionViewItemProvider(collectionView CollectionView /* not a class type */, itemProvider CollectionViewDiffableDataSourceItemProvider /* not a class type */) CollectionViewDiffableDataSource {
	instance := getCollectionViewDiffableDataSourceClass().Alloc()
	rv := objc.Send[CollectionViewDiffableDataSource](instance.ID, objc.Sel("initWithCollectionView:itemProvider:"), collectionView, itemProvider)
	rv.Autorelease()
	return rv
}

















// Updates the UI to reflect the state of the data in the specified snapshot, optionally animating the UI changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/applySnapshot(_:animatingDifferences:)
func (c_ CollectionViewDiffableDataSource) ApplySnapshotAnimatingDifferences(snapshot unsafe.Pointer, animatingDifferences bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("applySnapshot:animatingDifferences:"), snapshot, animatingDifferences)
}


// Returns an index path for the item with the specified identifier in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/indexPath(forItemIdentifier:)
func (c_ CollectionViewDiffableDataSource) IndexPathForItemIdentifier(identifier objectivec.IObject) foundation.IndexPath {
	rv := objc.Send[foundation.IndexPath](c_.ID, objc.Sel("indexPathForItemIdentifier:"), identifier)
	return rv
}


// Returns an identifier for the item at the specified index path in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/itemIdentifier(for:)
func (c_ CollectionViewDiffableDataSource) ItemIdentifierForIndexPath(indexPath foundation.foundation.INSIndexPath) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("itemIdentifierForIndexPath:"), indexPath)
	return rv
}


// Returns a representation of the current state of the data in the collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/snapshot()
func (c_ CollectionViewDiffableDataSource) Snapshot() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("snapshot"))
	return rv
}







// The closure that configures and returns the collection view’s supplementary views, such as headers and footers, from the diffable data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/supplementaryViewProvider
func (c_ CollectionViewDiffableDataSource) SupplementaryViewProvider() CollectionViewDiffableDataSourceSupplementaryViewProvider /* not a class type */ {
	rv := objc.Send[CollectionViewDiffableDataSourceSupplementaryViewProvider](c_.ID, objc.Sel("supplementaryViewProvider"))
	return rv
}


// The closure that configures and returns the collection view’s supplementary views, such as headers and footers, from the diffable data source.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewDiffableDataSourceReference/supplementaryViewProvider
func (c_ CollectionViewDiffableDataSource) SetSupplementaryViewProvider(value CollectionViewDiffableDataSourceSupplementaryViewProvider /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSupplementaryViewProvider:"), value)
}







