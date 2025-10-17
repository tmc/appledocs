// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [KeyedUnarchiver] class.
var KeyedUnarchiverClass objc.Class

func init() {
	KeyedUnarchiverClass = objc.GetClass("NSKeyedUnarchiver")
}

type KeyedUnarchiver struct {
	objc.ID
}

func KeyedUnarchiverFrom(ptr unsafe.Pointer) KeyedUnarchiver {
	return KeyedUnarchiver{
		ID: objc.ID(ptr),
	}
}




