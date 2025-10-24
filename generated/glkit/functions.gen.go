// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

/* debug [functions.gen.go]: Generating 48 functions for GLKit */
import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// GLKit Functions (48 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_GLKMathProject func(GLKVector3, GLKMatrix4, GLKMatrix4, []int) GLKVector3
	_GLKMathUnproject func(GLKVector3, GLKMatrix4, GLKMatrix4, []int, unsafe.Pointer) GLKVector3
	_GLKMatrix3Invert func(GLKMatrix3, unsafe.Pointer) GLKMatrix3
	_GLKMatrix3InvertAndTranspose func(GLKMatrix3, unsafe.Pointer) GLKMatrix3
	_GLKMatrix4Invert func(GLKMatrix4, unsafe.Pointer) GLKMatrix4
	_GLKMatrix4InvertAndTranspose func(GLKMatrix4, unsafe.Pointer) GLKMatrix4
	_GLKMatrixStackCreate func(AllocatorRef) GLKMatrixStackRef
	_GLKMatrixStackGetMatrix2 func(GLKMatrixStackRef) GLKMatrix2
	_GLKMatrixStackGetMatrix3 func(GLKMatrixStackRef) GLKMatrix3
	_GLKMatrixStackGetMatrix3Inverse func(GLKMatrixStackRef) GLKMatrix3
	_GLKMatrixStackGetMatrix3InverseTranspose func(GLKMatrixStackRef) GLKMatrix3
	_GLKMatrixStackGetMatrix4 func(GLKMatrixStackRef) GLKMatrix4
	_GLKMatrixStackGetMatrix4Inverse func(GLKMatrixStackRef) GLKMatrix4
	_GLKMatrixStackGetMatrix4InverseTranspose func(GLKMatrixStackRef) GLKMatrix4
	_GLKMatrixStackGetTypeID func() TypeID
	_GLKMatrixStackLoadMatrix4 func(GLKMatrixStackRef, GLKMatrix4)
	_GLKMatrixStackMultiplyMatrix4 func(GLKMatrixStackRef, GLKMatrix4)
	_GLKMatrixStackMultiplyMatrixStack func(GLKMatrixStackRef, GLKMatrixStackRef)
	_GLKMatrixStackPop func(GLKMatrixStackRef)
	_GLKMatrixStackPush func(GLKMatrixStackRef)
	_GLKMatrixStackRotate func(GLKMatrixStackRef, float32, float32, float32, float32)
	_GLKMatrixStackRotateWithVector3 func(GLKMatrixStackRef, float32, GLKVector3)
	_GLKMatrixStackRotateWithVector4 func(GLKMatrixStackRef, float32, GLKVector4)
	_GLKMatrixStackRotateX func(GLKMatrixStackRef, float32)
	_GLKMatrixStackRotateY func(GLKMatrixStackRef, float32)
	_GLKMatrixStackRotateZ func(GLKMatrixStackRef, float32)
	_GLKMatrixStackScale func(GLKMatrixStackRef, float32, float32, float32)
	_GLKMatrixStackScaleWithVector3 func(GLKMatrixStackRef, GLKVector3)
	_GLKMatrixStackScaleWithVector4 func(GLKMatrixStackRef, GLKVector4)
	_GLKMatrixStackSize func(GLKMatrixStackRef) int
	_GLKMatrixStackTranslate func(GLKMatrixStackRef, float32, float32, float32)
	_GLKMatrixStackTranslateWithVector3 func(GLKMatrixStackRef, GLKVector3)
	_GLKMatrixStackTranslateWithVector4 func(GLKMatrixStackRef, GLKVector4)
	_GLKQuaternionAngle func(GLKQuaternion) float32
	_GLKQuaternionAxis func(GLKQuaternion) GLKVector3
	_GLKQuaternionMakeWithMatrix3 func(GLKMatrix3) GLKQuaternion
	_GLKQuaternionMakeWithMatrix4 func(GLKMatrix4) GLKQuaternion
	_GLKQuaternionRotateVector3Array func(GLKQuaternion, unsafe.Pointer, uintptr)
	_GLKQuaternionRotateVector4Array func(GLKQuaternion, unsafe.Pointer, uintptr)
	_GLKQuaternionSlerp func(GLKQuaternion, GLKQuaternion, float32) GLKQuaternion
	_GLKVertexAttributeParametersFromModelIO func(unsafe.Pointer) GLKVertexAttributeParameters
	_NSStringFromGLKMatrix2 func(GLKMatrix2) unsafe.Pointer
	_NSStringFromGLKMatrix3 func(GLKMatrix3) unsafe.Pointer
	_NSStringFromGLKMatrix4 func(GLKMatrix4) unsafe.Pointer
	_NSStringFromGLKQuaternion func(GLKQuaternion) unsafe.Pointer
	_NSStringFromGLKVector2 func(GLKVector2) unsafe.Pointer
	_NSStringFromGLKVector3 func(GLKVector3) unsafe.Pointer
	_NSStringFromGLKVector4 func(GLKVector4) unsafe.Pointer
)

