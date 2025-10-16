
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [viewDidChangeEffectiveAppearance] class.
var viewDidChangeEffectiveAppearanceClass _viewDidChangeEffectiveAppearanceClass

func init() {
	viewDidChangeEffectiveAppearanceClass = _viewDidChangeEffectiveAppearanceClass{objc.GetClass("viewDidChangeEffectiveAppearance")}
}

type _viewDidChangeEffectiveAppearanceClass struct {
	objc.Class
}

// An interface definition for the [viewDidChangeEffectiveAppearance] class.
type IviewDidChangeEffectiveAppearance interface {
	ID() objc.ID
}

type viewDidChangeEffectiveAppearance struct {
	id objc.ID
}

func viewDidChangeEffectiveAppearanceFrom(ptr unsafe.Pointer) viewDidChangeEffectiveAppearance {
	return viewDidChangeEffectiveAppearance{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (v_ viewDidChangeEffectiveAppearance) ID() objc.ID {
	return v_.id
}

// Alloc allocates a new instance without initialization.
func (vc _viewDidChangeEffectiveAppearanceClass) Alloc() viewDidChangeEffectiveAppearance {
	rv := objc.Send[viewDidChangeEffectiveAppearance](objc.ID(vc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (vc _viewDidChangeEffectiveAppearanceClass) New() viewDidChangeEffectiveAppearance {
	rv := objc.Send[viewDidChangeEffectiveAppearance](objc.ID(vc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewviewDidChangeEffectiveAppearance creates and returns a new initialized instance.
func NewviewDidChangeEffectiveAppearance() viewDidChangeEffectiveAppearance {
	return viewDidChangeEffectiveAppearanceClass.New()
}

// Init initializes the instance.
func (v_ viewDidChangeEffectiveAppearance) Init() viewDidChangeEffectiveAppearance {
	rv := objc.Send[viewDidChangeEffectiveAppearance](v_.ID(), selInit)
	return rv
}
