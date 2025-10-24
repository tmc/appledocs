package occ2go

import "testing"

func TestClassToInterfaceName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSButton", "IButton"},
		{"NSView", "IView"},
		{"CGContext", "IContext"},
		{"", ""},
		{"CALayer", "ILayer"},
		{"CIImage", "IImage"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ClassToInterfaceName(tt.input)
			if got != tt.want {
				t.Errorf("ClassToInterfaceName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestClassToStructName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSButton", "Button"},
		{"NSView", "View"},
		{"CGContext", "Context"},
		{"CALayer", "Layer"},
		{"CIImage", "Image"},
		{"", ""},
		{"MTLDevice", "Device"},
		{"SCStream", "Stream"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ClassToStructName(tt.input)
			if got != tt.want {
				t.Errorf("ClassToStructName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestClassToVarName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSButton", "ButtonClass"},
		{"NSView", "ViewClass"},
		{"CGContext", "ContextClass"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ClassToVarName(tt.input)
			if got != tt.want {
				t.Errorf("ClassToVarName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestStripObjCPrefix(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSButton", "Button"},
		{"NSView", "View"},
		{"CGContext", "Context"},
		{"CALayer", "Layer"},
		{"CIImage", "Image"},
		{"MTLDevice", "Device"},
		{"SCStream", "Stream"},
		{"UIView", "View"},
		{"WKWebView", "WebView"},
		{"SKScene", "Scene"},
		{"MPSNNGraph", "Graph"},
		{"SMJobBless", "JobBless"},
		{"NoPrefix", "NoPrefix"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := StripObjCPrefix(tt.input)
			if got != tt.want {
				t.Errorf("StripObjCPrefix(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestMethodToGoName(t *testing.T) {
	tests := []struct {
		selector string
		want     string
	}{
		{"setTitle:", "SetTitle"},
		{"buttonWithTitle:image:", "ButtonWithTitleImage"},
		{"initWithFrame:", "InitWithFrame"},
		{"isEnabled", "IsEnabled"},
		{"title", "Title"},
		{"", ""},
		{"addSubview:", "AddSubview"},
		{"setTitle:forState:", "SetTitleForState"},
	}

	for _, tt := range tests {
		t.Run(tt.selector, func(t *testing.T) {
			got := MethodToGoName(tt.selector)
			if got != tt.want {
				t.Errorf("MethodToGoName(%q) = %q, want %q", tt.selector, got, tt.want)
			}
		})
	}
}

func TestInitMethodToConstructorName(t *testing.T) {
	tests := []struct {
		className string
		selector  string
		want      string
	}{
		{"NSButton", "init", "NewButton"},
		{"NSButton", "initWithFrame:", "NewButtonWithFrame"},
		{"NSButton", "buttonWithTitle:target:action:", "NewButtonWithTitleTargetAction"},
		{"CIBlendKernel", "kernelWithString:", "NewBlendKernelWithString"},
		{"NSView", "init", "NewView"},
		{"CGContext", "initWithData:", "NewContextWithData"},
	}

	for _, tt := range tests {
		t.Run(tt.className+":"+tt.selector, func(t *testing.T) {
			got := InitMethodToConstructorName(tt.className, tt.selector)
			if got != tt.want {
				t.Errorf("InitMethodToConstructorName(%q, %q) = %q, want %q", tt.className, tt.selector, got, tt.want)
			}
		})
	}
}

func TestPropertyToGoName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"title", "Title"},
		{"backgroundColor", "BackgroundColor"},
		{"object", "GetObject"}, // Special case
		{"", ""},
		{"isEnabled", "IsEnabled"},
		{"frame", "Frame"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := PropertyToGoName(tt.input)
			if got != tt.want {
				t.Errorf("PropertyToGoName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestCapitalizeFirst(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"title", "Title"},
		{"backgroundColor", "BackgroundColor"},
		{"false", "False"},
		{"type", "Type"},
		{"object", "GetObject"}, // Special case
		{"", ""},
		{"True", "True"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := CapitalizeFirst(tt.input)
			if got != tt.want {
				t.Errorf("CapitalizeFirst(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSButton", "ns_button"},
		{"NSTableView", "ns_table_view"},
		{"ICCameraDevice", "ic_camera_device"},
		{"URLRequest", "url_request"},
		{"SimpleClass", "simple_class"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ToSnakeCase(tt.input)
			if got != tt.want {
				t.Errorf("ToSnakeCase(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestClassFileName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSButton", "ns_button.gen.go"},
		{"NSTableView", "ns_table_view.gen.go"},
		{"ICCameraDevice", "ic_camera_device.gen.go"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ClassFileName(tt.input)
			if got != tt.want {
				t.Errorf("ClassFileName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestProtocolFileName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSCopying", "copying_protocol.gen.go"},
		{"NSTableViewDataSource", "table_view_data_source_protocol.gen.go"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ProtocolFileName(tt.input)
			if got != tt.want {
				t.Errorf("ProtocolFileName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestClassTestFileName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"NSButton", "ns_button.gen_test.go"},
		{"NSTableView", "ns_table_view.gen_test.go"},
		{"NSURLRequest", "nsurl_request.gen_test.go"}, // ToSnakeCase treats NSURL as single word
		{"ICCameraDevice", "ic_camera_device.gen_test.go"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := ClassTestFileName(tt.input)
			if got != tt.want {
				t.Errorf("ClassTestFileName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestReceiverName(t *testing.T) {
	tests := []struct {
		className string
		isClass   bool
		want      string
	}{
		{"Button", false, "b_"},
		{"Button", true, "bc"},
		{"View", false, "v_"},
		{"View", true, "vc"},
		{"", false, "x"},
		{"", true, "x"},
	}

	for _, tt := range tests {
		name := tt.className
		if name == "" {
			name = "empty"
		}
		t.Run(name, func(t *testing.T) {
			got := ReceiverName(tt.className, tt.isClass)
			if got != tt.want {
				t.Errorf("ReceiverName(%q, %v) = %q, want %q", tt.className, tt.isClass, got, tt.want)
			}
		})
	}
}

func TestCleanConstantName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"kCGBlendModeNormal", "BlendModeNormal"},
		{"kCAFillModeForwards", "FillModeForwards"},
		{"kCIAttributeFilterName", "AttributeFilterName"},
		{"kNSLeftTextAlignment", "LeftTextAlignment"},
		{"kMyConstant", "MyConstant"},
		{"MyConstant", "MyConstant"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := CleanConstantName(tt.input)
			if got != tt.want {
				t.Errorf("CleanConstantName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestPropertyConflictsWithParent(t *testing.T) {
	tests := []struct {
		className    string
		superClass   string
		propertyName string
		want         bool
	}{
		{"NSControl", "NSView", "view", true},  // view property would conflict with View field
		{"NSButton", "NSView", "title", false}, // title doesn't conflict
		{"NSView", "NSResponder", "window", true},
		{"NSButton", "NSControl", "button", false},
		{"", "NSView", "view", false},          // empty class name
		{"NSButton", "", "view", false},        // empty superclass
		{"NSButton", "NSView", "", false},      // empty property name
	}

	for _, tt := range tests {
		name := tt.className + ":" + tt.superClass + ":" + tt.propertyName
		t.Run(name, func(t *testing.T) {
			got := PropertyConflictsWithParent(tt.className, tt.superClass, tt.propertyName)
			if got != tt.want {
				t.Errorf("PropertyConflictsWithParent(%q, %q, %q) = %v, want %v",
					tt.className, tt.superClass, tt.propertyName, got, tt.want)
			}
		})
	}
}

func TestCommentLine(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"Simple text", "Simple text"},
		{"Text with\nnewline", "Text with newline"},
		{"Multiple  spaces", "Multiple spaces"},
		{"Tabs\tand\nnewlines\r\n", "Tabs and newlines"},
		{"  Leading and trailing  ", "Leading and trailing"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := CommentLine(tt.input)
			if got != tt.want {
				t.Errorf("CommentLine(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestDisambiguateMethodName(t *testing.T) {
	tests := []struct {
		name   string
		method *ParsedMethod
		want   string
	}{
		{
			name: "no parameters",
			method: &ParsedMethod{
				Selector:   "imageByInsertingIntermediate",
				Parameters: []Parameter{},
			},
			want: "ImageByInsertingIntermediate",
		},
		{
			name: "single parameter",
			method: &ParsedMethod{
				Selector: "imageByInsertingIntermediate:",
				Parameters: []Parameter{
					{Name: "cache", Type: "BOOL"},
				},
			},
			want: "ImageByInsertingIntermediateWithCache",
		},
		{
			name: "multiple parameters",
			method: &ParsedMethod{
				Selector: "setTitle:forState:",
				Parameters: []Parameter{
					{Name: "title", Type: "NSString"},
					{Name: "state", Type: "UIControlState"},
				},
			},
			want: "SetTitleForState",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DisambiguateMethodName(tt.method)
			if got != tt.want {
				t.Errorf("DisambiguateMethodName(%q) = %q, want %q", tt.method.Selector, got, tt.want)
			}
		})
	}
}

func TestMethodGoName(t *testing.T) {
	tests := []struct {
		name       string
		method     *ParsedMethod
		allMethods []*ParsedMethod
		want       string
	}{
		{
			name: "unique method name",
			method: &ParsedMethod{
				Selector:   "setTitle:",
				Parameters: []Parameter{{Name: "title", Type: "NSString"}},
			},
			allMethods: []*ParsedMethod{
				{Selector: "setTitle:", Parameters: []Parameter{{Name: "title", Type: "NSString"}}},
				{Selector: "getTitle", Parameters: []Parameter{}},
			},
			want: "SetTitle",
		},
		{
			name: "conflicting method name",
			method: &ParsedMethod{
				Selector:   "imageByInsertingIntermediate:",
				Parameters: []Parameter{{Name: "cache", Type: "BOOL"}},
			},
			allMethods: []*ParsedMethod{
				{Selector: "imageByInsertingIntermediate", Parameters: []Parameter{}},
				{Selector: "imageByInsertingIntermediate:", Parameters: []Parameter{{Name: "cache", Type: "BOOL"}}},
			},
			want: "ImageByInsertingIntermediateWithCache",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MethodGoName(tt.method, tt.allMethods)
			if got != tt.want {
				t.Errorf("MethodGoName(%q) = %q, want %q", tt.method.Selector, got, tt.want)
			}
		})
	}
}
