package config

type Config struct {
	Host                string   `json:"host"`
	AllocatableResource Resource `json:"allocatableResource"`
	RequestedResource   Resource `json:"requestedResource"`
	Stats               Stats    `json:"stats"`
}

type Resource struct {
	MilliCPU int64 `json:"milliCPU"`
	Memory   int64 `json:"memory"`
	Storage  int64 `json:"storage"`
}

type Stats struct {
	CPU     Readable `json:"cpu"`
	OS      string   `json:"os"`
	Memory  Readable `json:"memory"`
	Storage Readable `json:"storage"`
}

type Readable struct {
	Total string `json:"total"`
	Used  string `json:"used"`
}

var (
	Build   string
	Version string
)

func New() *Config {
	return &Config{}
}
