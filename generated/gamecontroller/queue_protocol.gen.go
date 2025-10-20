// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import "github.com/ebitengine/purego/objc"

// queueProtocol is the queue protocol.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 14.0+
//   - tvOS 17.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to queue.
var queueProtocol *objc.Protocol

func init() {
	queueProtocol = objc.GetProtocol("queue")
}


