// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PPlayerItemMetadataOutputPushDelegate is the AVPlayerItemMetadataOutputPushDelegate protocol interface.
//
// Methods you can implement to provide additional metadata.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlayerItemMetadataOutputPushDelegate
type PPlayerItemMetadataOutputPushDelegate interface {
	// Optional methods
	MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(output IAVPlayerItemMetadataOutput, groups []TimedMetadataGroup, track IAVPlayerItemTrack)
	HasMetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack() bool
}

// PlayerItemMetadataOutputPushDelegate is a delegate implementation builder for the PPlayerItemMetadataOutputPushDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type PlayerItemMetadataOutputPushDelegate struct {
	_MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack func(output IAVPlayerItemMetadataOutput, groups []TimedMetadataGroup, track IAVPlayerItemTrack)
}

// SetMetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack sets the handler for the MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack delegate method.
//
// Tells the delegate a new collection of metadata items is available.
func (d *PlayerItemMetadataOutputPushDelegate) SetMetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(f func(output IAVPlayerItemMetadataOutput, groups []TimedMetadataGroup, track IAVPlayerItemTrack)) {
	d._MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack = f
}

// MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack implements the PPlayerItemMetadataOutputPushDelegate interface.
func (d *PlayerItemMetadataOutputPushDelegate) MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(output IAVPlayerItemMetadataOutput, groups []TimedMetadataGroup, track IAVPlayerItemTrack) {
	if d._MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack != nil {
		d._MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack(output, groups, track)
	}
}

// HasMetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack returns true if a handler for MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack has been set.
func (d *PlayerItemMetadataOutputPushDelegate) HasMetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack() bool {
	return d._MetadataOutputDidOutputTimedMetadataGroupsFromPlayerItemTrack != nil
}
