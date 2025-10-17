// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ParagraphStyle] class.
var ParagraphStyleClass objc.Class

func init() {
	ParagraphStyleClass = objc.GetClass("NSParagraphStyle")
}

type ParagraphStyle struct {
	objc.ID
}

func ParagraphStyleFrom(ptr unsafe.Pointer) ParagraphStyle {
	return ParagraphStyle{
		ID: objc.ID(ptr),
	}
}



