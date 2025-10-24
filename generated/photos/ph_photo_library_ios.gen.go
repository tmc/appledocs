//go:build darwin && ios

// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for PHPhotoLibrary

// Prompts the user to update their limited library selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/presentLimitedLibraryPicker(from:)
func (p_ PHPhotoLibrary) PresentLimitedLibraryPickerFromViewController(controller objc.IObject /* cross-framework: ViewController */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("presentLimitedLibraryPickerFromViewController:"), controller)
}

// Prompts the user to update their limited library selection with a callback providing newly selected identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/presentLimitedLibraryPicker(from:completionHandler:)
func (p_ PHPhotoLibrary) PresentLimitedLibraryPickerFromViewControllerCompletionHandler(controller objc.IObject /* cross-framework: ViewController */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("presentLimitedLibraryPickerFromViewController:completionHandler:"), controller, completionHandler)
}

// Enables or disables the background asset resource upload job processing. This must be called before creating , by the extension’s host application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/setUploadJobExtensionEnabled(_:)
func (p_ PHPhotoLibrary) SetUploadJobExtensionEnabledError(enable bool, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("setUploadJobExtensionEnabled:error:"), enable, error_)
	return rv
}

// iOS-only properties

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Photos/PHPhotoLibrary/uploadJobExtensionEnabled
func (p_ PHPhotoLibrary) UploadJobExtensionEnabled() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("uploadJobExtensionEnabled"))
	return rv
}
