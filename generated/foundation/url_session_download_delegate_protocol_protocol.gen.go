// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PURLSessionDownloadDelegate is the NSURLSessionDownloadDelegate protocol interface.
//
// A protocol that defines methods that URL session instances call on their delegates to handle task-level events specific to download tasks.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.9+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/URLSessionDownloadDelegate
type PURLSessionDownloadDelegate interface {
	// Required methods
	URLSessionDownloadTaskDidFinishDownloadingToURL(session IURLSession, downloadTask IURLSessionDownloadTask, location IURL)
	// Optional methods
	URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes(session IURLSession, downloadTask IURLSessionDownloadTask, fileOffset int64, expectedTotalBytes int64)
	HasURLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes() bool
	URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite(session IURLSession, downloadTask IURLSessionDownloadTask, bytesWritten int64, totalBytesWritten int64, totalBytesExpectedToWrite int64)
	HasURLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite() bool
}

// URLSessionDownloadDelegate is a delegate implementation builder for the PURLSessionDownloadDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type URLSessionDownloadDelegate struct {
	_URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes func(session IURLSession, downloadTask IURLSessionDownloadTask, fileOffset int64, expectedTotalBytes int64)
	_URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite func(session IURLSession, downloadTask IURLSessionDownloadTask, bytesWritten int64, totalBytesWritten int64, totalBytesExpectedToWrite int64)
	_URLSessionDownloadTaskDidFinishDownloadingToURL func(session IURLSession, downloadTask IURLSessionDownloadTask, location IURL)
}

// SetURLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes sets the handler for the URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes delegate method.
//
// Tells the delegate that the download task has resumed downloading.
func (d *URLSessionDownloadDelegate) SetURLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes(f func(session IURLSession, downloadTask IURLSessionDownloadTask, fileOffset int64, expectedTotalBytes int64)) {
	d._URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes = f
}

// SetURLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite sets the handler for the URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite delegate method.
//
// Periodically informs the delegate about the download’s progress.
func (d *URLSessionDownloadDelegate) SetURLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite(f func(session IURLSession, downloadTask IURLSessionDownloadTask, bytesWritten int64, totalBytesWritten int64, totalBytesExpectedToWrite int64)) {
	d._URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite = f
}

// SetURLSessionDownloadTaskDidFinishDownloadingToURL sets the handler for the URLSessionDownloadTaskDidFinishDownloadingToURL delegate method.
//
// Tells the delegate that a download task has finished downloading.
func (d *URLSessionDownloadDelegate) SetURLSessionDownloadTaskDidFinishDownloadingToURL(f func(session IURLSession, downloadTask IURLSessionDownloadTask, location IURL)) {
	d._URLSessionDownloadTaskDidFinishDownloadingToURL = f
}

// URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes implements the PURLSessionDownloadDelegate interface.
func (d *URLSessionDownloadDelegate) URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes(session IURLSession, downloadTask IURLSessionDownloadTask, fileOffset int64, expectedTotalBytes int64) {
	if d._URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes != nil {
		d._URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes(session, downloadTask, fileOffset, expectedTotalBytes)
	}
}

// HasURLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes returns true if a handler for URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes has been set.
func (d *URLSessionDownloadDelegate) HasURLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes() bool {
	return d._URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes != nil
}

// URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite implements the PURLSessionDownloadDelegate interface.
func (d *URLSessionDownloadDelegate) URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite(session IURLSession, downloadTask IURLSessionDownloadTask, bytesWritten int64, totalBytesWritten int64, totalBytesExpectedToWrite int64) {
	if d._URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite != nil {
		d._URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite(session, downloadTask, bytesWritten, totalBytesWritten, totalBytesExpectedToWrite)
	}
}

// HasURLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite returns true if a handler for URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite has been set.
func (d *URLSessionDownloadDelegate) HasURLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite() bool {
	return d._URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite != nil
}

// URLSessionDownloadTaskDidFinishDownloadingToURL implements the PURLSessionDownloadDelegate interface.
func (d *URLSessionDownloadDelegate) URLSessionDownloadTaskDidFinishDownloadingToURL(session IURLSession, downloadTask IURLSessionDownloadTask, location IURL) {
	if d._URLSessionDownloadTaskDidFinishDownloadingToURL != nil {
		d._URLSessionDownloadTaskDidFinishDownloadingToURL(session, downloadTask, location)
	}
}

// HasURLSessionDownloadTaskDidFinishDownloadingToURL returns true if a handler for URLSessionDownloadTaskDidFinishDownloadingToURL has been set.
func (d *URLSessionDownloadDelegate) HasURLSessionDownloadTaskDidFinishDownloadingToURL() bool {
	return d._URLSessionDownloadTaskDidFinishDownloadingToURL != nil
}

// URLSessionDownloadDelegateObject wraps an existing Objective-C object that conforms to the PURLSessionDownloadDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type URLSessionDownloadDelegateObject struct {
	objectivec.Object
}

// NewURLSessionDownloadDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSURLSessionDownloadDelegate protocol.
func NewURLSessionDownloadDelegateObject(obj objectivec.Object) *URLSessionDownloadDelegateObject {
	return &URLSessionDownloadDelegateObject{obj}
}

// Make sure URLSessionDownloadDelegateObject implements PURLSessionDownloadDelegate.
var _ PURLSessionDownloadDelegate = (*URLSessionDownloadDelegateObject)(nil)

// URLSessionDownloadTaskDidFinishDownloadingToURL implements the PURLSessionDownloadDelegate interface.
// This required method is always available on objects conforming to URLSessionDownloadTaskDidFinishDownloadingToURL.
func (o *URLSessionDownloadDelegateObject) URLSessionDownloadTaskDidFinishDownloadingToURL(session IURLSession, downloadTask IURLSessionDownloadTask, location IURL) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:downloadTask:didFinishDownloadingToURL:"), session, downloadTask, location)
}

// URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes implements the PURLSessionDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionDownloadDelegateObject) URLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes(session IURLSession, downloadTask IURLSessionDownloadTask, fileOffset int64, expectedTotalBytes int64) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:downloadTask:didResumeAtOffset:expectedTotalBytes:"), session, downloadTask, fileOffset, expectedTotalBytes)
}

// HasURLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes returns true; this is a placeholder for optional method checks.
func (o *URLSessionDownloadDelegateObject) HasURLSessionDownloadTaskDidResumeAtOffsetExpectedTotalBytes() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite implements the PURLSessionDownloadDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionDownloadDelegateObject) URLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite(session IURLSession, downloadTask IURLSessionDownloadTask, bytesWritten int64, totalBytesWritten int64, totalBytesExpectedToWrite int64) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:downloadTask:didWriteData:totalBytesWritten:totalBytesExpectedToWrite:"), session, downloadTask, bytesWritten, totalBytesWritten, totalBytesExpectedToWrite)
}

// HasURLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite returns true; this is a placeholder for optional method checks.
func (o *URLSessionDownloadDelegateObject) HasURLSessionDownloadTaskDidWriteDataTotalBytesWrittenTotalBytesExpectedToWrite() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
