// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
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
	URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, mediaSelection IAVMediaSelection)
	HasURLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection() bool
	URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, timeRange TimeRange /* not a class type */, loadedTimeRanges []foundation.Value, timeRangeExpectedToLoad TimeRange /* not a class type */, mediaSelection IAVMediaSelection)
	HasURLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection() bool
	URLSessionAggregateAssetDownloadTaskWillDownloadToURL(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */)
	HasURLSessionAggregateAssetDownloadTaskWillDownloadToURL() bool
	URLSessionAssetDownloadTaskDidFinishDownloadingToURL(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */)
	HasURLSessionAssetDownloadTaskDidFinishDownloadingToURL() bool
	URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, timeRange TimeRange /* not a class type */, loadedTimeRanges []foundation.Value, timeRangeExpectedToLoad TimeRange /* not a class type */)
	HasURLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad() bool
	URLSessionAssetDownloadTaskDidReceiveMetricEvent(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, metricEvent IAVMetricEvent)
	HasURLSessionAssetDownloadTaskDidReceiveMetricEvent() bool
	URLSessionAssetDownloadTaskDidResolveMediaSelection(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, resolvedMediaSelection IAVMediaSelection)
	HasURLSessionAssetDownloadTaskDidResolveMediaSelection() bool
	URLSessionAssetDownloadTaskWillDownloadToURL(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */)
	HasURLSessionAssetDownloadTaskWillDownloadToURL() bool
	URLSessionAssetDownloadTaskWillDownloadVariants(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, variants []AssetVariant)
	HasURLSessionAssetDownloadTaskWillDownloadVariants() bool
}

// AssetDownloadDelegate is a delegate implementation builder for the PAssetDownloadDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AssetDownloadDelegate struct {
	_URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection func(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, mediaSelection IAVMediaSelection)
	_URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection func(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, timeRange TimeRange /* not a class type */, loadedTimeRanges []foundation.Value, timeRangeExpectedToLoad TimeRange /* not a class type */, mediaSelection IAVMediaSelection)
	_URLSessionAggregateAssetDownloadTaskWillDownloadToURL func(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */)
	_URLSessionAssetDownloadTaskDidFinishDownloadingToURL func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */)
	_URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, timeRange TimeRange /* not a class type */, loadedTimeRanges []foundation.Value, timeRangeExpectedToLoad TimeRange /* not a class type */)
	_URLSessionAssetDownloadTaskDidReceiveMetricEvent func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, metricEvent IAVMetricEvent)
	_URLSessionAssetDownloadTaskDidResolveMediaSelection func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, resolvedMediaSelection IAVMediaSelection)
	_URLSessionAssetDownloadTaskWillDownloadToURL func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */)
	_URLSessionAssetDownloadTaskWillDownloadVariants func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, variants []AssetVariant)
}

// SetURLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection sets the handler for the URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection delegate method.
//
// Tells the delegate that a child task finished downloading a media selection.
func (d *AssetDownloadDelegate) SetURLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection(f func(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, mediaSelection IAVMediaSelection)) {
	d._URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection = f
}

// SetURLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection sets the handler for the URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection delegate method.
//
// Tells the delegate that the aggregate download task loaded a new time range.
func (d *AssetDownloadDelegate) SetURLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection(f func(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, timeRange TimeRange /* not a class type */, loadedTimeRanges []foundation.Value, timeRangeExpectedToLoad TimeRange /* not a class type */, mediaSelection IAVMediaSelection)) {
	d._URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection = f
}

// SetURLSessionAggregateAssetDownloadTaskWillDownloadToURL sets the handler for the URLSessionAggregateAssetDownloadTaskWillDownloadToURL delegate method.
//
// Tells the delegate the final location of the asset when the download completes.
func (d *AssetDownloadDelegate) SetURLSessionAggregateAssetDownloadTaskWillDownloadToURL(f func(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */)) {
	d._URLSessionAggregateAssetDownloadTaskWillDownloadToURL = f
}

// SetURLSessionAssetDownloadTaskDidFinishDownloadingToURL sets the handler for the URLSessionAssetDownloadTaskDidFinishDownloadingToURL delegate method.
//
// Tells the delegate that a download task finished downloading the requested asset.
func (d *AssetDownloadDelegate) SetURLSessionAssetDownloadTaskDidFinishDownloadingToURL(f func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */)) {
	d._URLSessionAssetDownloadTaskDidFinishDownloadingToURL = f
}

// SetURLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad sets the handler for the URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad delegate method.
//
// Tells the delegate that a download task loaded a new time range.
func (d *AssetDownloadDelegate) SetURLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad(f func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, timeRange TimeRange /* not a class type */, loadedTimeRanges []foundation.Value, timeRangeExpectedToLoad TimeRange /* not a class type */)) {
	d._URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad = f
}

