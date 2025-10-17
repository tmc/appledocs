// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SplitViewItemAccessoryViewController] class.
var SplitViewItemAccessoryViewControllerClass objc.Class

func init() {
	SplitViewItemAccessoryViewControllerClass = objc.GetClass("NSSplitViewItemAccessoryViewController")
}

type SplitViewItemAccessoryViewController struct {
	objc.ID
}

func SplitViewItemAccessoryViewControllerFrom(ptr unsafe.Pointer) SplitViewItemAccessoryViewController {
	return SplitViewItemAccessoryViewController{
		ID: objc.ID(ptr),
	}
}




