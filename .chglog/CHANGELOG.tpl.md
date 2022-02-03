{{ if .Versions -}}
## Unreleased

{{ if .Unreleased.CommitGroups -}}
{{ range .Unreleased.CommitGroups -}}
### {{ .Title }}
{{ range .Commits -}}
- {{ if .Scope }}**{{ .Scope }}:** {{ end }}{{ .Subject }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}

{{ range .Versions }}## {{ if .Tag.Previous }}{{ .Tag.Name }}{{ else }}{{ .Tag.Name }}{{ end }} ({{ datetime "2006-01-02" .Tag.Date }})

{{ if .Tag.Previous }}### Code Diff
[{{ .Tag.Previous.Name }}...{{ .Tag.Name }} Code Diff]({{ $.Info.RepositoryURL }}/compare/{{ .Tag.Previous.Name }}...{{ .Tag.Name }}){{ else }}{{ .Tag.Name }}{{ end }}

### Code Changes
{{ range .CommitGroups -}}

#### {{ .Title }}
{{ range .Commits -}}
* {{ .Subject }}
{{ end }}
{{ end -}}

{{- if .RevertCommits -}}
#### Reverts
{{ range .RevertCommits -}}
* {{ .Revert.Header }}
{{ end }}
{{ end -}}

{{- if .MergeCommits -}}
#### Pull Requests
{{ range .MergeCommits -}}
* {{ .Header }}
{{ end }}
{{ end -}}

{{- if .NoteGroups -}}
{{ range .NoteGroups -}}
#### {{ .Title }}
{{ range .Notes }}
{{ .Body }}
{{ end }}
{{ end -}}
{{ end -}}
{{ end -}}