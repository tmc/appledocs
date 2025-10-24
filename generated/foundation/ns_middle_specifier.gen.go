// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSMiddleSpecifier */


/* debug [class_header]: Header for NSMiddleSpecifier */
// The class instance for the [MiddleSpecifier] class.
var (
	MiddleSpecifierClass     _MiddleSpecifierClass
	MiddleSpecifierClassOnce sync.Once
)

func getMiddleSpecifierClass() _MiddleSpecifierClass {
	MiddleSpecifierClassOnce.Do(func() {
		MiddleSpecifierClass = _MiddleSpecifierClass{objc.GetClass("NSMiddleSpecifier")}
	})
	return MiddleSpecifierClass
}

type _MiddleSpecifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MiddleSpecifier */
// An interface definition for the [MiddleSpecifier] class.
type IMiddleSpecifier interface {
	IScriptObjectSpecifier
	
/* debug [class_interface_properties]: Properties for MiddleSpecifier */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MiddleSpecifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MiddleSpecifier */
// Alloc allocates a new instance without initialization.
func (mc _MiddleSpecifierClass) Alloc() MiddleSpecifier {
	rv := objc.Send[MiddleSpecifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MiddleSpecifierClass) New() MiddleSpecifier {
	rv := objc.Send[MiddleSpecifier](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MiddleSpecifier) Init() MiddleSpecifier {
	rv := objc.Send[MiddleSpecifier](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MiddleSpecifier) Autorelease() MiddleSpecifier {
	rv := objc.Send[MiddleSpecifier](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMiddleSpecifier creates a new MiddleSpecifier instance.
func NewMiddleSpecifier() MiddleSpecifier {
	return getMiddleSpecifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MiddleSpecifier */
// A specifier indicating the middle object in a collection or, if not a one-to-many relationship, the sole object.
//
// You don’t typically subclass .


// A specifier indicating the middle object in a collection or, if not a one-to-many relationship, the sole object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMiddleSpecifier
type MiddleSpecifier struct {
	ScriptObjectSpecifier
}

// MiddleSpecifierFrom constructs a [MiddleSpecifier] from an unsafe.Pointer.
//
// A specifier indicating the middle object in a collection or, if not a one-to-many relationship, the sole object.
func MiddleSpecifierFrom(ptr unsafe.Pointer) MiddleSpecifier {
	return MiddleSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MiddleSpecifier *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MiddleSpecifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MiddleSpecifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MiddleSpecifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MiddleSpecifier */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSMiddleSpecifier */



