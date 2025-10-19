// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PageLayout] class.
var (
	pageLayoutClass     _PageLayoutClass
	pageLayoutClassOnce sync.Once
)

func getPageLayoutClass() _PageLayoutClass {
	pageLayoutClassOnce.Do(func() {
		pageLayoutClass = _PageLayoutClass{objc.GetClass("NSPageLayout")}
	})
	return pageLayoutClass
}

type _PageLayoutClass struct {
	class objc.Class
}

// An interface definition for the [PageLayout] class.
type IPageLayout interface {
	objectivec.IObject
}

// A panel that queries the user for information such as paper type and orientation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout
type PageLayout struct {
	objectivec.Object
}

// PageLayoutFrom constructs a [PageLayout] from an unsafe.Pointer.
//
// A panel that queries the user for information such as paper type and orientation.
func PageLayoutFrom(ptr unsafe.Pointer) PageLayout {
	return PageLayout{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PageLayoutClass) Alloc() PageLayout {
	rv := objc.Send[PageLayout](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PageLayoutClass) New() PageLayout {
	rv := objc.Send[PageLayout](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PageLayout) Init() PageLayout {
	rv := objc.Send[PageLayout](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PageLayout) Autorelease() PageLayout {
	rv := objc.Send[PageLayout](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPageLayout creates a new PageLayout instance.
func NewPageLayout() PageLayout {
	return getPageLayoutClass().New()
}




