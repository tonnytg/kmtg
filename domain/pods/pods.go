package pods

import "time"

type Pods struct {
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
				Run                            string `json:"run,omitempty"`
				App                            string `json:"app,omitempty"`
				ControllerRevisionHash         string `json:"controller-revision-hash,omitempty"`
				StatefulsetKubernetesIoPodName string `json:"statefulset.kubernetes.io/pod-name,omitempty"`
				K8SApp                         string `json:"k8s-app,omitempty"`
				PodTemplateHash                string `json:"pod-template-hash,omitempty"`
				Component                      string `json:"component,omitempty"`
				Tier                           string `json:"tier,omitempty"`
				PodTemplateGeneration          string `json:"pod-template-generation,omitempty"`
				AddonmanagerKubernetesIoMode   string `json:"addonmanager.kubernetes.io/mode,omitempty"`
				IntegrationTest                string `json:"integration-test,omitempty"`
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
							FRun struct {
							} `json:"f:run,omitempty"`
							FApp struct {
							} `json:"f:app,omitempty"`
							FControllerRevisionHash struct {
							} `json:"f:controller-revision-hash,omitempty"`
							FStatefulsetKubernetesIoPodName struct {
							} `json:"f:statefulset.kubernetes.io/pod-name,omitempty"`
							FK8SApp struct {
							} `json:"f:k8s-app,omitempty"`
							FPodTemplateHash struct {
							} `json:"f:pod-template-hash,omitempty"`
							FComponent struct {
							} `json:"f:component,omitempty"`
							FTier struct {
							} `json:"f:tier,omitempty"`
							FPodTemplateGeneration struct {
							} `json:"f:pod-template-generation,omitempty"`
							FAddonmanagerKubernetesIoMode struct {
							} `json:"f:addonmanager.kubernetes.io/mode,omitempty"`
							FIntegrationTest struct {
							} `json:"f:integration-test,omitempty"`
						} `json:"f:labels"`
						FAnnotations struct {
							Field1 struct {
							} `json:"."`
							FKubectlKubernetesIoLastAppliedConfiguration struct {
							} `json:"f:kubectl.kubernetes.io/last-applied-configuration,omitempty"`
							FKubeadmKubernetesIoEtcdAdvertiseClientUrls struct {
							} `json:"f:kubeadm.kubernetes.io/etcd.advertise-client-urls,omitempty"`
							FKubernetesIoConfigHash struct {
							} `json:"f:kubernetes.io/config.hash,omitempty"`
							FKubernetesIoConfigMirror struct {
							} `json:"f:kubernetes.io/config.mirror,omitempty"`
							FKubernetesIoConfigSeen struct {
							} `json:"f:kubernetes.io/config.seen,omitempty"`
							FKubernetesIoConfigSource struct {
							} `json:"f:kubernetes.io/config.source,omitempty"`
							FKubeadmKubernetesIoKubeApiserverAdvertiseAddressEndpoint struct {
							} `json:"f:kubeadm.kubernetes.io/kube-apiserver.advertise-address.endpoint,omitempty"`
						} `json:"f:annotations,omitempty"`
						FGenerateName struct {
						} `json:"f:generateName,omitempty"`
						FOwnerReferences struct {
							Field1 struct {
							} `json:"."`
							KUid028F808CF3F74C06A85551Cd1E3D73D8 struct {
							} `json:"k:{"uid":"028f808c-f3f7-4c06-a855-51cd1e3d73d8"},omitempty"`
							KUidFa8258A9264C4Aa288889D623C5F106E struct {
							} `json:"k:{"uid":"fa8258a9-264c-4aa2-8888-9d623c5f106e"},omitempty"`
							KUid07Cf72Ea67364733B94510A36B89804C struct {
							} `json:"k:{"uid":"07cf72ea-6736-4733-b945-10a36b89804c"},omitempty"`
							KUidBa3Fbfb27A0548B99A58466D27E24200 struct {
							} `json:"k:{"uid":"ba3fbfb2-7a05-48b9-9a58-466d27e24200"},omitempty"`
						} `json:"f:ownerReferences,omitempty"`
					} `json:"f:metadata,omitempty"`
					FSpec struct {
						FContainers struct {
							KNamePod1 struct {
								Field1 struct {
								} `json:"."`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FName struct {
								} `json:"f:name"`
								FResources struct {
								} `json:"f:resources"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
							} `json:"k:{"name":"pod1"},omitempty"`
							KNamePod2 struct {
								Field1 struct {
								} `json:"."`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FName struct {
								} `json:"f:name"`
								FResources struct {
								} `json:"f:resources"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
							} `json:"k:{"name":"pod2"},omitempty"`
							KNamePod4 struct {
								Field1 struct {
								} `json:"."`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FName struct {
								} `json:"f:name"`
								FResources struct {
								} `json:"f:resources"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
							} `json:"k:{"name":"pod4"},omitempty"`
							KNameReadyIfServicesReady struct {
								Field1 struct {
								} `json:"."`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FLivenessProbe struct {
									Field1 struct {
									} `json:"."`
									FExec struct {
										Field1 struct {
										} `json:"."`
										FCommand struct {
										} `json:"f:command"`
									} `json:"f:exec"`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:livenessProbe"`
								FName struct {
								} `json:"f:name"`
								FReadinessProbe struct {
									Field1 struct {
									} `json:"."`
									FExec struct {
										Field1 struct {
										} `json:"."`
										FCommand struct {
										} `json:"f:command"`
									} `json:"f:exec"`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:readinessProbe"`
								FResources struct {
								} `json:"f:resources"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
							} `json:"k:{"name":"ready-if-services-ready"},omitempty"`
							KNameNginx struct {
								Field1 struct {
								} `json:"."`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FName struct {
								} `json:"f:name"`
								FPorts struct {
									Field1 struct {
									} `json:"."`
									KContainerPort80ProtocolTCP struct {
										Field1 struct {
										} `json:"."`
										FContainerPort struct {
										} `json:"f:containerPort"`
										FName struct {
										} `json:"f:name"`
										FProtocol struct {
										} `json:"f:protocol"`
									} `json:"k:{"containerPort":80,"protocol":"TCP"}"`
								} `json:"f:ports"`
								FResources struct {
								} `json:"f:resources"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
								FVolumeMounts struct {
									Field1 struct {
									} `json:"."`
									KMountPathUsrShareNginxHtml struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
									} `json:"k:{"mountPath":"/usr/share/nginx/html"}"`
								} `json:"f:volumeMounts"`
							} `json:"k:{"name":"nginx"},omitempty"`
							KNameCoredns struct {
								Field1 struct {
								} `json:"."`
								FArgs struct {
								} `json:"f:args"`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FLivenessProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FInitialDelaySeconds struct {
									} `json:"f:initialDelaySeconds"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:livenessProbe"`
								FName struct {
								} `json:"f:name"`
								FPorts struct {
									Field1 struct {
									} `json:"."`
									KContainerPort53ProtocolTCP struct {
										Field1 struct {
										} `json:"."`
										FContainerPort struct {
										} `json:"f:containerPort"`
										FName struct {
										} `json:"f:name"`
										FProtocol struct {
										} `json:"f:protocol"`
									} `json:"k:{"containerPort":53,"protocol":"TCP"}"`
									KContainerPort53ProtocolUDP struct {
										Field1 struct {
										} `json:"."`
										FContainerPort struct {
										} `json:"f:containerPort"`
										FName struct {
										} `json:"f:name"`
										FProtocol struct {
										} `json:"f:protocol"`
									} `json:"k:{"containerPort":53,"protocol":"UDP"}"`
									KContainerPort9153ProtocolTCP struct {
										Field1 struct {
										} `json:"."`
										FContainerPort struct {
										} `json:"f:containerPort"`
										FName struct {
										} `json:"f:name"`
										FProtocol struct {
										} `json:"f:protocol"`
									} `json:"k:{"containerPort":9153,"protocol":"TCP"}"`
								} `json:"f:ports"`
								FReadinessProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:readinessProbe"`
								FResources struct {
									Field1 struct {
									} `json:"."`
									FLimits struct {
										Field1 struct {
										} `json:"."`
										FMemory struct {
										} `json:"f:memory"`
									} `json:"f:limits"`
									FRequests struct {
										Field1 struct {
										} `json:"."`
										FCpu struct {
										} `json:"f:cpu"`
										FMemory struct {
										} `json:"f:memory"`
									} `json:"f:requests"`
								} `json:"f:resources"`
								FSecurityContext struct {
									Field1 struct {
									} `json:"."`
									FAllowPrivilegeEscalation struct {
									} `json:"f:allowPrivilegeEscalation"`
									FCapabilities struct {
										Field1 struct {
										} `json:"."`
										FAdd struct {
										} `json:"f:add"`
										FDrop struct {
										} `json:"f:drop"`
									} `json:"f:capabilities"`
									FReadOnlyRootFilesystem struct {
									} `json:"f:readOnlyRootFilesystem"`
								} `json:"f:securityContext"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
								FVolumeMounts struct {
									Field1 struct {
									} `json:"."`
									KMountPathEtcCoredns struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
										FReadOnly struct {
										} `json:"f:readOnly"`
									} `json:"k:{"mountPath":"/etc/coredns"}"`
								} `json:"f:volumeMounts"`
							} `json:"k:{"name":"coredns"},omitempty"`
							KNameEtcd struct {
								Field1 struct {
								} `json:"."`
								FCommand struct {
								} `json:"f:command"`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FLivenessProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FHost struct {
										} `json:"f:host"`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FInitialDelaySeconds struct {
									} `json:"f:initialDelaySeconds"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:livenessProbe"`
								FName struct {
								} `json:"f:name"`
								FResources struct {
									Field1 struct {
									} `json:"."`
									FRequests struct {
										Field1 struct {
										} `json:"."`
										FCpu struct {
										} `json:"f:cpu"`
										FMemory struct {
										} `json:"f:memory"`
									} `json:"f:requests"`
								} `json:"f:resources"`
								FStartupProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FHost struct {
										} `json:"f:host"`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FInitialDelaySeconds struct {
									} `json:"f:initialDelaySeconds"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:startupProbe"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
								FVolumeMounts struct {
									Field1 struct {
									} `json:"."`
									KMountPathVarLibMinikubeCertsEtcd struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
									} `json:"k:{"mountPath":"/var/lib/minikube/certs/etcd"}"`
									KMountPathVarLibMinikubeEtcd struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
									} `json:"k:{"mountPath":"/var/lib/minikube/etcd"}"`
								} `json:"f:volumeMounts"`
							} `json:"k:{"name":"etcd"},omitempty"`
							KNameKubeApiserver struct {
								Field1 struct {
								} `json:"."`
								FCommand struct {
								} `json:"f:command"`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FLivenessProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FHost struct {
										} `json:"f:host"`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FInitialDelaySeconds struct {
									} `json:"f:initialDelaySeconds"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:livenessProbe"`
								FName struct {
								} `json:"f:name"`
								FReadinessProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FHost struct {
										} `json:"f:host"`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:readinessProbe"`
								FResources struct {
									Field1 struct {
									} `json:"."`
									FRequests struct {
										Field1 struct {
										} `json:"."`
										FCpu struct {
										} `json:"f:cpu"`
									} `json:"f:requests"`
								} `json:"f:resources"`
								FStartupProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FHost struct {
										} `json:"f:host"`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FInitialDelaySeconds struct {
									} `json:"f:initialDelaySeconds"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:startupProbe"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
								FVolumeMounts struct {
									Field1 struct {
									} `json:"."`
									KMountPathEtcSslCerts struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
										FReadOnly struct {
										} `json:"f:readOnly"`
									} `json:"k:{"mountPath":"/etc/ssl/certs"}"`
									KMountPathUsrShareCaCertificates struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
										FReadOnly struct {
										} `json:"f:readOnly"`
									} `json:"k:{"mountPath":"/usr/share/ca-certificates"}"`
									KMountPathVarLibMinikubeCerts struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
										FReadOnly struct {
										} `json:"f:readOnly"`
									} `json:"k:{"mountPath":"/var/lib/minikube/certs"}"`
								} `json:"f:volumeMounts"`
							} `json:"k:{"name":"kube-apiserver"},omitempty"`
							KNameKubeControllerManager struct {
								Field1 struct {
								} `json:"."`
								FCommand struct {
								} `json:"f:command"`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FLivenessProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FHost struct {
										} `json:"f:host"`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FInitialDelaySeconds struct {
									} `json:"f:initialDelaySeconds"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:livenessProbe"`
								FName struct {
								} `json:"f:name"`
								FResources struct {
									Field1 struct {
									} `json:"."`
									FRequests struct {
										Field1 struct {
										} `json:"."`
										FCpu struct {
										} `json:"f:cpu"`
									} `json:"f:requests"`
								} `json:"f:resources"`
								FStartupProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FHost struct {
										} `json:"f:host"`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FInitialDelaySeconds struct {
									} `json:"f:initialDelaySeconds"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:startupProbe"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
								FVolumeMounts struct {
									Field1 struct {
									} `json:"."`
									KMountPathEtcKubernetesControllerManagerConf struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
										FReadOnly struct {
										} `json:"f:readOnly"`
									} `json:"k:{"mountPath":"/etc/kubernetes/controller-manager.conf"}"`
									KMountPathEtcSslCerts struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
										FReadOnly struct {
										} `json:"f:readOnly"`
									} `json:"k:{"mountPath":"/etc/ssl/certs"}"`
									KMountPathUsrLibexecKubernetesKubeletPluginsVolumeExec struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
									} `json:"k:{"mountPath":"/usr/libexec/kubernetes/kubelet-plugins/volume/exec"}"`
									KMountPathUsrShareCaCertificates struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
										FReadOnly struct {
										} `json:"f:readOnly"`
									} `json:"k:{"mountPath":"/usr/share/ca-certificates"}"`
									KMountPathVarLibMinikubeCerts struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
										FReadOnly struct {
										} `json:"f:readOnly"`
									} `json:"k:{"mountPath":"/var/lib/minikube/certs"}"`
								} `json:"f:volumeMounts"`
							} `json:"k:{"name":"kube-controller-manager"},omitempty"`
							KNameKubeProxy struct {
								Field1 struct {
								} `json:"."`
								FCommand struct {
								} `json:"f:command"`
								FEnv struct {
									Field1 struct {
									} `json:"."`
									KNameNODENAME struct {
										Field1 struct {
										} `json:"."`
										FName struct {
										} `json:"f:name"`
										FValueFrom struct {
											Field1 struct {
											} `json:"."`
											FFieldRef struct {
											} `json:"f:fieldRef"`
										} `json:"f:valueFrom"`
									} `json:"k:{"name":"NODE_NAME"}"`
								} `json:"f:env"`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FName struct {
								} `json:"f:name"`
								FResources struct {
								} `json:"f:resources"`
								FSecurityContext struct {
									Field1 struct {
									} `json:"."`
									FPrivileged struct {
									} `json:"f:privileged"`
								} `json:"f:securityContext"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
								FVolumeMounts struct {
									Field1 struct {
									} `json:"."`
									KMountPathLibModules struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
										FReadOnly struct {
										} `json:"f:readOnly"`
									} `json:"k:{"mountPath":"/lib/modules"}"`
									KMountPathRunXtablesLock struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
									} `json:"k:{"mountPath":"/run/xtables.lock"}"`
									KMountPathVarLibKubeProxy struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
									} `json:"k:{"mountPath":"/var/lib/kube-proxy"}"`
								} `json:"f:volumeMounts"`
							} `json:"k:{"name":"kube-proxy"},omitempty"`
							KNameKubeScheduler struct {
								Field1 struct {
								} `json:"."`
								FCommand struct {
								} `json:"f:command"`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FLivenessProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FHost struct {
										} `json:"f:host"`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FInitialDelaySeconds struct {
									} `json:"f:initialDelaySeconds"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:livenessProbe"`
								FName struct {
								} `json:"f:name"`
								FResources struct {
									Field1 struct {
									} `json:"."`
									FRequests struct {
										Field1 struct {
										} `json:"."`
										FCpu struct {
										} `json:"f:cpu"`
									} `json:"f:requests"`
								} `json:"f:resources"`
								FStartupProbe struct {
									Field1 struct {
									} `json:"."`
									FFailureThreshold struct {
									} `json:"f:failureThreshold"`
									FHttpGet struct {
										Field1 struct {
										} `json:"."`
										FHost struct {
										} `json:"f:host"`
										FPath struct {
										} `json:"f:path"`
										FPort struct {
										} `json:"f:port"`
										FScheme struct {
										} `json:"f:scheme"`
									} `json:"f:httpGet"`
									FInitialDelaySeconds struct {
									} `json:"f:initialDelaySeconds"`
									FPeriodSeconds struct {
									} `json:"f:periodSeconds"`
									FSuccessThreshold struct {
									} `json:"f:successThreshold"`
									FTimeoutSeconds struct {
									} `json:"f:timeoutSeconds"`
								} `json:"f:startupProbe"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
								FVolumeMounts struct {
									Field1 struct {
									} `json:"."`
									KMountPathEtcKubernetesSchedulerConf struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
										FReadOnly struct {
										} `json:"f:readOnly"`
									} `json:"k:{"mountPath":"/etc/kubernetes/scheduler.conf"}"`
								} `json:"f:volumeMounts"`
							} `json:"k:{"name":"kube-scheduler"},omitempty"`
							KNameStorageProvisioner struct {
								Field1 struct {
								} `json:"."`
								FCommand struct {
								} `json:"f:command"`
								FImage struct {
								} `json:"f:image"`
								FImagePullPolicy struct {
								} `json:"f:imagePullPolicy"`
								FName struct {
								} `json:"f:name"`
								FResources struct {
								} `json:"f:resources"`
								FTerminationMessagePath struct {
								} `json:"f:terminationMessagePath"`
								FTerminationMessagePolicy struct {
								} `json:"f:terminationMessagePolicy"`
								FVolumeMounts struct {
									Field1 struct {
									} `json:"."`
									KMountPathTmp struct {
										Field1 struct {
										} `json:"."`
										FMountPath struct {
										} `json:"f:mountPath"`
										FName struct {
										} `json:"f:name"`
									} `json:"k:{"mountPath":"/tmp"}"`
								} `json:"f:volumeMounts"`
							} `json:"k:{"name":"storage-provisioner"},omitempty"`
						} `json:"f:containers"`
						FDnsPolicy struct {
						} `json:"f:dnsPolicy"`
						FEnableServiceLinks struct {
						} `json:"f:enableServiceLinks"`
						FRestartPolicy struct {
						} `json:"f:restartPolicy"`
						FSchedulerName struct {
						} `json:"f:schedulerName"`
						FSecurityContext struct {
							Field1 struct {
							} `json:".,omitempty"`
							FSeccompProfile struct {
								Field1 struct {
								} `json:"."`
								FType struct {
								} `json:"f:type"`
							} `json:"f:seccompProfile,omitempty"`
						} `json:"f:securityContext"`
						FTerminationGracePeriodSeconds struct {
						} `json:"f:terminationGracePeriodSeconds"`
						FHostname struct {
						} `json:"f:hostname,omitempty"`
						FSubdomain struct {
						} `json:"f:subdomain,omitempty"`
						FVolumes struct {
							Field1 struct {
							} `json:"."`
							KNameWww struct {
								Field1 struct {
								} `json:"."`
								FName struct {
								} `json:"f:name"`
								FPersistentVolumeClaim struct {
									Field1 struct {
									} `json:"."`
									FClaimName struct {
									} `json:"f:claimName"`
								} `json:"f:persistentVolumeClaim"`
							} `json:"k:{"name":"www"},omitempty"`
							KNameConfigVolume struct {
								Field1 struct {
								} `json:"."`
								FConfigMap struct {
									Field1 struct {
									} `json:"."`
									FDefaultMode struct {
									} `json:"f:defaultMode"`
									FItems struct {
									} `json:"f:items"`
									FName struct {
									} `json:"f:name"`
								} `json:"f:configMap"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"config-volume"},omitempty"`
							KNameEtcdCerts struct {
								Field1 struct {
								} `json:"."`
								FHostPath struct {
									Field1 struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FType struct {
									} `json:"f:type"`
								} `json:"f:hostPath"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"etcd-certs"},omitempty"`
							KNameEtcdData struct {
								Field1 struct {
								} `json:"."`
								FHostPath struct {
									Field1 struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FType struct {
									} `json:"f:type"`
								} `json:"f:hostPath"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"etcd-data"},omitempty"`
							KNameCaCerts struct {
								Field1 struct {
								} `json:"."`
								FHostPath struct {
									Field1 struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FType struct {
									} `json:"f:type"`
								} `json:"f:hostPath"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"ca-certs"},omitempty"`
							KNameK8SCerts struct {
								Field1 struct {
								} `json:"."`
								FHostPath struct {
									Field1 struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FType struct {
									} `json:"f:type"`
								} `json:"f:hostPath"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"k8s-certs"},omitempty"`
							KNameUsrShareCaCertificates struct {
								Field1 struct {
								} `json:"."`
								FHostPath struct {
									Field1 struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FType struct {
									} `json:"f:type"`
								} `json:"f:hostPath"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"usr-share-ca-certificates"},omitempty"`
							KNameFlexvolumeDir struct {
								Field1 struct {
								} `json:"."`
								FHostPath struct {
									Field1 struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FType struct {
									} `json:"f:type"`
								} `json:"f:hostPath"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"flexvolume-dir"},omitempty"`
							KNameKubeconfig struct {
								Field1 struct {
								} `json:"."`
								FHostPath struct {
									Field1 struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FType struct {
									} `json:"f:type"`
								} `json:"f:hostPath"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"kubeconfig"},omitempty"`
							KNameKubeProxy struct {
								Field1 struct {
								} `json:"."`
								FConfigMap struct {
									Field1 struct {
									} `json:"."`
									FDefaultMode struct {
									} `json:"f:defaultMode"`
									FName struct {
									} `json:"f:name"`
								} `json:"f:configMap"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"kube-proxy"},omitempty"`
							KNameLibModules struct {
								Field1 struct {
								} `json:"."`
								FHostPath struct {
									Field1 struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FType struct {
									} `json:"f:type"`
								} `json:"f:hostPath"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"lib-modules"},omitempty"`
							KNameXtablesLock struct {
								Field1 struct {
								} `json:"."`
								FHostPath struct {
									Field1 struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FType struct {
									} `json:"f:type"`
								} `json:"f:hostPath"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"xtables-lock"},omitempty"`
							KNameTmp struct {
								Field1 struct {
								} `json:"."`
								FHostPath struct {
									Field1 struct {
									} `json:"."`
									FPath struct {
									} `json:"f:path"`
									FType struct {
									} `json:"f:type"`
								} `json:"f:hostPath"`
								FName struct {
								} `json:"f:name"`
							} `json:"k:{"name":"tmp"},omitempty"`
						} `json:"f:volumes,omitempty"`
						FNodeSelector struct {
						} `json:"f:nodeSelector,omitempty"`
						FPriorityClassName struct {
						} `json:"f:priorityClassName,omitempty"`
						FServiceAccount struct {
						} `json:"f:serviceAccount,omitempty"`
						FServiceAccountName struct {
						} `json:"f:serviceAccountName,omitempty"`
						FTolerations struct {
						} `json:"f:tolerations,omitempty"`
						FHostNetwork struct {
						} `json:"f:hostNetwork,omitempty"`
						FNodeName struct {
						} `json:"f:nodeName,omitempty"`
						FAffinity struct {
							Field1 struct {
							} `json:"."`
							FNodeAffinity struct {
								Field1 struct {
								} `json:"."`
								FRequiredDuringSchedulingIgnoredDuringExecution struct {
								} `json:"f:requiredDuringSchedulingIgnoredDuringExecution"`
							} `json:"f:nodeAffinity"`
						} `json:"f:affinity,omitempty"`
					} `json:"f:spec,omitempty"`
					FStatus struct {
						FConditions struct {
							KTypeContainersReady struct {
								Field1 struct {
								} `json:"."`
								FLastProbeTime struct {
								} `json:"f:lastProbeTime"`
								FLastTransitionTime struct {
								} `json:"f:lastTransitionTime"`
								FStatus struct {
								} `json:"f:status"`
								FType struct {
								} `json:"f:type"`
								FMessage struct {
								} `json:"f:message,omitempty"`
								FReason struct {
								} `json:"f:reason,omitempty"`
							} `json:"k:{"type":"ContainersReady"},omitempty"`
							KTypeInitialized struct {
								Field1 struct {
								} `json:"."`
								FLastProbeTime struct {
								} `json:"f:lastProbeTime"`
								FLastTransitionTime struct {
								} `json:"f:lastTransitionTime"`
								FStatus struct {
								} `json:"f:status"`
								FType struct {
								} `json:"f:type"`
							} `json:"k:{"type":"Initialized"},omitempty"`
							KTypeReady struct {
								Field1 struct {
								} `json:"."`
								FLastProbeTime struct {
								} `json:"f:lastProbeTime"`
								FLastTransitionTime struct {
								} `json:"f:lastTransitionTime"`
								FStatus struct {
								} `json:"f:status"`
								FType struct {
								} `json:"f:type"`
								FMessage struct {
								} `json:"f:message,omitempty"`
								FReason struct {
								} `json:"f:reason,omitempty"`
							} `json:"k:{"type":"Ready"},omitempty"`
							Field4 struct {
							} `json:".,omitempty"`
							KTypePodScheduled struct {
								Field1 struct {
								} `json:"."`
								FLastProbeTime struct {
								} `json:"f:lastProbeTime"`
								FLastTransitionTime struct {
								} `json:"f:lastTransitionTime"`
								FMessage struct {
								} `json:"f:message,omitempty"`
								FReason struct {
								} `json:"f:reason,omitempty"`
								FStatus struct {
								} `json:"f:status"`
								FType struct {
								} `json:"f:type"`
							} `json:"k:{"type":"PodScheduled"},omitempty"`
						} `json:"f:conditions"`
						FContainerStatuses struct {
						} `json:"f:containerStatuses,omitempty"`
						FHostIP struct {
						} `json:"f:hostIP,omitempty"`
						FPhase struct {
						} `json:"f:phase,omitempty"`
						FPodIP struct {
						} `json:"f:podIP,omitempty"`
						FPodIPs struct {
							Field1 struct {
							} `json:"."`
							KIp1721703 struct {
								Field1 struct {
								} `json:"."`
								FIp struct {
								} `json:"f:ip"`
							} `json:"k:{"ip":"172.17.0.3"},omitempty"`
							KIp1721706 struct {
								Field1 struct {
								} `json:"."`
								FIp struct {
								} `json:"f:ip"`
							} `json:"k:{"ip":"172.17.0.6"},omitempty"`
							KIp1721705 struct {
								Field1 struct {
								} `json:"."`
								FIp struct {
								} `json:"f:ip"`
							} `json:"k:{"ip":"172.17.0.5"},omitempty"`
							KIp1721704 struct {
								Field1 struct {
								} `json:"."`
								FIp struct {
								} `json:"f:ip"`
							} `json:"k:{"ip":"172.17.0.4"},omitempty"`
							KIp1721707 struct {
								Field1 struct {
								} `json:"."`
								FIp struct {
								} `json:"f:ip"`
							} `json:"k:{"ip":"172.17.0.7"},omitempty"`
							KIp1721702 struct {
								Field1 struct {
								} `json:"."`
								FIp struct {
								} `json:"f:ip"`
							} `json:"k:{"ip":"172.17.0.2"},omitempty"`
							KIp1921686411 struct {
								Field1 struct {
								} `json:"."`
								FIp struct {
								} `json:"f:ip"`
							} `json:"k:{"ip":"192.168.64.11"},omitempty"`
						} `json:"f:podIPs,omitempty"`
						FStartTime struct {
						} `json:"f:startTime,omitempty"`
					} `json:"f:status,omitempty"`
				} `json:"fieldsV1"`
				Subresource string `json:"subresource,omitempty"`
			} `json:"managedFields"`
			Annotations struct {
				KubectlKubernetesIoLastAppliedConfiguration              string    `json:"kubectl.kubernetes.io/last-applied-configuration,omitempty"`
				KubeadmKubernetesIoEtcdAdvertiseClientUrls               string    `json:"kubeadm.kubernetes.io/etcd.advertise-client-urls,omitempty"`
				KubernetesIoConfigHash                                   string    `json:"kubernetes.io/config.hash,omitempty"`
				KubernetesIoConfigMirror                                 string    `json:"kubernetes.io/config.mirror,omitempty"`
				KubernetesIoConfigSeen                                   time.Time `json:"kubernetes.io/config.seen,omitempty"`
				KubernetesIoConfigSource                                 string    `json:"kubernetes.io/config.source,omitempty"`
				SeccompSecurityAlphaKubernetesIoPod                      string    `json:"seccomp.security.alpha.kubernetes.io/pod,omitempty"`
				KubeadmKubernetesIoKubeApiserverAdvertiseAddressEndpoint string    `json:"kubeadm.kubernetes.io/kube-apiserver.advertise-address.endpoint,omitempty"`
			} `json:"annotations,omitempty"`
			GenerateName    string `json:"generateName,omitempty"`
			OwnerReferences []struct {
				ApiVersion         string `json:"apiVersion"`
				Kind               string `json:"kind"`
				Name               string `json:"name"`
				Uid                string `json:"uid"`
				Controller         bool   `json:"controller"`
				BlockOwnerDeletion bool   `json:"blockOwnerDeletion,omitempty"`
			} `json:"ownerReferences,omitempty"`
		} `json:"metadata"`
		Spec struct {
			Volumes []struct {
				Name      string `json:"name"`
				Projected struct {
					Sources []struct {
						ServiceAccountToken struct {
							ExpirationSeconds int    `json:"expirationSeconds"`
							Path              string `json:"path"`
						} `json:"serviceAccountToken,omitempty"`
						ConfigMap struct {
							Name  string `json:"name"`
							Items []struct {
								Key  string `json:"key"`
								Path string `json:"path"`
							} `json:"items"`
						} `json:"configMap,omitempty"`
						DownwardAPI struct {
							Items []struct {
								Path     string `json:"path"`
								FieldRef struct {
									ApiVersion string `json:"apiVersion"`
									FieldPath  string `json:"fieldPath"`
								} `json:"fieldRef"`
							} `json:"items"`
						} `json:"downwardAPI,omitempty"`
					} `json:"sources"`
					DefaultMode int `json:"defaultMode"`
				} `json:"projected,omitempty"`
				PersistentVolumeClaim struct {
					ClaimName string `json:"claimName"`
				} `json:"persistentVolumeClaim,omitempty"`
				ConfigMap struct {
					Name  string `json:"name"`
					Items []struct {
						Key  string `json:"key"`
						Path string `json:"path"`
					} `json:"items,omitempty"`
					DefaultMode int `json:"defaultMode"`
				} `json:"configMap,omitempty"`
				HostPath struct {
					Path string `json:"path"`
					Type string `json:"type"`
				} `json:"hostPath,omitempty"`
			} `json:"volumes"`
			Containers []struct {
				Name      string `json:"name"`
				Image     string `json:"image"`
				Resources struct {
					Limits struct {
						Memory string `json:"memory"`
					} `json:"limits,omitempty"`
					Requests struct {
						Cpu    string `json:"cpu"`
						Memory string `json:"memory,omitempty"`
					} `json:"requests,omitempty"`
				} `json:"resources"`
				VolumeMounts []struct {
					Name      string `json:"name"`
					ReadOnly  bool   `json:"readOnly,omitempty"`
					MountPath string `json:"mountPath"`
				} `json:"volumeMounts"`
				TerminationMessagePath   string `json:"terminationMessagePath"`
				TerminationMessagePolicy string `json:"terminationMessagePolicy"`
				ImagePullPolicy          string `json:"imagePullPolicy"`
				LivenessProbe            struct {
					Exec struct {
						Command []string `json:"command"`
					} `json:"exec,omitempty"`
					TimeoutSeconds   int `json:"timeoutSeconds"`
					PeriodSeconds    int `json:"periodSeconds"`
					SuccessThreshold int `json:"successThreshold"`
					FailureThreshold int `json:"failureThreshold"`
					HttpGet          struct {
						Path   string `json:"path"`
						Port   int    `json:"port"`
						Scheme string `json:"scheme"`
						Host   string `json:"host,omitempty"`
					} `json:"httpGet,omitempty"`
					InitialDelaySeconds int `json:"initialDelaySeconds,omitempty"`
				} `json:"livenessProbe,omitempty"`
				ReadinessProbe struct {
					Exec struct {
						Command []string `json:"command"`
					} `json:"exec,omitempty"`
					TimeoutSeconds   int `json:"timeoutSeconds"`
					PeriodSeconds    int `json:"periodSeconds"`
					SuccessThreshold int `json:"successThreshold"`
					FailureThreshold int `json:"failureThreshold"`
					HttpGet          struct {
						Path   string `json:"path"`
						Port   int    `json:"port"`
						Scheme string `json:"scheme"`
						Host   string `json:"host,omitempty"`
					} `json:"httpGet,omitempty"`
				} `json:"readinessProbe,omitempty"`
				Ports []struct {
					Name          string `json:"name"`
					ContainerPort int    `json:"containerPort"`
					Protocol      string `json:"protocol"`
				} `json:"ports,omitempty"`
				Args            []string `json:"args,omitempty"`
				SecurityContext struct {
					Capabilities struct {
						Add  []string `json:"add"`
						Drop []string `json:"drop"`
					} `json:"capabilities,omitempty"`
					ReadOnlyRootFilesystem   bool `json:"readOnlyRootFilesystem,omitempty"`
					AllowPrivilegeEscalation bool `json:"allowPrivilegeEscalation,omitempty"`
					Privileged               bool `json:"privileged,omitempty"`
				} `json:"securityContext,omitempty"`
				Command      []string `json:"command,omitempty"`
				StartupProbe struct {
					HttpGet struct {
						Path   string `json:"path"`
						Port   int    `json:"port"`
						Host   string `json:"host"`
						Scheme string `json:"scheme"`
					} `json:"httpGet"`
					InitialDelaySeconds int `json:"initialDelaySeconds"`
					TimeoutSeconds      int `json:"timeoutSeconds"`
					PeriodSeconds       int `json:"periodSeconds"`
					SuccessThreshold    int `json:"successThreshold"`
					FailureThreshold    int `json:"failureThreshold"`
				} `json:"startupProbe,omitempty"`
				Env []struct {
					Name      string `json:"name"`
					ValueFrom struct {
						FieldRef struct {
							ApiVersion string `json:"apiVersion"`
							FieldPath  string `json:"fieldPath"`
						} `json:"fieldRef"`
					} `json:"valueFrom"`
				} `json:"env,omitempty"`
			} `json:"containers"`
			RestartPolicy                 string `json:"restartPolicy"`
			TerminationGracePeriodSeconds int    `json:"terminationGracePeriodSeconds"`
			DnsPolicy                     string `json:"dnsPolicy"`
			ServiceAccountName            string `json:"serviceAccountName,omitempty"`
			ServiceAccount                string `json:"serviceAccount,omitempty"`
			NodeName                      string `json:"nodeName"`
			SecurityContext               struct {
				SeccompProfile struct {
					Type string `json:"type"`
				} `json:"seccompProfile,omitempty"`
			} `json:"securityContext"`
			SchedulerName string `json:"schedulerName"`
			Tolerations   []struct {
				Key               string `json:"key,omitempty"`
				Operator          string `json:"operator,omitempty"`
				Effect            string `json:"effect,omitempty"`
				TolerationSeconds int    `json:"tolerationSeconds,omitempty"`
			} `json:"tolerations"`
			Priority           int    `json:"priority"`
			EnableServiceLinks bool   `json:"enableServiceLinks"`
			PreemptionPolicy   string `json:"preemptionPolicy"`
			Hostname           string `json:"hostname,omitempty"`
			Subdomain          string `json:"subdomain,omitempty"`
			NodeSelector       struct {
				KubernetesIoOs string `json:"kubernetes.io/os"`
			} `json:"nodeSelector,omitempty"`
			PriorityClassName string `json:"priorityClassName,omitempty"`
			HostNetwork       bool   `json:"hostNetwork,omitempty"`
			Affinity          struct {
				NodeAffinity struct {
					RequiredDuringSchedulingIgnoredDuringExecution struct {
						NodeSelectorTerms []struct {
							MatchFields []struct {
								Key      string   `json:"key"`
								Operator string   `json:"operator"`
								Values   []string `json:"values"`
							} `json:"matchFields"`
						} `json:"nodeSelectorTerms"`
					} `json:"requiredDuringSchedulingIgnoredDuringExecution"`
				} `json:"nodeAffinity"`
			} `json:"affinity,omitempty"`
		} `json:"spec"`
		Status struct {
			Phase      string `json:"phase"`
			Conditions []struct {
				Type               string      `json:"type"`
				Status             string      `json:"status"`
				LastProbeTime      interface{} `json:"lastProbeTime"`
				LastTransitionTime time.Time   `json:"lastTransitionTime"`
				Reason             string      `json:"reason,omitempty"`
				Message            string      `json:"message,omitempty"`
			} `json:"conditions"`
			HostIP string `json:"hostIP"`
			PodIP  string `json:"podIP"`
			PodIPs []struct {
				Ip string `json:"ip"`
			} `json:"podIPs"`
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
				Started      bool   `json:"started"`
			} `json:"containerStatuses"`
			QosClass string `json:"qosClass"`
		} `json:"status"`
	} `json:"items"`
}
