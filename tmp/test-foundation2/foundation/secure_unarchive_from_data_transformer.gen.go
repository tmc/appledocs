// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var secureUnarchiveFromDataTransformerClass _SecureUnarchiveFromDataTransformerClass

func init() {
	secureUnarchiveFromDataTransformerClass = _SecureUnarchiveFromDataTransformerClass{objc.GetClass("NSSecureUnarchiveFromDataTransformer")}
}

type _SecureUnarchiveFromDataTransformerClass struct {
	class objc.Class
}

type SecureUnarchiveFromDataTransformer struct {
	objc.ID
}

func SecureUnarchiveFromDataTransformerFrom(ptr unsafe.Pointer) SecureUnarchiveFromDataTransformer {
	return SecureUnarchiveFromDataTransformer{
		ID: objc.ID(ptr),
	}
}




