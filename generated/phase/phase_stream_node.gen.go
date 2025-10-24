// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfaudio"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PHASEStreamNode] class.
var (
	PHASEStreamNodeClass     _PHASEStreamNodeClass
	PHASEStreamNodeClassOnce sync.Once
)

func getPHASEStreamNodeClass() _PHASEStreamNodeClass {
	PHASEStreamNodeClassOnce.Do(func() {
		PHASEStreamNodeClass = _PHASEStreamNodeClass{objc.GetClass("PHASEStreamNode")}
	})
	return PHASEStreamNodeClass
}

type _PHASEStreamNodeClass struct {
	class objc.Class
}

// An interface definition for the [PHASEStreamNode] class.
type IPHASEStreamNode interface {
	objectivec.IObject
	// properties:
	Format() objc.IObject /* cross-framework: AudioFormat */
	GainMetaParameter() IPHASENumberMetaParameter
	Mixer() IPHASEMixer
	RateMetaParameter() IPHASENumberMetaParameter
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStreamNode
type PHASEStreamNode struct {
	objectivec.Object
}

// PHASEStreamNodeFrom constructs a [PHASEStreamNode] from an unsafe.Pointer.
func PHASEStreamNodeFrom(ptr unsafe.Pointer) PHASEStreamNode {
	return PHASEStreamNode{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PHASEStreamNodeClass) Alloc() PHASEStreamNode {
	rv := objc.Send[PHASEStreamNode](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PHASEStreamNodeClass) New() PHASEStreamNode {
	rv := objc.Send[PHASEStreamNode](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PHASEStreamNode) Init() PHASEStreamNode {
	rv := objc.Send[PHASEStreamNode](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PHASEStreamNode) Autorelease() PHASEStreamNode {
	rv := objc.Send[PHASEStreamNode](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPHASEStreamNode creates a new PHASEStreamNode instance.
func NewPHASEStreamNode() PHASEStreamNode {
	return getPHASEStreamNodeClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStreamNode/format
func (p_ PHASEStreamNode) Format() objc.IObject /* cross-framework: AudioFormat */ {
	rv := objc.Send[avfaudio.AudioFormat](p_.ID, objc.Sel("format"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStreamNode/gainMetaParameter
func (p_ PHASEStreamNode) GainMetaParameter() IPHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](p_.ID, objc.Sel("gainMetaParameter"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStreamNode/mixer
func (p_ PHASEStreamNode) Mixer() IPHASEMixer {
	rv := objc.Send[PHASEMixer](p_.ID, objc.Sel("mixer"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/PHASE/PHASEStreamNode/rateMetaParameter
func (p_ PHASEStreamNode) RateMetaParameter() IPHASENumberMetaParameter {
	rv := objc.Send[PHASENumberMetaParameter](p_.ID, objc.Sel("rateMetaParameter"))
	return rv
}



