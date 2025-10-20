// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var enumeratorClass _EnumeratorClass

func init() {
	enumeratorClass = _EnumeratorClass{objc.GetClass("NSEnumerator")}
}

type _EnumeratorClass struct {
	class objc.Class
}

type Enumerator struct {
	objc.ID
}

func EnumeratorFrom(ptr unsafe.Pointer) Enumerator {
	return Enumerator{
		ID: objc.ID(ptr),
	}
}




