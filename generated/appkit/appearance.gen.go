
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Appearance] class.
var AppearanceClass _AppearanceClass

func init() {
	AppearanceClass = _AppearanceClass{objc.GetClass("NSAppearance")}
}

type _AppearanceClass struct {
	objc.Class
}

// An interface definition for the [Appearance] class.
type IAppearance interface {
	ID() objc.ID
}

type Appearance struct {
	id objc.ID
}

func AppearanceFrom(ptr unsafe.Pointer) Appearance {
	return Appearance{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ Appearance) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _AppearanceClass) Alloc() Appearance {
	rv := objc.Send[Appearance](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _AppearanceClass) New() Appearance {
	rv := objc.Send[Appearance](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewAppearance creates and returns a new initialized instance.
func NewAppearance() Appearance {
	return AppearanceClass.New()
}

// Init initializes the instance.
func (a_ Appearance) Init() Appearance {
	rv := objc.Send[Appearance](a_.ID(), selInit)
	return rv
}
