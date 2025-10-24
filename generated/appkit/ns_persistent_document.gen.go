// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coredata"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PersistentDocument] class.
var (
	PersistentDocumentClass     _PersistentDocumentClass
	PersistentDocumentClassOnce sync.Once
)

func getPersistentDocumentClass() _PersistentDocumentClass {
	PersistentDocumentClassOnce.Do(func() {
		PersistentDocumentClass = _PersistentDocumentClass{objc.GetClass("NSPersistentDocument")}
	})
	return PersistentDocumentClass
}

type _PersistentDocumentClass struct {
	class objc.Class
}

// An interface definition for the [PersistentDocument] class.
type IPersistentDocument interface {
	IDocument
	// properties:
	HasUndoManager() bool
	SetHasUndoManager(value bool)
	IsDocumentEdited() bool
	SetIsDocumentEdited(value bool)
	UndoManager() objc.IObject /* cross-framework: UndoManager */
	SetUndoManager(value objc.IObject /* cross-framework: UndoManager */)
	ManagedObjectContext() objc.IObject /* cross-framework: ManagedObjectContext */
	SetManagedObjectContext(value objc.IObject /* cross-framework: ManagedObjectContext */)
	ManagedObjectModel() objc.IObject /* cross-framework: ManagedObjectModel */
	SetManagedObjectModel(value objc.IObject /* cross-framework: ManagedObjectModel */)
	// methods:
}

// A document object that can integrate with Core Data.
//
// The class is a subclass of that is designed to easily integrate into the Core Data framework. It provides methods to access a document-wide object, and provides default implementations of methods to read and write files using the persistence framework. In a persistent document, the undo manager functionality is taken over by managed object context. Standard document behavior is implemented as follows: Opening a document invokes with the new URL, and adds a store of the default type (XML). Objects are loaded from the persistent store on demand through the document’s context. Saving a new document adds a store of the default type with the chosen URL and invokes save: on the context. For an existing document, a save just invokes on the context. Save As for a new document simply invokes save. For an opened document, it migrates the persistent store to the new URL and invokes on the context. Revert resets the document’s managed object context. Objects are subsequently loaded from the persistent store on demand, as with opening a new document. By default an instance creates its own ready-to-use persistence stack including managed object context, persistent object store coordinator and persistent store. There is a one-to-one mapping between the document and the backing object store. You can customize the architecture of the persistence stack by overriding the property and method. You might wish to do this, for example, to specify a particular managed object model.


// A document object that can integrate with Core Data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPersistentDocument
type PersistentDocument struct {
	Document
}

// PersistentDocumentFrom constructs a [PersistentDocument] from an unsafe.Pointer.
//
// A document object that can integrate with Core Data.
func PersistentDocumentFrom(ptr unsafe.Pointer) PersistentDocument {
	return PersistentDocument{
		Document: DocumentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentDocumentClass) Alloc() PersistentDocument {
	rv := objc.Send[PersistentDocument](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentDocumentClass) New() PersistentDocument {
	rv := objc.Send[PersistentDocument](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentDocument) Init() PersistentDocument {
	rv := objc.Send[PersistentDocument](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentDocument) Autorelease() PersistentDocument {
	rv := objc.Send[PersistentDocument](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentDocument creates a new PersistentDocument instance.
func NewPersistentDocument() PersistentDocument {
	return getPersistentDocumentClass().New()
}



// A Boolean value that indicates whether the document owns an undo manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/hasundomanager
func (p_ PersistentDocument) HasUndoManager() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("hasUndoManager"))
	return rv
}


// A Boolean value that indicates whether the document owns an undo manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/hasundomanager
func (p_ PersistentDocument) SetHasUndoManager(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHasUndoManager:"), value)
}


// A Boolean value that indicates whether the document has unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isdocumentedited
func (p_ PersistentDocument) IsDocumentEdited() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDocumentEdited"))
	return rv
}


// A Boolean value that indicates whether the document has unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isdocumentedited
func (p_ PersistentDocument) SetIsDocumentEdited(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsDocumentEdited:"), value)
}


// The object that the document uses to support undo/redo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/undomanager
func (p_ PersistentDocument) UndoManager() objc.IObject /* cross-framework: UndoManager */ {
	rv := objc.Send[foundation.UndoManager](p_.ID, objc.Sel("undoManager"))
	return rv
}


// The object that the document uses to support undo/redo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/undomanager
func (p_ PersistentDocument) SetUndoManager(value objc.IObject /* cross-framework: UndoManager */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUndoManager:"), value)
}


// The managed object context for the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspersistentdocument/managedobjectcontext
func (p_ PersistentDocument) ManagedObjectContext() objc.IObject /* cross-framework: ManagedObjectContext */ {
	rv := objc.Send[coredata.ManagedObjectContext](p_.ID, objc.Sel("managedObjectContext"))
	return rv
}


// The managed object context for the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspersistentdocument/managedobjectcontext
func (p_ PersistentDocument) SetManagedObjectContext(value objc.IObject /* cross-framework: ManagedObjectContext */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setManagedObjectContext:"), value)
}


// The managed object model of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspersistentdocument/managedobjectmodel
func (p_ PersistentDocument) ManagedObjectModel() objc.IObject /* cross-framework: ManagedObjectModel */ {
	rv := objc.Send[coredata.ManagedObjectModel](p_.ID, objc.Sel("managedObjectModel"))
	return rv
}


// The managed object model of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nspersistentdocument/managedobjectmodel
func (p_ PersistentDocument) SetManagedObjectModel(value objc.IObject /* cross-framework: ManagedObjectModel */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setManagedObjectModel:"), value)
}



