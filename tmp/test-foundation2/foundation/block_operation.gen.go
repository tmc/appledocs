// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var blockOperationClass _BlockOperationClass

func init() {
	blockOperationClass = _BlockOperationClass{objc.GetClass("NSBlockOperation")}
}

type _BlockOperationClass struct {
	class objc.Class
}

type BlockOperation struct {
	objc.ID
}

func BlockOperationFrom(ptr unsafe.Pointer) BlockOperation {
	return BlockOperation{
		ID: objc.ID(ptr),
	}
}




