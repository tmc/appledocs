// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PURLSessionDataDelegate is the NSURLSessionDataDelegate protocol interface.
//
// A protocol that defines methods that URL session instances call on their delegates to handle task-level events specific to data and upload tasks.
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
// See: doc://com.apple.foundation/documentation/Foundation/URLSessionDataDelegate
type PURLSessionDataDelegate interface {
	// Optional methods
	URLSessionDataTaskDidBecomeDownloadTask(session IURLSession, dataTask IURLSessionDataTask, downloadTask IURLSessionDownloadTask)
	HasURLSessionDataTaskDidBecomeDownloadTask() bool
	URLSessionDataTaskDidBecomeStreamTask(session IURLSession, dataTask IURLSessionDataTask, streamTask IURLSessionStreamTask)
	HasURLSessionDataTaskDidBecomeStreamTask() bool
	URLSessionDataTaskDidReceiveData(session IURLSession, dataTask IURLSessionDataTask, data IData)
	HasURLSessionDataTaskDidReceiveData() bool
}

// URLSessionDataDelegate is a delegate implementation builder for the PURLSessionDataDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type URLSessionDataDelegate struct {
	_URLSessionDataTaskDidBecomeDownloadTask func(session IURLSession, dataTask IURLSessionDataTask, downloadTask IURLSessionDownloadTask)
	_URLSessionDataTaskDidBecomeStreamTask func(session IURLSession, dataTask IURLSessionDataTask, streamTask IURLSessionStreamTask)
	_URLSessionDataTaskDidReceiveData func(session IURLSession, dataTask IURLSessionDataTask, data IData)
}

// SetURLSessionDataTaskDidBecomeDownloadTask sets the handler for the URLSessionDataTaskDidBecomeDownloadTask delegate method.
//
// Tells the delegate that the data task was changed to a download task.
func (d *URLSessionDataDelegate) SetURLSessionDataTaskDidBecomeDownloadTask(f func(session IURLSession, dataTask IURLSessionDataTask, downloadTask IURLSessionDownloadTask)) {
	d._URLSessionDataTaskDidBecomeDownloadTask = f
}

// SetURLSessionDataTaskDidBecomeStreamTask sets the handler for the URLSessionDataTaskDidBecomeStreamTask delegate method.
//
// Tells the delegate that the data task was changed to a stream task.
func (d *URLSessionDataDelegate) SetURLSessionDataTaskDidBecomeStreamTask(f func(session IURLSession, dataTask IURLSessionDataTask, streamTask IURLSessionStreamTask)) {
	d._URLSessionDataTaskDidBecomeStreamTask = f
}

// SetURLSessionDataTaskDidReceiveData sets the handler for the URLSessionDataTaskDidReceiveData delegate method.
//
// Tells the delegate that the data task has received some of the expected data.
func (d *URLSessionDataDelegate) SetURLSessionDataTaskDidReceiveData(f func(session IURLSession, dataTask IURLSessionDataTask, data IData)) {
	d._URLSessionDataTaskDidReceiveData = f
}

// URLSessionDataTaskDidBecomeDownloadTask implements the PURLSessionDataDelegate interface.
func (d *URLSessionDataDelegate) URLSessionDataTaskDidBecomeDownloadTask(session IURLSession, dataTask IURLSessionDataTask, downloadTask IURLSessionDownloadTask) {
	if d._URLSessionDataTaskDidBecomeDownloadTask != nil {
		d._URLSessionDataTaskDidBecomeDownloadTask(session, dataTask, downloadTask)
	}
}

// HasURLSessionDataTaskDidBecomeDownloadTask returns true if a handler for URLSessionDataTaskDidBecomeDownloadTask has been set.
func (d *URLSessionDataDelegate) HasURLSessionDataTaskDidBecomeDownloadTask() bool {
	return d._URLSessionDataTaskDidBecomeDownloadTask != nil
}

// URLSessionDataTaskDidBecomeStreamTask implements the PURLSessionDataDelegate interface.
func (d *URLSessionDataDelegate) URLSessionDataTaskDidBecomeStreamTask(session IURLSession, dataTask IURLSessionDataTask, streamTask IURLSessionStreamTask) {
	if d._URLSessionDataTaskDidBecomeStreamTask != nil {
		d._URLSessionDataTaskDidBecomeStreamTask(session, dataTask, streamTask)
	}
}

// HasURLSessionDataTaskDidBecomeStreamTask returns true if a handler for URLSessionDataTaskDidBecomeStreamTask has been set.
func (d *URLSessionDataDelegate) HasURLSessionDataTaskDidBecomeStreamTask() bool {
	return d._URLSessionDataTaskDidBecomeStreamTask != nil
}

// URLSessionDataTaskDidReceiveData implements the PURLSessionDataDelegate interface.
func (d *URLSessionDataDelegate) URLSessionDataTaskDidReceiveData(session IURLSession, dataTask IURLSessionDataTask, data IData) {
	if d._URLSessionDataTaskDidReceiveData != nil {
		d._URLSessionDataTaskDidReceiveData(session, dataTask, data)
	}
}

// HasURLSessionDataTaskDidReceiveData returns true if a handler for URLSessionDataTaskDidReceiveData has been set.
func (d *URLSessionDataDelegate) HasURLSessionDataTaskDidReceiveData() bool {
	return d._URLSessionDataTaskDidReceiveData != nil
}
