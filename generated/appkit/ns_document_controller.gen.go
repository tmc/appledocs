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
	AddDocument(document IDocument)
	BeginOpenPanelForTypesCompletionHandler(openPanel IOpenPanel, inTypes []string, completionHandler unsafe.Pointer)
	BeginOpenPanelWithCompletionHandler(completionHandler unsafe.Pointer)
	ClearRecentDocuments(sender objectivec.IObject)
	CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo(delegate objectivec.IObject, didCloseAllSelector objc.SEL, contextInfo unsafe.Pointer)
	DisplayNameForType(typeName string) foundation.String
	DocumentForWindow(window IWindow) Document
	DocumentForURL(url foundation.IURL) Document
	DocumentClassForType(typeName string) objc.Class
	DocumentForFileName(fileName string) objc.ID
	DuplicateDocumentWithContentsOfURLCopyingDisplayNameError(url foundation.IURL, duplicateByCopying bool, displayNameOrNil string, outError unsafe.Pointer) Document
	FileExtensionsFromType(typeName string) foundation.Array
	FileNamesFromRunningOpenPanel() foundation.Array
	MakeDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError unsafe.Pointer) Document
	MakeDocumentWithContentsOfURLOfTypeError(url foundation.IURL, typeName string, outError unsafe.Pointer) Document
	MakeDocumentWithContentsOfFileOfType(fileName string, type_ string) objc.ID
	MakeDocumentWithContentsOfURLOfType(url foundation.IURL, type_ string) objc.ID
	MakeUntitledDocumentOfTypeError(typeName string, outError unsafe.Pointer) Document
	MakeUntitledDocumentOfType(type_ string) objc.ID
	NewDocument(sender objectivec.IObject)
	NoteNewRecentDocument(document IDocument)
	NoteNewRecentDocumentURL(url foundation.IURL)
	OpenDocument(sender objectivec.IObject)
	OpenDocumentWithContentsOfURLDisplayCompletionHandler(url foundation.IURL, displayDocument bool, completionHandler unsafe.Pointer)
	OpenDocumentWithContentsOfFileDisplay(fileName string, display bool) objc.ID
	OpenDocumentWithContentsOfURLDisplay(url foundation.IURL, display bool) objc.ID
	OpenDocumentWithContentsOfURLDisplayError(url foundation.IURL, displayDocument bool, outError unsafe.Pointer) objc.ID
	OpenUntitledDocumentAndDisplayError(displayDocument bool, outError unsafe.Pointer) Document
	OpenUntitledDocumentOfTypeDisplay(type_ string, display bool) objc.ID
	PresentError(error_ IError) bool
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ IError, window IWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo unsafe.Pointer)
	RemoveDocument(document IDocument)
	ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler(urlOrNil foundation.IURL, contentsURL foundation.IURL, displayDocument bool, completionHandler unsafe.Pointer)
	ReopenDocumentForURLWithContentsOfURLError(url foundation.IURL, contentsURL foundation.IURL, outError unsafe.Pointer) bool
	ReviewUnsavedDocumentsWithAlertTitleCancellableDelegateDidReviewAllSelectorContextInfo(title string, cancellable bool, delegate objectivec.IObject, didReviewAllSelector objc.SEL, contextInfo unsafe.Pointer)
	RunModalOpenPanelForTypes(openPanel IOpenPanel, types []string) int
	SaveAllDocuments(sender objectivec.IObject)
	SetShouldCreateUI(flag bool)
	ShouldCreateUI() bool
	StandardShareMenuItem() MenuItem
	TypeForContentsOfURLError(url foundation.IURL, outError unsafe.Pointer) foundation.String
	TypeFromFileExtension(fileNameExtensionOrHFSFileType string) foundation.String
	URLsFromRunningOpenPanel() []foundation.URL
	ValidateUserInterfaceItem(item objectivec.IObject) bool
	WillPresentError(error_ IError) Error
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




