// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [WritingToolsCoordinatorContext] class.
var WritingToolsCoordinatorContextClass objc.Class

func init() {
	WritingToolsCoordinatorContextClass = objc.GetClass("NSWritingToolsCoordinatorContext")
}

type WritingToolsCoordinatorContext struct {
	objc.ID
}

func WritingToolsCoordinatorContextFrom(ptr unsafe.Pointer) WritingToolsCoordinatorContext {
	return WritingToolsCoordinatorContext{
		ID: objc.ID(ptr),
	}
}



