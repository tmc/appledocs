// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [RandomDistribution] class.
var (
	RandomDistributionClass     _RandomDistributionClass
	RandomDistributionClassOnce sync.Once
)

func getRandomDistributionClass() _RandomDistributionClass {
	RandomDistributionClassOnce.Do(func() {
		RandomDistributionClass = _RandomDistributionClass{objc.GetClass("GKRandomDistribution")}
	})
	return RandomDistributionClass
}

type _RandomDistributionClass struct {
	class objc.Class
}

// An interface definition for the [RandomDistribution] class.
type IRandomDistribution interface {
	objectivec.IObject
	NextBool() bool
	NextInt() int
	NextIntWithUpperBound(upperBound uint) uint
	NextUniform() unsafe.Pointer
}

// A generator for random numbers that fall within a specific range and that exhibit a specific distribution over multiple samplings.
//
// You choose the algorithm that randomizes source values for a distribution by initializing it with an instance of any class that implements the protocol, such as a basic random source (a subclass of ) or another random distribution. The class itself implements a uniform distribution—for more specialized distributions use one of the subclasses and . In a distribution, the probability of generating any number in a specified range (between the values of the distribution’s and properties) is approximately equal. In other words, there is no bias toward any possible outcome. To generate random numbers in this range, use the methods from the protocol listed in Generating Random Numbers below. For more information on choosing and using randomizers in GameplayKit, read in .
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution
type RandomDistribution struct {
	objectivec.Object
}

// RandomDistributionFrom constructs a [RandomDistribution] from an unsafe.Pointer.
//
// A generator for random numbers that fall within a specific range and that exhibit a specific distribution over multiple samplings.
func RandomDistributionFrom(ptr unsafe.Pointer) RandomDistribution {
	return RandomDistribution{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _RandomDistributionClass) Alloc() RandomDistribution {
	rv := objc.Send[RandomDistribution](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _RandomDistributionClass) New() RandomDistribution {
	rv := objc.Send[RandomDistribution](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RandomDistribution) Init() RandomDistribution {
	rv := objc.Send[RandomDistribution](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RandomDistribution) Autorelease() RandomDistribution {
	rv := objc.Send[RandomDistribution](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRandomDistribution creates a new RandomDistribution instance.
func NewRandomDistribution() RandomDistribution {
	return getRandomDistributionClass().New()
}




// Creates a random distribution equivalent to a die with the specified number of sides.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/init(forDieWithSideCount:)
func NewRandomDistributionForDieWithSideCount(sideCount int) RandomDistribution {
	rv := objc.Send[RandomDistribution](objc.ID(getRandomDistributionClass().class), objc.Sel("distributionForDieWithSideCount:"), sideCount)
	return rv
}



// Creates a random distribution with the specified lower and upper bounds, using the Arc4 randomizer.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/init(lowestValue:highestValue:)
func NewRandomDistributionWithLowestValueHighestValue(lowestInclusive int, highestInclusive int) RandomDistribution {
	rv := objc.Send[RandomDistribution](objc.ID(getRandomDistributionClass().class), objc.Sel("distributionWithLowestValue:highestValue:"), lowestInclusive, highestInclusive)
	return rv
}



// Initializes a uniform random distribution with the specified lower and upper bounds, using the specified source randomizer.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/init(randomSource:lowestValue:highestValue:)
func NewRandomDistributionWithRandomSourceLowestValueHighestValue(source objectivec.IObject, lowestInclusive int, highestInclusive int) RandomDistribution {
	instance := getRandomDistributionClass().Alloc()
	rv := objc.Send[RandomDistribution](instance.ID, objc.Sel("initWithRandomSource:lowestValue:highestValue:"), source, lowestInclusive, highestInclusive)
	rv.Autorelease()
	return rv
}


// Creates a random distribution equivalent to a twenty-sided die.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/d20()
func (rc _RandomDistributionClass) D20() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("d20"))
	return rv
}

// Creates a random distribution equivalent to a six-sided die.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/d6()
func (rc _RandomDistributionClass) D6() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("d6"))
	return rv
}

// Creates a random distribution equivalent to a die with the specified number of sides.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/init(forDieWithSideCount:)
func (rc _RandomDistributionClass) DistributionForDieWithSideCount(sideCount int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("distributionForDieWithSideCount:"), sideCount)
	return rv
}

// Creates a random distribution with the specified lower and upper bounds, using the Arc4 randomizer.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/init(lowestValue:highestValue:)
func (rc _RandomDistributionClass) DistributionWithLowestValueHighestValue(lowestInclusive int, highestInclusive int) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(rc.class), objc.Sel("distributionWithLowestValue:highestValue:"), lowestInclusive, highestInclusive)
	return rv
}

// Generates and returns a new random Boolean value within the characteristics of the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/nextBool()
func (r_ RandomDistribution) NextBool() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("nextBool"))
	return rv
}

// Generates and returns a new random integer within the bounds of the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/nextInt()
func (r_ RandomDistribution) NextInt() int {
	rv := objc.Send[int](r_.ID, objc.Sel("nextInt"))
	return rv
}

// Generates and returns a new random integer within the bounds of the distribution and less than the specified limit.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/nextInt(upperBound:)
func (r_ RandomDistribution) NextIntWithUpperBound(upperBound uint) uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("nextIntWithUpperBound:"), upperBound)
	return rv
}

// Generates and returns a new random floating-point value within the characteristics of the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/nextUniform()
func (r_ RandomDistribution) NextUniform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("nextUniform"))
	return rv
}

// The highest value to be produced by the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/highestValue
func (r_ RandomDistribution) HighestValue() int {
	rv := objc.Send[int](r_.ID, objc.Sel("highestValue"))
	return rv
}

// The lowest value to be produced by the distribution.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/lowestValue
func (r_ RandomDistribution) LowestValue() int {
	rv := objc.Send[int](r_.ID, objc.Sel("lowestValue"))
	return rv
}

// The number of unique values the distribution can generate.
//
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKRandomDistribution/numberOfPossibleOutcomes
func (r_ RandomDistribution) NumberOfPossibleOutcomes() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("numberOfPossibleOutcomes"))
	return rv
}


