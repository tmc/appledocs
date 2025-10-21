// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PrintInfo] class.
var (
	PrintInfoClass     _PrintInfoClass
	PrintInfoClassOnce sync.Once
)

func getPrintInfoClass() _PrintInfoClass {
	PrintInfoClassOnce.Do(func() {
		PrintInfoClass = _PrintInfoClass{objc.GetClass("NSPrintInfo")}
	})
	return PrintInfoClass
}

type _PrintInfoClass struct {
	class objc.Class
}

// An interface definition for the [PrintInfo] class.
type IPrintInfo interface {
	objectivec.IObject
	Dictionary() unsafe.Pointer
}

// An object that stores information that’s used to generate printed output.
//
// A shared object is automatically created for an app and is used by default for all printing jobs for that app. The printing information in an object is stored in a dictionary. To access the standard attributes in the dictionary directly, this class defines a set of keys and provides the method. You can also initialize an instance of this class using the method. You can use this dictionary to store custom information associated with a print job. Any non-object values should be stored as or objects in the dictionary. See for a list of types which should be stored as numbers. For other non-object values, use the class. To store custom information that belongs in printing presets you should use the dictionary returned by the method.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo
type PrintInfo struct {
	objectivec.Object
}

// PrintInfoFrom constructs a [PrintInfo] from an unsafe.Pointer.
//
// An object that stores information that’s used to generate printed output.
func PrintInfoFrom(ptr unsafe.Pointer) PrintInfo {
	return PrintInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PrintInfoClass) Alloc() PrintInfo {
	rv := objc.Send[PrintInfo](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PrintInfoClass) New() PrintInfo {
	rv := objc.Send[PrintInfo](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PrintInfo) Init() PrintInfo {
	rv := objc.Send[PrintInfo](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PrintInfo) Autorelease() PrintInfo {
	rv := objc.Send[PrintInfo](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPrintInfo creates a new PrintInfo instance.
func NewPrintInfo() PrintInfo {
	return getPrintInfoClass().New()
}


// Returns the print info’s dictionary that contains the printing attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/dictionary()
func (p_ PrintInfo) Dictionary() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("dictionary"))
	return rv
}

// A Boolean value that indicates whether the image is centered horizontally.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isHorizontallyCentered
func (p_ PrintInfo) HorizontallyCentered() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("horizontallyCentered"))
	return rv
}


// SetHorizontallyCentered sets the value of the horizontallyCentered property.
// A Boolean value that indicates whether the image is centered horizontally.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isHorizontallyCentered
func (p_ PrintInfo) SetHorizontallyCentered(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setHorizontallyCentered:"), value)
}

// A Boolean value that indicates whether only the currently selected contents should be printed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isSelectionOnly
func (p_ PrintInfo) SelectionOnly() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("selectionOnly"))
	return rv
}


// SetSelectionOnly sets the value of the selectionOnly property.
// A Boolean value that indicates whether only the currently selected contents should be printed.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/isSelectionOnly
func (p_ PrintInfo) SetSelectionOnly(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSelectionOnly:"), value)
}

// The action specified for the job.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property
func (p_ PrintInfo) JobDisposition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("jobDisposition"))
	return rv
}


// SetJobDisposition sets the value of the jobDisposition property.
// The action specified for the job.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property
func (p_ PrintInfo) SetJobDisposition(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setJobDisposition:"), value)
}



