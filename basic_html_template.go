// Package main implements the HTML templates for appledocs
package main

// BasicHTMLTemplate is a simplified HTML template for the index page
const BasicHTMLTemplate = `<\!DOCTYPE html>
<html>
<head>
    <title>Apple Documentation JSON Mirror</title>
    <style>
        body { font-family: system-ui, sans-serif; }
        .tree ul { padding-left: 20px; }
        .folder { color: blue; }
        .file { color: black; }
        .container { display: flex; gap: 20px; }
    </style>
</head>
<body>
    <h1>Apple Documentation JSON Mirror</h1>
    <p>Found {{.FileCount}} JSON files across {{.DirCount}} directories.</p>
    
    <div class="container">
        <div class="tree">{{template "tree" .Root}}</div>
        <div class="content">
            <div id="jsonContent"></div>
        </div>
    </div>
    
    <script>
        // Tree toggle functionality
        document.addEventListener('DOMContentLoaded', function() {
            // Simple JavaScript for expanding tree nodes
            document.querySelectorAll('.folder').forEach(folder => {
                folder.addEventListener('click', function() {
                    const nestedList = this.nextElementSibling;
                    if (nestedList) nestedList.style.display = 
                        nestedList.style.display === 'none' ? 'block' : 'none';
                });
            });
        });
    </script>
</body>
</html>

{{define "tree"}}
<ul>
    {{range .Children}}
        <li>
            {{if .IsDir}}
                <span class="folder">{{.Name}}</span>
                {{template "tree" .}}
            {{else}}
                <span class="file">{{.Name}}</span>
            {{end}}
        </li>
    {{end}}
</ul>
{{end}}`
