// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// FileProviderUserInteractionSuppressingProtocol is the NSFileProviderUserInteractionSuppressing protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to NSFileProviderUserInteractionSuppressing.
var FileProviderUserInteractionSuppressingProtocol *objc.Protocol

func init() {
	FileProviderUserInteractionSuppressingProtocol = objc.GetProtocol("NSFileProviderUserInteractionSuppressing")
}
