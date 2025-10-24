// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSRangeSpecifier */


/* debug [class_header]: Header for NSRangeSpecifier */
// The class instance for the [RangeSpecifier] class.
var (
	RangeSpecifierClass     _RangeSpecifierClass
	RangeSpecifierClassOnce sync.Once
)

func getRangeSpecifierClass() _RangeSpecifierClass {
	RangeSpecifierClassOnce.Do(func() {
		RangeSpecifierClass = _RangeSpecifierClass{objc.GetClass("NSRangeSpecifier")}
	})
	return RangeSpecifierClass
}

type _RangeSpecifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RangeSpecifier */
// An interface definition for the [RangeSpecifier] class.
type IRangeSpecifier interface {
	IScriptObjectSpecifier
	
/* debug [class_interface_properties]: Properties for RangeSpecifier */
	// properties:
	StartSpecifier() IScriptObjectSpecifier
	SetStartSpecifier(value IScriptObjectSpecifier)
	EndSpecifier() IScriptObjectSpecifier
	SetEndSpecifier(value IScriptObjectSpecifier)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RangeSpecifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RangeSpecifier */
// Alloc allocates a new instance without initialization.
func (rc _RangeSpecifierClass) Alloc() RangeSpecifier {
	rv := objc.Send[RangeSpecifier](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RangeSpecifierClass) New() RangeSpecifier {
	rv := objc.Send[RangeSpecifier](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RangeSpecifier) Init() RangeSpecifier {
	rv := objc.Send[RangeSpecifier](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RangeSpecifier) Autorelease() RangeSpecifier {
	rv := objc.Send[RangeSpecifier](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRangeSpecifier creates a new RangeSpecifier instance.
func NewRangeSpecifier() RangeSpecifier {
	return getRangeSpecifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RangeSpecifier */
// A specifier for a range of objects in a container.
//
// An object specifies a range (that is, an uninterrupted series) of objects in a container through two delimiting objects. The range is represented by two object specifiers, a start specifier and an end specifier, which can be of any specifier type (such as or object). These specifiers are evaluated in the context of the same container object as the range specifier itself. You don’t normally subclass .


// A specifier for a range of objects in a container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier
type RangeSpecifier struct {
	ScriptObjectSpecifier
}

// RangeSpecifierFrom constructs a [RangeSpecifier] from an unsafe.Pointer.
//
// A specifier for a range of objects in a container.
func RangeSpecifierFrom(ptr unsafe.Pointer) RangeSpecifier {
	return RangeSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RangeSpecifier */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/init(coder:)
func NewRangeSpecifierWithCoder(inCoder ICoder) RangeSpecifier {
	instance := getRangeSpecifierClass().Alloc()
	rv := objc.Send[RangeSpecifier](instance.ID, objc.Sel("initWithCoder:"), inCoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewRangeSpecifierWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RangeSpecifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RangeSpecifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RangeSpecifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RangeSpecifier */

// Returns the object specifier representing the first object of the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/startSpecifier
func (r_ RangeSpecifier) StartSpecifier() IScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](r_.ID, objc.Sel("startSpecifier"))
	return rv
}/* debug [instance_properties/getter]: startSpecifier */


// Returns the object specifier representing the first object of the range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangeSpecifier/startSpecifier
func (r_ RangeSpecifier) SetStartSpecifier(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setStartSpecifier:"), value)
}/* debug [instance_properties/setter]: startSpecifier */


// Sets the object specifier representing the last object of the range to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrangespecifier/endspecifier
func (r_ RangeSpecifier) EndSpecifier() IScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](r_.ID, objc.Sel("endSpecifier"))
	return rv
}/* debug [instance_properties/getter]: endSpecifier */


// Sets the object specifier representing the last object of the range to a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsrangespecifier/endspecifier
func (r_ RangeSpecifier) SetEndSpecifier(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setEndSpecifier:"), value)
}/* debug [instance_properties/setter]: endSpecifier */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSRangeSpecifier */


