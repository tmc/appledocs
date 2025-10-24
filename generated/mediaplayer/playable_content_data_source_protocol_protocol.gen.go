// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPlayableContentDataSource is the MPPlayableContentDataSource protocol interface.
//
// The data source providing media metadata to external media players so they can build user interfaces displaying your app’s content.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 14.0)
//   - iOS 7.1+ (Deprecated in 14.0)
//   - iPadOS 7.1+ (Deprecated in 14.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// See: doc://com.apple.mediaplayer/documentation/MediaPlayer/MPPlayableContentDataSource
type PPlayableContentDataSource interface {
	// Required methods
	ContentItemAtIndexPath(indexPath foundation.IndexPath) ContentItem/* debug [protocol_interface/required_method]: ContentItemAtIndexPath */
	NumberOfChildItemsAtIndexPath(indexPath foundation.IndexPath) int/* debug [protocol_interface/required_method]: NumberOfChildItemsAtIndexPath */
	// Optional methods
	BeginLoadingChildItemsAtIndexPathCompletionHandler(indexPath foundation.IndexPath, completionHandler unsafe.Pointer)
	HasBeginLoadingChildItemsAtIndexPathCompletionHandler() bool
	ChildItemsDisplayPlaybackProgressAtIndexPath(indexPath foundation.IndexPath) bool
	HasChildItemsDisplayPlaybackProgressAtIndexPath() bool
	ContentItemForIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	HasContentItemForIdentifierCompletionHandler() bool
}

// PlayableContentDataSource is a delegate implementation builder for the PPlayableContentDataSource protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayableContentDataSource struct {
	_BeginLoadingChildItemsAtIndexPathCompletionHandler func(indexPath foundation.IndexPath, completionHandler unsafe.Pointer)
	_ChildItemsDisplayPlaybackProgressAtIndexPath func(indexPath foundation.IndexPath) bool
	_ContentItemForIdentifierCompletionHandler func(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	_ContentItemAtIndexPath func(indexPath foundation.IndexPath) ContentItem
	_NumberOfChildItemsAtIndexPath func(indexPath foundation.IndexPath) int
}

// SetBeginLoadingChildItemsAtIndexPathCompletionHandler sets the handler for the BeginLoadingChildItemsAtIndexPathCompletionHandler delegate method.
//
// Starts to load the children of the indicated index.
func (d *PlayableContentDataSource) SetBeginLoadingChildItemsAtIndexPathCompletionHandler(f func(indexPath foundation.IndexPath, completionHandler unsafe.Pointer)) {
	d._BeginLoadingChildItemsAtIndexPathCompletionHandler = f
}

// SetChildItemsDisplayPlaybackProgressAtIndexPath sets the handler for the ChildItemsDisplayPlaybackProgressAtIndexPath delegate method.
//
// Returns a Boolean value indicating whether the provided content supports playback progress.
func (d *PlayableContentDataSource) SetChildItemsDisplayPlaybackProgressAtIndexPath(f func(indexPath foundation.IndexPath) bool) {
	d._ChildItemsDisplayPlaybackProgressAtIndexPath = f
}

// SetContentItemForIdentifierCompletionHandler sets the handler for the ContentItemForIdentifierCompletionHandler delegate method.
//
// Retrieves the content item associated with the provided identifier.
func (d *PlayableContentDataSource) SetContentItemForIdentifierCompletionHandler(f func(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)) {
	d._ContentItemForIdentifierCompletionHandler = f
}

// SetContentItemAtIndexPath sets the handler for the ContentItemAtIndexPath delegate method.
//
// Retrieves the media item at the specified index.
func (d *PlayableContentDataSource) SetContentItemAtIndexPath(f func(indexPath foundation.IndexPath) ContentItem) {
	d._ContentItemAtIndexPath = f
}

