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
	DocumentClass     _DocumentClass
	DocumentClassOnce sync.Once
)

func getDocumentClass() _DocumentClass {
	DocumentClassOnce.Do(func() {
		DocumentClass = _DocumentClass{objc.GetClass("NSDocument")}
	})
	return DocumentClass
}

type _DocumentClass struct {
	class objc.Class
}

// An interface definition for the [Document] class.
type IDocument interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other AppKit classes.


// A parent class referenced by other AppKit classes. [Full Topic]
type Document struct {
	objectivec.Object
}

// DocumentFrom constructs a [Document] from an unsafe.Pointer.
//
// A parent class referenced by other AppKit classes.
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




