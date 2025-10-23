// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PositionalSpecifier] class.
type IPositionalSpecifier interface {
	objectivec.IObject
	Evaluate()
	SetInsertionClassDescription(classDescription IScriptClassDescription)
	InsertionContainer() objc.ID
	InsertionIndex() int
	InsertionKey() string
	InsertionReplaces() bool
	ObjectSpecifier() NSScriptObjectSpecifier
	Position() InsertionPosition
}

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

// Alloc allocates a new instance without initialization.
func (pc _PositionalSpecifierClass) Alloc() PositionalSpecifier {
	rv := objc.Send[PositionalSpecifier](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Initializes a positional specifier with a given position relative to another given specifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/init(position:objectSpecifier:)
func NewPositionalSpecifierWithPositionObjectSpecifier(position NSInsertionPosition, specifier IScriptObjectSpecifier) PositionalSpecifier {
	instance := getPositionalSpecifierClass().Alloc()
	rv := objc.Send[PositionalSpecifier](instance.ID, objc.Sel("initWithPosition:objectSpecifier:"), position, specifier)
	rv.Autorelease()
	return rv
}



// Causes the receiver to evaluate its position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/evaluate()
func (p_ PositionalSpecifier) Evaluate() {
	objc.Send[objc.ID](p_.ID, objc.Sel("evaluate"))
}


// Sets the class description for the object or objects to be inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/setInsertionClassDescription(_:)
func (p_ PositionalSpecifier) SetInsertionClassDescription(classDescription IScriptClassDescription) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setInsertionClassDescription:"), classDescription)
}


// Returns the container in which the new or copied object or objects should be placed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/insertionContainer
func (p_ PositionalSpecifier) InsertionContainer() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("insertionContainer"))
	return rv
}


// Returns an insertion index that indicates where the new or copied object or objects should be placed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/insertionIndex
func (p_ PositionalSpecifier) InsertionIndex() int {
	rv := objc.Send[int](p_.ID, objc.Sel("insertionIndex"))
	return rv
}


// Returns the key that identifies the relationship into which the new or copied object or objects should be inserted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/insertionKey
func (p_ PositionalSpecifier) InsertionKey() string {
	rv := objc.Send[string](p_.ID, objc.Sel("insertionKey"))
	return rv
}


// Returns a Boolean value that indicates whether evaluation has been successful and the object to be inserted should actually replace the keyed, indexed object in the insertion container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/insertionReplaces
func (p_ PositionalSpecifier) InsertionReplaces() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("insertionReplaces"))
	return rv
}


// Returns the object specifier specified at initialization time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/objectSpecifier
func (p_ PositionalSpecifier) ObjectSpecifier() NSScriptObjectSpecifier {
	rv := objc.Send[NSScriptObjectSpecifier](p_.ID, objc.Sel("objectSpecifier"))
	return rv
}


// Returns the insertion position specified at initialization time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/position
func (p_ PositionalSpecifier) Position() InsertionPosition {
	rv := objc.Send[InsertionPosition](p_.ID, objc.Sel("position"))
	return rv
}


