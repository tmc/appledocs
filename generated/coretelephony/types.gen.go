// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony
import (
	"unsafe"
)


// C struct types
// CTError - A type representing a Core Telephony error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTError
type CTError struct {
	Domain unsafe.Pointer // A numeric indication of the error domain.
	Error unsafe.Pointer // A code indicating the specific error.
}



