# AlibabaCloud LogService Exporter

This exporter supports sending OpenTelemetry data(Traces, Logs, Metrics) to [LogService](https://www.alibabacloud.com/product/log-service).

# Configuration options:

- `endpoint` (required): LogService's [Endpoint](https://www.alibabacloud.com/help/doc-detail/29008.htm).
- `project` (required): LogService's Project Name.
- `logstore` (required): LogService's store Name. For metrics data, you should use metric store.
- `access_key_id` (optional): AlibabaCloud access key id.
- `access_key_secret` (optional): AlibabaCloud access key secret.
- `ecs_ram_role` (optional): set AlibabaCLoud ECS ram role if you are using ACK.
- `token_file_path` (optional): Set token file path if you are using ACK.

# Example:
## Simple Trace Data

```yaml
receivers:
  examplereceiver:

exporters:
  alibabacloud_logservice:
    endpoint: "cn-hangzhou.log.aliyuncs.com"
    project: "demo-project"
    logstore: "traces-store"
    access_key_id: "access-key-id"
    access_key_secret: "access-key-secret"

service:
  pipelines:
    traces:
      receivers: [examplereceiver]
      exporters: [alibabacloud_logservice]
```


## All Telemetry Data
If you are using OpenTelemetry Collector to collect different types of telemetry data, you should send to different LogService's store.

```yaml
receivers:
  examplereceiver:

exporters:
  alibabacloud_logservice/logs:
    endpoint: "cn-hangzhou.log.aliyuncs.com"
    project: "demo-project"
    logstore: "logs-store"
    access_key_id: "access-key-id"
    access_key_secret: "access-key-secret"
  alibabacloud_logservice/metrics:
    endpoint: "cn-hangzhou.log.aliyuncs.com"
    project: "demo-project"
    logstore: "metrics-store"
    access_key_id: "access-key-id"
    access_key_secret: "access-key-secret"
  alibabacloud_logservice/traces:
    endpoint: "cn-hangzhou.log.aliyuncs.com"
    project: "demo-project"
    logstore: "traces-store"
    access_key_id: "access-key-id"
    access_key_secret: "access-key-secret"

service:
  pipelines:
    traces:
      receivers: [examplereceiver]
      exporters: [alibabacloud_logservice/traces]
    logs:
      receivers: [examplereceiver]
      exporters: [alibabacloud_logservice/logs]
    metrics:
      receivers: [examplereceiver]
      exporters: [alibabacloud_logservice/metrics]
```