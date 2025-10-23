// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	AutosavingIsImplicitlyCancellable() bool
	BackupFileURL() foundation.URL
	FileNameExtensionWasHiddenInLastRunSavePanel() bool
	FileTypeFromLastRunSavePanel() string
	FileURL() foundation.URL
	SetFileURL(value foundation.URL)
	HasUnautosavedChanges() bool
	Draft() bool
	SetDraft(value bool)
	Locked() bool
	KeepBackupFile() bool
	PrintInfo() IPrintInfo
	SetPrintInfo(value IPrintInfo)
	WindowNibName() NibName
	AllowsDocumentSharing() bool
	SetAllowsDocumentSharing(value bool)
	AutosavedContentsFileURL() foundation.URL
	SetAutosavedContentsFileURL(value foundation.URL)
	AutosavingFileType() string
	SetAutosavingFileType(value string)
	DisplayName() string
	SetDisplayName(value string)
	FileModificationDate() foundation.Date
	SetFileModificationDate(value foundation.Date)
	FileType() string
	SetFileType(value string)
	HasUndoManager() bool
	SetHasUndoManager(value bool)
	IsBrowsingVersions() bool
	SetIsBrowsingVersions(value bool)
	IsDocumentEdited() bool
	SetIsDocumentEdited(value bool)
	IsDraft() bool
	SetIsDraft(value bool)
	IsEntireFileLoaded() bool
	SetIsEntireFileLoaded(value bool)
	IsInViewingMode() bool
	SetIsInViewingMode(value bool)
	IsLocked() bool
	SetIsLocked(value bool)
	LastComponentOfFileName() string
	SetLastComponentOfFileName(value string)
	ObjectSpecifier() foundation.ScriptObjectSpecifier
	SetObjectSpecifier(value foundation.ScriptObjectSpecifier)
	ObservedPresentedItemUbiquityAttributes() unsafe.Pointer
	SetObservedPresentedItemUbiquityAttributes(value unsafe.Pointer)
	PdfPrintOperation() IPrintOperation
	SetPdfPrintOperation(value IPrintOperation)
	PresentedItemURL() foundation.URL
	SetPresentedItemURL(value foundation.URL)
	PreviewRepresentableActivityItems() unsafe.Pointer
	SetPreviewRepresentableActivityItems(value unsafe.Pointer)
	SavePanelShowsFileFormatsControl() bool
	SetSavePanelShowsFileFormatsControl(value bool)
	ShouldRunSavePanelWithAccessoryView() bool
	SetShouldRunSavePanelWithAccessoryView(value bool)
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.UndoManager)
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.UserActivity)
	WindowControllers() IWindowController
	SetWindowControllers(value IWindowController)
	WindowForSheet() IWindow
	SetWindowForSheet(value IWindow)
	NSUserActivityDocumentURLKey() string
	AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objectivec.IObject, didAutosaveSelector objc.SEL, contextInfo unsafe.Pointer)
	Close()
	ContinueAsynchronousWorkOnMainThreadUsingBlock(block unsafe.Pointer)
	DataOfTypeError(typeName string, outError unsafe.Pointer) foundation.Data
	DefaultDraftName() foundation.String
	DuplicateDocument(sender objectivec.IObject)
	DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objectivec.IObject, didDuplicateSelector objc.SEL, contextInfo unsafe.Pointer)
	LockDocument(sender objectivec.IObject)
	LockWithCompletionHandler(completionHandler unsafe.Pointer)
	LockDocumentWithCompletionHandler(completionHandler unsafe.Pointer)
	PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block unsafe.Pointer)
	PreparePageLayout(pageLayout IPageLayout) bool
	PrepareSavePanel(savePanel SavePanel) bool
	PresentedItemDidMoveToURL(newURL foundation.URL)
	ReadFromURLOfTypeError(url foundation.URL, typeName string, outError unsafe.Pointer) bool
	RelinquishPresentedItemToReader(reader unsafe.Pointer)
	RelinquishPresentedItemToWriter(writer unsafe.Pointer)
	RemoveWindowController(windowController IWindowController)
	RenameDocument(sender objectivec.IObject)
	RevertDocumentToSaved(sender objectivec.IObject)
	RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo(printInfo IPrintInfo, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo unsafe.Pointer)
	RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation IPrintOperation, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo unsafe.Pointer)
	RunPageLayout(sender objectivec.IObject)
	SaveToURLOfTypeForSaveOperationCompletionHandler(url foundation.URL, typeName string, saveOperation unsafe.Pointer, completionHandler unsafe.Pointer)
	SetWindow(window IWindow)
	StopBrowsingVersionsWithCompletionHandler(completionHandler unsafe.Pointer)
	UnlockDocumentWithCompletionHandler(completionHandler unsafe.Pointer)
	UpdateChangeCount(change unsafe.Pointer)
	ValidateUserInterfaceItem(item objectivec.IObject) bool
}

