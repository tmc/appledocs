// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"unsafe"

	"github.com/ebitengine/purego"
)


// GLKit Functions (48 total)
//
// Type-safe package-level functions with graceful error handling.
// Missing symbols are silently ignored during init; functions will panic when called if unavailable.

var (
	_GLKMathProject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMathUnproject func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrix3Invert func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrix3InvertAndTranspose func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrix4Invert func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrix4InvertAndTranspose func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackCreate func(unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackGetMatrix2 func(unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackGetMatrix3 func(unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackGetMatrix3Inverse func(unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackGetMatrix3InverseTranspose func(unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackGetMatrix4 func(unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackGetMatrix4Inverse func(unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackGetMatrix4InverseTranspose func(unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackGetTypeID func() unsafe.Pointer
	_GLKMatrixStackLoadMatrix4 func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackMultiplyMatrix4 func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackMultiplyMatrixStack func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackPop func(unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackPush func(unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackRotate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackRotateWithVector3 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackRotateWithVector4 func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackRotateX func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackRotateY func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackRotateZ func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackScale func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackScaleWithVector3 func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackScaleWithVector4 func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackSize func(unsafe.Pointer) int
	_GLKMatrixStackTranslate func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackTranslateWithVector3 func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKMatrixStackTranslateWithVector4 func(unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKQuaternionAngle func(unsafe.Pointer) unsafe.Pointer
	_GLKQuaternionAxis func(unsafe.Pointer) unsafe.Pointer
	_GLKQuaternionMakeWithMatrix3 func(unsafe.Pointer) unsafe.Pointer
	_GLKQuaternionMakeWithMatrix4 func(unsafe.Pointer) unsafe.Pointer
	_GLKQuaternionRotateVector3Array func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKQuaternionRotateVector4Array func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKQuaternionSlerp func(unsafe.Pointer, unsafe.Pointer, unsafe.Pointer) unsafe.Pointer
	_GLKVertexAttributeParametersFromModelIO func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromGLKMatrix2 func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromGLKMatrix3 func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromGLKMatrix4 func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromGLKQuaternion func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromGLKVector2 func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromGLKVector3 func(unsafe.Pointer) unsafe.Pointer
	_NSStringFromGLKVector4 func(unsafe.Pointer) unsafe.Pointer
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



// Projects a point in object space into the window coordinate system. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMathProject(_:_:_:_:)
func GLKMathProject(object unsafe.Pointer, model unsafe.Pointer, projection unsafe.Pointer, viewport unsafe.Pointer) unsafe.Pointer {
	return _GLKMathProject(object, model, projection, viewport)
	}


// Projects a point in view space into object space. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMathUnproject(_:_:_:_:_:)
func GLKMathUnproject(window unsafe.Pointer, model unsafe.Pointer, projection unsafe.Pointer, viewport unsafe.Pointer, success unsafe.Pointer) unsafe.Pointer {
	return _GLKMathUnproject(window, model, projection, viewport, success)
	}


// Returns the inverse of a matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrix3Invert(_:_:)
func GLKMatrix3Invert(matrix unsafe.Pointer, isInvertible unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrix3Invert(matrix, isInvertible)
	}


// Returns the inverse transpose of a matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrix3InvertAndTranspose(_:_:)
func GLKMatrix3InvertAndTranspose(matrix unsafe.Pointer, isInvertible unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrix3InvertAndTranspose(matrix, isInvertible)
	}


// Returns the inverse of a matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrix4Invert(_:_:)
func GLKMatrix4Invert(matrix unsafe.Pointer, isInvertible unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrix4Invert(matrix, isInvertible)
	}


// Returns the inverse transpose of a matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrix4InvertAndTranspose(_:_:)
func GLKMatrix4InvertAndTranspose(matrix unsafe.Pointer, isInvertible unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrix4InvertAndTranspose(matrix, isInvertible)
	}


// Allocates and returns a new matrix stack. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackCreate(_:)
func GLKMatrixStackCreate(alloc unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrixStackCreate(alloc)
	}


// Returns the top-left corner of the top matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix2(_:)
func GLKMatrixStackGetMatrix2(stack unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrixStackGetMatrix2(stack)
	}


// Returns the top-left corner of the top matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix3(_:)
func GLKMatrixStackGetMatrix3(stack unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrixStackGetMatrix3(stack)
	}


// Fetches the top-left corner of the top matrix and returns its inverse. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix3Inverse(_:)
func GLKMatrixStackGetMatrix3Inverse(stack unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrixStackGetMatrix3Inverse(stack)
	}


// Fetches the top-left corner of the top matrix and returns its inverse transpose. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix3InverseTranspose(_:)
func GLKMatrixStackGetMatrix3InverseTranspose(stack unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrixStackGetMatrix3InverseTranspose(stack)
	}


// Returns a copy of the top matrix on the stack. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix4(_:)
func GLKMatrixStackGetMatrix4(stack unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrixStackGetMatrix4(stack)
	}


// Returns the inverse of the top matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix4Inverse(_:)
func GLKMatrixStackGetMatrix4Inverse(stack unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrixStackGetMatrix4Inverse(stack)
	}


// Returns the inverse transpose of the top matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetMatrix4InverseTranspose(_:)
func GLKMatrixStackGetMatrix4InverseTranspose(stack unsafe.Pointer) unsafe.Pointer {
	return _GLKMatrixStackGetMatrix4InverseTranspose(stack)
	}


// Returns the Core Foundation type for a matrix stack. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackGetTypeID()
func GLKMatrixStackGetTypeID() unsafe.Pointer {
	return _GLKMatrixStackGetTypeID()
	}


// Replaces the contents of the top matrix with a new matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackLoadMatrix4(_:_:)
func GLKMatrixStackLoadMatrix4(stack unsafe.Pointer, matrix unsafe.Pointer) {
	_GLKMatrixStackLoadMatrix4(stack, matrix)
	}


// Replaces the contents of the top matrix with a matrix calculated by multiplying the contents of the top matrix by another matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackMultiplyMatrix4(_:_:)
func GLKMatrixStackMultiplyMatrix4(stack unsafe.Pointer, matrix unsafe.Pointer) {
	_GLKMatrixStackMultiplyMatrix4(stack, matrix)
	}


// Replaces the contents of the top matrix with a matrix calculated by multiplying the contents of the top matrix by the top matrix of another matrix stack. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackMultiplyMatrixStack(_:_:)
func GLKMatrixStackMultiplyMatrixStack(stackLeft unsafe.Pointer, stackRight unsafe.Pointer) {
	_GLKMatrixStackMultiplyMatrixStack(stackLeft, stackRight)
	}


// Removes the topmost entry from the stack. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackPop(_:)
func GLKMatrixStackPop(stack unsafe.Pointer) {
	_GLKMatrixStackPop(stack)
	}


// Push a copy of the topmost matrix onto the top of the stack. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackPush(_:)
func GLKMatrixStackPush(stack unsafe.Pointer) {
	_GLKMatrixStackPush(stack)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotate(_:_:_:_:_:)
func GLKMatrixStackRotate(stack unsafe.Pointer, radians unsafe.Pointer, x unsafe.Pointer, y unsafe.Pointer, z unsafe.Pointer) {
	_GLKMatrixStackRotate(stack, radians, x, y, z)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateWithVector3(_:_:_:)
func GLKMatrixStackRotateWithVector3(stack unsafe.Pointer, radians unsafe.Pointer, axisVector unsafe.Pointer) {
	_GLKMatrixStackRotateWithVector3(stack, radians, axisVector)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around an arbitrary axis. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateWithVector4(_:_:_:)
func GLKMatrixStackRotateWithVector4(stack unsafe.Pointer, radians unsafe.Pointer, axisVector unsafe.Pointer) {
	_GLKMatrixStackRotateWithVector4(stack, radians, axisVector)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-x axis. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateX(_:_:)
func GLKMatrixStackRotateX(stack unsafe.Pointer, radians unsafe.Pointer) {
	_GLKMatrixStackRotateX(stack, radians)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-y axis. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateY(_:_:)
func GLKMatrixStackRotateY(stack unsafe.Pointer, radians unsafe.Pointer) {
	_GLKMatrixStackRotateY(stack, radians)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a rotation around the positive-z axis. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackRotateZ(_:_:)
func GLKMatrixStackRotateZ(stack unsafe.Pointer, radians unsafe.Pointer) {
	_GLKMatrixStackRotateZ(stack, radians)
	}


// Replaces the contents of the top matrix with a matrix calculated by scaling the contents of the top matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackScale(_:_:_:_:)
func GLKMatrixStackScale(stack unsafe.Pointer, sx unsafe.Pointer, sy unsafe.Pointer, sz unsafe.Pointer) {
	_GLKMatrixStackScale(stack, sx, sy, sz)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a scaling operation. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackScaleWithVector3(_:_:)
func GLKMatrixStackScaleWithVector3(stack unsafe.Pointer, scaleVector unsafe.Pointer) {
	_GLKMatrixStackScaleWithVector3(stack, scaleVector)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a scaling operation defined by a vector. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackScaleWithVector4(_:_:)
func GLKMatrixStackScaleWithVector4(stack unsafe.Pointer, scaleVector unsafe.Pointer) {
	_GLKMatrixStackScaleWithVector4(stack, scaleVector)
	}


// Returns the number of matrices present on the matrix stack. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackSize(_:)
func GLKMatrixStackSize(stack unsafe.Pointer) int {
	return _GLKMatrixStackSize(stack)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation operation. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackTranslate(_:_:_:_:)
func GLKMatrixStackTranslate(stack unsafe.Pointer, tx unsafe.Pointer, ty unsafe.Pointer, tz unsafe.Pointer) {
	_GLKMatrixStackTranslate(stack, tx, ty, tz)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation defined by a vector. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackTranslateWithVector3(_:_:)
func GLKMatrixStackTranslateWithVector3(stack unsafe.Pointer, translationVector unsafe.Pointer) {
	_GLKMatrixStackTranslateWithVector3(stack, translationVector)
	}


// Replaces the contents of the top matrix with a matrix calculated by composing the top matrix with a translation defined by a vector. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKMatrixStackTranslateWithVector4(_:_:)
func GLKMatrixStackTranslateWithVector4(stack unsafe.Pointer, translationVector unsafe.Pointer) {
	_GLKMatrixStackTranslateWithVector4(stack, translationVector)
	}


// Returns the rotation angle of a quaternion. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionAngle(_:)
func GLKQuaternionAngle(quaternion unsafe.Pointer) unsafe.Pointer {
	return _GLKQuaternionAngle(quaternion)
	}


// Returns the axis of rotation of a quaternion. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionAxis(_:)
func GLKQuaternionAxis(quaternion unsafe.Pointer) unsafe.Pointer {
	return _GLKQuaternionAxis(quaternion)
	}


// Creates a quaternion from a rotation matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionMakeWithMatrix3(_:)
func GLKQuaternionMakeWithMatrix3(matrix unsafe.Pointer) unsafe.Pointer {
	return _GLKQuaternionMakeWithMatrix3(matrix)
	}


// Creates a quaternion from a rotation matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionMakeWithMatrix4(_:)
func GLKQuaternionMakeWithMatrix4(matrix unsafe.Pointer) unsafe.Pointer {
	return _GLKQuaternionMakeWithMatrix4(matrix)
	}


// Applies a quaternion rotation to an array of vectors. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionRotateVector3Array(_:_:_:)
func GLKQuaternionRotateVector3Array(quaternion unsafe.Pointer, vectors unsafe.Pointer, vectorCount unsafe.Pointer) {
	_GLKQuaternionRotateVector3Array(quaternion, vectors, vectorCount)
	}


// Applies a quaternion rotation to an array of vectors. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionRotateVector4Array(_:_:_:)
func GLKQuaternionRotateVector4Array(quaternion unsafe.Pointer, vectors unsafe.Pointer, vectorCount unsafe.Pointer) {
	_GLKQuaternionRotateVector4Array(quaternion, vectors, vectorCount)
	}


// Returns the spherical linear interpolation of two quaternions. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKQuaternionSlerp(_:_:_:)
func GLKQuaternionSlerp(quaternionStart unsafe.Pointer, quaternionEnd unsafe.Pointer, t unsafe.Pointer) unsafe.Pointer {
	return _GLKQuaternionSlerp(quaternionStart, quaternionEnd, t)
	}


// GLKVertexAttributeParametersFromModelIO is a GLKit function. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKVertexAttributeParametersFromModelIO(_:)
func GLKVertexAttributeParametersFromModelIO(vertexFormat unsafe.Pointer) unsafe.Pointer {
	return _GLKVertexAttributeParametersFromModelIO(vertexFormat)
	}


// Returns a string that represents the contents of a matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKMatrix2(_:)
func NSStringFromGLKMatrix2(matrix unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromGLKMatrix2(matrix)
	}


// Returns a string that represents the contents of a matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKMatrix3(_:)
func NSStringFromGLKMatrix3(matrix unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromGLKMatrix3(matrix)
	}


// Returns a string that represents the contents of a matrix. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKMatrix4(_:)
func NSStringFromGLKMatrix4(matrix unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromGLKMatrix4(matrix)
	}


// Returns a string that represents the contents of a quaternion. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKQuaternion(_:)
func NSStringFromGLKQuaternion(quaternion unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromGLKQuaternion(quaternion)
	}


// Returns a string that represents the contents of a vector. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKVector2(_:)
func NSStringFromGLKVector2(vector unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromGLKVector2(vector)
	}


// Returns a string that represents the contents of a vector. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKVector3(_:)
func NSStringFromGLKVector3(vector unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromGLKVector3(vector)
	}


// Returns a string that represents the contents of a vector. [Full Topic]
//
// Added in macOS 10.8.
//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/NSStringFromGLKVector4(_:)
func NSStringFromGLKVector4(vector unsafe.Pointer) unsafe.Pointer {
	return _NSStringFromGLKVector4(vector)
	}




