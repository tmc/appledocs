// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSPropertySpecifier */


/* debug [class_header]: Header for NSPropertySpecifier */
// The class instance for the [PropertySpecifier] class.
var (
	PropertySpecifierClass     _PropertySpecifierClass
	PropertySpecifierClassOnce sync.Once
)

func getPropertySpecifierClass() _PropertySpecifierClass {
	PropertySpecifierClassOnce.Do(func() {
		PropertySpecifierClass = _PropertySpecifierClass{objc.GetClass("NSPropertySpecifier")}
	})
	return PropertySpecifierClass
}

type _PropertySpecifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PropertySpecifier */
// An interface definition for the [PropertySpecifier] class.
type IPropertySpecifier interface {
	IScriptObjectSpecifier
	
/* debug [class_interface_properties]: Properties for PropertySpecifier */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PropertySpecifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PropertySpecifier */
// Alloc allocates a new instance without initialization.
func (pc _PropertySpecifierClass) Alloc() PropertySpecifier {
	rv := objc.Send[PropertySpecifier](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PropertySpecifierClass) New() PropertySpecifier {
	rv := objc.Send[PropertySpecifier](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PropertySpecifier) Init() PropertySpecifier {
	rv := objc.Send[PropertySpecifier](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PropertySpecifier) Autorelease() PropertySpecifier {
	rv := objc.Send[PropertySpecifier](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPropertySpecifier creates a new PropertySpecifier instance.
func NewPropertySpecifier() PropertySpecifier {
	return getPropertySpecifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PropertySpecifier */
// A specifier for a simple attribute value, a one-to-one relationship, or all elements of a to-many relationship.
//
// You don’t typically subclass .


// A specifier for a simple attribute value, a one-to-one relationship, or all elements of a to-many relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPropertySpecifier
type PropertySpecifier struct {
	ScriptObjectSpecifier
}

// PropertySpecifierFrom constructs a [PropertySpecifier] from an unsafe.Pointer.
//
// A specifier for a simple attribute value, a one-to-one relationship, or all elements of a to-many relationship.
func PropertySpecifierFrom(ptr unsafe.Pointer) PropertySpecifier {
	return PropertySpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PropertySpecifier *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PropertySpecifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PropertySpecifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PropertySpecifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PropertySpecifier */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPropertySpecifier */



