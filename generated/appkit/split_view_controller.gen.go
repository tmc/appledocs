
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SplitViewController] class.
var SplitViewControllerClass _SplitViewControllerClass

func init() {
	SplitViewControllerClass = _SplitViewControllerClass{objc.GetClass("NSSplitViewController")}
}

type _SplitViewControllerClass struct {
	objc.Class
}

// An interface definition for the [SplitViewController] class.
type ISplitViewController interface {
	ID() objc.ID
}

type SplitViewController struct {
	id objc.ID
}

func SplitViewControllerFrom(ptr unsafe.Pointer) SplitViewController {
	return SplitViewController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SplitViewController) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SplitViewControllerClass) Alloc() SplitViewController {
	rv := objc.Send[SplitViewController](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SplitViewControllerClass) New() SplitViewController {
	rv := objc.Send[SplitViewController](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSplitViewController creates and returns a new initialized instance.
func NewSplitViewController() SplitViewController {
	return SplitViewControllerClass.New()
}

// Init initializes the instance.
func (s_ SplitViewController) Init() SplitViewController {
	rv := objc.Send[SplitViewController](s_.ID(), selInit)
	return rv
}
