// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PrintInfo] class.
var PrintInfoClass objc.Class

func init() {
	PrintInfoClass = objc.GetClass("NSPrintInfo")
}

type PrintInfo struct {
	objc.ID
}

func PrintInfoFrom(ptr unsafe.Pointer) PrintInfo {
	return PrintInfo{
		ID: objc.ID(ptr),
	}
}




