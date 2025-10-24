// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// downloadProtocol is the download: protocol.
//
// Availability:
//   - Mac Catalyst 14.5+
//   - iOS 14.5+
//   - iPadOS 14.5+
//   - macOS 11.3+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to download:.
var downloadProtocol *objc.Protocol

func init() {
	downloadProtocol = objc.GetProtocol("download:")
}
