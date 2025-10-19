// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlugIn] class.
var plugInClass = _PlugInClass{objc.GetClass("CIPlugIn")}

type _PlugInClass struct {
	class objc.Class
}

// The mechanism for loading image units in macOS. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn

type PlugIn struct {
	objectivec.Object
}

// PlugInFrom constructs a [PlugIn] from an unsafe.Pointer.
//
// The mechanism for loading image units in macOS.
func PlugInFrom(ptr unsafe.Pointer) PlugIn {
	return PlugIn{objectivec.Object{objc.ID(ptr)}}
}

// Loads filters from an image unit that have the appropriate executable status. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/load(_:allowExecutableCode:)
func (pc _PlugInClass) LoadPlugInAllowExecutableCode(url unsafe.Pointer, allowExecutableCode bool) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadPlugIn:allowExecutableCode:"), url, allowExecutableCode)
}
// Scans directories for files that have the extension and then loads the image units. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadAllPlugIns()
func (pc _PlugInClass) LoadAllPlugIns() {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadAllPlugIns"))
}
// Loads a non-executable plug-in specified by its URL. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadNonExecutablePlugIn(_:)
func (pc _PlugInClass) LoadNonExecutablePlugIn(url unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadNonExecutablePlugIn:"), url)
}
// Scans directories for plugins. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadNonExecutablePlugIns()
func (pc _PlugInClass) LoadNonExecutablePlugIns() {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadNonExecutablePlugIns"))
}
// Loads filters from an image unit that have the appropriate executable status. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadPlugIn:allowNonExecutable:
func (pc _PlugInClass) LoadPlugInAllowNonExecutable(url unsafe.Pointer, allowNonExecutable bool) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadPlugIn:allowNonExecutable:"), url, allowNonExecutable)
}


