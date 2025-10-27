// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PPlayerItemMetadataCollectorPushDelegate is the AVPlayerItemMetadataCollectorPushDelegate protocol interface.
//
// A protocol you implement to receive metadata callbacks from a player item metadata collector.
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
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVPlayerItemMetadataCollectorPushDelegate
type PPlayerItemMetadataCollectorPushDelegate interface {
	// Required methods
	MetadataCollectorDidCollectDateRangeMetadataGroupsIndexesOfNewGroupsIndexesOfModifiedGroups(metadataCollector IAVPlayerItemMetadataCollector, metadataGroups []DateRangeMetadataGroup, indexesOfNewGroups foundation.IndexSet, indexesOfModifiedGroups foundation.IndexSet)
}
