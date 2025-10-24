// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/vision"
)

/* debug [class.gen.go]: Generating class PDFDocument */


/* debug [class_header]: Header for PDFDocument */
// The class instance for the [PDFDocument] class.
var (
	PDFDocumentClass     _PDFDocumentClass
	PDFDocumentClassOnce sync.Once
)

func getPDFDocumentClass() _PDFDocumentClass {
	PDFDocumentClassOnce.Do(func() {
		PDFDocumentClass = _PDFDocumentClass{objc.GetClass("PDFDocument")}
	})
	return PDFDocumentClass
}

type _PDFDocumentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PDFDocument */
// An interface definition for the [PDFDocument] class.
type IPDFDocument interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PDFDocument */
	// properties:
	AccessPermissions() PDFAccessPermissions
	AllowsCommenting() bool
	AllowsContentAccessibility() bool
	AllowsCopying() bool
	AllowsDocumentAssembly() bool
	AllowsDocumentChanges() bool
	AllowsFormFieldEntry() bool
	AllowsPrinting() bool
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DocumentAttributes() objc.IObject /* cross-framework: NSDictionary */
	SetDocumentAttributes(value objc.IObject /* cross-framework: NSDictionary */)
	DocumentRef() PDFDocumentRef /* not a class type */
	DocumentURL() objc.IObject /* cross-framework: NSURL */
	IsEncrypted() bool
	IsFinding() bool
	IsLocked() bool
	MajorVersion() int
	MinorVersion() int
	OutlineRoot() IPDFOutline
	SetOutlineRoot(value IPDFOutline)
	PageClass() objc.Class
	PageCount() uint
	PermissionsStatus() PDFDocumentPermissions
	SelectionForEntireDocument() IPDFSelection
	String() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PDFDocument */
	// methods:
	BeginFindStringWithOptions(string_ objc.IObject /* cross-framework: NSString */, options StringCompareOptions /* not a class type */)
	BeginFindStringsWithOptions(strings []string, options StringCompareOptions /* not a class type */)
	CancelFindString()
	DataRepresentation() foundation.Data
	DataRepresentationWithOptions(options objc.IObject /* cross-framework: NSDictionary */) foundation.Data
	ExchangePageAtIndexWithPageAtIndex(indexA uint, indexB uint)
	FindStringFromSelectionWithOptions(string_ objc.IObject /* cross-framework: NSString */, selection IPDFSelection, options StringCompareOptions /* not a class type */) IPDFSelection
	FindStringWithOptions(string_ objc.IObject /* cross-framework: NSString */, options StringCompareOptions /* not a class type */) []PDFSelection
	IndexForPage(page IPDFPage) uint
	InsertPageAtIndex(page IPDFPage, index uint)
	OutlineItemForSelection(selection IPDFSelection) IPDFOutline
	PageAtIndex(index uint) IPDFPage
	PrintOperationForPrintInfoScalingModeAutoRotate(printInfo appkit.PrintInfo, scaleMode PDFPrintScalingMode, doRotate bool) appkit.PrintOperation
	RemovePageAtIndex(index uint)
	SelectionFromPageAtPointToPageAtPoint(startPage IPDFPage, startPoint vision.Point, endPage IPDFPage, endPoint vision.Point) IPDFSelection
	SelectionFromPageAtPointToPageAtPointWithGranularity(startPage IPDFPage, startPoint corefoundation.CGPoint, endPage IPDFPage, endPoint corefoundation.CGPoint, granularity PDFSelectionGranularity) IPDFSelection
	SelectionFromPageAtCharacterIndexToPageAtCharacterIndex(startPage IPDFPage, startCharacter uint, endPage IPDFPage, endCharacter uint) IPDFSelection
	UnlockWithPassword(password objc.IObject /* cross-framework: NSString */) bool
	WriteToURL(url objc.IObject /* cross-framework: NSURL */) bool
	WriteToURLWithOptions(url objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary) bool
	WriteToFile(path objc.IObject /* cross-framework: NSString */) bool
	WriteToFileWithOptions(path objc.IObject /* cross-framework: NSString */, options foundation.IDictionary) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PDFDocument */
