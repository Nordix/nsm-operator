# NSM Operator configuration reference

This document is a configuration reference for Network Service Mesh Operator. It includes information about configuration parameters of NSM Operator and NSM components

## NSM Compoments
| Component                              | Description                                                                     |
| -------------------------------------- | ------------------------------------------------------------------------------- |
| [NSMGR](#nsmgr)                        | Network Service Manager                                                         |
| [Exclude-prefixes](#exclude-prefixes)  | Exclude-prefixes service                                                        |
| Registry                               | Network Service Registry [k8s](#registry-k8s)/[memory](#registry-memory)        |
| [Forwarders](#forwarders)              | [VPP](#forwarder-vpp)/[OVS](#forwarder-ovs)/[SRIOV](#forwarder-sriov) Forwarder |
| [Webhook](#webhook)                    | Admission Webhook                                                               |

### NSMGR
|KEY|TYPE|DEFAULT|DESCRIPTION|
|-|--|-|-|
|NSM_NAME|                              String|                            nmgr|                                                                                              Name of Network service manager|
|NSM_LISTEN_ON|                         Comma-separated list of URL|       unix:///var/lib/networkservicemesh/nsm.io.sock|                                                    url to listen on. tcp:// one will be used a public to register NSM.|
|NSM_REGISTRY_URL|                      URL|                               tcp://localhost:5001|                                                                              A NSE registry url to use|
|NSM_MAX_TOKEN_LIFETIME|                Duration|                          10m|                                                                                               maximum lifetime of tokens|
|NSM_REGISTRY_SERVER_POLICIES|          Comma-separated list of String|    etc/nsm/opa/common/.*.rego,etc/nsm/opa/registry/.*.rego,etc/nsm/opa/server/.*.rego|                paths to files and directories that contain registry server policies|
|NSM_REGISTRY_CLIENT_POLICIES|          Comma-separated list of String|    etc/nsm/opa/common/.*.rego,etc/nsm/opa/registry/.*.rego,etc/nsm/opa/client/.*.rego|                paths to files and directories that contain registry client policies|
|NSM_LOG_LEVEL|                         String|                            INFO|                                                                                              Log level \<TRACE\|DEBUG\|INFO\|WARN\|ERROR\>|
|NSM_DIAL_TIMEOUT|                      Duration|                          100ms|                                                                                             Timeout for the dial the next endpoint|
|NSM_FORWARDER_NETWORK_SERVICE_NAME|    String|                            forwarder|                                                                                         the default service name for forwarder discovering|
|NSM_OPENTELEMETRYENDPOINT|             String|                            otel-collector.observability.svc.cluster.local:4317|                                               OpenTelemetry Collector Endpoint|


### Exclude-prefixes
|KEY|TYPE|DEFAULT|DESCRIPTION|
|-|--|-|-|
|NSM_EXCLUDED_PREFIXES         |Comma-separated list of String|                                                                            |List of excluded prefixes|
|NSM_CONFIG_MAP_NAMESPACE      |String                            |default                                                                  |Namespace of user config map|
|NSM_CONFIG_MAP_NAME           |String                            |excluded-prefixes-config                                                 |Name of user config map|
|NSM_CONFIG_MAP_KEY            |String                            |excluded_prefixes_input.yaml                                             |key in the input configmap by which we retrieve data('filename' in data section in configmap specification yaml file)|
|NSM_OUTPUT_CONFIG_MAP_NAME    |String                            |nsm-config                                                               |Name of nsm config map|
|NSM_OUTPUT_CONFIG_MAP_KEY     |String                            |excluded_prefixes_output.yaml                                            |key in the output configmap by which we retrieve data('filename' in data section in configmap specification yaml file)|
|NSM_OUTPUT_FILE_PATH          |String                            |/var/lib/networkservicemesh/config/excluded_prefixes.yaml                |Path of output prefixes file|
|NSM_PREFIXES_OUTPUT_TYPE      |String                            |file                                                                     |Where to write excluded prefixes|
|NSM_LOG_LEVEL                 |String                            |INFO                                                                     |Log level \<TRACE\|DEBUG\|INFO\|WARN\|ERROR\>
|NSM_OPENTELEMETRYENDPOINT     |String                            |otel-collector.observability.svc.cluster.local:4317                      |OpenTelemetry Collector Endpoint|


### Registry k8s
|KEY|TYPE|DEFAULT|DESCRIPTION|
|-|--|-|-|
|NSM_NAMESPACE                   |String                            |default                                                                                           |namespace where is deployed registry-k8s instance|
|NSM_PROXY_REGISTRY_URL          |URL|                                                                                                                                 |url to the proxy registry that handles this domain|
|NSM_EXPIRE_PERIOD               |Duration                          |1m                                                                                                |period to check expired NSEs|
|NSM_CHAINCTX                    |context.Context                   |||                                                                                                  
|NSM_CLIENTSET                   |versioned.Interface|||                                                                                                                 
|NSM_LISTEN_ON                   |Comma-separated list of URL       |unix:///listen.on.socket                                                                          |url to listen on.|
|NSM_MAX_TOKEN_LIFETIME          |Duration                          |10m                                                                                               |maximum lifetime of tokens|
|NSM_REGISTRY_SERVER_POLICIES    |Comma-separated list of String    |etc/nsm/opa/common/.*.rego,etc/nsm/opa/registry/.*.rego,etc/nsm/opa/server/.*.rego                |paths to files and directories that contain registry server policies|
|NSM_REGISTRY_CLIENT_POLICIES    |Comma-separated list of String    |etc/nsm/opa/common/.*.rego,etc/nsm/opa/registry/.*.rego,etc/nsm/opa/client/.*.rego                |paths to files and directories that contain registry client policies|
|NSM_LOG_LEVEL                   |String                            |INFO                                                                                              |Log level \<TRACE\|DEBUG\|INFO\|WARN\|ERROR\>|
|NSM_OPENTELEMETRYENDPOINT       |String                            |otel-collector.observability.svc.cluster.local:4317                                               |OpenTelemetry Collector Endpoint|


### Registry memory
|KEY|TYPE|DEFAULT|DESCRIPTION|
|-|--|-|-|
|NSM_LISTEN_ON                |Comma-separated list of URL    |unix:///listen.on.socket                                           |url to listen on.|
|NSM_MAX_TOKEN_LIFETIME          |Duration                          |10m                                                                                               |maximum lifetime of tokens|
|NSM_REGISTRY_SERVER_POLICIES    |Comma-separated list of String    |etc/nsm/opa/common/.*.rego,etc/nsm/opa/registry/.*.rego,etc/nsm/opa/server/.*.rego                |paths to files and directories that contain registry server policies|
|NSM_REGISTRY_CLIENT_POLICIES    |Comma-separated list of String    |etc/nsm/opa/common/.*.rego,etc/nsm/opa/registry/.*.rego,etc/nsm/opa/client/.*.rego                |paths to files and directories that contain registry client policies|
|NSM_PROXY_REGISTRY_URL       |URL||                                                                                               url to the proxy registry that handles this domain|
|NSM_EXPIRE_PERIOD            |Duration                       |1s                                                                 |period to check expired NSEs|
|NSM_LOG_LEVEL                |String                         |INFO                                                          |     Log level \<TRACE\|DEBUG\|INFO\|WARN\|ERROR\>|
|NSM_OPENTELEMETRYENDPOINT    |String                         |otel-collector.observability.svc.cluster.local:4317           |     OpenTelemetry Collector Endpoint|


## Forwarders
### forwarder-vpp
| KEY | TYPE | DEFAULT | DESCRIPTION |
| - | - | - | - |
|`NSM_NAME`|String|forwarder||Name of Endpoint|
|`NSM_LABELS`|                   Comma-separated list of String:String pairs    |p2p:true |                                                         Labels related to this forwarder-vpp instance|
|`NSM_NSNAME`                   |String                                         |forwarder|                                                          Name of Network Service to Register with Registry|
|`NSM_CONNECT_TO`               |URL|                                            unix:///connect.to.socket                                          |URL to connect to|
|`NSM_LISTEN_ON`|                URL|                                            unix:///listen.on.socket                                           |URL to listen on|
|`NSM_MAX_TOKEN_LIFETIME`       |Duration                                       |10m|                                                                maximum lifetime of tokens|
|`NSM_REGISTRY_CLIENT_POLICIES`| Comma-separated list of String |         etc/nsm/opa/common/.*.rego,etc/nsm/opa/registry/.*.rego,etc/nsm/opa/client/.*.rego| paths to files and directories that contain registry client policies|
|`NSM_LOG_LEVEL`               |String                                         |INFO|                                                               Log level \<TRACE\|DEBUG\|INFO\|WARN\|ERROR\>|
|`NSM_DIAL_TIMEOUT`|             Duration                                       |100ms|                                                              Timeout for the dial the next endpoint|
|`NSM_OPENTELEMETRYENDPOINT`    |String|                                         otel-collector.observability.svc.cluster.local:4317|                OpenTelemetry Collector Endpoint
|`NSM_TUNNEL_IP`                |String|                                                                    |                                        IP to use for tunnels|
|`NSM_VXLAN_PORT`|               Unsigned Integer                               |0|                                                                  VXLAN port to use|
|`NSM_VPP_API_SOCKET`           |String|                                         /var/run/vpp/external/vpp-api.sock                               |  filename of socket to connect to existing VPP instance.  If empty a VPP instance is run in forwarder|
|`NSM_VPP_INIT`                 |Func                                           |NONE|                                                               type of VPP initialization. Must be NONE or AF_PACKET|
|`NSM_RESOURCE_POLL_TIMEOUT`    |Duration                                       |30s|                                                                device plugin polling timeout|
|`NSM_DEVICE_PLUGIN_PATH`       |String|                                         /var/lib/kubelet/device-plugins/|                                   path to the device plugin directory|
|`NSM_POD_RESOURCES_PATH`       |String|                                         /var/lib/kubelet/pod-resources/|                                    path to the pod resources directory|
|`NSM_DEVICE_SELECTOR_FILE`     |String||                                                                                                            config file for device name to label matching|
|`NSM_SRIOV_CONFIG_FILE`        |String||                                                                                                            PCI resources config path|
|`NSM_PCI_DEVICES_PATH`         |String|                                         /sys/bus/pci/devices|                                               path to the PCI devices directory|
|`NSM_PCI_DRIVERS_PATH`         |String|                                         /sys/bus/pci/drivers|                                               path to the PCI drivers directory|
|`NSM_CGROUP_PATH`              |String|                                         /host/sys/fs/cgroup/devices|                                        path to the host cgroup directory|
|`NSM_VFIO_PATH`                |String|                                         /host/dev/vfio|                                                     path to the host VFIO directory|
|`NSM_MECHANISM_PRIORITY`|   Comma-separated list of String||      sets priorities for mechanisms|


### forwarder-ovs
|KEY|TYPE|DEFAULT|DESCRIPTION|
|-|-|-|-|
|`NSM_NAME`|String|forwarder||Name of Endpoint|
|`NSM_LABELS`|                   Comma-separated list of String:String pairs    |p2p:true |                                                         Labels related to this forwarder-ovs instance|
|`NSM_NSNAME`                   |String                                         |forwarder|                                                          Name of Network Service to Register with Registry|
|`NSM_BRIDGENAME`                   |String|                                         br-nsm|                                                             Name of the OvS bridge|
|`NSM_TUNNEL_IP`                |String|                                                                    |                                        IP to use for tunnels|
|`NSM_CONNECT_TO`               |URL|                                            unix:///connect.to.socket                                          |URL to connect to|
|`NSM_DIAL_TIMEOUT`|             Duration                                       |50ms|                                                              Timeout for the dial the next endpoint|
|`NSM_MAX_TOKEN_LIFETIME`       |Duration                                       |24h|                                                                maximum lifetime of tokens|
|`NSM_REGISTRY_CLIENT_POLICIES`| Comma-separated list of String |         etc/nsm/opa/common/.*.rego,etc/nsm/opa/registry/.*.rego,etc/nsm/opa/client/.*.rego| paths to files and directories that contain registry client policies|
|`NSM_RESOURCE_POLL_TIMEOUT`    |Duration                                       |30s|                                                                device plugin polling timeout|
|`NSM_DEVICE_PLUGIN_PATH`       |String|                                         /var/lib/kubelet/device-plugins/|                                   path to the device plugin directory|
|`NSM_POD_RESOURCES_PATH`       |String|                                         /var/lib/kubelet/pod-resources/|                                    path to the pod resources directory|
|`NSM_SRIOV_CONFIG_FILE`        |String|pci.config|                                                                                                            PCI resources config path|
|`NSM_L2_RESOURCE_SELECTOR_FILE`     |String||                                                                                                            config file for device name to label matching|
|`NSM_PCI_DEVICES_PATH`         |String|                                         /sys/bus/pci/devices|                                               path to the PCI devices directory|
|`NSM_PCI_DRIVERS_PATH`         |String|                                         /sys/bus/pci/drivers|                                               path to the PCI drivers directory|
|`NSM_CGROUP_PATH`              |String|                                         /host/sys/fs/cgroup/devices|                                        path to the host cgroup directory|
|`NSM_VFIO_PATH`                |String|                                         /host/dev/vfio|                                                     path to the host VFIO directory|
|`NSM_LOG_LEVEL`               |String                                         |INFO|                                                               Log level \<TRACE\|DEBUG\|INFO\|WARN\|ERROR\>|
|`NSM_OPENTELEMETRYENDPOINT`    |String|                                         otel-collector.observability.svc.cluster.local:4317|                OpenTelemetry Collector Endpoint|


### forwarder-sriov
|KEY|TYPE|DEFAULT|DESCRIPTION|
|-|--|-|-|
|`NSM_NAME`|String|sriov-forwarder|Name of Endpoint|
|`NSM_LABELS`|                   Comma-separated list of String:String pairs    |p2p:true |                                                         Labels related to this forwarder-sriov instance|
|`NSM_NSNAME`                   |String                                         |forwarder|                                                          Name of Network Service to Register with Registry|
|`NSM_CONNECT_TO`               |URL|                                            unix:///connect.to.socket|                                          URL to connect to|
|`NSM_DIAL_TIMEOUT`|             Duration                                       |50ms|                                                              Timeout for the dial the next endpoint|
|`NSM_MAX_TOKEN_LIFETIME`       |Duration                                       |10m|                                                                maximum lifetime of tokens|
|`NSM_REGISTRY_CLIENT_POLICIES`| Comma-separated list of String |         etc/nsm/opa/common/.*.rego,etc/nsm/opa/registry/.*.rego,etc/nsm/opa/client/.*.rego| paths to files and directories that contain registry client policies|
|`NSM_RESOURCE_POLL_TIMEOUT`    |Duration                                       |30s|                                                                device plugin polling timeout|
|`NSM_DEVICE_PLUGIN_PATH`       |String|                                         /var/lib/kubelet/device-plugins/|                                   path to the device plugin directory|
|`NSM_POD_RESOURCES_PATH`       |String|                                         /var/lib/kubelet/pod-resources/|                                    path to the pod resources directory|
|`NSM_SRIOV_CONFIG_FILE`        |String|pci.config|                                                                                                            PCI resources config path|
|`NSM_PCI_DEVICES_PATH`         |String|                                         /sys/bus/pci/devices|                                               path to the PCI devices directory|
|`NSM_PCI_DRIVERS_PATH`         |String|                                         /sys/bus/pci/drivers|                                               path to the PCI drivers directory|
|`NSM_CGROUP_PATH`              |String|                                         /host/sys/fs/cgroup/devices|                                        path to the host cgroup directory|
|`NSM_VFIO_PATH`                |String|                                         /host/dev/vfio|                                                     path to the host VFIO directory|
|`NSM_LOG_LEVEL`               |String                                         |INFO|                                                               Log level \<TRACE\|DEBUG\|INFO\|WARN\|ERROR\>|
|`NSM_OPENTELEMETRYENDPOINT`    |String|                                         otel-collector.observability.svc.cluster.local:4317|                OpenTelemetry Collector Endpoint|


### Webhook
|KEY|TYPE|DEFAULT|DESCRIPTION|
|-|--|-|-|
|NSM_NAME|                     String|                                         admission-webhook-k8s|                                       Name of current admission webhook instance|
|NSM_SERVICE_NAME|             String|                                         default|                                                     Name of service that related to this admission webhook instance|
|NSM_NAMESPACE|                String|                                         default|                                                     Namespace where admission webhook is deployed|
|NSM_ANNOTATION|               String|                                         networkservicemesh.io|                                       Name of annotation that means that the resource can be handled by admission-webhook|
|NSM_LABELS|                   Comma-separated list of String:String pairs||                                                                Map of labels and their values that should be appended for each deployment that has Config.Annotation|
|NSM_NSURL_ENV_NAME|           String|                                         NSM_NETWORK_SERVICES|                                        Name of env that contains NSURL in initContainers/Containers|
|NSM_INIT_CONTAINER_IMAGES|    Comma-separated list of String||                                                                             List of init containers that should be appended for each deployment that has Config.Annotation|
|NSM_CONTAINER_IMAGES|         Comma-separated list of String||                                                                             List of containers that should be appended for each deployment that has Config.Annotation|
|NSM_ENVS|                     Comma-separated list of String||                                                                             Additional Envs that should be appended for each Config.ContainerImages and Config.InitContainerImages|
|NSM_CERT_FILE_PATH|           String||                                                                                                     Path to certificate|
|NSM_KEY_FILE_PATH|            String||                                                                                                     Path to RSA/Ed25519 related to Config.CertFilePath|
|NSM_CA_BUNDLE_FILE_PATH|      String||                                                                                                     Path to cabundle file related to Config.CertFilePath|
|NSM_OPENTELEMETRYENDPOINT|    String|                                         otel-collector.observability.svc.cluster.local:4317|         OpenTelemetry Collector Endpoint|
