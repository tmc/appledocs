// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SecureUnarchiveFromDataTransformer] class.
var SecureUnarchiveFromDataTransformerClass objc.Class

func init() {
	SecureUnarchiveFromDataTransformerClass = objc.GetClass("NSSecureUnarchiveFromDataTransformer")
}

type SecureUnarchiveFromDataTransformer struct {
	objc.ID
}

func SecureUnarchiveFromDataTransformerFrom(ptr unsafe.Pointer) SecureUnarchiveFromDataTransformer {
	return SecureUnarchiveFromDataTransformer{
		ID: objc.ID(ptr),
	}
}



