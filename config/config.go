package config

type Config struct {
	Host                string   `json:"host"`
	AllocatableResource Resource `json:"allocatableResource"`
	RequestedResource   Resource `json:"requestedResource"`
}

type Resource struct {
	MilliCPU int64 `json:"milliCPU"`
	Memory   int64 `json:"memory"`
	Storage  int64 `json:"storage"`
}

var (
	Build   string
	Version string
)

func New() *Config {
	return &Config{}
}
