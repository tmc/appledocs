// Code generated from Apple documentation for ContactsUI. DO NOT EDIT.

package contactsui

import "github.com/ebitengine/purego/objc"

// contactPickerProtocol is the contactPicker: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to contactPicker:.
var contactPickerProtocol *objc.Protocol

func init() {
	contactPickerProtocol = objc.GetProtocol("contactPicker:")
}
