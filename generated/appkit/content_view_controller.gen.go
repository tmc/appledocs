
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [contentViewController] class.
var contentViewControllerClass _contentViewControllerClass

func init() {
	contentViewControllerClass = _contentViewControllerClass{objc.GetClass("contentViewController")}
}

type _contentViewControllerClass struct {
	objc.Class
}

// An interface definition for the [contentViewController] class.
type IcontentViewController interface {
	ID() objc.ID
}

type contentViewController struct {
	id objc.ID
}

func contentViewControllerFrom(ptr unsafe.Pointer) contentViewController {
	return contentViewController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (c_ contentViewController) ID() objc.ID {
	return c_.id
}

// Alloc allocates a new instance without initialization.
func (cc _contentViewControllerClass) Alloc() contentViewController {
	rv := objc.Send[contentViewController](objc.ID(cc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (cc _contentViewControllerClass) New() contentViewController {
	rv := objc.Send[contentViewController](objc.ID(cc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewcontentViewController creates and returns a new initialized instance.
func NewcontentViewController() contentViewController {
	return contentViewControllerClass.New()
}

// Init initializes the instance.
func (c_ contentViewController) Init() contentViewController {
	rv := objc.Send[contentViewController](c_.ID(), selInit)
	return rv
}
