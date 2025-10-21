// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PrintOperation] class.
var (
	PrintOperationClass     _PrintOperationClass
	PrintOperationClassOnce sync.Once
)

func getPrintOperationClass() _PrintOperationClass {
	PrintOperationClassOnce.Do(func() {
		PrintOperationClass = _PrintOperationClass{objc.GetClass("NSPrintOperation")}
	})
	return PrintOperationClass
}

type _PrintOperationClass struct {
	class objc.Class
}

// An interface definition for the [PrintOperation] class.
type IPrintOperation interface {
	objectivec.IObject
	CleanUpOperation()
	RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow unsafe.Pointer, delegate objc.ID, didRunSelector objc.SEL, contextInfo unsafe.Pointer)
}

// An object that controls operations that generate Encapsulated PostScript (EPS) code, Portable Document Format (PDF) code, or print jobs.
//
// An object works in conjunction with two other objects: an object, which specifies how the code should be generated, and an object, which generates the actual code. It is important to note that the majority of methods in copy the instance of passed into them. Future changes to that print info are not reflected in the print info retained by the current object. All changes should be made to the print info before passing to the methods of this class. The only method in which does not copy the instance is .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation
type PrintOperation struct {
	objectivec.Object
}

// PrintOperationFrom constructs a [PrintOperation] from an unsafe.Pointer.
//
// An object that controls operations that generate Encapsulated PostScript (EPS) code, Portable Document Format (PDF) code, or print jobs.
func PrintOperationFrom(ptr unsafe.Pointer) PrintOperation {
	return PrintOperation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PrintOperationClass) Alloc() PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PrintOperationClass) New() PrintOperation {
	rv := objc.Send[PrintOperation](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PrintOperation) Init() PrintOperation {
	rv := objc.Send[PrintOperation](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PrintOperation) Autorelease() PrintOperation {
	rv := objc.Send[PrintOperation](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPrintOperation creates a new PrintOperation instance.
func NewPrintOperation() PrintOperation {
	return getPrintOperationClass().New()
}


// Creates and returns a new print operation object ready to control the copying of EPS graphics from the specified view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/epsOperation(with:inside:to:)
func (pc _PrintOperationClass) EPSOperationWithViewInsideRectToData(view unsafe.Pointer, rect coregraphics.CGRect, data unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("EPSOperationWithView:insideRect:toData:"), view, rect, data)
	return rv
}

// Called at the end of a print operation to remove the print operation as the current operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/cleanUp()
func (p_ PrintOperation) CleanUpOperation() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cleanUpOperation"))
}

// Runs the print operation, calling your custom delegate method upon completion.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/runModal(for:delegate:didRun:contextInfo:)
func (p_ PrintOperation) RunOperationModalForWindowDelegateDidRunSelectorContextInfo(docWindow unsafe.Pointer, delegate objc.ID, didRunSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("runOperationModalForWindow:delegate:didRunSelector:contextInfo:"), docWindow, delegate, didRunSelector, contextInfo)
}

// A Boolean value that indicates whether the print operation is an EPS or PDF copy operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/isCopyingOperation
func (p_ PrintOperation) CopyingOperation() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("copyingOperation"))
	return rv
}

// The PDF panel object to use during the operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfPanel
func (p_ PrintOperation) PDFPanel() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("PDFPanel"))
	return rv
}


// SetPDFPanel sets the value of the PDFPanel property.
// The PDF panel object to use during the operation.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/pdfPanel
func (p_ PrintOperation) SetPDFPanel(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPDFPanel:"), value)
}

// The printing information associated with the print operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printInfo
func (p_ PrintOperation) PrintInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("printInfo"))
	return rv
}


// SetPrintInfo sets the value of the printInfo property.
// The printing information associated with the print operation.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/printInfo
func (p_ PrintOperation) SetPrintInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPrintInfo:"), value)
}

// The view object that generates the actual data for the print operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/view
func (p_ PrintOperation) View() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("view"))
	return rv
}



