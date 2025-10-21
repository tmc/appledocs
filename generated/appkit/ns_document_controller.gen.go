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
	BeginOpenPanelWithCompletionHandler(completionHandler unsafe.Pointer)
	MakeDocumentWithContentsOfURLOfTypeError(url foundation.URL, typeName string, outError unsafe.Pointer) unsafe.Pointer
	OpenDocumentWithContentsOfURLDisplayCompletionHandler(url foundation.URL, displayDocument bool, completionHandler unsafe.Pointer)
}

// An object that manages an app’s documents.
//
// As the first-responder target of New and Open menu commands, creates and opens documents and tracks them throughout a session of the app. When opening documents, a document controller runs and manages the modal Open panel. objects also maintain and manage the mappings of document types, extensions, and subclasses as specified in the property loaded from the information property list ( ). You can use various methods to get a list of the current documents, get the current document (which is the document whose window is currently key), get documents based on a given filename or window, and find out about a document’s extension, type, display name, and document class. In some situations, it’s worthwhile to subclass in non- -based apps to get some of its features. For example, the management of the Open Recent menu is useful in apps that don’t use subclasses of .
//
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


// Presents an Open dialog and delivers the results to a completion handler as an array of URLs for the chosen files, or nil.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/beginOpenPanel(completionHandler:)
func (d_ DocumentController) BeginOpenPanelWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("beginOpenPanelWithCompletionHandler:"), completionHandler)
}

// Instantiates a document located by a URL, of a specified type, and returns it if successful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeDocument(withContentsOf:ofType:)
func (d_ DocumentController) MakeDocumentWithContentsOfURLOfTypeError(url foundation.URL, typeName string, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("makeDocumentWithContentsOfURL:ofType:error:"), url, objc.String(typeName), outError)
	return rv
}

// Opens a document located by a URL, optionally presents its user interface, and calls the passed-in completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocument(withContentsOf:display:completionHandler:)
func (d_ DocumentController) OpenDocumentWithContentsOfURLDisplayCompletionHandler(url foundation.URL, displayDocument bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("openDocumentWithContentsOfURL:display:completionHandler:"), url, displayDocument, completionHandler)
}

// The list of recent-document URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/recentDocumentURLs
func (d_ DocumentController) RecentDocumentURLs() []foundation.URL {
	rv := objc.Send[[]foundation.URL](d_.ID, objc.Sel("recentDocumentURLs"))
	return rv
}

// A Boolean value that the system uses to insert a Share menu in the File menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/allowsautomaticsharemenu
func (d_ DocumentController) AllowsAutomaticShareMenu() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("allowsAutomaticShareMenu"))
	return rv
}


// SetAllowsAutomaticShareMenu sets the value of the allowsAutomaticShareMenu property.
// A Boolean value that the system uses to insert a Share menu in the File menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/allowsautomaticsharemenu
func (d_ DocumentController) SetAllowsAutomaticShareMenu(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAllowsAutomaticShareMenu:"), value)
}

// The time interval (in seconds) for periodic autosaving.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/autosavingdelay
func (d_ DocumentController) AutosavingDelay() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("autosavingDelay"))
	return rv
}


// SetAutosavingDelay sets the value of the autosavingDelay property.
// The time interval (in seconds) for periodic autosaving.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/autosavingdelay
func (d_ DocumentController) SetAutosavingDelay(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavingDelay:"), value)
}

// The directory path to use as the starting point in the Open dialog.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/currentdirectory
func (d_ DocumentController) CurrentDirectory() string {
	rv := objc.Send[string](d_.ID, objc.Sel("currentDirectory"))
	return rv
}


// SetCurrentDirectory sets the value of the currentDirectory property.
// The directory path to use as the starting point in the Open dialog.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/currentdirectory
func (d_ DocumentController) SetCurrentDirectory(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCurrentDirectory:"), objc.String(value))
}

// The document object associated with the main window.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/currentdocument
func (d_ DocumentController) CurrentDocument() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("currentDocument"))
	return rv
}


// SetCurrentDocument sets the value of the currentDocument property.
// The document object associated with the main window.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/currentdocument
func (d_ DocumentController) SetCurrentDocument(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setCurrentDocument:"), value)
}

// Returns the name of the document type that should be used when creating new documents.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/defaulttype
func (d_ DocumentController) DefaultType() string {
	rv := objc.Send[string](d_.ID, objc.Sel("defaultType"))
	return rv
}


// SetDefaultType sets the value of the defaultType property.
// Returns the name of the document type that should be used when creating new documents.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/defaulttype
func (d_ DocumentController) SetDefaultType(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDefaultType:"), objc.String(value))
}

// An array of strings representing the custom document classes supported by this app.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/documentclassnames
func (d_ DocumentController) DocumentClassNames() string {
	rv := objc.Send[string](d_.ID, objc.Sel("documentClassNames"))
	return rv
}


// SetDocumentClassNames sets the value of the documentClassNames property.
// An array of strings representing the custom document classes supported by this app.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/documentclassnames
func (d_ DocumentController) SetDocumentClassNames(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDocumentClassNames:"), objc.String(value))
}

// The document objects managed by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/documents
func (d_ DocumentController) Documents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("documents"))
	return rv
}


// SetDocuments sets the value of the documents property.
// The document objects managed by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/documents
func (d_ DocumentController) SetDocuments(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDocuments:"), value)
}

// A Boolean value indicating whether the receiver has any documents with unsaved changes.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/hasediteddocuments
func (d_ DocumentController) HasEditedDocuments() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("hasEditedDocuments"))
	return rv
}


// SetHasEditedDocuments sets the value of the hasEditedDocuments property.
// A Boolean value indicating whether the receiver has any documents with unsaved changes.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/hasediteddocuments
func (d_ DocumentController) SetHasEditedDocuments(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHasEditedDocuments:"), value)
}

// The maximum number of items that may be presented in the standard Open Recent menu.
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/maximumrecentdocumentcount
func (d_ DocumentController) MaximumRecentDocumentCount() int {
	rv := objc.Send[int](d_.ID, objc.Sel("maximumRecentDocumentCount"))
	return rv
}


// SetMaximumRecentDocumentCount sets the value of the maximumRecentDocumentCount property.
// The maximum number of items that may be presented in the standard Open Recent menu.

//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocumentcontroller/maximumrecentdocumentcount
func (d_ DocumentController) SetMaximumRecentDocumentCount(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setMaximumRecentDocumentCount:"), value)
}



