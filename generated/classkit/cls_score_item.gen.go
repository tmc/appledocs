// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CLSScoreItem */


/* debug [class_header]: Header for CLSScoreItem */
// The class instance for the [SScoreItem] class.
var (
	SScoreItemClass     _SScoreItemClass
	SScoreItemClassOnce sync.Once
)

func getSScoreItemClass() _SScoreItemClass {
	SScoreItemClassOnce.Do(func() {
		SScoreItemClass = _SScoreItemClass{objc.GetClass("CLSScoreItem")}
	})
	return SScoreItemClass
}

type _SScoreItemClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SScoreItem */
// An interface definition for the [SScoreItem] class.
type ISScoreItem interface {
	ISActivityItem
	
/* debug [class_interface_properties]: Properties for SScoreItem */
	// properties:
	MaxScore() float64
	SetMaxScore(value float64)
	Score() float64
	SetScore(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SScoreItem */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SScoreItem */
// Alloc allocates a new instance without initialization.
func (sc _SScoreItemClass) Alloc() SScoreItem {
	rv := objc.Send[SScoreItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SScoreItemClass) New() SScoreItem {
	rv := objc.Send[SScoreItem](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SScoreItem) Init() SScoreItem {
	rv := objc.Send[SScoreItem](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SScoreItem) Autorelease() SScoreItem {
	rv := objc.Send[SScoreItem](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSScoreItem creates a new SScoreItem instance.
func NewSScoreItem() SScoreItem {
	return getSScoreItemClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SScoreItem */
// Activity information that signifies a score out of a possible maximum.
//
// Use an activity item of this type to indicate the relative success in completing a task, like the number of correctly answered questions on a quiz.


// Activity information that signifies a score out of a possible maximum.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem
type SScoreItem struct {
	SActivityItem
}

// SScoreItemFrom constructs a [SScoreItem] from an unsafe.Pointer.
//
// Activity information that signifies a score out of a possible maximum.
func SScoreItemFrom(ptr unsafe.Pointer) SScoreItem {
	return SScoreItem{
		SActivityItem: SActivityItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SScoreItem */

// Initializes an activity item that holds a score value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem/init(identifier:title:score:maxScore:)
func NewSScoreItemWithIdentifierTitleScoreMaxScore(identifier objc.IObject /* cross-framework: NSString */, title objc.IObject /* cross-framework: NSString */, score float64, maxScore float64) SScoreItem {
	instance := getSScoreItemClass().Alloc()
	rv := objc.Send[SScoreItem](instance.ID, objc.Sel("initWithIdentifier:title:score:maxScore:"), identifier, title, score, maxScore)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSScoreItemWithIdentifierTitleScoreMaxScore */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SScoreItem */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SScoreItem */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SScoreItem */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SScoreItem */

// The maximum possible score that the user can earn on a given task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem/maxScore
func (s_ SScoreItem) MaxScore() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("maxScore"))
	return rv
}/* debug [instance_properties/getter]: maxScore */


// The maximum possible score that the user can earn on a given task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem/maxScore
func (s_ SScoreItem) SetMaxScore(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxScore:"), value)
}/* debug [instance_properties/setter]: maxScore */


// The score earned by a user in completing the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem/score
func (s_ SScoreItem) Score() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("score"))
	return rv
}/* debug [instance_properties/getter]: score */


// The score earned by a user in completing the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem/score
func (s_ SScoreItem) SetScore(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScore:"), value)
}/* debug [instance_properties/setter]: score */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CLSScoreItem */


