package cmd

type EdgeadmConfig struct {
	IsEnableEdge        bool
	WorkerPath          string
	ManifestsDir        string
	InstallPkgPath      string
	Kubeconfig          string
	TunnelCloudToken    string
	KubeVIPInterface    string
	DefaultHA           string
	ContainerRuntime    string
	Version             string
	PodInfraContainer   string
	EdgeImageRepository string
	EdgeVirtualAddr     string
}
