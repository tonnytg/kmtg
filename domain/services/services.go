package services

import "time"

type Services struct {
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata   struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Namespace         string    `json:"namespace"`
			Uid               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Labels            struct {
				Component                  string `json:"component,omitempty"`
				Provider                   string `json:"provider,omitempty"`
				App                        string `json:"app,omitempty"`
				K8SApp                     string `json:"k8s-app,omitempty"`
				KubernetesIoClusterService string `json:"kubernetes.io/cluster-service,omitempty"`
				KubernetesIoName           string `json:"kubernetes.io/name,omitempty"`
			} `json:"labels"`
			ManagedFields []struct {
				Manager    string    `json:"manager"`
				Operation  string    `json:"operation"`
				ApiVersion string    `json:"apiVersion"`
				Time       time.Time `json:"time"`
				FieldsType string    `json:"fieldsType"`
				FieldsV1   struct {
					FMetadata struct {
						FLabels struct {
							Field1 struct {
							} `json:"."`
							FComponent struct {
							} `json:"f:component,omitempty"`
							FProvider struct {
							} `json:"f:provider,omitempty"`
							FApp struct {
							} `json:"f:app,omitempty"`
							FK8SApp struct {
							} `json:"f:k8s-app,omitempty"`
							FKubernetesIoClusterService struct {
							} `json:"f:kubernetes.io/cluster-service,omitempty"`
							FKubernetesIoName struct {
							} `json:"f:kubernetes.io/name,omitempty"`
						} `json:"f:labels"`
						FAnnotations struct {
							Field1 struct {
							} `json:"."`
							FKubectlKubernetesIoLastAppliedConfiguration struct {
							} `json:"f:kubectl.kubernetes.io/last-applied-configuration,omitempty"`
							FPrometheusIoPort struct {
							} `json:"f:prometheus.io/port,omitempty"`
							FPrometheusIoScrape struct {
							} `json:"f:prometheus.io/scrape,omitempty"`
						} `json:"f:annotations,omitempty"`
					} `json:"f:metadata"`
					FSpec struct {
						FClusterIP struct {
						} `json:"f:clusterIP"`
						FInternalTrafficPolicy struct {
						} `json:"f:internalTrafficPolicy"`
						FIpFamilyPolicy struct {
						} `json:"f:ipFamilyPolicy,omitempty"`
						FPorts struct {
							Field1 struct {
							} `json:"."`
							KPort443ProtocolTCP struct {
								Field1 struct {
								} `json:"."`
								FName struct {
								} `json:"f:name"`
								FPort struct {
								} `json:"f:port"`
								FProtocol struct {
								} `json:"f:protocol"`
								FTargetPort struct {
								} `json:"f:targetPort"`
							} `json:"k:{"port":443,"protocol":"TCP"},omitempty"`
							KPort80ProtocolTCP struct {
								Field1 struct {
								} `json:"."`
								FName struct {
								} `json:"f:name"`
								FPort struct {
								} `json:"f:port"`
								FProtocol struct {
								} `json:"f:protocol"`
								FTargetPort struct {
								} `json:"f:targetPort"`
							} `json:"k:{"port":80,"protocol":"TCP"},omitempty"`
							KPort53ProtocolTCP struct {
								Field1 struct {
								} `json:"."`
								FName struct {
								} `json:"f:name"`
								FPort struct {
								} `json:"f:port"`
								FProtocol struct {
								} `json:"f:protocol"`
								FTargetPort struct {
								} `json:"f:targetPort"`
							} `json:"k:{"port":53,"protocol":"TCP"},omitempty"`
							KPort53ProtocolUDP struct {
								Field1 struct {
								} `json:"."`
								FName struct {
								} `json:"f:name"`
								FPort struct {
								} `json:"f:port"`
								FProtocol struct {
								} `json:"f:protocol"`
								FTargetPort struct {
								} `json:"f:targetPort"`
							} `json:"k:{"port":53,"protocol":"UDP"},omitempty"`
							KPort9153ProtocolTCP struct {
								Field1 struct {
								} `json:"."`
								FName struct {
								} `json:"f:name"`
								FPort struct {
								} `json:"f:port"`
								FProtocol struct {
								} `json:"f:protocol"`
								FTargetPort struct {
								} `json:"f:targetPort"`
							} `json:"k:{"port":9153,"protocol":"TCP"},omitempty"`
						} `json:"f:ports"`
						FSessionAffinity struct {
						} `json:"f:sessionAffinity"`
						FType struct {
						} `json:"f:type"`
						FSelector struct {
						} `json:"f:selector,omitempty"`
					} `json:"f:spec"`
				} `json:"fieldsV1"`
			} `json:"managedFields"`
			Annotations struct {
				KubectlKubernetesIoLastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration,omitempty"`
				PrometheusIoPort                            string `json:"prometheus.io/port,omitempty"`
				PrometheusIoScrape                          string `json:"prometheus.io/scrape,omitempty"`
			} `json:"annotations,omitempty"`
		} `json:"metadata"`
		Spec struct {
			Ports []struct {
				Name       string `json:"name"`
				Protocol   string `json:"protocol"`
				Port       int    `json:"port"`
				TargetPort int    `json:"targetPort"`
			} `json:"ports"`
			ClusterIP             string   `json:"clusterIP"`
			ClusterIPs            []string `json:"clusterIPs"`
			Type                  string   `json:"type"`
			SessionAffinity       string   `json:"sessionAffinity"`
			IpFamilies            []string `json:"ipFamilies"`
			IpFamilyPolicy        string   `json:"ipFamilyPolicy"`
			InternalTrafficPolicy string   `json:"internalTrafficPolicy"`
			Selector              struct {
				App    string `json:"app,omitempty"`
				K8SApp string `json:"k8s-app,omitempty"`
			} `json:"selector,omitempty"`
		} `json:"spec"`
		Status struct {
			LoadBalancer struct {
			} `json:"loadBalancer"`
		} `json:"status"`
	} `json:"items"`
}
