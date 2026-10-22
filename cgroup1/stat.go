/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package cgroup1

import (
	"errors"
	"fmt"

	v1 "github.com/containerd/cgroups/v3/cgroup1/stats"
)

// StatFiltered returns metrics for the selected active subsystems without
// reading files or resolving paths for other subsystems. An empty list selects
// all subsystems, as in Cgroup.Stat. Duplicate and inactive names are ignored.
// Select both Cpu and Cpuacct to collect CPU throttling and usage statistics.
// The CPU, CPU.Usage and CPU.Throttling fields are always initialized, as in Stat.
//
// This function keeps the Cgroup interface compatible with existing implementations.
// Cgroups returned by New, Load and Cgroup.New support filtered stats. Other
// implementations can opt in by providing this method:
//
//	StatFiltered([]Name, ...ErrorHandler) (*stats.Metrics, error)
//
// If c does not implement that method, StatFiltered returns an error wrapping
// errors.ErrUnsupported.
func StatFiltered(c Cgroup, subsystems []Name, handlers ...ErrorHandler) (*v1.Metrics, error) {
	if filtered, ok := c.(interface {
		StatFiltered([]Name, ...ErrorHandler) (*v1.Metrics, error)
	}); ok {
		return filtered.StatFiltered(subsystems, handlers...)
	}
	return nil, fmt.Errorf("cgroups: filtered stats: %w", errors.ErrUnsupported)
}
