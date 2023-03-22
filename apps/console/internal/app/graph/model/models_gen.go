// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type AppSpec struct {
	Frozen         *bool                  `json:"frozen"`
	Hpa            *AppSpecHpa            `json:"hpa"`
	NodeSelector   map[string]interface{} `json:"nodeSelector"`
	Region         string                 `json:"region"`
	ServiceAccount *string                `json:"serviceAccount"`
	AccountName    string                 `json:"accountName"`
	Containers     []*AppSpecContainers   `json:"containers"`
	Interception   *AppSpecInterception   `json:"interception"`
	Replicas       *int                   `json:"replicas"`
	Services       []*AppSpecServices     `json:"services"`
	Tolerations    []*AppSpecTolerations  `json:"tolerations"`
}

type AppSpecContainers struct {
	Image           string                           `json:"image"`
	ImagePullPolicy *string                          `json:"imagePullPolicy"`
	LivenessProbe   *AppSpecContainersLivenessProbe  `json:"livenessProbe"`
	ReadinessProbe  *AppSpecContainersReadinessProbe `json:"readinessProbe"`
	ResourceMemory  *AppSpecContainersResourceMemory `json:"resourceMemory"`
	Command         []*string                        `json:"command"`
	Env             []*AppSpecContainersEnv          `json:"env"`
	EnvFrom         []*AppSpecContainersEnvFrom      `json:"envFrom"`
	Volumes         []*AppSpecContainersVolumes      `json:"volumes"`
	Args            []*string                        `json:"args"`
	Name            string                           `json:"name"`
	ResourceCPU     *AppSpecContainersResourceCPU    `json:"resourceCpu"`
}

type AppSpecContainersEnv struct {
	Key     string  `json:"key"`
	RefKey  *string `json:"refKey"`
	RefName *string `json:"refName"`
	Type    *string `json:"type"`
	Value   *string `json:"value"`
}

type AppSpecContainersEnvFrom struct {
	RefName string `json:"refName"`
	Type    string `json:"type"`
}

type AppSpecContainersEnvFromIn struct {
	RefName string `json:"refName"`
	Type    string `json:"type"`
}

type AppSpecContainersEnvIn struct {
	Key     string  `json:"key"`
	RefKey  *string `json:"refKey"`
	RefName *string `json:"refName"`
	Type    *string `json:"type"`
	Value   *string `json:"value"`
}

type AppSpecContainersIn struct {
	Image           string                             `json:"image"`
	ImagePullPolicy *string                            `json:"imagePullPolicy"`
	LivenessProbe   *AppSpecContainersLivenessProbeIn  `json:"livenessProbe"`
	ReadinessProbe  *AppSpecContainersReadinessProbeIn `json:"readinessProbe"`
	ResourceMemory  *AppSpecContainersResourceMemoryIn `json:"resourceMemory"`
	Command         []*string                          `json:"command"`
	Env             []*AppSpecContainersEnvIn          `json:"env"`
	EnvFrom         []*AppSpecContainersEnvFromIn      `json:"envFrom"`
	Volumes         []*AppSpecContainersVolumesIn      `json:"volumes"`
	Args            []*string                          `json:"args"`
	Name            string                             `json:"name"`
	ResourceCPU     *AppSpecContainersResourceCPUIn    `json:"resourceCpu"`
}

type AppSpecContainersLivenessProbe struct {
	TCP              *AppSpecContainersLivenessProbeTCP     `json:"tcp"`
	Type             string                                 `json:"type"`
	FailureThreshold *int                                   `json:"failureThreshold"`
	HTTPGet          *AppSpecContainersLivenessProbeHTTPGet `json:"httpGet"`
	InitialDelay     *int                                   `json:"initialDelay"`
	Interval         *int                                   `json:"interval"`
	Shell            *AppSpecContainersLivenessProbeShell   `json:"shell"`
}

type AppSpecContainersLivenessProbeHTTPGet struct {
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
	HTTPHeaders map[string]interface{} `json:"httpHeaders"`
}

type AppSpecContainersLivenessProbeHTTPGetIn struct {
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
	HTTPHeaders map[string]interface{} `json:"httpHeaders"`
}

