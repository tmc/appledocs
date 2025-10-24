// Code generated from Apple documentation for AddressBook. DO NOT EDIT.

package addressbook

import "github.com/ebitengine/purego/objc"

// personPickerDidCloseProtocol is the personPickerDidClose: protocol.
//
// Availability:
//   - macOS 10.9+
//
// Use this protocol when registering custom classes that conform to personPickerDidClose:.
var personPickerDidCloseProtocol *objc.Protocol

func init() {
	personPickerDidCloseProtocol = objc.GetProtocol("personPickerDidClose:")
}

