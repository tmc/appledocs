// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var personNameComponentsClass _PersonNameComponentsClass

func init() {
	personNameComponentsClass = _PersonNameComponentsClass{objc.GetClass("NSPersonNameComponents")}
}

type _PersonNameComponentsClass struct {
	class objc.Class
}

type PersonNameComponents struct {
	objc.ID
}

func PersonNameComponentsFrom(ptr unsafe.Pointer) PersonNameComponents {
	return PersonNameComponents{
		ID: objc.ID(ptr),
	}
}




