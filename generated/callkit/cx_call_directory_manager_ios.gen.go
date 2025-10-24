//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CXCallDirectoryManager


// Asynchronously returns the enabled status of the extension with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/getEnabledStatusForExtension(withIdentifier:completionHandler:)
func (c_ CXCallDirectoryManager) GetEnabledStatusForExtensionWithIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("getEnabledStatusForExtensionWithIdentifier:completionHandler:"), identifier, completion)
}

// Opens the iOS Settings app and shows the Call Blocking & Identification settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/openSettings(completionHandler:)
func (c_ CXCallDirectoryManager) OpenSettingsWithCompletionHandler(completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("openSettingsWithCompletionHandler:"), completion)
}

// Asynchronously reloads the extension with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallDirectoryManager/reloadExtension(withIdentifier:completionHandler:)
func (c_ CXCallDirectoryManager) ReloadExtensionWithIdentifierCompletionHandler(identifier objc.IObject /* cross-framework: NSString */, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reloadExtensionWithIdentifier:completionHandler:"), identifier, completion)
}

// iOS-only properties