func init() {
	lib, err := purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}
	tryRegister(&_GLKMathProject, lib, "GLKMathProject")
	tryRegister(&_GLKMathUnproject, lib, "GLKMathUnproject")
	tryRegister(&_GLKMatrix3Invert, lib, "GLKMatrix3Invert")
	tryRegister(&_GLKMatrix3InvertAndTranspose, lib, "GLKMatrix3InvertAndTranspose")
	tryRegister(&_GLKMatrix4Invert, lib, "GLKMatrix4Invert")
	tryRegister(&_GLKMatrix4InvertAndTranspose, lib, "GLKMatrix4InvertAndTranspose")
	tryRegister(&_GLKMatrixStackCreate, lib, "GLKMatrixStackCreate")
	tryRegister(&_GLKMatrixStackGetMatrix2, lib, "GLKMatrixStackGetMatrix2")
	tryRegister(&_GLKMatrixStackGetMatrix3, lib, "GLKMatrixStackGetMatrix3")
	tryRegister(&_GLKMatrixStackGetMatrix3Inverse, lib, "GLKMatrixStackGetMatrix3Inverse")
	tryRegister(&_GLKMatrixStackGetMatrix3InverseTranspose, lib, "GLKMatrixStackGetMatrix3InverseTranspose")
	tryRegister(&_GLKMatrixStackGetMatrix4, lib, "GLKMatrixStackGetMatrix4")
	tryRegister(&_GLKMatrixStackGetMatrix4Inverse, lib, "GLKMatrixStackGetMatrix4Inverse")
	tryRegister(&_GLKMatrixStackGetMatrix4InverseTranspose, lib, "GLKMatrixStackGetMatrix4InverseTranspose")
	tryRegister(&_GLKMatrixStackGetTypeID, lib, "GLKMatrixStackGetTypeID")
	tryRegister(&_GLKMatrixStackLoadMatrix4, lib, "GLKMatrixStackLoadMatrix4")
	tryRegister(&_GLKMatrixStackMultiplyMatrix4, lib, "GLKMatrixStackMultiplyMatrix4")
	tryRegister(&_GLKMatrixStackMultiplyMatrixStack, lib, "GLKMatrixStackMultiplyMatrixStack")
	tryRegister(&_GLKMatrixStackPop, lib, "GLKMatrixStackPop")
	tryRegister(&_GLKMatrixStackPush, lib, "GLKMatrixStackPush")
	tryRegister(&_GLKMatrixStackRotate, lib, "GLKMatrixStackRotate")
	tryRegister(&_GLKMatrixStackRotateWithVector3, lib, "GLKMatrixStackRotateWithVector3")
	tryRegister(&_GLKMatrixStackRotateWithVector4, lib, "GLKMatrixStackRotateWithVector4")
	tryRegister(&_GLKMatrixStackRotateX, lib, "GLKMatrixStackRotateX")
	tryRegister(&_GLKMatrixStackRotateY, lib, "GLKMatrixStackRotateY")
	tryRegister(&_GLKMatrixStackRotateZ, lib, "GLKMatrixStackRotateZ")
	tryRegister(&_GLKMatrixStackScale, lib, "GLKMatrixStackScale")
	tryRegister(&_GLKMatrixStackScaleWithVector3, lib, "GLKMatrixStackScaleWithVector3")
	tryRegister(&_GLKMatrixStackScaleWithVector4, lib, "GLKMatrixStackScaleWithVector4")
	tryRegister(&_GLKMatrixStackSize, lib, "GLKMatrixStackSize")
	tryRegister(&_GLKMatrixStackTranslate, lib, "GLKMatrixStackTranslate")
	tryRegister(&_GLKMatrixStackTranslateWithVector3, lib, "GLKMatrixStackTranslateWithVector3")
	tryRegister(&_GLKMatrixStackTranslateWithVector4, lib, "GLKMatrixStackTranslateWithVector4")
	tryRegister(&_GLKQuaternionAngle, lib, "GLKQuaternionAngle")
	tryRegister(&_GLKQuaternionAxis, lib, "GLKQuaternionAxis")
	tryRegister(&_GLKQuaternionMakeWithMatrix3, lib, "GLKQuaternionMakeWithMatrix3")
	tryRegister(&_GLKQuaternionMakeWithMatrix4, lib, "GLKQuaternionMakeWithMatrix4")
	tryRegister(&_GLKQuaternionRotateVector3Array, lib, "GLKQuaternionRotateVector3Array")
	tryRegister(&_GLKQuaternionRotateVector4Array, lib, "GLKQuaternionRotateVector4Array")
	tryRegister(&_GLKQuaternionSlerp, lib, "GLKQuaternionSlerp")
	tryRegister(&_GLKVertexAttributeParametersFromModelIO, lib, "GLKVertexAttributeParametersFromModelIO")
	tryRegister(&_NSStringFromGLKMatrix2, lib, "NSStringFromGLKMatrix2")
	tryRegister(&_NSStringFromGLKMatrix3, lib, "NSStringFromGLKMatrix3")
	tryRegister(&_NSStringFromGLKMatrix4, lib, "NSStringFromGLKMatrix4")
	tryRegister(&_NSStringFromGLKQuaternion, lib, "NSStringFromGLKQuaternion")
	tryRegister(&_NSStringFromGLKVector2, lib, "NSStringFromGLKVector2")
	tryRegister(&_NSStringFromGLKVector3, lib, "NSStringFromGLKVector3")
	tryRegister(&_NSStringFromGLKVector4, lib, "NSStringFromGLKVector4")
}

