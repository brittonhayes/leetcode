# {{.Title}} :terminal:

##### {{.Subtitle}}

>>> {{.Description}}

<img src="{{.ImageURL}}" width="{{.ImageSize}}" alt="leetcode logo">

{{ range .Sections -}}
    <h2>
        {{.Title}}
    </h2>{{.Emoji}}
    <pre>{{.Body}}</pre>
{{end}}

## Packages :package:

| Title | Description | Path |
| :--- | :--- | :--- |
{{range .Packages -}}
    | {{.Title}} | {{.Description}} | `{{.Path}}` |
{{end}}