// An abstract class that defines the interface for macOS documents.
//
// A document is an object that can internally represent data displayed in a window and that can read data from and write data to a file or file package. Documents create and manage one or more window controllers and are in turn managed by a document controller. Documents respond to first-responder action messages to save, revert, and print their data. Conceptually, a document is a container for a body of information identified by a name under which it is stored in a disk file. In this sense, however, the document is not the same as the file but is an object in memory that owns and manages the document data. In the context of AppKit, a document is an instance of a custom subclass that knows how to represent internally, in one or more formats, persistent data that is displayed in windows. A document can read that data from a file and write it to a file. It is also the first-responder target for many menu commands related to documents, such as Save, Revert, and Print. A document manages its window’s edited status and is set up to perform undo and redo operations. When a window is closing, the document is asked before the window delegate to approve the closing. is one of the triad of AppKit classes that establish an architectural basis for document-based apps (the others being and ). For more information about using in a document-based app, see .


// An abstract class that defines the interface for macOS documents.
//
// [Full Topic]
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



// Returns a Boolean value that indicates whether the document can read and write the data natively.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isNativeType(_:)
func (dc _DocumentClass) IsNativeType(type_ string) bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("isNativeType:"), objc.String(type_))
	return rv
}


// Returns the types of data the receiver can read natively and any types filterable to that native type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/readableTypes
func (dc _DocumentClass) ReadableTypes() []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("readableTypes"))
	return rv
}

// Returns an array of key paths that represent the restorable attributes of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/restorableStateKeyPaths
func (dc _DocumentClass) RestorableStateKeyPaths() []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("restorableStateKeyPaths"))
	return rv
}

// Returns whether the document object stores its contents in the user’s iCloud document storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/usesUbiquitousStorage
func (dc _DocumentClass) UsesUbiquitousStorage() bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("usesUbiquitousStorage"))
	return rv
}

// Autosaves the document’s contents to an appropriate location in the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosave(withDelegate:didAutosave:contextInfo:)
func (d_ Document) AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objectivec.IObject, didAutosaveSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("autosaveDocumentWithDelegate:didAutosaveSelector:contextInfo:"), delegate, didAutosaveSelector, contextInfo)
}


// Closes all of the document’s windows and removes the document from its document controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/close()
func (d_ Document) Close() {
	objc.Send[objc.ID](d_.ID, objc.Sel("close"))
}


// Invokes the passed-in block on the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/continueAsynchronousWorkOnMainThread(_:)
func (d_ Document) ContinueAsynchronousWorkOnMainThreadUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("continueAsynchronousWorkOnMainThreadUsingBlock:"), block)
}


// Creates and returns a data object that contains the contents of the document, formatted to a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/data(ofType:)
func (d_ Document) DataOfTypeError(typeName string, outError unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](d_.ID, objc.Sel("dataOfType:error:"), objc.String(typeName), outError)
	return rv
}


// Returns the default draft name for the document subclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/defaultDraftName()
func (d_ Document) DefaultDraftName() foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("defaultDraftName"))
	return rv
}


// Creates a copy of the receiving document in response to the user choosing Duplicate from the File menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate(_:)
func (d_ Document) DuplicateDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("duplicateDocument:"), sender)
}


// Creates a new document whose contents are the same as the current document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate(withDelegate:didDuplicate:contextInfo:)
func (d_ Document) DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objectivec.IObject, didDuplicateSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("duplicateDocumentWithDelegate:didDuplicateSelector:contextInfo:"), delegate, didDuplicateSelector, contextInfo)
}


// Locks the document in response to the user choosing the Lock menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lock(_:)
func (d_ Document) LockDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("lockDocument:"), sender)
}


