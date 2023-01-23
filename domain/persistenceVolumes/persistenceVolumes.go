package persistenceVolumes

import "time"

type PersistentVolumes struct {
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata   struct {
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name              string    `json:"name"`
			Uid               string    `json:"uid"`
			ResourceVersion   string    `json:"resourceVersion"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
			Annotations       struct {
				HostPathProvisionerIdentity string `json:"hostPathProvisionerIdentity"`
				PvKubernetesIoProvisionedBy string `json:"pv.kubernetes.io/provisioned-by"`
			} `json:"annotations"`
			Finalizers    []string `json:"finalizers"`
			ManagedFields []struct {
				Manager    string    `json:"manager"`
				Operation  string    `json:"operation"`
				ApiVersion string    `json:"apiVersion"`
				Time       time.Time `json:"time"`
				FieldsType string    `json:"fieldsType"`
				FieldsV1   struct {
					FStatus struct {
						FPhase struct {
						} `json:"f:phase"`
					} `json:"f:status,omitempty"`
					FMetadata struct {
						FAnnotations struct {
							Field1 struct {
							} `json:"."`
							FHostPathProvisionerIdentity struct {
							} `json:"f:hostPathProvisionerIdentity"`
							FPvKubernetesIoProvisionedBy struct {
							} `json:"f:pv.kubernetes.io/provisioned-by"`
						} `json:"f:annotations"`
					} `json:"f:metadata,omitempty"`
					FSpec struct {
						FAccessModes struct {
						} `json:"f:accessModes"`
						FCapacity struct {
							Field1 struct {
							} `json:"."`
							FStorage struct {
							} `json:"f:storage"`
						} `json:"f:capacity"`
						FClaimRef struct {
						} `json:"f:claimRef"`
						FHostPath struct {
							Field1 struct {
							} `json:"."`
							FPath struct {
							} `json:"f:path"`
							FType struct {
							} `json:"f:type"`
						} `json:"f:hostPath"`
						FPersistentVolumeReclaimPolicy struct {
						} `json:"f:persistentVolumeReclaimPolicy"`
						FStorageClassName struct {
						} `json:"f:storageClassName"`
						FVolumeMode struct {
						} `json:"f:volumeMode"`
					} `json:"f:spec,omitempty"`
				} `json:"fieldsV1"`
				Subresource string `json:"subresource,omitempty"`
			} `json:"managedFields"`
		} `json:"metadata"`
		Spec struct {
			Capacity struct {
				Storage string `json:"storage"`
			} `json:"capacity"`
			HostPath struct {
				Path string `json:"path"`
				Type string `json:"type"`
			} `json:"hostPath"`
			AccessModes []string `json:"accessModes"`
			ClaimRef    struct {
				Kind            string `json:"kind"`
				Namespace       string `json:"namespace"`
				Name            string `json:"name"`
				Uid             string `json:"uid"`
				ApiVersion      string `json:"apiVersion"`
				ResourceVersion string `json:"resourceVersion"`
			} `json:"claimRef"`
			PersistentVolumeReclaimPolicy string `json:"persistentVolumeReclaimPolicy"`
			StorageClassName              string `json:"storageClassName"`
			VolumeMode                    string `json:"volumeMode"`
		} `json:"spec"`
		Status struct {
			Phase string `json:"phase"`
		} `json:"status"`
	} `json:"items"`
}
