// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
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
	// properties:
	AllowsDocumentSharing() bool
	AutosavedContentsFileURL() objc.IObject /* cross-framework: NSURL */
	SetAutosavedContentsFileURL(value objc.IObject /* cross-framework: NSURL */)
	AutosavingFileType() objc.IObject /* cross-framework: NSString */
	AutosavingIsImplicitlyCancellable() bool
	BackupFileURL() objc.IObject /* cross-framework: NSURL */
	DisplayName() objc.IObject /* cross-framework: NSString */
	SetDisplayName(value objc.IObject /* cross-framework: NSString */)
	FileModificationDate() objc.IObject /* cross-framework: NSDate */
	SetFileModificationDate(value objc.IObject /* cross-framework: NSDate */)
	FileNameExtensionWasHiddenInLastRunSavePanel() bool
	FileType() objc.IObject /* cross-framework: NSString */
	SetFileType(value objc.IObject /* cross-framework: NSString */)
	FileTypeFromLastRunSavePanel() objc.IObject /* cross-framework: NSString */
	FileURL() objc.IObject /* cross-framework: NSURL */
	SetFileURL(value objc.IObject /* cross-framework: NSURL */)
	HasUnautosavedChanges() bool
	HasUndoManager() bool
	SetHasUndoManager(value bool)
	BrowsingVersions() bool
	DocumentEdited() bool
	Draft() bool
	SetDraft(value bool)
	EntireFileLoaded() bool
	InViewingMode() bool
	Locked() bool
	KeepBackupFile() bool
	LastComponentOfFileName() objc.IObject /* cross-framework: NSString */
	SetLastComponentOfFileName(value objc.IObject /* cross-framework: NSString */)
	ObjectSpecifier() foundation.ScriptObjectSpecifier
	ObservedPresentedItemUbiquityAttributes() unsafe.Pointer
	PDFPrintOperation() IPrintOperation
	PresentedItemURL() objc.IObject /* cross-framework: NSURL */
	PreviewRepresentableActivityItems() []objc.ID
	SetPreviewRepresentableActivityItems(value []objc.ID)
	PrintInfo() IPrintInfo
	SetPrintInfo(value IPrintInfo)
	SavePanelShowsFileFormatsControl() bool
	ShouldRunSavePanelWithAccessoryView() bool
	UndoManager() foundation.UndoManager
	SetUndoManager(value foundation.UndoManager)
	UserActivity() foundation.UserActivity
	SetUserActivity(value foundation.UserActivity)
	WindowControllers() []WindowController
	WindowForSheet() IWindow
	WindowNibName() objc.IObject /* cross-framework: NibName */
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
	NSUserActivityDocumentURLKey() objc.IObject /* cross-framework: NSString */
	// methods:
	AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler unsafe.Pointer)
	AddWindowController(windowController IWindowController)
	AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objc.IObject, didAutosaveSelector objc.SEL, contextInfo unsafe.Pointer)
	AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler unsafe.Pointer)
	BrowseDocumentVersions(sender objc.IObject)
	CanAsynchronouslyWriteToURLOfTypeForSaveOperation(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType) bool
	CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo(delegate objc.IObject, shouldCloseSelector objc.SEL, contextInfo unsafe.Pointer)
	ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.ID
	CheckAutosavingSafetyAndReturnError(outError unsafe.Pointer) bool
	Close()
	ContinueActivityUsingBlock(block unsafe.Pointer)
	ContinueAsynchronousWorkOnMainThreadUsingBlock(block unsafe.Pointer)
	DataOfTypeError(typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) foundation.Data
	DefaultDraftName() foundation.String
	DuplicateAndReturnError(outError unsafe.Pointer) IDocument
	DuplicateDocument(sender objc.IObject)
	DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objc.IObject, didDuplicateSelector objc.SEL, contextInfo unsafe.Pointer)
	EncodeRestorableStateWithCoder(coder foundation.Coder)
	EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.Coder, queue foundation.OperationQueue)
	FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType, absoluteOriginalContentsURL objc.IObject /* cross-framework: NSURL */, outError unsafe.Pointer) foundation.IDictionary
	FileNameExtensionForTypeSaveOperation(typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType) foundation.String
	FileWrapperOfTypeError(typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) foundation.FileWrapper
	HandleCloseScriptCommand(command foundation.CloseCommand) objc.ID
	HandlePrintScriptCommand(command foundation.ScriptCommand) objc.ID
	HandleSaveScriptCommand(command foundation.ScriptCommand) objc.ID
	InvalidateRestorableState()
	LockDocument(sender objc.IObject)
	LockWithCompletionHandler(completionHandler unsafe.Pointer)
	LockDocumentWithCompletionHandler(completionHandler unsafe.Pointer)
	MakeWindowControllers()
	MoveDocument(sender objc.IObject)
	MoveDocumentWithCompletionHandler(completionHandler unsafe.Pointer)
	MoveToURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer)
	MoveDocumentToUbiquityContainer(sender objc.IObject)
	PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block unsafe.Pointer)
	PerformAsynchronousFileAccessUsingBlock(block unsafe.Pointer)
	PerformSynchronousFileAccessUsingBlock(block unsafe.Pointer)
	PrepareSharingServicePicker(sharingServicePicker ISharingServicePicker)
	PreparePageLayout(pageLayout IPageLayout) bool
	PrepareSavePanel(savePanel ISavePanel) bool
	PresentError(error_ objc.IObject /* cross-framework: Error */) bool
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ objc.IObject /* cross-framework: Error */, window IWindow, delegate objc.IObject, didPresentSelector objc.SEL, contextInfo unsafe.Pointer)
	PresentedItemDidChange()
	PresentedItemDidChangeUbiquityAttributes(attributes unsafe.Pointer)
	PresentedItemDidGainVersion(version foundation.FileVersion)
	PresentedItemDidLoseVersion(version foundation.FileVersion)
	PresentedItemDidMoveToURL(newURL objc.IObject /* cross-framework: NSURL */)
	PresentedItemDidResolveConflictVersion(version foundation.FileVersion)
	PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo(printSettings foundation.IDictionary, showPrintPanel bool, delegate objc.IObject, didPrintSelector objc.SEL, contextInfo unsafe.Pointer)
	PrintDocument(sender objc.IObject)
	PrintOperationWithSettingsError(printSettings foundation.IDictionary, outError unsafe.Pointer) IPrintOperation
	ReadFromURLOfTypeError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool
	ReadFromFileWrapperOfTypeError(fileWrapper foundation.FileWrapper, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool
	ReadFromDataOfTypeError(data objc.IObject /* cross-framework: NSData */, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool
	RelinquishPresentedItemToReader(reader unsafe.Pointer)
	RelinquishPresentedItemToWriter(writer unsafe.Pointer)
	RemoveWindowController(windowController IWindowController)
	RenameDocument(sender objc.IObject)
	RestoreStateWithCoder(coder foundation.Coder)
	RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */, state foundation.Coder, completionHandler unsafe.Pointer)
	RevertToContentsOfURLOfTypeError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool
	RevertDocumentToSaved(sender objc.IObject)
	RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo(printInfo IPrintInfo, delegate objc.IObject, didRunSelector objc.SEL, contextInfo unsafe.Pointer)
	RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation IPrintOperation, delegate objc.IObject, didRunSelector objc.SEL, contextInfo unsafe.Pointer)
	RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo(saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.SEL, contextInfo unsafe.Pointer)
	RunPageLayout(sender objc.IObject)
	SaveDocument(sender objc.IObject)
	SaveToURLOfTypeForSaveOperationCompletionHandler(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType, completionHandler unsafe.Pointer)
	SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.SEL, contextInfo unsafe.Pointer)
	SaveDocumentWithDelegateDidSaveSelectorContextInfo(delegate objc.IObject, didSaveSelector objc.SEL, contextInfo unsafe.Pointer)
	SaveDocumentAs(sender objc.IObject)
	SavePresentedItemChangesWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveDocumentTo(sender objc.IObject)
	SaveDocumentToPDF(sender objc.IObject)
	ScheduleAutosaving()
	SetWindow(window IWindow)
	ShareDocumentWithSharingServiceCompletionHandler(sharingService ISharingService, completionHandler unsafe.Pointer)
	ShouldChangePrintInfo(newPrintInfo IPrintInfo) bool
	ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo(windowController IWindowController, delegate objc.IObject, shouldCloseSelector objc.SEL, contextInfo unsafe.Pointer)
	ShowWindows()
	StopBrowsingVersionsWithCompletionHandler(completionHandler unsafe.Pointer)
	UnblockUserInteraction()
	UnlockDocument(sender objc.IObject)
	UnlockWithCompletionHandler(completionHandler unsafe.Pointer)
	UnlockDocumentWithCompletionHandler(completionHandler unsafe.Pointer)
	UpdateChangeCount(change DocumentChangeType)
	UpdateChangeCountWithTokenForSaveOperation(changeCountToken objc.IObject, saveOperation SaveOperationType)
	UpdateUserActivityState(activity foundation.UserActivity)
	ValidateUserInterfaceItem(item objc.IObject) bool
	WillNotPresentError(error_ objc.IObject /* cross-framework: Error */)
	WillPresentError(error_ objc.IObject /* cross-framework: Error */) objc.IObject /* cross-framework: Error */
	WindowControllerDidLoadNib(windowController IWindowController)
	WindowControllerWillLoadNib(windowController IWindowController)
	WritableTypesForSaveOperation(saveOperation SaveOperationType) []string
	WriteToURLOfTypeError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool
	WriteToURLOfTypeForSaveOperationOriginalContentsURLError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType, absoluteOriginalContentsURL objc.IObject /* cross-framework: NSURL */, outError unsafe.Pointer) bool
	WriteSafelyToURLOfTypeForSaveOperationError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType, outError unsafe.Pointer) bool
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



