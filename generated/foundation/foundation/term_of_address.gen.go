// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TermOfAddress] class.
var TermOfAddressClass objc.Class

func init() {
	TermOfAddressClass = objc.GetClass("NSTermOfAddress")
}

type TermOfAddress struct {
	objc.ID
}

func TermOfAddressFrom(ptr unsafe.Pointer) TermOfAddress {
	return TermOfAddress{
		ID: objc.ID(ptr),
	}
}




