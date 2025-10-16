
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Document] class.
var DocumentClass _DocumentClass

func init() {
	DocumentClass = _DocumentClass{objc.GetClass("NSDocument")}
}

type _DocumentClass struct {
	objc.Class
}

// An interface definition for the [Document] class.
type IDocument interface {
	ID() objc.ID
	EncodeRestorableStateWithCoder(coder unsafe.Pointer)
	ValidateUserInterfaceItem(item unsafe.Pointer) bool
}

type Document struct {
	id objc.ID
}

func DocumentFrom(ptr unsafe.Pointer) Document {
	return Document{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (d_ Document) ID() objc.ID {
	return d_.id
}

// Alloc allocates a new instance without initialization.
func (dc _DocumentClass) Alloc() Document {
	rv := objc.Send[Document](objc.ID(dc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (dc _DocumentClass) New() Document {
	rv := objc.Send[Document](objc.ID(dc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewDocument creates and returns a new initialized instance.
func NewDocument() Document {
	return DocumentClass.New()
}

// Init initializes the instance.
func (d_ Document) Init() Document {
	rv := objc.Send[Document](d_.ID(), selInit)
	return rv
}
// Returns the classes that support secure coding. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSDocument/allowedClasses(forRestorableStateKeyPath:)
func (dc _DocumentClass) AllowedClassesForRestorableStateKeyPath(keyPath unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.Class), objc.RegisterName("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}
// Saves the interface-related state of the document. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSDocument/encodeRestorableState(with:)
func (d_ Document) EncodeRestorableStateWithCoder(coder unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID(), objc.RegisterName("encodeRestorableStateWithCoder:"), coder)
}
// Validates the specified user interface item that the receiver manages. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSDocument/validateUserInterfaceItem(_:)
func (d_ Document) ValidateUserInterfaceItem(item unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID(), objc.RegisterName("validateUserInterfaceItem:"), item)
	return rv
}
