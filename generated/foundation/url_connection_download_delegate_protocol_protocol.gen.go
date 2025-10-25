// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PURLConnectionDownloadDelegate is the NSURLConnectionDownloadDelegate protocol interface.
//
// A protocol that delegates of a URL connection created with Newsstand Kit implement to receive data associated with a download.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/NSURLConnectionDownloadDelegate
type PURLConnectionDownloadDelegate interface {
	// Required methods
	ConnectionDidFinishDownloadingDestinationURL(connection IURLConnection, destinationURL IURL)/* debug [protocol_interface/required_method]: ConnectionDidFinishDownloadingDestinationURL */
	// Optional methods
	ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes(connection IURLConnection, bytesWritten objectivec.IObject, totalBytesWritten objectivec.IObject, expectedTotalBytes objectivec.IObject)
	HasConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes() bool
	ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes(connection IURLConnection, totalBytesWritten objectivec.IObject, expectedTotalBytes objectivec.IObject)
	HasConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes() bool
}

// URLConnectionDownloadDelegate is a delegate implementation builder for the PURLConnectionDownloadDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type URLConnectionDownloadDelegate struct {
	_ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes func(connection IURLConnection, bytesWritten objectivec.IObject, totalBytesWritten objectivec.IObject, expectedTotalBytes objectivec.IObject)
	_ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes func(connection IURLConnection, totalBytesWritten objectivec.IObject, expectedTotalBytes objectivec.IObject)
	_ConnectionDidFinishDownloadingDestinationURL func(connection IURLConnection, destinationURL IURL)
}

// SetConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes sets the handler for the ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes delegate method.
//
// Sent to the delegate to deliver progress information for a download of a URL asset to a destination file.
func (d *URLConnectionDownloadDelegate) SetConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes(f func(connection IURLConnection, bytesWritten objectivec.IObject, totalBytesWritten objectivec.IObject, expectedTotalBytes objectivec.IObject)) {
	d._ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes = f
}

// SetConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes sets the handler for the ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes delegate method.
//
// Sent to the delegate when an URL connection resumes downloading a URL asset that was earlier suspended.
func (d *URLConnectionDownloadDelegate) SetConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes(f func(connection IURLConnection, totalBytesWritten objectivec.IObject, expectedTotalBytes objectivec.IObject)) {
	d._ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes = f
}

// SetConnectionDidFinishDownloadingDestinationURL sets the handler for the ConnectionDidFinishDownloadingDestinationURL delegate method.
//
// Sent to the delegate when the URL connection has successfully downloaded the URL asset to a destination file.
func (d *URLConnectionDownloadDelegate) SetConnectionDidFinishDownloadingDestinationURL(f func(connection IURLConnection, destinationURL IURL)) {
	d._ConnectionDidFinishDownloadingDestinationURL = f
}

// ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes implements the PURLConnectionDownloadDelegate interface.
func (d *URLConnectionDownloadDelegate) ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes(connection IURLConnection, bytesWritten objectivec.IObject, totalBytesWritten objectivec.IObject, expectedTotalBytes objectivec.IObject) {
	if d._ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes != nil {
		d._ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes(connection, bytesWritten, totalBytesWritten, expectedTotalBytes)
	}
}

// HasConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes returns true if a handler for ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes has been set.
func (d *URLConnectionDownloadDelegate) HasConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes() bool {
	return d._ConnectionDidWriteDataTotalBytesWrittenExpectedTotalBytes != nil
}

// ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes implements the PURLConnectionDownloadDelegate interface.
func (d *URLConnectionDownloadDelegate) ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes(connection IURLConnection, totalBytesWritten objectivec.IObject, expectedTotalBytes objectivec.IObject) {
	if d._ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes != nil {
		d._ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes(connection, totalBytesWritten, expectedTotalBytes)
	}
}

// HasConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes returns true if a handler for ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes has been set.
func (d *URLConnectionDownloadDelegate) HasConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes() bool {
	return d._ConnectionDidResumeDownloadingTotalBytesWrittenExpectedTotalBytes != nil
}

// ConnectionDidFinishDownloadingDestinationURL implements the PURLConnectionDownloadDelegate interface.
func (d *URLConnectionDownloadDelegate) ConnectionDidFinishDownloadingDestinationURL(connection IURLConnection, destinationURL IURL) {
	if d._ConnectionDidFinishDownloadingDestinationURL != nil {
		d._ConnectionDidFinishDownloadingDestinationURL(connection, destinationURL)
	}
}

// HasConnectionDidFinishDownloadingDestinationURL returns true if a handler for ConnectionDidFinishDownloadingDestinationURL has been set.
func (d *URLConnectionDownloadDelegate) HasConnectionDidFinishDownloadingDestinationURL() bool {
	return d._ConnectionDidFinishDownloadingDestinationURL != nil
}
