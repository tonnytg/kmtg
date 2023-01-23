package configMaps

import "time"

type ConfigMaps struct {
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
			Annotations       struct {
				KubernetesIoDescription string `json:"kubernetes.io/description"`
			} `json:"annotations,omitempty"`
			ManagedFields []struct {
				Manager    string    `json:"manager"`
				Operation  string    `json:"operation"`
				ApiVersion string    `json:"apiVersion"`
				Time       time.Time `json:"time"`
				FieldsType string    `json:"fieldsType"`
				FieldsV1   struct {
					FData struct {
						Field1 struct {
						} `json:".,omitempty"`
						FCaCrt struct {
						} `json:"f:ca.crt,omitempty"`
						FKubeconfig struct {
						} `json:"f:kubeconfig,omitempty"`
						FCorefile struct {
						} `json:"f:Corefile,omitempty"`
						FClientCaFile struct {
						} `json:"f:client-ca-file,omitempty"`
						FRequestheaderAllowedNames struct {
						} `json:"f:requestheader-allowed-names,omitempty"`
						FRequestheaderClientCaFile struct {
						} `json:"f:requestheader-client-ca-file,omitempty"`
						FRequestheaderExtraHeadersPrefix struct {
						} `json:"f:requestheader-extra-headers-prefix,omitempty"`
						FRequestheaderGroupHeaders struct {
						} `json:"f:requestheader-group-headers,omitempty"`
						FRequestheaderUsernameHeaders struct {
						} `json:"f:requestheader-username-headers,omitempty"`
						FConfigConf struct {
						} `json:"f:config.conf,omitempty"`
						FKubeconfigConf struct {
						} `json:"f:kubeconfig.conf,omitempty"`
						FClusterConfiguration struct {
						} `json:"f:ClusterConfiguration,omitempty"`
						FKubelet struct {
						} `json:"f:kubelet,omitempty"`
					} `json:"f:data"`
					FMetadata struct {
						FAnnotations struct {
							Field1 struct {
							} `json:"."`
							FKubernetesIoDescription struct {
							} `json:"f:kubernetes.io/description"`
						} `json:"f:annotations,omitempty"`
						FLabels struct {
							Field1 struct {
							} `json:"."`
							FApp struct {
							} `json:"f:app"`
						} `json:"f:labels,omitempty"`
					} `json:"f:metadata,omitempty"`
				} `json:"fieldsV1"`
			} `json:"managedFields"`
			Labels struct {
				App string `json:"app"`
			} `json:"labels,omitempty"`
		} `json:"metadata"`
		Data struct {
			CaCrt                           string `json:"ca.crt,omitempty"`
			Kubeconfig                      string `json:"kubeconfig,omitempty"`
			Corefile                        string `json:"Corefile,omitempty"`
			ClientCaFile                    string `json:"client-ca-file,omitempty"`
			RequestheaderAllowedNames       string `json:"requestheader-allowed-names,omitempty"`
			RequestheaderClientCaFile       string `json:"requestheader-client-ca-file,omitempty"`
			RequestheaderExtraHeadersPrefix string `json:"requestheader-extra-headers-prefix,omitempty"`
			RequestheaderGroupHeaders       string `json:"requestheader-group-headers,omitempty"`
			RequestheaderUsernameHeaders    string `json:"requestheader-username-headers,omitempty"`
			ConfigConf                      string `json:"config.conf,omitempty"`
			KubeconfigConf                  string `json:"kubeconfig.conf,omitempty"`
			ClusterConfiguration            string `json:"ClusterConfiguration,omitempty"`
			Kubelet                         string `json:"kubelet,omitempty"`
		} `json:"data"`
	} `json:"items"`
}