// Prevents the user from making changes to the document’s file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lock(completionHandler:)-161qv
func (d_ Document) LockWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("lockWithCompletionHandler:"), completionHandler)
}


// Prevents the user from making further changes to the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lock(completionHandler:)-6zuhh
func (d_ Document) LockDocumentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("lockDocumentWithCompletionHandler:"), completionHandler)
}


// Waits for any work scheduled by previous invocations of this method to complete, then invokes the passed-in block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/performActivity(withSynchronousWaiting:using:)
func (d_ Document) PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("performActivityWithSynchronousWaiting:usingBlock:"), waitSynchronously, block)
}


// Adds document-specific content to the Page Layout panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/preparePageLayout(_:)
func (d_ Document) PreparePageLayout(pageLayout IPageLayout) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("preparePageLayout:"), pageLayout)
	return rv
}


// Tells the document to customize the specified Save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/prepareSavePanel(_:)
func (d_ Document) PrepareSavePanel(savePanel SavePanel) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("prepareSavePanel:"), savePanel)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidMove(to:)
func (d_ Document) PresentedItemDidMoveToURL(newURL foundation.URL) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidMoveToURL:"), newURL)
}


// Sets the contents of this document by reading from a file or file package, of a specified type, located by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/read(from:ofType:)-1vttv
func (d_ Document) ReadFromURLOfTypeError(url foundation.URL, typeName string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("readFromURL:ofType:error:"), url, objc.String(typeName), outError)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/relinquishPresentedItem(toReader:)
func (d_ Document) RelinquishPresentedItemToReader(reader unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("relinquishPresentedItemToReader:"), reader)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/relinquishPresentedItem(toWriter:)
func (d_ Document) RelinquishPresentedItemToWriter(writer unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("relinquishPresentedItemToWriter:"), writer)
}


// Removes the specified window controller from the receiver’s array of window controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/removeWindowController(_:)
func (d_ Document) RemoveWindowController(windowController IWindowController) {
	objc.Send[objc.ID](d_.ID, objc.Sel("removeWindowController:"), windowController)
}


// Renames the current document in response to the user choosing the Rename menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/rename(_:)
func (d_ Document) RenameDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("renameDocument:"), sender)
}


// The action of the File menu item Revert in a document-based app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/revertToSaved(_:)
func (d_ Document) RevertDocumentToSaved(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("revertDocumentToSaved:"), sender)
}


// Runs the modal page layout panel with the receiver’s printing information object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runModalPageLayout(with:delegate:didRun:contextInfo:)
func (d_ Document) RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo(printInfo IPrintInfo, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runModalPageLayoutWithPrintInfo:delegate:didRunSelector:contextInfo:"), printInfo, delegate, didRunSelector, contextInfo)
}


// Runs the specified print operation modally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runModalPrintOperation(_:delegate:didRun:contextInfo:)
func (d_ Document) RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation IPrintOperation, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runModalPrintOperation:delegate:didRunSelector:contextInfo:"), printOperation, delegate, didRunSelector, contextInfo)
}


// The action method invoked in the receiver as first responder when the user chooses the Page Setup menu command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runPageLayout(_:)
func (d_ Document) RunPageLayout(sender objectivec.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runPageLayout:"), sender)
}


// Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation, and invokes the passed-in completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/save(to:ofType:for:completionHandler:)
func (d_ Document) SaveToURLOfTypeForSaveOperationCompletionHandler(url foundation.URL, typeName string, saveOperation unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveToURL:ofType:forSaveOperation:completionHandler:"), url, objc.String(typeName), saveOperation, completionHandler)
}


// Sets the window outlet of this document to the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/setWindow(_:)
func (d_ Document) SetWindow(window IWindow) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWindow:"), window)
}


// Dismiss the Versions browser for the current document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/stopBrowsingVersions(completionHandler:)
func (d_ Document) StopBrowsingVersionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("stopBrowsingVersionsWithCompletionHandler:"), completionHandler)
}


// Allows the user to make modifications to the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/unlock(completionHandler:)-8p8zd
func (d_ Document) UnlockDocumentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("unlockDocumentWithCompletionHandler:"), completionHandler)
}


