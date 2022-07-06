package global

{{- if .HasGlobal }}

import "github.com/yunxiaoyang01/go-mall/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}