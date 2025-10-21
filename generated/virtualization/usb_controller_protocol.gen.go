// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import "github.com/ebitengine/purego/objc"

// usbControllerProtocol is the usbController protocol.
//
// Availability:
//   - macOS 15.0+
//
// Use this protocol when registering custom classes that conform to usbController.
var usbControllerProtocol *objc.Protocol

func init() {
	usbControllerProtocol = objc.GetProtocol("usbController")
}
