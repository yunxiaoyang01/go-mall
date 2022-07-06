package global

{{- if .HasGlobal }}

import "github.com/yunxiaoyang01/go-mall/go-mall-admin/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}