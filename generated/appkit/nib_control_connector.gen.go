// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NibControlConnector] class.
var nibControlConnectorClass = _NibControlConnectorClass{objc.GetClass("NSNibControlConnector")}

type _NibControlConnectorClass struct {
	class objc.Class
}

// A control connection between two Interface Builder objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibControlConnector

type NibControlConnector struct {
	NibConnector
}

// NibControlConnectorFrom constructs a [NibControlConnector] from an unsafe.Pointer.
//
// A control connection between two Interface Builder objects.
func NibControlConnectorFrom(ptr unsafe.Pointer) NibControlConnector {
	return NibControlConnector{
		NibConnector: NibConnectorFrom(ptr),
	}
}



