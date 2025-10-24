// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [DocumentController] class.
var (
	DocumentControllerClass     _DocumentControllerClass
	DocumentControllerClassOnce sync.Once
)

func getDocumentControllerClass() _DocumentControllerClass {
	DocumentControllerClassOnce.Do(func() {
		DocumentControllerClass = _DocumentControllerClass{objc.GetClass("NSDocumentController")}
	})
	return DocumentControllerClass
}

type _DocumentControllerClass struct {
	class objc.Class
}

// An interface definition for the [DocumentController] class.
type IDocumentController interface {
	objectivec.IObject
	// properties:
	AllowsAutomaticShareMenu() bool
	SetAllowsAutomaticShareMenu(value bool)
	AutosavingDelay() float64
	SetAutosavingDelay(value float64)
	CurrentDirectory() objc.IObject /* cross-framework: NSString */
	SetCurrentDirectory(value objc.IObject /* cross-framework: NSString */)
	CurrentDocument() IDocument
	SetCurrentDocument(value IDocument)
	DefaultType() objc.IObject /* cross-framework: NSString */
	SetDefaultType(value objc.IObject /* cross-framework: NSString */)
	DocumentClassNames() objc.IObject /* cross-framework: NSString */
	SetDocumentClassNames(value objc.IObject /* cross-framework: NSString */)
	Documents() IDocument
	SetDocuments(value IDocument)
	HasEditedDocuments() bool
	SetHasEditedDocuments(value bool)
	MaximumRecentDocumentCount() int
	SetMaximumRecentDocumentCount(value int)
	RecentDocumentURLs() objc.IObject /* cross-framework: URL */
	SetRecentDocumentURLs(value objc.IObject /* cross-framework: URL */)
	// methods:
}

// An object that manages an app’s documents.
//
// As the first-responder target of New and Open menu commands, creates and opens documents and tracks them throughout a session of the app. When opening documents, a document controller runs and manages the modal Open panel. objects also maintain and manage the mappings of document types, extensions, and subclasses as specified in the property loaded from the information property list ( ). You can use various methods to get a list of the current documents, get the current document (which is the document whose window is currently key), get documents based on a given filename or window, and find out about a document’s extension, type, display name, and document class. In some situations, it’s worthwhile to subclass in non- -based apps to get some of its features. For example, the management of the Open Recent menu is useful in apps that don’t use subclasses of .


// An object that manages an app’s documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController
type DocumentController struct {
	objectivec.Object
}

// DocumentControllerFrom constructs a [DocumentController] from an unsafe.Pointer.
//
// An object that manages an app’s documents.
func DocumentControllerFrom(ptr unsafe.Pointer) DocumentController {
	return DocumentController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DocumentControllerClass) Alloc() DocumentController {
	rv := objc.Send[DocumentController](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DocumentControllerClass) New() DocumentController {
	rv := objc.Send[DocumentController](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DocumentController) Init() DocumentController {
	rv := objc.Send[DocumentController](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DocumentController) Autorelease() DocumentController {
	rv := objc.Send[DocumentController](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDocumentController creates a new DocumentController instance.
func NewDocumentController() DocumentController {
	return getDocumentControllerClass().New()
}



// A Boolean value that the system uses to insert a Share menu in the File menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/allowsautomaticsharemenu
func (d_ DocumentController) AllowsAutomaticShareMenu() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("allowsAutomaticShareMenu"))
	return rv
}


// A Boolean value that the system uses to insert a Share menu in the File menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/allowsautomaticsharemenu
func (d_ DocumentController) SetAllowsAutomaticShareMenu(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAllowsAutomaticShareMenu:"), value)
}


// The time interval (in seconds) for periodic autosaving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/autosavingdelay
func (d_ DocumentController) AutosavingDelay() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("autosavingDelay"))
	return rv
}


// The time interval (in seconds) for periodic autosaving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/autosavingdelay
func (d_ DocumentController) SetAutosavingDelay(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavingDelay:"), value)
}


// The directory path to use as the starting point in the Open dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/currentdirectory
func (d_ DocumentController) CurrentDirectory() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("currentDirectory"))
	return rv
}


// The directory path to use as the starting point in the Open dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/currentdirectory
func (d_ DocumentController) SetCurrentDirectory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCurrentDirectory:"), value)
}


// The document object associated with the main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/currentdocument
func (d_ DocumentController) CurrentDocument() IDocument {
	rv := objc.Send[Document](d_.ID, objc.Sel("currentDocument"))
	return rv
}


// The document object associated with the main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/currentdocument
func (d_ DocumentController) SetCurrentDocument(value IDocument) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCurrentDocument:"), value)
}


// Returns the name of the document type that should be used when creating new documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/defaulttype
func (d_ DocumentController) DefaultType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("defaultType"))
	return rv
}


// Returns the name of the document type that should be used when creating new documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/defaulttype
func (d_ DocumentController) SetDefaultType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultType:"), value)
}


// An array of strings representing the custom document classes supported by this app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/documentclassnames
func (d_ DocumentController) DocumentClassNames() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("documentClassNames"))
	return rv
}


// An array of strings representing the custom document classes supported by this app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/documentclassnames
func (d_ DocumentController) SetDocumentClassNames(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDocumentClassNames:"), value)
}


// The document objects managed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/documents
func (d_ DocumentController) Documents() IDocument {
	rv := objc.Send[Document](d_.ID, objc.Sel("documents"))
	return rv
}


// The document objects managed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/documents
func (d_ DocumentController) SetDocuments(value IDocument) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDocuments:"), value)
}


// A Boolean value indicating whether the receiver has any documents with unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/hasediteddocuments
func (d_ DocumentController) HasEditedDocuments() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("hasEditedDocuments"))
	return rv
}


// A Boolean value indicating whether the receiver has any documents with unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/hasediteddocuments
func (d_ DocumentController) SetHasEditedDocuments(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHasEditedDocuments:"), value)
}


// The maximum number of items that may be presented in the standard Open Recent menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/maximumrecentdocumentcount
func (d_ DocumentController) MaximumRecentDocumentCount() int {
	rv := objc.Send[int](d_.ID, objc.Sel("maximumRecentDocumentCount"))
	return rv
}


// The maximum number of items that may be presented in the standard Open Recent menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/maximumrecentdocumentcount
func (d_ DocumentController) SetMaximumRecentDocumentCount(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumRecentDocumentCount:"), value)
}


// The list of recent-document URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/recentdocumenturls
func (d_ DocumentController) RecentDocumentURLs() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("recentDocumentURLs"))
	return rv
}


// The list of recent-document URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/recentdocumenturls
func (d_ DocumentController) SetRecentDocumentURLs(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setRecentDocumentURLs:"), value)
}