// Updates the receiver’s change count according to the given change type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/updateChangeCount(_:)
func (d_ Document) UpdateChangeCount(change unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("updateChangeCount:"), change)
}


// Validates the specified user interface item that the receiver manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/validateUserInterfaceItem(_:)
func (d_ Document) ValidateUserInterfaceItem(item objectivec.IObject) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}


// A Boolean value that indicates whether you can cancel an in-progress autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavingIsImplicitlyCancellable
func (d_ Document) AutosavingIsImplicitlyCancellable() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("autosavingIsImplicitlyCancellable"))
	return rv
}


// The URL for the document’s backup file that was created during an autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/backupFileURL
func (d_ Document) BackupFileURL() foundation.URL {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("backupFileURL"))
	return rv
}


// A Boolean value that indicates whether the user chose to hide the document’s filename extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileNameExtensionWasHiddenInLastRunSavePanel
func (d_ Document) FileNameExtensionWasHiddenInLastRunSavePanel() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileNameExtensionWasHiddenInLastRunSavePanel"))
	return rv
}


// The file type that was last selected in the Save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileTypeFromLastRunSavePanel
func (d_ Document) FileTypeFromLastRunSavePanel() string {
	rv := objc.Send[string](d_.ID, objc.Sel("fileTypeFromLastRunSavePanel"))
	return rv
}


// The location of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileURL
func (d_ Document) FileURL() foundation.URL {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("fileURL"))
	return rv
}


// The location of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileURL
func (d_ Document) SetFileURL(value foundation.URL) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileURL:"), value)
}


// A Boolean value that indicates whether the document has changes that have not been autosaved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/hasUnautosavedChanges
func (d_ Document) HasUnautosavedChanges() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("hasUnautosavedChanges"))
	return rv
}


// A Boolean value that indicates whether the document is a draft that the user has not yet saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isDraft
func (d_ Document) Draft() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("draft"))
	return rv
}


// A Boolean value that indicates whether the document is a draft that the user has not yet saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isDraft
func (d_ Document) SetDraft(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDraft:"), value)
}


// A Boolean value that indicates whether or not the file can be written to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isLocked
func (d_ Document) Locked() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("locked"))
	return rv
}


// A Boolean value that indicates whether the document archives previously saved versions of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/keepBackupFile
func (d_ Document) KeepBackupFile() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("keepBackupFile"))
	return rv
}


// The printing information associated with the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/printInfo
func (d_ Document) PrintInfo() IPrintInfo {
	rv := objc.Send[PrintInfo](d_.ID, objc.Sel("printInfo"))
	return rv
}


// The printing information associated with the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/printInfo
func (d_ Document) SetPrintInfo(value IPrintInfo) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPrintInfo:"), value)
}


// Returns the types of data the receiver can read natively and any types filterable to that native type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/readableTypes
func (d_ Document) ReadableTypes() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("readableTypes"))
	return rv
}


// Returns an array of key paths that represent the restorable attributes of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/restorableStateKeyPaths
func (d_ Document) RestorableStateKeyPaths() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("restorableStateKeyPaths"))
	return rv
}


// Returns whether the document object stores its contents in the user’s iCloud document storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/usesUbiquitousStorage
func (d_ Document) UsesUbiquitousStorage() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("usesUbiquitousStorage"))
	return rv
}


// The name of the document’s sole nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowNibName
func (d_ Document) WindowNibName() NibName {
	rv := objc.Send[NibName](d_.ID, objc.Sel("windowNibName"))
	return rv
}


// A Boolean value that indicates whether the document is shareable from the standard Share menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/allowsdocumentsharing
func (d_ Document) AllowsDocumentSharing() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("allowsDocumentSharing"))
	return rv
}


// A Boolean value that indicates whether the document is shareable from the standard Share menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/allowsdocumentsharing
func (d_ Document) SetAllowsDocumentSharing(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAllowsDocumentSharing:"), value)
}


// The location of the most recently autosaved document contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/autosavedcontentsfileurl
func (d_ Document) AutosavedContentsFileURL() foundation.URL {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("autosavedContentsFileURL"))
	return rv
}


// The location of the most recently autosaved document contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/autosavedcontentsfileurl
func (d_ Document) SetAutosavedContentsFileURL(value foundation.URL) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavedContentsFileURL:"), value)
}


