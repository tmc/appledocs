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
	AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler unsafe.Pointer)
	AddWindowController(windowController unsafe.Pointer)
	AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objc.ID, didAutosaveSelector objc.SEL, contextInfo unsafe.Pointer)
	AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler unsafe.Pointer)
	BrowseDocumentVersions(sender objc.ID)
	CanAsynchronouslyWriteToURLOfTypeForSaveOperation(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer) bool
	CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo(delegate objc.ID, shouldCloseSelector objc.SEL, contextInfo unsafe.Pointer)
	ChangeCountTokenForSaveOperation(saveOperation unsafe.Pointer) objc.ID
	CheckAutosavingSafetyAndReturnError(outError unsafe.Pointer) bool
	Close()
	ContinueActivityUsingBlock(block unsafe.Pointer)
	ContinueAsynchronousWorkOnMainThreadUsingBlock(block unsafe.Pointer)
	DataOfTypeError(typeName string, outError unsafe.Pointer) unsafe.Pointer
	DefaultDraftName() unsafe.Pointer
	DuplicateAndReturnError(outError unsafe.Pointer) unsafe.Pointer
	DuplicateDocument(sender objc.ID)
	DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objc.ID, didDuplicateSelector objc.SEL, contextInfo unsafe.Pointer)
	EncodeRestorableStateWithCoder(coder unsafe.Pointer)
	EncodeRestorableStateWithCoderBackgroundQueue(coder unsafe.Pointer, queue unsafe.Pointer)
	FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, absoluteOriginalContentsURL unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer
	FileNameExtensionForTypeSaveOperation(typeName string, saveOperation unsafe.Pointer) unsafe.Pointer
	FileWrapperOfTypeError(typeName string, outError unsafe.Pointer) unsafe.Pointer
	HandleCloseScriptCommand(command unsafe.Pointer) objc.ID
	HandlePrintScriptCommand(command unsafe.Pointer) objc.ID
	HandleSaveScriptCommand(command unsafe.Pointer) objc.ID
	InvalidateRestorableState()
	LockDocument(sender objc.ID)
	LockWithCompletionHandler(completionHandler unsafe.Pointer)
	LockDocumentWithCompletionHandler(completionHandler unsafe.Pointer)
	MakeWindowControllers()
	MoveDocument(sender objc.ID)
	MoveDocumentWithCompletionHandler(completionHandler unsafe.Pointer)
	MoveToURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer)
	MoveDocumentToUbiquityContainer(sender objc.ID)
	PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block unsafe.Pointer)
	PerformAsynchronousFileAccessUsingBlock(block unsafe.Pointer)
	PerformSynchronousFileAccessUsingBlock(block unsafe.Pointer)
	PrepareSharingServicePicker(sharingServicePicker unsafe.Pointer)
	PreparePageLayout(pageLayout unsafe.Pointer) bool
	PrepareSavePanel(savePanel unsafe.Pointer) bool
	PresentError(error unsafe.Pointer) bool
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error unsafe.Pointer, window unsafe.Pointer, delegate objc.ID, didPresentSelector objc.SEL, contextInfo unsafe.Pointer)
	PresentedItemDidChange()
	PresentedItemDidChangeUbiquityAttributes(attributes unsafe.Pointer)
	PresentedItemDidGainVersion(version unsafe.Pointer)
	PresentedItemDidLoseVersion(version unsafe.Pointer)
	PresentedItemDidMoveToURL(newURL unsafe.Pointer)
	PresentedItemDidResolveConflictVersion(version unsafe.Pointer)
	PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo(printSettings unsafe.Pointer, showPrintPanel bool, delegate objc.ID, didPrintSelector objc.SEL, contextInfo unsafe.Pointer)
	PrintDocument(sender objc.ID)
	PrintOperationWithSettingsError(printSettings unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer
	ReadFromURLOfTypeError(url unsafe.Pointer, typeName string, outError unsafe.Pointer) bool
	ReadFromFileWrapperOfTypeError(fileWrapper unsafe.Pointer, typeName string, outError unsafe.Pointer) bool
	ReadFromDataOfTypeError(data unsafe.Pointer, typeName string, outError unsafe.Pointer) bool
	RelinquishPresentedItemToReader(reader unsafe.Pointer)
	RelinquishPresentedItemToWriter(writer unsafe.Pointer)
	RemoveWindowController(windowController unsafe.Pointer)
	RenameDocument(sender objc.ID)
	RestoreStateWithCoder(coder unsafe.Pointer)
	RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier unsafe.Pointer, state unsafe.Pointer, completionHandler unsafe.Pointer)
	RevertToContentsOfURLOfTypeError(url unsafe.Pointer, typeName string, outError unsafe.Pointer) bool
	RevertDocumentToSaved(sender objc.ID)
	RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo(printInfo unsafe.Pointer, delegate objc.ID, didRunSelector objc.SEL, contextInfo unsafe.Pointer)
	RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation unsafe.Pointer, delegate objc.ID, didRunSelector objc.SEL, contextInfo unsafe.Pointer)
	RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo(saveOperation unsafe.Pointer, delegate objc.ID, didSaveSelector objc.SEL, contextInfo unsafe.Pointer)
	RunPageLayout(sender objc.ID)
	SaveDocument(sender objc.ID)
	SaveToURLOfTypeForSaveOperationCompletionHandler(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, completionHandler unsafe.Pointer)
	SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, delegate objc.ID, didSaveSelector objc.SEL, contextInfo unsafe.Pointer)
	SaveDocumentWithDelegateDidSaveSelectorContextInfo(delegate objc.ID, didSaveSelector objc.SEL, contextInfo unsafe.Pointer)
	SaveDocumentAs(sender objc.ID)
	SavePresentedItemChangesWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveDocumentTo(sender objc.ID)
	SaveDocumentToPDF(sender objc.ID)
	SaveToURLOfTypeForSaveOperationError(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, outError unsafe.Pointer) bool
	ScheduleAutosaving()
	SetWindow(window unsafe.Pointer)
	ShareDocumentWithSharingServiceCompletionHandler(sharingService unsafe.Pointer, completionHandler unsafe.Pointer)
	ShouldChangePrintInfo(newPrintInfo unsafe.Pointer) bool
	ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo(windowController unsafe.Pointer, delegate objc.ID, shouldCloseSelector objc.SEL, contextInfo unsafe.Pointer)
	ShowWindows()
	StopBrowsingVersionsWithCompletionHandler(completionHandler unsafe.Pointer)
	UnblockUserInteraction()
	UnlockDocument(sender objc.ID)
	UnlockWithCompletionHandler(completionHandler unsafe.Pointer)
	UnlockDocumentWithCompletionHandler(completionHandler unsafe.Pointer)
	UpdateChangeCount(change unsafe.Pointer)
	UpdateChangeCountWithTokenForSaveOperation(changeCountToken objc.ID, saveOperation unsafe.Pointer)
	UpdateUserActivityState(activity unsafe.Pointer)
	ValidateUserInterfaceItem(item objc.ID) bool
	WillNotPresentError(error unsafe.Pointer)
	WillPresentError(error unsafe.Pointer) unsafe.Pointer
	WindowControllerDidLoadNib(windowController unsafe.Pointer)
	WindowControllerWillLoadNib(windowController unsafe.Pointer)
	WritableTypesForSaveOperation(saveOperation unsafe.Pointer) []string
	WriteToURLOfTypeError(url unsafe.Pointer, typeName string, outError unsafe.Pointer) bool
	WriteToURLOfTypeForSaveOperationOriginalContentsURLError(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, absoluteOriginalContentsURL unsafe.Pointer, outError unsafe.Pointer) bool
	WriteSafelyToURLOfTypeForSaveOperationError(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, outError unsafe.Pointer) bool
}

