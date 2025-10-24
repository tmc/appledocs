// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCSSearchableIndexDelegate is the CSSearchableIndexDelegate protocol interface.
//
// A protocol that defines methods a delegate object or app extension uses to handle communication from the on-device index.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - visionOS 1.0+
//
// See: doc://com.apple.corespotlight/documentation/CoreSpotlight/CSSearchableIndexDelegate
type PCSSearchableIndexDelegate interface {
	// Required methods
	SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex ICSSearchableIndex, acknowledgementHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler */
	SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex ICSSearchableIndex, identifiers []string, acknowledgementHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler */
	// Optional methods
	DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex ICSSearchableIndex, itemIdentifier objc.IObject /* cross-framework: NSString */, typeIdentifier objc.IObject /* cross-framework: NSString */, outError objectivec.IObject) foundation.Data
	HasDataForSearchableIndexItemIdentifierTypeIdentifierError() bool
	FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex ICSSearchableIndex, itemIdentifier objc.IObject /* cross-framework: NSString */, typeIdentifier objc.IObject /* cross-framework: NSString */, inPlace bool, outError objectivec.IObject) foundation.URL
	HasFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError() bool
	SearchableIndexDidFinishThrottle(searchableIndex ICSSearchableIndex)
	HasSearchableIndexDidFinishThrottle() bool
	SearchableIndexDidThrottle(searchableIndex ICSSearchableIndex)
	HasSearchableIndexDidThrottle() bool
	SearchableItemsForIdentifiersSearchableItemsHandler(identifiers []string, searchableItemsHandler unsafe.Pointer)
	HasSearchableItemsForIdentifiersSearchableItemsHandler() bool
	SearchableItemsDidUpdate(items []CSSearchableItem)
	HasSearchableItemsDidUpdate() bool
}

// CSSearchableIndexDelegate is a delegate implementation builder for the PCSSearchableIndexDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CSSearchableIndexDelegate struct {
	_DataForSearchableIndexItemIdentifierTypeIdentifierError func(searchableIndex ICSSearchableIndex, itemIdentifier objc.IObject /* cross-framework: NSString */, typeIdentifier objc.IObject /* cross-framework: NSString */, outError objectivec.IObject) foundation.Data
	_FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError func(searchableIndex ICSSearchableIndex, itemIdentifier objc.IObject /* cross-framework: NSString */, typeIdentifier objc.IObject /* cross-framework: NSString */, inPlace bool, outError objectivec.IObject) foundation.URL
	_SearchableIndexDidFinishThrottle func(searchableIndex ICSSearchableIndex)
	_SearchableIndexDidThrottle func(searchableIndex ICSSearchableIndex)
	_SearchableItemsForIdentifiersSearchableItemsHandler func(identifiers []string, searchableItemsHandler unsafe.Pointer)
	_SearchableItemsDidUpdate func(items []CSSearchableItem)
	_SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler func(searchableIndex ICSSearchableIndex, acknowledgementHandler unsafe.Pointer)
	_SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler func(searchableIndex ICSSearchableIndex, identifiers []string, acknowledgementHandler unsafe.Pointer)
}

// SetDataForSearchableIndexItemIdentifierTypeIdentifierError sets the handler for the DataForSearchableIndexItemIdentifierTypeIdentifierError delegate method.
func (d *CSSearchableIndexDelegate) SetDataForSearchableIndexItemIdentifierTypeIdentifierError(f func(searchableIndex ICSSearchableIndex, itemIdentifier objc.IObject /* cross-framework: NSString */, typeIdentifier objc.IObject /* cross-framework: NSString */, outError objectivec.IObject) foundation.Data) {
	d._DataForSearchableIndexItemIdentifierTypeIdentifierError = f
}

// SetFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError sets the handler for the FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError delegate method.
func (d *CSSearchableIndexDelegate) SetFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(f func(searchableIndex ICSSearchableIndex, itemIdentifier objc.IObject /* cross-framework: NSString */, typeIdentifier objc.IObject /* cross-framework: NSString */, inPlace bool, outError objectivec.IObject) foundation.URL) {
	d._FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError = f
}

