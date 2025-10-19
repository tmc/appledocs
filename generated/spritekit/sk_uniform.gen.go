// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKUniform] class.
var (
	sKUniformClass     _SKUniformClass
	sKUniformClassOnce sync.Once
)

func getSKUniformClass() _SKUniformClass {
	sKUniformClassOnce.Do(func() {
		sKUniformClass = _SKUniformClass{objc.GetClass("SKUniform")}
	})
	return sKUniformClass
}

type _SKUniformClass struct {
	class objc.Class
}

// An interface definition for the [SKUniform] class.
type ISKUniform interface {
	objectivec.IObject
}

// A container for uniform shader data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform
type SKUniform struct {
	objectivec.Object
}

// SKUniformFrom constructs a [SKUniform] from an unsafe.Pointer.
//
// A container for uniform shader data.
func SKUniformFrom(ptr unsafe.Pointer) SKUniform {
	return SKUniform{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKUniformClass) Alloc() SKUniform {
	rv := objc.Send[SKUniform](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKUniformClass) New() SKUniform {
	rv := objc.Send[SKUniform](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKUniform) Init() SKUniform {
	rv := objc.Send[SKUniform](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKUniform) Autorelease() SKUniform {
	rv := objc.Send[SKUniform](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKUniform creates a new SKUniform instance.
func NewSKUniform() SKUniform {
	return getSKUniformClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:vectorFloat4:)
func NewSKUniformWithNameVectorFloat4(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:vectorFloat4:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:vectorFloat3:)
func NewSKUniformWithNameVectorFloat3(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:vectorFloat3:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
// Initializes a new uniform object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:)
func NewSKUniformWithName(name string) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:"), objc.String(name))
	rv.Autorelease()
	return rv
}
// Initializes a new uniform object that holds a floating-point number. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:float:)-48rln
func NewSKUniformWithNameFloat(name string, value float32) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:float:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
// Initializes a new uniform object that holds a vector of two floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:float:)-9g5vj
func NewSKUniformWithNameFloatVector2(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:floatVector2:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
// Creates and initializes a new uniform object that holds a vector of three floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:float:)-9g6a7
func NewSKUniformWithNameFloatVector3(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:floatVector3:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:matrixFloat2x2:)
func NewSKUniformWithNameMatrixFloat2x2(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:matrixFloat2x2:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:matrixFloat3x3:)
func NewSKUniformWithNameMatrixFloat3x3(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:matrixFloat3x3:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:vectorFloat2:)
func NewSKUniformWithNameVectorFloat2(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:vectorFloat2:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
// Initializes a new uniform object that holds a matrix of floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:float:)-60zbm
func NewSKUniformWithNameFloatMatrix4(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:floatMatrix4:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
// Initializes a new uniform object that holds a matrix of floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:float:)-6110m
func NewSKUniformWithNameFloatMatrix2(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:floatMatrix2:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
// Initializes a new uniform object that holds a matrix of floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:float:)-611hs
func NewSKUniformWithNameFloatMatrix3(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:floatMatrix3:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
// Initializes a new uniform object that holds a vector of four floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:float:)-9g7j7
func NewSKUniformWithNameFloatVector4(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:floatVector4:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:matrixFloat4x4:)
func NewSKUniformWithNameMatrixFloat4x4(name string, value unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:matrixFloat4x4:"), objc.String(name), value)
	rv.Autorelease()
	return rv
}
// Initializes a new uniform object that holds a reference to a texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/init(name:texture:)
func NewSKUniformWithNameTexture(name string, texture unsafe.Pointer) SKUniform {
	instance := getSKUniformClass().Alloc()
	rv := objc.Send[SKUniform](instance.ID, objc.Sel("initWithName:texture:"), objc.String(name), texture)
	rv.Autorelease()
	return rv
}


// Creates and initializes a new uniform object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:
func (sc _SKUniformClass) UniformWithName(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:"), objc.String(name))
	return rv
}
// Creates and initializes a new uniform object that holds a floating-point number. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:float:
func (sc _SKUniformClass) UniformWithNameFloat(name string, value float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:float:"), objc.String(name), value)
	return rv
}
// Creates and initializes a new uniform object that holds a matrix of floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:floatMatrix2:
func (sc _SKUniformClass) UniformWithNameFloatMatrix2(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:floatMatrix2:"), objc.String(name), value)
	return rv
}
// Creates and initializes a new uniform object that holds a matrix of floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:floatMatrix3:
func (sc _SKUniformClass) UniformWithNameFloatMatrix3(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:floatMatrix3:"), objc.String(name), value)
	return rv
}
// Creates and initializes a new uniform object that holds a matrix of floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:floatMatrix4:
func (sc _SKUniformClass) UniformWithNameFloatMatrix4(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:floatMatrix4:"), objc.String(name), value)
	return rv
}
// Creates and initializes a new uniform object that holds a vector of two floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:floatVector2:
func (sc _SKUniformClass) UniformWithNameFloatVector2(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:floatVector2:"), objc.String(name), value)
	return rv
}
// Creates and initializes a new uniform object that holds a vector of three floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:floatVector3:
func (sc _SKUniformClass) UniformWithNameFloatVector3(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:floatVector3:"), objc.String(name), value)
	return rv
}
// Creates and initializes a new uniform object that holds a vector of four floating-point numbers. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:floatVector4:
func (sc _SKUniformClass) UniformWithNameFloatVector4(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:floatVector4:"), objc.String(name), value)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:matrixFloat2x2:
func (sc _SKUniformClass) UniformWithNameMatrixFloat2x2(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:matrixFloat2x2:"), objc.String(name), value)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:matrixFloat3x3:
func (sc _SKUniformClass) UniformWithNameMatrixFloat3x3(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:matrixFloat3x3:"), objc.String(name), value)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:matrixFloat4x4:
func (sc _SKUniformClass) UniformWithNameMatrixFloat4x4(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:matrixFloat4x4:"), objc.String(name), value)
	return rv
}
// Creates and initializes a new uniform object that holds a reference to a texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:texture:
func (sc _SKUniformClass) UniformWithNameTexture(name string, texture unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:texture:"), objc.String(name), texture)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:vectorFloat2:
func (sc _SKUniformClass) UniformWithNameVectorFloat2(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:vectorFloat2:"), objc.String(name), value)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:vectorFloat3:
func (sc _SKUniformClass) UniformWithNameVectorFloat3(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:vectorFloat3:"), objc.String(name), value)
	return rv
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKUniform/uniformWithName:vectorFloat4:
func (sc _SKUniformClass) UniformWithNameVectorFloat4(name string, value unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("uniformWithName:vectorFloat4:"), objc.String(name), value)
	return rv
}

