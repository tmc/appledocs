// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PBADownloadManagerDelegate is the BADownloadManagerDelegate protocol interface.
//
// An interface for reacting to asset download events and processing concluded downloads.
//
// Availability:
//   - Mac Catalyst 16.1+
//   - iOS 16.1+
//   - iPadOS 16.1+
//   - macOS 13.0+
//   - tvOS 18.4+
//   - visionOS 2.4+
//
// See: doc://com.apple.backgroundassets/documentation/BackgroundAssets/BADownloadManagerDelegate
type PBADownloadManagerDelegate interface {
	// Optional methods
	DownloadDidReceiveChallengeCompletionHandler(download IBADownload, challenge foundation.URLAuthenticationChallenge, completionHandler unsafe.Pointer)
	HasDownloadDidReceiveChallengeCompletionHandler() bool
	DownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite(download IBADownload, bytesWritten int64, totalBytesWritten int64, totalExpectedBytes int64)
	HasDownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite() bool
	DownloadFailedWithError(download IBADownload, error_ objc.IObject /* cross-framework: Error */)
	HasDownloadFailedWithError() bool
	DownloadFinishedWithFileURL(download IBADownload, fileURL objc.IObject /* cross-framework: NSURL */)
	HasDownloadFinishedWithFileURL() bool
	DownloadDidBegin(download IBADownload)
	HasDownloadDidBegin() bool
	DownloadDidPause(download IBADownload)
	HasDownloadDidPause() bool
}

// BADownloadManagerDelegate is a delegate implementation builder for the PBADownloadManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type BADownloadManagerDelegate struct {
	_DownloadDidReceiveChallengeCompletionHandler func(download IBADownload, challenge foundation.URLAuthenticationChallenge, completionHandler unsafe.Pointer)
	_DownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite func(download IBADownload, bytesWritten int64, totalBytesWritten int64, totalExpectedBytes int64)
	_DownloadFailedWithError func(download IBADownload, error_ objc.IObject /* cross-framework: Error */)
	_DownloadFinishedWithFileURL func(download IBADownload, fileURL objc.IObject /* cross-framework: NSURL */)
	_DownloadDidBegin func(download IBADownload)
	_DownloadDidPause func(download IBADownload)
}

// SetDownloadDidReceiveChallengeCompletionHandler sets the handler for the DownloadDidReceiveChallengeCompletionHandler delegate method.
//
// Tells the delegate to resolve the specified URL authentication challenge.
func (d *BADownloadManagerDelegate) SetDownloadDidReceiveChallengeCompletionHandler(f func(download IBADownload, challenge foundation.URLAuthenticationChallenge, completionHandler unsafe.Pointer)) {
	d._DownloadDidReceiveChallengeCompletionHandler = f
}

// SetDownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite sets the handler for the DownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite delegate method.
//
// Informs the delegate about the progress of the specified asset download.
func (d *BADownloadManagerDelegate) SetDownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite(f func(download IBADownload, bytesWritten int64, totalBytesWritten int64, totalExpectedBytes int64)) {
	d._DownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite = f
}

// SetDownloadFailedWithError sets the handler for the DownloadFailedWithError delegate method.
//
// Informs the delegate about a failed asset download.
func (d *BADownloadManagerDelegate) SetDownloadFailedWithError(f func(download IBADownload, error_ objc.IObject /* cross-framework: Error */)) {
	d._DownloadFailedWithError = f
}

// SetDownloadFinishedWithFileURL sets the handler for the DownloadFinishedWithFileURL delegate method.
//
// Informs the delegate about a finished asset download and provides the on-disk location.
func (d *BADownloadManagerDelegate) SetDownloadFinishedWithFileURL(f func(download IBADownload, fileURL objc.IObject /* cross-framework: NSURL */)) {
	d._DownloadFinishedWithFileURL = f
}

// SetDownloadDidBegin sets the handler for the DownloadDidBegin delegate method.
//
// Informs the delegate about a started asset download.
func (d *BADownloadManagerDelegate) SetDownloadDidBegin(f func(download IBADownload)) {
	d._DownloadDidBegin = f
}

