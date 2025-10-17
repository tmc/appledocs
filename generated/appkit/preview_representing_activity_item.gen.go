
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PreviewRepresentingActivityItem] class.
var PreviewRepresentingActivityItemClass _PreviewRepresentingActivityItemClass

func init() {
	PreviewRepresentingActivityItemClass = _PreviewRepresentingActivityItemClass{objc.GetClass("NSPreviewRepresentingActivityItem")}
}

type _PreviewRepresentingActivityItemClass struct {
	objc.Class
}

// An interface definition for the [PreviewRepresentingActivityItem] class.
type IPreviewRepresentingActivityItem interface {
	ID() objc.ID
}

type PreviewRepresentingActivityItem struct {
	id objc.ID
}

func PreviewRepresentingActivityItemFrom(ptr unsafe.Pointer) PreviewRepresentingActivityItem {
	return PreviewRepresentingActivityItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PreviewRepresentingActivityItem) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PreviewRepresentingActivityItemClass) Alloc() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PreviewRepresentingActivityItemClass) New() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPreviewRepresentingActivityItem creates and returns a new initialized instance.
func NewPreviewRepresentingActivityItem() PreviewRepresentingActivityItem {
	return PreviewRepresentingActivityItemClass.New()
}

// Init initializes the instance.
func (p_ PreviewRepresentingActivityItem) Init() PreviewRepresentingActivityItem {
	rv := objc.Send[PreviewRepresentingActivityItem](p_.ID(), selInit)
	return rv
}
