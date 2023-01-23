package pod

import "time"

type Pods struct {
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata   struct {
		SelfLink        string `json:"selfLink"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Items []struct {
		Metadata struct {
			Name      string `json:"name"`
			Namespace string `json:"namespace"`
			Labels    struct {
				App string `json:"app"`
			} `json:"labels"`
			Annotations struct {
			} `json:"annotations"`
			CreationTimestamp time.Time `json:"creationTimestamp"`
		} `json:"metadata"`
		Spec struct {
			Containers []struct {
				Name  string `json:"name"`
				Image string `json:"image"`
				Ports []struct {
					Name          string `json:"name"`
					ContainerPort int    `json:"containerPort"`
				} `json:"ports"`
				Resources struct {
				} `json:"resources"`
				VolumeMounts   []interface{} `json:"volumeMounts"`
				LivenessProbe  interface{}   `json:"livenessProbe"`
				ReadinessProbe interface{}   `json:"readinessProbe"`
				Lifecycle      interface{}   `json:"lifecycle"`
			} `json:"containers"`
			Volumes                       []interface{} `json:"volumes"`
			RestartPolicy                 string        `json:"restartPolicy"`
			TerminationGracePeriodSeconds int           `json:"terminationGracePeriodSeconds"`
			DnsPolicy                     string        `json:"dnsPolicy"`
			ServiceAccountName            string        `json:"serviceAccountName"`
			ServiceAccount                string        `json:"serviceAccount"`
			NodeName                      string        `json:"nodeName"`
			SecurityContext               struct {
			} `json:"securityContext"`
			SchedulerName string `json:"schedulerName"`
		} `json:"spec"`
		Status struct {
			Phase      string `json:"phase"`
			Conditions []struct {
				Type               string      `json:"type"`
				Status             string      `json:"status"`
				LastProbeTime      interface{} `json:"lastProbeTime"`
				LastTransitionTime time.Time   `json:"lastTransitionTime"`
			} `json:"conditions"`
			HostIP            string    `json:"hostIP"`
			PodIP             string    `json:"podIP"`
			StartTime         time.Time `json:"startTime"`
			ContainerStatuses []struct {
				Name  string `json:"name"`
				State struct {
					Running struct {
						StartedAt time.Time `json:"startedAt"`
					} `json:"running"`
				} `json:"state"`
				LastState struct {
				} `json:"lastState"`
				Ready        bool   `json:"ready"`
				RestartCount int    `json:"restartCount"`
				Image        string `json:"image"`
				ImageID      string `json:"imageID"`
				ContainerID  string `json:"containerID"`
			} `json:"containerStatuses"`
			QosClass string `json:"qosClass"`
		} `json:"status"`
	} `json:"items"`
}
