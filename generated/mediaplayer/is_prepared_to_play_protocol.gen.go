// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import "github.com/ebitengine/purego/objc"

// isPreparedToPlayProtocol is the isPreparedToPlay protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to isPreparedToPlay.
var isPreparedToPlayProtocol *objc.Protocol

func init() {
	isPreparedToPlayProtocol = objc.GetProtocol("isPreparedToPlay")
}
