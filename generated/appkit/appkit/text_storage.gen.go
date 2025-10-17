// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextStorage] class.
var TextStorageClass objc.Class

func init() {
	TextStorageClass = objc.GetClass("NSTextStorage")
}

type TextStorage struct {
	objc.ID
}

func TextStorageFrom(ptr unsafe.Pointer) TextStorage {
	return TextStorage{
		ID: objc.ID(ptr),
	}
}



