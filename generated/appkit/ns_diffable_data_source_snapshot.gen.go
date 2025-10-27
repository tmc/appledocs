// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [DiffableDataSourceSnapshot] class.
var (
	DiffableDataSourceSnapshotClass     _DiffableDataSourceSnapshotClass
	DiffableDataSourceSnapshotClassOnce sync.Once
)

func getDiffableDataSourceSnapshotClass() _DiffableDataSourceSnapshotClass {
	DiffableDataSourceSnapshotClassOnce.Do(func() {
		DiffableDataSourceSnapshotClass = _DiffableDataSourceSnapshotClass{objc.GetClass("NSDiffableDataSourceSnapshot")}
	})
	return DiffableDataSourceSnapshotClass
}

type _DiffableDataSourceSnapshotClass struct {
	class objc.Class
}





// An interface definition for the [DiffableDataSourceSnapshot] class.
type IDiffableDataSourceSnapshot interface {
	objectivec.IObject
	

	// properties:
	ItemIdentifiers() []objc.ID
	NumberOfItems() int
	NumberOfSections() int
	SectionIdentifiers() []objc.ID


	

	// methods:
	AppendItemsWithIdentifiers(identifiers []objc.ID)
	AppendItemsWithIdentifiersIntoSectionWithIdentifier(identifiers []objc.ID, sectionIdentifier objectivec.IObject)
	AppendSectionsWithIdentifiers(sectionIdentifiers []objc.ID)
	DeleteAllItems()
	DeleteItemsWithIdentifiers(identifiers []objc.ID)
	DeleteSectionsWithIdentifiers(sectionIdentifiers []objc.ID)
	IndexOfItemIdentifier(itemIdentifier objectivec.IObject) int
	IndexOfSectionIdentifier(sectionIdentifier objectivec.IObject) int
	InsertItemsWithIdentifiersAfterItemWithIdentifier(identifiers []objc.ID, itemIdentifier objectivec.IObject)
	InsertItemsWithIdentifiersBeforeItemWithIdentifier(identifiers []objc.ID, itemIdentifier objectivec.IObject)
	InsertSectionsWithIdentifiersAfterSectionWithIdentifier(sectionIdentifiers []objc.ID, toSectionIdentifier objectivec.IObject)
	InsertSectionsWithIdentifiersBeforeSectionWithIdentifier(sectionIdentifiers []objc.ID, toSectionIdentifier objectivec.IObject)
	ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objectivec.IObject) []objc.ID
	MoveItemWithIdentifierAfterItemWithIdentifier(fromIdentifier objectivec.IObject, toIdentifier objectivec.IObject)
	MoveItemWithIdentifierBeforeItemWithIdentifier(fromIdentifier objectivec.IObject, toIdentifier objectivec.IObject)
	MoveSectionWithIdentifierAfterSectionWithIdentifier(fromSectionIdentifier objectivec.IObject, toSectionIdentifier objectivec.IObject)
	MoveSectionWithIdentifierBeforeSectionWithIdentifier(fromSectionIdentifier objectivec.IObject, toSectionIdentifier objectivec.IObject)
	NumberOfItemsInSection(sectionIdentifier objectivec.IObject) int
	ReloadItemsWithIdentifiers(identifiers []objc.ID)
	ReloadSectionsWithIdentifiers(sectionIdentifiers []objc.ID)
	SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objectivec.IObject) objectivec.IObject


}





// Alloc allocates a new instance without initialization.
func (dc _DiffableDataSourceSnapshotClass) Alloc() DiffableDataSourceSnapshot {
	rv := objc.Send[DiffableDataSourceSnapshot](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DiffableDataSourceSnapshotClass) New() DiffableDataSourceSnapshot {
	rv := objc.Send[DiffableDataSourceSnapshot](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DiffableDataSourceSnapshot) Init() DiffableDataSourceSnapshot {
	rv := objc.Send[DiffableDataSourceSnapshot](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DiffableDataSourceSnapshot) Autorelease() DiffableDataSourceSnapshot {
	rv := objc.Send[DiffableDataSourceSnapshot](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDiffableDataSourceSnapshot creates a new DiffableDataSourceSnapshot instance.
func NewDiffableDataSourceSnapshot() DiffableDataSourceSnapshot {
	return getDiffableDataSourceSnapshotClass().New()
}





// A representation of the state of the data in a view at a specific point in time.
//
// Diffable data sources use to provide data for collection views and table views. Through a snapshot, you set up the initial state of the data that displays in a view, and later update that data. The data in a snapshot is made up of the sections and items you want to display, in the specific order you want to display them. You configure what to display by adding, deleting, or moving the sections and items. To display data in a view using a snapshot: Create a snapshot and populate it with the state of the data you want to display. Apply the snapshot to reflect the changes in the UI. You can create and configure a snapshot in one of these ways: Create an empty snapshot, then append sections and items to it. Get the current snapshot by calling the diffable data source’s method, then modify that snapshot to reflect the new state of the data that you want to display. For example, the following code creates an empty snapshot, and populates it with a single section with three items. Then, it applies the snapshot, animating the UI updates between the previous state and the new state represented in the snapshot. For more information, see the diffable data source types:


// A representation of the state of the data in a view at a specific point in time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference
type DiffableDataSourceSnapshot struct {
	objectivec.Object
}

// DiffableDataSourceSnapshotFrom constructs a [DiffableDataSourceSnapshot] from an unsafe.Pointer.
//
// A representation of the state of the data in a view at a specific point in time.
func DiffableDataSourceSnapshotFrom(ptr unsafe.Pointer) DiffableDataSourceSnapshot {
	return DiffableDataSourceSnapshot{objectivec.Object{objc.ID(ptr)}}
}




















// Adds the items with the specified identifiers to the last section of the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/appendItems(withIdentifiers:)
func (d_ DiffableDataSourceSnapshot) AppendItemsWithIdentifiers(identifiers []objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("appendItemsWithIdentifiers:"), identifiers)
}


// Adds the items with the specified identifiers to the specified section of the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/appendItems(withIdentifiers:intoSectionWithIdentifier:)
func (d_ DiffableDataSourceSnapshot) AppendItemsWithIdentifiersIntoSectionWithIdentifier(identifiers []objc.ID, sectionIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("appendItemsWithIdentifiers:intoSectionWithIdentifier:"), identifiers, sectionIdentifier)
}


// Adds the sections with the specified identifiers to the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/appendSections(withIdentifiers:)
func (d_ DiffableDataSourceSnapshot) AppendSectionsWithIdentifiers(sectionIdentifiers []objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("appendSectionsWithIdentifiers:"), sectionIdentifiers)
}


// Deletes all of the items from the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/deleteAllItems()
func (d_ DiffableDataSourceSnapshot) DeleteAllItems() {
	objc.Send[objc.ID](d_.ID, objc.Sel("deleteAllItems"))
}


// Deletes the items with the specified identifiers from the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/deleteItems(withIdentifiers:)
func (d_ DiffableDataSourceSnapshot) DeleteItemsWithIdentifiers(identifiers []objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("deleteItemsWithIdentifiers:"), identifiers)
}


// Deletes the sections with the specified identifiers from the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/deleteSections(withIdentifiers:)
func (d_ DiffableDataSourceSnapshot) DeleteSectionsWithIdentifiers(sectionIdentifiers []objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("deleteSectionsWithIdentifiers:"), sectionIdentifiers)
}


// Returns the index of the item in the snapshot with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/index(ofItemIdentifier:)
func (d_ DiffableDataSourceSnapshot) IndexOfItemIdentifier(itemIdentifier objectivec.IObject) int {
	rv := objc.Send[int](d_.ID, objc.Sel("indexOfItemIdentifier:"), itemIdentifier)
	return rv
}


// Returns the index of the section of the snapshot with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/index(ofSectionIdentifier:)
func (d_ DiffableDataSourceSnapshot) IndexOfSectionIdentifier(sectionIdentifier objectivec.IObject) int {
	rv := objc.Send[int](d_.ID, objc.Sel("indexOfSectionIdentifier:"), sectionIdentifier)
	return rv
}


// Inserts the provided items immediately after the item with the specified identifier in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/insertItems(withIdentifiers:afterItemWithIdentifier:)
func (d_ DiffableDataSourceSnapshot) InsertItemsWithIdentifiersAfterItemWithIdentifier(identifiers []objc.ID, itemIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("insertItemsWithIdentifiers:afterItemWithIdentifier:"), identifiers, itemIdentifier)
}


// Inserts the provided items immediately before the item with the specified identifier in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/insertItems(withIdentifiers:beforeItemWithIdentifier:)
func (d_ DiffableDataSourceSnapshot) InsertItemsWithIdentifiersBeforeItemWithIdentifier(identifiers []objc.ID, itemIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("insertItemsWithIdentifiers:beforeItemWithIdentifier:"), identifiers, itemIdentifier)
}


// Inserts the provided sections immediately after the section with the specified identifier in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/insertSections(withIdentifiers:afterSectionWithIdentifier:)
func (d_ DiffableDataSourceSnapshot) InsertSectionsWithIdentifiersAfterSectionWithIdentifier(sectionIdentifiers []objc.ID, toSectionIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("insertSectionsWithIdentifiers:afterSectionWithIdentifier:"), sectionIdentifiers, toSectionIdentifier)
}


