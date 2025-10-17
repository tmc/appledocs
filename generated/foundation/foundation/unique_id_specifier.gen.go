// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [UniqueIDSpecifier] class.
var UniqueIDSpecifierClass objc.Class

func init() {
	UniqueIDSpecifierClass = objc.GetClass("NSUniqueIDSpecifier")
}

type UniqueIDSpecifier struct {
	objc.ID
}

func UniqueIDSpecifierFrom(ptr unsafe.Pointer) UniqueIDSpecifier {
	return UniqueIDSpecifier{
		ID: objc.ID(ptr),
	}
}




