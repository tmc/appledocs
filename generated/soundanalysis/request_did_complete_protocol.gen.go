// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import "github.com/ebitengine/purego/objc"

// requestDidCompleteProtocol is the requestDidComplete: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// Use this protocol when registering custom classes that conform to requestDidComplete:.
var requestDidCompleteProtocol *objc.Protocol

func init() {
	requestDidCompleteProtocol = objc.GetProtocol("requestDidComplete:")
}