// Initializes a document with the specified contents, and places the resulting document’s file at the designated location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/init(for:withContentsOf:ofType:)
func NewDocumentForURLWithContentsOfURLOfTypeError(urlOrNil objc.IObject /* cross-framework: NSURL */, contentsURL objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) Document {
	instance := getDocumentClass().Alloc()
	rv := objc.Send[Document](instance.ID, objc.Sel("initForURL:withContentsOfURL:ofType:error:"), urlOrNil, contentsURL, typeName, outError)
	rv.Autorelease()
	return rv
}


// Initializes and returns a document object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/initWithContentsOfFile:ofType:
func NewDocumentWithContentsOfFileOfType(absolutePath objc.IObject /* cross-framework: NSString */, typeName objc.IObject /* cross-framework: NSString */) Document {
	instance := getDocumentClass().Alloc()
	rv := objc.Send[Document](instance.ID, objc.Sel("initWithContentsOfFile:ofType:"), absolutePath, typeName)
	rv.Autorelease()
	return rv
}


// Initializes and returns a document object of a given type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/initWithContentsOfURL:ofType:
func NewDocumentWithContentsOfURLOfType(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */) Document {
	instance := getDocumentClass().Alloc()
	rv := objc.Send[Document](instance.ID, objc.Sel("initWithContentsOfURL:ofType:"), url, typeName)
	rv.Autorelease()
	return rv
}


