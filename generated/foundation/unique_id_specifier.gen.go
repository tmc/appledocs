// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UniqueIDSpecifier] class.
var UniqueIDSpecifierClass = _UniqueIDSpecifierClass{objc.GetClass("NSUniqueIDSpecifier")}

type _UniqueIDSpecifierClass struct {
	class objc.Class
}

type UniqueIDSpecifier struct {
	objc.ID
}

func UniqueIDSpecifierFrom(ptr unsafe.Pointer) UniqueIDSpecifier {
	return UniqueIDSpecifier{
		ID: objc.ID(ptr),
	}
}




