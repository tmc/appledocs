// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import "github.com/ebitengine/purego/objc"

// deviceDidChangeNameProtocol is the deviceDidChangeName: protocol.
//
// Availability:
//   - macOS 10.4+
//
// Use this protocol when registering custom classes that conform to deviceDidChangeName:.
var deviceDidChangeNameProtocol *objc.Protocol

func init() {
	deviceDidChangeNameProtocol = objc.GetProtocol("deviceDidChangeName:")
}

