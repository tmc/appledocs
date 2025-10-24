// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [PHASEPullStreamNode] class.
type IPHASEPullStreamNode interface {
	IPHASEStreamNode
	// properties:
	RenderBlock() unsafe.Pointer
	SetRenderBlock(value unsafe.Pointer)
	RenderHandler() unsafe.Pointer
	SetRenderHandler(value unsafe.Pointer)
	// methods:
}



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

// Alloc allocates a new instance without initialization.
func (pc _PHASEPullStreamNodeClass) Alloc() PHASEPullStreamNode {
	rv := objc.Send[PHASEPullStreamNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNode/renderHandler
func (p_ PHASEPullStreamNode) RenderBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("renderBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPullStreamNode/renderHandler
func (p_ PHASEPullStreamNode) SetRenderBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRenderBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepullstreamnode/renderhandler
func (p_ PHASEPullStreamNode) RenderHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("renderHandler"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/phase/phasepullstreamnode/renderhandler
func (p_ PHASEPullStreamNode) SetRenderHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRenderHandler:"), value)
}