// Initializes a document located by a URL of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/init(contentsOf:ofType:)
func NewDocumentWithContentsOfURLOfTypeError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) Document {
	instance := getDocumentClass().Alloc()
	rv := objc.Send[Document](instance.ID, objc.Sel("initWithContentsOfURL:ofType:error:"), url, typeName, outError)
	rv.Autorelease()
	return rv
}


// Initializes a document of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/init(type:)
func NewDocumentWithTypeError(typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) Document {
	instance := getDocumentClass().Alloc()
	rv := objc.Send[Document](instance.ID, objc.Sel("initWithType:error:"), typeName, outError)
	rv.Autorelease()
	return rv
}



// Returns the classes that support secure coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/allowedClasses(forRestorableStateKeyPath:)
func (dc _DocumentClass) AllowedClassesForRestorableStateKeyPath(keyPath objc.IObject /* cross-framework: NSString */) []objc.Class {
	rv := objc.Send[[]objc.Class](objc.ID(dc.class), objc.Sel("allowedClassesForRestorableStateKeyPath:"), keyPath)
	return rv
}


// Returns a Boolean value that indicates whether the receiver reads multiple documents of the given type concurrently.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/canConcurrentlyReadDocuments(ofType:)
func (dc _DocumentClass) CanConcurrentlyReadDocumentsOfType(typeName objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("canConcurrentlyReadDocumentsOfType:"), typeName)
	return rv
}


// Returns a Boolean value that indicates whether the document can read and write the data natively.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isNativeType(_:)
func (dc _DocumentClass) IsNativeType(type_ objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("isNativeType:"), type_)
	return rv
}


// A Boolean value that indicates whether the document subclass supports autosaving of drafts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavesDrafts
func (dc _DocumentClass) AutosavesDrafts() bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("autosavesDrafts"))
	return rv
}

// A Boolean value that indicates whether the document subclass supports autosaving in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavesInPlace
func (dc _DocumentClass) AutosavesInPlace() bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("autosavesInPlace"))
	return rv
}

// A Boolean value that indicates whether the document subclass supports version management.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/preservesVersions
func (dc _DocumentClass) PreservesVersions() bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("preservesVersions"))
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

// Returns the types of data the receiver can write natively and any types filterable to that native type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/writableTypes
func (dc _DocumentClass) WritableTypes() []string {
	rv := objc.Send[[]string](objc.ID(dc.class), objc.Sel("writableTypes"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/accommodatePresentedItemDeletion(completionHandler:)
func (d_ Document) AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("accommodatePresentedItemDeletionWithCompletionHandler:"), completionHandler)
}


// Adds the specified window controller to the current document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/addWindowController(_:)
func (d_ Document) AddWindowController(windowController IWindowController) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addWindowController:"), windowController)
}


// Autosaves the document’s contents to an appropriate location in the file system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosave(withDelegate:didAutosave:contextInfo:)
func (d_ Document) AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objc.IObject, didAutosaveSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("autosaveDocumentWithDelegate:didAutosaveSelector:contextInfo:"), delegate, didAutosaveSelector, contextInfo)
}


// Autosaves the document’s contents to an appropriate file-system location, as needed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosave(withImplicitCancellability:completionHandler:)
func (d_ Document) AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("autosaveWithImplicitCancellability:completionHandler:"), autosavingIsImplicitlyCancellable, completionHandler)
}


// Opens the Versions browser in the document’s main window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/browseVersions(_:)
func (d_ Document) BrowseDocumentVersions(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("browseDocumentVersions:"), sender)
}


// Returns whether the receiver can concurrently write to a file or file package located by a URL, that is formatted for a specific type, for a specific kind of save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/canAsynchronouslyWrite(to:ofType:for:)
func (d_ Document) CanAsynchronouslyWriteToURLOfTypeForSaveOperation(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("canAsynchronouslyWriteToURL:ofType:forSaveOperation:"), url, typeName, saveOperation)
	return rv
}


// Determines whether to close the document, prompting the user as needed to choose a course of action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/canClose(withDelegate:shouldClose:contextInfo:)
func (d_ Document) CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo(delegate objc.IObject, shouldCloseSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("canCloseDocumentWithDelegate:shouldCloseSelector:contextInfo:"), delegate, shouldCloseSelector, contextInfo)
}


// Returns an object that encapsulates the current record of document changes at the beginning of a save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/changeCountToken(for:)
func (d_ Document) ChangeCountTokenForSaveOperation(saveOperation SaveOperationType) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("changeCountTokenForSaveOperation:"), saveOperation)
	return rv
}


// Returns a Boolean value that indicates whether it is safe to autosave document changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/checkAutosavingSafety()
func (d_ Document) CheckAutosavingSafetyAndReturnError(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("checkAutosavingSafetyAndReturnError:"), outError)
	return rv
}


// Closes all of the document’s windows and removes the document from its document controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/close()
func (d_ Document) Close() {
	objc.Send[objc.ID](d_.ID, objc.Sel("close"))
}


