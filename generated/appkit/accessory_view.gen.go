
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [accessoryView] class.
var accessoryViewClass _accessoryViewClass

func init() {
	accessoryViewClass = _accessoryViewClass{objc.GetClass("accessoryView")}
}

type _accessoryViewClass struct {
	objc.Class
}

// An interface definition for the [accessoryView] class.
type IaccessoryView interface {
	ID() objc.ID
}

type accessoryView struct {
	id objc.ID
}

func accessoryViewFrom(ptr unsafe.Pointer) accessoryView {
	return accessoryView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (a_ accessoryView) ID() objc.ID {
	return a_.id
}

// Alloc allocates a new instance without initialization.
func (ac _accessoryViewClass) Alloc() accessoryView {
	rv := objc.Send[accessoryView](objc.ID(ac.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ac _accessoryViewClass) New() accessoryView {
	rv := objc.Send[accessoryView](objc.ID(ac.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewaccessoryView creates and returns a new initialized instance.
func NewaccessoryView() accessoryView {
	return accessoryViewClass.New()
}

// Init initializes the instance.
func (a_ accessoryView) Init() accessoryView {
	rv := objc.Send[accessoryView](a_.ID(), selInit)
	return rv
}
