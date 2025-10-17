// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WritingToolsCoordinatorAnimationParameters] class.
var WritingToolsCoordinatorAnimationParametersClass objc.Class

func init() {
	WritingToolsCoordinatorAnimationParametersClass = objc.GetClass("NSWritingToolsCoordinatorAnimationParameters")
}

type WritingToolsCoordinatorAnimationParameters struct {
	objc.ID
}

func WritingToolsCoordinatorAnimationParametersFrom(ptr unsafe.Pointer) WritingToolsCoordinatorAnimationParameters {
	return WritingToolsCoordinatorAnimationParameters{
		ID: objc.ID(ptr),
	}
}