// An abstract class that defines the interface for macOS documents.
//
// A document is an object that can internally represent data displayed in a window and that can read data from and write data to a file or file package. Documents create and manage one or more window controllers and are in turn managed by a document controller. Documents respond to first-responder action messages to save, revert, and print their data. Conceptually, a document is a container for a body of information identified by a name under which it is stored in a disk file. In this sense, however, the document is not the same as the file but is an object in memory that owns and manages the document data. In the context of AppKit, a document is an instance of a custom subclass that knows how to represent internally, in one or more formats, persistent data that is displayed in windows. A document can read that data from a file and write it to a file. It is also the first-responder target for many menu commands related to documents, such as Save, Revert, and Print. A document manages its window’s edited status and is set up to perform undo and redo operations. When a window is closing, the document is asked before the window delegate to approve the closing. is one of the triad of AppKit classes that establish an architectural basis for document-based apps (the others being and ). For more information about using in a document-based app, see .
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


// Initializes a document located by a URL of a specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/init(contentsOf:ofType:)
func NewDocumentWithContentsOfURLOfTypeError(url unsafe.Pointer, typeName string, outError unsafe.Pointer) Document {
	instance := getDocumentClass().Alloc()
	rv := objc.Send[Document](instance.ID, objc.Sel("initWithContentsOfURL:ofType:error:"), url, objc.String(typeName), outError)
	rv.Autorelease()
	return rv
}

// Initializes a document with the specified contents, and places the resulting document’s file at the designated location.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/init(for:withContentsOf:ofType:)
func NewDocumentForURLWithContentsOfURLOfTypeError(urlOrNil unsafe.Pointer, contentsURL unsafe.Pointer, typeName string, outError unsafe.Pointer) Document {
	instance := getDocumentClass().Alloc()
	rv := objc.Send[Document](instance.ID, objc.Sel("initForURL:withContentsOfURL:ofType:error:"), urlOrNil, contentsURL, objc.String(typeName), outError)
	rv.Autorelease()
	return rv
}

// Initializes a document of a specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/init(type:)
func NewDocumentWithTypeError(typeName string, outError unsafe.Pointer) Document {
	instance := getDocumentClass().Alloc()
	rv := objc.Send[Document](instance.ID, objc.Sel("initWithType:error:"), objc.String(typeName), outError)
	rv.Autorelease()
	return rv
}


// Returns the classes that support secure coding.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/allowedClasses(forRestorableStateKeyPath:)
func (dc _DocumentClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.Send[[]objc.Class](objc.ID(dc.class), objc.Sel("allowedClassesForRestorableStateKeyPath:"), objc.String(keyPath))
	return rv
}

