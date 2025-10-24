// Code generated from Apple documentation for MediaPlayer. DO NOT EDIT.

package mediaplayer

import "github.com/ebitengine/purego/objc"

// nowPlayingSessionDidChangeActiveProtocol is the nowPlayingSessionDidChangeActive: protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to nowPlayingSessionDidChangeActive:.
var nowPlayingSessionDidChangeActiveProtocol *objc.Protocol

func init() {
	nowPlayingSessionDidChangeActiveProtocol = objc.GetProtocol("nowPlayingSessionDidChangeActive:")
}

