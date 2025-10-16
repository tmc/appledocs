
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [title] class.
var titleClass _titleClass

func init() {
	titleClass = _titleClass{objc.GetClass("title")}
}

type _titleClass struct {
	objc.Class
}

// An interface definition for the [title] class.
type Ititle interface {
	ID() objc.ID
}

type title struct {
	id objc.ID
}

func titleFrom(ptr unsafe.Pointer) title {
	return title{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ title) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _titleClass) Alloc() title {
	rv := objc.Send[title](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _titleClass) New() title {
	rv := objc.Send[title](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// Newtitle creates and returns a new initialized instance.
func Newtitle() title {
	return titleClass.New()
}

// Init initializes the instance.
func (t_ title) Init() title {
	rv := objc.Send[title](t_.ID(), selInit)
	return rv
}
