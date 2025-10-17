// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PrintInfo] class.
var printInfoClass = _PrintInfoClass{objc.GetClass("NSPrintInfo")}

type _PrintInfoClass struct {
	class objc.Class
}

// An object that stores information that’s used to generate printed output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo

type PrintInfo struct {
	objectivec.Object
}

// PrintInfoFrom constructs a [PrintInfo] from an unsafe.Pointer.
//
// An object that stores information that’s used to generate printed output.
func PrintInfoFrom(ptr unsafe.Pointer) PrintInfo {
	return PrintInfo{objectivec.Object{objc.ID(ptr)}}
}



