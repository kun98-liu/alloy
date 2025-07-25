// Package all imports all known component packages.
package all

import (
	_ "github.com/grafana/alloy/internal/component/beyla/ebpf"                               // Import beyla.ebpf
	_ "github.com/grafana/alloy/internal/component/database_observability/mysql"             // Import database_observability.mysql
	_ "github.com/grafana/alloy/internal/component/discovery/aws"                            // Import discovery.aws.ec2 and discovery.aws.lightsail
	_ "github.com/grafana/alloy/internal/component/discovery/azure"                          // Import discovery.azure
	_ "github.com/grafana/alloy/internal/component/discovery/consul"                         // Import discovery.consul
	_ "github.com/grafana/alloy/internal/component/discovery/consulagent"                    // Import discovery.consulagent
	_ "github.com/grafana/alloy/internal/component/discovery/digitalocean"                   // Import discovery.digitalocean
	_ "github.com/grafana/alloy/internal/component/discovery/dns"                            // Import discovery.dns
	_ "github.com/grafana/alloy/internal/component/discovery/docker"                         // Import discovery.docker
	_ "github.com/grafana/alloy/internal/component/discovery/dockerswarm"                    // Import discovery.dockerswarm
	_ "github.com/grafana/alloy/internal/component/discovery/eureka"                         // Import discovery.eureka
	_ "github.com/grafana/alloy/internal/component/discovery/file"                           // Import discovery.file
	_ "github.com/grafana/alloy/internal/component/discovery/gce"                            // Import discovery.gce
	_ "github.com/grafana/alloy/internal/component/discovery/hetzner"                        // Import discovery.hetzner
	_ "github.com/grafana/alloy/internal/component/discovery/http"                           // Import discovery.http
	_ "github.com/grafana/alloy/internal/component/discovery/ionos"                          // Import discovery.ionos
	_ "github.com/grafana/alloy/internal/component/discovery/kubelet"                        // Import discovery.kubelet
	_ "github.com/grafana/alloy/internal/component/discovery/kubernetes"                     // Import discovery.kubernetes
	_ "github.com/grafana/alloy/internal/component/discovery/kuma"                           // Import discovery.kuma
	_ "github.com/grafana/alloy/internal/component/discovery/linode"                         // Import discovery.linode
	_ "github.com/grafana/alloy/internal/component/discovery/marathon"                       // Import discovery.marathon
	_ "github.com/grafana/alloy/internal/component/discovery/nerve"                          // Import discovery.nerve
	_ "github.com/grafana/alloy/internal/component/discovery/nomad"                          // Import discovery.nomad
	_ "github.com/grafana/alloy/internal/component/discovery/openstack"                      // Import discovery.openstack
	_ "github.com/grafana/alloy/internal/component/discovery/ovhcloud"                       // Import discovery.ovhcloud
	_ "github.com/grafana/alloy/internal/component/discovery/process"                        // Import discovery.process
	_ "github.com/grafana/alloy/internal/component/discovery/puppetdb"                       // Import discovery.puppetdb
	_ "github.com/grafana/alloy/internal/component/discovery/relabel"                        // Import discovery.relabel
	_ "github.com/grafana/alloy/internal/component/discovery/scaleway"                       // Import discovery.scaleway
	_ "github.com/grafana/alloy/internal/component/discovery/serverset"                      // Import discovery.serverset
	_ "github.com/grafana/alloy/internal/component/discovery/triton"                         // Import discovery.triton
	_ "github.com/grafana/alloy/internal/component/discovery/uyuni"                          // Import discovery.uyuni
	_ "github.com/grafana/alloy/internal/component/faro/receiver"                            // Import faro.receiver
	_ "github.com/grafana/alloy/internal/component/local/file"                               // Import local.file
	_ "github.com/grafana/alloy/internal/component/local/file_match"                         // Import local.file_match
	_ "github.com/grafana/alloy/internal/component/loki/echo"                                // Import loki.echo
	_ "github.com/grafana/alloy/internal/component/loki/enrich"                              // Import loki.enrich
	_ "github.com/grafana/alloy/internal/component/loki/process"                             // Import loki.process
	_ "github.com/grafana/alloy/internal/component/loki/relabel"                             // Import loki.relabel
	_ "github.com/grafana/alloy/internal/component/loki/rules/kubernetes"                    // Import loki.rules.kubernetes
	_ "github.com/grafana/alloy/internal/component/loki/secretfilter"                        // Import loki.secretfilter
	_ "github.com/grafana/alloy/internal/component/loki/source/api"                          // Import loki.source.api
	_ "github.com/grafana/alloy/internal/component/loki/source/aws_firehose"                 // Import loki.source.awsfirehose
	_ "github.com/grafana/alloy/internal/component/loki/source/azure_event_hubs"             // Import loki.source.azure_event_hubs
	_ "github.com/grafana/alloy/internal/component/loki/source/cloudflare"                   // Import loki.source.cloudflare
	_ "github.com/grafana/alloy/internal/component/loki/source/docker"                       // Import loki.source.docker
	_ "github.com/grafana/alloy/internal/component/loki/source/file"                         // Import loki.source.file
	_ "github.com/grafana/alloy/internal/component/loki/source/gcplog"                       // Import loki.source.gcplog
	_ "github.com/grafana/alloy/internal/component/loki/source/gelf"                         // Import loki.source.gelf
	_ "github.com/grafana/alloy/internal/component/loki/source/heroku"                       // Import loki.source.heroku
	_ "github.com/grafana/alloy/internal/component/loki/source/journal"                      // Import loki.source.journal
	_ "github.com/grafana/alloy/internal/component/loki/source/kafka"                        // Import loki.source.kafka
	_ "github.com/grafana/alloy/internal/component/loki/source/kubernetes"                   // Import loki.source.kubernetes
	_ "github.com/grafana/alloy/internal/component/loki/source/kubernetes_events"            // Import loki.source.kubernetes_events
	_ "github.com/grafana/alloy/internal/component/loki/source/podlogs"                      // Import loki.source.podlogs
	_ "github.com/grafana/alloy/internal/component/loki/source/syslog"                       // Import loki.source.syslog
	_ "github.com/grafana/alloy/internal/component/loki/source/windowsevent"                 // Import loki.source.windowsevent
	_ "github.com/grafana/alloy/internal/component/loki/write"                               // Import loki.write
	_ "github.com/grafana/alloy/internal/component/mimir/rules/kubernetes"                   // Import mimir.rules.kubernetes
	_ "github.com/grafana/alloy/internal/component/otelcol/auth/basic"                       // Import otelcol.auth.basic
	_ "github.com/grafana/alloy/internal/component/otelcol/auth/bearer"                      // Import otelcol.auth.bearer
	_ "github.com/grafana/alloy/internal/component/otelcol/auth/headers"                     // Import otelcol.auth.headers
	_ "github.com/grafana/alloy/internal/component/otelcol/auth/oauth2"                      // Import otelcol.auth.oauth2
	_ "github.com/grafana/alloy/internal/component/otelcol/auth/sigv4"                       // Import otelcol.auth.sigv4
	_ "github.com/grafana/alloy/internal/component/otelcol/connector/host_info"              // Import otelcol.connector.host_info
	_ "github.com/grafana/alloy/internal/component/otelcol/connector/servicegraph"           // Import otelcol.connector.servicegraph
	_ "github.com/grafana/alloy/internal/component/otelcol/connector/spanlogs"               // Import otelcol.connector.spanlogs
	_ "github.com/grafana/alloy/internal/component/otelcol/connector/spanmetrics"            // Import otelcol.connector.spanmetrics
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/awss3"                   // Import otelcol.exporter.awss3exporter
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/datadog"                 // Import otelcol.exporter.datadog
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/debug"                   // Import otelcol.exporter.debug
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/faro"                    // Import otelcol.exporter.faro
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/googlecloud"             // Import otelcol.exporter.googlecloud
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/kafka"                   // Import otelcol.exporter.kafka
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/loadbalancing"           // Import otelcol.exporter.loadbalancing
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/loki"                    // Import otelcol.exporter.loki
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/otlp"                    // Import otelcol.exporter.otlp
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/otlphttp"                // Import otelcol.exporter.otlphttp
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/prometheus"              // Import otelcol.exporter.prometheus
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/splunkhec"               // Import otelcol.exporter.splunkhec
	_ "github.com/grafana/alloy/internal/component/otelcol/exporter/syslog"                  // Import otelcol.exporter.syslog
	_ "github.com/grafana/alloy/internal/component/otelcol/extension/jaeger_remote_sampling" // Import otelcol.extension.jaeger_remote_sampling
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/attributes"             // Import otelcol.processor.attributes
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/batch"                  // Import otelcol.processor.batch
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/cumulativetodelta"      // Import otelcol.processor.cumulativetodelta
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/deltatocumulative"      // Import otelcol.processor.deltatocumulative
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/discovery"              // Import otelcol.processor.discovery
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/filter"                 // Import otelcol.processor.filter
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/groupbyattrs"           // Import otelcol.processor.groupbyattrs
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/interval"               // Import otelcol.processor.interval
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/k8sattributes"          // Import otelcol.processor.k8sattributes
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/memorylimiter"          // Import otelcol.processor.memory_limiter
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/probabilistic_sampler"  // Import otelcol.processor.probabilistic_sampler
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/resourcedetection"      // Import otelcol.processor.resourcedetection
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/span"                   // Import otelcol.processor.span
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/tail_sampling"          // Import otelcol.processor.tail_sampling
	_ "github.com/grafana/alloy/internal/component/otelcol/processor/transform"              // Import otelcol.processor.transform
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/awscloudwatch"           // Import otelcol.receiver.awscloudwatch
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/datadog"                 // Import otelcol.receiver.datadog
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/file_stats"              // Import otelcol.receiver.file_stats
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/filelog"                 // Import otelcol.receiver.filelog
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/influxdb"                // Import otelcol.receiver.influxdb
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/jaeger"                  // Import otelcol.receiver.jaeger
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/kafka"                   // Import otelcol.receiver.kafka
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/loki"                    // Import otelcol.receiver.loki
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/opencensus"              // Import otelcol.receiver.opencensus
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/otlp"                    // Import otelcol.receiver.otlp
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/prometheus"              // Import otelcol.receiver.prometheus
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/solace"                  // Import otelcol.receiver.solace
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/splunkhec"               // Import otelcol.receiver.splunkhec
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/syslog"                  // Import otelcol.receiver.syslog
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/tcplog"                  // Import otelcol.receiver.tcplog
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/vcenter"                 // Import otelcol.receiver.vcenter
	_ "github.com/grafana/alloy/internal/component/otelcol/receiver/zipkin"                  // Import otelcol.receiver.zipkin
	_ "github.com/grafana/alloy/internal/component/otelcol/storage/file"                     // Import otelcol.storage.file
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/apache"               // Import prometheus.exporter.apache
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/azure"                // Import prometheus.exporter.azure
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/blackbox"             // Import prometheus.exporter.blackbox
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/cadvisor"             // Import prometheus.exporter.cadvisor
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/catchpoint"           // Import prometheus.exporter.catchpoint
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/cloudwatch"           // Import prometheus.exporter.cloudwatch
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/consul"               // Import prometheus.exporter.consul
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/dnsmasq"              // Import prometheus.exporter.dnsmasq
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/elasticsearch"        // Import prometheus.exporter.elasticsearch
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/gcp"                  // Import prometheus.exporter.gcp
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/github"               // Import prometheus.exporter.github
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/kafka"                // Import prometheus.exporter.kafka
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/memcached"            // Import prometheus.exporter.memcached
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/mongodb"              // Import prometheus.exporter.mongodb
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/mssql"                // Import prometheus.exporter.mssql
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/mysql"                // Import prometheus.exporter.mysql
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/oracledb"             // Import prometheus.exporter.oracledb
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/postgres"             // Import prometheus.exporter.postgres
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/process"              // Import prometheus.exporter.process
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/redis"                // Import prometheus.exporter.redis
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/self"                 // Import prometheus.exporter.self
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/snmp"                 // Import prometheus.exporter.snmp
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/snowflake"            // Import prometheus.exporter.snowflake
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/squid"                // Import prometheus.exporter.squid
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/statsd"               // Import prometheus.exporter.statsd
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/unix"                 // Import prometheus.exporter.unix
	_ "github.com/grafana/alloy/internal/component/prometheus/exporter/windows"              // Import prometheus.exporter.windows
	_ "github.com/grafana/alloy/internal/component/prometheus/operator/podmonitors"          // Import prometheus.operator.podmonitors
	_ "github.com/grafana/alloy/internal/component/prometheus/operator/probes"               // Import prometheus.operator.probes
	_ "github.com/grafana/alloy/internal/component/prometheus/operator/scrapeconfigs"        // Import prometheus.operator.scrapeconfigs
	_ "github.com/grafana/alloy/internal/component/prometheus/operator/servicemonitors"      // Import prometheus.operator.servicemonitors
	_ "github.com/grafana/alloy/internal/component/prometheus/receive_http"                  // Import prometheus.receive_http
	_ "github.com/grafana/alloy/internal/component/prometheus/relabel"                       // Import prometheus.relabel
	_ "github.com/grafana/alloy/internal/component/prometheus/remotewrite"                   // Import prometheus.remote_write
	_ "github.com/grafana/alloy/internal/component/prometheus/scrape"                        // Import prometheus.scrape
	_ "github.com/grafana/alloy/internal/component/prometheus/write/queue"                   // Import prometheus.write.queue
	_ "github.com/grafana/alloy/internal/component/pyroscope/ebpf"                           // Import pyroscope.ebpf
	_ "github.com/grafana/alloy/internal/component/pyroscope/java"                           // Import pyroscope.java
	_ "github.com/grafana/alloy/internal/component/pyroscope/receive_http"                   // Import pyroscope.receive_http
	_ "github.com/grafana/alloy/internal/component/pyroscope/relabel"                        // Import pyroscope.relabel
	_ "github.com/grafana/alloy/internal/component/pyroscope/scrape"                         // Import pyroscope.scrape
	_ "github.com/grafana/alloy/internal/component/pyroscope/write"                          // Import pyroscope.write
	_ "github.com/grafana/alloy/internal/component/remote/http"                              // Import remote.http
	_ "github.com/grafana/alloy/internal/component/remote/kubernetes/configmap"              // Import remote.kubernetes.configmap
	_ "github.com/grafana/alloy/internal/component/remote/kubernetes/secret"                 // Import remote.kubernetes.secret
	_ "github.com/grafana/alloy/internal/component/remote/s3"                                // Import remote.s3
	_ "github.com/grafana/alloy/internal/component/remote/vault"                             // Import remote.vault

	_ "github.com/grafana/alloy/internal/util/otelfeaturegatefix" // Gracefully handle duplicate OTEL feature gates
)
