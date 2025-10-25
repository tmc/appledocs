// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSIndexSpecifier */


/* debug [class_header]: Header for NSIndexSpecifier */
// The class instance for the [IndexSpecifier] class.
var (
	IndexSpecifierClass     _IndexSpecifierClass
	IndexSpecifierClassOnce sync.Once
)

func getIndexSpecifierClass() _IndexSpecifierClass {
	IndexSpecifierClassOnce.Do(func() {
		IndexSpecifierClass = _IndexSpecifierClass{objc.GetClass("NSIndexSpecifier")}
	})
	return IndexSpecifierClass
}

type _IndexSpecifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for IndexSpecifier */
// An interface definition for the [IndexSpecifier] class.
type IIndexSpecifier interface {
	IScriptObjectSpecifier
	
/* debug [class_interface_properties]: Properties for IndexSpecifier */
	// properties:
	Index() int
	SetIndex(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for IndexSpecifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for IndexSpecifier */
// Alloc allocates a new instance without initialization.
func (ic _IndexSpecifierClass) Alloc() IndexSpecifier {
	rv := objc.Send[IndexSpecifier](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ic _IndexSpecifierClass) New() IndexSpecifier {
	rv := objc.Send[IndexSpecifier](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IndexSpecifier) Init() IndexSpecifier {
	rv := objc.Send[IndexSpecifier](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IndexSpecifier) Autorelease() IndexSpecifier {
	rv := objc.Send[IndexSpecifier](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIndexSpecifier creates a new IndexSpecifier instance.
func NewIndexSpecifier() IndexSpecifier {
	return getIndexSpecifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for IndexSpecifier */
// A specifier representing an object in a collection (or container) with an index number.
//
// The script terms and specify the object with index , while specifies the object with index of . A negative index indicates a location by counting backward from the last object in the collection. You don’t normally subclass .


// A specifier representing an object in a collection (or container) with an index number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier
type IndexSpecifier struct {
	ScriptObjectSpecifier
}

// IndexSpecifierFrom constructs a [IndexSpecifier] from an unsafe.Pointer.
//
// A specifier representing an object in a collection (or container) with an index number.
func IndexSpecifierFrom(ptr unsafe.Pointer) IndexSpecifier {
	return IndexSpecifier{
		ScriptObjectSpecifier: ScriptObjectSpecifierFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for IndexSpecifier */

// Initializes an allocated object with a class description, container specifier, collection key, and object index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier/init(containerClassDescription:containerSpecifier:key:index:)
func NewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKeyIndex(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property IString, index int) IndexSpecifier {
	instance := getIndexSpecifierClass().Alloc()
	rv := objc.Send[IndexSpecifier](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:index:"), classDesc, container, property, index)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKeyIndex */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for IndexSpecifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for IndexSpecifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for IndexSpecifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for IndexSpecifier */

// Sets the value of the receiver’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexspecifier/index
func (i_ IndexSpecifier) Index() int {
	rv := objc.Send[int](i_.ID, objc.Sel("index"))
	return rv
}/* debug [instance_properties/getter]: index */


// Sets the value of the receiver’s
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsindexspecifier/index
func (i_ IndexSpecifier) SetIndex(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIndex:"), value)
}/* debug [instance_properties/setter]: index */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSIndexSpecifier */


