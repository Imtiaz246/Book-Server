{{- define "setDeploymentName" }}
{{- if .Values.deployments.nameOverWriteWithReleaseName }}
name: {{.Values.deployments.name }}
{{ else }}
name: {{.Release.Name}}
{{- end }}
{{- end }}

{{- define "attachLabels" }}
{{- end }}