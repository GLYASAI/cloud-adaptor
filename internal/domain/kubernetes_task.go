// RAINBOND, Application Management Platform
// Copyright (C) 2020-2021 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package domain

// KubernetesTask -
type KubernetesTask struct {
	EnterpriseID       string `json:"eid"`
	ClusterID          string `json:"clusterID"`
	Version            int    `json:"version"`
	TaskID             string `json:"taskID"`
	Name               string `json:"name"`
	WorkerResourceType string `json:"resourceType"`
	WorkerNum          int    `json:"workerNum"`
	Provider           string `json:"provider"`
	Region             string `json:"region"`
	Status             string `json:"status"`
}
