
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [reservedThicknessForAccessoryView] class.
var reservedThicknessForAccessoryViewClass _reservedThicknessForAccessoryViewClass

func init() {
	reservedThicknessForAccessoryViewClass = _reservedThicknessForAccessoryViewClass{objc.GetClass("reservedThicknessForAccessoryView")}
}

type _reservedThicknessForAccessoryViewClass struct {
	objc.Class
}

// An interface definition for the [reservedThicknessForAccessoryView] class.
type IreservedThicknessForAccessoryView interface {
	ID() objc.ID
}

type reservedThicknessForAccessoryView struct {
	id objc.ID
}

func reservedThicknessForAccessoryViewFrom(ptr unsafe.Pointer) reservedThicknessForAccessoryView {
	return reservedThicknessForAccessoryView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (r_ reservedThicknessForAccessoryView) ID() objc.ID {
	return r_.id
}

// Alloc allocates a new instance without initialization.
func (rc _reservedThicknessForAccessoryViewClass) Alloc() reservedThicknessForAccessoryView {
	rv := objc.Send[reservedThicknessForAccessoryView](objc.ID(rc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (rc _reservedThicknessForAccessoryViewClass) New() reservedThicknessForAccessoryView {
	rv := objc.Send[reservedThicknessForAccessoryView](objc.ID(rc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewreservedThicknessForAccessoryView creates and returns a new initialized instance.
func NewreservedThicknessForAccessoryView() reservedThicknessForAccessoryView {
	return reservedThicknessForAccessoryViewClass.New()
}

// Init initializes the instance.
func (r_ reservedThicknessForAccessoryView) Init() reservedThicknessForAccessoryView {
	rv := objc.Send[reservedThicknessForAccessoryView](r_.ID(), selInit)
	return rv
}
