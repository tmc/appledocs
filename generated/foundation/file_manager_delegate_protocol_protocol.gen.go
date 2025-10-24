// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (

	"github.com/tmc/appledocs/generated/objc"
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
	FileManagerShouldLinkItemAtURLToURL(fileManager IFileManager, srcURL IURL, dstURL IURL) bool
	HasFileManagerShouldLinkItemAtURLToURL() bool
	FileManagerShouldMoveItemAtPathToPath(fileManager IFileManager, srcPath IString, dstPath IString) bool
	HasFileManagerShouldMoveItemAtPathToPath() bool
	FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool
	HasFileManagerShouldProceedAfterErrorCopyingItemAtURLToURL() bool
}

// FileManagerDelegate is a delegate implementation builder for the PFileManagerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type FileManagerDelegate struct {
	_FileManagerShouldLinkItemAtURLToURL func(fileManager IFileManager, srcURL IURL, dstURL IURL) bool
	_FileManagerShouldMoveItemAtPathToPath func(fileManager IFileManager, srcPath IString, dstPath IString) bool
	_FileManagerShouldProceedAfterErrorCopyingItemAtURLToURL func(fileManager IFileManager, error_ IError, srcURL IURL, dstURL IURL) bool
}

// SetFileManagerShouldLinkItemAtURLToURL sets the handler for the FileManagerShouldLinkItemAtURLToURL delegate method.
//
// Asks the delegate if a hard link should be created between the items at the two URLs.
func (d *FileManagerDelegate) SetFileManagerShouldLinkItemAtURLToURL(f func(fileManager IFileManager, srcURL IURL, dstURL IURL) bool) {
	d._FileManagerShouldLinkItemAtURLToURL = f
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
