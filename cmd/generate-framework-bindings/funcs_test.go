package main

import (
	"testing"

	"github.com/tmc/appledocs/occ2go"
)

func TestClassToInterfaceName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"NSButton", "IButton"},
		{"NSView", "IView"},
		{"NSWindow", "IWindow"},
		{"CGContext", "IContext"},
		{"CGColor", "IColor"},
		{"CFString", "IString"},
		{"CALayer", "ILayer"},
		{"CIFilter", "IFilter"},
		{"CLLocation", "ILocation"},
		{"CMTime", "ITime"},
		{"CVPixelBuffer", "IPixelBuffer"},
		{"CTFont", "IFont"},
		{"", ""},
		{"NoPrefix", "INoPrefix"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := classToInterfaceName(tt.input)
			if result != tt.expected {
				t.Errorf("classToInterfaceName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestClassToStructName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"NSButton", "Button"},
		{"NSView", "View"},
		{"NSWindow", "Window"},
		{"CGContext", "Context"},
		{"CGColor", "Color"},
		{"CFString", "String"},
		{"CALayer", "Layer"},
		{"CIFilter", "Filter"},
		{"CLLocation", "Location"},
		{"CMTime", "Time"},
		{"CVPixelBuffer", "PixelBuffer"},
		{"CTFont", "Font"},
		{"", ""},
		{"NoPrefix", "NoPrefix"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := classToStructName(tt.input)
			if result != tt.expected {
				t.Errorf("classToStructName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestClassToVarName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"NSButton", "ButtonClass"},
		{"NSView", "ViewClass"},
		{"NSWindow", "WindowClass"},
		{"CGContext", "ContextClass"},
		{"CGColor", "ColorClass"},
		{"CFString", "StringClass"},
		{"", ""},
		{"NoPrefix", "NoPrefixClass"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := classToVarName(tt.input)
			if result != tt.expected {
				t.Errorf("classToVarName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestStripObjCPrefix(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"NSButton", "Button"},
		{"NSView", "View"},
		{"CGContext", "Context"},
		{"CFString", "String"},
		{"CALayer", "Layer"},
		{"CIFilter", "Filter"},
		{"CLLocation", "Location"},
		{"CMTime", "Time"},
		{"CVPixelBuffer", "PixelBuffer"},
		{"CTFont", "Font"},
		{"NoPrefix", "NoPrefix"},
		{"NSone", "NSone"}, // Should not strip if next char is lowercase
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := stripObjCPrefix(tt.input)
			if result != tt.expected {
				t.Errorf("stripObjCPrefix(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestMethodToGoName(t *testing.T) {
	tests := []struct {
		selector string
		expected string
	}{
		{"setTitle:", "SetTitle"},
		{"title", "Title"},
		{"isEnabled", "IsEnabled"},
		{"setEnabled:", "SetEnabled"},
		{"initWithFrame:", "InitWithFrame"},
		{"buttonWithTitle:image:", "ButtonWithTitleImage"},
		{"buttonWithTitle:image:target:action:", "ButtonWithTitleImageTargetAction"},
		{"init", "Init"},
		{"alloc", "Alloc"},
		{"new", "New"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			method := MethodInfo{Selector: tt.selector}
			result := methodToGoName(method)
			if result != tt.expected {
				t.Errorf("methodToGoName(%q) = %q, want %q", tt.selector, result, tt.expected)
			}
		})
	}
}

func TestObjcSelectorFromMethod(t *testing.T) {
	tests := []struct {
		selector string
		expected string
	}{
		{"setTitle:", "setTitle:"},
		{"title", "title"},
		{"buttonWithTitle:image:", "buttonWithTitle:image:"},
		{"init", "init"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			method := MethodInfo{Selector: tt.selector}
			result := objcSelectorFromMethod(method)
			if result != tt.expected {
				t.Errorf("objcSelectorFromMethod(%q) = %q, want %q", tt.selector, result, tt.expected)
			}
		})
	}
}

func TestParameterToGoType(t *testing.T) {
	tests := []struct {
		paramType string
		framework string
		expected  string
	}{
		{"NSString *", "AppKit", "unsafe.Pointer"},
		{"int", "AppKit", "int"},
		{"BOOL", "AppKit", "bool"},
		{"CGFloat", "CoreGraphics", "CGFloat"},
		{"CGRect", "CoreGraphics", "CGRect"},
		{"CGPoint", "CoreGraphics", "CGPoint"},
		{"double", "AppKit", "float64"},
		{"float", "AppKit", "float32"},
		{"void *", "AppKit", "unsafe.Pointer"},
		{"id", "AppKit", "unsafe.Pointer"},
	}

	for _, tt := range tests {
		t.Run(tt.paramType, func(t *testing.T) {
			param := occ2go.Parameter{Type: tt.paramType}
			result := parameterToGoType(param, tt.framework)
			if result != tt.expected {
				t.Errorf("parameterToGoType(%q, %q) = %q, want %q", tt.paramType, tt.framework, result, tt.expected)
			}
		})
	}
}

func TestGenerateConstructorName(t *testing.T) {
	tests := []struct {
		selector string
		expected string
	}{
		{"init", "New"},
		{"initWithFrame:", "NewWithFrame"},
		{"initWithTitle:", "NewWithTitle"},
		{"initWithTitle:image:", "NewWithTitleImage"},
		{"initWithCoder:", "NewWithCoder"},
		{"", "New"},
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			method := MethodInfo{Selector: tt.selector}
			result := generateConstructorName(method)
			if result != tt.expected {
				t.Errorf("generateConstructorName(%q) = %q, want %q", tt.selector, result, tt.expected)
			}
		})
	}
}

func TestIsPropertyGetter(t *testing.T) {
	tests := []struct {
		name       string
		selector   string
		params     []occ2go.Parameter
		returnType string
		expected   bool
	}{
		{
			name:       "simple getter",
			selector:   "title",
			params:     []occ2go.Parameter{},
			returnType: "NSString *",
			expected:   true,
		},
		{
			name:       "boolean getter",
			selector:   "isEnabled",
			params:     []occ2go.Parameter{},
			returnType: "BOOL",
			expected:   true,
		},
		{
			name:       "not a getter - has parameters",
			selector:   "setTitle:",
			params:     []occ2go.Parameter{{Name: "title", Type: "NSString *"}},
			returnType: "void",
			expected:   false,
		},
		{
			name:       "not a getter - returns void",
			selector:   "doSomething",
			params:     []occ2go.Parameter{},
			returnType: "void",
			expected:   false,
		},
		{
			name:       "not a getter - init method",
			selector:   "init",
			params:     []occ2go.Parameter{},
			returnType: "instancetype",
			expected:   false,
		},
		{
			name:       "not a getter - alloc method",
			selector:   "alloc",
			params:     []occ2go.Parameter{},
			returnType: "instancetype",
			expected:   false,
		},
		{
			name:       "not a getter - new method",
			selector:   "new",
			params:     []occ2go.Parameter{},
			returnType: "instancetype",
			expected:   false,
		},
		{
			name:       "not a getter - copy method",
			selector:   "copy",
			params:     []occ2go.Parameter{},
			returnType: "id",
			expected:   false,
		},
		{
			name:       "not a getter - mutableCopy method",
			selector:   "mutableCopy",
			params:     []occ2go.Parameter{},
			returnType: "id",
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method := MethodInfo{
				Selector:   tt.selector,
				Parameters: tt.params,
				ReturnType: tt.returnType,
			}
			result := isPropertyGetter(method)
			if result != tt.expected {
				t.Errorf("isPropertyGetter(%+v) = %v, want %v", method, result, tt.expected)
			}
		})
	}
}

func TestMapObjCTypeToGo_ArraySyntax(t *testing.T) {
	tests := []struct {
		name      string
		objcType  string
		framework string
		expected  string
	}{
		{
			name:      "NSUInteger array with space",
			objcType:  "NSUInteger []",
			framework: "Foundation",
			expected:  "[]uint",
		},
		{
			name:      "NSUInteger array without space",
			objcType:  "NSUInteger[]",
			framework: "Foundation",
			expected:  "[]uint",
		},
		{
			name:      "const NSUInteger array",
			objcType:  "const NSUInteger []",
			framework: "Foundation",
			expected:  "[]uint",
		},
		{
			name:      "NSInteger array",
			objcType:  "NSInteger []",
			framework: "Foundation",
			expected:  "[]int",
		},
		{
			name:      "CGFloat array",
			objcType:  "CGFloat []",
			framework: "CoreGraphics",
			expected:  "[]float64",
		},
		{
			name:      "non-array type unchanged",
			objcType:  "NSString *",
			framework: "Foundation",
			expected:  "string",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapObjCTypeToGo(tt.objcType, tt.framework)
			if result != tt.expected {
				t.Errorf("mapObjCTypeToGo(%q, %q) = %q, want %q", tt.objcType, tt.framework, result, tt.expected)
			}
		})
	}
}

func TestIsPropertySetter(t *testing.T) {
	tests := []struct {
		name       string
		selector   string
		params     []occ2go.Parameter
		returnType string
		expected   bool
	}{
		{
			name:       "simple setter",
			selector:   "setTitle:",
			params:     []occ2go.Parameter{{Name: "title", Type: "NSString *"}},
			returnType: "void",
			expected:   true,
		},
		{
			name:       "boolean setter",
			selector:   "setEnabled:",
			params:     []occ2go.Parameter{{Name: "enabled", Type: "BOOL"}},
			returnType: "void",
			expected:   true,
		},
		{
			name:       "not a setter - no parameters",
			selector:   "title",
			params:     []occ2go.Parameter{},
			returnType: "NSString *",
			expected:   false,
		},
		{
			name:     "not a setter - too many parameters",
			selector: "setTitle:subtitle:",
			params: []occ2go.Parameter{
				{Name: "title", Type: "NSString *"},
				{Name: "subtitle", Type: "NSString *"},
			},
			returnType: "void",
			expected:   false,
		},
		{
			name:       "not a setter - returns value",
			selector:   "setTitle:",
			params:     []occ2go.Parameter{{Name: "title", Type: "NSString *"}},
			returnType: "BOOL",
			expected:   false,
		},
		{
			name:       "not a setter - doesn't start with set",
			selector:   "updateTitle:",
			params:     []occ2go.Parameter{{Name: "title", Type: "NSString *"}},
			returnType: "void",
			expected:   false,
		},
		{
			name:       "not a setter - set with lowercase next char",
			selector:   "setup:",
			params:     []occ2go.Parameter{{Name: "config", Type: "id"}},
			returnType: "void",
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			method := MethodInfo{
				Selector:   tt.selector,
				Parameters: tt.params,
				ReturnType: tt.returnType,
			}
			result := isPropertySetter(method)
			if result != tt.expected {
				t.Errorf("isPropertySetter(%+v) = %v, want %v", method, result, tt.expected)
			}
		})
	}
}

// TestMapObjCTypeToGo_BlockTypes tests mapping of Objective-C block types to Go function types.
// This covers appledocs-506: Map Objective-C blocks to Go function types
func TestMapObjCTypeToGo_BlockTypes(t *testing.T) {
	tests := []struct {
		name      string
		objcType  string
		framework string
		expected  string
	}{
		{
			name:      "simple void block",
			objcType:  "void (^)(void)",
			framework: "Foundation",
			expected:  "func()",
		},
		{
			name:      "array of void blocks (NSBlockOperation.executionBlocks)",
			objcType:  "NSArray<void (^)(void)> *",
			framework: "Foundation",
			expected:  "[]func()",
		},
		{
			name:      "block with error parameter",
			objcType:  "void (^)(NSError *)",
			framework: "Foundation",
			expected:  "func(unsafe.Pointer)", // Note: NSError * maps to unsafe.Pointer, not objc.ID
		},
		{
			name:      "block with bool parameter",
			objcType:  "void (^)(BOOL)",
			framework: "Foundation",
			expected:  "func(bool)", // Note: parameter names are not included in type signatures
		},
		{
			name:      "block with return value",
			objcType:  "BOOL (^)(id, NSError *)",
			framework: "Foundation",
			expected:  "func(unsafe.Pointer, unsafe.Pointer) bool", // Note: parameters map to unsafe.Pointer
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := mapObjCTypeToGo(tt.objcType, tt.framework)
			if result != tt.expected {
				t.Errorf("mapObjCTypeToGo(%q, %q) = %q, want %q", tt.objcType, tt.framework, result, tt.expected)
			}
		})
	}
}
