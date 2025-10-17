// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TitlebarAccessoryViewController] class.
var TitlebarAccessoryViewControllerClass objc.Class

func init() {
	TitlebarAccessoryViewControllerClass = objc.GetClass("NSTitlebarAccessoryViewController")
}

type TitlebarAccessoryViewController struct {
	objc.ID
}

func TitlebarAccessoryViewControllerFrom(ptr unsafe.Pointer) TitlebarAccessoryViewController {
	return TitlebarAccessoryViewController{
		ID: objc.ID(ptr),
	}
}



