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
	// properties:
	AllowsDocumentSharing() bool
	SetAllowsDocumentSharing(value bool)
	AutosavedContentsFileURL() objc.IObject /* cross-framework: URL */
	SetAutosavedContentsFileURL(value objc.IObject /* cross-framework: URL */)
	AutosavingFileType() objc.IObject /* cross-framework: NSString */
	SetAutosavingFileType(value objc.IObject /* cross-framework: NSString */)
	AutosavingIsImplicitlyCancellable() bool
	SetAutosavingIsImplicitlyCancellable(value bool)
	BackupFileURL() objc.IObject /* cross-framework: URL */
	SetBackupFileURL(value objc.IObject /* cross-framework: URL */)
	DisplayName() objc.IObject /* cross-framework: NSString */
	SetDisplayName(value objc.IObject /* cross-framework: NSString */)
	FileModificationDate() objc.IObject /* cross-framework: Date */
	SetFileModificationDate(value objc.IObject /* cross-framework: Date */)
	FileNameExtensionWasHiddenInLastRunSavePanel() bool
	SetFileNameExtensionWasHiddenInLastRunSavePanel(value bool)
	FileType() objc.IObject /* cross-framework: NSString */
	SetFileType(value objc.IObject /* cross-framework: NSString */)
	FileTypeFromLastRunSavePanel() objc.IObject /* cross-framework: NSString */
	SetFileTypeFromLastRunSavePanel(value objc.IObject /* cross-framework: NSString */)
	FileURL() objc.IObject /* cross-framework: URL */
	SetFileURL(value objc.IObject /* cross-framework: URL */)
	HasUnautosavedChanges() bool
	SetHasUnautosavedChanges(value bool)
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
	KeepBackupFile() bool
	SetKeepBackupFile(value bool)
	LastComponentOfFileName() objc.IObject /* cross-framework: NSString */
	SetLastComponentOfFileName(value objc.IObject /* cross-framework: NSString */)
	ObjectSpecifier() objc.IObject /* cross-framework: ScriptObjectSpecifier */
	SetObjectSpecifier(value objc.IObject /* cross-framework: ScriptObjectSpecifier */)
	ObservedPresentedItemUbiquityAttributes() unsafe.Pointer
	SetObservedPresentedItemUbiquityAttributes(value unsafe.Pointer)
	PdfPrintOperation() IPrintOperation
	SetPdfPrintOperation(value IPrintOperation)
	PresentedItemURL() objc.IObject /* cross-framework: URL */
	SetPresentedItemURL(value objc.IObject /* cross-framework: URL */)
	PreviewRepresentableActivityItems() PreviewRepresentableActivityItem /* not a class type */
	SetPreviewRepresentableActivityItems(value PreviewRepresentableActivityItem /* not a class type */)
	PrintInfo() IPrintInfo
	SetPrintInfo(value IPrintInfo)
	SavePanelShowsFileFormatsControl() bool
	SetSavePanelShowsFileFormatsControl(value bool)
	ShouldRunSavePanelWithAccessoryView() bool
	SetShouldRunSavePanelWithAccessoryView(value bool)
	UndoManager() objc.IObject /* cross-framework: UndoManager */
	SetUndoManager(value objc.IObject /* cross-framework: UndoManager */)
	UserActivity() objc.IObject /* cross-framework: UserActivity */
	SetUserActivity(value objc.IObject /* cross-framework: UserActivity */)
	WindowControllers() objc.IObject /* cross-framework: WindowController */
	SetWindowControllers(value objc.IObject /* cross-framework: WindowController */)
	WindowForSheet() objc.IObject /* cross-framework: Window */
	SetWindowForSheet(value objc.IObject /* cross-framework: Window */)
	WindowNibName() unsafe.Pointer
	SetWindowNibName(value unsafe.Pointer)
	NSUserActivityDocumentURLKey() objc.IObject /* cross-framework: NSString */
	// methods:
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
func (d_ Document) AutosavedContentsFileURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("autosavedContentsFileURL"))
	return rv
}


