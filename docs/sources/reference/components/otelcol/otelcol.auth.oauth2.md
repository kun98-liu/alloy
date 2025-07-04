---
canonical: https://grafana.com/docs/alloy/latest/reference/components/otelcol/otelcol.auth.oauth2/
aliases:
  - ../otelcol.auth.oauth2/ # /docs/alloy/latest/reference/components/otelcol.auth.oauth2/
description: Learn about otelcol.auth.oauth2
labels:
  stage: general-availability
  products:
    - oss
title: otelcol.auth.oauth2
---

# `otelcol.auth.oauth2`

`otelcol.auth.oauth2` exposes a `handler` that other `otelcol` components can use to authenticate requests using OAuth 2.0.

This component only supports client authentication.

The authorization tokens can be used by HTTP and gRPC based OpenTelemetry exporters.
This component can fetch and refresh expired tokens automatically.
Refer to the [OAuth 2.0 Authorization Framework](https://datatracker.ietf.org/doc/html/rfc6749#section-4.4) for more information about the Auth 2.0 Client Credentials flow.

{{< admonition type="note" >}}
`otelcol.auth.oauth2` is a wrapper over the upstream OpenTelemetry Collector [`oauth2client`][] extension.
Bug reports or feature requests will be redirected to the upstream repository, if necessary.

[`oauth2client`]: https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/{{< param "OTEL_VERSION" >}}/extension/oauth2clientauthextension
{{< /admonition >}}

You can specify multiple `otelcol.auth.oauth2` components by giving them different labels.

## Usage

```alloy
otelcol.auth.oauth2 "<LABEL>" {
    client_id     = "<CLIENT_ID>"
    client_secret = "<CLIENT_SECRET>"
    token_url     = "<TOKEN_URL>"
}
```

## Arguments

You can use the following arguments with `otelcol.auth.oauth2`:

| Name                 | Type                | Description                                                                        | Default | Required |
| -------------------- | ------------------- | ---------------------------------------------------------------------------------- | ------- | -------- |
| `token_url`          | `string`            | The server endpoint URL from which to get tokens.                                  |         | yes      |
| `client_id_file`     | `string`            | The file path to retrieve the client identifier issued to the client.              |         | no       |
| `client_id`          | `string`            | The client identifier issued to the client.                                        |         | no       |
| `client_secret_file` | `string`            | The file path to retrieve the secret string associated with the client identifier. |         | no       |
| `client_secret`      | `secret`            | The secret string associated with the client identifier.                           |         | no       |
| `endpoint_params`    | `map(list(string))` | Additional parameters that are sent to the token endpoint.                         | `{}`    | no       |
| `scopes`             | `list(string)`      | Requested permissions associated for the client.                                   | `[]`    | no       |
| `timeout`            | `duration`          | The timeout on the client connecting to `token_url`.                               | `"0s"`  | no       |

The `timeout` argument is used both for requesting initial tokens and for refreshing tokens. `"0s"` implies no timeout.

At least one of the `client_id` and `client_id_file` pair of arguments must be set.
If both are set, `client_id_file` takes precedence.

Similarly, at least one of the `client_secret` and `client_secret_file` pair of arguments must be set.
If both are set, `client_secret_file` also takes precedence.

## Blocks

You can use the following blocks with `otelcol.auth.oauth2`:

| Block                            | Description                                                                | Required |
| -------------------------------- | -------------------------------------------------------------------------- | -------- |
| [`debug_metrics`][debug_metrics] | Configures the metrics that this component generates to monitor its state. | no       |
| [`tls`][tls]                     | TLS settings for the token client.                                         | no       |
| `tls` > [`tpm`][tpm]             | TPM settings for the TLS key_file.                                         | no       |

[tls]: #tls
[tpm]: #tpm
[debug_metrics]: #debug_metrics

### `debug_metrics`

{{< docs/shared lookup="reference/components/otelcol-debug-metrics-block.md" source="alloy" version="<ALLOY_VERSION>" >}}

### `tls`

The `tls` block configures TLS settings used for connecting to the token client.
If the `tls` block isn't provided, TLS won't be used for communication.

{{< docs/shared lookup="reference/components/otelcol-tls-client-block.md" source="alloy" version="<ALLOY_VERSION>" >}}

### `tpm`

The `tpm` block configures retrieving the TLS `key_file` from a trusted device.

{{< docs/shared lookup="reference/components/otelcol-tls-tpm-block.md" source="alloy" version="<ALLOY_VERSION>" >}}

## Exported fields

The following fields are exported and can be referenced by other components:

| Name      | Type                       | Description                                                     |
| --------- | -------------------------- | --------------------------------------------------------------- |
| `handler` | `capsule(otelcol.Handler)` | A value that other components can use to authenticate requests. |

## Component health

`otelcol.auth.oauth2` is only reported as unhealthy if given an invalid configuration.

## Debug information

`otelcol.auth.oauth2` doesn't expose any component-specific debug information.

## Example

This example configures [`otelcol.exporter.otlp`][otelcol.exporter.otlp] to use OAuth 2.0 for authentication:

```alloy
otelcol.exporter.otlp "example" {
  client {
    endpoint = "my-otlp-grpc-server:4317"
    auth     = otelcol.auth.oauth2.creds.handler
  }
}

otelcol.auth.oauth2 "creds" {
    client_id     = "someclientid"
    client_secret = "someclientsecret"
    token_url     = "https://example.com/oauth2/default/v1/token"
}
```

Here is another example with some optional attributes specified:

```alloy
otelcol.exporter.otlp "example" {
  client {
    endpoint = "my-otlp-grpc-server:4317"
    auth     = otelcol.auth.oauth2.creds.handler
  }
}

otelcol.auth.oauth2 "creds" {
    client_id       = "someclientid2"
    client_secret   = "someclientsecret2"
    token_url       = "https://example.com/oauth2/default/v1/token"
    endpoint_params = {"audience" = ["someaudience"]}
    scopes          = ["api.metrics"]
    timeout         = "3600s"
}
```

[otelcol.exporter.otlp]: ../otelcol.exporter.otlp/
