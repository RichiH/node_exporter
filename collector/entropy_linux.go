// Copyright 2015 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !noentropy

package collector

import (
//	"io/ioutil"
//	"strconv"
//	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type entropyCollector struct {
	entropy_avail *prometheus.Desc
//	poolsize      *prometheus.Desc
}

func init() {
	Factories["entropy"] = NewEntropyCollector
}

// Takes a prometheus registry and returns a new Collector exposing
// entropy stats
func NewEntropyCollector() (Collector, error) {
	return &entropyCollector{
		entropy_avail: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "entropy_avail"),
			"Bits of available entropy.",
			nil, nil,
		),
//		limit: prometheus.NewDesc(
//			prometheus.BuildFQName(Namespace, "", "nf_entropy_entries_limit"),
//			"Maximum size of connection tracking table.",
//			nil, nil,
//		),
	}, nil
}

func (c *entropyCollector) Update(ch chan<- prometheus.Metric) (err error) {
	value, err := readUintFromFile(procFilePath("sys/kernel/random/entropy_avail"))
	if err != nil {
		return nil
	}
	ch <- prometheus.MustNewConstMetric(
		c.entropy_avail, prometheus.GaugeValue, float64(value))

//	value, err = readUintFromFile(procFilePath("sys/net/netfilter/nf_entropy_max"))
//	if err != nil {
//		return nil
//	}
//	ch <- prometheus.MustNewConstMetric(
//		c.limit, prometheus.GaugeValue, float64(value))

	return nil
}

//func readUintFromFile(path string) (uint64, error) {
//	data, err := ioutil.ReadFile(path)
//	if err != nil {
//		return 0, err
//	}
//	value, err := strconv.ParseUint(strings.TrimSpace(string(data)), 10, 64)
//	if err != nil {
//		return 0, err
//	}
//	return value, nil
//}
