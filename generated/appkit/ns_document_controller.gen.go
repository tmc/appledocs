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
	AllowsAutomaticShareMenu() bool /* primitive/slice/pointer. */
	AutosavingDelay() float64 /* primitive/slice/pointer. */
	SetAutosavingDelay(value float64 /* primitive/slice/pointer. */)
	CurrentDirectory() string /* primitive/slice/pointer. */
	CurrentDocument() objc.IObject /* cross-framework: Document */
	DefaultType() string /* primitive/slice/pointer. */
	DocumentClassNames() []string /* primitive/slice/pointer. */
	Documents() []Document /* primitive/slice/pointer. */
	HasEditedDocuments() bool /* primitive/slice/pointer. */
	MaximumRecentDocumentCount() uint /* primitive/slice/pointer. */
	RecentDocumentURLs() []foundation.objc.IObject /* cross-framework: URL */
	// methods:
	AddDocument(document objc.IObject /* cross-framework Document */)
	BeginOpenPanelForTypesCompletionHandler(openPanel IOpenPanel, inTypes []string /* primitive/slice/pointer. */, completionHandler unsafe.Pointer)
	BeginOpenPanelWithCompletionHandler(completionHandler unsafe.Pointer)
	ClearRecentDocuments(sender objectivec.IObject)
	CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo(delegate objectivec.IObject, didCloseAllSelector objc.SEL, contextInfo unsafe.Pointer)
	DisplayNameForType(typeName string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */
	DocumentForWindow(window IWindow) objc.IObject /* cross-framework: Document */
	DocumentForURL(url foundation.objc.IObject /* cross-framework URL */) objc.IObject /* cross-framework: Document */
	DocumentClassForType(typeName string /* primitive/slice/pointer. */) objc.Class
	DuplicateDocumentWithContentsOfURLCopyingDisplayNameError(url foundation.objc.IObject /* cross-framework URL */, duplicateByCopying bool /* primitive/slice/pointer. */, displayNameOrNil string /* primitive/slice/pointer. */, outError unsafe.Pointer) objc.IObject /* cross-framework: Document */
	MakeDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.objc.IObject /* cross-framework URL */, contentsURL foundation.objc.IObject /* cross-framework URL */, typeName string /* primitive/slice/pointer. */, outError unsafe.Pointer) objc.IObject /* cross-framework: Document */
	MakeDocumentWithContentsOfURLOfTypeError(url foundation.objc.IObject /* cross-framework URL */, typeName string /* primitive/slice/pointer. */, outError unsafe.Pointer) objc.IObject /* cross-framework: Document */
	MakeUntitledDocumentOfTypeError(typeName string /* primitive/slice/pointer. */, outError unsafe.Pointer) objc.IObject /* cross-framework: Document */
	NewDocument(sender objectivec.IObject)
	NoteNewRecentDocument(document objc.IObject /* cross-framework Document */)
	NoteNewRecentDocumentURL(url foundation.objc.IObject /* cross-framework URL */)
	OpenDocument(sender objectivec.IObject)
	OpenDocumentWithContentsOfURLDisplayCompletionHandler(url foundation.objc.IObject /* cross-framework URL */, displayDocument bool /* primitive/slice/pointer. */, completionHandler unsafe.Pointer)
	OpenUntitledDocumentAndDisplayError(displayDocument bool /* primitive/slice/pointer. */, outError unsafe.Pointer) objc.IObject /* cross-framework: Document */
	PresentError(error_ Error /* not a class type */) bool /* primitive/slice/pointer. */
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ Error /* not a class type */, window IWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo unsafe.Pointer)
	RemoveDocument(document objc.IObject /* cross-framework Document */)
	ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler(urlOrNil foundation.objc.IObject /* cross-framework URL */, contentsURL foundation.objc.IObject /* cross-framework URL */, displayDocument bool /* primitive/slice/pointer. */, completionHandler unsafe.Pointer)
	ReviewUnsavedDocumentsWithAlertTitleCancellableDelegateDidReviewAllSelectorContextInfo(title string /* primitive/slice/pointer. */, cancellable bool /* primitive/slice/pointer. */, delegate objectivec.IObject, didReviewAllSelector objc.SEL, contextInfo unsafe.Pointer)
	RunModalOpenPanelForTypes(openPanel IOpenPanel, types []string /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */
	SaveAllDocuments(sender objectivec.IObject)
	StandardShareMenuItem() objc.IObject /* cross-framework: MenuItem */
	TypeForContentsOfURLError(url foundation.objc.IObject /* cross-framework URL */, outError unsafe.Pointer) objc.IObject /* cross-framework: String */
	URLsFromRunningOpenPanel() []foundation.objc.IObject /* cross-framework: URL */
	ValidateUserInterfaceItem(item objectivec.IObject) bool /* primitive/slice/pointer. */
	WillPresentError(error_ Error /* not a class type */) Error /* not a class type */
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



// This method initializes a new NSDocumentController from the coder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/init(coder:)
func NewDocumentControllerWithCoder(coder Coder /* not a class type */) DocumentController {
	instance := getDocumentControllerClass().Alloc()
	rv := objc.Send[DocumentController](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}



// Returns the shared instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/shared
func (dc _DocumentControllerClass) SharedDocumentController() DocumentController {
	rv := objc.Send[DocumentController](objc.ID(dc.class), objc.Sel("sharedDocumentController"))
	return rv
}

// Adds the given document to the list of open documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/addDocument(_:)
func (d_ DocumentController) AddDocument(document objc.IObject /* cross-framework Document */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addDocument:"), document)
}


// Presents a nonmodal Open dialog that displays files you can open from a list of UTIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/beginOpenPanel(_:forTypes:completionHandler:)
func (d_ DocumentController) BeginOpenPanelForTypesCompletionHandler(openPanel IOpenPanel, inTypes []string /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("beginOpenPanel:forTypes:completionHandler:"), openPanel, inTypes, completionHandler)
}


// Presents an Open dialog and delivers the results to a completion handler as an array of URLs for the chosen files, or nil.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/beginOpenPanel(completionHandler:)
func (d_ DocumentController) BeginOpenPanelWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("beginOpenPanelWithCompletionHandler:"), completionHandler)
}


