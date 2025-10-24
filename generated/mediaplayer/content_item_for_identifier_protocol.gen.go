// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import "github.com/ebitengine/purego/objc"

// contentItemForIdentifierProtocol is the contentItemForIdentifier: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 14.0)
//   - iOS 10.0+ (Deprecated in 14.0)
//   - iPadOS 10.0+ (Deprecated in 14.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Use this protocol when registering custom classes that conform to contentItemForIdentifier:.
var contentItemForIdentifierProtocol *objc.Protocol

func init() {
	contentItemForIdentifierProtocol = objc.GetProtocol("contentItemForIdentifier:")
}

