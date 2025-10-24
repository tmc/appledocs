// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import "github.com/ebitengine/purego/objc"

// speechRecognizerProtocol is the speechRecognizer: protocol.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to speechRecognizer:.
var speechRecognizerProtocol *objc.Protocol

func init() {
	speechRecognizerProtocol = objc.GetProtocol("speechRecognizer:")
}
