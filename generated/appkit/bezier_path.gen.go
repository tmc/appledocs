// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [BezierPath] class.
var bezierPathClass = _BezierPathClass{objc.GetClass("NSBezierPath")}

type _BezierPathClass struct {
	class objc.Class
}

// An object that can create paths using PostScript-style commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBezierPath

type BezierPath struct {
	objectivec.Object
}

// BezierPathFrom constructs a [BezierPath] from an unsafe.Pointer.
//
// An object that can create paths using PostScript-style commands.
func BezierPathFrom(ptr unsafe.Pointer) BezierPath {
	return BezierPath{objectivec.Object{objc.ID(ptr)}}
}



