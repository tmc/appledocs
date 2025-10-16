
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [prepareLayout] class.
var prepareLayoutClass _prepareLayoutClass

func init() {
	prepareLayoutClass = _prepareLayoutClass{objc.GetClass("prepareLayout")}
}

type _prepareLayoutClass struct {
	objc.Class
}

// An interface definition for the [prepareLayout] class.
type IprepareLayout interface {
	ID() objc.ID
}

type prepareLayout struct {
	id objc.ID
}

func prepareLayoutFrom(ptr unsafe.Pointer) prepareLayout {
	return prepareLayout{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ prepareLayout) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _prepareLayoutClass) Alloc() prepareLayout {
	rv := objc.Send[prepareLayout](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _prepareLayoutClass) New() prepareLayout {
	rv := objc.Send[prepareLayout](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewprepareLayout creates and returns a new initialized instance.
func NewprepareLayout() prepareLayout {
	return prepareLayoutClass.New()
}

// Init initializes the instance.
func (p_ prepareLayout) Init() prepareLayout {
	rv := objc.Send[prepareLayout](p_.ID(), selInit)
	return rv
}