// Inserts the provided sections immediately before the section with the specified identifier in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/insertSections(withIdentifiers:beforeSectionWithIdentifier:)
func (d_ DiffableDataSourceSnapshot) InsertSectionsWithIdentifiersBeforeSectionWithIdentifier(sectionIdentifiers []objc.ID, toSectionIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("insertSectionsWithIdentifiers:beforeSectionWithIdentifier:"), sectionIdentifiers, toSectionIdentifier)
}


// Returns the identifiers of all of the items in the specified section of the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/itemIdentifiersInSection(withIdentifier:)
func (d_ DiffableDataSourceSnapshot) ItemIdentifiersInSectionWithIdentifier(sectionIdentifier objectivec.IObject) []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("itemIdentifiersInSectionWithIdentifier:"), sectionIdentifier)
	return rv
}


// Moves the item from its current position in the snapshot to the position immediately after the specified item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/moveItem(withIdentifier:afterItemWithIdentifier:)
func (d_ DiffableDataSourceSnapshot) MoveItemWithIdentifierAfterItemWithIdentifier(fromIdentifier objectivec.IObject, toIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveItemWithIdentifier:afterItemWithIdentifier:"), fromIdentifier, toIdentifier)
}


// Moves the item from its current position in the snapshot to the position immediately before the specified item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/moveItem(withIdentifier:beforeItemWithIdentifier:)
func (d_ DiffableDataSourceSnapshot) MoveItemWithIdentifierBeforeItemWithIdentifier(fromIdentifier objectivec.IObject, toIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveItemWithIdentifier:beforeItemWithIdentifier:"), fromIdentifier, toIdentifier)
}


// Moves the section from its current position in the snapshot to the position immediately after the specified section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/moveSection(withIdentifier:afterSectionWithIdentifier:)
func (d_ DiffableDataSourceSnapshot) MoveSectionWithIdentifierAfterSectionWithIdentifier(fromSectionIdentifier objectivec.IObject, toSectionIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveSectionWithIdentifier:afterSectionWithIdentifier:"), fromSectionIdentifier, toSectionIdentifier)
}


// Moves the section from its current position in the snapshot to the position immediately before the specified section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/moveSection(withIdentifier:beforeSectionWithIdentifier:)
func (d_ DiffableDataSourceSnapshot) MoveSectionWithIdentifierBeforeSectionWithIdentifier(fromSectionIdentifier objectivec.IObject, toSectionIdentifier objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveSectionWithIdentifier:beforeSectionWithIdentifier:"), fromSectionIdentifier, toSectionIdentifier)
}


// Returns the number of items in the specified section of the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/numberOfItems(inSection:)
func (d_ DiffableDataSourceSnapshot) NumberOfItemsInSection(sectionIdentifier objectivec.IObject) int {
	rv := objc.Send[int](d_.ID, objc.Sel("numberOfItemsInSection:"), sectionIdentifier)
	return rv
}


// Reloads the data within the specified items in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/reloadItems(withIdentifiers:)
func (d_ DiffableDataSourceSnapshot) ReloadItemsWithIdentifiers(identifiers []objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("reloadItemsWithIdentifiers:"), identifiers)
}


// Reloads the data within the specified sections of the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/reloadSections(withIdentifiers:)
func (d_ DiffableDataSourceSnapshot) ReloadSectionsWithIdentifiers(sectionIdentifiers []objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("reloadSectionsWithIdentifiers:"), sectionIdentifiers)
}


// Returns the identifier of the section containing the specified item in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/sectionIdentifier(forSectionContainingItemIdentifier:)
func (d_ DiffableDataSourceSnapshot) SectionIdentifierForSectionContainingItemIdentifier(itemIdentifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](d_.ID, objc.Sel("sectionIdentifierForSectionContainingItemIdentifier:"), itemIdentifier)
	return rv
}







// The identifiers of all of the items in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/itemIdentifiers
func (d_ DiffableDataSourceSnapshot) ItemIdentifiers() []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("itemIdentifiers"))
	return rv
}


// The number of items in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/numberOfItems
func (d_ DiffableDataSourceSnapshot) NumberOfItems() int {
	rv := objc.Send[int](d_.ID, objc.Sel("numberOfItems"))
	return rv
}


// The number of sections in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/numberOfSections
func (d_ DiffableDataSourceSnapshot) NumberOfSections() int {
	rv := objc.Send[int](d_.ID, objc.Sel("numberOfSections"))
	return rv
}


// The identifiers of all of the sections in the snapshot.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDiffableDataSourceSnapshotReference/sectionIdentifiers
func (d_ DiffableDataSourceSnapshot) SectionIdentifiers() []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("sectionIdentifiers"))
	return rv
}








