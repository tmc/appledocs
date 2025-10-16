
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [style] class.
var styleClass _styleClass

func init() {
	styleClass = _styleClass{objc.GetClass("style")}
}

type _styleClass struct {
	objc.Class
}

// An interface definition for the [style] class.
type Istyle interface {
	ID() objc.ID
}

type style struct {
	id objc.ID
}

func styleFrom(ptr unsafe.Pointer) style {
	return style{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ style) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _styleClass) Alloc() style {
	rv := objc.Send[style](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _styleClass) New() style {
	rv := objc.Send[style](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newstyle creates and returns a new initialized instance.
func Newstyle() style {
	return styleClass.New()
}

// Init initializes the instance.
func (s_ style) Init() style {
	rv := objc.Send[style](s_.ID(), selInit)
	return rv
}
