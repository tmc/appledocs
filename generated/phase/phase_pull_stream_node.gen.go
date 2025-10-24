// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class PHASEPullStreamNode */


/* debug [class_header]: Header for PHASEPullStreamNode */
// The class instance for the [PHASEPullStreamNode] class.
var (
	PHASEPullStreamNodeClass     _PHASEPullStreamNodeClass
	PHASEPullStreamNodeClassOnce sync.Once
)

func getPHASEPullStreamNodeClass() _PHASEPullStreamNodeClass {
	PHASEPullStreamNodeClassOnce.Do(func() {
		PHASEPullStreamNodeClass = _PHASEPullStreamNodeClass{objc.GetClass("PHASEPullStreamNode")}
	})
	return PHASEPullStreamNodeClass
}

type _PHASEPullStreamNodeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PHASEPullStreamNode */
// An interface definition for the [PHASEPullStreamNode] class.
type IPHASEPullStreamNode interface {
	IPHASEStreamNode
	
/* debug [class_interface_properties]: Properties for PHASEPullStreamNode */
	// properties:
	RenderBlock() unsafe.Pointer
	SetRenderBlock(value unsafe.Pointer)
	RenderHandler() unsafe.Pointer
	SetRenderHandler(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PHASEPullStreamNode */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PHASEPullStreamNode */
// Alloc allocates a new instance without initialization.
func (pc _PHASEPullStreamNodeClass) Alloc() PHASEPullStreamNode {
	rv := objc.Send[PHASEPullStreamNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PHASEPullStreamNodeClass) New() PHASEPullStreamNode {
	rv := objc.Send[PHASEPullStreamNode](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEPullStreamNode) Init() PHASEPullStreamNode {
	rv := objc.Send[PHASEPullStreamNode](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEPullStreamNode) Autorelease() PHASEPullStreamNode {
	rv := objc.Send[PHASEPullStreamNode](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEPullStreamNode creates a new PHASEPullStreamNode instance.
func NewPHASEPullStreamNode() PHASEPullStreamNode {
	return getPHASEPullStreamNodeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PHASEPullStreamNode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNode
type PHASEPullStreamNode struct {
	PHASEStreamNode
}

// PHASEPullStreamNodeFrom constructs a [PHASEPullStreamNode] from an unsafe.Pointer.
func PHASEPullStreamNodeFrom(ptr unsafe.Pointer) PHASEPullStreamNode {
	return PHASEPullStreamNode{
		PHASEStreamNode: PHASEStreamNodeFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PHASEPullStreamNode *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PHASEPullStreamNode */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PHASEPullStreamNode */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PHASEPullStreamNode */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PHASEPullStreamNode */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNode/renderHandler
func (p_ PHASEPullStreamNode) RenderBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("renderBlock"))
	return rv
}/* debug [instance_properties/getter]: renderBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNode/renderHandler
func (p_ PHASEPullStreamNode) SetRenderBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRenderBlock:"), value)
}/* debug [instance_properties/setter]: renderBlock */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepullstreamnode/renderhandler
func (p_ PHASEPullStreamNode) RenderHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("renderHandler"))
	return rv
}/* debug [instance_properties/getter]: renderHandler */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepullstreamnode/renderhandler
func (p_ PHASEPullStreamNode) SetRenderHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRenderHandler:"), value)
}/* debug [instance_properties/setter]: renderHandler */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class PHASEPullStreamNode */



