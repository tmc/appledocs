// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import "github.com/ebitengine/purego/objc"

// preferredSizeOfLayerProtocol is the preferredSizeOfLayer: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.5+
//
// Use this protocol when registering custom classes that conform to preferredSizeOfLayer:.
var preferredSizeOfLayerProtocol *objc.Protocol

func init() {
	preferredSizeOfLayerProtocol = objc.GetProtocol("preferredSizeOfLayer:")
}
