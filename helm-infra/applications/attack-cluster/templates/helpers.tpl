{{- define "attack-cluster.fullname" -}}
{{- if eq .Release.Name .Chart.Name -}}
{{ .Release.Name }}
{{- else -}}
{{ printf "%s-%s" .Release.Name .Chart.Name }}
{{- end -}}
{{- end }}

{{- define "attack-cluster.labels" -}}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "attack-cluster.selectorLabels" -}}
app.kubernetes.io/name: {{ .Chart.Name }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "attack-cluster.serviceAccountName" -}}
{{ .Release.Name }}-sa
{{- end }}