# {{.Title}} :terminal:

##### {{.Subtitle}}

>>> {{.Description}}

<img src="{{.ImageURL}}" width="{{.ImageSize}}" alt="leetcode logo">

{{ range .Sections -}}
    ## {{.Title}}&nbsp;{{with .Emoji -}}{{.}}{{end}}
    {{.Body}}
{{end}}

## Packages :package:

| Title | Description | Path |
| :--- | :--- | :--- |
{{range .Packages -}}
    | {{.Title}} | {{.Description}} | `{{.Path}}` |
{{end}}