// Continues to perform the task for a user activity object using a different block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/continueActivity(_:)
func (d_ Document) ContinueActivityUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("continueActivityUsingBlock:"), block)
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
func (d_ Document) DataOfTypeError(typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) foundation.Data {
	rv := objc.Send[foundation.Data](d_.ID, objc.Sel("dataOfType:error:"), typeName, outError)
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


// Creates a new document whose contents are the same as the receiver and returns an error object if unsuccessful.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate()
func (d_ Document) DuplicateAndReturnError(outError unsafe.Pointer) IDocument {
	rv := objc.Send[Document](d_.ID, objc.Sel("duplicateAndReturnError:"), outError)
	return rv
}


// Creates a copy of the receiving document in response to the user choosing Duplicate from the File menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate(_:)
func (d_ Document) DuplicateDocument(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("duplicateDocument:"), sender)
}


// Creates a new document whose contents are the same as the current document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate(withDelegate:didDuplicate:contextInfo:)
func (d_ Document) DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objc.IObject, didDuplicateSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("duplicateDocumentWithDelegate:didDuplicateSelector:contextInfo:"), delegate, didDuplicateSelector, contextInfo)
}


// Saves the interface-related state of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/encodeRestorableState(with:)
func (d_ Document) EncodeRestorableStateWithCoder(coder foundation.Coder) {
	objc.Send[objc.ID](d_.ID, objc.Sel("encodeRestorableStateWithCoder:"), coder)
}


// Saves the interface-related state of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/encodeRestorableState(with:backgroundQueue:)
func (d_ Document) EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.Coder, queue foundation.OperationQueue) {
	objc.Send[objc.ID](d_.ID, objc.Sel("encodeRestorableStateWithCoder:backgroundQueue:"), coder, queue)
}


// Returns the attributes to write to the file or file package at the specified URL, and targeting the specified type of save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileAttributesToWrite(to:ofType:for:originalContentsURL:)
func (d_ Document) FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType, absoluteOriginalContentsURL objc.IObject /* cross-framework: NSURL */, outError unsafe.Pointer) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](d_.ID, objc.Sel("fileAttributesToWriteToURL:ofType:forSaveOperation:originalContentsURL:error:"), url, typeName, saveOperation, absoluteOriginalContentsURL, outError)
	return rv
}


// Returns a filename extension that can be appended to a base filename, for a specified file type and kind of save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileNameExtension(forType:saveOperation:)
func (d_ Document) FileNameExtensionForTypeSaveOperation(typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType) foundation.String {
	rv := objc.Send[foundation.String](d_.ID, objc.Sel("fileNameExtensionForType:saveOperation:"), typeName, saveOperation)
	return rv
}


// Creates and returns a file wrapper that contains the contents of the document, formatted to the specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileWrapper(ofType:)
func (d_ Document) FileWrapperOfTypeError(typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) foundation.FileWrapper {
	rv := objc.Send[foundation.FileWrapper](d_.ID, objc.Sel("fileWrapperOfType:error:"), typeName, outError)
	return rv
}


// Handles the Close AppleScript command by attempting to close the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/handleClose(_:)
func (d_ Document) HandleCloseScriptCommand(command foundation.CloseCommand) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("handleCloseScriptCommand:"), command)
	return rv
}


// Handles the Print AppleScript command by attempting to print the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/handlePrint(_:)
func (d_ Document) HandlePrintScriptCommand(command foundation.ScriptCommand) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("handlePrintScriptCommand:"), command)
	return rv
}


// Handles the Save AppleScript command by attempting to save the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/handleSave(_:)
func (d_ Document) HandleSaveScriptCommand(command foundation.ScriptCommand) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("handleSaveScriptCommand:"), command)
	return rv
}


// Marks the document’s interface-related state as dirty.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/invalidateRestorableState()
func (d_ Document) InvalidateRestorableState() {
	objc.Send[objc.ID](d_.ID, objc.Sel("invalidateRestorableState"))
}


// Locks the document in response to the user choosing the Lock menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lock(_:)
func (d_ Document) LockDocument(sender objc.IObject) {
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


// Creates the window controller objects that the document uses to display its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/makeWindowControllers()
func (d_ Document) MakeWindowControllers() {
	objc.Send[objc.ID](d_.ID, objc.Sel("makeWindowControllers"))
}


// Moves the document to a new location in response to the user choosing the Move To… menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/move(_:)
func (d_ Document) MoveDocument(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveDocument:"), sender)
}


// Moves the document to a user-selected location.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/move(completionHandler:)
func (d_ Document) MoveDocumentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveDocumentWithCompletionHandler:"), completionHandler)
}


// Moves the document’s file to the given URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/move(to:completionHandler:)
func (d_ Document) MoveToURLCompletionHandler(url objc.IObject /* cross-framework: NSURL */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveToURL:completionHandler:"), url, completionHandler)
}


// Moves the document to the user’s iCloud storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/moveToUbiquityContainer(_:)
func (d_ Document) MoveDocumentToUbiquityContainer(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveDocumentToUbiquityContainer:"), sender)
}


// Waits for any work scheduled by previous invocations of this method to complete, then invokes the passed-in block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/performActivity(withSynchronousWaiting:using:)
func (d_ Document) PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("performActivityWithSynchronousWaiting:usingBlock:"), waitSynchronously, block)
}


