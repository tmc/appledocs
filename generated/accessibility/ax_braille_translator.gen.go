// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AXBrailleTranslator] class.
var (
	AXBrailleTranslatorClass     _AXBrailleTranslatorClass
	AXBrailleTranslatorClassOnce sync.Once
)

func getAXBrailleTranslatorClass() _AXBrailleTranslatorClass {
	AXBrailleTranslatorClassOnce.Do(func() {
		AXBrailleTranslatorClass = _AXBrailleTranslatorClass{objc.GetClass("AXBrailleTranslator")}
	})
	return AXBrailleTranslatorClass
}

type _AXBrailleTranslatorClass struct {
	class objc.Class
}

// An interface definition for the [AXBrailleTranslator] class.
type IAXBrailleTranslator interface {
	objectivec.IObject
	BackTranslateBraille(braille string) AXBrailleTranslationResult
	TranslatePrintText(printText string) AXBrailleTranslationResult
}

// Translates print text to Braille and Braille to print text according to the given Braille table.


// Translates print text to Braille and Braille to print text according to the given Braille table.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator
type AXBrailleTranslator struct {
	objectivec.Object
}

// AXBrailleTranslatorFrom constructs a [AXBrailleTranslator] from an unsafe.Pointer.
//
// Translates print text to Braille and Braille to print text according to the given Braille table.
func AXBrailleTranslatorFrom(ptr unsafe.Pointer) AXBrailleTranslator {
	return AXBrailleTranslator{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AXBrailleTranslatorClass) Alloc() AXBrailleTranslator {
	rv := objc.Send[AXBrailleTranslator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AXBrailleTranslatorClass) New() AXBrailleTranslator {
	rv := objc.Send[AXBrailleTranslator](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AXBrailleTranslator) Init() AXBrailleTranslator {
	rv := objc.Send[AXBrailleTranslator](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AXBrailleTranslator) Autorelease() AXBrailleTranslator {
	rv := objc.Send[AXBrailleTranslator](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAXBrailleTranslator creates a new AXBrailleTranslator instance.
func NewAXBrailleTranslator() AXBrailleTranslator {
	return getAXBrailleTranslatorClass().New()
}



// Input Braille should use the unicode Braille characters (0x2800-0x28FF).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator/backTranslateBraille(_:)
func (a_ AXBrailleTranslator) BackTranslateBraille(braille string) AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](a_.ID, objc.Sel("backTranslateBraille:"), objc.String(braille))
	return rv
}


// Output Braille uses the unicode Braille characters (0x2800-0x28FF).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator/translatePrintText(_:)
func (a_ AXBrailleTranslator) TranslatePrintText(printText string) AXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](a_.ID, objc.Sel("translatePrintText:"), objc.String(printText))
	return rv
}



