// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [UnitMass] class.
var (
	UnitMassClass     _UnitMassClass
	UnitMassClassOnce sync.Once
)

func getUnitMassClass() _UnitMassClass {
	UnitMassClassOnce.Do(func() {
		UnitMassClass = _UnitMassClass{objc.GetClass("NSUnitMass")}
	})
	return UnitMassClass
}

type _UnitMassClass struct {
	class objc.Class
}

// An interface definition for the [UnitMass] class.
type IUnitMass interface {
	IDimension
}

// A unit of measure for mass.
//
// You typically use instances of to represent specific quantities of mass using the class.


// A unit of measure for mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass
type UnitMass struct {
	Dimension
}

// UnitMassFrom constructs a [UnitMass] from an unsafe.Pointer.
//
// A unit of measure for mass.
func UnitMassFrom(ptr unsafe.Pointer) UnitMass {
	return UnitMass{
		Dimension: DimensionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (uc _UnitMassClass) Alloc() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _UnitMassClass) New() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ UnitMass) Init() UnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ UnitMass) Autorelease() UnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitMass creates a new UnitMass instance.
func NewUnitMass() UnitMass {
	return getUnitMassClass().New()
}



// The carats unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/carats
func (uc _UnitMassClass) Carats() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("carats"))
	return rv
}

// The centigrams unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/centigrams
func (uc _UnitMassClass) Centigrams() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("centigrams"))
	return rv
}

// The decigrams unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/decigrams
func (uc _UnitMassClass) Decigrams() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("decigrams"))
	return rv
}

// The grams unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/grams
func (uc _UnitMassClass) Grams() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("grams"))
	return rv
}

// The kilograms unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/kilograms
func (uc _UnitMassClass) Kilograms() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("kilograms"))
	return rv
}

// The metric tons unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/metricTons
func (uc _UnitMassClass) MetricTons() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("metricTons"))
	return rv
}

// The micrograms unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/micrograms
func (uc _UnitMassClass) Micrograms() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("micrograms"))
	return rv
}

// The milligrams unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/milligrams
func (uc _UnitMassClass) Milligrams() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("milligrams"))
	return rv
}

// The nanograms unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/nanograms
func (uc _UnitMassClass) Nanograms() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("nanograms"))
	return rv
}

// The ounces unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/ounces
func (uc _UnitMassClass) Ounces() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("ounces"))
	return rv
}

// The ounces troy unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/ouncesTroy
func (uc _UnitMassClass) OuncesTroy() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("ouncesTroy"))
	return rv
}

// The picograms unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/picograms
func (uc _UnitMassClass) Picograms() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("picograms"))
	return rv
}

// The pounds unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/pounds
func (uc _UnitMassClass) PoundsMass() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("poundsMass"))
	return rv
}

// The short tons unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/shortTons
func (uc _UnitMassClass) ShortTons() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("shortTons"))
	return rv
}

// The slugs unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/slugs
func (uc _UnitMassClass) Slugs() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("slugs"))
	return rv
}

// The stone unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/stones
func (uc _UnitMassClass) Stones() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("stones"))
	return rv
}

// The carats unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/carats
func (u_ UnitMass) Carats() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("carats"))
	return rv
}


// The centigrams unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/centigrams
func (u_ UnitMass) Centigrams() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("centigrams"))
	return rv
}


// The decigrams unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/decigrams
func (u_ UnitMass) Decigrams() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("decigrams"))
	return rv
}


// The grams unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/grams
func (u_ UnitMass) Grams() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("grams"))
	return rv
}


// The kilograms unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/kilograms
func (u_ UnitMass) Kilograms() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("kilograms"))
	return rv
}


// The metric tons unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/metricTons
func (u_ UnitMass) MetricTons() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("metricTons"))
	return rv
}


// The micrograms unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/micrograms
func (u_ UnitMass) Micrograms() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("micrograms"))
	return rv
}


// The milligrams unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/milligrams
func (u_ UnitMass) Milligrams() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("milligrams"))
	return rv
}


// The nanograms unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/nanograms
func (u_ UnitMass) Nanograms() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("nanograms"))
	return rv
}


// The ounces unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/ounces
func (u_ UnitMass) Ounces() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("ounces"))
	return rv
}


// The ounces troy unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/ouncesTroy
func (u_ UnitMass) OuncesTroy() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("ouncesTroy"))
	return rv
}


// The picograms unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/picograms
func (u_ UnitMass) Picograms() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("picograms"))
	return rv
}


// The pounds unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/pounds
func (u_ UnitMass) PoundsMass() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("poundsMass"))
	return rv
}


// The short tons unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/shortTons
func (u_ UnitMass) ShortTons() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("shortTons"))
	return rv
}


// The slugs unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/slugs
func (u_ UnitMass) Slugs() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("slugs"))
	return rv
}


// The stone unit of mass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/UnitMass/stones
func (u_ UnitMass) Stones() IUnitMass {
	rv := objc.Send[UnitMass](u_.ID, objc.Sel("stones"))
	return rv
}



