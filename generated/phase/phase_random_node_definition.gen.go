// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASERandomNodeDefinition */


/* debug [class_header]: Header for PHASERandomNodeDefinition */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASERandomNodeDefinition */
// An interface definition for the [PHASERandomNodeDefinition] class.
type IPHASERandomNodeDefinition interface {
	IPHASESoundEventNodeDefinition
	
/* debug [class_interface_properties]: Properties for PHASERandomNodeDefinition */
	// properties:
	UniqueSelectionQueueLength() int
	SetUniqueSelectionQueueLength(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASERandomNodeDefinition */
	// methods:
	AddSubtreeWeight(subtree IPHASESoundEventNodeDefinition, weight objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASERandomNodeDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASERandomNodeDefinitionClass) Alloc() PHASERandomNodeDefinition {
	rv := objc.Send[PHASERandomNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASERandomNodeDefinition */
// A sound event node that invokes one of its child nodes at random.
//
// When the framework invokes a random node, it passes the invocation on to one of its children at random. The weight you choose for a child node in the argument skews the node’s selection chances.


// A sound event node that invokes one of its child nodes at random.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASERandomNodeDefinition */

// Creates a random node with the name you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASERandomNodeDefinition/init(identifier:)
func NewPHASERandomNodeDefinitionWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) PHASERandomNodeDefinition {
	instance := getPHASERandomNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASERandomNodeDefinition](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASERandomNodeDefinitionWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASERandomNodeDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASERandomNodeDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASERandomNodeDefinition */

// Adds a node tree that’s one of the random-selection options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASERandomNodeDefinition/addSubtree(_:weight:)
func (p_ PHASERandomNodeDefinition) AddSubtreeWeight(subtree IPHASESoundEventNodeDefinition, weight objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addSubtree:weight:"), subtree, weight)
}/* debug [instance_methods/method]: AddSubtreeWeight */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASERandomNodeDefinition */

// The length of the unique selection queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASERandomNodeDefinition/uniqueSelectionQueueLength
func (p_ PHASERandomNodeDefinition) UniqueSelectionQueueLength() int {
	rv := objc.Send[int](p_.ID, objc.Sel("uniqueSelectionQueueLength"))
	return rv
}/* debug [instance_properties/getter]: uniqueSelectionQueueLength */


// The length of the unique selection queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASERandomNodeDefinition/uniqueSelectionQueueLength
func (p_ PHASERandomNodeDefinition) SetUniqueSelectionQueueLength(value int) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUniqueSelectionQueueLength:"), value)
}/* debug [instance_properties/setter]: uniqueSelectionQueueLength */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASERandomNodeDefinition */


