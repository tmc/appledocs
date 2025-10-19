// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PrintInfo] class.
var (
	printInfoClass     _PrintInfoClass
	printInfoClassOnce sync.Once
)

func getPrintInfoClass() _PrintInfoClass {
	printInfoClassOnce.Do(func() {
		printInfoClass = _PrintInfoClass{objc.GetClass("NSPrintInfo")}
	})
	return printInfoClass
}

type _PrintInfoClass struct {
	class objc.Class
}

// An interface definition for the [PrintInfo] class.
type IPrintInfo interface {
	objectivec.IObject
}

// An object that stores information that’s used to generate printed output. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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




