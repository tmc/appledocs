
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [customizationLabel] class.
var customizationLabelClass _customizationLabelClass

func init() {
	customizationLabelClass = _customizationLabelClass{objc.GetClass("customizationLabel")}
}

type _customizationLabelClass struct {
	objc.Class
}

// An interface definition for the [customizationLabel] class.
type IcustomizationLabel interface {
	ID() objc.ID
}

type customizationLabel struct {
	id objc.ID
}

func customizationLabelFrom(ptr unsafe.Pointer) customizationLabel {
	return customizationLabel{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ customizationLabel) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _customizationLabelClass) Alloc() customizationLabel {
	rv := objc.Send[customizationLabel](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _customizationLabelClass) New() customizationLabel {
	rv := objc.Send[customizationLabel](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcustomizationLabel creates and returns a new initialized instance.
func NewcustomizationLabel() customizationLabel {
	return customizationLabelClass.New()
}

// Init initializes the instance.
func (c_ customizationLabel) Init() customizationLabel {
	rv := objc.Send[customizationLabel](c_.ID(), selInit)
	return rv
}
