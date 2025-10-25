// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptionRuby */


/* debug [class_header]: Header for AVCaptionRuby */
// The class instance for the [CaptionRuby] class.
var (
	CaptionRubyClass     _CaptionRubyClass
	CaptionRubyClassOnce sync.Once
)

func getCaptionRubyClass() _CaptionRubyClass {
	CaptionRubyClassOnce.Do(func() {
		CaptionRubyClass = _CaptionRubyClass{objc.GetClass("AVCaptionRuby")}
	})
	return CaptionRubyClass
}

type _CaptionRubyClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptionRuby */
// An interface definition for the [CaptionRuby] class.
type ICaptionRuby interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptionRuby */
	// properties:
	Alignment() CaptionRubyAlignment
	Position() CaptionRubyPosition
	Text() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptionRuby */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptionRuby */
// Alloc allocates a new instance without initialization.
func (cc _CaptionRubyClass) Alloc() CaptionRuby {
	rv := objc.Send[CaptionRuby](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionRubyClass) New() CaptionRuby {
	rv := objc.Send[CaptionRuby](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionRuby) Init() CaptionRuby {
	rv := objc.Send[CaptionRuby](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionRuby) Autorelease() CaptionRuby {
	rv := objc.Send[CaptionRuby](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionRuby creates a new CaptionRuby instance.
func NewCaptionRuby() CaptionRuby {
	return getCaptionRubyClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptionRuby */
// An object that presents ruby characters.
//
// Ruby characters are small annotations, typically used in Japanese content, that render alongside the base text.


// An object that presents ruby characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby
type CaptionRuby struct {
	objectivec.Object
}

// CaptionRubyFrom constructs a [CaptionRuby] from an unsafe.Pointer.
//
// An object that presents ruby characters.
func CaptionRubyFrom(ptr unsafe.Pointer) CaptionRuby {
	return CaptionRuby{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptionRuby */

// Creates ruby text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/init(text:)
func NewCaptionRubyWithText(text objc.IObject /* cross-framework: NSString */) CaptionRuby {
	instance := getCaptionRubyClass().Alloc()
	rv := objc.Send[CaptionRuby](instance.ID, objc.Sel("initWithText:"), text)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptionRubyWithText */


// Creates ruby text with position and alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/init(text:position:alignment:)
func NewCaptionRubyWithTextPositionAlignment(text objc.IObject /* cross-framework: NSString */, position CaptionRubyPosition, alignment CaptionRubyAlignment) CaptionRuby {
	instance := getCaptionRubyClass().Alloc()
	rv := objc.Send[CaptionRuby](instance.ID, objc.Sel("initWithText:position:alignment:"), text, position, alignment)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptionRubyWithTextPositionAlignment */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptionRuby */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptionRuby */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptionRuby */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptionRuby */

// The ruby text alignment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/alignment
func (c_ CaptionRuby) Alignment() CaptionRubyAlignment {
	rv := objc.Send[CaptionRubyAlignment](c_.ID, objc.Sel("alignment"))
	return rv
}/* debug [instance_properties/getter]: alignment */


// The ruby text position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/position
func (c_ CaptionRuby) Position() CaptionRubyPosition {
	rv := objc.Send[CaptionRubyPosition](c_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */


// The ruby text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption/Ruby/text
func (c_ CaptionRuby) Text() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("text"))
	return rv
}/* debug [instance_properties/getter]: text */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptionRuby */


