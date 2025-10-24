// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSGlyphGenerator */


/* debug [class_header]: Header for NSGlyphGenerator */
// The class instance for the [GlyphGenerator] class.
var (
	GlyphGeneratorClass     _GlyphGeneratorClass
	GlyphGeneratorClassOnce sync.Once
)

func getGlyphGeneratorClass() _GlyphGeneratorClass {
	GlyphGeneratorClassOnce.Do(func() {
		GlyphGeneratorClass = _GlyphGeneratorClass{objc.GetClass("NSGlyphGenerator")}
	})
	return GlyphGeneratorClass
}

type _GlyphGeneratorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GlyphGenerator */
// An interface definition for the [GlyphGenerator] class.
type IGlyphGenerator interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for GlyphGenerator */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GlyphGenerator */
	// methods:
	GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorage unsafe.Pointer, nChars uint, glyphIndex uint, charIndex uint)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GlyphGenerator */
// Alloc allocates a new instance without initialization.
func (gc _GlyphGeneratorClass) Alloc() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GlyphGeneratorClass) New() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GlyphGenerator) Init() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GlyphGenerator) Autorelease() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGlyphGenerator creates a new GlyphGenerator instance.
func NewGlyphGenerator() GlyphGenerator {
	return getGlyphGeneratorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GlyphGenerator */
// An object that performs the initial, nominal glyph generation phase in the layout process.
//
// The nominal glyph generation pass essentially generates one glyph per character; the typesetter may later make substitutions in the glyph stream, for example, changing an acute accent glyph followed by an “e” glyph into a single acute-accented “é” glyph. communicates via the protocol. An example of a class that conforms to the protocol is .


// An object that performs the initial, nominal glyph generation phase in the layout process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphGenerator
type GlyphGenerator struct {
	objectivec.Object
}

// GlyphGeneratorFrom constructs a [GlyphGenerator] from an unsafe.Pointer.
//
// An object that performs the initial, nominal glyph generation phase in the layout process.
func GlyphGeneratorFrom(ptr unsafe.Pointer) GlyphGenerator {
	return GlyphGenerator{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GlyphGenerator *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GlyphGenerator */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GlyphGenerator */

// Returns a shared instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphGenerator/shared
func (gc _GlyphGeneratorClass) SharedGlyphGenerator() GlyphGenerator {
	rv := objc.Send[GlyphGenerator](objc.ID(gc.class), objc.Sel("sharedGlyphGenerator"))
	return rv
}/* debug [class_properties_class/property]: sharedGlyphGenerator */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GlyphGenerator */

// Generates glyphs for the specified glyph storage object ( by default).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphGenerator/generateGlyphs(for:desiredNumberOfCharacters:glyphIndex:characterIndex:)
func (g_ GlyphGenerator) GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex(glyphStorage unsafe.Pointer, nChars uint, glyphIndex uint, charIndex uint) {
	objc.Send[objc.ID](g_.ID, objc.Sel("generateGlyphsForGlyphStorage:desiredNumberOfCharacters:glyphIndex:characterIndex:"), glyphStorage, nChars, glyphIndex, charIndex)
}/* debug [instance_methods/method]: GenerateGlyphsForGlyphStorageDesiredNumberOfCharactersGlyphIndexCharacterIndex */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GlyphGenerator */

// Returns a shared instance of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphGenerator/shared
func (g_ GlyphGenerator) SharedGlyphGenerator() IGlyphGenerator {
	rv := objc.Send[GlyphGenerator](g_.ID, objc.Sel("sharedGlyphGenerator"))
	return rv
}/* debug [instance_properties/getter]: sharedGlyphGenerator */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSGlyphGenerator */



