// Code generated from Apple documentation for CoreWLAN. DO NOT EDIT.

package corewlan

import "github.com/ebitengine/purego/objc"

// clientConnectionInterruptedProtocol is the clientConnectionInterrupted protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.6+
//
// Use this protocol when registering custom classes that conform to clientConnectionInterrupted.
var clientConnectionInterruptedProtocol *objc.Protocol

func init() {
	clientConnectionInterruptedProtocol = objc.GetProtocol("clientConnectionInterrupted")
}
