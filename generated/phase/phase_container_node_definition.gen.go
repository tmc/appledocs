// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class PHASEContainerNodeDefinition */


/* debug [class_header]: Header for PHASEContainerNodeDefinition */
// The class instance for the [PHASEContainerNodeDefinition] class.
var (
	PHASEContainerNodeDefinitionClass     _PHASEContainerNodeDefinitionClass
	PHASEContainerNodeDefinitionClassOnce sync.Once
)

func getPHASEContainerNodeDefinitionClass() _PHASEContainerNodeDefinitionClass {
	PHASEContainerNodeDefinitionClassOnce.Do(func() {
		PHASEContainerNodeDefinitionClass = _PHASEContainerNodeDefinitionClass{objc.GetClass("PHASEContainerNodeDefinition")}
	})
	return PHASEContainerNodeDefinitionClass
}

type _PHASEContainerNodeDefinitionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEContainerNodeDefinition */
// An interface definition for the [PHASEContainerNodeDefinition] class.
type IPHASEContainerNodeDefinition interface {
	IPHASESoundEventNodeDefinition
	
/* debug [class_interface_properties]: Properties for PHASEContainerNodeDefinition */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEContainerNodeDefinition */
	// methods:
	AddSubtree(subtree IPHASESoundEventNodeDefinition)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEContainerNodeDefinition */
// Alloc allocates a new instance without initialization.
func (pc _PHASEContainerNodeDefinitionClass) Alloc() PHASEContainerNodeDefinition {
	rv := objc.Send[PHASEContainerNodeDefinition](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEContainerNodeDefinitionClass) New() PHASEContainerNodeDefinition {
	rv := objc.Send[PHASEContainerNodeDefinition](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEContainerNodeDefinition) Init() PHASEContainerNodeDefinition {
	rv := objc.Send[PHASEContainerNodeDefinition](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEContainerNodeDefinition) Autorelease() PHASEContainerNodeDefinition {
	rv := objc.Send[PHASEContainerNodeDefinition](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEContainerNodeDefinition creates a new PHASEContainerNodeDefinition instance.
func NewPHASEContainerNodeDefinition() PHASEContainerNodeDefinition {
	return getPHASEContainerNodeDefinitionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEContainerNodeDefinition */
// A node that plays all its children at the same time.
//
// This node adds structure to the sound event tree while performing no conditional logic or audio playback of its own. By passing invocation to all its children at once, this class invokes the child nodes’ actions simultaneously.


// A node that plays all its children at the same time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEContainerNodeDefinition
type PHASEContainerNodeDefinition struct {
	PHASESoundEventNodeDefinition
}

// PHASEContainerNodeDefinitionFrom constructs a [PHASEContainerNodeDefinition] from an unsafe.Pointer.
//
// A node that plays all its children at the same time.
func PHASEContainerNodeDefinitionFrom(ptr unsafe.Pointer) PHASEContainerNodeDefinition {
	return PHASEContainerNodeDefinition{
		PHASESoundEventNodeDefinition: PHASESoundEventNodeDefinitionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEContainerNodeDefinition */

// Creates a container node with the given name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEContainerNodeDefinition/init(identifier:)
func NewPHASEContainerNodeDefinitionWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) PHASEContainerNodeDefinition {
	instance := getPHASEContainerNodeDefinitionClass().Alloc()
	rv := objc.Send[PHASEContainerNodeDefinition](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPHASEContainerNodeDefinitionWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEContainerNodeDefinition */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEContainerNodeDefinition */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEContainerNodeDefinition */

// Adds a sound event node as a child.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEContainerNodeDefinition/addSubtree(_:)
func (p_ PHASEContainerNodeDefinition) AddSubtree(subtree IPHASESoundEventNodeDefinition) {
	objc.Send[objc.ID](p_.ID, objc.Sel("addSubtree:"), subtree)
}/* debug [instance_methods/method]: AddSubtree */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEContainerNodeDefinition */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEContainerNodeDefinition */


