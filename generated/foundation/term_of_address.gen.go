// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TermOfAddress] class.
var TermOfAddressClass = _TermOfAddressClass{objc.GetClass("NSTermOfAddress")}

type _TermOfAddressClass struct {
	class objc.Class
}

type TermOfAddress struct {
	objc.ID
}

func TermOfAddressFrom(ptr unsafe.Pointer) TermOfAddress {
	return TermOfAddress{
		ID: objc.ID(ptr),
	}
}




