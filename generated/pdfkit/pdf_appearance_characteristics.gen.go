// Code generated from Apple documentation for PDFKit. DO NOT EDIT.

package pdfkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PDFAppearanceCharacteristics] class.
var (
	PDFAppearanceCharacteristicsClass     _PDFAppearanceCharacteristicsClass
	PDFAppearanceCharacteristicsClassOnce sync.Once
)

func getPDFAppearanceCharacteristicsClass() _PDFAppearanceCharacteristicsClass {
	PDFAppearanceCharacteristicsClassOnce.Do(func() {
		PDFAppearanceCharacteristicsClass = _PDFAppearanceCharacteristicsClass{objc.GetClass("PDFAppearanceCharacteristics")}
	})
	return PDFAppearanceCharacteristicsClass
}

type _PDFAppearanceCharacteristicsClass struct {
	class objc.Class
}

// An interface definition for the [PDFAppearanceCharacteristics] class.
type IPDFAppearanceCharacteristics interface {
	objectivec.IObject
}

// An object that represents appearance characteristics of a widget annotation.
//
// [Full Topic]: https://developer.apple.com/documentation/PDFKit/PDFAppearanceCharacteristics
type PDFAppearanceCharacteristics struct {
	objectivec.Object
}

// PDFAppearanceCharacteristicsFrom constructs a [PDFAppearanceCharacteristics] from an unsafe.Pointer.
//
// An object that represents appearance characteristics of a widget annotation.
func PDFAppearanceCharacteristicsFrom(ptr unsafe.Pointer) PDFAppearanceCharacteristics {
	return PDFAppearanceCharacteristics{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PDFAppearanceCharacteristicsClass) Alloc() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PDFAppearanceCharacteristicsClass) New() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PDFAppearanceCharacteristics) Init() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PDFAppearanceCharacteristics) Autorelease() PDFAppearanceCharacteristics {
	rv := objc.Send[PDFAppearanceCharacteristics](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPDFAppearanceCharacteristics creates a new PDFAppearanceCharacteristics instance.
func NewPDFAppearanceCharacteristics() PDFAppearanceCharacteristics {
	return getPDFAppearanceCharacteristicsClass().New()
}




