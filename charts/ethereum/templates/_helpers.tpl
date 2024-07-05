{{/*
Expand the name of the chart.
*/}}
{{- define "ethereum.geth.name" -}}
{{- default .Chart.Name .Values.ethereum.geth.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "ethereum.geth.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "ethereum.geth.labels" -}}
helm.sh/chart: {{ include "ethereum.geth.chart" . }}
{{ include "ethereum.geth.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "ethereum.geth.selectorLabels" -}}
app.kubernetes.io/name: {{ include "ethereum.geth.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the custom exporter service monitor
*/}}
{{- define "ethereum.geth.customServiceMonitorName" -}}
{{ .Release.Name }}-custom-exporter
{{- end }}
