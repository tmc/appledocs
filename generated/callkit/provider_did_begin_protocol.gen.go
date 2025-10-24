// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import "github.com/ebitengine/purego/objc"

// providerDidBeginProtocol is the providerDidBegin: protocol.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - visionOS 1.0+
//   - watchOS 9.0+
//
// Use this protocol when registering custom classes that conform to providerDidBegin:.
var providerDidBeginProtocol *objc.Protocol

func init() {
	providerDidBeginProtocol = objc.GetProtocol("providerDidBegin:")
}

