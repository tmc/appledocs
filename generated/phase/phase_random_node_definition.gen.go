// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASERandomNodeDefinition] class.
var (
	PHASERandomNodeDefinitionClass     _PHASERandomNodeDefinitionClass
	PHASERandomNodeDefinitionClassOnce sync.Once
)

func getPHASERandomNodeDefinitionClass() _PHASERandomNodeDefinitionClass {
	PHASERandomNodeDefinitionClassOnce.Do(func() {
		PHASERandomNodeDefinitionClass = _PHASERandomNodeDefinitionClass{objc.GetClass("PHASERandomNodeDefinition")}
	})
	return PHASERandomNodeDefinitionClass
}

type _PHASERandomNodeDefinitionClass struct {
	class objc.Class
}

// An interface definition for the [PHASERandomNodeDefinition] class.
type IPHASERandomNodeDefinition interface {
	IPHASESoundEventNodeDefinition
	AddSubtreeWeight(subtree unsafe.Pointer, weight unsafe.Pointer)
}

// A sound event node that invokes one of its child nodes at random.
//
// When the framework invokes a random node, it passes the invocation on to one of its children at random. The weight you choose for a child node in the argument skews the node’s selection chances.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASERandomNodeDefinition
type PHASERandomNodeDefinition struct {
	PHASESoundEventNodeDefinition
}

// PHASERandomNodeDefinitionFrom constructs a [PHASERandomNodeDefinition] from an unsafe.Pointer.
//
// A sound event node that invokes one of its child nodes at random.
func PHASERandomNodeDefinitionFrom(ptr unsafe.Pointer) PHASERandomNodeDefinition {
	return PHASERandomNodeDefinition{
		PHASESoundEventNodeDefinition: PHASESoundEventNodeDefinitionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASERandomNodeDefinitionClass) Alloc() PHASERandomNodeDefinition {
	rv := objc.Send[PHASERandomNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASERandomNodeDefinitionClass) New() PHASERandomNodeDefinition {
	rv := objc.Send[PHASERandomNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASERandomNodeDefinition) Init() PHASERandomNodeDefinition {
	rv := objc.Send[PHASERandomNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASERandomNodeDefinition) Autorelease() PHASERandomNodeDefinition {
	rv := objc.Send[PHASERandomNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASERandomNodeDefinition creates a new PHASERandomNodeDefinition instance.
func NewPHASERandomNodeDefinition() PHASERandomNodeDefinition {
	return getPHASERandomNodeDefinitionClass().New()
}




// Creates a random node with the name you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASERandomNodeDefinition/init(identifier:)
func NewPHASERandomNodeDefinitionWithIdentifier(identifier string) PHASERandomNodeDefinition {
	instance := getPHASERandomNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASERandomNodeDefinition](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	rv.Autorelease()
	return rv
}


// Adds a node tree that’s one of the random-selection options.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASERandomNodeDefinition/addSubtree(_:weight:)
func (p_ PHASERandomNodeDefinition) AddSubtreeWeight(subtree unsafe.Pointer, weight unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addSubtree:weight:"), subtree, weight)
}

// The length of the unique selection queue.
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASERandomNodeDefinition/uniqueSelectionQueueLength
func (p_ PHASERandomNodeDefinition) UniqueSelectionQueueLength() int {
	rv := objc.Send[int](p_.ID, objc.Sel("uniqueSelectionQueueLength"))
	return rv
}


// SetUniqueSelectionQueueLength sets the value of the uniqueSelectionQueueLength property.
// The length of the unique selection queue.

//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASERandomNodeDefinition/uniqueSelectionQueueLength
func (p_ PHASERandomNodeDefinition) SetUniqueSelectionQueueLength(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUniqueSelectionQueueLength:"), value)
}


