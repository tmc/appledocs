// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
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
	CollectionViewNumberOfItemsInSection(collectionView CollectionView /* not a class type */, section int) int
	CollectionViewItemForRepresentedObjectAtIndexPath(collectionView CollectionView /* not a class type */, indexPath foundation.foundation.INSIndexPath) ICollectionViewItem
	// Optional methods
	CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView CollectionView /* not a class type */, kind CollectionViewSupplementaryElementKind, indexPath foundation.foundation.INSIndexPath) IView
	HasCollectionViewViewForSupplementaryElementOfKindAtIndexPath() bool
	NumberOfSectionsInCollectionView(collectionView CollectionView /* not a class type */) int
	HasNumberOfSectionsInCollectionView() bool
}

// CollectionViewDataSource is a delegate implementation builder for the PCollectionViewDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CollectionViewDataSource struct {
	_CollectionViewViewForSupplementaryElementOfKindAtIndexPath func(collectionView CollectionView /* not a class type */, kind CollectionViewSupplementaryElementKind, indexPath foundation.foundation.INSIndexPath) IView
	_NumberOfSectionsInCollectionView func(collectionView CollectionView /* not a class type */) int
	_CollectionViewNumberOfItemsInSection func(collectionView CollectionView /* not a class type */, section int) int
	_CollectionViewItemForRepresentedObjectAtIndexPath func(collectionView CollectionView /* not a class type */, indexPath foundation.foundation.INSIndexPath) ICollectionViewItem
}

// SetCollectionViewViewForSupplementaryElementOfKindAtIndexPath sets the handler for the CollectionViewViewForSupplementaryElementOfKindAtIndexPath delegate method.
//
// Asks your data source object to provide the supplementary view at the specified location in a section of the collection view.
func (d *CollectionViewDataSource) SetCollectionViewViewForSupplementaryElementOfKindAtIndexPath(f func(collectionView CollectionView /* not a class type */, kind CollectionViewSupplementaryElementKind, indexPath foundation.foundation.INSIndexPath) IView) {
	d._CollectionViewViewForSupplementaryElementOfKindAtIndexPath = f
}

// SetNumberOfSectionsInCollectionView sets the handler for the NumberOfSectionsInCollectionView delegate method.
//
// Asks your data source object to provide the total number of sections.
func (d *CollectionViewDataSource) SetNumberOfSectionsInCollectionView(f func(collectionView CollectionView /* not a class type */) int) {
	d._NumberOfSectionsInCollectionView = f
}

// SetCollectionViewNumberOfItemsInSection sets the handler for the CollectionViewNumberOfItemsInSection delegate method.
//
// Asks your data source object to provide the number of items in the specified section.
func (d *CollectionViewDataSource) SetCollectionViewNumberOfItemsInSection(f func(collectionView CollectionView /* not a class type */, section int) int) {
	d._CollectionViewNumberOfItemsInSection = f
}

// SetCollectionViewItemForRepresentedObjectAtIndexPath sets the handler for the CollectionViewItemForRepresentedObjectAtIndexPath delegate method.
//
// Asks your data source object to provide the item at the specified location in the collection view.
func (d *CollectionViewDataSource) SetCollectionViewItemForRepresentedObjectAtIndexPath(f func(collectionView CollectionView /* not a class type */, indexPath foundation.foundation.INSIndexPath) ICollectionViewItem) {
	d._CollectionViewItemForRepresentedObjectAtIndexPath = f
}

// CollectionViewViewForSupplementaryElementOfKindAtIndexPath implements the PCollectionViewDataSource interface.
func (d *CollectionViewDataSource) CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView CollectionView /* not a class type */, kind CollectionViewSupplementaryElementKind, indexPath foundation.foundation.INSIndexPath) IView {
	if d._CollectionViewViewForSupplementaryElementOfKindAtIndexPath != nil {
		return d._CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView, kind, indexPath)
	}
	var zero IView
	return zero
}

// HasCollectionViewViewForSupplementaryElementOfKindAtIndexPath returns true if a handler for CollectionViewViewForSupplementaryElementOfKindAtIndexPath has been set.
func (d *CollectionViewDataSource) HasCollectionViewViewForSupplementaryElementOfKindAtIndexPath() bool {
	return d._CollectionViewViewForSupplementaryElementOfKindAtIndexPath != nil
}

