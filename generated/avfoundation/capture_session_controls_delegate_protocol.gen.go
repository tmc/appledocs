// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import "github.com/ebitengine/purego/objc"

// CaptureSessionControlsDelegateProtocol is the AVCaptureSessionControlsDelegate protocol.
//
// Availability:
//   - Mac Catalyst 18.0+
//   - iOS 18.0+
//   - iPadOS 18.0+
//   - macOS 15.0+
//   - tvOS 18.0+
//
// Use this protocol when registering custom classes that conform to AVCaptureSessionControlsDelegate.
var CaptureSessionControlsDelegateProtocol *objc.Protocol

func init() {
	CaptureSessionControlsDelegateProtocol = objc.GetProtocol("AVCaptureSessionControlsDelegate")
}
