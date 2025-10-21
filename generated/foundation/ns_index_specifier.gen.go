// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [IndexSpecifier] class.
type IIndexSpecifier interface {
	IScriptObjectSpecifier
}

// A specifier representing an object in a collection (or container) with an index number.
//
// The script terms and specify the object with index , while specifies the object with index of . A negative index indicates a location by counting backward from the last object in the collection. You don’t normally subclass .
//
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

// Alloc allocates a new instance without initialization.
func (ic _IndexSpecifierClass) Alloc() IndexSpecifier {
	rv := objc.Send[IndexSpecifier](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes an allocated object with a class description, container specifier, collection key, and object index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier/init(containerClassDescription:containerSpecifier:key:index:)
func NewIndexSpecifierWithContainerClassDescriptionContainerSpecifierKeyIndex(classDesc IScriptClassDescription, container IScriptObjectSpecifier, property appkit.string, index int) IndexSpecifier {
	instance := getIndexSpecifierClass().Alloc()
	rv := objc.Send[IndexSpecifier](instance.ID, objc.Sel("initWithContainerClassDescription:containerSpecifier:key:index:"), classDesc, container, property, index)
	rv.Autorelease()
	return rv
}


// Sets the value of the receiver’s property.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier/index
func (i_ IndexSpecifier) Index() int {
	rv := objc.Send[int](i_.ID, objc.Sel("index"))
	return rv
}


// SetIndex sets the value of the index property.
// Sets the value of the receiver’s property.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSpecifier/index
func (i_ IndexSpecifier) SetIndex(value int) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setIndex:"), value)
}


