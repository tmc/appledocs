// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import "github.com/ebitengine/purego/objc"

// extensionWillTerminateProtocol is the extensionWillTerminate protocol.
//
// Availability:
//   - Mac Catalyst 16.1+ (Deprecated in 16.4)
//   - iOS 16.1+ (Deprecated in 16.4)
//   - iPadOS 16.1+ (Deprecated in 16.4)
//   - macOS 13.0+ (Deprecated in 13.3)
//
// Use this protocol when registering custom classes that conform to extensionWillTerminate.
var extensionWillTerminateProtocol *objc.Protocol

func init() {
	extensionWillTerminateProtocol = objc.GetProtocol("extensionWillTerminate")
}
