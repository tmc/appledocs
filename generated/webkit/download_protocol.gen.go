// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// downloadProtocol is the download: protocol.
//
// Availability:
//   - Mac Catalyst 18.2+
//   - iOS 18.2+
//   - iPadOS 18.2+
//   - macOS 11.3+
//   - visionOS 2.2+
//
// Use this protocol when registering custom classes that conform to download:.
var downloadProtocol *objc.Protocol

func init() {
	downloadProtocol = objc.GetProtocol("download:")
}