// This method initializes a new NSDocumentController from the coder.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/init(coder:)
func NewDocumentControllerWithCoder(coder ICoder) DocumentController {
	instance := getDocumentControllerClass().Alloc()
	rv := objc.Send[DocumentController](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Returns the shared instance.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/shared
func (dc _DocumentControllerClass) SharedDocumentController() NSDocumentController {
	rv := objc.Send[NSDocumentController](objc.ID(dc.class), objc.Sel("sharedDocumentController"))
	return rv
}
// Adds the given document to the list of open documents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/addDocument(_:)
func (d_ DocumentController) AddDocument(document IDocument) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addDocument:"), document)
}

// Presents a nonmodal Open dialog that displays files you can open from a list of UTIs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/beginOpenPanel(_:forTypes:completionHandler:)
func (d_ DocumentController) BeginOpenPanelForTypesCompletionHandler(openPanel IOpenPanel, inTypes []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("beginOpenPanel:forTypes:completionHandler:"), openPanel, inTypes, completionHandler)
}

// Presents an Open dialog and delivers the results to a completion handler as an array of URLs for the chosen files, or nil.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/beginOpenPanel(completionHandler:)
func (d_ DocumentController) BeginOpenPanelWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("beginOpenPanelWithCompletionHandler:"), completionHandler)
}

// Empties the recent documents list for the application.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/clearRecentDocuments(_:)
func (d_ DocumentController) ClearRecentDocuments(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("clearRecentDocuments:"), sender)
}

// Iterates through all the open documents and tries to close them one by one using the specified delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/closeAllDocuments(withDelegate:didCloseAllSelector:contextInfo:)
func (d_ DocumentController) CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo(delegate objectivec.IObject, didCloseAllSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("closeAllDocumentsWithDelegate:didCloseAllSelector:contextInfo:"), delegate, didCloseAllSelector, contextInfo)
}

// Returns the descriptive name for the specified document type, which is used in the File Format pop-up menu of the Save As dialog.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/displayName(forType:)
func (d_ DocumentController) DisplayNameForType(typeName string) foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("displayNameForType:"), objc.String(typeName))
	return rv
}

// Returns the document object whose window controller owns a specified window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/document(for:)-a5yd
func (d_ DocumentController) DocumentForWindow(window IWindow) Document {
	rv := objc.Send[Document](d_.ID, objc.Sel("documentForWindow:"), window)
	return rv
}

// Returns, for a given URL, the open document whose file or file package is located by the URL, or if there is no such open document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/document(for:)-i5zi
func (d_ DocumentController) DocumentForURL(url foundation.IURL) Document {
	rv := objc.Send[Document](d_.ID, objc.Sel("documentForURL:"), url)
	return rv
}

// Returns the subclass associated with a given document type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/documentClass(forType:)
func (d_ DocumentController) DocumentClassForType(typeName string) objc.Class {
	rv := objc.Send[objc.Class](d_.ID, objc.Sel("documentClassForType:"), objc.String(typeName))
	return rv
}

// Returns the document object for the file in which the document data is stored.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/documentForFileName:
func (d_ DocumentController) DocumentForFileName(fileName string) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("documentForFileName:"), objc.String(fileName))
	return rv
}

// Creates a new document by reading the contents for the document from another URL, presents its user interface, and returns the document if successful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/duplicateDocument(withContentsOf:copying:displayName:)
func (d_ DocumentController) DuplicateDocumentWithContentsOfURLCopyingDisplayNameError(url foundation.IURL, duplicateByCopying bool, displayNameOrNil string, outError unsafe.Pointer) Document {
	rv := objc.Send[Document](d_.ID, objc.Sel("duplicateDocumentWithContentsOfURL:copying:displayName:error:"), url, duplicateByCopying, objc.String(displayNameOrNil), outError)
	return rv
}

// Returns the allowable file extensions for the given document type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/fileExtensionsFromType:
func (d_ DocumentController) FileExtensionsFromType(typeName string) foundation.Array {
	rv := objc.Send[foundation.Array](d_.ID, objc.Sel("fileExtensionsFromType:"), objc.String(typeName))
	return rv
}

// Returns a selection of files chosen by the user in the Open panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/fileNamesFromRunningOpenPanel
func (d_ DocumentController) FileNamesFromRunningOpenPanel() foundation.Array {
	rv := objc.Send[foundation.Array](d_.ID, objc.Sel("fileNamesFromRunningOpenPanel"))
	return rv
}

// Instantiates a document located by a URL, of a specified type, but by reading the contents for the document from another URL, and returns it if successful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeDocument(for:withContentsOf:ofType:)
func (d_ DocumentController) MakeDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.IURL, contentsURL foundation.IURL, typeName string, outError unsafe.Pointer) Document {
	rv := objc.Send[Document](d_.ID, objc.Sel("makeDocumentForURL:withContentsOfURL:ofType:error:"), urlOrNil, contentsURL, objc.String(typeName), outError)
	return rv
}

// Instantiates a document located by a URL, of a specified type, and returns it if successful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeDocument(withContentsOf:ofType:)
func (d_ DocumentController) MakeDocumentWithContentsOfURLOfTypeError(url foundation.IURL, typeName string, outError unsafe.Pointer) Document {
	rv := objc.Send[Document](d_.ID, objc.Sel("makeDocumentWithContentsOfURL:ofType:error:"), url, objc.String(typeName), outError)
	return rv
}

// Creates and returns a document object of a given document type from the contents of a file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeDocumentWithContentsOfFile:ofType:
func (d_ DocumentController) MakeDocumentWithContentsOfFileOfType(fileName string, type_ string) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("makeDocumentWithContentsOfFile:ofType:"), objc.String(fileName), objc.String(type_))
	return rv
}

// Creates and returns a document object for the given document type from the contents of a given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeDocumentWithContentsOfURL:ofType:
func (d_ DocumentController) MakeDocumentWithContentsOfURLOfType(url foundation.IURL, type_ string) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("makeDocumentWithContentsOfURL:ofType:"), url, objc.String(type_))
	return rv
}

// Instantiates a new untitled document of the specified type and returns it if successful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeUntitledDocument(ofType:)
func (d_ DocumentController) MakeUntitledDocumentOfTypeError(typeName string, outError unsafe.Pointer) Document {
	rv := objc.Send[Document](d_.ID, objc.Sel("makeUntitledDocumentOfType:error:"), objc.String(typeName), outError)
	return rv
}

// Creates and returns a document object for document type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeUntitledDocumentOfType:
func (d_ DocumentController) MakeUntitledDocumentOfType(type_ string) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("makeUntitledDocumentOfType:"), objc.String(type_))
	return rv
}

// An action method called by the New menu command, this method creates a new object and adds it to the list of such objects managed by the document controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/newDocument(_:)
func (d_ DocumentController) NewDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("newDocument:"), sender)
}

// Adds or replaces an Open Recent menu item corresponding to the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/noteNewRecentDocument(_:)
func (d_ DocumentController) NoteNewRecentDocument(document IDocument) {
	objc.Send[objc.ID](d_.ID, objc.Sel("noteNewRecentDocument:"), document)
}

// Adds or replaces an Open Recent menu item corresponding to the data located by the URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/noteNewRecentDocumentURL(_:)
func (d_ DocumentController) NoteNewRecentDocumentURL(url foundation.IURL) {
	objc.Send[objc.ID](d_.ID, objc.Sel("noteNewRecentDocumentURL:"), url)
}

// An action method called by the Open menu command, it runs the modal Open panel and, based on the selected filenames, creates one or more objects from the contents of the files.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocument(_:)
func (d_ DocumentController) OpenDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("openDocument:"), sender)
}

// Opens a document located by a URL, optionally presents its user interface, and calls the passed-in completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocument(withContentsOf:display:completionHandler:)
func (d_ DocumentController) OpenDocumentWithContentsOfURLDisplayCompletionHandler(url foundation.IURL, displayDocument bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("openDocumentWithContentsOfURL:display:completionHandler:"), url, displayDocument, completionHandler)
}

// Returns a document object created from the contents of a given file and optionally displays it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocumentWithContentsOfFile:display:
func (d_ DocumentController) OpenDocumentWithContentsOfFileDisplay(fileName string, display bool) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("openDocumentWithContentsOfFile:display:"), objc.String(fileName), display)
	return rv
}

// Returns a document object created from the contents of a given URL and optionally displays it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocumentWithContentsOfURL:display:
func (d_ DocumentController) OpenDocumentWithContentsOfURLDisplay(url foundation.IURL, display bool) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("openDocumentWithContentsOfURL:display:"), url, display)
	return rv
}

