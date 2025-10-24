// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSRandomSpecifier */


/* debug [class_header]: Header for NSRandomSpecifier */
// The class instance for the [RandomSpecifier] class.
var (
	RandomSpecifierClass     _RandomSpecifierClass
	RandomSpecifierClassOnce sync.Once
)

func getRandomSpecifierClass() _RandomSpecifierClass {
	RandomSpecifierClassOnce.Do(func() {
		RandomSpecifierClass = _RandomSpecifierClass{objc.GetClass("NSRandomSpecifier")}
	})
	return RandomSpecifierClass
}

type _RandomSpecifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RandomSpecifier */
// An interface definition for the [RandomSpecifier] class.
type IRandomSpecifier interface {
	IScriptObjectSpecifier
	
/* debug [class_interface_properties]: Properties for RandomSpecifier */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RandomSpecifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RandomSpecifier */
// Alloc allocates a new instance without initialization.
func (rc _RandomSpecifierClass) Alloc() RandomSpecifier {
	rv := objc.Send[RandomSpecifier](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RandomSpecifierClass) New() RandomSpecifier {
	rv := objc.Send[RandomSpecifier](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RandomSpecifier) Init() RandomSpecifier {
	rv := objc.Send[RandomSpecifier](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RandomSpecifier) Autorelease() RandomSpecifier {
	rv := objc.Send[RandomSpecifier](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRandomSpecifier creates a new RandomSpecifier instance.
func NewRandomSpecifier() RandomSpecifier {
	return getRandomSpecifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RandomSpecifier */
// A specifier for an arbitrary object in a collection or, if not a one-to-many relationship, the sole object.


// A specifier for an arbitrary object in a collection or, if not a one-to-many relationship, the sole object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRandomSpecifier
type RandomSpecifier struct {
	ScriptObjectSpecifier
}

// RandomSpecifierFrom constructs a [RandomSpecifier] from an unsafe.Pointer.
//
// A specifier for an arbitrary object in a collection or, if not a one-to-many relationship, the sole object.
func RandomSpecifierFrom(ptr unsafe.Pointer) RandomSpecifier {
	return RandomSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RandomSpecifier *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RandomSpecifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RandomSpecifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RandomSpecifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RandomSpecifier */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSRandomSpecifier */



