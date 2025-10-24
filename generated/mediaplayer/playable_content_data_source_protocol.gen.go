// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import "github.com/ebitengine/purego/objc"

// PlayableContentDataSourceProtocol is the MPPlayableContentDataSource protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 14.0)
//   - iOS 7.1+ (Deprecated in 14.0)
//   - iPadOS 7.1+ (Deprecated in 14.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Use this protocol when registering custom classes that conform to MPPlayableContentDataSource.
var PlayableContentDataSourceProtocol *objc.Protocol

func init() {
	PlayableContentDataSourceProtocol = objc.GetProtocol("MPPlayableContentDataSource")
}