// Empties the recent documents list for the application.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/clearRecentDocuments(_:)
func (d_ DocumentController) ClearRecentDocuments(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("clearRecentDocuments:"), sender)
}


// Iterates through all the open documents and tries to close them one by one using the specified delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/closeAllDocuments(withDelegate:didCloseAllSelector:contextInfo:)
func (d_ DocumentController) CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo(delegate objectivec.IObject, didCloseAllSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("closeAllDocumentsWithDelegate:didCloseAllSelector:contextInfo:"), delegate, didCloseAllSelector, contextInfo)
}


// Returns the descriptive name for the specified document type, which is used in the File Format pop-up menu of the Save As dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/displayName(forType:)
func (d_ DocumentController) DisplayNameForType(typeName string /* primitive/slice/pointer. */) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](d_.ID, objc.Sel("displayNameForType:"), objc.String(typeName))
	return rv
}


// Returns the document object whose window controller owns a specified window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/document(for:)-a5yd
func (d_ DocumentController) DocumentForWindow(window IWindow) objc.IObject /* cross-framework: Document */ {
	rv := objc.Send[Document](d_.ID, objc.Sel("documentForWindow:"), window)
	return rv
}


// Returns, for a given URL, the open document whose file or file package is located by the URL, or if there is no such open document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/document(for:)-i5zi
func (d_ DocumentController) DocumentForURL(url foundation.objc.IObject /* cross-framework URL */) objc.IObject /* cross-framework: Document */ {
	rv := objc.Send[Document](d_.ID, objc.Sel("documentForURL:"), url)
	return rv
}


// Returns the subclass associated with a given document type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/documentClass(forType:)
func (d_ DocumentController) DocumentClassForType(typeName string /* primitive/slice/pointer. */) objc.Class {
	rv := objc.Send[objc.Class](d_.ID, objc.Sel("documentClassForType:"), objc.String(typeName))
	return rv
}


// Creates a new document by reading the contents for the document from another URL, presents its user interface, and returns the document if successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/duplicateDocument(withContentsOf:copying:displayName:)
func (d_ DocumentController) DuplicateDocumentWithContentsOfURLCopyingDisplayNameError(url foundation.objc.IObject /* cross-framework URL */, duplicateByCopying bool /* primitive/slice/pointer. */, displayNameOrNil string /* primitive/slice/pointer. */, outError unsafe.Pointer) objc.IObject /* cross-framework: Document */ {
	rv := objc.Send[Document](d_.ID, objc.Sel("duplicateDocumentWithContentsOfURL:copying:displayName:error:"), url, duplicateByCopying, objc.String(displayNameOrNil), outError)
	return rv
}


// Instantiates a document located by a URL, of a specified type, but by reading the contents for the document from another URL, and returns it if successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeDocument(for:withContentsOf:ofType:)
func (d_ DocumentController) MakeDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.objc.IObject /* cross-framework URL */, contentsURL foundation.objc.IObject /* cross-framework URL */, typeName string /* primitive/slice/pointer. */, outError unsafe.Pointer) objc.IObject /* cross-framework: Document */ {
	rv := objc.Send[Document](d_.ID, objc.Sel("makeDocumentForURL:withContentsOfURL:ofType:error:"), urlOrNil, contentsURL, objc.String(typeName), outError)
	return rv
}


