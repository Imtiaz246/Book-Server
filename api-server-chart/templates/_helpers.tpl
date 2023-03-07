{{- define "setDeploymentName" }}
{{- if .Values.deployments.nameOverWriteWithReleaseName }}
name: {{.Values.deployments.name }}
{{ else }}
name: {{.Release.Name}}
{{- end }}
{{- end }}

{{- define "attachLabels" }}
chartName: {{ .Chart.Name | quote }}
releaseName: {{ .Release.Name | quote }}
version: {{ .Values.deployments.version }}
{{- range $key, $value := .Values.labels }}
{{ $key }}: {{ $value | quote }}
{{- end }}
{{- end }}