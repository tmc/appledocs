// Code generated from Apple documentation for BackgroundAssets. DO NOT EDIT.

package backgroundassets

import "github.com/ebitengine/purego/objc"

// downloadProtocol is the download: protocol.
//
// Availability:
//   - Mac Catalyst 16.1+
//   - iOS 16.1+
//   - iPadOS 16.1+
//   - macOS 13.0+
//   - tvOS 18.4+
//   - visionOS 2.4+
//
// Use this protocol when registering custom classes that conform to download:.
var downloadProtocol *objc.Protocol

func init() {
	downloadProtocol = objc.GetProtocol("download:")
}