// NumberOfSectionsInCollectionView implements the PCollectionViewDataSource interface.
func (d *CollectionViewDataSource) NumberOfSectionsInCollectionView(collectionView CollectionView /* not a class type */) int {
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
func (d *CollectionViewDataSource) CollectionViewNumberOfItemsInSection(collectionView CollectionView /* not a class type */, section int) int {
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
func (d *CollectionViewDataSource) CollectionViewItemForRepresentedObjectAtIndexPath(collectionView CollectionView /* not a class type */, indexPath foundation.foundation.INSIndexPath) ICollectionViewItem {
	if d._CollectionViewItemForRepresentedObjectAtIndexPath != nil {
		return d._CollectionViewItemForRepresentedObjectAtIndexPath(collectionView, indexPath)
	}
	var zero ICollectionViewItem
	return zero
}

// HasCollectionViewItemForRepresentedObjectAtIndexPath returns true if a handler for CollectionViewItemForRepresentedObjectAtIndexPath has been set.
func (d *CollectionViewDataSource) HasCollectionViewItemForRepresentedObjectAtIndexPath() bool {
	return d._CollectionViewItemForRepresentedObjectAtIndexPath != nil
}

// CollectionViewDataSourceObject wraps an existing Objective-C object that conforms to the PCollectionViewDataSource protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type CollectionViewDataSourceObject struct {
	objectivec.Object
}

// NewCollectionViewDataSourceObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSCollectionViewDataSource protocol.
func NewCollectionViewDataSourceObject(obj objectivec.Object) *CollectionViewDataSourceObject {
	return &CollectionViewDataSourceObject{obj}
}

// Make sure CollectionViewDataSourceObject implements PCollectionViewDataSource.
var _ PCollectionViewDataSource = (*CollectionViewDataSourceObject)(nil)

// CollectionViewNumberOfItemsInSection implements the PCollectionViewDataSource interface.
// This required method is always available on objects conforming to CollectionViewNumberOfItemsInSection.
func (o *CollectionViewDataSourceObject) CollectionViewNumberOfItemsInSection(collectionView CollectionView /* not a class type */, section int) int {
	return objc.Send[int](o.ID, objc.Sel("collectionView:numberOfItemsInSection:"), collectionView, section)
}

// CollectionViewItemForRepresentedObjectAtIndexPath implements the PCollectionViewDataSource interface.
// This required method is always available on objects conforming to CollectionViewItemForRepresentedObjectAtIndexPath.
func (o *CollectionViewDataSourceObject) CollectionViewItemForRepresentedObjectAtIndexPath(collectionView CollectionView /* not a class type */, indexPath foundation.foundation.INSIndexPath) ICollectionViewItem {
	return objc.Send[ICollectionViewItem](o.ID, objc.Sel("collectionView:itemForRepresentedObjectAtIndexPath:"), collectionView, indexPath)
}

// CollectionViewViewForSupplementaryElementOfKindAtIndexPath implements the PCollectionViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *CollectionViewDataSourceObject) CollectionViewViewForSupplementaryElementOfKindAtIndexPath(collectionView CollectionView /* not a class type */, kind CollectionViewSupplementaryElementKind, indexPath foundation.foundation.INSIndexPath) IView {
	return objc.Send[IView](o.ID, objc.Sel("collectionView:viewForSupplementaryElementOfKind:atIndexPath:"), collectionView, kind, indexPath)
}

// HasCollectionViewViewForSupplementaryElementOfKindAtIndexPath returns true; this is a placeholder for optional method checks.
func (o *CollectionViewDataSourceObject) HasCollectionViewViewForSupplementaryElementOfKindAtIndexPath() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// NumberOfSectionsInCollectionView implements the PCollectionViewDataSource interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *CollectionViewDataSourceObject) NumberOfSectionsInCollectionView(collectionView CollectionView /* not a class type */) int {
	return objc.Send[int](o.ID, objc.Sel("numberOfSectionsInCollectionView:"), collectionView)
}

// HasNumberOfSectionsInCollectionView returns true; this is a placeholder for optional method checks.
func (o *CollectionViewDataSourceObject) HasNumberOfSectionsInCollectionView() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
