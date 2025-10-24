// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PFilePromiseProviderDelegate is the NSFilePromiseProviderDelegate protocol interface.
//
// A set of methods that provides the name of the promised file and writes the file to the destination directory when the file promise is fulfilled.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSFilePromiseProviderDelegate
type PFilePromiseProviderDelegate interface {
	// Required methods
	FilePromiseProviderFileNameForType(filePromiseProvider IFilePromiseProvider, fileType objc.IObject /* cross-framework: NSString */) foundation.String/* debug [protocol_interface/required_method]: FilePromiseProviderFileNameForType */
	FilePromiseProviderWritePromiseToURLCompletionHandler(filePromiseProvider IFilePromiseProvider, url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: FilePromiseProviderWritePromiseToURLCompletionHandler */
	// Optional methods
	OperationQueueForFilePromiseProvider(filePromiseProvider IFilePromiseProvider) foundation.OperationQueue
	HasOperationQueueForFilePromiseProvider() bool
}

// FilePromiseProviderDelegate is a delegate implementation builder for the PFilePromiseProviderDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type FilePromiseProviderDelegate struct {
	_OperationQueueForFilePromiseProvider func(filePromiseProvider IFilePromiseProvider) foundation.OperationQueue
	_FilePromiseProviderFileNameForType func(filePromiseProvider IFilePromiseProvider, fileType objc.IObject /* cross-framework: NSString */) foundation.String
	_FilePromiseProviderWritePromiseToURLCompletionHandler func(filePromiseProvider IFilePromiseProvider, url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)
}

// SetOperationQueueForFilePromiseProvider sets the handler for the OperationQueueForFilePromiseProvider delegate method.
//
// Returns the operation queue from which to issue the write request.
func (d *FilePromiseProviderDelegate) SetOperationQueueForFilePromiseProvider(f func(filePromiseProvider IFilePromiseProvider) foundation.OperationQueue) {
	d._OperationQueueForFilePromiseProvider = f
}

// SetFilePromiseProviderFileNameForType sets the handler for the FilePromiseProviderFileNameForType delegate method.
//
// Provides the drag destination file’s name.
func (d *FilePromiseProviderDelegate) SetFilePromiseProviderFileNameForType(f func(filePromiseProvider IFilePromiseProvider, fileType objc.IObject /* cross-framework: NSString */) foundation.String) {
	d._FilePromiseProviderFileNameForType = f
}

// SetFilePromiseProviderWritePromiseToURLCompletionHandler sets the handler for the FilePromiseProviderWritePromiseToURLCompletionHandler delegate method.
//
// Writes the contents of a promise to the specified URL.
func (d *FilePromiseProviderDelegate) SetFilePromiseProviderWritePromiseToURLCompletionHandler(f func(filePromiseProvider IFilePromiseProvider, url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)) {
	d._FilePromiseProviderWritePromiseToURLCompletionHandler = f
}

// OperationQueueForFilePromiseProvider implements the PFilePromiseProviderDelegate interface.
func (d *FilePromiseProviderDelegate) OperationQueueForFilePromiseProvider(filePromiseProvider IFilePromiseProvider) foundation.OperationQueue {
	if d._OperationQueueForFilePromiseProvider != nil {
		return d._OperationQueueForFilePromiseProvider(filePromiseProvider)
	}
	var zero foundation.OperationQueue
	return zero
}

// HasOperationQueueForFilePromiseProvider returns true if a handler for OperationQueueForFilePromiseProvider has been set.
func (d *FilePromiseProviderDelegate) HasOperationQueueForFilePromiseProvider() bool {
	return d._OperationQueueForFilePromiseProvider != nil
}

// FilePromiseProviderFileNameForType implements the PFilePromiseProviderDelegate interface.
func (d *FilePromiseProviderDelegate) FilePromiseProviderFileNameForType(filePromiseProvider IFilePromiseProvider, fileType objc.IObject /* cross-framework: NSString */) foundation.String {
	if d._FilePromiseProviderFileNameForType != nil {
		return d._FilePromiseProviderFileNameForType(filePromiseProvider, fileType)
	}
	var zero foundation.String
	return zero
}

// HasFilePromiseProviderFileNameForType returns true if a handler for FilePromiseProviderFileNameForType has been set.
func (d *FilePromiseProviderDelegate) HasFilePromiseProviderFileNameForType() bool {
	return d._FilePromiseProviderFileNameForType != nil
}

// FilePromiseProviderWritePromiseToURLCompletionHandler implements the PFilePromiseProviderDelegate interface.
func (d *FilePromiseProviderDelegate) FilePromiseProviderWritePromiseToURLCompletionHandler(filePromiseProvider IFilePromiseProvider, url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	if d._FilePromiseProviderWritePromiseToURLCompletionHandler != nil {
		d._FilePromiseProviderWritePromiseToURLCompletionHandler(filePromiseProvider, url, completionHandler)
	}
}

// HasFilePromiseProviderWritePromiseToURLCompletionHandler returns true if a handler for FilePromiseProviderWritePromiseToURLCompletionHandler has been set.
func (d *FilePromiseProviderDelegate) HasFilePromiseProviderWritePromiseToURLCompletionHandler() bool {
	return d._FilePromiseProviderWritePromiseToURLCompletionHandler != nil
}
