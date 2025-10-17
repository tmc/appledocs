// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PageLayout] class.
var PageLayoutClass objc.Class

func init() {
	PageLayoutClass = objc.GetClass("NSPageLayout")
}

type PageLayout struct {
	objc.ID
}

func PageLayoutFrom(ptr unsafe.Pointer) PageLayout {
	return PageLayout{
		ID: objc.ID(ptr),
	}
}



