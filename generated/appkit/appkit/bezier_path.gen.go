// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [BezierPath] class.
var BezierPathClass objc.Class

func init() {
	BezierPathClass = objc.GetClass("NSBezierPath")
}

type BezierPath struct {
	objc.ID
}

func BezierPathFrom(ptr unsafe.Pointer) BezierPath {
	return BezierPath{
		ID: objc.ID(ptr),
	}
}




