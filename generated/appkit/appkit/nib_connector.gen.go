// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NibConnector] class.
var NibConnectorClass objc.Class

func init() {
	NibConnectorClass = objc.GetClass("NSNibConnector")
}

type NibConnector struct {
	objc.ID
}

func NibConnectorFrom(ptr unsafe.Pointer) NibConnector {
	return NibConnector{
		ID: objc.ID(ptr),
	}
}




