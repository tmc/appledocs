// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPlayableContentDelegate is the MPPlayableContentDelegate protocol interface.
//
// The protocol used to let external media players send playback commands to an app.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 14.0)
//   - iOS 7.1+ (Deprecated in 14.0)
//   - iPadOS 7.1+ (Deprecated in 14.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// See: doc://com.apple.mediaplayer/documentation/MediaPlayer/MPPlayableContentDelegate
type PPlayableContentDelegate interface {
	// Optional methods
	PlayableContentManagerDidUpdateContext(contentManager IMPPlayableContentManager, context IMPPlayableContentManagerContext)
	HasPlayableContentManagerDidUpdateContext() bool
	PlayableContentManagerInitializePlaybackQueueWithCompletionHandler(contentManager IMPPlayableContentManager, completionHandler unsafe.Pointer)
	HasPlayableContentManagerInitializePlaybackQueueWithCompletionHandler() bool
	PlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler(contentManager IMPPlayableContentManager, contentItems objc.IObject /* cross-framework: NSArray */, completionHandler unsafe.Pointer)
	HasPlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler() bool
	PlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler(contentManager IMPPlayableContentManager, indexPath foundation.IndexPath, completionHandler unsafe.Pointer)
	HasPlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler() bool
}

// PlayableContentDelegate is a delegate implementation builder for the PPlayableContentDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayableContentDelegate struct {
	_PlayableContentManagerDidUpdateContext func(contentManager IMPPlayableContentManager, context IMPPlayableContentManagerContext)
	_PlayableContentManagerInitializePlaybackQueueWithCompletionHandler func(contentManager IMPPlayableContentManager, completionHandler unsafe.Pointer)
	_PlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler func(contentManager IMPPlayableContentManager, contentItems objc.IObject /* cross-framework: NSArray */, completionHandler unsafe.Pointer)
	_PlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler func(contentManager IMPPlayableContentManager, indexPath foundation.IndexPath, completionHandler unsafe.Pointer)
}

// SetPlayableContentManagerDidUpdateContext sets the handler for the PlayableContentManagerDidUpdateContext delegate method.
//
// Notifies the delegate that the playable content manager’s context information has changed.
func (d *PlayableContentDelegate) SetPlayableContentManagerDidUpdateContext(f func(contentManager IMPPlayableContentManager, context IMPPlayableContentManagerContext)) {
	d._PlayableContentManagerDidUpdateContext = f
}

// SetPlayableContentManagerInitializePlaybackQueueWithCompletionHandler sets the handler for the PlayableContentManagerInitializePlaybackQueueWithCompletionHandler delegate method.
//
// Asks the delegate to prepare suggested content for playback.
func (d *PlayableContentDelegate) SetPlayableContentManagerInitializePlaybackQueueWithCompletionHandler(f func(contentManager IMPPlayableContentManager, completionHandler unsafe.Pointer)) {
	d._PlayableContentManagerInitializePlaybackQueueWithCompletionHandler = f
}

// SetPlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler sets the handler for the PlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler delegate method.
//
// Asks the delegate to prepare suggested content for playback.
func (d *PlayableContentDelegate) SetPlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler(f func(contentManager IMPPlayableContentManager, contentItems objc.IObject /* cross-framework: NSArray */, completionHandler unsafe.Pointer)) {
	d._PlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler = f
}

// SetPlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler sets the handler for the PlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler delegate method.
//
// Asks the delegate to begin playback of the specified content item.
func (d *PlayableContentDelegate) SetPlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler(f func(contentManager IMPPlayableContentManager, indexPath foundation.IndexPath, completionHandler unsafe.Pointer)) {
	d._PlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler = f
}

// PlayableContentManagerDidUpdateContext implements the PPlayableContentDelegate interface.
func (d *PlayableContentDelegate) PlayableContentManagerDidUpdateContext(contentManager IMPPlayableContentManager, context IMPPlayableContentManagerContext) {
	if d._PlayableContentManagerDidUpdateContext != nil {
		d._PlayableContentManagerDidUpdateContext(contentManager, context)
	}
}

// HasPlayableContentManagerDidUpdateContext returns true if a handler for PlayableContentManagerDidUpdateContext has been set.
func (d *PlayableContentDelegate) HasPlayableContentManagerDidUpdateContext() bool {
	return d._PlayableContentManagerDidUpdateContext != nil
}

// PlayableContentManagerInitializePlaybackQueueWithCompletionHandler implements the PPlayableContentDelegate interface.
func (d *PlayableContentDelegate) PlayableContentManagerInitializePlaybackQueueWithCompletionHandler(contentManager IMPPlayableContentManager, completionHandler unsafe.Pointer) {
	if d._PlayableContentManagerInitializePlaybackQueueWithCompletionHandler != nil {
		d._PlayableContentManagerInitializePlaybackQueueWithCompletionHandler(contentManager, completionHandler)
	}
}

// HasPlayableContentManagerInitializePlaybackQueueWithCompletionHandler returns true if a handler for PlayableContentManagerInitializePlaybackQueueWithCompletionHandler has been set.
func (d *PlayableContentDelegate) HasPlayableContentManagerInitializePlaybackQueueWithCompletionHandler() bool {
	return d._PlayableContentManagerInitializePlaybackQueueWithCompletionHandler != nil
}

// PlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler implements the PPlayableContentDelegate interface.
func (d *PlayableContentDelegate) PlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler(contentManager IMPPlayableContentManager, contentItems objc.IObject /* cross-framework: NSArray */, completionHandler unsafe.Pointer) {
	if d._PlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler != nil {
		d._PlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler(contentManager, contentItems, completionHandler)
	}
}

// HasPlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler returns true if a handler for PlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler has been set.
func (d *PlayableContentDelegate) HasPlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler() bool {
	return d._PlayableContentManagerInitializePlaybackQueueWithContentItemsCompletionHandler != nil
}

// PlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler implements the PPlayableContentDelegate interface.
func (d *PlayableContentDelegate) PlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler(contentManager IMPPlayableContentManager, indexPath foundation.IndexPath, completionHandler unsafe.Pointer) {
	if d._PlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler != nil {
		d._PlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler(contentManager, indexPath, completionHandler)
	}
}

// HasPlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler returns true if a handler for PlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler has been set.
func (d *PlayableContentDelegate) HasPlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler() bool {
	return d._PlayableContentManagerInitiatePlaybackOfContentItemAtIndexPathCompletionHandler != nil
}