// Returns a Boolean value that indicates whether the receiver reads multiple documents of the given type concurrently.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/canConcurrentlyReadDocuments(ofType:)
func (dc _DocumentClass) CanConcurrentlyReadDocumentsOfType(typeName string) bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("canConcurrentlyReadDocumentsOfType:"), objc.String(typeName))
	return rv
}

// Returns a Boolean value that indicates whether the document can read and write the data natively.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isNativeType(_:)
func (dc _DocumentClass) IsNativeType(type_ string) bool {
	rv := objc.Send[bool](objc.ID(dc.class), objc.Sel("isNativeType:"), objc.String(type_))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/accommodatePresentedItemDeletion(completionHandler:)
func (d_ Document) AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("accommodatePresentedItemDeletionWithCompletionHandler:"), completionHandler)
}

// Adds the specified window controller to the current document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/addWindowController(_:)
func (d_ Document) AddWindowController(windowController unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("addWindowController:"), windowController)
}

// Autosaves the document’s contents to an appropriate location in the file system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosave(withDelegate:didAutosave:contextInfo:)
func (d_ Document) AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objc.ID, didAutosaveSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("autosaveDocumentWithDelegate:didAutosaveSelector:contextInfo:"), delegate, didAutosaveSelector, contextInfo)
}

// Autosaves the document’s contents to an appropriate file-system location, as needed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosave(withImplicitCancellability:completionHandler:)
func (d_ Document) AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("autosaveWithImplicitCancellability:completionHandler:"), autosavingIsImplicitlyCancellable, completionHandler)
}

// Opens the Versions browser in the document’s main window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/browseVersions(_:)
func (d_ Document) BrowseDocumentVersions(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("browseDocumentVersions:"), sender)
}

// Returns whether the receiver can concurrently write to a file or file package located by a URL, that is formatted for a specific type, for a specific kind of save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/canAsynchronouslyWrite(to:ofType:for:)
func (d_ Document) CanAsynchronouslyWriteToURLOfTypeForSaveOperation(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("canAsynchronouslyWriteToURL:ofType:forSaveOperation:"), url, objc.String(typeName), saveOperation)
	return rv
}

// Determines whether to close the document, prompting the user as needed to choose a course of action.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/canClose(withDelegate:shouldClose:contextInfo:)
func (d_ Document) CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo(delegate objc.ID, shouldCloseSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("canCloseDocumentWithDelegate:shouldCloseSelector:contextInfo:"), delegate, shouldCloseSelector, contextInfo)
}

// Returns an object that encapsulates the current record of document changes at the beginning of a save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/changeCountToken(for:)
func (d_ Document) ChangeCountTokenForSaveOperation(saveOperation unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("changeCountTokenForSaveOperation:"), saveOperation)
	return rv
}

// Returns a Boolean value that indicates whether it is safe to autosave document changes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/checkAutosavingSafety()
func (d_ Document) CheckAutosavingSafetyAndReturnError(outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("checkAutosavingSafetyAndReturnError:"), outError)
	return rv
}

// Closes all of the document’s windows and removes the document from its document controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/close()
func (d_ Document) Close() {
	objc.Send[objc.ID](d_.ID, objc.Sel("close"))
}

// Continues to perform the task for a user activity object using a different block.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/continueActivity(_:)
func (d_ Document) ContinueActivityUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("continueActivityUsingBlock:"), block)
}

// Invokes the passed-in block on the main thread.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/continueAsynchronousWorkOnMainThread(_:)
func (d_ Document) ContinueAsynchronousWorkOnMainThreadUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("continueAsynchronousWorkOnMainThreadUsingBlock:"), block)
}

// Creates and returns a data object that contains the contents of the document, formatted to a specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/data(ofType:)
func (d_ Document) DataOfTypeError(typeName string, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("dataOfType:error:"), objc.String(typeName), outError)
	return rv
}

// Returns the default draft name for the document subclass.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/defaultDraftName()
func (d_ Document) DefaultDraftName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("defaultDraftName"))
	return rv
}

// Creates a new document whose contents are the same as the receiver and returns an error object if unsuccessful.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate()
func (d_ Document) DuplicateAndReturnError(outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("duplicateAndReturnError:"), outError)
	return rv
}

// Creates a copy of the receiving document in response to the user choosing Duplicate from the File menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate(_:)
func (d_ Document) DuplicateDocument(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("duplicateDocument:"), sender)
}

// Creates a new document whose contents are the same as the current document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate(withDelegate:didDuplicate:contextInfo:)
func (d_ Document) DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objc.ID, didDuplicateSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("duplicateDocumentWithDelegate:didDuplicateSelector:contextInfo:"), delegate, didDuplicateSelector, contextInfo)
}

// Saves the interface-related state of the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/encodeRestorableState(with:)
func (d_ Document) EncodeRestorableStateWithCoder(coder unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("encodeRestorableStateWithCoder:"), coder)
}

// Saves the interface-related state of the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/encodeRestorableState(with:backgroundQueue:)
func (d_ Document) EncodeRestorableStateWithCoderBackgroundQueue(coder unsafe.Pointer, queue unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("encodeRestorableStateWithCoder:backgroundQueue:"), coder, queue)
}

// Returns the attributes to write to the file or file package at the specified URL, and targeting the specified type of save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileAttributesToWrite(to:ofType:for:originalContentsURL:)
func (d_ Document) FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, absoluteOriginalContentsURL unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileAttributesToWriteToURL:ofType:forSaveOperation:originalContentsURL:error:"), url, objc.String(typeName), saveOperation, absoluteOriginalContentsURL, outError)
	return rv
}

