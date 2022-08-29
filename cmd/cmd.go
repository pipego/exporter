package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/pipego/exporter/config"
)

const (
	// Duration Duration: 10s = 10*1000ms = 10*1000000000ns
	Duration = 10 * 1000000000
	Milli    = 1000
)

const (
	Dev  = "/dev/"
	Home = "/home"
	Root = "/"
)

var (
	app = kingpin.New("exporter", "pipego exporter").Version(config.Version + "-build-" + config.Build)
)

func Run(ctx context.Context) error {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	c := config.New()

	c.Host = host()
	c.AllocatableResource.MilliCPU, c.RequestedResource.MilliCPU = milliCPU()
	c.AllocatableResource.Memory, c.RequestedResource.Memory = memory()
	c.AllocatableResource.Storage, c.RequestedResource.Storage = storage()

	export(c)

	return nil
}

func host() string {
	conn, _ := net.Dial("udp", "8.8.8.8:8")
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	buf := conn.LocalAddr().(*net.UDPAddr)

	return strings.Split(buf.String(), ":")[0]
}

func milliCPU() (alloc, request int64) {
	c, err := cpu.Counts(true)
	if err != nil {
		return -1, -1
	}

	if c*Milli > math.MaxInt64 {
		return -1, -1
	}

	// FIXME: Got error on MacOS 10.13.6
	p, err := cpu.Percent(Duration, false)
	if err != nil {
		return -1, -1
	}

	used := float64(c) * p[0] * 0.01
	if used > math.MaxInt64 {
		return -1, -1
	}

	return int64(c * Milli), int64(used * Milli)
}

func memory() (alloc, request int64) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return -1, -1
	}

	if v.Total > math.MaxInt64 || v.Used > math.MaxInt64 {
		return -1, -1
	}

	return int64(v.Total), int64(v.Used)
}

func storage() (alloc, request int64) {
	helper := func(path string) bool {
		found := false
		p, _ := disk.Partitions(false)
		for _, item := range p {
			if strings.HasPrefix(item.Device, Dev) && item.Mountpoint == path {
				found = true
				break
			}
		}
		return found
	}

	r, err := disk.Usage(Root)
	if err != nil {
		return -1, -1
	}

	total := r.Total
	used := r.Used

	if helper(Home) {
		h, err := disk.Usage(Home)
		if err != nil {
			return -1, -1
		}
		total = h.Total
		used = h.Used
	}

	if total > math.MaxInt64 || used > math.MaxInt64 {
		return -1, -1
	}

	return int64(total), int64(used)
}

func export(cfg *config.Config) {
	b, _ := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(string(b))
}