// tryRegister attempts to register a function, silently ignoring failures.
// This allows the library to load even if some symbols are missing.
func tryRegister(fn interface{}, lib uintptr, name string) {
	defer func() {
		if r := recover(); r != nil {
			// Symbol not found - function will remain nil and panic when called
			// This is expected for inline functions, macros, or version-specific APIs
		}
	}()
	purego.RegisterLibFunc(fn, lib, name)
}



// Projects a point in object space into the window coordinate system.
//
// Added in macOS 10.8.
// Projects a point in object space into the window coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMathProject(_:_:_:_:)
func GLKMathProject(object GLKVector3, model GLKMatrix4, projection GLKMatrix4, viewport []int) GLKVector3 {
	return _GLKMathProject(object, model, projection, viewport)
}/* debug [functions.gen.go/function]: GLKMathProject */

// Projects a point in view space into object space.
//
// Added in macOS 10.8.
// Projects a point in view space into object space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMathUnproject(_:_:_:_:_:)
func GLKMathUnproject(window GLKVector3, model GLKMatrix4, projection GLKMatrix4, viewport []int, success unsafe.Pointer) GLKVector3 {
	return _GLKMathUnproject(window, model, projection, viewport, success)
}/* debug [functions.gen.go/function]: GLKMathUnproject */

// Returns the inverse of a matrix.
//
// Added in macOS 10.8.
// Returns the inverse of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrix3Invert(_:_:)
func GLKMatrix3Invert(matrix GLKMatrix3, isInvertible unsafe.Pointer) GLKMatrix3 {
	return _GLKMatrix3Invert(matrix, isInvertible)
}/* debug [functions.gen.go/function]: GLKMatrix3Invert */

// Returns the inverse transpose of a matrix.
//
// Added in macOS 10.8.
// Returns the inverse transpose of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrix3InvertAndTranspose(_:_:)
func GLKMatrix3InvertAndTranspose(matrix GLKMatrix3, isInvertible unsafe.Pointer) GLKMatrix3 {
	return _GLKMatrix3InvertAndTranspose(matrix, isInvertible)
}/* debug [functions.gen.go/function]: GLKMatrix3InvertAndTranspose */

// Returns the inverse of a matrix.
//
// Added in macOS 10.8.
// Returns the inverse of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrix4Invert(_:_:)
func GLKMatrix4Invert(matrix GLKMatrix4, isInvertible unsafe.Pointer) GLKMatrix4 {
	return _GLKMatrix4Invert(matrix, isInvertible)
}/* debug [functions.gen.go/function]: GLKMatrix4Invert */

