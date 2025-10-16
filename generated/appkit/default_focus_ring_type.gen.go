
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [defaultFocusRingType] class.
var defaultFocusRingTypeClass _defaultFocusRingTypeClass

func init() {
	defaultFocusRingTypeClass = _defaultFocusRingTypeClass{objc.GetClass("defaultFocusRingType")}
}

type _defaultFocusRingTypeClass struct {
	objc.Class
}

// An interface definition for the [defaultFocusRingType] class.
type IdefaultFocusRingType interface {
	ID() objc.ID
}

type defaultFocusRingType struct {
	id objc.ID
}

func defaultFocusRingTypeFrom(ptr unsafe.Pointer) defaultFocusRingType {
	return defaultFocusRingType{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ defaultFocusRingType) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _defaultFocusRingTypeClass) Alloc() defaultFocusRingType {
	rv := objc.Send[defaultFocusRingType](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _defaultFocusRingTypeClass) New() defaultFocusRingType {
	rv := objc.Send[defaultFocusRingType](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewdefaultFocusRingType creates and returns a new initialized instance.
func NewdefaultFocusRingType() defaultFocusRingType {
	return defaultFocusRingTypeClass.New()
}

// Init initializes the instance.
func (d_ defaultFocusRingType) Init() defaultFocusRingType {
	rv := objc.Send[defaultFocusRingType](d_.ID(), selInit)
	return rv
}