// SetDownloadDidPause sets the handler for the DownloadDidPause delegate method.
//
// Informs the delegate about a paused asset download.
func (d *BADownloadManagerDelegate) SetDownloadDidPause(f func(download IBADownload)) {
	d._DownloadDidPause = f
}

// DownloadDidReceiveChallengeCompletionHandler implements the PBADownloadManagerDelegate interface.
func (d *BADownloadManagerDelegate) DownloadDidReceiveChallengeCompletionHandler(download IBADownload, challenge foundation.URLAuthenticationChallenge, completionHandler unsafe.Pointer) {
	if d._DownloadDidReceiveChallengeCompletionHandler != nil {
		d._DownloadDidReceiveChallengeCompletionHandler(download, challenge, completionHandler)
	}
}

// HasDownloadDidReceiveChallengeCompletionHandler returns true if a handler for DownloadDidReceiveChallengeCompletionHandler has been set.
func (d *BADownloadManagerDelegate) HasDownloadDidReceiveChallengeCompletionHandler() bool {
	return d._DownloadDidReceiveChallengeCompletionHandler != nil
}

// DownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite implements the PBADownloadManagerDelegate interface.
func (d *BADownloadManagerDelegate) DownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite(download IBADownload, bytesWritten int64, totalBytesWritten int64, totalExpectedBytes int64) {
	if d._DownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite != nil {
		d._DownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite(download, bytesWritten, totalBytesWritten, totalExpectedBytes)
	}
}

// HasDownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite returns true if a handler for DownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite has been set.
func (d *BADownloadManagerDelegate) HasDownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite() bool {
	return d._DownloadDidWriteBytesTotalBytesWrittenTotalBytesExpectedToWrite != nil
}

// DownloadFailedWithError implements the PBADownloadManagerDelegate interface.
func (d *BADownloadManagerDelegate) DownloadFailedWithError(download IBADownload, error_ objc.IObject /* cross-framework: Error */) {
	if d._DownloadFailedWithError != nil {
		d._DownloadFailedWithError(download, error_)
	}
}

// HasDownloadFailedWithError returns true if a handler for DownloadFailedWithError has been set.
func (d *BADownloadManagerDelegate) HasDownloadFailedWithError() bool {
	return d._DownloadFailedWithError != nil
}

// DownloadFinishedWithFileURL implements the PBADownloadManagerDelegate interface.
func (d *BADownloadManagerDelegate) DownloadFinishedWithFileURL(download IBADownload, fileURL objc.IObject /* cross-framework: NSURL */) {
	if d._DownloadFinishedWithFileURL != nil {
		d._DownloadFinishedWithFileURL(download, fileURL)
	}
}

// HasDownloadFinishedWithFileURL returns true if a handler for DownloadFinishedWithFileURL has been set.
func (d *BADownloadManagerDelegate) HasDownloadFinishedWithFileURL() bool {
	return d._DownloadFinishedWithFileURL != nil
}

// DownloadDidBegin implements the PBADownloadManagerDelegate interface.
func (d *BADownloadManagerDelegate) DownloadDidBegin(download IBADownload) {
	if d._DownloadDidBegin != nil {
		d._DownloadDidBegin(download)
	}
}

// HasDownloadDidBegin returns true if a handler for DownloadDidBegin has been set.
func (d *BADownloadManagerDelegate) HasDownloadDidBegin() bool {
	return d._DownloadDidBegin != nil
}

// DownloadDidPause implements the PBADownloadManagerDelegate interface.
func (d *BADownloadManagerDelegate) DownloadDidPause(download IBADownload) {
	if d._DownloadDidPause != nil {
		d._DownloadDidPause(download)
	}
}

// HasDownloadDidPause returns true if a handler for DownloadDidPause has been set.
func (d *BADownloadManagerDelegate) HasDownloadDidPause() bool {
	return d._DownloadDidPause != nil
}
