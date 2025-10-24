// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSPositionalSpecifier */


/* debug [class_header]: Header for NSPositionalSpecifier */
// The class instance for the [PositionalSpecifier] class.
var (
	PositionalSpecifierClass     _PositionalSpecifierClass
	PositionalSpecifierClassOnce sync.Once
)

func getPositionalSpecifierClass() _PositionalSpecifierClass {
	PositionalSpecifierClassOnce.Do(func() {
		PositionalSpecifierClass = _PositionalSpecifierClass{objc.GetClass("NSPositionalSpecifier")}
	})
	return PositionalSpecifierClass
}

type _PositionalSpecifierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PositionalSpecifier */
// An interface definition for the [PositionalSpecifier] class.
type IPositionalSpecifier interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PositionalSpecifier */
	// properties:
	InsertionIndex() int
	InsertionKey() IString
	SetInsertionKey(value IString)
	InsertionReplaces() bool
	SetInsertionReplaces(value bool)
	ObjectSpecifier() IScriptObjectSpecifier
	SetObjectSpecifier(value IScriptObjectSpecifier)
	Position() objectivec.IObject
	SetPosition(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PositionalSpecifier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PositionalSpecifier */
// Alloc allocates a new instance without initialization.
func (pc _PositionalSpecifierClass) Alloc() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PositionalSpecifierClass) New() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PositionalSpecifier) Init() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PositionalSpecifier) Autorelease() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPositionalSpecifier creates a new PositionalSpecifier instance.
func NewPositionalSpecifier() PositionalSpecifier {
	return getPositionalSpecifierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PositionalSpecifier */
// A specifier for an insertion point in a container relative to another object in the container.
//
// Instances of specify an insertion point in a container relative to another object in the container, for example, or . The container is specified by an instance of . objects commonly encapsulate object specifiers used as arguments to the ( ) and commands and indicate where the created or moved object is to be inserted relative to the object represented by an object specifier. Invoking an accessor method to obtain information about an instance of causes the object to be evaluated if it hasn’t been already. You don’t normally subclass .


// A specifier for an insertion point in a container relative to another object in the container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier
type PositionalSpecifier struct {
	objectivec.Object
}

// PositionalSpecifierFrom constructs a [PositionalSpecifier] from an unsafe.Pointer.
//
// A specifier for an insertion point in a container relative to another object in the container.
func PositionalSpecifierFrom(ptr unsafe.Pointer) PositionalSpecifier {
	return PositionalSpecifier{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PositionalSpecifier *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PositionalSpecifier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PositionalSpecifier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PositionalSpecifier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PositionalSpecifier */

// Returns an insertion index that indicates where the new or copied object or objects should be placed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/insertionIndex
func (p_ PositionalSpecifier) InsertionIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("insertionIndex"))
	return rv
}/* debug [instance_properties/getter]: insertionIndex */


// Returns the key that identifies the relationship into which the new or copied object or objects should be inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/insertionkey
func (p_ PositionalSpecifier) InsertionKey() IString {
	rv := objc.Send[String](p_.ID, objc.Sel("insertionKey"))
	return rv
}/* debug [instance_properties/getter]: insertionKey */


// Returns the key that identifies the relationship into which the new or copied object or objects should be inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/insertionkey
func (p_ PositionalSpecifier) SetInsertionKey(value IString) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInsertionKey:"), value)
}/* debug [instance_properties/setter]: insertionKey */


// Returns a Boolean value that indicates whether evaluation has been successful and the object to be inserted should actually replace the keyed, indexed object in the insertion container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/insertionreplaces
func (p_ PositionalSpecifier) InsertionReplaces() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("insertionReplaces"))
	return rv
}/* debug [instance_properties/getter]: insertionReplaces */


// Returns a Boolean value that indicates whether evaluation has been successful and the object to be inserted should actually replace the keyed, indexed object in the insertion container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/insertionreplaces
func (p_ PositionalSpecifier) SetInsertionReplaces(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInsertionReplaces:"), value)
}/* debug [instance_properties/setter]: insertionReplaces */


// Returns the object specifier specified at initialization time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/objectspecifier
func (p_ PositionalSpecifier) ObjectSpecifier() IScriptObjectSpecifier {
	rv := objc.Send[ScriptObjectSpecifier](p_.ID, objc.Sel("objectSpecifier"))
	return rv
}/* debug [instance_properties/getter]: objectSpecifier */


// Returns the object specifier specified at initialization time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/objectspecifier
func (p_ PositionalSpecifier) SetObjectSpecifier(value IScriptObjectSpecifier) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setObjectSpecifier:"), value)
}/* debug [instance_properties/setter]: objectSpecifier */


// Returns the insertion position specified at initialization time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/position
func (p_ PositionalSpecifier) Position() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */


// Returns the insertion position specified at initialization time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspositionalspecifier/position
func (p_ PositionalSpecifier) SetPosition(value objectivec.IObject) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPosition:"), value)
}/* debug [instance_properties/setter]: position */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSPositionalSpecifier */



