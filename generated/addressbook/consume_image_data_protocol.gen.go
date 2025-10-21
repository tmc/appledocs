// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import "github.com/ebitengine/purego/objc"

// consumeImageDataProtocol is the consumeImageData: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to consumeImageData:.
var consumeImageDataProtocol *objc.Protocol

func init() {
	consumeImageDataProtocol = objc.GetProtocol("consumeImageData:")
}
