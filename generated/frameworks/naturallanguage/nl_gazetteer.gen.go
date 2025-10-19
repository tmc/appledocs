// Code generated from Apple documentation for NaturalLanguage. DO NOT EDIT.

package naturallanguage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NLGazetteer] class.
var nLGazetteerClass = _NLGazetteerClass{objc.GetClass("NLGazetteer")}

type _NLGazetteerClass struct {
	class objc.Class
}

// An interface definition for the [NLGazetteer] class.
type INLGazetteer interface {
	objectivec.IObject
}

// A collection of terms and their labels, which take precedence over a word tagger. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/NaturalLanguage/NLGazetteer

type NLGazetteer struct {
	objectivec.Object
}

// NLGazetteerFrom constructs a [NLGazetteer] from an unsafe.Pointer.
//
// A collection of terms and their labels, which take precedence over a word tagger.
func NLGazetteerFrom(ptr unsafe.Pointer) NLGazetteer {
	return NLGazetteer{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (nc _NLGazetteerClass) Alloc() NLGazetteer {
	rv := objc.Send[NLGazetteer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (nc _NLGazetteerClass) New() NLGazetteer {
	rv := objc.Send[NLGazetteer](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NLGazetteer) Init() NLGazetteer {
	rv := objc.Send[NLGazetteer](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NLGazetteer) Autorelease() NLGazetteer {
	rv := objc.Send[NLGazetteer](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNLGazetteer creates a new NLGazetteer instance.
func NewNLGazetteer() NLGazetteer {
	return nLGazetteerClass.New()
}