// The location of the most recently autosaved document contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/autosavedcontentsfileurl
func (d_ Document) SetAutosavedContentsFileURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavedContentsFileURL:"), value)
}


// The document type to use for an autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/autosavingfiletype
func (d_ Document) AutosavingFileType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("autosavingFileType"))
	return rv
}


// The document type to use for an autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/autosavingfiletype
func (d_ Document) SetAutosavingFileType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavingFileType:"), value)
}


// A Boolean value that indicates whether you can cancel an in-progress autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/autosavingisimplicitlycancellable
func (d_ Document) AutosavingIsImplicitlyCancellable() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("autosavingIsImplicitlyCancellable"))
	return rv
}


// A Boolean value that indicates whether you can cancel an in-progress autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/autosavingisimplicitlycancellable
func (d_ Document) SetAutosavingIsImplicitlyCancellable(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setAutosavingIsImplicitlyCancellable:"), value)
}


// The URL for the document’s backup file that was created during an autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/backupfileurl
func (d_ Document) BackupFileURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("backupFileURL"))
	return rv
}


// The URL for the document’s backup file that was created during an autosave operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/backupfileurl
func (d_ Document) SetBackupFileURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setBackupFileURL:"), value)
}


// The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/displayname
func (d_ Document) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("displayName"))
	return rv
}


// The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/displayname
func (d_ Document) SetDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setDisplayName:"), value)
}


// The last-known modification date of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filemodificationdate
func (d_ Document) FileModificationDate() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](d_.ID, objc.Sel("fileModificationDate"))
	return rv
}


// The last-known modification date of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filemodificationdate
func (d_ Document) SetFileModificationDate(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileModificationDate:"), value)
}


// A Boolean value that indicates whether the user chose to hide the document’s filename extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filenameextensionwashiddeninlastrunsavepanel
func (d_ Document) FileNameExtensionWasHiddenInLastRunSavePanel() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("fileNameExtensionWasHiddenInLastRunSavePanel"))
	return rv
}


// A Boolean value that indicates whether the user chose to hide the document’s filename extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filenameextensionwashiddeninlastrunsavepanel
func (d_ Document) SetFileNameExtensionWasHiddenInLastRunSavePanel(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileNameExtensionWasHiddenInLastRunSavePanel:"), value)
}


// The name of the document type, as specified in the app’s information property-list file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filetype
func (d_ Document) FileType() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("fileType"))
	return rv
}


// The name of the document type, as specified in the app’s information property-list file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filetype
func (d_ Document) SetFileType(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileType:"), value)
}


// The file type that was last selected in the Save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filetypefromlastrunsavepanel
func (d_ Document) FileTypeFromLastRunSavePanel() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("fileTypeFromLastRunSavePanel"))
	return rv
}


// The file type that was last selected in the Save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/filetypefromlastrunsavepanel
func (d_ Document) SetFileTypeFromLastRunSavePanel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileTypeFromLastRunSavePanel:"), value)
}


// The location of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/fileurl
func (d_ Document) FileURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("fileURL"))
	return rv
}


// The location of the document’s on-disk representation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/fileurl
func (d_ Document) SetFileURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setFileURL:"), value)
}


// A Boolean value that indicates whether the document has changes that have not been autosaved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/hasunautosavedchanges
func (d_ Document) HasUnautosavedChanges() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("hasUnautosavedChanges"))
	return rv
}


// A Boolean value that indicates whether the document has changes that have not been autosaved.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/hasunautosavedchanges
func (d_ Document) SetHasUnautosavedChanges(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHasUnautosavedChanges:"), value)
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


// A Boolean value that indicates whether the document archives previously saved versions of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/keepbackupfile
func (d_ Document) KeepBackupFile() bool {
	rv := objc.Send[bool](d_.ID, objc.Sel("keepBackupFile"))
	return rv
}


// A Boolean value that indicates whether the document archives previously saved versions of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/keepbackupfile
func (d_ Document) SetKeepBackupFile(value bool) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setKeepBackupFile:"), value)
}


// The name of the document seen by the user in AppleScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/lastcomponentoffilename
func (d_ Document) LastComponentOfFileName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("lastComponentOfFileName"))
	return rv
}