type AppSpecContainersLivenessProbeIn struct {
	TCP              *AppSpecContainersLivenessProbeTCPIn     `json:"tcp"`
	Type             string                                   `json:"type"`
	FailureThreshold *int                                     `json:"failureThreshold"`
	HTTPGet          *AppSpecContainersLivenessProbeHTTPGetIn `json:"httpGet"`
	InitialDelay     *int                                     `json:"initialDelay"`
	Interval         *int                                     `json:"interval"`
	Shell            *AppSpecContainersLivenessProbeShellIn   `json:"shell"`
}

type AppSpecContainersLivenessProbeShell struct {
	Command []*string `json:"command"`
}

type AppSpecContainersLivenessProbeShellIn struct {
	Command []*string `json:"command"`
}

type AppSpecContainersLivenessProbeTCP struct {
	Port int `json:"port"`
}

type AppSpecContainersLivenessProbeTCPIn struct {
	Port int `json:"port"`
}

type AppSpecContainersReadinessProbe struct {
	InitialDelay     *int                                    `json:"initialDelay"`
	Interval         *int                                    `json:"interval"`
	Shell            *AppSpecContainersReadinessProbeShell   `json:"shell"`
	TCP              *AppSpecContainersReadinessProbeTCP     `json:"tcp"`
	Type             string                                  `json:"type"`
	FailureThreshold *int                                    `json:"failureThreshold"`
	HTTPGet          *AppSpecContainersReadinessProbeHTTPGet `json:"httpGet"`
}

type AppSpecContainersReadinessProbeHTTPGet struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type AppSpecContainersReadinessProbeHTTPGetIn struct {
	HTTPHeaders map[string]interface{} `json:"httpHeaders"`
	Path        string                 `json:"path"`
	Port        int                    `json:"port"`
}

type AppSpecContainersReadinessProbeIn struct {
	InitialDelay     *int                                      `json:"initialDelay"`
	Interval         *int                                      `json:"interval"`
	Shell            *AppSpecContainersReadinessProbeShellIn   `json:"shell"`
	TCP              *AppSpecContainersReadinessProbeTCPIn     `json:"tcp"`
	Type             string                                    `json:"type"`
	FailureThreshold *int                                      `json:"failureThreshold"`
	HTTPGet          *AppSpecContainersReadinessProbeHTTPGetIn `json:"httpGet"`
}

type AppSpecContainersReadinessProbeShell struct {
	Command []*string `json:"command"`
}

type AppSpecContainersReadinessProbeShellIn struct {
	Command []*string `json:"command"`
}

type AppSpecContainersReadinessProbeTCP struct {
	Port int `json:"port"`
}

type AppSpecContainersReadinessProbeTCPIn struct {
	Port int `json:"port"`
}

type AppSpecContainersResourceCPU struct {
	Max *string `json:"max"`
	Min *string `json:"min"`
}

type AppSpecContainersResourceCPUIn struct {
	Max *string `json:"max"`
	Min *string `json:"min"`
}

type AppSpecContainersResourceMemory struct {
	Max *string `json:"max"`
	Min *string `json:"min"`
}

type AppSpecContainersResourceMemoryIn struct {
	Max *string `json:"max"`
	Min *string `json:"min"`
}

type AppSpecContainersVolumes struct {
	Items     []*AppSpecContainersVolumesItems `json:"items"`
	MountPath string                           `json:"mountPath"`
	RefName   string                           `json:"refName"`
	Type      string                           `json:"type"`
}

type AppSpecContainersVolumesIn struct {
	Items     []*AppSpecContainersVolumesItemsIn `json:"items"`
	MountPath string                             `json:"mountPath"`
	RefName   string                             `json:"refName"`
	Type      string                             `json:"type"`
}

type AppSpecContainersVolumesItems struct {
	FileName *string `json:"fileName"`
	Key      string  `json:"key"`
}

type AppSpecContainersVolumesItemsIn struct {
	FileName *string `json:"fileName"`
	Key      string  `json:"key"`
}

type AppSpecHpa struct {
	MaxReplicas     *int  `json:"maxReplicas"`
	MinReplicas     *int  `json:"minReplicas"`
	ThresholdCPU    *int  `json:"thresholdCpu"`
	ThresholdMemory *int  `json:"thresholdMemory"`
	Enabled         *bool `json:"enabled"`
}

type AppSpecHpaIn struct {
	MaxReplicas     *int  `json:"maxReplicas"`
	MinReplicas     *int  `json:"minReplicas"`
	ThresholdCPU    *int  `json:"thresholdCpu"`
	ThresholdMemory *int  `json:"thresholdMemory"`
	Enabled         *bool `json:"enabled"`
}

type AppSpecIn struct {
	Frozen         *bool                   `json:"frozen"`
	Hpa            *AppSpecHpaIn           `json:"hpa"`
	NodeSelector   map[string]interface{}  `json:"nodeSelector"`
	Region         string                  `json:"region"`
	ServiceAccount *string                 `json:"serviceAccount"`
	AccountName    string                  `json:"accountName"`
	Containers     []*AppSpecContainersIn  `json:"containers"`
	Interception   *AppSpecInterceptionIn  `json:"interception"`
	Replicas       *int                    `json:"replicas"`
	Services       []*AppSpecServicesIn    `json:"services"`
	Tolerations    []*AppSpecTolerationsIn `json:"tolerations"`
}

type AppSpecInterception struct {
	Enabled   *bool  `json:"enabled"`
	ForDevice string `json:"forDevice"`
}

type AppSpecInterceptionIn struct {
	Enabled   *bool  `json:"enabled"`
	ForDevice string `json:"forDevice"`
}

type AppSpecServices struct {
	Type       *string `json:"type"`
	Name       *string `json:"name"`
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort"`
}

type AppSpecServicesIn struct {
	Type       *string `json:"type"`
	Name       *string `json:"name"`
	Port       int     `json:"port"`
	TargetPort *int    `json:"targetPort"`
}

type AppSpecTolerations struct {
	Effect            *string `json:"effect"`
	Key               *string `json:"key"`
	Operator          *string `json:"operator"`
	TolerationSeconds *int    `json:"tolerationSeconds"`
	Value             *string `json:"value"`
}

type AppSpecTolerationsIn struct {
	Effect            *string `json:"effect"`
	Key               *string `json:"key"`
	Operator          *string `json:"operator"`
	TolerationSeconds *int    `json:"tolerationSeconds"`
	Value             *string `json:"value"`
}

type Check struct {
	Status     *bool   `json:"status"`
	Message    *string `json:"message"`
	Generation *int    `json:"generation"`
}

type ManagedResourceSpec struct {
	MresKind *ManagedResourceSpecMresKind `json:"mresKind"`
	MsvcRef  *ManagedResourceSpecMsvcRef  `json:"msvcRef"`
	Inputs   map[string]interface{}       `json:"inputs"`
}

type ManagedResourceSpecIn struct {
	MresKind *ManagedResourceSpecMresKindIn `json:"mresKind"`
	MsvcRef  *ManagedResourceSpecMsvcRefIn  `json:"msvcRef"`
	Inputs   map[string]interface{}         `json:"inputs"`
}

type ManagedResourceSpecMresKind struct {
	Kind string `json:"kind"`
}

type ManagedResourceSpecMresKindIn struct {
	Kind string `json:"kind"`
}

type ManagedResourceSpecMsvcRef struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind"`
	Name       string  `json:"name"`
}

type ManagedResourceSpecMsvcRefIn struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind"`
	Name       string  `json:"name"`
}

type ManagedServiceSpec struct {
	Region       string                           `json:"region"`
	Tolerations  []*ManagedServiceSpecTolerations `json:"tolerations"`
	Inputs       map[string]interface{}           `json:"inputs"`
	MsvcKind     *ManagedServiceSpecMsvcKind      `json:"msvcKind"`
	NodeSelector map[string]interface{}           `json:"nodeSelector"`
}

type ManagedServiceSpecIn struct {
	Region       string                             `json:"region"`
	Tolerations  []*ManagedServiceSpecTolerationsIn `json:"tolerations"`
	Inputs       map[string]interface{}             `json:"inputs"`
	MsvcKind     *ManagedServiceSpecMsvcKindIn      `json:"msvcKind"`
	NodeSelector map[string]interface{}             `json:"nodeSelector"`
}

type ManagedServiceSpecMsvcKind struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind"`
}