// Alloc allocates a new instance without initialization.
func (pc _PDFDocumentClass) Alloc() PDFDocument {
	rv := objc.Send[PDFDocument](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PDFDocumentClass) New() PDFDocument {
	rv := objc.Send[PDFDocument](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFDocument) Init() PDFDocument {
	rv := objc.Send[PDFDocument](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFDocument) Autorelease() PDFDocument {
	rv := objc.Send[PDFDocument](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFDocument creates a new PDFDocument instance.
func NewPDFDocument() PDFDocument {
	return getPDFDocumentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PDFDocument */
// An object that represents PDF data or a PDF file and defines methods for writing, searching, and selecting PDF data.
//
// The other utility classes are either instantiated from methods in , as are and ; or support it, as do and . You initialize a object with PDF data or with a URL to a PDF file. You can then ask for the page count, add or delete pages, perform a find, or parse selected content into an object.


// An object that represents PDF data or a PDF file and defines methods for writing, searching, and selecting PDF data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument
type PDFDocument struct {
	objectivec.Object
}

// PDFDocumentFrom constructs a [PDFDocument] from an unsafe.Pointer.
//
// An object that represents PDF data or a PDF file and defines methods for writing, searching, and selecting PDF data.
func PDFDocumentFrom(ptr unsafe.Pointer) PDFDocument {
	return PDFDocument{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PDFDocument */

// Initializes a object with the passed-in data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/init(data:)
func NewPDFDocumentWithData(data objc.IObject /* cross-framework: NSData */) PDFDocument {
	instance := getPDFDocumentClass().Alloc()
	rv := objc.Send[PDFDocument](instance.ID, objc.Sel("initWithData:"), data)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFDocumentWithData */


// Initializes a object with the contents at the specified URL (if the URL is invalid, this method returns ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/init(url:)
func NewPDFDocumentWithURL(url objc.IObject /* cross-framework: NSURL */) PDFDocument {
	instance := getPDFDocumentClass().Alloc()
	rv := objc.Send[PDFDocument](instance.ID, objc.Sel("initWithURL:"), url)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPDFDocumentWithURL */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PDFDocument */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PDFDocument */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PDFDocument */

// Asynchronously finds all instances of the specified string in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/beginFindString(_:withOptions:)
func (p_ PDFDocument) BeginFindStringWithOptions(string_ objc.IObject /* cross-framework: NSString */, options StringCompareOptions /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginFindString:withOptions:"), string_, options)
}/* debug [instance_methods/method]: BeginFindStringWithOptions */


// Asynchronously finds all instances of the specified array of strings in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/beginFindStrings(_:withOptions:)
func (p_ PDFDocument) BeginFindStringsWithOptions(strings []string, options StringCompareOptions /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("beginFindStrings:withOptions:"), strings, options)
}/* debug [instance_methods/method]: BeginFindStringsWithOptions */


// Cancels a search initiated with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/cancelFindString()
func (p_ PDFDocument) CancelFindString() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancelFindString"))
}/* debug [instance_methods/method]: CancelFindString */


// Returns a representation of the document as an object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/dataRepresentation()
func (p_ PDFDocument) DataRepresentation() foundation.Data {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("dataRepresentation"))
	return rv
}/* debug [instance_methods/method]: DataRepresentation */


// Returns a representation of the document as an object with additional options applied, such as filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/dataRepresentation(options:)
func (p_ PDFDocument) DataRepresentationWithOptions(options objc.IObject /* cross-framework: NSDictionary */) foundation.Data {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("dataRepresentationWithOptions:"), options)
	return rv
}/* debug [instance_methods/method]: DataRepresentationWithOptions */


// Swaps one page with another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/exchangePage(at:withPageAt:)
func (p_ PDFDocument) ExchangePageAtIndexWithPageAtIndex(indexA uint, indexB uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("exchangePageAtIndex:withPageAtIndex:"), indexA, indexB)
}/* debug [instance_methods/method]: ExchangePageAtIndexWithPageAtIndex */


// Synchronously finds the next occurance of a string after the specified selection (or before the selection if you specified as a search option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/findString(_:fromSelection:withOptions:)
func (p_ PDFDocument) FindStringFromSelectionWithOptions(string_ objc.IObject /* cross-framework: NSString */, selection IPDFSelection, options StringCompareOptions /* not a class type */) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("findString:fromSelection:withOptions:"), string_, selection, options)
	return rv
}/* debug [instance_methods/method]: FindStringFromSelectionWithOptions */


// Synchronously finds all instances of the specified string in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/findString(_:withOptions:)
func (p_ PDFDocument) FindStringWithOptions(string_ objc.IObject /* cross-framework: NSString */, options StringCompareOptions /* not a class type */) []PDFSelection {
	rv := objc.Send[[]PDFSelection](p_.ID, objc.Sel("findString:withOptions:"), string_, options)
	return rv
}/* debug [instance_methods/method]: FindStringWithOptions */


// Gets the index number for the specified page.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/index(for:)
func (p_ PDFDocument) IndexForPage(page IPDFPage) uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("indexForPage:"), page)
	return rv
}/* debug [instance_methods/method]: IndexForPage */


// Inserts a page at the specified index point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/insert(_:at:)
func (p_ PDFDocument) InsertPageAtIndex(page IPDFPage, index uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("insertPage:atIndex:"), page, index)
}/* debug [instance_methods/method]: InsertPageAtIndex */


