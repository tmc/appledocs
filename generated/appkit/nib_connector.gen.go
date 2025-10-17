// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NibConnector] class.
var nibConnectorClass = _NibConnectorClass{objc.GetClass("NSNibConnector")}

type _NibConnectorClass struct {
	class objc.Class
}

// An interface definition for the [NibConnector] class.
type INibConnector interface {
	objectivec.IObject
}

// A connection between two nibs. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNibConnector

type NibConnector struct {
	objectivec.Object
}

// NibConnectorFrom constructs a [NibConnector] from an unsafe.Pointer.
//
// A connection between two nibs.
func NibConnectorFrom(ptr unsafe.Pointer) NibConnector {
	return NibConnector{objectivec.Object{objc.ID(ptr)}}
}



