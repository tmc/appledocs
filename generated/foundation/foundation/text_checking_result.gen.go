// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextCheckingResult] class.
var TextCheckingResultClass objc.Class

func init() {
	TextCheckingResultClass = objc.GetClass("NSTextCheckingResult")
}

type TextCheckingResult struct {
	objc.ID
}

func TextCheckingResultFrom(ptr unsafe.Pointer) TextCheckingResult {
	return TextCheckingResult{
		ID: objc.ID(ptr),
	}
}




