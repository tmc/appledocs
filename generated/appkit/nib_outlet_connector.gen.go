// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [NibOutletConnector] class.
var NibOutletConnectorClass objc.Class

func init() {
	NibOutletConnectorClass = objc.GetClass("NSNibOutletConnector")
}

type NibOutletConnector struct {
	objc.ID
}

func NibOutletConnectorFrom(ptr unsafe.Pointer) NibOutletConnector {
	return NibOutletConnector{
		ID: objc.ID(ptr),
	}
}