type ManagedServiceSpecMsvcKindIn struct {
	APIVersion string  `json:"apiVersion"`
	Kind       *string `json:"kind"`
}

type ManagedServiceSpecTolerations struct {
	Effect            *string `json:"effect"`
	Key               *string `json:"key"`
	Operator          *string `json:"operator"`
	TolerationSeconds *int    `json:"tolerationSeconds"`
	Value             *string `json:"value"`
}

type ManagedServiceSpecTolerationsIn struct {
	Effect            *string `json:"effect"`
	Key               *string `json:"key"`
	Operator          *string `json:"operator"`
	TolerationSeconds *int    `json:"tolerationSeconds"`
	Value             *string `json:"value"`
}

type Patch struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type PatchIn struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

type ProjectSpec struct {
	AccountName     string  `json:"accountName"`
	DisplayName     *string `json:"displayName"`
	Logo            *string `json:"logo"`
	TargetNamespace string  `json:"targetNamespace"`
}

type ProjectSpecIn struct {
	AccountName     string  `json:"accountName"`
	DisplayName     *string `json:"displayName"`
	Logo            *string `json:"logo"`
	TargetNamespace string  `json:"targetNamespace"`
}

type RouterSpec struct {
	Region          *string              `json:"region"`
	Routes          []*RouterSpecRoutes  `json:"routes"`
	BasicAuth       *RouterSpecBasicAuth `json:"basicAuth"`
	Cors            *RouterSpecCors      `json:"cors"`
	Domains         []*string            `json:"domains"`
	HTTPS           *RouterSpecHTTPS     `json:"https"`
	MaxBodySizeInMb *int                 `json:"maxBodySizeInMB"`
	RateLimit       *RouterSpecRateLimit `json:"rateLimit"`
}

type RouterSpecBasicAuth struct {
	Enabled    bool    `json:"enabled"`
	SecretName *string `json:"secretName"`
	Username   *string `json:"username"`
}

type RouterSpecBasicAuthIn struct {
	Enabled    bool    `json:"enabled"`
	SecretName *string `json:"secretName"`
	Username   *string `json:"username"`
}

type RouterSpecCors struct {
	AllowCredentials *bool     `json:"allowCredentials"`
	Enabled          *bool     `json:"enabled"`
	Origins          []*string `json:"origins"`
}

type RouterSpecCorsIn struct {
	AllowCredentials *bool     `json:"allowCredentials"`
	Enabled          *bool     `json:"enabled"`
	Origins          []*string `json:"origins"`
}

type RouterSpecHTTPS struct {
	ForceRedirect *bool `json:"forceRedirect"`
	Enabled       bool  `json:"enabled"`
}

type RouterSpecHTTPSIn struct {
	ForceRedirect *bool `json:"forceRedirect"`
	Enabled       bool  `json:"enabled"`
}

type RouterSpecIn struct {
	Region          *string                `json:"region"`
	Routes          []*RouterSpecRoutesIn  `json:"routes"`
	BasicAuth       *RouterSpecBasicAuthIn `json:"basicAuth"`
	Cors            *RouterSpecCorsIn      `json:"cors"`
	Domains         []*string              `json:"domains"`
	HTTPS           *RouterSpecHTTPSIn     `json:"https"`
	MaxBodySizeInMb *int                   `json:"maxBodySizeInMB"`
	RateLimit       *RouterSpecRateLimitIn `json:"rateLimit"`
}

type RouterSpecRateLimit struct {
	Connections *int  `json:"connections"`
	Enabled     *bool `json:"enabled"`
	Rpm         *int  `json:"rpm"`
	Rps         *int  `json:"rps"`
}

type RouterSpecRateLimitIn struct {
	Connections *int  `json:"connections"`
	Enabled     *bool `json:"enabled"`
	Rpm         *int  `json:"rpm"`
	Rps         *int  `json:"rps"`
}

type RouterSpecRoutes struct {
	App     *string `json:"app"`
	Lambda  *string `json:"lambda"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
	Rewrite *bool   `json:"rewrite"`
}

type RouterSpecRoutesIn struct {
	App     *string `json:"app"`
	Lambda  *string `json:"lambda"`
	Path    string  `json:"path"`
	Port    int     `json:"port"`
	Rewrite *bool   `json:"rewrite"`
}