// Returns a filename extension that can be appended to a base filename, for a specified file type and kind of save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileNameExtension(forType:saveOperation:)
func (d_ Document) FileNameExtensionForTypeSaveOperation(typeName string, saveOperation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileNameExtensionForType:saveOperation:"), objc.String(typeName), saveOperation)
	return rv
}

// Creates and returns a file wrapper that contains the contents of the document, formatted to the specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileWrapper(ofType:)
func (d_ Document) FileWrapperOfTypeError(typeName string, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileWrapperOfType:error:"), objc.String(typeName), outError)
	return rv
}

// Handles the Close AppleScript command by attempting to close the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/handleClose(_:)
func (d_ Document) HandleCloseScriptCommand(command unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("handleCloseScriptCommand:"), command)
	return rv
}

// Handles the Print AppleScript command by attempting to print the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/handlePrint(_:)
func (d_ Document) HandlePrintScriptCommand(command unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("handlePrintScriptCommand:"), command)
	return rv
}

// Handles the Save AppleScript command by attempting to save the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/handleSave(_:)
func (d_ Document) HandleSaveScriptCommand(command unsafe.Pointer) objc.ID {
	rv := objc.Send[objc.ID](d_.ID, objc.Sel("handleSaveScriptCommand:"), command)
	return rv
}

// Marks the document’s interface-related state as dirty.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/invalidateRestorableState()
func (d_ Document) InvalidateRestorableState() {
	objc.Send[objc.ID](d_.ID, objc.Sel("invalidateRestorableState"))
}

// Locks the document in response to the user choosing the Lock menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lock(_:)
func (d_ Document) LockDocument(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("lockDocument:"), sender)
}

// Prevents the user from making changes to the document’s file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lock(completionHandler:)-161qv
func (d_ Document) LockWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("lockWithCompletionHandler:"), completionHandler)
}

// Prevents the user from making further changes to the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lock(completionHandler:)-6zuhh
func (d_ Document) LockDocumentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("lockDocumentWithCompletionHandler:"), completionHandler)
}

// Creates the window controller objects that the document uses to display its content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/makeWindowControllers()
func (d_ Document) MakeWindowControllers() {
	objc.Send[objc.ID](d_.ID, objc.Sel("makeWindowControllers"))
}

// Moves the document to a new location in response to the user choosing the Move To… menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/move(_:)
func (d_ Document) MoveDocument(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveDocument:"), sender)
}

// Moves the document to a user-selected location.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/move(completionHandler:)
func (d_ Document) MoveDocumentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveDocumentWithCompletionHandler:"), completionHandler)
}

// Moves the document’s file to the given URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/move(to:completionHandler:)
func (d_ Document) MoveToURLCompletionHandler(url unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveToURL:completionHandler:"), url, completionHandler)
}

// Moves the document to the user’s iCloud storage.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/moveToUbiquityContainer(_:)
func (d_ Document) MoveDocumentToUbiquityContainer(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("moveDocumentToUbiquityContainer:"), sender)
}

// Waits for any work scheduled by previous invocations of this method to complete, then invokes the passed-in block.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/performActivity(withSynchronousWaiting:using:)
func (d_ Document) PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("performActivityWithSynchronousWaiting:usingBlock:"), waitSynchronously, block)
}

// Waits for any scheduled file access to complete but without blocking the main thread, then invokes the passed-in block.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/performAsynchronousFileAccess(_:)
func (d_ Document) PerformAsynchronousFileAccessUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("performAsynchronousFileAccessUsingBlock:"), block)
}

// Waits for any scheduled file access to complete, then invokes the passed-in block.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/performSynchronousFileAccess(_:)
func (d_ Document) PerformSynchronousFileAccessUsingBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("performSynchronousFileAccessUsingBlock:"), block)
}

// Perform any custom setup associated with a sharing service picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/prepare(_:)
func (d_ Document) PrepareSharingServicePicker(sharingServicePicker unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("prepareSharingServicePicker:"), sharingServicePicker)
}

// Adds document-specific content to the Page Layout panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/preparePageLayout(_:)
func (d_ Document) PreparePageLayout(pageLayout unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("preparePageLayout:"), pageLayout)
	return rv
}

// Tells the document to customize the specified Save panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/prepareSavePanel(_:)
func (d_ Document) PrepareSavePanel(savePanel unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("prepareSavePanel:"), savePanel)
	return rv
}

// Presents an error alert to the user as a modal panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentError(_:)
func (d_ Document) PresentError(error unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("presentError:"), error)
	return rv
}

