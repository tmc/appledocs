// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NibControlConnector] class.
var NibControlConnectorClass objc.Class

func init() {
	NibControlConnectorClass = objc.GetClass("NSNibControlConnector")
}

type NibControlConnector struct {
	objc.ID
}

func NibControlConnectorFrom(ptr unsafe.Pointer) NibControlConnector {
	return NibControlConnector{
		ID: objc.ID(ptr),
	}
}




