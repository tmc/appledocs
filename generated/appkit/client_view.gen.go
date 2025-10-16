
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [clientView] class.
var clientViewClass _clientViewClass

func init() {
	clientViewClass = _clientViewClass{objc.GetClass("clientView")}
}

type _clientViewClass struct {
	objc.Class
}

// An interface definition for the [clientView] class.
type IclientView interface {
	ID() objc.ID
}

type clientView struct {
	id objc.ID
}

func clientViewFrom(ptr unsafe.Pointer) clientView {
	return clientView{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ clientView) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _clientViewClass) Alloc() clientView {
	rv := objc.Send[clientView](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _clientViewClass) New() clientView {
	rv := objc.Send[clientView](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewclientView creates and returns a new initialized instance.
func NewclientView() clientView {
	return clientViewClass.New()
}

// Init initializes the instance.
func (c_ clientView) Init() clientView {
	rv := objc.Send[clientView](c_.ID(), selInit)
	return rv
}
