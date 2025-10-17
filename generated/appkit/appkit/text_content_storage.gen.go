// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextContentStorage] class.
var TextContentStorageClass objc.Class

func init() {
	TextContentStorageClass = objc.GetClass("NSTextContentStorage")
}

type TextContentStorage struct {
	objc.ID
}

func TextContentStorageFrom(ptr unsafe.Pointer) TextContentStorage {
	return TextContentStorage{
		ID: objc.ID(ptr),
	}
}




