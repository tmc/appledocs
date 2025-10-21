package occ2go

import "testing"

func TestIsConstructorSelector(t *testing.T) {
	tests := []struct {
		selector string
		want     bool
	}{
		{"init", true},
		{"initWithFrame:", true},
		{"initWithTitle:image:", true},
		{"alloc", true},
		{"allocWithZone:", true},
		{"new", true},
		{"buttonWithTitle:", false},
		{"description", false},
		{"setTitle:", false},
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			got := IsConstructorSelector(tt.selector)
			if got != tt.want {
				t.Errorf("IsConstructorSelector(%q) = %v, want %v", tt.selector, got, tt.want)
			}
		})
	}
}

func TestIsFactoryMethodForClass(t *testing.T) {
	tests := []struct {
		selector      string
		className     string
		isClassMethod bool
		want          bool
	}{
		{"buttonWithTitle:", "NSButton", true, true},
		{"imageNamed:", "NSImage", true, true},
		{"alloc", "NSButton", true, true},
		{"new", "NSButton", true, true},
		{"description", "NSButton", true, false},
		{"init", "NSButton", false, false}, // Instance method
		{"buttonWithTitle:", "NSButton", false, false}, // Instance method
	}

	for _, tt := range tests {
		name := tt.selector + ":" + tt.className
		t.Run(name, func(t *testing.T) {
			got := IsFactoryMethodForClass(tt.selector, tt.className, tt.isClassMethod)
			if got != tt.want {
				t.Errorf("IsFactoryMethodForClass(%q, %q, %v) = %v, want %v",
					tt.selector, tt.className, tt.isClassMethod, got, tt.want)
			}
		})
	}
}

func TestIsEssentialSelector(t *testing.T) {
	tests := []struct {
		selector string
		want     bool
	}{
		{"alloc", true},
		{"allocWithZone:", true},
		{"new", true},
		{"init", true},
		{"autorelease", true},
		{"copy", true},
		{"copyWithZone:", true},
		{"mutableCopy", true},
		{"mutableCopyWithZone:", true},
		{"retain", false},
		{"release", false},
		{"description", false},
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			got := IsEssentialSelector(tt.selector)
			if got != tt.want {
				t.Errorf("IsEssentialSelector(%q) = %v, want %v", tt.selector, got, tt.want)
			}
		})
	}
}

func TestIsInheritedFromNSObject(t *testing.T) {
	tests := []struct {
		selector string
		want     bool
	}{
		{"retain", true},
		{"release", true},
		{"autorelease", true},
		{"isEqual:", true},
		{"hash", true},
		{"description", true},
		{"valueForKey:", true},
		{"copy", true},
		{"customMethod", false},
		{"setTitle:", false},
		{"buttonWithTitle:", false},
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			got := IsInheritedFromNSObject(tt.selector)
			if got != tt.want {
				t.Errorf("IsInheritedFromNSObject(%q) = %v, want %v", tt.selector, got, tt.want)
			}
		})
	}
}

func TestIsMemoryManagementSelector(t *testing.T) {
	tests := []struct {
		selector string
		want     bool
	}{
		{"retain", true},
		{"release", true},
		{"autorelease", true},
		{"dealloc", true},
		{"retainCount", true},
		{"init", false},
		{"copy", false},
		{"description", false},
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			got := IsMemoryManagementSelector(tt.selector)
			if got != tt.want {
				t.Errorf("IsMemoryManagementSelector(%q) = %v, want %v", tt.selector, got, tt.want)
			}
		})
	}
}

func TestIsPropertyAccessor(t *testing.T) {
	tests := []struct {
		selector string
		want     bool
	}{
		{"title", true},          // Getter
		{"setTitle:", true},      // Setter
		{"isEnabled", true},      // Boolean getter
		{"buttonWithTitle:", false}, // Factory method
		{"initWithFrame:", false},   // Init method
		{"setValue:forKey:", false}, // KVC method (2 params)
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			got := IsPropertyAccessor(tt.selector)
			if got != tt.want {
				t.Errorf("IsPropertyAccessor(%q) = %v, want %v", tt.selector, got, tt.want)
			}
		})
	}
}

func TestIsBooleanGetter(t *testing.T) {
	tests := []struct {
		selector string
		want     bool
	}{
		{"isEnabled", true},
		{"hasChildren", true},
		{"canBecomeKey", true},
		{"shouldClose", true},
		{"willAppear", true},
		{"didLoad", true},
		{"title", false},
		{"island", false}, // Starts with "is" but not boolean getter
		{"setTitle:", false},
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			got := IsBooleanGetter(tt.selector)
			if got != tt.want {
				t.Errorf("IsBooleanGetter(%q) = %v, want %v", tt.selector, got, tt.want)
			}
		})
	}
}

func TestHasInitMethods(t *testing.T) {
	tests := []struct {
		name    string
		methods []*ParsedMethod
		want    bool
	}{
		{
			name: "has instance init",
			methods: []*ParsedMethod{
				{Selector: "init", IsClassMethod: false},
				{Selector: "description", IsClassMethod: false},
			},
			want: true,
		},
		{
			name: "has class factory initializer",
			methods: []*ParsedMethod{
				{Selector: "buttonWithTitle:", IsClassMethod: true, IsInitializer: true},
			},
			want: true,
		},
		{
			name: "no init methods",
			methods: []*ParsedMethod{
				{Selector: "description", IsClassMethod: false},
				{Selector: "setTitle:", IsClassMethod: false},
			},
			want: false,
		},
		{
			name:    "empty methods",
			methods: []*ParsedMethod{},
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := HasInitMethods(tt.methods)
			if got != tt.want {
				t.Errorf("HasInitMethods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilterInitMethods(t *testing.T) {
	tests := []struct {
		name    string
		methods []*ParsedMethod
		want    int
	}{
		{
			name: "mixed methods",
			methods: []*ParsedMethod{
				{Selector: "init", IsClassMethod: false},
				{Selector: "initWithFrame:", IsClassMethod: false},
				{Selector: "alloc", IsClassMethod: true, IsInitializer: false},
				{Selector: "description", IsClassMethod: false},
				{Selector: "buttonWithTitle:", IsClassMethod: true, IsInitializer: true},
			},
			want: 4, // init, initWithFrame, alloc (essential), buttonWithTitle (initializer)
		},
		{
			name: "no init methods",
			methods: []*ParsedMethod{
				{Selector: "description", IsClassMethod: false},
				{Selector: "setTitle:", IsClassMethod: false},
			},
			want: 0,
		},
		{
			name:    "empty methods",
			methods: []*ParsedMethod{},
			want:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FilterInitMethods(tt.methods)
			if len(got) != tt.want {
				t.Errorf("FilterInitMethods() returned %d methods, want %d", len(got), tt.want)
			}
		})
	}
}

func TestShouldSkipMethod(t *testing.T) {
	tests := []struct {
		selector string
		want     bool
	}{
		{"retain", true},          // Memory management
		{"release", true},         // Memory management
		{"description", true},     // Inherited from NSObject
		{"isEqual:", true},        // Inherited from NSObject
		{"init", false},           // Constructor, don't skip
		{"customMethod", false},   // Custom method, don't skip
		{"setTitle:", false},      // Property setter, don't skip
		{"buttonWithTitle:", false}, // Factory method, don't skip
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			got := ShouldSkipMethod(tt.selector)
			if got != tt.want {
				t.Errorf("ShouldSkipMethod(%q) = %v, want %v", tt.selector, got, tt.want)
			}
		})
	}
}
