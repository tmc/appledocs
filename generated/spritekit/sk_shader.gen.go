// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [SKShader] class.
var (
	sKShaderClass     _SKShaderClass
	sKShaderClassOnce sync.Once
)

func getSKShaderClass() _SKShaderClass {
	sKShaderClassOnce.Do(func() {
		sKShaderClass = _SKShaderClass{objc.GetClass("SKShader")}
	})
	return sKShaderClass
}

type _SKShaderClass struct {
	class objc.Class
}

// An interface definition for the [SKShader] class.
type ISKShader interface {
	objectivec.IObject
	AddUniform(uniform unsafe.Pointer)
	RemoveUniformNamed(name string)
	UniformNamed(name string) unsafe.Pointer
}

// An object that allows you to apply a custom fragment shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader
type SKShader struct {
	objectivec.Object
}

// SKShaderFrom constructs a [SKShader] from an unsafe.Pointer.
//
// An object that allows you to apply a custom fragment shader.
func SKShaderFrom(ptr unsafe.Pointer) SKShader {
	return SKShader{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SKShaderClass) Alloc() SKShader {
	rv := objc.Send[SKShader](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKShaderClass) New() SKShader {
	rv := objc.Send[SKShader](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKShader) Init() SKShader {
	rv := objc.Send[SKShader](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKShader) Autorelease() SKShader {
	rv := objc.Send[SKShader](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKShader creates a new SKShader instance.
func NewSKShader() SKShader {
	return getSKShaderClass().New()
}


// Creates a new shader object by loading the source for a fragment shader from a file stored in the app’s bundle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader/init(fileNamed:)
func NewSKShaderWithFileNamed(name string) SKShader {
	rv := objc.Send[SKShader](objc.ID(getSKShaderClass().class), objc.Sel("shaderWithFileNamed:"), objc.String(name))
	return rv
}
// Initializes a new shader object using the specified source code. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader/init(source:)
func NewSKShaderWithSource(source string) SKShader {
	instance := getSKShaderClass().Alloc()
	rv := objc.Send[SKShader](instance.ID, objc.Sel("initWithSource:"), objc.String(source))
	rv.Autorelease()
	return rv
}
// Initializes a new shader object using the specified source and uniform data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader/init(source:uniforms:)
func NewSKShaderWithSourceUniforms(source string, uniforms unsafe.Pointer) SKShader {
	instance := getSKShaderClass().Alloc()
	rv := objc.Send[SKShader](instance.ID, objc.Sel("initWithSource:uniforms:"), objc.String(source), uniforms)
	rv.Autorelease()
	return rv
}


// Creates a new shader object by loading the source for a fragment shader from a file stored in the app’s bundle. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader/init(fileNamed:)
func (sc _SKShaderClass) ShaderWithFileNamed(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("shaderWithFileNamed:"), objc.String(name))
	return rv
}
// Creates a new empty shader object. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader/shader
func (sc _SKShaderClass) Shader() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("shader"))
	return rv
}
// Creates a new shader object using the specified source code. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader/shaderWithSource:
func (sc _SKShaderClass) ShaderWithSource(source string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("shaderWithSource:"), objc.String(source))
	return rv
}
// Creates a new shader object using the specified source and uniform data. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader/shaderWithSource:uniforms:
func (sc _SKShaderClass) ShaderWithSourceUniforms(source string, uniforms unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(sc.class), objc.Sel("shaderWithSource:uniforms:"), objc.String(source), uniforms)
	return rv
}
// Adds a uniform to the shader. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader/addUniform(_:)
func (s_ SKShader) AddUniform(uniform unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addUniform:"), uniform)
}
// Removes a uniform from the shader. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader/removeUniformNamed(_:)
func (s_ SKShader) RemoveUniformNamed(name string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeUniformNamed:"), objc.String(name))
}
// Returns the uniform object corresponding to a particular uniform variable. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKShader/uniformNamed(_:)
func (s_ SKShader) UniformNamed(name string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("uniformNamed:"), objc.String(name))
	return rv
}