// Returns the inverse transpose of a matrix.
//
// Added in macOS 10.8.
// Returns the inverse transpose of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrix4InvertAndTranspose(_:_:)
func GLKMatrix4InvertAndTranspose(matrix GLKMatrix4, isInvertible unsafe.Pointer) GLKMatrix4 {
	return _GLKMatrix4InvertAndTranspose(matrix, isInvertible)
}/* debug [functions.gen.go/function]: GLKMatrix4InvertAndTranspose */

// Allocates and returns a new matrix stack.
//
// Added in macOS 10.8.
// Allocates and returns a new matrix stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackCreate(_:)
func GLKMatrixStackCreate(alloc AllocatorRef) GLKMatrixStackRef {
	return _GLKMatrixStackCreate(alloc)
}/* debug [functions.gen.go/function]: GLKMatrixStackCreate */

// Returns the top-left corner of the top matrix.
//
// Added in macOS 10.8.
// Returns the top-left corner of the top matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix2(_:)
func GLKMatrixStackGetMatrix2(stack GLKMatrixStackRef) GLKMatrix2 {
	return _GLKMatrixStackGetMatrix2(stack)
}/* debug [functions.gen.go/function]: GLKMatrixStackGetMatrix2 */

// Returns the top-left corner of the top matrix.
//
// Added in macOS 10.8.
// Returns the top-left corner of the top matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix3(_:)
func GLKMatrixStackGetMatrix3(stack GLKMatrixStackRef) GLKMatrix3 {
	return _GLKMatrixStackGetMatrix3(stack)
}/* debug [functions.gen.go/function]: GLKMatrixStackGetMatrix3 */

// Fetches the top-left corner of the top matrix and returns its inverse.
//
// Added in macOS 10.8.
// Fetches the top-left corner of the top matrix and returns its inverse.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix3Inverse(_:)
func GLKMatrixStackGetMatrix3Inverse(stack GLKMatrixStackRef) GLKMatrix3 {
	return _GLKMatrixStackGetMatrix3Inverse(stack)
}/* debug [functions.gen.go/function]: GLKMatrixStackGetMatrix3Inverse */

// Fetches the top-left corner of the top matrix and returns its inverse transpose.
//
// Added in macOS 10.8.
// Fetches the top-left corner of the top matrix and returns its inverse transpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix3InverseTranspose(_:)
func GLKMatrixStackGetMatrix3InverseTranspose(stack GLKMatrixStackRef) GLKMatrix3 {
	return _GLKMatrixStackGetMatrix3InverseTranspose(stack)
}/* debug [functions.gen.go/function]: GLKMatrixStackGetMatrix3InverseTranspose */

// Returns a copy of the top matrix on the stack.
//
// Added in macOS 10.8.
// Returns a copy of the top matrix on the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix4(_:)
func GLKMatrixStackGetMatrix4(stack GLKMatrixStackRef) GLKMatrix4 {
	return _GLKMatrixStackGetMatrix4(stack)
}/* debug [functions.gen.go/function]: GLKMatrixStackGetMatrix4 */

// Returns the inverse of the top matrix.
//
// Added in macOS 10.8.
// Returns the inverse of the top matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix4Inverse(_:)
func GLKMatrixStackGetMatrix4Inverse(stack GLKMatrixStackRef) GLKMatrix4 {
	return _GLKMatrixStackGetMatrix4Inverse(stack)
}/* debug [functions.gen.go/function]: GLKMatrixStackGetMatrix4Inverse */

// Returns the inverse transpose of the top matrix.
//
// Added in macOS 10.8.
// Returns the inverse transpose of the top matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix4InverseTranspose(_:)
func GLKMatrixStackGetMatrix4InverseTranspose(stack GLKMatrixStackRef) GLKMatrix4 {
	return _GLKMatrixStackGetMatrix4InverseTranspose(stack)
}/* debug [functions.gen.go/function]: GLKMatrixStackGetMatrix4InverseTranspose */

// Returns the Core Foundation type for a matrix stack.
//
// Added in macOS 10.8.
// Returns the Core Foundation type for a matrix stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetTypeID()
func GLKMatrixStackGetTypeID() TypeID {
	return _GLKMatrixStackGetTypeID()
}/* debug [functions.gen.go/function]: GLKMatrixStackGetTypeID */

