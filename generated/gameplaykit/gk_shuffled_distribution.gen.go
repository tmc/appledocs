// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GKShuffledDistribution */


/* debug [class_header]: Header for GKShuffledDistribution */
// The class instance for the [ShuffledDistribution] class.
var (
	ShuffledDistributionClass     _ShuffledDistributionClass
	ShuffledDistributionClassOnce sync.Once
)

func getShuffledDistributionClass() _ShuffledDistributionClass {
	ShuffledDistributionClassOnce.Do(func() {
		ShuffledDistributionClass = _ShuffledDistributionClass{objc.GetClass("GKShuffledDistribution")}
	})
	return ShuffledDistributionClass
}

type _ShuffledDistributionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ShuffledDistribution */
// An interface definition for the [ShuffledDistribution] class.
type IShuffledDistribution interface {
	IRandomDistribution
	
/* debug [class_interface_properties]: Properties for ShuffledDistribution */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ShuffledDistribution */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ShuffledDistribution */
// Alloc allocates a new instance without initialization.
func (sc _ShuffledDistributionClass) Alloc() ShuffledDistribution {
	rv := objc.Send[ShuffledDistribution](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _ShuffledDistributionClass) New() ShuffledDistribution {
	rv := objc.Send[ShuffledDistribution](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ShuffledDistribution) Init() ShuffledDistribution {
	rv := objc.Send[ShuffledDistribution](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ShuffledDistribution) Autorelease() ShuffledDistribution {
	rv := objc.Send[ShuffledDistribution](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewShuffledDistribution creates a new ShuffledDistribution instance.
func NewShuffledDistribution() ShuffledDistribution {
	return getShuffledDistributionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ShuffledDistribution */
// A generator for random numbers that are uniformly distributed across many samplings, but where short sequences of similar values are unlikely.
//
// The behavior of a shuffled distribution is sometimes called “fair” randomization, because true randomness in games can result in extended “lucky streaks” or “unlucky streaks” for players. To create a shuffled distribution and use it to generate random numbers, use the methods defined by its superclass . The class inherits its entire interface from its superclass—to initialize and use a shuffled distribution, use the methods listed in . A shuffled distribution differs from its superclass in behavior only. Consider the code snippets below: In this example, each distribution generates 100 random integers from a simulated six-sided die. In both cases, the distribution of results is roughly uniform—that is, the number of occurrences of any specific value is about the same as that of any other value. However, the shuffled distribution makes sure not to repeat any one value until it has used all of its possible values. In this example, if the die rolls a 1, the shuffled distribution will not generate another 1 for at least five more rolls. For more information on choosing and using randomizers in GameplayKit, read in .


// A generator for random numbers that are uniformly distributed across many samplings, but where short sequences of similar values are unlikely.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameplayKit/GKShuffledDistribution
type ShuffledDistribution struct {
	RandomDistribution
}

// ShuffledDistributionFrom constructs a [ShuffledDistribution] from an unsafe.Pointer.
//
// A generator for random numbers that are uniformly distributed across many samplings, but where short sequences of similar values are unlikely.
func ShuffledDistributionFrom(ptr unsafe.Pointer) ShuffledDistribution {
	return ShuffledDistribution{
		RandomDistribution: RandomDistributionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ShuffledDistribution *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ShuffledDistribution */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ShuffledDistribution */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ShuffledDistribution */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ShuffledDistribution */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKShuffledDistribution */



