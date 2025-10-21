// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PHASEPushStreamNode] class.
var (
	PHASEPushStreamNodeClass     _PHASEPushStreamNodeClass
	PHASEPushStreamNodeClassOnce sync.Once
)

func getPHASEPushStreamNodeClass() _PHASEPushStreamNodeClass {
	PHASEPushStreamNodeClassOnce.Do(func() {
		PHASEPushStreamNodeClass = _PHASEPushStreamNodeClass{objc.GetClass("PHASEPushStreamNode")}
	})
	return PHASEPushStreamNodeClass
}

type _PHASEPushStreamNodeClass struct {
	class objc.Class
}

// An interface definition for the [PHASEPushStreamNode] class.
type IPHASEPushStreamNode interface {
	IPHASEStreamNode
}

// An audio stream you manage to provide a sound buffer data.
//
// A sound event’s dictionary populates with an instance of this class when PHASE invokes a in your event node tree. Your app provides the audio data that the sound event plays by calling one or more of this class’s buffer-scheduling functions, for example, .
//
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEPushStreamNode
type PHASEPushStreamNode struct {
	PHASEStreamNode
}

// PHASEPushStreamNodeFrom constructs a [PHASEPushStreamNode] from an unsafe.Pointer.
//
// An audio stream you manage to provide a sound buffer data.
func PHASEPushStreamNodeFrom(ptr unsafe.Pointer) PHASEPushStreamNode {
	return PHASEPushStreamNode{
		PHASEStreamNode: PHASEStreamNodeFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEPushStreamNodeClass) Alloc() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEPushStreamNodeClass) New() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEPushStreamNode) Init() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEPushStreamNode) Autorelease() PHASEPushStreamNode {
	rv := objc.Send[PHASEPushStreamNode](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEPushStreamNode creates a new PHASEPushStreamNode instance.
func NewPHASEPushStreamNode() PHASEPushStreamNode {
	return getPHASEPushStreamNodeClass().New()
}