// Replaces the contents of the top matrix with a new matrix.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a new matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackLoadMatrix4(_:_:)
func GLKMatrixStackLoadMatrix4(stack GLKMatrixStackRef, matrix GLKMatrix4) {
	_GLKMatrixStackLoadMatrix4(stack, matrix)
}/* debug [functions.gen.go/function]: GLKMatrixStackLoadMatrix4 */

// Replaces the contents of the top matrix with a matrix calculated by multiplying the contents of the top matrix by another matrix.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by multiplying the contents of the top matrix by another matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackMultiplyMatrix4(_:_:)
func GLKMatrixStackMultiplyMatrix4(stack GLKMatrixStackRef, matrix GLKMatrix4) {
	_GLKMatrixStackMultiplyMatrix4(stack, matrix)
}/* debug [functions.gen.go/function]: GLKMatrixStackMultiplyMatrix4 */

// Replaces the contents of the top matrix with a matrix calculated by multiplying the contents of the top matrix by the top matrix of another matrix stack.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by multiplying the contents of the top matrix by the top matrix of another matrix stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackMultiplyMatrixStack(_:_:)
func GLKMatrixStackMultiplyMatrixStack(stackLeft GLKMatrixStackRef, stackRight GLKMatrixStackRef) {
	_GLKMatrixStackMultiplyMatrixStack(stackLeft, stackRight)
}/* debug [functions.gen.go/function]: GLKMatrixStackMultiplyMatrixStack */

// Removes the topmost entry from the stack.
//
// Added in macOS 10.8.
// Removes the topmost entry from the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackPop(_:)
func GLKMatrixStackPop(stack GLKMatrixStackRef) {
	_GLKMatrixStackPop(stack)
}/* debug [functions.gen.go/function]: GLKMatrixStackPop */

// Push a copy of the topmost matrix onto the top of the stack.
//
// Added in macOS 10.8.
// Push a copy of the topmost matrix onto the top of the stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackPush(_:)
func GLKMatrixStackPush(stack GLKMatrixStackRef) {
	_GLKMatrixStackPush(stack)
}/* debug [functions.gen.go/function]: GLKMatrixStackPush */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotate(_:_:_:_:_:)
func GLKMatrixStackRotate(stack GLKMatrixStackRef, radians float32, x float32, y float32, z float32) {
	_GLKMatrixStackRotate(stack, radians, x, y, z)
}/* debug [functions.gen.go/function]: GLKMatrixStackRotate */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateWithVector3(_:_:_:)
func GLKMatrixStackRotateWithVector3(stack GLKMatrixStackRef, radians float32, axisVector GLKVector3) {
	_GLKMatrixStackRotateWithVector3(stack, radians, axisVector)
}/* debug [functions.gen.go/function]: GLKMatrixStackRotateWithVector3 */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateWithVector4(_:_:_:)
func GLKMatrixStackRotateWithVector4(stack GLKMatrixStackRef, radians float32, axisVector GLKVector4) {
	_GLKMatrixStackRotateWithVector4(stack, radians, axisVector)
}/* debug [functions.gen.go/function]: GLKMatrixStackRotateWithVector4 */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-x axis.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-x axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateX(_:_:)
func GLKMatrixStackRotateX(stack GLKMatrixStackRef, radians float32) {
	_GLKMatrixStackRotateX(stack, radians)
}/* debug [functions.gen.go/function]: GLKMatrixStackRotateX */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-y axis.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-y axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateY(_:_:)
func GLKMatrixStackRotateY(stack GLKMatrixStackRef, radians float32) {
	_GLKMatrixStackRotateY(stack, radians)
}/* debug [functions.gen.go/function]: GLKMatrixStackRotateY */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-z axis.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-z axis.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateZ(_:_:)
func GLKMatrixStackRotateZ(stack GLKMatrixStackRef, radians float32) {
	_GLKMatrixStackRotateZ(stack, radians)
}/* debug [functions.gen.go/function]: GLKMatrixStackRotateZ */

