
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SplitViewItemAccessoryViewController] class.
var SplitViewItemAccessoryViewControllerClass _SplitViewItemAccessoryViewControllerClass

func init() {
	SplitViewItemAccessoryViewControllerClass = _SplitViewItemAccessoryViewControllerClass{objc.GetClass("NSSplitViewItemAccessoryViewController")}
}

type _SplitViewItemAccessoryViewControllerClass struct {
	objc.Class
}

// An interface definition for the [SplitViewItemAccessoryViewController] class.
type ISplitViewItemAccessoryViewController interface {
	ID() objc.ID
}

type SplitViewItemAccessoryViewController struct {
	id objc.ID
}

func SplitViewItemAccessoryViewControllerFrom(ptr unsafe.Pointer) SplitViewItemAccessoryViewController {
	return SplitViewItemAccessoryViewController{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (s_ SplitViewItemAccessoryViewController) ID() objc.ID {
	return s_.id
}

// Alloc allocates a new instance without initialization.
func (sc _SplitViewItemAccessoryViewControllerClass) Alloc() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](objc.ID(sc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (sc _SplitViewItemAccessoryViewControllerClass) New() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](objc.ID(sc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewSplitViewItemAccessoryViewController creates and returns a new initialized instance.
func NewSplitViewItemAccessoryViewController() SplitViewItemAccessoryViewController {
	return SplitViewItemAccessoryViewControllerClass.New()
}

// Init initializes the instance.
func (s_ SplitViewItemAccessoryViewController) Init() SplitViewItemAccessoryViewController {
	rv := objc.Send[SplitViewItemAccessoryViewController](s_.ID(), selInit)
	return rv
}
