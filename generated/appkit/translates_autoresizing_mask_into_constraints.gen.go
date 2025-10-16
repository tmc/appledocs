
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [translatesAutoresizingMaskIntoConstraints] class.
var translatesAutoresizingMaskIntoConstraintsClass _translatesAutoresizingMaskIntoConstraintsClass

func init() {
	translatesAutoresizingMaskIntoConstraintsClass = _translatesAutoresizingMaskIntoConstraintsClass{objc.GetClass("translatesAutoresizingMaskIntoConstraints")}
}

type _translatesAutoresizingMaskIntoConstraintsClass struct {
	objc.Class
}

// An interface definition for the [translatesAutoresizingMaskIntoConstraints] class.
type ItranslatesAutoresizingMaskIntoConstraints interface {
	ID() objc.ID
}

type translatesAutoresizingMaskIntoConstraints struct {
	id objc.ID
}

func translatesAutoresizingMaskIntoConstraintsFrom(ptr unsafe.Pointer) translatesAutoresizingMaskIntoConstraints {
	return translatesAutoresizingMaskIntoConstraints{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ translatesAutoresizingMaskIntoConstraints) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _translatesAutoresizingMaskIntoConstraintsClass) Alloc() translatesAutoresizingMaskIntoConstraints {
	rv := objc.Send[translatesAutoresizingMaskIntoConstraints](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _translatesAutoresizingMaskIntoConstraintsClass) New() translatesAutoresizingMaskIntoConstraints {
	rv := objc.Send[translatesAutoresizingMaskIntoConstraints](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewtranslatesAutoresizingMaskIntoConstraints creates and returns a new initialized instance.
func NewtranslatesAutoresizingMaskIntoConstraints() translatesAutoresizingMaskIntoConstraints {
	return translatesAutoresizingMaskIntoConstraintsClass.New()
}

// Init initializes the instance.
func (t_ translatesAutoresizingMaskIntoConstraints) Init() translatesAutoresizingMaskIntoConstraints {
	rv := objc.Send[translatesAutoresizingMaskIntoConstraints](t_.ID(), selInit)
	return rv
}