// Presents an error alert to the user as a modal panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (d_ Document) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error unsafe.Pointer, window unsafe.Pointer, delegate objc.ID, didPresentSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error, window, delegate, didPresentSelector, contextInfo)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidChange()
func (d_ Document) PresentedItemDidChange() {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidChange"))
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidChangeUbiquityAttributes(_:)
func (d_ Document) PresentedItemDidChangeUbiquityAttributes(attributes unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidChangeUbiquityAttributes:"), attributes)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidGain(_:)
func (d_ Document) PresentedItemDidGainVersion(version unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidGainVersion:"), version)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidLose(_:)
func (d_ Document) PresentedItemDidLoseVersion(version unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidLoseVersion:"), version)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidMove(to:)
func (d_ Document) PresentedItemDidMoveToURL(newURL unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidMoveToURL:"), newURL)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidResolveConflict(_:)
func (d_ Document) PresentedItemDidResolveConflictVersion(version unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("presentedItemDidResolveConflictVersion:"), version)
}

// Prints the document’s contents, optionally displaying a print panel to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/print(withSettings:showPrintPanel:delegate:didPrint:contextInfo:)
func (d_ Document) PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo(printSettings unsafe.Pointer, showPrintPanel bool, delegate objc.ID, didPrintSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("printDocumentWithSettings:showPrintPanel:delegate:didPrintSelector:contextInfo:"), printSettings, showPrintPanel, delegate, didPrintSelector, contextInfo)
}

// Prints the receiver in response to the user choosing the Print menu command.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/printDocument(_:)
func (d_ Document) PrintDocument(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("printDocument:"), sender)
}

// Creates and returns a print operation for the document’s contents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/printOperation(withSettings:)
func (d_ Document) PrintOperationWithSettingsError(printSettings unsafe.Pointer, outError unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("printOperationWithSettings:error:"), printSettings, outError)
	return rv
}

// Sets the contents of this document by reading from a file or file package, of a specified type, located by a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/read(from:ofType:)-1vttv
func (d_ Document) ReadFromURLOfTypeError(url unsafe.Pointer, typeName string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("readFromURL:ofType:error:"), url, objc.String(typeName), outError)
	return rv
}

// Sets the contents of this document by reading from a file wrapper of a specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/read(from:ofType:)-3rzsi
func (d_ Document) ReadFromFileWrapperOfTypeError(fileWrapper unsafe.Pointer, typeName string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("readFromFileWrapper:ofType:error:"), fileWrapper, objc.String(typeName), outError)
	return rv
}

// Sets the contents of this document by reading from data of a specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/read(from:ofType:)-6g6ai
func (d_ Document) ReadFromDataOfTypeError(data unsafe.Pointer, typeName string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("readFromData:ofType:error:"), data, objc.String(typeName), outError)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/relinquishPresentedItem(toReader:)
func (d_ Document) RelinquishPresentedItemToReader(reader unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("relinquishPresentedItemToReader:"), reader)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/relinquishPresentedItem(toWriter:)
func (d_ Document) RelinquishPresentedItemToWriter(writer unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("relinquishPresentedItemToWriter:"), writer)
}

// Removes the specified window controller from the receiver’s array of window controllers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/removeWindowController(_:)
func (d_ Document) RemoveWindowController(windowController unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("removeWindowController:"), windowController)
}

// Renames the current document in response to the user choosing the Rename menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/rename(_:)
func (d_ Document) RenameDocument(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("renameDocument:"), sender)
}

// Restores the interface-related state of the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/restoreState(with:)
func (d_ Document) RestoreStateWithCoder(coder unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("restoreStateWithCoder:"), coder)
}

// Restores a window that was associated with a document, after that document is reopened.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/restoreWindow(withIdentifier:state:completionHandler:)
func (d_ Document) RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier unsafe.Pointer, state unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("restoreDocumentWindowWithIdentifier:state:completionHandler:"), identifier, state, completionHandler)
}

// Discards all unsaved document modifications and replaces the document’s contents by reading a file or file package located by a URL of a specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/revert(toContentsOf:ofType:)
func (d_ Document) RevertToContentsOfURLOfTypeError(url unsafe.Pointer, typeName string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("revertToContentsOfURL:ofType:error:"), url, objc.String(typeName), outError)
	return rv
}

// The action of the File menu item Revert in a document-based app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/revertToSaved(_:)
func (d_ Document) RevertDocumentToSaved(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("revertDocumentToSaved:"), sender)
}

// Runs the modal page layout panel with the receiver’s printing information object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runModalPageLayout(with:delegate:didRun:contextInfo:)
func (d_ Document) RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo(printInfo unsafe.Pointer, delegate objc.ID, didRunSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runModalPageLayoutWithPrintInfo:delegate:didRunSelector:contextInfo:"), printInfo, delegate, didRunSelector, contextInfo)
}

// Runs the specified print operation modally.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runModalPrintOperation(_:delegate:didRun:contextInfo:)
func (d_ Document) RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation unsafe.Pointer, delegate objc.ID, didRunSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runModalPrintOperation:delegate:didRunSelector:contextInfo:"), printOperation, delegate, didRunSelector, contextInfo)
}

// Presents a modal Save panel to the user, then tries to save the document if the user approves the operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runModalSavePanel(for:delegate:didSave:contextInfo:)
func (d_ Document) RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo(saveOperation unsafe.Pointer, delegate objc.ID, didSaveSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runModalSavePanelForSaveOperation:delegate:didSaveSelector:contextInfo:"), saveOperation, delegate, didSaveSelector, contextInfo)
}

// The action method invoked in the receiver as first responder when the user chooses the Page Setup menu command.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/runPageLayout(_:)
func (d_ Document) RunPageLayout(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("runPageLayout:"), sender)
}

// The action method invoked in the receiver as first responder when the user chooses the Save menu command.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/save(_:)
func (d_ Document) SaveDocument(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveDocument:"), sender)
}

// Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation, and invokes the passed-in completion handler.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/save(to:ofType:for:completionHandler:)
func (d_ Document) SaveToURLOfTypeForSaveOperationCompletionHandler(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveToURL:ofType:forSaveOperation:completionHandler:"), url, objc.String(typeName), saveOperation, completionHandler)
}

// Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/save(to:ofType:for:delegate:didSave:contextInfo:)
func (d_ Document) SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, delegate objc.ID, didSaveSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveToURL:ofType:forSaveOperation:delegate:didSaveSelector:contextInfo:"), url, objc.String(typeName), saveOperation, delegate, didSaveSelector, contextInfo)
}

// Saves the document and delivers the results to the provided delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/save(withDelegate:didSave:contextInfo:)
func (d_ Document) SaveDocumentWithDelegateDidSaveSelectorContextInfo(delegate objc.ID, didSaveSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveDocumentWithDelegate:didSaveSelector:contextInfo:"), delegate, didSaveSelector, contextInfo)
}

// The action method invoked in the receiver as first responder when the user chooses the Save As menu command.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/saveAs(_:)
func (d_ Document) SaveDocumentAs(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveDocumentAs:"), sender)
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/savePresentedItemChanges(completionHandler:)
func (d_ Document) SavePresentedItemChangesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("savePresentedItemChangesWithCompletionHandler:"), completionHandler)
}

// The action method invoked in the receiver as first responder when the user chooses the Save To menu command.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/saveTo(_:)
func (d_ Document) SaveDocumentTo(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveDocumentTo:"), sender)
}

// Exports a PDF representation of the document’s current contents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/saveToPDF(_:)
func (d_ Document) SaveDocumentToPDF(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("saveDocumentToPDF:"), sender)
}

// Saves the contents of the document to a file or file package located by a URL, formatted to a specified type, for a particular kind of save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/saveToURL:ofType:forSaveOperation:error:
func (d_ Document) SaveToURLOfTypeForSaveOperationError(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("saveToURL:ofType:forSaveOperation:error:"), url, objc.String(typeName), saveOperation, outError)
	return rv
}

// Schedules periodic autosaving for the purpose of crash protection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/scheduleAutosaving()
func (d_ Document) ScheduleAutosaving() {
	objc.Send[objc.ID](d_.ID, objc.Sel("scheduleAutosaving"))
}

// Sets the window outlet of this document to the specified value.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/setWindow(_:)
func (d_ Document) SetWindow(window unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWindow:"), window)
}

// Share the document’s file using the specified sharing service.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/share(with:completionHandler:)
func (d_ Document) ShareDocumentWithSharingServiceCompletionHandler(sharingService unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("shareDocumentWithSharingService:completionHandler:"), sharingService, completionHandler)
}

// Returns a Boolean value that indicates whether the document allows changes to the default printing information.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/shouldChangePrintInfo(_:)
func (d_ Document) ShouldChangePrintInfo(newPrintInfo unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldChangePrintInfo:"), newPrintInfo)
	return rv
}

// Determines whether the system should close the document and its associated window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/shouldCloseWindowController(_:delegate:shouldClose:contextInfo:)
func (d_ Document) ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo(windowController unsafe.Pointer, delegate objc.ID, shouldCloseSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("shouldCloseWindowController:delegate:shouldCloseSelector:contextInfo:"), windowController, delegate, shouldCloseSelector, contextInfo)
}

// Displays all of the document’s windows, bringing them to the front and making them main or key as necessary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/showWindows()
func (d_ Document) ShowWindows() {
	objc.Send[objc.ID](d_.ID, objc.Sel("showWindows"))
}

// Dismiss the Versions browser for the current document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/stopBrowsingVersions(completionHandler:)
func (d_ Document) StopBrowsingVersionsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("stopBrowsingVersionsWithCompletionHandler:"), completionHandler)
}

// Unblocks the main thread during asynchronous saving.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/unblockUserInteraction()
func (d_ Document) UnblockUserInteraction() {
	objc.Send[objc.ID](d_.ID, objc.Sel("unblockUserInteraction"))
}

// Unlocks the document in response to the user choosing the Unlock menu item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/unlock(_:)
func (d_ Document) UnlockDocument(sender objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("unlockDocument:"), sender)
}

// Allows the user to make modifications to the document’s file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/unlock(completionHandler:)-6m7rh
func (d_ Document) UnlockWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("unlockWithCompletionHandler:"), completionHandler)
}

// Allows the user to make modifications to the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/unlock(completionHandler:)-8p8zd
func (d_ Document) UnlockDocumentWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("unlockDocumentWithCompletionHandler:"), completionHandler)
}

// Updates the receiver’s change count according to the given change type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/updateChangeCount(_:)
func (d_ Document) UpdateChangeCount(change unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("updateChangeCount:"), change)
}

// Updates the document’s change count settings after a successful save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/updateChangeCount(withToken:for:)
func (d_ Document) UpdateChangeCountWithTokenForSaveOperation(changeCountToken objc.ID, saveOperation unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("updateChangeCountWithToken:forSaveOperation:"), changeCountToken, saveOperation)
}

// Updates the state of the given user activity.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/updateUserActivityState(_:)
func (d_ Document) UpdateUserActivityState(activity unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("updateUserActivityState:"), activity)
}