// Returns the most likely parent PDF outline object for the selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/outlineItem(for:)
func (p_ PDFDocument) OutlineItemForSelection(selection IPDFSelection) IPDFOutline {
	rv := objc.Send[PDFOutline](p_.ID, objc.Sel("outlineItemForSelection:"), selection)
	return rv
}/* debug [instance_methods/method]: OutlineItemForSelection */


// Returns the page at the specified index number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/page(at:)
func (p_ PDFDocument) PageAtIndex(index uint) IPDFPage {
	rv := objc.Send[PDFPage](p_.ID, objc.Sel("pageAtIndex:"), index)
	return rv
}/* debug [instance_methods/method]: PageAtIndex */


// Returns a print operation suitable for printing the PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/printOperation(for:scalingMode:autoRotate:)
func (p_ PDFDocument) PrintOperationForPrintInfoScalingModeAutoRotate(printInfo appkit.PrintInfo, scaleMode PDFPrintScalingMode, doRotate bool) appkit.PrintOperation {
	rv := objc.Send[appkit.PrintOperation](p_.ID, objc.Sel("printOperationForPrintInfo:scalingMode:autoRotate:"), printInfo, scaleMode, doRotate)
	return rv
}/* debug [instance_methods/method]: PrintOperationForPrintInfoScalingModeAutoRotate */


// Removes the page at the specified index point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/removePage(at:)
func (p_ PDFDocument) RemovePageAtIndex(index uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removePageAtIndex:"), index)
}/* debug [instance_methods/method]: RemovePageAtIndex */


// Returns the specified selection based on starting and ending points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/selection(from:at:to:at:)
func (p_ PDFDocument) SelectionFromPageAtPointToPageAtPoint(startPage IPDFPage, startPoint vision.Point, endPage IPDFPage, endPoint vision.Point) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionFromPage:atPoint:toPage:atPoint:"), startPage, startPoint, endPage, endPoint)
	return rv
}/* debug [instance_methods/method]: SelectionFromPageAtPointToPageAtPoint */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/selection(from:at:to:at:with:)
func (p_ PDFDocument) SelectionFromPageAtPointToPageAtPointWithGranularity(startPage IPDFPage, startPoint corefoundation.CGPoint, endPage IPDFPage, endPoint corefoundation.CGPoint, granularity PDFSelectionGranularity) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionFromPage:atPoint:toPage:atPoint:withGranularity:"), startPage, startPoint, endPage, endPoint, granularity)
	return rv
}/* debug [instance_methods/method]: SelectionFromPageAtPointToPageAtPointWithGranularity */


