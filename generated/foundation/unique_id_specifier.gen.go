// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UniqueIDSpecifier] class.
var uniqueIDSpecifierClass = _UniqueIDSpecifierClass{objc.GetClass("NSUniqueIDSpecifier")}

type _UniqueIDSpecifierClass struct {
	class objc.Class
}

// A specifier for an object in a collection (or container) by unique ID. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUniqueIDSpecifier

type UniqueIDSpecifier struct {
	ScriptObjectSpecifier
}

// UniqueIDSpecifierFrom constructs a [UniqueIDSpecifier] from an unsafe.Pointer.
//
// A specifier for an object in a collection (or container) by unique ID.
func UniqueIDSpecifierFrom(ptr unsafe.Pointer) UniqueIDSpecifier {
	return UniqueIDSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}