// Validates the specified user interface item that the receiver manages.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/validateUserInterfaceItem(_:)
func (d_ Document) ValidateUserInterfaceItem(item objc.ID) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}

// Confirms that the error object is not to be presented to the user and the error cannot be recovered from, so cleanup can be done.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/willNotPresentError(_:)
func (d_ Document) WillNotPresentError(error unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("willNotPresentError:"), error)
}

// Called when the receiver is about to present an error.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/willPresentError(_:)
func (d_ Document) WillPresentError(error unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("willPresentError:"), error)
	return rv
}

// Called after one of the document’s window controllers loads its nib file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowControllerDidLoadNib(_:)
func (d_ Document) WindowControllerDidLoadNib(windowController unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("windowControllerDidLoadNib:"), windowController)
}

// Called before one of the document’s window controllers loads its nib file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowControllerWillLoadNib(_:)
func (d_ Document) WindowControllerWillLoadNib(windowController unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("windowControllerWillLoadNib:"), windowController)
}

// Returns the names of the types to which this document can be saved for a specified kind of save operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/writableTypes(for:)
func (d_ Document) WritableTypesForSaveOperation(saveOperation unsafe.Pointer) []string {
	rv := objc.Send[[]string](d_.ID, objc.Sel("writableTypesForSaveOperation:"), saveOperation)
	return rv
}

// Writes the contents of the document to a file or file package located by a URL, that is formatted to a specified type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/write(to:ofType:)
func (d_ Document) WriteToURLOfTypeError(url unsafe.Pointer, typeName string, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:ofType:error:"), url, objc.String(typeName), outError)
	return rv
}

// Writes the contents of the document to a file or file package located by a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/write(to:ofType:for:originalContentsURL:)
func (d_ Document) WriteToURLOfTypeForSaveOperationOriginalContentsURLError(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, absoluteOriginalContentsURL unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeToURL:ofType:forSaveOperation:originalContentsURL:error:"), url, objc.String(typeName), saveOperation, absoluteOriginalContentsURL, outError)
	return rv
}

// Writes the contents of the document to a file or file package located by a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/writeSafely(to:ofType:for:)
func (d_ Document) WriteSafelyToURLOfTypeForSaveOperationError(url unsafe.Pointer, typeName string, saveOperation unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("writeSafelyToURL:ofType:forSaveOperation:error:"), url, objc.String(typeName), saveOperation, outError)
	return rv
}

// A Boolean value that indicates whether the document is shareable from the standard Share menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/allowsDocumentSharing
func (d_ Document) AllowsDocumentSharing() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("allowsDocumentSharing"))
	return rv
}

// The location of the most recently autosaved document contents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavedContentsFileURL
func (d_ Document) AutosavedContentsFileURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("autosavedContentsFileURL"))
	return rv
}


// SetAutosavedContentsFileURL sets the value of the autosavedContentsFileURL property.
// The location of the most recently autosaved document contents.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavedContentsFileURL
func (d_ Document) SetAutosavedContentsFileURL(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavedContentsFileURL:"), value)
}
// The document type to use for an autosave operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavingFileType
func (d_ Document) AutosavingFileType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("autosavingFileType"))
	return rv
}

// A Boolean value that indicates whether you can cancel an in-progress autosave operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/autosavingIsImplicitlyCancellable
func (d_ Document) AutosavingIsImplicitlyCancellable() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("autosavingIsImplicitlyCancellable"))
	return rv
}

// The URL for the document’s backup file that was created during an autosave operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/backupFileURL
func (d_ Document) BackupFileURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("backupFileURL"))
	return rv
}

// The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/displayName
func (d_ Document) DisplayName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("displayName"))
	return rv
}


// SetDisplayName sets the value of the displayName property.
// The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/displayName
func (d_ Document) SetDisplayName(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisplayName:"), value)
}
// The last-known modification date of the document’s on-disk representation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileModificationDate
func (d_ Document) FileModificationDate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileModificationDate"))
	return rv
}


// SetFileModificationDate sets the value of the fileModificationDate property.
// The last-known modification date of the document’s on-disk representation.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileModificationDate
func (d_ Document) SetFileModificationDate(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileModificationDate:"), value)
}
// A Boolean value that indicates whether the user chose to hide the document’s filename extension.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileNameExtensionWasHiddenInLastRunSavePanel
func (d_ Document) FileNameExtensionWasHiddenInLastRunSavePanel() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileNameExtensionWasHiddenInLastRunSavePanel"))
	return rv
}

// The name of the document type, as specified in the app’s information property-list file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileType
func (d_ Document) FileType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileType"))
	return rv
}


// SetFileType sets the value of the fileType property.
// The name of the document type, as specified in the app’s information property-list file.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileType
func (d_ Document) SetFileType(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileType:"), value)
}
// The file type that was last selected in the Save panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileTypeFromLastRunSavePanel
func (d_ Document) FileTypeFromLastRunSavePanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileTypeFromLastRunSavePanel"))
	return rv
}

// The location of the document’s on-disk representation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileURL
func (d_ Document) FileURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("fileURL"))
	return rv
}


// SetFileURL sets the value of the fileURL property.
// The location of the document’s on-disk representation.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/fileURL
func (d_ Document) SetFileURL(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileURL:"), value)
}
// A Boolean value that indicates whether the document has changes that have not been autosaved.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/hasUnautosavedChanges
func (d_ Document) HasUnautosavedChanges() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("hasUnautosavedChanges"))
	return rv
}

