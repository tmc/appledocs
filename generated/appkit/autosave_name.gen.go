
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [autosaveName] class.
var autosaveNameClass _autosaveNameClass

func init() {
	autosaveNameClass = _autosaveNameClass{objc.GetClass("autosaveName")}
}

type _autosaveNameClass struct {
	objc.Class
}

// An interface definition for the [autosaveName] class.
type IautosaveName interface {
	ID() objc.ID
}

type autosaveName struct {
	id objc.ID
}

func autosaveNameFrom(ptr unsafe.Pointer) autosaveName {
	return autosaveName{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ autosaveName) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _autosaveNameClass) Alloc() autosaveName {
	rv := objc.Send[autosaveName](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _autosaveNameClass) New() autosaveName {
	rv := objc.Send[autosaveName](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewautosaveName creates and returns a new initialized instance.
func NewautosaveName() autosaveName {
	return autosaveNameClass.New()
}

// Init initializes the instance.
func (a_ autosaveName) Init() autosaveName {
	rv := objc.Send[autosaveName](a_.ID(), selInit)
	return rv
}
