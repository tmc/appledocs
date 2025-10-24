// Code generated from Apple documentation for HealthKit. DO NOT EDIT.

package healthkit

import "github.com/ebitengine/purego/objc"

// workoutSessionProtocol is the workoutSession: protocol.
//
// Availability:
//   - Mac Catalyst 17.0+
//   - iOS 17.0+
//   - iPadOS 17.0+
//   - macOS 13.0+
//   - visionOS 1.0+
//   - watchOS 9.0+
//
// Use this protocol when registering custom classes that conform to workoutSession:.
var workoutSessionProtocol *objc.Protocol

func init() {
	workoutSessionProtocol = objc.GetProtocol("workoutSession:")
}