// SetURLSessionAssetDownloadTaskDidReceiveMetricEvent sets the handler for the URLSessionAssetDownloadTaskDidReceiveMetricEvent delegate method.
//
// Sent when a download task receives an AVMetricEvent.
func (d *AssetDownloadDelegate) SetURLSessionAssetDownloadTaskDidReceiveMetricEvent(f func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, metricEvent IAVMetricEvent)) {
	d._URLSessionAssetDownloadTaskDidReceiveMetricEvent = f
}

// SetURLSessionAssetDownloadTaskDidResolveMediaSelection sets the handler for the URLSessionAssetDownloadTaskDidResolveMediaSelection delegate method.
//
// Tells the delegate that a download task resolved the media selection to download, including any automatic selections.
func (d *AssetDownloadDelegate) SetURLSessionAssetDownloadTaskDidResolveMediaSelection(f func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, resolvedMediaSelection IAVMediaSelection)) {
	d._URLSessionAssetDownloadTaskDidResolveMediaSelection = f
}

// SetURLSessionAssetDownloadTaskWillDownloadToURL sets the handler for the URLSessionAssetDownloadTaskWillDownloadToURL delegate method.
//
// Tells the delegate when a download task determines its download location.
func (d *AssetDownloadDelegate) SetURLSessionAssetDownloadTaskWillDownloadToURL(f func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */)) {
	d._URLSessionAssetDownloadTaskWillDownloadToURL = f
}

// SetURLSessionAssetDownloadTaskWillDownloadVariants sets the handler for the URLSessionAssetDownloadTaskWillDownloadVariants delegate method.
//
// Tells the delegate that a download task completed variant selection.
func (d *AssetDownloadDelegate) SetURLSessionAssetDownloadTaskWillDownloadVariants(f func(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, variants []AssetVariant)) {
	d._URLSessionAssetDownloadTaskWillDownloadVariants = f
}

// URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection implements the PAssetDownloadDelegate interface.
func (d *AssetDownloadDelegate) URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, mediaSelection IAVMediaSelection) {
	if d._URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection != nil {
		d._URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection(session, aggregateAssetDownloadTask, mediaSelection)
	}
}

// HasURLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection returns true if a handler for URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection has been set.
func (d *AssetDownloadDelegate) HasURLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection() bool {
	return d._URLSessionAggregateAssetDownloadTaskDidCompleteForMediaSelection != nil
}

// URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection implements the PAssetDownloadDelegate interface.
func (d *AssetDownloadDelegate) URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, timeRange TimeRange /* not a class type */, loadedTimeRanges []foundation.Value, timeRangeExpectedToLoad TimeRange /* not a class type */, mediaSelection IAVMediaSelection) {
	if d._URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection != nil {
		d._URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection(session, aggregateAssetDownloadTask, timeRange, loadedTimeRanges, timeRangeExpectedToLoad, mediaSelection)
	}
}

// HasURLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection returns true if a handler for URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection has been set.
func (d *AssetDownloadDelegate) HasURLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection() bool {
	return d._URLSessionAggregateAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoadForMediaSelection != nil
}

// URLSessionAggregateAssetDownloadTaskWillDownloadToURL implements the PAssetDownloadDelegate interface.
func (d *AssetDownloadDelegate) URLSessionAggregateAssetDownloadTaskWillDownloadToURL(session foundation.URLSession, aggregateAssetDownloadTask IAVAggregateAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */) {
	if d._URLSessionAggregateAssetDownloadTaskWillDownloadToURL != nil {
		d._URLSessionAggregateAssetDownloadTaskWillDownloadToURL(session, aggregateAssetDownloadTask, location)
	}
}

// HasURLSessionAggregateAssetDownloadTaskWillDownloadToURL returns true if a handler for URLSessionAggregateAssetDownloadTaskWillDownloadToURL has been set.
func (d *AssetDownloadDelegate) HasURLSessionAggregateAssetDownloadTaskWillDownloadToURL() bool {
	return d._URLSessionAggregateAssetDownloadTaskWillDownloadToURL != nil
}

// URLSessionAssetDownloadTaskDidFinishDownloadingToURL implements the PAssetDownloadDelegate interface.
func (d *AssetDownloadDelegate) URLSessionAssetDownloadTaskDidFinishDownloadingToURL(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */) {
	if d._URLSessionAssetDownloadTaskDidFinishDownloadingToURL != nil {
		d._URLSessionAssetDownloadTaskDidFinishDownloadingToURL(session, assetDownloadTask, location)
	}
}

