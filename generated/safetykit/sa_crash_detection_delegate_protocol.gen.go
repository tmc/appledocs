// Code generated from Apple documentation for SafetyKit. DO NOT EDIT.

package safetykit

import "github.com/ebitengine/purego/objc"

// SACrashDetectionDelegateProtocol is the SACrashDetectionDelegate protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - watchOS 10.1+
//
// Use this protocol when registering custom classes that conform to SACrashDetectionDelegate.
var SACrashDetectionDelegateProtocol *objc.Protocol

func init() {
	SACrashDetectionDelegateProtocol = objc.GetProtocol("SACrashDetectionDelegate")
}