// Returns the specified selection based on starting and ending character indexes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/selection(from:atCharacterIndex:to:atCharacterIndex:)
func (p_ PDFDocument) SelectionFromPageAtCharacterIndexToPageAtCharacterIndex(startPage IPDFPage, startCharacter uint, endPage IPDFPage, endCharacter uint) IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionFromPage:atCharacterIndex:toPage:atCharacterIndex:"), startPage, startCharacter, endPage, endCharacter)
	return rv
}/* debug [instance_methods/method]: SelectionFromPageAtCharacterIndexToPageAtCharacterIndex */


// Attempts to unlock an encrypted document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/unlock(withPassword:)
func (p_ PDFDocument) UnlockWithPassword(password objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("unlockWithPassword:"), password)
	return rv
}/* debug [instance_methods/method]: UnlockWithPassword */


// Writes the document to a location specified by the passed-in URL.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/write(to:)
func (p_ PDFDocument) WriteToURL(url objc.IObject /* cross-framework: NSURL */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("writeToURL:"), url)
	return rv
}/* debug [instance_methods/method]: WriteToURL */


// Writes the document to the specified URL with the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/write(to:withOptions:)
func (p_ PDFDocument) WriteToURLWithOptions(url objc.IObject /* cross-framework: NSURL */, options foundation.IDictionary) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("writeToURL:withOptions:"), url, options)
	return rv
}/* debug [instance_methods/method]: WriteToURLWithOptions */


// Writes the document to a file at the specified path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/write(toFile:)
func (p_ PDFDocument) WriteToFile(path objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("writeToFile:"), path)
	return rv
}/* debug [instance_methods/method]: WriteToFile */


// Writes the document to a file at the specified path with the specified options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/write(toFile:withOptions:)
func (p_ PDFDocument) WriteToFileWithOptions(path objc.IObject /* cross-framework: NSString */, options foundation.IDictionary) bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("writeToFile:withOptions:"), path, options)
	return rv
}/* debug [instance_methods/method]: WriteToFileWithOptions */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PDFDocument */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/accessPermissions
func (p_ PDFDocument) AccessPermissions() PDFAccessPermissions {
	rv := objc.Send[PDFAccessPermissions](p_.ID, objc.Sel("accessPermissions"))
	return rv
}/* debug [instance_properties/getter]: accessPermissions */


// A Boolean value indicating whether you can create or modify document annotations, including form field entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsCommenting
func (p_ PDFDocument) AllowsCommenting() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsCommenting"))
	return rv
}/* debug [instance_properties/getter]: allowsCommenting */


// A Boolean value indicating whether you can extract content from the document, but only for the purpose of accessibility.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsContentAccessibility
func (p_ PDFDocument) AllowsContentAccessibility() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsContentAccessibility"))
	return rv
}/* debug [instance_properties/getter]: allowsContentAccessibility */


// A Boolean value indicating whether the document allows copying of content to the Pasteboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsCopying
func (p_ PDFDocument) AllowsCopying() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsCopying"))
	return rv
}/* debug [instance_properties/getter]: allowsCopying */


// A Boolean value indicating whether you can manage a document by inserting, deleting, and rotating pages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsDocumentAssembly
func (p_ PDFDocument) AllowsDocumentAssembly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsDocumentAssembly"))
	return rv
}/* debug [instance_properties/getter]: allowsDocumentAssembly */


// A Boolean value indicating whether you can modify the document contents except for document attributes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsDocumentChanges
func (p_ PDFDocument) AllowsDocumentChanges() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsDocumentChanges"))
	return rv
}/* debug [instance_properties/getter]: allowsDocumentChanges */


// A Boolean value indicating whether you can modify form field entries even if you can’t edit document annotations.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsFormFieldEntry
func (p_ PDFDocument) AllowsFormFieldEntry() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsFormFieldEntry"))
	return rv
}/* debug [instance_properties/getter]: allowsFormFieldEntry */


// A Boolean value indicating whether the document allows printing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/allowsPrinting
func (p_ PDFDocument) AllowsPrinting() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("allowsPrinting"))
	return rv
}/* debug [instance_properties/getter]: allowsPrinting */


