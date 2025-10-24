// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import "github.com/ebitengine/purego/objc"

// MediaTimingProtocol is the CAMediaTiming protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.5+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to CAMediaTiming.
var MediaTimingProtocol *objc.Protocol

func init() {
	MediaTimingProtocol = objc.GetProtocol("CAMediaTiming")
}

