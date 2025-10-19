// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKTileGroupRule] class.
var (
	sKTileGroupRuleClass     _SKTileGroupRuleClass
	sKTileGroupRuleClassOnce sync.Once
)

func getSKTileGroupRuleClass() _SKTileGroupRuleClass {
	sKTileGroupRuleClassOnce.Do(func() {
		sKTileGroupRuleClass = _SKTileGroupRuleClass{objc.GetClass("SKTileGroupRule")}
	})
	return sKTileGroupRuleClass
}

type _SKTileGroupRuleClass struct {
	class objc.Class
}

// An interface definition for the [SKTileGroupRule] class.
type ISKTileGroupRule interface {
	objectivec.IObject
}

// Rules that describe how various tiles should be placed in a map.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule
type SKTileGroupRule struct {
	objectivec.Object
}

// SKTileGroupRuleFrom constructs a [SKTileGroupRule] from an unsafe.Pointer.
//
// Rules that describe how various tiles should be placed in a map.
func SKTileGroupRuleFrom(ptr unsafe.Pointer) SKTileGroupRule {
	return SKTileGroupRule{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKTileGroupRuleClass) Alloc() SKTileGroupRule {
	rv := objc.Send[SKTileGroupRule](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKTileGroupRuleClass) New() SKTileGroupRule {
	rv := objc.Send[SKTileGroupRule](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKTileGroupRule) Init() SKTileGroupRule {
	rv := objc.Send[SKTileGroupRule](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKTileGroupRule) Autorelease() SKTileGroupRule {
	rv := objc.Send[SKTileGroupRule](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTileGroupRule creates a new SKTileGroupRule instance.
func NewSKTileGroupRule() SKTileGroupRule {
	return getSKTileGroupRuleClass().New()
}


// Initializes a new tile group rule with adjacency rules and tile definitions.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule/init(adjacency:tileDefinitions:)
func NewSKTileGroupRuleWithAdjacencyTileDefinitions(adjacency unsafe.Pointer, tileDefinitions unsafe.Pointer) SKTileGroupRule {
	instance := getSKTileGroupRuleClass().Alloc()
	rv := objc.Send[SKTileGroupRule](instance.ID, objc.Sel("initWithAdjacency:tileDefinitions:"), adjacency, tileDefinitions)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKTileGroupRule/tileGroupRuleWithAdjacency:tileDefinitions:
func (sc _SKTileGroupRuleClass) TileGroupRuleWithAdjacencyTileDefinitions(adjacency unsafe.Pointer, tileDefinitions unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("tileGroupRuleWithAdjacency:tileDefinitions:"), adjacency, tileDefinitions)
	return rv
}