// The object acting as the delegate for the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/delegate
func (p_ PDFDocument) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The object acting as the delegate for the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/delegate
func (p_ PDFDocument) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A dictionary of document metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/documentAttributes
func (p_ PDFDocument) DocumentAttributes() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](p_.ID, objc.Sel("documentAttributes"))
	return rv
}/* debug [instance_properties/getter]: documentAttributes */


// A dictionary of document metadata.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/documentAttributes
func (p_ PDFDocument) SetDocumentAttributes(value objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDocumentAttributes:"), value)
}/* debug [instance_properties/setter]: documentAttributes */


// The associated with the object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/documentRef
func (p_ PDFDocument) DocumentRef() PDFDocumentRef /* not a class type */ {
	rv := objc.Send[PDFDocumentRef](p_.ID, objc.Sel("documentRef"))
	return rv
}/* debug [instance_properties/getter]: documentRef */


// The URL for the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/documentURL
func (p_ PDFDocument) DocumentURL() objc.IObject /* cross-framework: NSURL */ {
	rv := objc.Send[foundation.NSURL](p_.ID, objc.Sel("documentURL"))
	return rv
}/* debug [instance_properties/getter]: documentURL */


// A Boolean value specifying whether the document is encrypted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/isEncrypted
func (p_ PDFDocument) IsEncrypted() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isEncrypted"))
	return rv
}/* debug [instance_properties/getter]: isEncrypted */


// Returns a Boolean value indicating whether an asynchronous find operation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/isFinding
func (p_ PDFDocument) IsFinding() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFinding"))
	return rv
}/* debug [instance_properties/getter]: isFinding */


// A Boolean value indicating whether the document is locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/isLocked
func (p_ PDFDocument) IsLocked() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isLocked"))
	return rv
}/* debug [instance_properties/getter]: isLocked */


// The major version of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/majorVersion
func (p_ PDFDocument) MajorVersion() int {
	rv := objc.Send[int](p_.ID, objc.Sel("majorVersion"))
	return rv
}/* debug [instance_properties/getter]: majorVersion */


// The minor version of the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/minorVersion
func (p_ PDFDocument) MinorVersion() int {
	rv := objc.Send[int](p_.ID, objc.Sel("minorVersion"))
	return rv
}/* debug [instance_properties/getter]: minorVersion */


// The document’s root outline to a PDF outline object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/outlineRoot
func (p_ PDFDocument) OutlineRoot() IPDFOutline {
	rv := objc.Send[PDFOutline](p_.ID, objc.Sel("outlineRoot"))
	return rv
}/* debug [instance_properties/getter]: outlineRoot */


// The document’s root outline to a PDF outline object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/outlineRoot
func (p_ PDFDocument) SetOutlineRoot(value IPDFOutline) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOutlineRoot:"), value)
}/* debug [instance_properties/setter]: outlineRoot */


// The class that is allocated and initialized when page objects are created for the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/pageClass
func (p_ PDFDocument) PageClass() objc.Class {
	rv := objc.Send[objc.Class](p_.ID, objc.Sel("pageClass"))
	return rv
}/* debug [instance_properties/getter]: pageClass */


// The number of pages in the document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/pageCount
func (p_ PDFDocument) PageCount() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("pageCount"))
	return rv
}/* debug [instance_properties/getter]: pageCount */


// The permissions status of the PDF document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/permissionsStatus
func (p_ PDFDocument) PermissionsStatus() PDFDocumentPermissions {
	rv := objc.Send[PDFDocumentPermissions](p_.ID, objc.Sel("permissionsStatus"))
	return rv
}/* debug [instance_properties/getter]: permissionsStatus */


// Returns a selection representing the textual content of the entire document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/selectionForEntireDocument
func (p_ PDFDocument) SelectionForEntireDocument() IPDFSelection {
	rv := objc.Send[PDFSelection](p_.ID, objc.Sel("selectionForEntireDocument"))
	return rv
}/* debug [instance_properties/getter]: selectionForEntireDocument */


// A string representing the textual content for the entire document.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFDocument/string
func (p_ PDFDocument) String() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("string"))
	return rv
}/* debug [instance_properties/getter]: string */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PDFDocument */