// The document type to use for an autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/autosavingfiletype
func (d_ Document) AutosavingFileType() string {
	rv := objc.Send[string](d_.ID, objc.Sel("autosavingFileType"))
	return rv
}


// The document type to use for an autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/autosavingfiletype
func (d_ Document) SetAutosavingFileType(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavingFileType:"), objc.String(value))
}


// The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/displayname
func (d_ Document) DisplayName() string {
	rv := objc.Send[string](d_.ID, objc.Sel("displayName"))
	return rv
}


// The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/displayname
func (d_ Document) SetDisplayName(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}


// The last-known modification date of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filemodificationdate
func (d_ Document) FileModificationDate() foundation.Date {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("fileModificationDate"))
	return rv
}


// The last-known modification date of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filemodificationdate
func (d_ Document) SetFileModificationDate(value foundation.Date) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileModificationDate:"), value)
}


// The name of the document type, as specified in the app’s information property-list file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filetype
func (d_ Document) FileType() string {
	rv := objc.Send[string](d_.ID, objc.Sel("fileType"))
	return rv
}


// The name of the document type, as specified in the app’s information property-list file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filetype
func (d_ Document) SetFileType(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileType:"), objc.String(value))
}


// A Boolean value that indicates whether the document owns an undo manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/hasundomanager
func (d_ Document) HasUndoManager() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("hasUndoManager"))
	return rv
}


// A Boolean value that indicates whether the document owns an undo manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/hasundomanager
func (d_ Document) SetHasUndoManager(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHasUndoManager:"), value)
}


// A Boolean value that indicates whether the document is currently displaying the Versions browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isbrowsingversions
func (d_ Document) IsBrowsingVersions() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isBrowsingVersions"))
	return rv
}


// A Boolean value that indicates whether the document is currently displaying the Versions browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isbrowsingversions
func (d_ Document) SetIsBrowsingVersions(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsBrowsingVersions:"), value)
}


// A Boolean value that indicates whether the document has unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isdocumentedited
func (d_ Document) IsDocumentEdited() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isDocumentEdited"))
	return rv
}


// A Boolean value that indicates whether the document has unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isdocumentedited
func (d_ Document) SetIsDocumentEdited(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsDocumentEdited:"), value)
}


// A Boolean value that indicates whether the document is a draft that the user has not yet saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isdraft
func (d_ Document) IsDraft() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isDraft"))
	return rv
}


// A Boolean value that indicates whether the document is a draft that the user has not yet saved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isdraft
func (d_ Document) SetIsDraft(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsDraft:"), value)
}


// A Boolean value that indicates whether the document’s file is completely loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isentirefileloaded
func (d_ Document) IsEntireFileLoaded() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isEntireFileLoaded"))
	return rv
}


// A Boolean value that indicates whether the document’s file is completely loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isentirefileloaded
func (d_ Document) SetIsEntireFileLoaded(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsEntireFileLoaded:"), value)
}


// A Boolean value that indicates whether the document is in read-only mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isinviewingmode
func (d_ Document) IsInViewingMode() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isInViewingMode"))
	return rv
}


// A Boolean value that indicates whether the document is in read-only mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/isinviewingmode
func (d_ Document) SetIsInViewingMode(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsInViewingMode:"), value)
}


// A Boolean value that indicates whether or not the file can be written to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/islocked
func (d_ Document) IsLocked() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("isLocked"))
	return rv
}


// A Boolean value that indicates whether or not the file can be written to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/islocked
func (d_ Document) SetIsLocked(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setIsLocked:"), value)
}


// The name of the document seen by the user in AppleScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/lastcomponentoffilename
func (d_ Document) LastComponentOfFileName() string {
	rv := objc.Send[string](d_.ID, objc.Sel("lastComponentOfFileName"))
	return rv
}


// The name of the document seen by the user in AppleScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/lastcomponentoffilename
func (d_ Document) SetLastComponentOfFileName(value string) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLastComponentOfFileName:"), objc.String(value))
}


// Returns the object specifier that represents the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/objectspecifier
func (d_ Document) ObjectSpecifier() foundation.ScriptObjectSpecifier {
	rv := objc.Send[foundation.ScriptObjectSpecifier](d_.ID, objc.Sel("objectSpecifier"))
	return rv
}


