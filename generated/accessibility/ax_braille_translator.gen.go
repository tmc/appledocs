// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AXBrailleTranslator */


/* debug [class_header]: Header for AXBrailleTranslator */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AXBrailleTranslator */
// An interface definition for the [AXBrailleTranslator] class.
type IAXBrailleTranslator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AXBrailleTranslator */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AXBrailleTranslator */
	// methods:
	BackTranslateBraille(braille objc.IObject /* cross-framework: NSString */) IAXBrailleTranslationResult
	TranslatePrintText(printText objc.IObject /* cross-framework: NSString */) IAXBrailleTranslationResult
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AXBrailleTranslator */
// Alloc allocates a new instance without initialization.
func (ac _AXBrailleTranslatorClass) Alloc() AXBrailleTranslator {
	rv := objc.Send[AXBrailleTranslator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AXBrailleTranslator */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AXBrailleTranslator */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator/init(brailleTable:)
func NewAXBrailleTranslatorWithBrailleTable(brailleTable IAXBrailleTable) AXBrailleTranslator {
	instance := getAXBrailleTranslatorClass().Alloc()
	rv := objc.Send[AXBrailleTranslator](instance.ID, objc.Sel("initWithBrailleTable:"), brailleTable)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAXBrailleTranslatorWithBrailleTable */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AXBrailleTranslator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AXBrailleTranslator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AXBrailleTranslator */

// Input Braille should use the unicode Braille characters (0x2800-0x28FF).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator/backTranslateBraille(_:)
func (a_ AXBrailleTranslator) BackTranslateBraille(braille objc.IObject /* cross-framework: NSString */) IAXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](a_.ID, objc.Sel("backTranslateBraille:"), braille)
	return rv
}/* debug [instance_methods/method]: BackTranslateBraille */


// Output Braille uses the unicode Braille characters (0x2800-0x28FF).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AXBrailleTranslator/translatePrintText(_:)
func (a_ AXBrailleTranslator) TranslatePrintText(printText objc.IObject /* cross-framework: NSString */) IAXBrailleTranslationResult {
	rv := objc.Send[AXBrailleTranslationResult](a_.ID, objc.Sel("translatePrintText:"), printText)
	return rv
}/* debug [instance_methods/method]: TranslatePrintText */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AXBrailleTranslator */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AXBrailleTranslator */