// Opens a document located by the given URL presents its user interface if requested, and returns the document if successful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocumentWithContentsOfURL:display:error:
func (d_ DocumentController) OpenDocumentWithContentsOfURLDisplayError(url foundation.IURL, displayDocument bool, outError unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("openDocumentWithContentsOfURL:display:error:"), url, displayDocument, outError)
	return rv
}

// Creates a new untitled document, presents its user interface if is , and returns the document if successful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openUntitledDocumentAndDisplay(_:)
func (d_ DocumentController) OpenUntitledDocumentAndDisplayError(displayDocument bool, outError unsafe.Pointer) Document {
	rv := objc.Send[Document](d_.ID, objc.Sel("openUntitledDocumentAndDisplay:error:"), displayDocument, outError)
	return rv
}

// Returns a document object instantiated from the subclass of the given document type and optionally displays it.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openUntitledDocumentOfType:display:
func (d_ DocumentController) OpenUntitledDocumentOfTypeDisplay(type_ string, display bool) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("openUntitledDocumentOfType:display:"), objc.String(type_), display)
	return rv
}

// Presents an error alert to the user as a modal panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/presentError(_:)
func (d_ DocumentController) PresentError(error_ IError) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("presentError:"), error_)
	return rv
}

// Presents an error alert to the user as a modal panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (d_ DocumentController) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ IError, window IWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error_, window, delegate, didPresentSelector, contextInfo)
}

// Removes the given document from the list of open documents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/removeDocument(_:)
func (d_ DocumentController) RemoveDocument(document IDocument) {
	objc.Send[objc.ID](d_.ID, objc.Sel("removeDocument:"), document)
}

// Reopens a document, optionally located by a URL, by reading the contents for the document from another URL, optionally presents its user interface, and calls the passed-in completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/reopenDocument(for:withContentsOf:display:completionHandler:)
func (d_ DocumentController) ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler(urlOrNil foundation.IURL, contentsURL foundation.IURL, displayDocument bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("reopenDocumentForURL:withContentsOfURL:display:completionHandler:"), urlOrNil, contentsURL, displayDocument, completionHandler)
}

// Reopens an autosaved document located by a URL, by reading the contents for the document from another URL, presents its user interface, and returns if successful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/reopenDocumentForURL:withContentsOfURL:error:
func (d_ DocumentController) ReopenDocumentForURLWithContentsOfURLError(url foundation.IURL, contentsURL foundation.IURL, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("reopenDocumentForURL:withContentsOfURL:error:"), url, contentsURL, outError)
	return rv
}

// Displays an alert asking if the user wants to review unsaved documents, quit regardless of unsaved documents, or cancel the save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/reviewUnsavedDocuments(withAlertTitle:cancellable:delegate:didReviewAllSelector:contextInfo:)
func (d_ DocumentController) ReviewUnsavedDocumentsWithAlertTitleCancellableDelegateDidReviewAllSelectorContextInfo(title string, cancellable bool, delegate objectivec.IObject, didReviewAllSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("reviewUnsavedDocumentsWithAlertTitle:cancellable:delegate:didReviewAllSelector:contextInfo:"), objc.String(title), cancellable, delegate, didReviewAllSelector, contextInfo)
}

// Presents a modal Open dialog and limits selection to specific file types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/runModalOpenPanel(_:forTypes:)
func (d_ DocumentController) RunModalOpenPanelForTypes(openPanel IOpenPanel, types []string) int {
	rv := objc.Send[int](d_.ID, objc.Sel("runModalOpenPanel:forTypes:"), openPanel, types)
	return rv
}

// As the action method called by the Save All command, saves all open documents of the application that need to be saved.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/saveAllDocuments(_:)
func (d_ DocumentController) SaveAllDocuments(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveAllDocuments:"), sender)
}

// Sets whether the window controllers of a document should be created when the document is created.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/setShouldCreateUI:
func (d_ DocumentController) SetShouldCreateUI(flag bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShouldCreateUI:"), flag)
}

// Returns a Boolean value that indicates whether the window controllers of a document should be created when the document is created.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/shouldCreateUI
func (d_ DocumentController) ShouldCreateUI() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldCreateUI"))
	return rv
}

// Returns a menu item that your app uses for sharing the current document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/standardShareMenuItem()
func (d_ DocumentController) StandardShareMenuItem() MenuItem {
	rv := objc.Send[MenuItem](d_.ID, objc.Sel("standardShareMenuItem"))
	return rv
}

