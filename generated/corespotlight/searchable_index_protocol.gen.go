// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import "github.com/ebitengine/purego/objc"

// searchableIndexProtocol is the searchableIndex: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to searchableIndex:.
var searchableIndexProtocol *objc.Protocol

func init() {
	searchableIndexProtocol = objc.GetProtocol("searchableIndex:")
}