// The name of the document seen by the user in AppleScript.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/lastcomponentoffilename
func (d_ Document) SetLastComponentOfFileName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setLastComponentOfFileName:"), value)
}


// Returns the object specifier that represents the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/objectspecifier
func (d_ Document) ObjectSpecifier() objc.IObject /* cross-framework: ScriptObjectSpecifier */ {
	rv := objc.Send[foundation.ScriptObjectSpecifier](d_.ID, objc.Sel("objectSpecifier"))
	return rv
}


// Returns the object specifier that represents the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/objectspecifier
func (d_ Document) SetObjectSpecifier(value objc.IObject /* cross-framework: ScriptObjectSpecifier */) {
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
func (d_ Document) PresentedItemURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](d_.ID, objc.Sel("presentedItemURL"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/presenteditemurl
func (d_ Document) SetPresentedItemURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPresentedItemURL:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/previewrepresentableactivityitems
func (d_ Document) PreviewRepresentableActivityItems() PreviewRepresentableActivityItem /* not a class type */ {
	rv := objc.Send[PreviewRepresentableActivityItem](d_.ID, objc.Sel("previewRepresentableActivityItems"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/previewrepresentableactivityitems
func (d_ Document) SetPreviewRepresentableActivityItems(value PreviewRepresentableActivityItem /* not a class type */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPreviewRepresentableActivityItems:"), value)
}


// The printing information associated with the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/printinfo
func (d_ Document) PrintInfo() IPrintInfo {
	rv := objc.Send[PrintInfo](d_.ID, objc.Sel("printInfo"))
	return rv
}


// The printing information associated with the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/printinfo
func (d_ Document) SetPrintInfo(value IPrintInfo) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setPrintInfo:"), value)
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
func (d_ Document) UndoManager() objc.IObject /* cross-framework: UndoManager */ {
	rv := objc.Send[foundation.UndoManager](d_.ID, objc.Sel("undoManager"))
	return rv
}


// The object that the document uses to support undo/redo operations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/undomanager
func (d_ Document) SetUndoManager(value objc.IObject /* cross-framework: UndoManager */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUndoManager:"), value)
}


// An object that encapsulates a user activity the document supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/useractivity
func (d_ Document) UserActivity() objc.IObject /* cross-framework: UserActivity */ {
	rv := objc.Send[foundation.UserActivity](d_.ID, objc.Sel("userActivity"))
	return rv
}


// An object that encapsulates a user activity the document supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/useractivity
func (d_ Document) SetUserActivity(value objc.IObject /* cross-framework: UserActivity */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setUserActivity:"), value)
}


// The document’s current window controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/windowcontrollers
func (d_ Document) WindowControllers() objc.IObject /* cross-framework: WindowController */ {
	rv := objc.Send[WindowController](d_.ID, objc.Sel("windowControllers"))
	return rv
}


// The document’s current window controllers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/windowcontrollers
func (d_ Document) SetWindowControllers(value objc.IObject /* cross-framework: WindowController */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWindowControllers:"), value)
}


// Returns the document window to use as the parent of a document-modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/windowforsheet
func (d_ Document) WindowForSheet() objc.IObject /* cross-framework: Window */ {
	rv := objc.Send[Window](d_.ID, objc.Sel("windowForSheet"))
	return rv
}


// Returns the document window to use as the parent of a document-modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/windowforsheet
func (d_ Document) SetWindowForSheet(value objc.IObject /* cross-framework: Window */) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWindowForSheet:"), value)
}


// The name of the document’s sole nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/windownibname
func (d_ Document) WindowNibName() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("windowNibName"))
	return rv
}


// The name of the document’s sole nib file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsdocument/windownibname
func (d_ Document) SetWindowNibName(value unsafe.Pointer) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWindowNibName:"), value)
}


// The key that identifies the document associated with a user activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsuseractivitydocumenturlkey
func (d_ Document) NSUserActivityDocumentURLKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](d_.ID, objc.Sel("NSUserActivityDocumentURLKey"))
	return rv
}



