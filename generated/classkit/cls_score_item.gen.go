// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [SScoreItem] class.
type ISScoreItem interface {
	ISActivityItem
}

// Activity information that signifies a score out of a possible maximum.
//
// Use an activity item of this type to indicate the relative success in completing a task, like the number of correctly answered questions on a quiz.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SScoreItemClass) Alloc() SScoreItem {
	rv := objc.Send[SScoreItem](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Initializes an activity item that holds a score value.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem/init(identifier:title:score:maxScore:)
func NewSScoreItemWithIdentifierTitleScoreMaxScore(identifier string, title string, score unsafe.Pointer, maxScore unsafe.Pointer) SScoreItem {
	instance := getSScoreItemClass().Alloc()
	rv := objc.Send[SScoreItem](instance.ID, objc.Sel("initWithIdentifier:title:score:maxScore:"), objc.String(identifier), objc.String(title), score, maxScore)
	rv.Autorelease()
	return rv
}


// The maximum possible score that the user can earn on a given task.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem/maxScore
func (s_ SScoreItem) MaxScore() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("maxScore"))
	return rv
}


// SetMaxScore sets the value of the maxScore property.
// The maximum possible score that the user can earn on a given task.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem/maxScore
func (s_ SScoreItem) SetMaxScore(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMaxScore:"), value)
}
// The score earned by a user in completing the task.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem/score
func (s_ SScoreItem) Score() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("score"))
	return rv
}


// SetScore sets the value of the score property.
// The score earned by a user in completing the task.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSScoreItem/score
func (s_ SScoreItem) SetScore(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setScore:"), value)
}

