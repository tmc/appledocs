// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import "github.com/ebitengine/purego/objc"

// OSLogEntryWithPayloadProtocol is the OSLogEntryWithPayload protocol.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 10.15+
//   - tvOS 15.0+
//   - visionOS 1.0+
//   - watchOS 8.0+
//
// Use this protocol when registering custom classes that conform to OSLogEntryWithPayload.
var OSLogEntryWithPayloadProtocol *objc.Protocol

func init() {
	OSLogEntryWithPayloadProtocol = objc.GetProtocol("OSLogEntryWithPayload")
}
