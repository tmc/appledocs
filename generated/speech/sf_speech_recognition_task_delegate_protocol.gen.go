// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import "github.com/ebitengine/purego/objc"

// SFSpeechRecognitionTaskDelegateProtocol is the SFSpeechRecognitionTaskDelegate protocol.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to SFSpeechRecognitionTaskDelegate.
var SFSpeechRecognitionTaskDelegateProtocol *objc.Protocol

func init() {
	SFSpeechRecognitionTaskDelegateProtocol = objc.GetProtocol("SFSpeechRecognitionTaskDelegate")
}
