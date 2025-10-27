// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PFileManagerDelegate is the NSFileManagerDelegate protocol interface.
//
// The interface a file manager’s delegate uses to intervene during operations or if an error occurs.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// See: doc://com.apple.foundation/documentation/Foundation/FileManagerDelegate
type PFileManagerDelegate interface {
	// Optional methods
	FileManagerShouldCopyItemAtURLToURL(fileManager IFileManager, srcURL IURL, dstURL IURL) bool
	HasFileManagerShouldCopyItemAtURLToURL() bool
	FileManagerShouldCopyItemAtPathToPath(fileManager IFileManager, srcPath IString, dstPath IString) bool
	HasFileManagerShouldCopyItemAtPathToPath() bool
	FileManagerShouldLinkItemAtURLToURL(fileManager IFileManager, srcURL IURL, dstURL IURL) bool
	HasFileManagerShouldLinkItemAtURLToURL() bool
	FileManagerShouldLinkItemAtPathToPath(fileManager IFileManager, srcPath IString, dstPath IString) bool
	HasFileManagerShouldLinkItemAtPathToPath() bool
	FileManagerShouldMoveItemAtURLToURL(fileManager IFileManager, srcURL IURL, dstURL IURL) bool
	HasFileManagerShouldMoveItemAtURLToURL() bool
	FileManagerShouldMoveItemAtPathToPath(fileManager IFileManager, srcPath IString, dstPath IString) bool
	HasFileManagerShouldMoveItemAtPathToPath() bool
	FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool
	HasFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL() bool
	FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool
	HasFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath() bool
	FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool
	HasFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL() bool
	FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool
	HasFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath() bool
	FileManagerShouldProceedAfterErrorMovingItemAtURLToURL(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool
	HasFileManagerShouldProceedAfterErrorMovingItemAtURLToURL() bool
	FileManagerShouldProceedAfterErrorMovingItemAtPathToPath(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool
	HasFileManagerShouldProceedAfterErrorMovingItemAtPathToPath() bool
	FileManagerShouldProceedAfterErrorRemovingItemAtURL(fileManager IFileManager, error_ IError, URL IURL) bool
	HasFileManagerShouldProceedAfterErrorRemovingItemAtURL() bool
	FileManagerShouldProceedAfterErrorRemovingItemAtPath(fileManager IFileManager, error_ IError, path IString) bool
	HasFileManagerShouldProceedAfterErrorRemovingItemAtPath() bool
	FileManagerShouldRemoveItemAtURL(fileManager IFileManager, URL IURL) bool
	HasFileManagerShouldRemoveItemAtURL() bool
	FileManagerShouldRemoveItemAtPath(fileManager IFileManager, path IString) bool
	HasFileManagerShouldRemoveItemAtPath() bool
}

// FileManagerDelegate is a delegate implementation builder for the PFileManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type FileManagerDelegate struct {
	_FileManagerShouldCopyItemAtURLToURL func(fileManager IFileManager, srcURL IURL, dstURL IURL) bool
	_FileManagerShouldCopyItemAtPathToPath func(fileManager IFileManager, srcPath IString, dstPath IString) bool
	_FileManagerShouldLinkItemAtURLToURL func(fileManager IFileManager, srcURL IURL, dstURL IURL) bool
	_FileManagerShouldLinkItemAtPathToPath func(fileManager IFileManager, srcPath IString, dstPath IString) bool
	_FileManagerShouldMoveItemAtURLToURL func(fileManager IFileManager, srcURL IURL, dstURL IURL) bool
	_FileManagerShouldMoveItemAtPathToPath func(fileManager IFileManager, srcPath IString, dstPath IString) bool
	_FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL func(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool
	_FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath func(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool
	_FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL func(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool
	_FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath func(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool
	_FileManagerShouldProceedAfterErrorMovingItemAtURLToURL func(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool
	_FileManagerShouldProceedAfterErrorMovingItemAtPathToPath func(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool
	_FileManagerShouldProceedAfterErrorRemovingItemAtURL func(fileManager IFileManager, error_ IError, URL IURL) bool
	_FileManagerShouldProceedAfterErrorRemovingItemAtPath func(fileManager IFileManager, error_ IError, path IString) bool
	_FileManagerShouldRemoveItemAtURL func(fileManager IFileManager, URL IURL) bool
	_FileManagerShouldRemoveItemAtPath func(fileManager IFileManager, path IString) bool
}

// SetFileManagerShouldCopyItemAtURLToURL sets the handler for the FileManagerShouldCopyItemAtURLToURL delegate method.
//
// Asks the delegate if the file manager should copy the specified item to the new URL.
func (d *FileManagerDelegate) SetFileManagerShouldCopyItemAtURLToURL(f func(fileManager IFileManager, srcURL IURL, dstURL IURL) bool) {
	d._FileManagerShouldCopyItemAtURLToURL = f
}

// SetFileManagerShouldCopyItemAtPathToPath sets the handler for the FileManagerShouldCopyItemAtPathToPath delegate method.
//
// Asks the delegate if the file manager should copy the specified item to the new path.
func (d *FileManagerDelegate) SetFileManagerShouldCopyItemAtPathToPath(f func(fileManager IFileManager, srcPath IString, dstPath IString) bool) {
	d._FileManagerShouldCopyItemAtPathToPath = f
}

// SetFileManagerShouldLinkItemAtURLToURL sets the handler for the FileManagerShouldLinkItemAtURLToURL delegate method.
//
// Asks the delegate if a hard link should be created between the items at the two URLs.
func (d *FileManagerDelegate) SetFileManagerShouldLinkItemAtURLToURL(f func(fileManager IFileManager, srcURL IURL, dstURL IURL) bool) {
	d._FileManagerShouldLinkItemAtURLToURL = f
}

// SetFileManagerShouldLinkItemAtPathToPath sets the handler for the FileManagerShouldLinkItemAtPathToPath delegate method.
//
// Asks the delegate if a hard link should be created between the items at the two paths.
func (d *FileManagerDelegate) SetFileManagerShouldLinkItemAtPathToPath(f func(fileManager IFileManager, srcPath IString, dstPath IString) bool) {
	d._FileManagerShouldLinkItemAtPathToPath = f
}

// SetFileManagerShouldMoveItemAtURLToURL sets the handler for the FileManagerShouldMoveItemAtURLToURL delegate method.
//
// Asks the delegate if the file manager should move the specified item to the new URL.
func (d *FileManagerDelegate) SetFileManagerShouldMoveItemAtURLToURL(f func(fileManager IFileManager, srcURL IURL, dstURL IURL) bool) {
	d._FileManagerShouldMoveItemAtURLToURL = f
}

// SetFileManagerShouldMoveItemAtPathToPath sets the handler for the FileManagerShouldMoveItemAtPathToPath delegate method.
//
// Asks the delegate if the file manager should move the specified item to the new path.
func (d *FileManagerDelegate) SetFileManagerShouldMoveItemAtPathToPath(f func(fileManager IFileManager, srcPath IString, dstPath IString) bool) {
	d._FileManagerShouldMoveItemAtPathToPath = f
}

// SetFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL sets the handler for the FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL delegate method.
//
// Asks the delegate if the move operation should continue after an error occurs while copying the item at the specified URL.
func (d *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(f func(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool) {
	d._FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL = f
}

// SetFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath sets the handler for the FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath delegate method.
//
// Asks the delegate if the move operation should continue after an error occurs while copying the item at the specified path.
func (d *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(f func(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool) {
	d._FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath = f
}

// SetFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL sets the handler for the FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL delegate method.
//
// Asks the delegate if the operation should continue after an error occurs while linking to the item at the specified URL.
func (d *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(f func(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool) {
	d._FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL = f
}

// SetFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath sets the handler for the FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath delegate method.
//
// Asks the delegate if the operation should continue after an error occurs while linking to the item at the specified path.
func (d *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(f func(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool) {
	d._FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath = f
}

// SetFileManagerShouldProceedAfterErrorMovingItemAtURLToURL sets the handler for the FileManagerShouldProceedAfterErrorMovingItemAtURLToURL delegate method.
//
// Asks the delegate if the move operation should continue after an error occurs while moving the item at the specified URL.
func (d *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorMovingItemAtURLToURL(f func(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool) {
	d._FileManagerShouldProceedAfterErrorMovingItemAtURLToURL = f
}

// SetFileManagerShouldProceedAfterErrorMovingItemAtPathToPath sets the handler for the FileManagerShouldProceedAfterErrorMovingItemAtPathToPath delegate method.
//
// Asks the delegate if the move operation should continue after an error occurs while moving the item at the specified path.
func (d *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorMovingItemAtPathToPath(f func(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool) {
	d._FileManagerShouldProceedAfterErrorMovingItemAtPathToPath = f
}

// SetFileManagerShouldProceedAfterErrorRemovingItemAtURL sets the handler for the FileManagerShouldProceedAfterErrorRemovingItemAtURL delegate method.
//
// Asks the delegate if the operation should continue after an error occurs while removing the item at the specified URL.
func (d *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorRemovingItemAtURL(f func(fileManager IFileManager, error_ IError, URL IURL) bool) {
	d._FileManagerShouldProceedAfterErrorRemovingItemAtURL = f
}

// SetFileManagerShouldProceedAfterErrorRemovingItemAtPath sets the handler for the FileManagerShouldProceedAfterErrorRemovingItemAtPath delegate method.
//
// Asks the delegate if the operation should continue after an error occurs while removing the item at the specified path.
func (d *FileManagerDelegate) SetFileManagerShouldProceedAfterErrorRemovingItemAtPath(f func(fileManager IFileManager, error_ IError, path IString) bool) {
	d._FileManagerShouldProceedAfterErrorRemovingItemAtPath = f
}

// SetFileManagerShouldRemoveItemAtURL sets the handler for the FileManagerShouldRemoveItemAtURL delegate method.
//
// Asks the delegate whether the item at the specified URL should be deleted.
func (d *FileManagerDelegate) SetFileManagerShouldRemoveItemAtURL(f func(fileManager IFileManager, URL IURL) bool) {
	d._FileManagerShouldRemoveItemAtURL = f
}

// SetFileManagerShouldRemoveItemAtPath sets the handler for the FileManagerShouldRemoveItemAtPath delegate method.
//
// Asks the delegate whether the item at the specified path should be deleted.
func (d *FileManagerDelegate) SetFileManagerShouldRemoveItemAtPath(f func(fileManager IFileManager, path IString) bool) {
	d._FileManagerShouldRemoveItemAtPath = f
}

// FileManagerShouldCopyItemAtURLToURL implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldCopyItemAtURLToURL(fileManager IFileManager, srcURL IURL, dstURL IURL) bool {
	if d._FileManagerShouldCopyItemAtURLToURL != nil {
		return d._FileManagerShouldCopyItemAtURLToURL(fileManager, srcURL, dstURL)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldCopyItemAtURLToURL returns true if a handler for FileManagerShouldCopyItemAtURLToURL has been set.
func (d *FileManagerDelegate) HasFileManagerShouldCopyItemAtURLToURL() bool {
	return d._FileManagerShouldCopyItemAtURLToURL != nil
}

// FileManagerShouldCopyItemAtPathToPath implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldCopyItemAtPathToPath(fileManager IFileManager, srcPath IString, dstPath IString) bool {
	if d._FileManagerShouldCopyItemAtPathToPath != nil {
		return d._FileManagerShouldCopyItemAtPathToPath(fileManager, srcPath, dstPath)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldCopyItemAtPathToPath returns true if a handler for FileManagerShouldCopyItemAtPathToPath has been set.
func (d *FileManagerDelegate) HasFileManagerShouldCopyItemAtPathToPath() bool {
	return d._FileManagerShouldCopyItemAtPathToPath != nil
}

// FileManagerShouldLinkItemAtURLToURL implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldLinkItemAtURLToURL(fileManager IFileManager, srcURL IURL, dstURL IURL) bool {
	if d._FileManagerShouldLinkItemAtURLToURL != nil {
		return d._FileManagerShouldLinkItemAtURLToURL(fileManager, srcURL, dstURL)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldLinkItemAtURLToURL returns true if a handler for FileManagerShouldLinkItemAtURLToURL has been set.
func (d *FileManagerDelegate) HasFileManagerShouldLinkItemAtURLToURL() bool {
	return d._FileManagerShouldLinkItemAtURLToURL != nil
}

// FileManagerShouldLinkItemAtPathToPath implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldLinkItemAtPathToPath(fileManager IFileManager, srcPath IString, dstPath IString) bool {
	if d._FileManagerShouldLinkItemAtPathToPath != nil {
		return d._FileManagerShouldLinkItemAtPathToPath(fileManager, srcPath, dstPath)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldLinkItemAtPathToPath returns true if a handler for FileManagerShouldLinkItemAtPathToPath has been set.
func (d *FileManagerDelegate) HasFileManagerShouldLinkItemAtPathToPath() bool {
	return d._FileManagerShouldLinkItemAtPathToPath != nil
}

// FileManagerShouldMoveItemAtURLToURL implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldMoveItemAtURLToURL(fileManager IFileManager, srcURL IURL, dstURL IURL) bool {
	if d._FileManagerShouldMoveItemAtURLToURL != nil {
		return d._FileManagerShouldMoveItemAtURLToURL(fileManager, srcURL, dstURL)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldMoveItemAtURLToURL returns true if a handler for FileManagerShouldMoveItemAtURLToURL has been set.
func (d *FileManagerDelegate) HasFileManagerShouldMoveItemAtURLToURL() bool {
	return d._FileManagerShouldMoveItemAtURLToURL != nil
}

// FileManagerShouldMoveItemAtPathToPath implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldMoveItemAtPathToPath(fileManager IFileManager, srcPath IString, dstPath IString) bool {
	if d._FileManagerShouldMoveItemAtPathToPath != nil {
		return d._FileManagerShouldMoveItemAtPathToPath(fileManager, srcPath, dstPath)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldMoveItemAtPathToPath returns true if a handler for FileManagerShouldMoveItemAtPathToPath has been set.
func (d *FileManagerDelegate) HasFileManagerShouldMoveItemAtPathToPath() bool {
	return d._FileManagerShouldMoveItemAtPathToPath != nil
}

// FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool {
	if d._FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL != nil {
		return d._FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(fileManager, error_, srcURL, dstURL)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL returns true if a handler for FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL has been set.
func (d *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL() bool {
	return d._FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL != nil
}

// FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool {
	if d._FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath != nil {
		return d._FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(fileManager, error_, srcPath, dstPath)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath returns true if a handler for FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath has been set.
func (d *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath() bool {
	return d._FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath != nil
}

// FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool {
	if d._FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL != nil {
		return d._FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(fileManager, error_, srcURL, dstURL)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL returns true if a handler for FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL has been set.
func (d *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL() bool {
	return d._FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL != nil
}

// FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool {
	if d._FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath != nil {
		return d._FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(fileManager, error_, srcPath, dstPath)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath returns true if a handler for FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath has been set.
func (d *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath() bool {
	return d._FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath != nil
}

// FileManagerShouldProceedAfterErrorMovingItemAtURLToURL implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldProceedAfterErrorMovingItemAtURLToURL(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool {
	if d._FileManagerShouldProceedAfterErrorMovingItemAtURLToURL != nil {
		return d._FileManagerShouldProceedAfterErrorMovingItemAtURLToURL(fileManager, error_, srcURL, dstURL)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldProceedAfterErrorMovingItemAtURLToURL returns true if a handler for FileManagerShouldProceedAfterErrorMovingItemAtURLToURL has been set.
func (d *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorMovingItemAtURLToURL() bool {
	return d._FileManagerShouldProceedAfterErrorMovingItemAtURLToURL != nil
}

// FileManagerShouldProceedAfterErrorMovingItemAtPathToPath implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldProceedAfterErrorMovingItemAtPathToPath(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool {
	if d._FileManagerShouldProceedAfterErrorMovingItemAtPathToPath != nil {
		return d._FileManagerShouldProceedAfterErrorMovingItemAtPathToPath(fileManager, error_, srcPath, dstPath)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldProceedAfterErrorMovingItemAtPathToPath returns true if a handler for FileManagerShouldProceedAfterErrorMovingItemAtPathToPath has been set.
func (d *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorMovingItemAtPathToPath() bool {
	return d._FileManagerShouldProceedAfterErrorMovingItemAtPathToPath != nil
}

// FileManagerShouldProceedAfterErrorRemovingItemAtURL implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldProceedAfterErrorRemovingItemAtURL(fileManager IFileManager, error_ IError, URL IURL) bool {
	if d._FileManagerShouldProceedAfterErrorRemovingItemAtURL != nil {
		return d._FileManagerShouldProceedAfterErrorRemovingItemAtURL(fileManager, error_, URL)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldProceedAfterErrorRemovingItemAtURL returns true if a handler for FileManagerShouldProceedAfterErrorRemovingItemAtURL has been set.
func (d *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorRemovingItemAtURL() bool {
	return d._FileManagerShouldProceedAfterErrorRemovingItemAtURL != nil
}

// FileManagerShouldProceedAfterErrorRemovingItemAtPath implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldProceedAfterErrorRemovingItemAtPath(fileManager IFileManager, error_ IError, path IString) bool {
	if d._FileManagerShouldProceedAfterErrorRemovingItemAtPath != nil {
		return d._FileManagerShouldProceedAfterErrorRemovingItemAtPath(fileManager, error_, path)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldProceedAfterErrorRemovingItemAtPath returns true if a handler for FileManagerShouldProceedAfterErrorRemovingItemAtPath has been set.
func (d *FileManagerDelegate) HasFileManagerShouldProceedAfterErrorRemovingItemAtPath() bool {
	return d._FileManagerShouldProceedAfterErrorRemovingItemAtPath != nil
}

// FileManagerShouldRemoveItemAtURL implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldRemoveItemAtURL(fileManager IFileManager, URL IURL) bool {
	if d._FileManagerShouldRemoveItemAtURL != nil {
		return d._FileManagerShouldRemoveItemAtURL(fileManager, URL)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldRemoveItemAtURL returns true if a handler for FileManagerShouldRemoveItemAtURL has been set.
func (d *FileManagerDelegate) HasFileManagerShouldRemoveItemAtURL() bool {
	return d._FileManagerShouldRemoveItemAtURL != nil
}

// FileManagerShouldRemoveItemAtPath implements the PFileManagerDelegate interface.
func (d *FileManagerDelegate) FileManagerShouldRemoveItemAtPath(fileManager IFileManager, path IString) bool {
	if d._FileManagerShouldRemoveItemAtPath != nil {
		return d._FileManagerShouldRemoveItemAtPath(fileManager, path)
	}
	var zero bool
	return zero
}

// HasFileManagerShouldRemoveItemAtPath returns true if a handler for FileManagerShouldRemoveItemAtPath has been set.
func (d *FileManagerDelegate) HasFileManagerShouldRemoveItemAtPath() bool {
	return d._FileManagerShouldRemoveItemAtPath != nil
}

// FileManagerDelegateObject wraps an existing Objective-C object that conforms to the PFileManagerDelegate protocol.
// This allows you to safely call protocol methods on any object that implements the protocol,
// with runtime checks for optional methods using RespondsToSelector.
type FileManagerDelegateObject struct {
	objectivec.Object
}

// NewFileManagerDelegateObject creates a new protocol wrapper for an existing Objective-C object.
// The object should implement the NSFileManagerDelegate protocol.
func NewFileManagerDelegateObject(obj objectivec.Object) *FileManagerDelegateObject {
	return &FileManagerDelegateObject{obj}
}

// Make sure FileManagerDelegateObject implements PFileManagerDelegate.
var _ PFileManagerDelegate = (*FileManagerDelegateObject)(nil)

// FileManagerShouldCopyItemAtURLToURL implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldCopyItemAtURLToURL(fileManager IFileManager, srcURL IURL, dstURL IURL) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldCopyItemAtURL:toURL:"), fileManager, srcURL, dstURL)
}

// HasFileManagerShouldCopyItemAtURLToURL returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldCopyItemAtURLToURL() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldCopyItemAtPathToPath implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldCopyItemAtPathToPath(fileManager IFileManager, srcPath IString, dstPath IString) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldCopyItemAtPath:toPath:"), fileManager, srcPath, dstPath)
}

// HasFileManagerShouldCopyItemAtPathToPath returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldCopyItemAtPathToPath() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldLinkItemAtURLToURL implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldLinkItemAtURLToURL(fileManager IFileManager, srcURL IURL, dstURL IURL) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldLinkItemAtURL:toURL:"), fileManager, srcURL, dstURL)
}

// HasFileManagerShouldLinkItemAtURLToURL returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldLinkItemAtURLToURL() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldLinkItemAtPathToPath implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldLinkItemAtPathToPath(fileManager IFileManager, srcPath IString, dstPath IString) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldLinkItemAtPath:toPath:"), fileManager, srcPath, dstPath)
}

// HasFileManagerShouldLinkItemAtPathToPath returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldLinkItemAtPathToPath() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldMoveItemAtURLToURL implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldMoveItemAtURLToURL(fileManager IFileManager, srcURL IURL, dstURL IURL) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldMoveItemAtURL:toURL:"), fileManager, srcURL, dstURL)
}

// HasFileManagerShouldMoveItemAtURLToURL returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldMoveItemAtURLToURL() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldMoveItemAtPathToPath implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldMoveItemAtPathToPath(fileManager IFileManager, srcPath IString, dstPath IString) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldMoveItemAtPath:toPath:"), fileManager, srcPath, dstPath)
}

// HasFileManagerShouldMoveItemAtPathToPath returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldMoveItemAtPathToPath() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:copyingItemAtURL:toURL:"), fileManager, error_, srcURL, dstURL)
}

// HasFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldProceedAfterErrorCopyingItemAtPathToPath(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:copyingItemAtPath:toPath:"), fileManager, error_, srcPath, dstPath)
}

// HasFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorCopyingItemAtPathToPath() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldProceedAfterErrorLinkingItemAtURLToURL(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:linkingItemAtURL:toURL:"), fileManager, error_, srcURL, dstURL)
}

// HasFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorLinkingItemAtURLToURL() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldProceedAfterErrorLinkingItemAtPathToPath(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:linkingItemAtPath:toPath:"), fileManager, error_, srcPath, dstPath)
}

// HasFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorLinkingItemAtPathToPath() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldProceedAfterErrorMovingItemAtURLToURL implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldProceedAfterErrorMovingItemAtURLToURL(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:movingItemAtURL:toURL:"), fileManager, error_, srcURL, dstURL)
}

// HasFileManagerShouldProceedAfterErrorMovingItemAtURLToURL returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorMovingItemAtURLToURL() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldProceedAfterErrorMovingItemAtPathToPath implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldProceedAfterErrorMovingItemAtPathToPath(fileManager IFileManager, error_ IError, srcPath IString, dstPath IString) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:movingItemAtPath:toPath:"), fileManager, error_, srcPath, dstPath)
}

// HasFileManagerShouldProceedAfterErrorMovingItemAtPathToPath returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorMovingItemAtPathToPath() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldProceedAfterErrorRemovingItemAtURL implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldProceedAfterErrorRemovingItemAtURL(fileManager IFileManager, error_ IError, URL IURL) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:removingItemAtURL:"), fileManager, error_, URL)
}

// HasFileManagerShouldProceedAfterErrorRemovingItemAtURL returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorRemovingItemAtURL() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldProceedAfterErrorRemovingItemAtPath implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldProceedAfterErrorRemovingItemAtPath(fileManager IFileManager, error_ IError, path IString) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldProceedAfterError:removingItemAtPath:"), fileManager, error_, path)
}

// HasFileManagerShouldProceedAfterErrorRemovingItemAtPath returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldProceedAfterErrorRemovingItemAtPath() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldRemoveItemAtURL implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldRemoveItemAtURL(fileManager IFileManager, URL IURL) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldRemoveItemAtURL:"), fileManager, URL)
}

// HasFileManagerShouldRemoveItemAtURL returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldRemoveItemAtURL() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}

// FileManagerShouldRemoveItemAtPath implements the PFileManagerDelegate interface.
// This optional method is called directly; checking selector availability is the caller's responsibility.
func (o *FileManagerDelegateObject) FileManagerShouldRemoveItemAtPath(fileManager IFileManager, path IString) bool {
	return objc.Send[bool](o.ID, objc.Sel("fileManager:shouldRemoveItemAtPath:"), fileManager, path)
}

// HasFileManagerShouldRemoveItemAtPath returns true; this is a placeholder for optional method checks.
func (o *FileManagerDelegateObject) HasFileManagerShouldRemoveItemAtPath() bool {
	return true // TODO: Implement proper selector checking when RespondsToSelector is available
}
