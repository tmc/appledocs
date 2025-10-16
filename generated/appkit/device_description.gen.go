
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [deviceDescription] class.
var deviceDescriptionClass _deviceDescriptionClass

func init() {
	deviceDescriptionClass = _deviceDescriptionClass{objc.GetClass("deviceDescription")}
}

type _deviceDescriptionClass struct {
	objc.Class
}

// An interface definition for the [deviceDescription] class.
type IdeviceDescription interface {
	ID() objc.ID
}

type deviceDescription struct {
	id objc.ID
}

func deviceDescriptionFrom(ptr unsafe.Pointer) deviceDescription {
	return deviceDescription{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ deviceDescription) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _deviceDescriptionClass) Alloc() deviceDescription {
	rv := objc.Send[deviceDescription](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _deviceDescriptionClass) New() deviceDescription {
	rv := objc.Send[deviceDescription](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdeviceDescription creates and returns a new initialized instance.
func NewdeviceDescription() deviceDescription {
	return deviceDescriptionClass.New()
}

// Init initializes the instance.
func (d_ deviceDescription) Init() deviceDescription {
	rv := objc.Send[deviceDescription](d_.ID(), selInit)
	return rv
}
