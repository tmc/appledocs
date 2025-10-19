// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [layerCount] class.
var layerCountClass = _layerCountClass{objc.GetClass("layerCount")}

type _layerCountClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIAztecCodeDescriptor/layerCount-c.ivar

type layerCount struct {
	objectivec.Object
}

// layerCountFrom constructs a [layerCount] from an unsafe.Pointer.
func layerCountFrom(ptr unsafe.Pointer) layerCount {
	return layerCount{objectivec.Object{objc.ID(ptr)}}
}



