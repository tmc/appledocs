
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PrintInfo] class.
var PrintInfoClass _PrintInfoClass

func init() {
	PrintInfoClass = _PrintInfoClass{objc.GetClass("NSPrintInfo")}
}

type _PrintInfoClass struct {
	objc.Class
}

// An interface definition for the [PrintInfo] class.
type IPrintInfo interface {
	ID() objc.ID
}

type PrintInfo struct {
	id objc.ID
}

func PrintInfoFrom(ptr unsafe.Pointer) PrintInfo {
	return PrintInfo{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (p_ PrintInfo) ID() objc.ID {
	return p_.id
}

// Alloc allocates a new instance without initialization.
func (pc _PrintInfoClass) Alloc() PrintInfo {
	rv := objc.Send[PrintInfo](objc.ID(pc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (pc _PrintInfoClass) New() PrintInfo {
	rv := objc.Send[PrintInfo](objc.ID(pc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewPrintInfo creates and returns a new initialized instance.
func NewPrintInfo() PrintInfo {
	return PrintInfoClass.New()
}

// Init initializes the instance.
func (p_ PrintInfo) Init() PrintInfo {
	rv := objc.Send[PrintInfo](p_.ID(), selInit)
	return rv
}
// The action specified for the job. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property
func (p_ PrintInfo) JobDisposition() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID(), objc.RegisterName("jobDisposition"))
	return rv
}
// SetJobDisposition sets the value of the jobDisposition property. [Full Topic]

//
// [Full Topic]: doc://com.apple.appkit/documentation/AppKit/NSPrintInfo/jobDisposition-swift.property
func (p_ PrintInfo) SetJobDisposition(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID(), objc.RegisterName("setJobDisposition:"), value)
}