// Waits for any scheduled file access to complete but without blocking the main thread, then invokes the passed-in block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/performAsynchronousFileAccess(_:)
func (d_ Document) PerformAsynchronousFileAccessUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("performAsynchronousFileAccessUsingBlock:"), block)
}


// Waits for any scheduled file access to complete, then invokes the passed-in block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/performSynchronousFileAccess(_:)
func (d_ Document) PerformSynchronousFileAccessUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("performSynchronousFileAccessUsingBlock:"), block)
}


// Perform any custom setup associated with a sharing service picker.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/prepare(_:)
func (d_ Document) PrepareSharingServicePicker(sharingServicePicker ISharingServicePicker) {
	objc.Send[objc.ID](d_.ID, objc.Sel("prepareSharingServicePicker:"), sharingServicePicker)
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
func (d_ Document) PrepareSavePanel(savePanel ISavePanel) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("prepareSavePanel:"), savePanel)
	return rv
}


// Presents an error alert to the user as a modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentError(_:)
func (d_ Document) PresentError(error_ objc.IObject /* cross-framework: Error */) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("presentError:"), error_)
	return rv
}


// Presents an error alert to the user as a modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (d_ Document) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ objc.IObject /* cross-framework: Error */, window IWindow, delegate objc.IObject, didPresentSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error_, window, delegate, didPresentSelector, contextInfo)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidChange()
func (d_ Document) PresentedItemDidChange() {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidChange"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidChangeUbiquityAttributes(_:)
func (d_ Document) PresentedItemDidChangeUbiquityAttributes(attributes unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidChangeUbiquityAttributes:"), attributes)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidGain(_:)
func (d_ Document) PresentedItemDidGainVersion(version foundation.FileVersion) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidGainVersion:"), version)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidLose(_:)
func (d_ Document) PresentedItemDidLoseVersion(version foundation.FileVersion) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidLoseVersion:"), version)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidMove(to:)
func (d_ Document) PresentedItemDidMoveToURL(newURL objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidMoveToURL:"), newURL)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidResolveConflict(_:)
func (d_ Document) PresentedItemDidResolveConflictVersion(version foundation.FileVersion) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidResolveConflictVersion:"), version)
}


// Prints the document’s contents, optionally displaying a print panel to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/print(withSettings:showPrintPanel:delegate:didPrint:contextInfo:)
func (d_ Document) PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo(printSettings foundation.IDictionary, showPrintPanel bool, delegate objc.IObject, didPrintSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("printDocumentWithSettings:showPrintPanel:delegate:didPrintSelector:contextInfo:"), printSettings, showPrintPanel, delegate, didPrintSelector, contextInfo)
}


// Prints the receiver in response to the user choosing the Print menu command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/printDocument(_:)
func (d_ Document) PrintDocument(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("printDocument:"), sender)
}


// Creates and returns a print operation for the document’s contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/printOperation(withSettings:)
func (d_ Document) PrintOperationWithSettingsError(printSettings foundation.IDictionary, outError unsafe.Pointer) IPrintOperation {
	rv := objc.Send[PrintOperation](d_.ID, objc.Sel("printOperationWithSettings:error:"), printSettings, outError)
	return rv
}


// Sets the contents of this document by reading from a file or file package, of a specified type, located by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/read(from:ofType:)-1vttv
func (d_ Document) ReadFromURLOfTypeError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("readFromURL:ofType:error:"), url, typeName, outError)
	return rv
}


// Sets the contents of this document by reading from a file wrapper of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/read(from:ofType:)-3rzsi
func (d_ Document) ReadFromFileWrapperOfTypeError(fileWrapper foundation.FileWrapper, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("readFromFileWrapper:ofType:error:"), fileWrapper, typeName, outError)
	return rv
}


// Sets the contents of this document by reading from data of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/read(from:ofType:)-6g6ai
func (d_ Document) ReadFromDataOfTypeError(data objc.IObject /* cross-framework: NSData */, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("readFromData:ofType:error:"), data, typeName, outError)
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
func (d_ Document) RenameDocument(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("renameDocument:"), sender)
}


// Restores the interface-related state of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/restoreState(with:)
func (d_ Document) RestoreStateWithCoder(coder foundation.Coder) {
	objc.Send[objc.ID](d_.ID, objc.Sel("restoreStateWithCoder:"), coder)
}


// Restores a window that was associated with a document, after that document is reopened.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/restoreWindow(withIdentifier:state:completionHandler:)
func (d_ Document) RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier objc.IObject /* cross-framework: UserInterfaceItemIdentifier */, state foundation.Coder, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("restoreDocumentWindowWithIdentifier:state:completionHandler:"), identifier, state, completionHandler)
}


// Discards all unsaved document modifications and replaces the document’s contents by reading a file or file package located by a URL of a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/revert(toContentsOf:ofType:)
func (d_ Document) RevertToContentsOfURLOfTypeError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("revertToContentsOfURL:ofType:error:"), url, typeName, outError)
	return rv
}


// The action of the File menu item Revert in a document-based app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/revertToSaved(_:)
func (d_ Document) RevertDocumentToSaved(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("revertDocumentToSaved:"), sender)
}


// Runs the modal page layout panel with the receiver’s printing information object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runModalPageLayout(with:delegate:didRun:contextInfo:)
func (d_ Document) RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo(printInfo IPrintInfo, delegate objc.IObject, didRunSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runModalPageLayoutWithPrintInfo:delegate:didRunSelector:contextInfo:"), printInfo, delegate, didRunSelector, contextInfo)
}


