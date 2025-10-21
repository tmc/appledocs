// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// presentedTimeProtocol is the presentedTime protocol.
//
// Availability:
//   - Mac Catalyst 13.4+
//   - iOS 10.3+
//   - iPadOS 10.3+
//   - macOS 10.15.4+
//   - tvOS 10.2+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to presentedTime.
var presentedTimeProtocol *objc.Protocol

func init() {
	presentedTimeProtocol = objc.GetProtocol("presentedTime")
}
