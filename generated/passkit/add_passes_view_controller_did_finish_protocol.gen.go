// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import "github.com/ebitengine/purego/objc"

// addPassesViewControllerDidFinishProtocol is the addPassesViewControllerDidFinish: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to addPassesViewControllerDidFinish:.
var addPassesViewControllerDidFinishProtocol *objc.Protocol

func init() {
	addPassesViewControllerDidFinishProtocol = objc.GetProtocol("addPassesViewControllerDidFinish:")
}