// Runs the specified print operation modally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runModalPrintOperation(_:delegate:didRun:contextInfo:)
func (d_ Document) RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation IPrintOperation, delegate objc.IObject, didRunSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runModalPrintOperation:delegate:didRunSelector:contextInfo:"), printOperation, delegate, didRunSelector, contextInfo)
}


// Presents a modal Save panel to the user, then tries to save the document if the user approves the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runModalSavePanel(for:delegate:didSave:contextInfo:)
func (d_ Document) RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo(saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runModalSavePanelForSaveOperation:delegate:didSaveSelector:contextInfo:"), saveOperation, delegate, didSaveSelector, contextInfo)
}


// The action method invoked in the receiver as first responder when the user chooses the Page Setup menu command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runPageLayout(_:)
func (d_ Document) RunPageLayout(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runPageLayout:"), sender)
}


// The action method invoked in the receiver as first responder when the user chooses the Save menu command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/save(_:)
func (d_ Document) SaveDocument(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveDocument:"), sender)
}


// Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation, and invokes the passed-in completion handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/save(to:ofType:for:completionHandler:)
func (d_ Document) SaveToURLOfTypeForSaveOperationCompletionHandler(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveToURL:ofType:forSaveOperation:completionHandler:"), url, typeName, saveOperation, completionHandler)
}


// Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/save(to:ofType:for:delegate:didSave:contextInfo:)
func (d_ Document) SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType, delegate objc.IObject, didSaveSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveToURL:ofType:forSaveOperation:delegate:didSaveSelector:contextInfo:"), url, typeName, saveOperation, delegate, didSaveSelector, contextInfo)
}


// Saves the document and delivers the results to the provided delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/save(withDelegate:didSave:contextInfo:)
func (d_ Document) SaveDocumentWithDelegateDidSaveSelectorContextInfo(delegate objc.IObject, didSaveSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveDocumentWithDelegate:didSaveSelector:contextInfo:"), delegate, didSaveSelector, contextInfo)
}


// The action method invoked in the receiver as first responder when the user chooses the Save As menu command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/saveAs(_:)
func (d_ Document) SaveDocumentAs(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveDocumentAs:"), sender)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/savePresentedItemChanges(completionHandler:)
func (d_ Document) SavePresentedItemChangesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("savePresentedItemChangesWithCompletionHandler:"), completionHandler)
}


// The action method invoked in the receiver as first responder when the user chooses the Save To menu command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/saveTo(_:)
func (d_ Document) SaveDocumentTo(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveDocumentTo:"), sender)
}


// Exports a PDF representation of the document’s current contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/saveToPDF(_:)
func (d_ Document) SaveDocumentToPDF(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveDocumentToPDF:"), sender)
}


// Schedules periodic autosaving for the purpose of crash protection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/scheduleAutosaving()
func (d_ Document) ScheduleAutosaving() {
	objc.Send[objc.ID](d_.ID, objc.Sel("scheduleAutosaving"))
}


// Sets the window outlet of this document to the specified value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/setWindow(_:)
func (d_ Document) SetWindow(window IWindow) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWindow:"), window)
}


// Share the document’s file using the specified sharing service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/share(with:completionHandler:)
func (d_ Document) ShareDocumentWithSharingServiceCompletionHandler(sharingService ISharingService, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("shareDocumentWithSharingService:completionHandler:"), sharingService, completionHandler)
}


// Returns a Boolean value that indicates whether the document allows changes to the default printing information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/shouldChangePrintInfo(_:)
func (d_ Document) ShouldChangePrintInfo(newPrintInfo IPrintInfo) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldChangePrintInfo:"), newPrintInfo)
	return rv
}


// Determines whether the system should close the document and its associated window.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/shouldCloseWindowController(_:delegate:shouldClose:contextInfo:)
func (d_ Document) ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo(windowController IWindowController, delegate objc.IObject, shouldCloseSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("shouldCloseWindowController:delegate:shouldCloseSelector:contextInfo:"), windowController, delegate, shouldCloseSelector, contextInfo)
}


// Displays all of the document’s windows, bringing them to the front and making them main or key as necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/showWindows()
func (d_ Document) ShowWindows() {
	objc.Send[objc.ID](d_.ID, objc.Sel("showWindows"))
}


// Dismiss the Versions browser for the current document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/stopBrowsingVersions(completionHandler:)
func (d_ Document) StopBrowsingVersionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("stopBrowsingVersionsWithCompletionHandler:"), completionHandler)
}


// Unblocks the main thread during asynchronous saving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/unblockUserInteraction()
func (d_ Document) UnblockUserInteraction() {
	objc.Send[objc.ID](d_.ID, objc.Sel("unblockUserInteraction"))
}


// Unlocks the document in response to the user choosing the Unlock menu item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/unlock(_:)
func (d_ Document) UnlockDocument(sender objc.IObject) {
	objc.Send[objc.ID](d_.ID, objc.Sel("unlockDocument:"), sender)
}


// Allows the user to make modifications to the document’s file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/unlock(completionHandler:)-6m7rh
func (d_ Document) UnlockWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("unlockWithCompletionHandler:"), completionHandler)
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
func (d_ Document) UpdateChangeCount(change DocumentChangeType) {
	objc.Send[objc.ID](d_.ID, objc.Sel("updateChangeCount:"), change)
}