// Replaces the contents of the top matrix with a matrix calculated by scaling the contents of the top matrix.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by scaling the contents of the top matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackScale(_:_:_:_:)
func GLKMatrixStackScale(stack GLKMatrixStackRef, sx float32, sy float32, sz float32) {
	_GLKMatrixStackScale(stack, sx, sy, sz)
}/* debug [functions.gen.go/function]: GLKMatrixStackScale */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a scaling operation.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a scaling operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackScaleWithVector3(_:_:)
func GLKMatrixStackScaleWithVector3(stack GLKMatrixStackRef, scaleVector GLKVector3) {
	_GLKMatrixStackScaleWithVector3(stack, scaleVector)
}/* debug [functions.gen.go/function]: GLKMatrixStackScaleWithVector3 */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a scaling operation defined by a vector.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a scaling operation defined by a vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackScaleWithVector4(_:_:)
func GLKMatrixStackScaleWithVector4(stack GLKMatrixStackRef, scaleVector GLKVector4) {
	_GLKMatrixStackScaleWithVector4(stack, scaleVector)
}/* debug [functions.gen.go/function]: GLKMatrixStackScaleWithVector4 */

// Returns the number of matrices present on the matrix stack.
//
// Added in macOS 10.8.
// Returns the number of matrices present on the matrix stack.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackSize(_:)
func GLKMatrixStackSize(stack GLKMatrixStackRef) int {
	return _GLKMatrixStackSize(stack)
}/* debug [functions.gen.go/function]: GLKMatrixStackSize */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation operation.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackTranslate(_:_:_:_:)
func GLKMatrixStackTranslate(stack GLKMatrixStackRef, tx float32, ty float32, tz float32) {
	_GLKMatrixStackTranslate(stack, tx, ty, tz)
}/* debug [functions.gen.go/function]: GLKMatrixStackTranslate */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation defined by a vector.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation defined by a vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackTranslateWithVector3(_:_:)
func GLKMatrixStackTranslateWithVector3(stack GLKMatrixStackRef, translationVector GLKVector3) {
	_GLKMatrixStackTranslateWithVector3(stack, translationVector)
}/* debug [functions.gen.go/function]: GLKMatrixStackTranslateWithVector3 */

// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation defined by a vector.
//
// Added in macOS 10.8.
// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation defined by a vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackTranslateWithVector4(_:_:)
func GLKMatrixStackTranslateWithVector4(stack GLKMatrixStackRef, translationVector GLKVector4) {
	_GLKMatrixStackTranslateWithVector4(stack, translationVector)
}/* debug [functions.gen.go/function]: GLKMatrixStackTranslateWithVector4 */

// Returns the rotation angle of a quaternion.
//
// Added in macOS 10.8.
// Returns the rotation angle of a quaternion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionAngle(_:)
func GLKQuaternionAngle(quaternion GLKQuaternion) float32 {
	return _GLKQuaternionAngle(quaternion)
}/* debug [functions.gen.go/function]: GLKQuaternionAngle */

// Returns the axis of rotation of a quaternion.
//
// Added in macOS 10.8.
// Returns the axis of rotation of a quaternion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionAxis(_:)
func GLKQuaternionAxis(quaternion GLKQuaternion) GLKVector3 {
	return _GLKQuaternionAxis(quaternion)
}/* debug [functions.gen.go/function]: GLKQuaternionAxis */

// Creates a quaternion from a rotation matrix.
//
// Added in macOS 10.8.
// Creates a quaternion from a rotation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionMakeWithMatrix3(_:)
func GLKQuaternionMakeWithMatrix3(matrix GLKMatrix3) GLKQuaternion {
	return _GLKQuaternionMakeWithMatrix3(matrix)
}/* debug [functions.gen.go/function]: GLKQuaternionMakeWithMatrix3 */

// Creates a quaternion from a rotation matrix.
//
// Added in macOS 10.8.
// Creates a quaternion from a rotation matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionMakeWithMatrix4(_:)
func GLKQuaternionMakeWithMatrix4(matrix GLKMatrix4) GLKQuaternion {
	return _GLKQuaternionMakeWithMatrix4(matrix)
}/* debug [functions.gen.go/function]: GLKQuaternionMakeWithMatrix4 */

// Applies a quaternion rotation to an array of vectors.
//
// Added in macOS 10.8.
// Applies a quaternion rotation to an array of vectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionRotateVector3Array(_:_:_:)
func GLKQuaternionRotateVector3Array(quaternion GLKQuaternion, vectors unsafe.Pointer, vectorCount uintptr) {
	_GLKQuaternionRotateVector3Array(quaternion, vectors, vectorCount)
}/* debug [functions.gen.go/function]: GLKQuaternionRotateVector3Array */