// HasURLSessionAssetDownloadTaskDidFinishDownloadingToURL returns true if a handler for URLSessionAssetDownloadTaskDidFinishDownloadingToURL has been set.
func (d *AssetDownloadDelegate) HasURLSessionAssetDownloadTaskDidFinishDownloadingToURL() bool {
	return d._URLSessionAssetDownloadTaskDidFinishDownloadingToURL != nil
}

// URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad implements the PAssetDownloadDelegate interface.
func (d *AssetDownloadDelegate) URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, timeRange TimeRange /* not a class type */, loadedTimeRanges []foundation.Value, timeRangeExpectedToLoad TimeRange /* not a class type */) {
	if d._URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad != nil {
		d._URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad(session, assetDownloadTask, timeRange, loadedTimeRanges, timeRangeExpectedToLoad)
	}
}

// HasURLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad returns true if a handler for URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad has been set.
func (d *AssetDownloadDelegate) HasURLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad() bool {
	return d._URLSessionAssetDownloadTaskDidLoadTimeRangeTotalTimeRangesLoadedTimeRangeExpectedToLoad != nil
}

// URLSessionAssetDownloadTaskDidReceiveMetricEvent implements the PAssetDownloadDelegate interface.
func (d *AssetDownloadDelegate) URLSessionAssetDownloadTaskDidReceiveMetricEvent(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, metricEvent IAVMetricEvent) {
	if d._URLSessionAssetDownloadTaskDidReceiveMetricEvent != nil {
		d._URLSessionAssetDownloadTaskDidReceiveMetricEvent(session, assetDownloadTask, metricEvent)
	}
}

// HasURLSessionAssetDownloadTaskDidReceiveMetricEvent returns true if a handler for URLSessionAssetDownloadTaskDidReceiveMetricEvent has been set.
func (d *AssetDownloadDelegate) HasURLSessionAssetDownloadTaskDidReceiveMetricEvent() bool {
	return d._URLSessionAssetDownloadTaskDidReceiveMetricEvent != nil
}

// URLSessionAssetDownloadTaskDidResolveMediaSelection implements the PAssetDownloadDelegate interface.
func (d *AssetDownloadDelegate) URLSessionAssetDownloadTaskDidResolveMediaSelection(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, resolvedMediaSelection IAVMediaSelection) {
	if d._URLSessionAssetDownloadTaskDidResolveMediaSelection != nil {
		d._URLSessionAssetDownloadTaskDidResolveMediaSelection(session, assetDownloadTask, resolvedMediaSelection)
	}
}

// HasURLSessionAssetDownloadTaskDidResolveMediaSelection returns true if a handler for URLSessionAssetDownloadTaskDidResolveMediaSelection has been set.
func (d *AssetDownloadDelegate) HasURLSessionAssetDownloadTaskDidResolveMediaSelection() bool {
	return d._URLSessionAssetDownloadTaskDidResolveMediaSelection != nil
}

// URLSessionAssetDownloadTaskWillDownloadToURL implements the PAssetDownloadDelegate interface.
func (d *AssetDownloadDelegate) URLSessionAssetDownloadTaskWillDownloadToURL(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, location objc.IObject /* cross-framework: NSURL */) {
	if d._URLSessionAssetDownloadTaskWillDownloadToURL != nil {
		d._URLSessionAssetDownloadTaskWillDownloadToURL(session, assetDownloadTask, location)
	}
}

// HasURLSessionAssetDownloadTaskWillDownloadToURL returns true if a handler for URLSessionAssetDownloadTaskWillDownloadToURL has been set.
func (d *AssetDownloadDelegate) HasURLSessionAssetDownloadTaskWillDownloadToURL() bool {
	return d._URLSessionAssetDownloadTaskWillDownloadToURL != nil
}

// URLSessionAssetDownloadTaskWillDownloadVariants implements the PAssetDownloadDelegate interface.
func (d *AssetDownloadDelegate) URLSessionAssetDownloadTaskWillDownloadVariants(session foundation.URLSession, assetDownloadTask IAVAssetDownloadTask, variants []AssetVariant) {
	if d._URLSessionAssetDownloadTaskWillDownloadVariants != nil {
		d._URLSessionAssetDownloadTaskWillDownloadVariants(session, assetDownloadTask, variants)
	}
}

// HasURLSessionAssetDownloadTaskWillDownloadVariants returns true if a handler for URLSessionAssetDownloadTaskWillDownloadVariants has been set.
func (d *AssetDownloadDelegate) HasURLSessionAssetDownloadTaskWillDownloadVariants() bool {
	return d._URLSessionAssetDownloadTaskWillDownloadVariants != nil
}
