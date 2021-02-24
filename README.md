# Overview

This is a golang example that uses
[OpenTelemetry](https://opentelemetry.io) to fulfill observability.

# Usage

## Build

```shell
$ make build-demo
```

## Help

```shell
$ ./release/demo -h
```

## Run and exports metrics to stdout

```shell
$ ./release/demo --use-stdout-exporter
```

## Run and exports metrics to [AlibabaCloud SLS](https://www.aliyun.com/product/sls)

1. Download `otel-contrib-collector` from
   [opentelemetry-collector-contrib releases](https://github.com/open-telemetry/opentelemetry-collector-contrib/releases)
   and make it executable and searchable from `$PATH`.

2. Prepare a yaml file of collector configuration:

```yaml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: "0.0.0.0:4317"

exporters:
  alibabacloud_logservice/metrics:
    # LogService's Endpoint, https://www.alibabacloud.com/help/doc-detail/29008.htm
    # for AlibabaCloud Kubernetes(or ECS), set {region-id}-intranet.log.aliyuncs.com, eg cn-hangzhou-intranet.log.aliyuncs.com;
    #  others set {region-id}.log.aliyuncs.com, eg cn-hangzhou.log.aliyuncs.com
    endpoint: "cn-hangzhou.log.aliyuncs.com"
    # LogService's Project Name
    project: ""
    # LogService's Logstore Name
    logstore: ""
    # AlibabaCloud access key id
    access_key_id: ""
    # AlibabaCloud access key secret
    access_key_secret: ""
  file:
    path: ./metrics.json

service:
  pipelines:
    metrics:
      receivers: [otlp]
      exporters: [alibabacloud_logservice/metrics, file]
```

You can find the detailed configuration of AlibabaCloud SLS [here](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/exporter/alibabacloudlogserviceexporter).

3. Run `otelcontribcol`

```shell
$ otelcontribcol --config /PATH/TO/<SLS config yaml>
```

4. Run app

```shell
$ ./release/demo --use-stdout-exporter=false --use-collector --collector-otlp-endpoint 0.0.0.0:4317
```
