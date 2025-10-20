// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var termOfAddressClass _TermOfAddressClass

func init() {
	termOfAddressClass = _TermOfAddressClass{objc.GetClass("NSTermOfAddress")}
}

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