// Updates the document’s change count settings after a successful save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/updateChangeCount(withToken:for:)
func (d_ Document) UpdateChangeCountWithTokenForSaveOperation(changeCountToken objc.IObject, saveOperation SaveOperationType) {
	objc.Send[objc.ID](d_.ID, objc.Sel("updateChangeCountWithToken:forSaveOperation:"), changeCountToken, saveOperation)
}


// Updates the state of the given user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/updateUserActivityState(_:)
func (d_ Document) UpdateUserActivityState(activity foundation.UserActivity) {
	objc.Send[objc.ID](d_.ID, objc.Sel("updateUserActivityState:"), activity)
}


// Validates the specified user interface item that the receiver manages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/validateUserInterfaceItem(_:)
func (d_ Document) ValidateUserInterfaceItem(item objc.IObject) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}


// Confirms that the error object is not to be presented to the user and the error cannot be recovered from, so cleanup can be done.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/willNotPresentError(_:)
func (d_ Document) WillNotPresentError(error_ objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("willNotPresentError:"), error_)
}


// Called when the receiver is about to present an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/willPresentError(_:)
func (d_ Document) WillPresentError(error_ objc.IObject /* cross-framework: Error */) objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](d_.ID, objc.Sel("willPresentError:"), error_)
	return rv
}


// Called after one of the document’s window controllers loads its nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowControllerDidLoadNib(_:)
func (d_ Document) WindowControllerDidLoadNib(windowController IWindowController) {
	objc.Send[objc.ID](d_.ID, objc.Sel("windowControllerDidLoadNib:"), windowController)
}


// Called before one of the document’s window controllers loads its nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowControllerWillLoadNib(_:)
func (d_ Document) WindowControllerWillLoadNib(windowController IWindowController) {
	objc.Send[objc.ID](d_.ID, objc.Sel("windowControllerWillLoadNib:"), windowController)
}


// Returns the names of the types to which this document can be saved for a specified kind of save operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/writableTypes(for:)
func (d_ Document) WritableTypesForSaveOperation(saveOperation SaveOperationType) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("writableTypesForSaveOperation:"), saveOperation)
	return rv
}


// Writes the contents of the document to a file or file package located by a URL, that is formatted to a specified type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/write(to:ofType:)
func (d_ Document) WriteToURLOfTypeError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:ofType:error:"), url, typeName, outError)
	return rv
}


// Writes the contents of the document to a file or file package located by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/write(to:ofType:for:originalContentsURL:)
func (d_ Document) WriteToURLOfTypeForSaveOperationOriginalContentsURLError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType, absoluteOriginalContentsURL objc.IObject /* cross-framework: NSURL */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:ofType:forSaveOperation:originalContentsURL:error:"), url, typeName, saveOperation, absoluteOriginalContentsURL, outError)
	return rv
}


// Writes the contents of the document to a file or file package located by a URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/writeSafely(to:ofType:for:)
func (d_ Document) WriteSafelyToURLOfTypeForSaveOperationError(url objc.IObject /* cross-framework: NSURL */, typeName objc.IObject /* cross-framework: NSString */, saveOperation SaveOperationType, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeSafelyToURL:ofType:forSaveOperation:error:"), url, typeName, saveOperation, outError)
	return rv
}


// A Boolean value that indicates whether the document is shareable from the standard Share menu.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/allowsDocumentSharing
func (d_ Document) AllowsDocumentSharing() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("allowsDocumentSharing"))
	return rv
}


// The location of the most recently autosaved document contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavedContentsFileURL
func (d_ Document) AutosavedContentsFileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("autosavedContentsFileURL"))
	return rv
}


// The location of the most recently autosaved document contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavedContentsFileURL
func (d_ Document) SetAutosavedContentsFileURL(value objc.IObject /* cross-framework: NSURL */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavedContentsFileURL:"), value)
}


// A Boolean value that indicates whether the document subclass supports autosaving of drafts.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavesDrafts
func (d_ Document) AutosavesDrafts() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("autosavesDrafts"))
	return rv
}


// A Boolean value that indicates whether the document subclass supports autosaving in place.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavesInPlace
func (d_ Document) AutosavesInPlace() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("autosavesInPlace"))
	return rv
}


// The document type to use for an autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavingFileType
func (d_ Document) AutosavingFileType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("autosavingFileType"))
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
func (d_ Document) BackupFileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("backupFileURL"))
	return rv
}


// The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/displayName
func (d_ Document) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("displayName"))
	return rv
}


// The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/displayName
func (d_ Document) SetDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisplayName:"), value)
}


// The last-known modification date of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileModificationDate
func (d_ Document) FileModificationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](d_.ID, objc.Sel("fileModificationDate"))
	return rv
}


// The last-known modification date of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileModificationDate
func (d_ Document) SetFileModificationDate(value objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileModificationDate:"), value)
}


// A Boolean value that indicates whether the user chose to hide the document’s filename extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileNameExtensionWasHiddenInLastRunSavePanel
func (d_ Document) FileNameExtensionWasHiddenInLastRunSavePanel() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileNameExtensionWasHiddenInLastRunSavePanel"))
	return rv
}


// The name of the document type, as specified in the app’s information property-list file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileType
func (d_ Document) FileType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("fileType"))
	return rv
}


// The name of the document type, as specified in the app’s information property-list file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileType
func (d_ Document) SetFileType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileType:"), value)
}


