
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PasteboardItem] class.
var PasteboardItemClass _PasteboardItemClass

func init() {
	PasteboardItemClass = _PasteboardItemClass{objc.GetClass("NSPasteboardItem")}
}

type _PasteboardItemClass struct {
	objc.Class
}

// An interface definition for the [PasteboardItem] class.
type IPasteboardItem interface {
	ID() objc.ID
}

type PasteboardItem struct {
	id objc.ID
}

func PasteboardItemFrom(ptr unsafe.Pointer) PasteboardItem {
	return PasteboardItem{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PasteboardItem) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PasteboardItemClass) Alloc() PasteboardItem {
	rv := objc.Send[PasteboardItem](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PasteboardItemClass) New() PasteboardItem {
	rv := objc.Send[PasteboardItem](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPasteboardItem creates and returns a new initialized instance.
func NewPasteboardItem() PasteboardItem {
	return PasteboardItemClass.New()
}

// Init initializes the instance.
func (p_ PasteboardItem) Init() PasteboardItem {
	rv := objc.Send[PasteboardItem](p_.ID(), selInit)
	return rv
}
// A model object you use for conveying data during a collaboration. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPasteboardItem/collaborationMetadata
func (p_ PasteboardItem) CollaborationMetadata() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID(), objc.RegisterName("collaborationMetadata"))
	return rv
}
// SetCollaborationMetadata sets the value of the collaborationMetadata property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPasteboardItem/collaborationMetadata
func (p_ PasteboardItem) SetCollaborationMetadata(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setCollaborationMetadata:"), value)
}
