// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SplitViewController] class.
var SplitViewControllerClass objc.Class

func init() {
	SplitViewControllerClass = objc.GetClass("NSSplitViewController")
}

type SplitViewController struct {
	objc.ID
}

func SplitViewControllerFrom(ptr unsafe.Pointer) SplitViewController {
	return SplitViewController{
		ID: objc.ID(ptr),
	}
}



