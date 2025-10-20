// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import "github.com/ebitengine/purego/objc"

// scannerDeviceProtocol is the scannerDevice: protocol.
//
// Availability:
//   - macOS 10.4+
//
// Use this protocol when registering custom classes that conform to scannerDevice:.
var scannerDeviceProtocol *objc.Protocol

func init() {
	scannerDeviceProtocol = objc.GetProtocol("scannerDevice:")
}

