
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TabViewController] class.
var TabViewControllerClass _TabViewControllerClass

func init() {
	TabViewControllerClass = _TabViewControllerClass{objc.GetClass("NSTabViewController")}
}

type _TabViewControllerClass struct {
	objc.Class
}

// An interface definition for the [TabViewController] class.
type ITabViewController interface {
	ID() objc.ID
}

type TabViewController struct {
	id objc.ID
}

func TabViewControllerFrom(ptr unsafe.Pointer) TabViewController {
	return TabViewController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (t_ TabViewController) ID() objc.ID {
	return t_.id
}

// Alloc allocates a new instance without initialization.
func (tc _TabViewControllerClass) Alloc() TabViewController {
	rv := objc.Send[TabViewController](objc.ID(tc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (tc _TabViewControllerClass) New() TabViewController {
	rv := objc.Send[TabViewController](objc.ID(tc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewTabViewController creates and returns a new initialized instance.
func NewTabViewController() TabViewController {
	return TabViewControllerClass.New()
}

// Init initializes the instance.
func (t_ TabViewController) Init() TabViewController {
	rv := objc.Send[TabViewController](t_.ID(), selInit)
	return rv
}
