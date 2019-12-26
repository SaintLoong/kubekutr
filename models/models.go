package models

// Config represents the structure to hold configuration loaded from an external data source.
type Config struct {
	Deployments  []Deployment  `koanf:"deployments"`
	Services     []Service     `koanf:"services"`
	Ingresses    []Ingress     `koanf:"ingresses"`
	StatefulSets []StatefulSet `koanf:"statefulsets"`
}

// Deployment represents
type Deployment struct {
	Name       string      `koanf:"name"`
	Replicas   string      `koanf:"replicas"`
	Containers []Container `koanf:"containers"`
	Labels     []Label     `koanf:"labels"`
}

// Container represents
type Container struct {
	Name          string   `koanf:"name"`
	Image         string   `koanf:"image"`
	EnvSecret     string   `koanf:"envSecret"`
	EnvVars       []EnvVar `koanf:"envVars"`
	Container     string   `koanf:"container"`
	PortInt       string   `koanf:"port"`
	Command       string   `koanf:"command"`
	Args          string   `koanf:"arg"`
	ConfigMapName string   `koanf:"configmap"`
}

type Service struct {
	Name       string `koanf:"name"`
	Port       int    `koanf:"port"`
	TargetPort int    `koanf:"targetPort"`
	Type       string `koanf:"type"`
}

type Ingress struct {
	Name        string        `koanf:"name"`
	Class       string        `koanf:"class"`
	Paths       []IngressPath `koanf:"ingressPaths"`
	Annotations []Annotation  `koanf:"annotations"`
}

type StatefulSet struct {
	Name        string        `koanf:"name"`
	Class       string        `koanf:"class"`
	Paths       []IngressPath `koanf:"ingressPaths"`
	Annotations []Annotation  `koanf:"annotations"`
}

type IngressPath struct {
	Path    string `koanf:"path"`
	Service string `koanf:"service"`
	Port    string `koanf:"port"`
}

// Resource is a set of common actions performed on Resource Types
type Resource interface {
	GetMetaData() ResourceMeta
}

type ResourceMeta struct {
	Name         string
	Config       map[string]interface{}
	TemplatePath string
	ManifestPath string
}

type Annotation struct {
	Annotation string `koanf:"annotation"`
}

type Label struct {
	Label string `koanf:"label"`
}

type EnvVar struct {
	Name  string `koanf:"name"`
	Value string `koanf:"value"`
}
