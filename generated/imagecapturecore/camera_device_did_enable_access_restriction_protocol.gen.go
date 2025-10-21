// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import "github.com/ebitengine/purego/objc"

// cameraDeviceDidEnableAccessRestrictionProtocol is the cameraDeviceDidEnableAccessRestriction: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to cameraDeviceDidEnableAccessRestriction:.
var cameraDeviceDidEnableAccessRestrictionProtocol *objc.Protocol

func init() {
	cameraDeviceDidEnableAccessRestrictionProtocol = objc.GetProtocol("cameraDeviceDidEnableAccessRestriction:")
}
