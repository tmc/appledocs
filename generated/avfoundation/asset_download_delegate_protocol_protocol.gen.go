// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PAssetDownloadDelegate is the AVAssetDownloadDelegate protocol interface.
//
// A protocol that defines the methods to implement to respond to asset-download events.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//   - watchOS 10.0+
//
// See: doc://com.apple.avfoundation/documentation/AVFoundation/AVAssetDownloadDelegate
type PAssetDownloadDelegate interface {
	// Optional methods
	URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection(session URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, mediaSelection IAVMediaSelection)
	HasURLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection() bool
	URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection(session URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, timeRange objectivec.IObject, loadedTimeRanges []foundation.Value, timeRangeExpectedToLoad objectivec.IObject, mediaSelection IAVMediaSelection)
	HasURLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection() bool
	URLSessionAggregateAssetDownloadTaskWillDownloadToURL(session URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, location foundation.foundation.INSURL)
	HasURLSessionAggregateAssetDownloadTaskWillDownloadToURL() bool
	URLSessionAssetDownloadTaskDidFinishDownloadingToURL(session URLSession, assetDownloadTask IAVAssetDownloadTask, location foundation.foundation.INSURL)
	HasURLSessionAssetDownloadTaskDidFinishDownloadingToURL() bool
	URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad(session URLSession, assetDownloadTask IAVAssetDownloadTask, timeRange objectivec.IObject, loadedTimeRanges []foundation.Value, timeRangeExpectedToLoad objectivec.IObject)
	HasURLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad() bool
	URLSessionAssetDownloadTaskDidReceiveMetricEvent(session URLSession, assetDownloadTask IAVAssetDownloadTask, metricEvent IAVMetricEvent)
	HasURLSessionAssetDownloadTaskDidReceiveMetricEvent() bool
	URLSessionAssetDownloadTaskDidResolveMediaSelection(session URLSession, assetDownloadTask IAVAssetDownloadTask, resolvedMediaSelection IAVMediaSelection)
	HasURLSessionAssetDownloadTaskDidResolveMediaSelection() bool
	URLSessionAssetDownloadTaskWillDownloadToURL(session URLSession, assetDownloadTask IAVAssetDownloadTask, location foundation.foundation.INSURL)
	HasURLSessionAssetDownloadTaskWillDownloadToURL() bool
	URLSessionAssetDownloadTaskWillDownloadVariants(session URLSession, assetDownloadTask IAVAssetDownloadTask, variants []AssetVariant)
	HasURLSessionAssetDownloadTaskWillDownloadVariants() bool
}
