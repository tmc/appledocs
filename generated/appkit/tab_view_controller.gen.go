// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TabViewController] class.
var TabViewControllerClass objc.Class

func init() {
	TabViewControllerClass = objc.GetClass("NSTabViewController")
}

type TabViewController struct {
	objc.ID
}

func TabViewControllerFrom(ptr unsafe.Pointer) TabViewController {
	return TabViewController{
		ID: objc.ID(ptr),
	}
}



