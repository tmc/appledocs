
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [needsPanelToBecomeKey] class.
var needsPanelToBecomeKeyClass _needsPanelToBecomeKeyClass

func init() {
	needsPanelToBecomeKeyClass = _needsPanelToBecomeKeyClass{objc.GetClass("needsPanelToBecomeKey")}
}

type _needsPanelToBecomeKeyClass struct {
	objc.Class
}

// An interface definition for the [needsPanelToBecomeKey] class.
type IneedsPanelToBecomeKey interface {
	ID() objc.ID
}

type needsPanelToBecomeKey struct {
	id objc.ID
}

func needsPanelToBecomeKeyFrom(ptr unsafe.Pointer) needsPanelToBecomeKey {
	return needsPanelToBecomeKey{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ needsPanelToBecomeKey) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _needsPanelToBecomeKeyClass) Alloc() needsPanelToBecomeKey {
	rv := objc.Send[needsPanelToBecomeKey](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _needsPanelToBecomeKeyClass) New() needsPanelToBecomeKey {
	rv := objc.Send[needsPanelToBecomeKey](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewneedsPanelToBecomeKey creates and returns a new initialized instance.
func NewneedsPanelToBecomeKey() needsPanelToBecomeKey {
	return needsPanelToBecomeKeyClass.New()
}

// Init initializes the instance.
func (n_ needsPanelToBecomeKey) Init() needsPanelToBecomeKey {
	rv := objc.Send[needsPanelToBecomeKey](n_.ID(), selInit)
	return rv
}
