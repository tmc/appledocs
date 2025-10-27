// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
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

// URLSessionDataDelegateObject wraps an existing Objective-C object that conforms to the PURLSessionDataDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type URLSessionDataDelegateObject struct {
	objectivec.Object
}

// NewURLSessionDataDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSURLSessionDataDelegate protocol.
func NewURLSessionDataDelegateObject(obj objectivec.Object) *URLSessionDataDelegateObject {
	return &URLSessionDataDelegateObject{obj}
}

// Make sure URLSessionDataDelegateObject implements PURLSessionDataDelegate.
var _ PURLSessionDataDelegate = (*URLSessionDataDelegateObject)(nil)

// URLSessionDataTaskDidBecomeDownloadTask implements the PURLSessionDataDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionDataDelegateObject) URLSessionDataTaskDidBecomeDownloadTask(session IURLSession, dataTask IURLSessionDataTask, downloadTask IURLSessionDownloadTask) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:dataTask:didBecomeDownloadTask:"), session, dataTask, downloadTask)
}

// HasURLSessionDataTaskDidBecomeDownloadTask returns true; this is a placeholder for optional method checks.
func (o *URLSessionDataDelegateObject) HasURLSessionDataTaskDidBecomeDownloadTask() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionDataTaskDidBecomeStreamTask implements the PURLSessionDataDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionDataDelegateObject) URLSessionDataTaskDidBecomeStreamTask(session IURLSession, dataTask IURLSessionDataTask, streamTask IURLSessionStreamTask) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:dataTask:didBecomeStreamTask:"), session, dataTask, streamTask)
}

// HasURLSessionDataTaskDidBecomeStreamTask returns true; this is a placeholder for optional method checks.
func (o *URLSessionDataDelegateObject) HasURLSessionDataTaskDidBecomeStreamTask() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// URLSessionDataTaskDidReceiveData implements the PURLSessionDataDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *URLSessionDataDelegateObject) URLSessionDataTaskDidReceiveData(session IURLSession, dataTask IURLSessionDataTask, data IData) {
	objc.Send[objc.ID](o.ID, objc.Sel("URLSession:dataTask:didReceiveData:"), session, dataTask, data)
}

// HasURLSessionDataTaskDidReceiveData returns true; this is a placeholder for optional method checks.
func (o *URLSessionDataDelegateObject) HasURLSessionDataTaskDidReceiveData() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
