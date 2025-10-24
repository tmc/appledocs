// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PCollectionViewDataSource is the NSCollectionViewDataSource protocol interface.
//
// A set of methods that a data source object implements to provide the information and view objects that a collection view requires to present content.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSCollectionViewDataSource
type PCollectionViewDataSource interface {
	// Required methods
	CollectionViewNumberOfItemsInSection(collectionView objc.IObject /* cross-framework: CollectionView */, section int) int
	CollectionViewItemForRepresentedObjectAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, indexPath foundation.IndexPath) CollectionViewItem
	// Optional methods
	CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, kind objc.IObject /* cross-framework: CollectionViewSupplementaryElementKind */, indexPath foundation.IndexPath) View
	HasCollectionViewViewForSupplementaryElementOfKindAtIndexPath() bool
	NumberOfSectionsInCollectionView(collectionView objc.IObject /* cross-framework: CollectionView */) int
	HasNumberOfSectionsInCollectionView() bool
}

// CollectionViewDataSource is a delegate implementation builder for the PCollectionViewDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CollectionViewDataSource struct {
	_CollectionViewViewForSupplementaryElementOfKindAtIndexPath func(collectionView objc.IObject /* cross-framework: CollectionView */, kind objc.IObject /* cross-framework: CollectionViewSupplementaryElementKind */, indexPath foundation.IndexPath) View
	_NumberOfSectionsInCollectionView func(collectionView objc.IObject /* cross-framework: CollectionView */) int
	_CollectionViewNumberOfItemsInSection func(collectionView objc.IObject /* cross-framework: CollectionView */, section int) int
	_CollectionViewItemForRepresentedObjectAtIndexPath func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPath foundation.IndexPath) CollectionViewItem
}

// SetCollectionViewViewForSupplementaryElementOfKindAtIndexPath sets the handler for the CollectionViewViewForSupplementaryElementOfKindAtIndexPath delegate method.
//
// Asks your data source object to provide the supplementary view at the specified location in a section of the collection view.
func (d *CollectionViewDataSource) SetCollectionViewViewForSupplementaryElementOfKindAtIndexPath(f func(collectionView objc.IObject /* cross-framework: CollectionView */, kind objc.IObject /* cross-framework: CollectionViewSupplementaryElementKind */, indexPath foundation.IndexPath) View) {
	d._CollectionViewViewForSupplementaryElementOfKindAtIndexPath = f
}

// SetNumberOfSectionsInCollectionView sets the handler for the NumberOfSectionsInCollectionView delegate method.
//
// Asks your data source object to provide the total number of sections.
func (d *CollectionViewDataSource) SetNumberOfSectionsInCollectionView(f func(collectionView objc.IObject /* cross-framework: CollectionView */) int) {
	d._NumberOfSectionsInCollectionView = f
}

// SetCollectionViewNumberOfItemsInSection sets the handler for the CollectionViewNumberOfItemsInSection delegate method.
//
// Asks your data source object to provide the number of items in the specified section.
func (d *CollectionViewDataSource) SetCollectionViewNumberOfItemsInSection(f func(collectionView objc.IObject /* cross-framework: CollectionView */, section int) int) {
	d._CollectionViewNumberOfItemsInSection = f
}

// SetCollectionViewItemForRepresentedObjectAtIndexPath sets the handler for the CollectionViewItemForRepresentedObjectAtIndexPath delegate method.
//
// Asks your data source object to provide the item at the specified location in the collection view.
func (d *CollectionViewDataSource) SetCollectionViewItemForRepresentedObjectAtIndexPath(f func(collectionView objc.IObject /* cross-framework: CollectionView */, indexPath foundation.IndexPath) CollectionViewItem) {
	d._CollectionViewItemForRepresentedObjectAtIndexPath = f
}

// CollectionViewViewForSupplementaryElementOfKindAtIndexPath implements the PCollectionViewDataSource interface.
func (d *CollectionViewDataSource) CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, kind objc.IObject /* cross-framework: CollectionViewSupplementaryElementKind */, indexPath foundation.IndexPath) View {
	if d._CollectionViewViewForSupplementaryElementOfKindAtIndexPath != nil {
		return d._CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView, kind, indexPath)
	}
	var zero View
	return zero
}

// HasCollectionViewViewForSupplementaryElementOfKindAtIndexPath returns true if a handler for CollectionViewViewForSupplementaryElementOfKindAtIndexPath has been set.
func (d *CollectionViewDataSource) HasCollectionViewViewForSupplementaryElementOfKindAtIndexPath() bool {
	return d._CollectionViewViewForSupplementaryElementOfKindAtIndexPath != nil
}

// NumberOfSectionsInCollectionView implements the PCollectionViewDataSource interface.
func (d *CollectionViewDataSource) NumberOfSectionsInCollectionView(collectionView objc.IObject /* cross-framework: CollectionView */) int {
	if d._NumberOfSectionsInCollectionView != nil {
		return d._NumberOfSectionsInCollectionView(collectionView)
	}
	var zero int
	return zero
}

// HasNumberOfSectionsInCollectionView returns true if a handler for NumberOfSectionsInCollectionView has been set.
func (d *CollectionViewDataSource) HasNumberOfSectionsInCollectionView() bool {
	return d._NumberOfSectionsInCollectionView != nil
}

// CollectionViewNumberOfItemsInSection implements the PCollectionViewDataSource interface.
func (d *CollectionViewDataSource) CollectionViewNumberOfItemsInSection(collectionView objc.IObject /* cross-framework: CollectionView */, section int) int {
	if d._CollectionViewNumberOfItemsInSection != nil {
		return d._CollectionViewNumberOfItemsInSection(collectionView, section)
	}
	var zero int
	return zero
}

// HasCollectionViewNumberOfItemsInSection returns true if a handler for CollectionViewNumberOfItemsInSection has been set.
func (d *CollectionViewDataSource) HasCollectionViewNumberOfItemsInSection() bool {
	return d._CollectionViewNumberOfItemsInSection != nil
}

// CollectionViewItemForRepresentedObjectAtIndexPath implements the PCollectionViewDataSource interface.
func (d *CollectionViewDataSource) CollectionViewItemForRepresentedObjectAtIndexPath(collectionView objc.IObject /* cross-framework: CollectionView */, indexPath foundation.IndexPath) CollectionViewItem {
	if d._CollectionViewItemForRepresentedObjectAtIndexPath != nil {
		return d._CollectionViewItemForRepresentedObjectAtIndexPath(collectionView, indexPath)
	}
	var zero CollectionViewItem
	return zero
}

// HasCollectionViewItemForRepresentedObjectAtIndexPath returns true if a handler for CollectionViewItemForRepresentedObjectAtIndexPath has been set.
func (d *CollectionViewDataSource) HasCollectionViewItemForRepresentedObjectAtIndexPath() bool {
	return d._CollectionViewItemForRepresentedObjectAtIndexPath != nil
}
