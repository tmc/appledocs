// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// receivedDataProtocol is the receivedData: protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to receivedData:.
var receivedDataProtocol *objc.Protocol

func init() {
	receivedDataProtocol = objc.GetProtocol("receivedData:")
}

