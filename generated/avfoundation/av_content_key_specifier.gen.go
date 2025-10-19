// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVContentKeySpecifier] class.
var aVContentKeySpecifierClass = _AVContentKeySpecifierClass{objc.GetClass("AVContentKeySpecifier")}

type _AVContentKeySpecifierClass struct {
	class objc.Class
}

// An object that uniquely identifies a content key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContentKeySpecifier

type AVContentKeySpecifier struct {
	objectivec.Object
}

// AVContentKeySpecifierFrom constructs a [AVContentKeySpecifier] from an unsafe.Pointer.
//
// An object that uniquely identifies a content key.
func AVContentKeySpecifierFrom(ptr unsafe.Pointer) AVContentKeySpecifier {
	return AVContentKeySpecifier{objectivec.Object{objc.ID(ptr)}}
}



