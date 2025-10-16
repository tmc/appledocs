
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [allowsFullHeightLayout] class.
var allowsFullHeightLayoutClass _allowsFullHeightLayoutClass

func init() {
	allowsFullHeightLayoutClass = _allowsFullHeightLayoutClass{objc.GetClass("allowsFullHeightLayout")}
}

type _allowsFullHeightLayoutClass struct {
	objc.Class
}

// An interface definition for the [allowsFullHeightLayout] class.
type IallowsFullHeightLayout interface {
	ID() objc.ID
}

type allowsFullHeightLayout struct {
	id objc.ID
}

func allowsFullHeightLayoutFrom(ptr unsafe.Pointer) allowsFullHeightLayout {
	return allowsFullHeightLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ allowsFullHeightLayout) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _allowsFullHeightLayoutClass) Alloc() allowsFullHeightLayout {
	rv := objc.Send[allowsFullHeightLayout](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _allowsFullHeightLayoutClass) New() allowsFullHeightLayout {
	rv := objc.Send[allowsFullHeightLayout](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewallowsFullHeightLayout creates and returns a new initialized instance.
func NewallowsFullHeightLayout() allowsFullHeightLayout {
	return allowsFullHeightLayoutClass.New()
}

// Init initializes the instance.
func (a_ allowsFullHeightLayout) Init() allowsFullHeightLayout {
	rv := objc.Send[allowsFullHeightLayout](a_.ID(), selInit)
	return rv
}