// Returns, for a specified URL, the document type identifier to use when opening the document at that location, if successful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/typeForContents(of:)
func (d_ DocumentController) TypeForContentsOfURLError(url foundation.IURL, outError unsafe.Pointer) foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("typeForContentsOfURL:error:"), url, outError)
	return rv
}

// Returns the document type associated with files having extension .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/typeFromFileExtension:
func (d_ DocumentController) TypeFromFileExtension(fileNameExtensionOrHFSFileType string) foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("typeFromFileExtension:"), objc.String(fileNameExtensionOrHFSFileType))
	return rv
}

// An array of URLs that correspond to the selected files in a running Open dialog.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/urlsFromRunningOpenPanel()
func (d_ DocumentController) URLsFromRunningOpenPanel() []foundation.URL {
	rv := objc.Send[[]foundation.URL](d_.ID, objc.Sel("URLsFromRunningOpenPanel"))
	return rv
}

// Returns a Boolean value that indicates whether a given user interface item should be enabled.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/validateUserInterfaceItem(_:)
func (d_ DocumentController) ValidateUserInterfaceItem(item objectivec.IObject) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}

// Indicates an error condition and provides the opportunity to return the same or a different error.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/willPresentError(_:)
func (d_ DocumentController) WillPresentError(error_ IError) Error {
	rv := objc.Send[Error](d_.ID, objc.Sel("willPresentError:"), error_)
	return rv
}

// A Boolean value that the system uses to insert a Share menu in the File menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/allowsAutomaticShareMenu
func (d_ DocumentController) AllowsAutomaticShareMenu() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("allowsAutomaticShareMenu"))
	return rv
}

// The time interval (in seconds) for periodic autosaving.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/autosavingDelay
func (d_ DocumentController) AutosavingDelay() float64 {
	rv := objc.Send[float64](d_.ID, objc.Sel("autosavingDelay"))
	return rv
}


// SetAutosavingDelay sets the value of the autosavingDelay property.
// The time interval (in seconds) for periodic autosaving.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/autosavingDelay
func (d_ DocumentController) SetAutosavingDelay(value float64) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavingDelay:"), value)
}

// The directory path to use as the starting point in the Open dialog.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/currentDirectory
func (d_ DocumentController) CurrentDirectory() string {
	rv := objc.Send[string](d_.ID, objc.Sel("currentDirectory"))
	return rv
}

// The document object associated with the main window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/currentDocument
func (d_ DocumentController) CurrentDocument() NSDocument {
	rv := objc.Send[NSDocument](d_.ID, objc.Sel("currentDocument"))
	return rv
}

// Returns the name of the document type that should be used when creating new documents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/defaultType
func (d_ DocumentController) DefaultType() string {
	rv := objc.Send[string](d_.ID, objc.Sel("defaultType"))
	return rv
}

// An array of strings representing the custom document classes supported by this app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/documentClassNames
func (d_ DocumentController) DocumentClassNames() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("documentClassNames"))
	return rv
}

// The document objects managed by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/documents
func (d_ DocumentController) Documents() []Document {
	rv := objc.Send[[]Document](d_.ID, objc.Sel("documents"))
	return rv
}

// A Boolean value indicating whether the receiver has any documents with unsaved changes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/hasEditedDocuments
func (d_ DocumentController) HasEditedDocuments() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("hasEditedDocuments"))
	return rv
}

// The maximum number of items that may be presented in the standard Open Recent menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/maximumRecentDocumentCount
func (d_ DocumentController) MaximumRecentDocumentCount() uint {
	rv := objc.Send[uint](d_.ID, objc.Sel("maximumRecentDocumentCount"))
	return rv
}

// The list of recent-document URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/recentDocumentURLs
func (d_ DocumentController) RecentDocumentURLs() []foundation.URL {
	rv := objc.Send[[]foundation.URL](d_.ID, objc.Sel("recentDocumentURLs"))
	return rv
}

// Returns the shared instance.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocumentController/shared
func (d_ DocumentController) SharedDocumentController() NSDocumentController {
	rv := objc.Send[NSDocumentController](d_.ID, objc.Sel("sharedDocumentController"))
	return rv
}


