# Mezmo Exporter

This exporter supports sending OpenTelemetry log data to [LogDNA (Mezmo)](https://logdna.com).

# Configuration options:

- `ingest_url` (optional): Specifies the URL to send ingested logs to.  If not specified, will default to `https://logs.logdna.com/log/ingest`.
- `ingest_key` (required): Ingestion key used to send log data to Mezmo.  See [Ingestion Keys](https://docs.logdna.com/docs/ingestion-key) for more details.
- `hostname` (optional): Defines the hostname used to identify the source of logs data in the Mezmo system. If not specified, will default `otel-collector`.

# Example:
## Simple Log Data

```yaml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: ":4317"

exporters:
  mezmo:
    ingest_url: "https://logs.logdna.com/log/ingest"
    ingest_key: "00000000000000000000000000000000"
    hostname: "collector.example.com"

service:
  pipelines:
    logs:
      receivers: [ otlp ]
      exporters: [ mezmo ]
```