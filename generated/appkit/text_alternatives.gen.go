
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TextAlternatives] class.
var TextAlternativesClass _TextAlternativesClass

func init() {
	TextAlternativesClass = _TextAlternativesClass{objc.GetClass("NSTextAlternatives")}
}

type _TextAlternativesClass struct {
	objc.Class
}

// An interface definition for the [TextAlternatives] class.
type ITextAlternatives interface {
	ID() objc.ID
}

type TextAlternatives struct {
	id objc.ID
}

func TextAlternativesFrom(ptr unsafe.Pointer) TextAlternatives {
	return TextAlternatives{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TextAlternatives) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TextAlternativesClass) Alloc() TextAlternatives {
	rv := objc.Send[TextAlternatives](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TextAlternativesClass) New() TextAlternatives {
	rv := objc.Send[TextAlternatives](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTextAlternatives creates and returns a new initialized instance.
func NewTextAlternatives() TextAlternatives {
	return TextAlternativesClass.New()
}

// Init initializes the instance.
func (t_ TextAlternatives) Init() TextAlternatives {
	rv := objc.Send[TextAlternatives](t_.ID(), selInit)
	return rv
}
