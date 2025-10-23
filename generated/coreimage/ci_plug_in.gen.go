// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PlugIn] class.
var (
	PlugInClass     _PlugInClass
	PlugInClassOnce sync.Once
)

func getPlugInClass() _PlugInClass {
	PlugInClassOnce.Do(func() {
		PlugInClass = _PlugInClass{objc.GetClass("CIPlugIn")}
	})
	return PlugInClass
}

type _PlugInClass struct {
	class objc.Class
}

// An interface definition for the [PlugIn] class.
type IPlugIn interface {
	objectivec.IObject
}

// The mechanism for loading image units in macOS.
//
// An image unit is an image processing bundle that contains one or more Core Image filters. Th extension indicates one or more filters packaged as an image unit.


// The mechanism for loading image units in macOS.
//
// [Full Topic]
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

// Alloc allocates a new instance without initialization.
func (pc _PlugInClass) Alloc() PlugIn {
	rv := objc.Send[PlugIn](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PlugInClass) New() PlugIn {
	rv := objc.Send[PlugIn](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PlugIn) Init() PlugIn {
	rv := objc.Send[PlugIn](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PlugIn) Autorelease() PlugIn {
	rv := objc.Send[PlugIn](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlugIn creates a new PlugIn instance.
func NewPlugIn() PlugIn {
	return getPlugInClass().New()
}



// Loads filters from an image unit that have the appropriate executable status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/load(_:allowExecutableCode:)
func (pc _PlugInClass) LoadPlugInAllowExecutableCode(url foundation.URL, allowExecutableCode bool) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadPlugIn:allowExecutableCode:"), url, allowExecutableCode)
}


// Scans directories for files that have the extension and then loads the image units.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadAllPlugIns()
func (pc _PlugInClass) LoadAllPlugIns() {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadAllPlugIns"))
}


// Loads a non-executable plug-in specified by its URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadNonExecutablePlugIn(_:)
func (pc _PlugInClass) LoadNonExecutablePlugIn(url foundation.URL) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadNonExecutablePlugIn:"), url)
}


// Scans directories for plugins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadNonExecutablePlugIns()
func (pc _PlugInClass) LoadNonExecutablePlugIns() {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadNonExecutablePlugIns"))
}


// Loads filters from an image unit that have the appropriate executable status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIPlugIn/loadPlugIn:allowNonExecutable:
func (pc _PlugInClass) LoadPlugInAllowNonExecutable(url foundation.URL, allowNonExecutable bool) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadPlugIn:allowNonExecutable:"), url, allowNonExecutable)
}



