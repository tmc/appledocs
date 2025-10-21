// Code generated from Apple documentation for SoundAnalysis. DO NOT EDIT.

package soundanalysis

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SNClassification] class.
var (
	SNClassificationClass     _SNClassificationClass
	SNClassificationClassOnce sync.Once
)

func getSNClassificationClass() _SNClassificationClass {
	SNClassificationClassOnce.Do(func() {
		SNClassificationClass = _SNClassificationClass{objc.GetClass("SNClassification")}
	})
	return SNClassificationClass
}

type _SNClassificationClass struct {
	class objc.Class
}

// An interface definition for the [SNClassification] class.
type ISNClassification interface {
	objectivec.IObject
}

// A type that pairs a sound classifier’s prediction with its confidence in that prediction.
//
// An represents a single sound classification prediction, and the sound classifier model’s confidence in that prediction.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassification
type SNClassification struct {
	objectivec.Object
}

// SNClassificationFrom constructs a [SNClassification] from an unsafe.Pointer.
//
// A type that pairs a sound classifier’s prediction with its confidence in that prediction.
func SNClassificationFrom(ptr unsafe.Pointer) SNClassification {
	return SNClassification{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SNClassificationClass) Alloc() SNClassification {
	rv := objc.Send[SNClassification](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SNClassificationClass) New() SNClassification {
	rv := objc.Send[SNClassification](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SNClassification) Init() SNClassification {
	rv := objc.Send[SNClassification](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SNClassification) Autorelease() SNClassification {
	rv := objc.Send[SNClassification](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSNClassification creates a new SNClassification instance.
func NewSNClassification() SNClassification {
	return getSNClassificationClass().New()
}


// A prediction label that’s one of the classifications a sound classifier’s underlying model defines.
//
// [Full Topic]: https://developer.apple.com/documentation/SoundAnalysis/SNClassification/identifier
func (s_ SNClassification) Identifier() string {
	rv := objc.Send[string](s_.ID, objc.Sel("identifier"))
	return rv
}