// Applies a quaternion rotation to an array of vectors.
//
// Added in macOS 10.8.
// Applies a quaternion rotation to an array of vectors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionRotateVector4Array(_:_:_:)
func GLKQuaternionRotateVector4Array(quaternion GLKQuaternion, vectors unsafe.Pointer, vectorCount uintptr) {
	_GLKQuaternionRotateVector4Array(quaternion, vectors, vectorCount)
}/* debug [functions.gen.go/function]: GLKQuaternionRotateVector4Array */

// Returns the spherical linear interpolation of two quaternions.
//
// Added in macOS 10.8.
// Returns the spherical linear interpolation of two quaternions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionSlerp(_:_:_:)
func GLKQuaternionSlerp(quaternionStart GLKQuaternion, quaternionEnd GLKQuaternion, t float32) GLKQuaternion {
	return _GLKQuaternionSlerp(quaternionStart, quaternionEnd, t)
}/* debug [functions.gen.go/function]: GLKQuaternionSlerp */

// GLKVertexAttributeParametersFromModelIO is a GLKit function.
//
// Added in macOS 10.8.
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKVertexAttributeParametersFromModelIO(_:)
func GLKVertexAttributeParametersFromModelIO(vertexFormat unsafe.Pointer) GLKVertexAttributeParameters {
	return _GLKVertexAttributeParametersFromModelIO(vertexFormat)
}/* debug [functions.gen.go/function]: GLKVertexAttributeParametersFromModelIO */

// Returns a string that represents the contents of a matrix.
//
// Added in macOS 10.8.
// Returns a string that represents the contents of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKMatrix2(_:)
func NSStringFromGLKMatrix2(matrix GLKMatrix2) unsafe.Pointer {
	return _NSStringFromGLKMatrix2(matrix)
}/* debug [functions.gen.go/function]: NSStringFromGLKMatrix2 */

// Returns a string that represents the contents of a matrix.
//
// Added in macOS 10.8.
// Returns a string that represents the contents of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKMatrix3(_:)
func NSStringFromGLKMatrix3(matrix GLKMatrix3) unsafe.Pointer {
	return _NSStringFromGLKMatrix3(matrix)
}/* debug [functions.gen.go/function]: NSStringFromGLKMatrix3 */

// Returns a string that represents the contents of a matrix.
//
// Added in macOS 10.8.
// Returns a string that represents the contents of a matrix.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKMatrix4(_:)
func NSStringFromGLKMatrix4(matrix GLKMatrix4) unsafe.Pointer {
	return _NSStringFromGLKMatrix4(matrix)
}/* debug [functions.gen.go/function]: NSStringFromGLKMatrix4 */

// Returns a string that represents the contents of a quaternion.
//
// Added in macOS 10.8.
// Returns a string that represents the contents of a quaternion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKQuaternion(_:)
func NSStringFromGLKQuaternion(quaternion GLKQuaternion) unsafe.Pointer {
	return _NSStringFromGLKQuaternion(quaternion)
}/* debug [functions.gen.go/function]: NSStringFromGLKQuaternion */

// Returns a string that represents the contents of a vector.
//
// Added in macOS 10.8.
// Returns a string that represents the contents of a vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKVector2(_:)
func NSStringFromGLKVector2(vector GLKVector2) unsafe.Pointer {
	return _NSStringFromGLKVector2(vector)
}/* debug [functions.gen.go/function]: NSStringFromGLKVector2 */

// Returns a string that represents the contents of a vector.
//
// Added in macOS 10.8.
// Returns a string that represents the contents of a vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKVector3(_:)
func NSStringFromGLKVector3(vector GLKVector3) unsafe.Pointer {
	return _NSStringFromGLKVector3(vector)
}/* debug [functions.gen.go/function]: NSStringFromGLKVector3 */

// Returns a string that represents the contents of a vector.
//
// Added in macOS 10.8.
// Returns a string that represents the contents of a vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKVector4(_:)
func NSStringFromGLKVector4(vector GLKVector4) unsafe.Pointer {
	return _NSStringFromGLKVector4(vector)
}/* debug [functions.gen.go/function]: NSStringFromGLKVector4 */




