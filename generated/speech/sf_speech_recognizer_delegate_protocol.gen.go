// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import "github.com/ebitengine/purego/objc"

// SFSpeechRecognizerDelegateProtocol is the SFSpeechRecognizerDelegate protocol.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to SFSpeechRecognizerDelegate.
var SFSpeechRecognizerDelegateProtocol *objc.Protocol

func init() {
	SFSpeechRecognizerDelegateProtocol = objc.GetProtocol("SFSpeechRecognizerDelegate")
}
