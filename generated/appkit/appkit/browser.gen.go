// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Browser] class.
var BrowserClass objc.Class

func init() {
	BrowserClass = objc.GetClass("NSBrowser")
}

type Browser struct {
	objc.ID
}

func BrowserFrom(ptr unsafe.Pointer) Browser {
	return Browser{
		ID: objc.ID(ptr),
	}
}



