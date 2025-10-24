// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import "github.com/ebitengine/purego/objc"

// SessionDelegateProtocol is the GKSessionDelegate protocol.
//
// Availability:
//   - macOS 10.8+ (Deprecated in 10.10)
//   - visionOS 1.0+ (Deprecated in 1.0)
//   - watchOS 2.0+ (Deprecated in 2.0)
//
// Use this protocol when registering custom classes that conform to GKSessionDelegate.
var SessionDelegateProtocol *objc.Protocol

func init() {
	SessionDelegateProtocol = objc.GetProtocol("GKSessionDelegate")
}

