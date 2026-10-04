{{- define "bandmate7.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "bandmate7.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 55 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 55 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 55 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{- define "bandmate7.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "bandmate7.labels" -}}
helm.sh/chart: {{ include "bandmate7.chart" . }}
{{ include "bandmate7.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{- define "bandmate7.selectorLabels" -}}
app.kubernetes.io/name: {{ include "bandmate7.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{- define "bandmate7.image" -}}
{{- $image := .image -}}
{{- if $image.digest }}
{{- printf "%s@%s" $image.repository $image.digest }}
{{- else }}
{{- $tag := $image.tag | default .ctx.Chart.AppVersion | toString }}
{{- if not $tag }}{{ fail (printf "Set %s.image.tag, %s.image.digest or Chart.appVersion" .component .component) }}{{ end }}
{{- printf "%s:%s" $image.repository $tag }}
{{- end }}
{{- end }}

{{- define "bandmate7.backend.image" -}}
{{- include "bandmate7.image" (dict "image" .Values.backend.image "ctx" . "component" "backend") }}
{{- end }}

{{- define "bandmate7.web.image" -}}
{{- include "bandmate7.image" (dict "image" .Values.web.image "ctx" . "component" "web") }}
{{- end }}

{{- define "bandmate7.scheme" -}}
{{- if or .Values.ingress.tlsEnabled (not (empty .Values.ingress.tls)) -}}
https
{{- else -}}
http
{{- end -}}
{{- end }}

{{- define "bandmate7.publicUrl" -}}
{{- if .Values.backend.publicUrl }}
{{- .Values.backend.publicUrl }}
{{- else if .Values.ingress.host }}
{{- printf "%s://%s" (include "bandmate7.scheme" .) .Values.ingress.host }}
{{- end }}
{{- end }}

{{- define "bandmate7.validate" -}}
{{- if and .Values.ingress.enabled (not .Values.ingress.host) }}
{{- fail "ingress.host is required when ingress.enabled is true" }}
{{- end }}
{{- if and .Values.ingress.enabled (eq .Values.ingress.apiPrefix "/") }}
{{- fail "ingress.apiPrefix must not be \"/\" - it would collide with the frontend route" }}
{{- end }}
{{- if and .Values.ingress.enabled (not (hasPrefix "/" .Values.ingress.apiPrefix)) }}
{{- fail "ingress.apiPrefix must start with /" }}
{{- end }}
{{- end }}