// The file type that was last selected in the Save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileTypeFromLastRunSavePanel
func (d_ Document) FileTypeFromLastRunSavePanel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("fileTypeFromLastRunSavePanel"))
	return rv
}


// The location of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileURL
func (d_ Document) FileURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("fileURL"))
	return rv
}


// The location of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileURL
func (d_ Document) SetFileURL(value objc.IObject /* cross-framework: NSURL */) {
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


// A Boolean value that indicates whether the document owns an undo manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/hasUndoManager
func (d_ Document) HasUndoManager() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("hasUndoManager"))
	return rv
}


// A Boolean value that indicates whether the document owns an undo manager object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/hasUndoManager
func (d_ Document) SetHasUndoManager(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHasUndoManager:"), value)
}


// A Boolean value that indicates whether the document is currently displaying the Versions browser.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isBrowsingVersions
func (d_ Document) BrowsingVersions() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("browsingVersions"))
	return rv
}


// A Boolean value that indicates whether the document has unsaved changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isDocumentEdited
func (d_ Document) DocumentEdited() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("documentEdited"))
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


// A Boolean value that indicates whether the document’s file is completely loaded into memory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isEntireFileLoaded
func (d_ Document) EntireFileLoaded() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("entireFileLoaded"))
	return rv
}


// A Boolean value that indicates whether the document is in read-only mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isInViewingMode
func (d_ Document) InViewingMode() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("inViewingMode"))
	return rv
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


// The name of the document seen by the user in AppleScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lastComponentOfFileName
func (d_ Document) LastComponentOfFileName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("lastComponentOfFileName"))
	return rv
}


// The name of the document seen by the user in AppleScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lastComponentOfFileName
func (d_ Document) SetLastComponentOfFileName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLastComponentOfFileName:"), value)
}


// Returns the object specifier that represents the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/objectSpecifier
func (d_ Document) ObjectSpecifier() foundation.ScriptObjectSpecifier {
	rv := objc.Send[foundation.ScriptObjectSpecifier](d_.ID, objc.Sel("objectSpecifier"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/observedPresentedItemUbiquityAttributes
func (d_ Document) ObservedPresentedItemUbiquityAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("observedPresentedItemUbiquityAttributes"))
	return rv
}


// A print operation you can use to create a PDF representation of the document’s current contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/pdfPrintOperation
func (d_ Document) PDFPrintOperation() IPrintOperation {
	rv := objc.Send[PrintOperation](d_.ID, objc.Sel("PDFPrintOperation"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemURL
func (d_ Document) PresentedItemURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](d_.ID, objc.Sel("presentedItemURL"))
	return rv
}


// A Boolean value that indicates whether the document subclass supports version management.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/preservesVersions
func (d_ Document) PreservesVersions() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("preservesVersions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/previewRepresentableActivityItems
func (d_ Document) PreviewRepresentableActivityItems() []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("previewRepresentableActivityItems"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/previewRepresentableActivityItems
func (d_ Document) SetPreviewRepresentableActivityItems(value []objc.ID) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreviewRepresentableActivityItems:"), nsArray)
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


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/savePanelShowsFileFormatsControl
func (d_ Document) SavePanelShowsFileFormatsControl() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("savePanelShowsFileFormatsControl"))
	return rv
}


// A Boolean value that indicates whether the document’s Save panel displays a list of supported writable document types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/shouldRunSavePanelWithAccessoryView
func (d_ Document) ShouldRunSavePanelWithAccessoryView() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldRunSavePanelWithAccessoryView"))
	return rv
}


// The object that the document uses to support undo/redo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/undoManager
func (d_ Document) UndoManager() foundation.UndoManager {
	rv := objc.Send[foundation.UndoManager](d_.ID, objc.Sel("undoManager"))
	return rv
}


// The object that the document uses to support undo/redo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/undoManager
func (d_ Document) SetUndoManager(value foundation.UndoManager) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUndoManager:"), value)
}


// An object that encapsulates a user activity the document supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/userActivity
func (d_ Document) UserActivity() foundation.UserActivity {
	rv := objc.Send[foundation.UserActivity](d_.ID, objc.Sel("userActivity"))
	return rv
}


// An object that encapsulates a user activity the document supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/userActivity
func (d_ Document) SetUserActivity(value foundation.UserActivity) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUserActivity:"), value)
}


// Returns whether the document object stores its contents in the user’s iCloud document storage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/usesUbiquitousStorage
func (d_ Document) UsesUbiquitousStorage() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("usesUbiquitousStorage"))
	return rv
}


// The document’s current window controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowControllers
func (d_ Document) WindowControllers() []WindowController {
	rv := objc.Send[[]WindowController](d_.ID, objc.Sel("windowControllers"))
	return rv
}


// Returns the document window to use as the parent of a document-modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowForSheet
func (d_ Document) WindowForSheet() IWindow {
	rv := objc.Send[Window](d_.ID, objc.Sel("windowForSheet"))
	return rv
}


// The name of the document’s sole nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowNibName
func (d_ Document) WindowNibName() objc.IObject /* cross-framework: NibName */ {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("windowNibName"))
	return rv
}


// Returns the types of data the receiver can write natively and any types filterable to that native type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/writableTypes
func (d_ Document) WritableTypes() []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("writableTypes"))
	return rv
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


// The key that identifies the document associated with a user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuseractivitydocumenturlkey
func (d_ Document) NSUserActivityDocumentURLKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("NSUserActivityDocumentURLKey"))
	return rv
}


