// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var OrthographyClass _OrthographyClass

func init() {
	OrthographyClass = _OrthographyClass{objc.GetClass("NSOrthography")}
}

type _OrthographyClass struct {
	class objc.Class
}

type Orthography struct {
	objc.ID
}

func OrthographyFrom(ptr unsafe.Pointer) Orthography {
	return Orthography{
		ID: objc.ID(ptr),
	}
}




