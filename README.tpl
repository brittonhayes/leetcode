# {{.Title}} :zap:

>>> {{.Description}}

{{if .ReportCard}}
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/leetcode)](https://goreportcard.com/report/github.com/brittonhayes/leetcode)
{{end}}

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