// Instantiates a document located by a URL, of a specified type, and returns it if successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeDocument(withContentsOf:ofType:)
func (d_ DocumentController) MakeDocumentWithContentsOfURLOfTypeError(url foundation.objc.IObject /* cross-framework URL */, typeName string /* primitive/slice/pointer. */, outError unsafe.Pointer) objc.IObject /* cross-framework: Document */ {
	rv := objc.Send[Document](d_.ID, objc.Sel("makeDocumentWithContentsOfURL:ofType:error:"), url, objc.String(typeName), outError)
	return rv
}


// Instantiates a new untitled document of the specified type and returns it if successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeUntitledDocument(ofType:)
func (d_ DocumentController) MakeUntitledDocumentOfTypeError(typeName string /* primitive/slice/pointer. */, outError unsafe.Pointer) objc.IObject /* cross-framework: Document */ {
	rv := objc.Send[Document](d_.ID, objc.Sel("makeUntitledDocumentOfType:error:"), objc.String(typeName), outError)
	return rv
}


// An action method called by the New menu command, this method creates a new object and adds it to the list of such objects managed by the document controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/newDocument(_:)
func (d_ DocumentController) NewDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("newDocument:"), sender)
}


// Adds or replaces an Open Recent menu item corresponding to the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/noteNewRecentDocument(_:)
func (d_ DocumentController) NoteNewRecentDocument(document objc.IObject /* cross-framework Document */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("noteNewRecentDocument:"), document)
}


// Adds or replaces an Open Recent menu item corresponding to the data located by the URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/noteNewRecentDocumentURL(_:)
func (d_ DocumentController) NoteNewRecentDocumentURL(url foundation.objc.IObject /* cross-framework URL */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("noteNewRecentDocumentURL:"), url)
}


// An action method called by the Open menu command, it runs the modal Open panel and, based on the selected filenames, creates one or more objects from the contents of the files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocument(_:)
func (d_ DocumentController) OpenDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("openDocument:"), sender)
}


// Opens a document located by a URL, optionally presents its user interface, and calls the passed-in completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocument(withContentsOf:display:completionHandler:)
func (d_ DocumentController) OpenDocumentWithContentsOfURLDisplayCompletionHandler(url foundation.objc.IObject /* cross-framework URL */, displayDocument bool /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("openDocumentWithContentsOfURL:display:completionHandler:"), url, displayDocument, completionHandler)
}


// Creates a new untitled document, presents its user interface if is , and returns the document if successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openUntitledDocumentAndDisplay(_:)
func (d_ DocumentController) OpenUntitledDocumentAndDisplayError(displayDocument bool /* primitive/slice/pointer. */, outError unsafe.Pointer) objc.IObject /* cross-framework: Document */ {
	rv := objc.Send[Document](d_.ID, objc.Sel("openUntitledDocumentAndDisplay:error:"), displayDocument, outError)
	return rv
}


// Presents an error alert to the user as a modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/presentError(_:)
func (d_ DocumentController) PresentError(error_ Error /* not a class type */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("presentError:"), error_)
	return rv
}


// Presents an error alert to the user as a modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (d_ DocumentController) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ Error /* not a class type */, window IWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error_, window, delegate, didPresentSelector, contextInfo)
}


// Removes the given document from the list of open documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/removeDocument(_:)
func (d_ DocumentController) RemoveDocument(document objc.IObject /* cross-framework Document */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("removeDocument:"), document)
}


// Reopens a document, optionally located by a URL, by reading the contents for the document from another URL, optionally presents its user interface, and calls the passed-in completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/reopenDocument(for:withContentsOf:display:completionHandler:)
func (d_ DocumentController) ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler(urlOrNil foundation.objc.IObject /* cross-framework URL */, contentsURL foundation.objc.IObject /* cross-framework URL */, displayDocument bool /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("reopenDocumentForURL:withContentsOfURL:display:completionHandler:"), urlOrNil, contentsURL, displayDocument, completionHandler)
}


// Displays an alert asking if the user wants to review unsaved documents, quit regardless of unsaved documents, or cancel the save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/reviewUnsavedDocuments(withAlertTitle:cancellable:delegate:didReviewAllSelector:contextInfo:)
func (d_ DocumentController) ReviewUnsavedDocumentsWithAlertTitleCancellableDelegateDidReviewAllSelectorContextInfo(title string /* primitive/slice/pointer. */, cancellable bool /* primitive/slice/pointer. */, delegate objectivec.IObject, didReviewAllSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("reviewUnsavedDocumentsWithAlertTitle:cancellable:delegate:didReviewAllSelector:contextInfo:"), objc.String(title), cancellable, delegate, didReviewAllSelector, contextInfo)
}