// SetSearchableIndexDidFinishThrottle sets the handler for the SearchableIndexDidFinishThrottle delegate method.
//
// Tells the delegate that the index throttling has finished.
func (d *CSSearchableIndexDelegate) SetSearchableIndexDidFinishThrottle(f func(searchableIndex ICSSearchableIndex)) {
	d._SearchableIndexDidFinishThrottle = f
}

// SetSearchableIndexDidThrottle sets the handler for the SearchableIndexDidThrottle delegate method.
//
// Tells the delegate that indexing is being throttled.
func (d *CSSearchableIndexDelegate) SetSearchableIndexDidThrottle(f func(searchableIndex ICSSearchableIndex)) {
	d._SearchableIndexDidThrottle = f
}

// SetSearchableItemsForIdentifiersSearchableItemsHandler sets the handler for the SearchableItemsForIdentifiersSearchableItemsHandler delegate method.
//
// Requests that the delegate provide searchable items for the provided identifiers.
func (d *CSSearchableIndexDelegate) SetSearchableItemsForIdentifiersSearchableItemsHandler(f func(identifiers []string, searchableItemsHandler unsafe.Pointer)) {
	d._SearchableItemsForIdentifiersSearchableItemsHandler = f
}

// SetSearchableItemsDidUpdate sets the handler for the SearchableItemsDidUpdate delegate method.
//
// Tells the delegate that the framework updated the list of searchable items.
func (d *CSSearchableIndexDelegate) SetSearchableItemsDidUpdate(f func(items []CSSearchableItem)) {
	d._SearchableItemsDidUpdate = f
}

// SetSearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler sets the handler for the SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler delegate method.
//
// Tells the delegate to reindex all searchable data and clear all local state information.
func (d *CSSearchableIndexDelegate) SetSearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(f func(searchableIndex ICSSearchableIndex, acknowledgementHandler unsafe.Pointer)) {
	d._SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler = f
}

// SetSearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler sets the handler for the SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler delegate method.
//
// Tells the delegate to reindex the searchable items associated with the specified identifiers.
func (d *CSSearchableIndexDelegate) SetSearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(f func(searchableIndex ICSSearchableIndex, identifiers []string, acknowledgementHandler unsafe.Pointer)) {
	d._SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler = f
}

// DataForSearchableIndexItemIdentifierTypeIdentifierError implements the PCSSearchableIndexDelegate interface.
func (d *CSSearchableIndexDelegate) DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex ICSSearchableIndex, itemIdentifier objc.IObject /* cross-framework: NSString */, typeIdentifier objc.IObject /* cross-framework: NSString */, outError objectivec.IObject) foundation.Data {
	if d._DataForSearchableIndexItemIdentifierTypeIdentifierError != nil {
		return d._DataForSearchableIndexItemIdentifierTypeIdentifierError(searchableIndex, itemIdentifier, typeIdentifier, outError)
	}
	var zero foundation.Data
	return zero
}

// HasDataForSearchableIndexItemIdentifierTypeIdentifierError returns true if a handler for DataForSearchableIndexItemIdentifierTypeIdentifierError has been set.
func (d *CSSearchableIndexDelegate) HasDataForSearchableIndexItemIdentifierTypeIdentifierError() bool {
	return d._DataForSearchableIndexItemIdentifierTypeIdentifierError != nil
}

// FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError implements the PCSSearchableIndexDelegate interface.
func (d *CSSearchableIndexDelegate) FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex ICSSearchableIndex, itemIdentifier objc.IObject /* cross-framework: NSString */, typeIdentifier objc.IObject /* cross-framework: NSString */, inPlace bool, outError objectivec.IObject) foundation.URL {
	if d._FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError != nil {
		return d._FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError(searchableIndex, itemIdentifier, typeIdentifier, inPlace, outError)
	}
	var zero foundation.URL
	return zero
}

// HasFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError returns true if a handler for FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError has been set.
func (d *CSSearchableIndexDelegate) HasFileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError() bool {
	return d._FileURLForSearchableIndexItemIdentifierTypeIdentifierInPlaceError != nil
}

// SearchableIndexDidFinishThrottle implements the PCSSearchableIndexDelegate interface.
func (d *CSSearchableIndexDelegate) SearchableIndexDidFinishThrottle(searchableIndex ICSSearchableIndex) {
	if d._SearchableIndexDidFinishThrottle != nil {
		d._SearchableIndexDidFinishThrottle(searchableIndex)
	}
}

// HasSearchableIndexDidFinishThrottle returns true if a handler for SearchableIndexDidFinishThrottle has been set.
func (d *CSSearchableIndexDelegate) HasSearchableIndexDidFinishThrottle() bool {
	return d._SearchableIndexDidFinishThrottle != nil
}

// SearchableIndexDidThrottle implements the PCSSearchableIndexDelegate interface.
func (d *CSSearchableIndexDelegate) SearchableIndexDidThrottle(searchableIndex ICSSearchableIndex) {
	if d._SearchableIndexDidThrottle != nil {
		d._SearchableIndexDidThrottle(searchableIndex)
	}
}

// HasSearchableIndexDidThrottle returns true if a handler for SearchableIndexDidThrottle has been set.
func (d *CSSearchableIndexDelegate) HasSearchableIndexDidThrottle() bool {
	return d._SearchableIndexDidThrottle != nil
}

// SearchableItemsForIdentifiersSearchableItemsHandler implements the PCSSearchableIndexDelegate interface.
func (d *CSSearchableIndexDelegate) SearchableItemsForIdentifiersSearchableItemsHandler(identifiers []string, searchableItemsHandler unsafe.Pointer) {
	if d._SearchableItemsForIdentifiersSearchableItemsHandler != nil {
		d._SearchableItemsForIdentifiersSearchableItemsHandler(identifiers, searchableItemsHandler)
	}
}

// HasSearchableItemsForIdentifiersSearchableItemsHandler returns true if a handler for SearchableItemsForIdentifiersSearchableItemsHandler has been set.
func (d *CSSearchableIndexDelegate) HasSearchableItemsForIdentifiersSearchableItemsHandler() bool {
	return d._SearchableItemsForIdentifiersSearchableItemsHandler != nil
}

// SearchableItemsDidUpdate implements the PCSSearchableIndexDelegate interface.
func (d *CSSearchableIndexDelegate) SearchableItemsDidUpdate(items []CSSearchableItem) {
	if d._SearchableItemsDidUpdate != nil {
		d._SearchableItemsDidUpdate(items)
	}
}

// HasSearchableItemsDidUpdate returns true if a handler for SearchableItemsDidUpdate has been set.
func (d *CSSearchableIndexDelegate) HasSearchableItemsDidUpdate() bool {
	return d._SearchableItemsDidUpdate != nil
}

// SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler implements the PCSSearchableIndexDelegate interface.
func (d *CSSearchableIndexDelegate) SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex ICSSearchableIndex, acknowledgementHandler unsafe.Pointer) {
	if d._SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler != nil {
		d._SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler(searchableIndex, acknowledgementHandler)
	}
}

// HasSearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler returns true if a handler for SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler has been set.
func (d *CSSearchableIndexDelegate) HasSearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler() bool {
	return d._SearchableIndexReindexAllSearchableItemsWithAcknowledgementHandler != nil
}

// SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler implements the PCSSearchableIndexDelegate interface.
func (d *CSSearchableIndexDelegate) SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex ICSSearchableIndex, identifiers []string, acknowledgementHandler unsafe.Pointer) {
	if d._SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler != nil {
		d._SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler(searchableIndex, identifiers, acknowledgementHandler)
	}
}

// HasSearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler returns true if a handler for SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler has been set.
func (d *CSSearchableIndexDelegate) HasSearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler() bool {
	return d._SearchableIndexReindexSearchableItemsWithIdentifiersAcknowledgementHandler != nil
}
