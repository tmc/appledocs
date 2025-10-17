// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [NibOutletConnector] class.
var nibOutletConnectorClass = _NibOutletConnectorClass{objc.GetClass("NSNibOutletConnector")}

type _NibOutletConnectorClass struct {
	class objc.Class
}

// An interface definition for the [NibOutletConnector] class.
type INibOutletConnector interface {
	INibConnector
}

// An outlet connection between Interface Builder objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibOutletConnector

type NibOutletConnector struct {
	NibConnector
}

// NibOutletConnectorFrom constructs a [NibOutletConnector] from an unsafe.Pointer.
//
// An outlet connection between Interface Builder objects.
func NibOutletConnectorFrom(ptr unsafe.Pointer) NibOutletConnector {
	return NibOutletConnector{
		NibConnector: NibConnectorFrom(ptr),
	}
}