// Returns the object specifier that represents the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/objectspecifier
func (d_ Document) SetObjectSpecifier(value foundation.ScriptObjectSpecifier) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setObjectSpecifier:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/observedpresenteditemubiquityattributes
func (d_ Document) ObservedPresentedItemUbiquityAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("observedPresentedItemUbiquityAttributes"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/observedpresenteditemubiquityattributes
func (d_ Document) SetObservedPresentedItemUbiquityAttributes(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setObservedPresentedItemUbiquityAttributes:"), value)
}


// A print operation you can use to create a PDF representation of the document’s current contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/pdfprintoperation
func (d_ Document) PdfPrintOperation() IPrintOperation {
	rv := objc.Send[PrintOperation](d_.ID, objc.Sel("pdfPrintOperation"))
	return rv
}


// A print operation you can use to create a PDF representation of the document’s current contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/pdfprintoperation
func (d_ Document) SetPdfPrintOperation(value IPrintOperation) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPdfPrintOperation:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/presenteditemurl
func (d_ Document) PresentedItemURL() foundation.URL {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("presentedItemURL"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/presenteditemurl
func (d_ Document) SetPresentedItemURL(value foundation.URL) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPresentedItemURL:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/previewrepresentableactivityitems
func (d_ Document) PreviewRepresentableActivityItems() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("previewRepresentableActivityItems"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/previewrepresentableactivityitems
func (d_ Document) SetPreviewRepresentableActivityItems(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreviewRepresentableActivityItems:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/savepanelshowsfileformatscontrol
func (d_ Document) SavePanelShowsFileFormatsControl() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("savePanelShowsFileFormatsControl"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/savepanelshowsfileformatscontrol
func (d_ Document) SetSavePanelShowsFileFormatsControl(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setSavePanelShowsFileFormatsControl:"), value)
}


// A Boolean value that indicates whether the document’s Save panel displays a list of supported writable document types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/shouldrunsavepanelwithaccessoryview
func (d_ Document) ShouldRunSavePanelWithAccessoryView() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldRunSavePanelWithAccessoryView"))
	return rv
}


// A Boolean value that indicates whether the document’s Save panel displays a list of supported writable document types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/shouldrunsavepanelwithaccessoryview
func (d_ Document) SetShouldRunSavePanelWithAccessoryView(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setShouldRunSavePanelWithAccessoryView:"), value)
}


// The object that the document uses to support undo/redo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/undomanager
func (d_ Document) UndoManager() foundation.UndoManager {
	rv := objc.Send[foundation.UndoManager](d_.ID, objc.Sel("undoManager"))
	return rv
}


// The object that the document uses to support undo/redo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/undomanager
func (d_ Document) SetUndoManager(value foundation.UndoManager) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUndoManager:"), value)
}


// An object that encapsulates a user activity the document supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/useractivity
func (d_ Document) UserActivity() foundation.UserActivity {
	rv := objc.Send[foundation.UserActivity](d_.ID, objc.Sel("userActivity"))
	return rv
}


// An object that encapsulates a user activity the document supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/useractivity
func (d_ Document) SetUserActivity(value foundation.UserActivity) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUserActivity:"), value)
}


// The document’s current window controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/windowcontrollers
func (d_ Document) WindowControllers() IWindowController {
	rv := objc.Send[WindowController](d_.ID, objc.Sel("windowControllers"))
	return rv
}


// The document’s current window controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/windowcontrollers
func (d_ Document) SetWindowControllers(value IWindowController) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWindowControllers:"), value)
}


// Returns the document window to use as the parent of a document-modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/windowforsheet
func (d_ Document) WindowForSheet() IWindow {
	rv := objc.Send[Window](d_.ID, objc.Sel("windowForSheet"))
	return rv
}


// Returns the document window to use as the parent of a document-modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/windowforsheet
func (d_ Document) SetWindowForSheet(value IWindow) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWindowForSheet:"), value)
}


// The key that identifies the document associated with a user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuseractivitydocumenturlkey
func (d_ Document) NSUserActivityDocumentURLKey() string {
	rv := objc.Send[string](d_.ID, objc.Sel("NSUserActivityDocumentURLKey"))
	return rv
}



