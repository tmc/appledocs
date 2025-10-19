// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Document] class.
var (
	documentClass     _DocumentClass
	documentClassOnce sync.Once
)

func getDocumentClass() _DocumentClass {
	documentClassOnce.Do(func() {
		documentClass = _DocumentClass{objc.GetClass("NSDocument")}
	})
	return documentClass
}

type _DocumentClass struct {
	class objc.Class
}

// An interface definition for the [Document] class.
type IDocument interface {
	objectivec.IObject
	EncodeRestorableStateWithCoder(coder unsafe.Pointer)
	ValidateUserInterfaceItem(item unsafe.Pointer) bool
}

// An abstract class that defines the interface for macOS documents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument
type Document struct {
	objectivec.Object
}

// DocumentFrom constructs a [Document] from an unsafe.Pointer.
//
// An abstract class that defines the interface for macOS documents.
func DocumentFrom(ptr unsafe.Pointer) Document {
	return Document{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DocumentClass) Alloc() Document {
	rv := objc.Send[Document](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DocumentClass) New() Document {
	rv := objc.Send[Document](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Document) Init() Document {
	rv := objc.Send[Document](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Document) Autorelease() Document {
	rv := objc.Send[Document](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDocument creates a new Document instance.
func NewDocument() Document {
	return getDocumentClass().New()
}


// Returns the classes that support secure coding.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/allowedClasses(forRestorableStateKeyPath:)
func (dc _DocumentClass) AllowedClassesForRestorableStateKeyPath(keyPath string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(dc.class), objc.Sel("allowedClassesForRestorableStateKeyPath:"), objc.String(keyPath))
	return rv
}

// Saves the interface-related state of the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/encodeRestorableState(with:)
func (d_ Document) EncodeRestorableStateWithCoder(coder unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("encodeRestorableStateWithCoder:"), coder)
}

// Validates the specified user interface item that the receiver manages.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/validateUserInterfaceItem(_:)
func (d_ Document) ValidateUserInterfaceItem(item unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}