// A Boolean value that indicates whether the document owns an undo manager object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/hasUndoManager
func (d_ Document) HasUndoManager() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("hasUndoManager"))
	return rv
}


// SetHasUndoManager sets the value of the hasUndoManager property.
// A Boolean value that indicates whether the document owns an undo manager object.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/hasUndoManager
func (d_ Document) SetHasUndoManager(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHasUndoManager:"), value)
}
// A Boolean value that indicates whether the document is currently displaying the Versions browser.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isBrowsingVersions
func (d_ Document) BrowsingVersions() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("browsingVersions"))
	return rv
}

// A Boolean value that indicates whether the document has unsaved changes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isDocumentEdited
func (d_ Document) DocumentEdited() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("documentEdited"))
	return rv
}

// A Boolean value that indicates whether the document is a draft that the user has not yet saved.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isDraft
func (d_ Document) Draft() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("draft"))
	return rv
}


// SetDraft sets the value of the draft property.
// A Boolean value that indicates whether the document is a draft that the user has not yet saved.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isDraft
func (d_ Document) SetDraft(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDraft:"), value)
}
// A Boolean value that indicates whether the document’s file is completely loaded into memory.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isEntireFileLoaded
func (d_ Document) EntireFileLoaded() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("entireFileLoaded"))
	return rv
}

// A Boolean value that indicates whether the document is in read-only mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isInViewingMode
func (d_ Document) InViewingMode() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("inViewingMode"))
	return rv
}

// A Boolean value that indicates whether or not the file can be written to.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/isLocked
func (d_ Document) Locked() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("locked"))
	return rv
}

// A Boolean value that indicates whether the document archives previously saved versions of the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/keepBackupFile
func (d_ Document) KeepBackupFile() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("keepBackupFile"))
	return rv
}

// The name of the document seen by the user in AppleScript.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lastComponentOfFileName
func (d_ Document) LastComponentOfFileName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("lastComponentOfFileName"))
	return rv
}


// SetLastComponentOfFileName sets the value of the lastComponentOfFileName property.
// The name of the document seen by the user in AppleScript.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/lastComponentOfFileName
func (d_ Document) SetLastComponentOfFileName(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLastComponentOfFileName:"), value)
}
// Returns the object specifier that represents the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/objectSpecifier
func (d_ Document) ObjectSpecifier() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("objectSpecifier"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/observedPresentedItemUbiquityAttributes
func (d_ Document) ObservedPresentedItemUbiquityAttributes() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("observedPresentedItemUbiquityAttributes"))
	return rv
}

// A print operation you can use to create a PDF representation of the document’s current contents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/pdfPrintOperation
func (d_ Document) PDFPrintOperation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("PDFPrintOperation"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemURL
func (d_ Document) PresentedItemURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("presentedItemURL"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/previewRepresentableActivityItems
func (d_ Document) PreviewRepresentableActivityItems() []objc.ID {
	rv := objc.Send[[]objc.ID](d_.ID, objc.Sel("previewRepresentableActivityItems"))
	return rv
}


// SetPreviewRepresentableActivityItems sets the value of the previewRepresentableActivityItems property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/previewRepresentableActivityItems
func (d_ Document) SetPreviewRepresentableActivityItems(value []objc.ID) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreviewRepresentableActivityItems:"), value)
}
// The printing information associated with the document.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/printInfo
func (d_ Document) PrintInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("printInfo"))
	return rv
}


// SetPrintInfo sets the value of the printInfo property.
// The printing information associated with the document.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/printInfo
func (d_ Document) SetPrintInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPrintInfo:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/savePanelShowsFileFormatsControl
func (d_ Document) SavePanelShowsFileFormatsControl() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("savePanelShowsFileFormatsControl"))
	return rv
}

// A Boolean value that indicates whether the document’s Save panel displays a list of supported writable document types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/shouldRunSavePanelWithAccessoryView
func (d_ Document) ShouldRunSavePanelWithAccessoryView() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("shouldRunSavePanelWithAccessoryView"))
	return rv
}

// The object that the document uses to support undo/redo operations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/undoManager
func (d_ Document) UndoManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("undoManager"))
	return rv
}


// SetUndoManager sets the value of the undoManager property.
// The object that the document uses to support undo/redo operations.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/undoManager
func (d_ Document) SetUndoManager(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUndoManager:"), value)
}
// An object that encapsulates a user activity the document supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/userActivity
func (d_ Document) UserActivity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("userActivity"))
	return rv
}


// SetUserActivity sets the value of the userActivity property.
// An object that encapsulates a user activity the document supports.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/userActivity
func (d_ Document) SetUserActivity(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUserActivity:"), value)
}
// The document’s current window controllers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowControllers
func (d_ Document) WindowControllers() []WindowController {
	rv := objc.Send[[]WindowController](d_.ID, objc.Sel("windowControllers"))
	return rv
}

// Returns the document window to use as the parent of a document-modal sheet.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowForSheet
func (d_ Document) WindowForSheet() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("windowForSheet"))
	return rv
}

// The name of the document’s sole nib file.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/windowNibName
func (d_ Document) WindowNibName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("windowNibName"))
	return rv
}


