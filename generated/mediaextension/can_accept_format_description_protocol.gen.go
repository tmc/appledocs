// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import "github.com/ebitengine/purego/objc"

// canAcceptFormatDescriptionProtocol is the canAcceptFormatDescription: protocol.
//
// Availability:
//   - macOS 14.0+
//
// Use this protocol when registering custom classes that conform to canAcceptFormatDescription:.
var canAcceptFormatDescriptionProtocol *objc.Protocol

func init() {
	canAcceptFormatDescriptionProtocol = objc.GetProtocol("canAcceptFormatDescription:")
}

