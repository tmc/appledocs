// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var TextCheckingResultClass _TextCheckingResultClass

func init() {
	TextCheckingResultClass = _TextCheckingResultClass{objc.GetClass("NSTextCheckingResult")}
}

type _TextCheckingResultClass struct {
	class objc.Class
}

type TextCheckingResult struct {
	objc.ID
}

func TextCheckingResultFrom(ptr unsafe.Pointer) TextCheckingResult {
	return TextCheckingResult{
		ID: objc.ID(ptr),
	}
}




