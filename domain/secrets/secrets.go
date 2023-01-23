package secrets

import "time"

type Secrets struct {
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
				KubectlKubernetesIoLastAppliedConfiguration string `json:"kubectl.kubernetes.io/last-applied-configuration"`
				KubernetesIoServiceAccountName              string `json:"kubernetes.io/service-account.name"`
				KubernetesIoServiceAccountUid               string `json:"kubernetes.io/service-account.uid"`
			} `json:"annotations"`
			ManagedFields []struct {
				Manager    string    `json:"manager"`
				Operation  string    `json:"operation"`
				ApiVersion string    `json:"apiVersion"`
				Time       time.Time `json:"time"`
				FieldsType string    `json:"fieldsType"`
				FieldsV1   struct {
					FData struct {
						Field1 struct {
						} `json:"."`
						FCaCrt struct {
						} `json:"f:ca.crt"`
						FNamespace struct {
						} `json:"f:namespace"`
						FToken struct {
						} `json:"f:token"`
					} `json:"f:data,omitempty"`
					FMetadata struct {
						FAnnotations struct {
							FKubernetesIoServiceAccountUid struct {
							} `json:"f:kubernetes.io/service-account.uid,omitempty"`
							Field2 struct {
							} `json:".,omitempty"`
							FKubectlKubernetesIoLastAppliedConfiguration struct {
							} `json:"f:kubectl.kubernetes.io/last-applied-configuration,omitempty"`
							FKubernetesIoServiceAccountName struct {
							} `json:"f:kubernetes.io/service-account.name,omitempty"`
						} `json:"f:annotations"`
					} `json:"f:metadata"`
					FType struct {
					} `json:"f:type,omitempty"`
				} `json:"fieldsV1"`
			} `json:"managedFields"`
		} `json:"metadata"`
		Data struct {
			CaCrt     string `json:"ca.crt"`
			Namespace string `json:"namespace"`
			Token     string `json:"token"`
		} `json:"data"`
		Type string `json:"type"`
	} `json:"items"`
}
