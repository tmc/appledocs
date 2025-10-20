// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var keyedUnarchiverClass _KeyedUnarchiverClass

func init() {
	keyedUnarchiverClass = _KeyedUnarchiverClass{objc.GetClass("NSKeyedUnarchiver")}
}

type _KeyedUnarchiverClass struct {
	class objc.Class
}

type KeyedUnarchiver struct {
	objc.ID
}

func KeyedUnarchiverFrom(ptr unsafe.Pointer) KeyedUnarchiver {
	return KeyedUnarchiver{
		ID: objc.ID(ptr),
	}
}




