// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [MutableParagraphStyle] class.
var MutableParagraphStyleClass objc.Class

func init() {
	MutableParagraphStyleClass = objc.GetClass("NSMutableParagraphStyle")
}

type MutableParagraphStyle struct {
	objc.ID
}

func MutableParagraphStyleFrom(ptr unsafe.Pointer) MutableParagraphStyle {
	return MutableParagraphStyle{
		ID: objc.ID(ptr),
	}
}




