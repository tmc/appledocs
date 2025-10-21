// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import "github.com/ebitengine/purego/objc"

// personPickerProtocol is the personPicker: protocol.
//
// Availability:
//   - macOS 10.9+
//
// Use this protocol when registering custom classes that conform to personPicker:.
var personPickerProtocol *objc.Protocol

func init() {
	personPickerProtocol = objc.GetProtocol("personPicker:")
}