// SetNumberOfChildItemsAtIndexPath sets the handler for the NumberOfChildItemsAtIndexPath delegate method.
//
// Provides the number of child nodes for the indicated node.
func (d *PlayableContentDataSource) SetNumberOfChildItemsAtIndexPath(f func(indexPath foundation.IndexPath) int) {
	d._NumberOfChildItemsAtIndexPath = f
}

// BeginLoadingChildItemsAtIndexPathCompletionHandler implements the PPlayableContentDataSource interface.
func (d *PlayableContentDataSource) BeginLoadingChildItemsAtIndexPathCompletionHandler(indexPath foundation.IndexPath, completionHandler unsafe.Pointer) {
	if d._BeginLoadingChildItemsAtIndexPathCompletionHandler != nil {
		d._BeginLoadingChildItemsAtIndexPathCompletionHandler(indexPath, completionHandler)
	}
}

// HasBeginLoadingChildItemsAtIndexPathCompletionHandler returns true if a handler for BeginLoadingChildItemsAtIndexPathCompletionHandler has been set.
func (d *PlayableContentDataSource) HasBeginLoadingChildItemsAtIndexPathCompletionHandler() bool {
	return d._BeginLoadingChildItemsAtIndexPathCompletionHandler != nil
}

// ChildItemsDisplayPlaybackProgressAtIndexPath implements the PPlayableContentDataSource interface.
func (d *PlayableContentDataSource) ChildItemsDisplayPlaybackProgressAtIndexPath(indexPath foundation.IndexPath) bool {
	if d._ChildItemsDisplayPlaybackProgressAtIndexPath != nil {
		return d._ChildItemsDisplayPlaybackProgressAtIndexPath(indexPath)
	}
	var zero bool
	return zero
}

// HasChildItemsDisplayPlaybackProgressAtIndexPath returns true if a handler for ChildItemsDisplayPlaybackProgressAtIndexPath has been set.
func (d *PlayableContentDataSource) HasChildItemsDisplayPlaybackProgressAtIndexPath() bool {
	return d._ChildItemsDisplayPlaybackProgressAtIndexPath != nil
}

// ContentItemForIdentifierCompletionHandler implements the PPlayableContentDataSource interface.
func (d *PlayableContentDataSource) ContentItemForIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	if d._ContentItemForIdentifierCompletionHandler != nil {
		d._ContentItemForIdentifierCompletionHandler(identifier, completionHandler)
	}
}

// HasContentItemForIdentifierCompletionHandler returns true if a handler for ContentItemForIdentifierCompletionHandler has been set.
func (d *PlayableContentDataSource) HasContentItemForIdentifierCompletionHandler() bool {
	return d._ContentItemForIdentifierCompletionHandler != nil
}

// ContentItemAtIndexPath implements the PPlayableContentDataSource interface.
func (d *PlayableContentDataSource) ContentItemAtIndexPath(indexPath foundation.IndexPath) ContentItem {
	if d._ContentItemAtIndexPath != nil {
		return d._ContentItemAtIndexPath(indexPath)
	}
	var zero ContentItem
	return zero
}

// HasContentItemAtIndexPath returns true if a handler for ContentItemAtIndexPath has been set.
func (d *PlayableContentDataSource) HasContentItemAtIndexPath() bool {
	return d._ContentItemAtIndexPath != nil
}

// NumberOfChildItemsAtIndexPath implements the PPlayableContentDataSource interface.
func (d *PlayableContentDataSource) NumberOfChildItemsAtIndexPath(indexPath foundation.IndexPath) int {
	if d._NumberOfChildItemsAtIndexPath != nil {
		return d._NumberOfChildItemsAtIndexPath(indexPath)
	}
	var zero int
	return zero
}

// HasNumberOfChildItemsAtIndexPath returns true if a handler for NumberOfChildItemsAtIndexPath has been set.
func (d *PlayableContentDataSource) HasNumberOfChildItemsAtIndexPath() bool {
	return d._NumberOfChildItemsAtIndexPath != nil
}