// Presents a modal Open dialog and limits selection to specific file types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/runModalOpenPanel(_:forTypes:)
func (d_ DocumentController) RunModalOpenPanelForTypes(openPanel IOpenPanel, types []string /* primitive/slice/pointer. */) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](d_.ID, objc.Sel("runModalOpenPanel:forTypes:"), openPanel, types)
	return rv
}


// As the action method called by the Save All command, saves all open documents of the application that need to be saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/saveAllDocuments(_:)
func (d_ DocumentController) SaveAllDocuments(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveAllDocuments:"), sender)
}


// Returns a menu item that your app uses for sharing the current document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/standardShareMenuItem()
func (d_ DocumentController) StandardShareMenuItem() objc.IObject /* cross-framework: MenuItem */ {
	rv := objc.Send[MenuItem](d_.ID, objc.Sel("standardShareMenuItem"))
	return rv
}


// Returns, for a specified URL, the document type identifier to use when opening the document at that location, if successful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/typeForContents(of:)
func (d_ DocumentController) TypeForContentsOfURLError(url foundation.objc.IObject /* cross-framework URL */, outError unsafe.Pointer) objc.IObject /* cross-framework: String */ {
	rv := objc.Send[String](d_.ID, objc.Sel("typeForContentsOfURL:error:"), url, outError)
	return rv
}


// An array of URLs that correspond to the selected files in a running Open dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/urlsFromRunningOpenPanel()
func (d_ DocumentController) URLsFromRunningOpenPanel() []foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[[]foundation.URL](d_.ID, objc.Sel("URLsFromRunningOpenPanel"))
	return rv
}


// Returns a Boolean value that indicates whether a given user interface item should be enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/validateUserInterfaceItem(_:)
func (d_ DocumentController) ValidateUserInterfaceItem(item objectivec.IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}


// Indicates an error condition and provides the opportunity to return the same or a different error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/willPresentError(_:)
func (d_ DocumentController) WillPresentError(error_ Error /* not a class type */) Error /* not a class type */ {
	rv := objc.Send[Error](d_.ID, objc.Sel("willPresentError:"), error_)
	return rv
}


// A Boolean value that the system uses to insert a Share menu in the File menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/allowsAutomaticShareMenu
func (d_ DocumentController) AllowsAutomaticShareMenu() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("allowsAutomaticShareMenu"))
	return rv
}


// The time interval (in seconds) for periodic autosaving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/autosavingDelay
func (d_ DocumentController) AutosavingDelay() float64 /* primitive/slice/pointer. */ {
	rv := objc.Send[float64](d_.ID, objc.Sel("autosavingDelay"))
	return rv
}


// The time interval (in seconds) for periodic autosaving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/autosavingDelay
func (d_ DocumentController) SetAutosavingDelay(value float64 /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavingDelay:"), value)
}


// The directory path to use as the starting point in the Open dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/currentDirectory
func (d_ DocumentController) CurrentDirectory() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("currentDirectory"))
	return rv
}


// The document object associated with the main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/currentDocument
func (d_ DocumentController) CurrentDocument() objc.IObject /* cross-framework: Document */ {
	rv := objc.Send[Document](d_.ID, objc.Sel("currentDocument"))
	return rv
}


// Returns the name of the document type that should be used when creating new documents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/defaultType
func (d_ DocumentController) DefaultType() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](d_.ID, objc.Sel("defaultType"))
	return rv
}


// An array of strings representing the custom document classes supported by this app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/documentClassNames
func (d_ DocumentController) DocumentClassNames() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](d_.ID, objc.Sel("documentClassNames"))
	return rv
}


// The document objects managed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/documents
func (d_ DocumentController) Documents() []Document /* primitive/slice/pointer. */ {
	rv := objc.Send[[]Document](d_.ID, objc.Sel("documents"))
	return rv
}


// A Boolean value indicating whether the receiver has any documents with unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/hasEditedDocuments
func (d_ DocumentController) HasEditedDocuments() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](d_.ID, objc.Sel("hasEditedDocuments"))
	return rv
}


// The maximum number of items that may be presented in the standard Open Recent menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/maximumRecentDocumentCount
func (d_ DocumentController) MaximumRecentDocumentCount() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](d_.ID, objc.Sel("maximumRecentDocumentCount"))
	return rv
}


// The list of recent-document URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/recentDocumentURLs
func (d_ DocumentController) RecentDocumentURLs() []foundation.objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[[]foundation.URL](d_.ID, objc.Sel("recentDocumentURLs"))
	return rv
}


// Returns the shared instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/shared
func (d_ DocumentController) SharedDocumentController() IDocumentController {
	rv := objc.Send[DocumentController](d_.ID, objc.Sel("sharedDocumentController"))
	return rv
}


