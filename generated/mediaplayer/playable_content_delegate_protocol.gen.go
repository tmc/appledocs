// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import "github.com/ebitengine/purego/objc"

// PlayableContentDelegateProtocol is the MPPlayableContentDelegate protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 14.0)
//   - iOS 7.1+ (Deprecated in 14.0)
//   - iPadOS 7.1+ (Deprecated in 14.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Use this protocol when registering custom classes that conform to MPPlayableContentDelegate.
var PlayableContentDelegateProtocol *objc.Protocol

func init() {
	PlayableContentDelegateProtocol = objc.GetProtocol("MPPlayableContentDelegate")
}

