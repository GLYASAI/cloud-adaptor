// RAINBOND, Application Management Platform
// Copyright (C) 2020-2020 Goodrain Co., Ltd.

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

package dao

import (
	"github.com/pkg/errors"
	"goodrain.com/cloud-adaptor/internal/model"
	"goodrain.com/cloud-adaptor/pkg/bcode"
	"goodrain.com/cloud-adaptor/pkg/util/uuidutil"
	"gorm.io/gorm"
)

// KubernetesTaskDao enterprise create kubernetes task
type KubernetesTaskDao interface {
	Transaction(tx *gorm.DB) KubernetesTaskDao
	Create(ent *model.CreateKubernetesTask) error
	GetLastTask(eid string, providerName string) (*model.CreateKubernetesTask, error)
	UpdateStatus(eid string, taskID string, status string) error
	GetTask(eid string, taskID string) (*model.CreateKubernetesTask, error)
	GetLastOneByClusterID(clusterID string) (*model.CreateKubernetesTask, error)
}

// CreateKubernetesTaskDao enterprise create kubernetes task
type CreateKubernetesTaskDao struct {
	DB *gorm.DB `inject:""`
}

// NewKubernetesTaskRepo creates a new KubernetesTaskDao.
func NewKubernetesTaskRepo(db *gorm.DB) KubernetesTaskDao {
	return &CreateKubernetesTaskDao{DB: db}
}

// Transaction -
func (c *CreateKubernetesTaskDao) Transaction(tx *gorm.DB) KubernetesTaskDao {
	return &CreateKubernetesTaskDao{DB: tx}
}

//Create create a task
func (c *CreateKubernetesTaskDao) Create(ck *model.CreateKubernetesTask) error {
	var old model.CreateKubernetesTask
	if ck.TaskID == "" {
		ck.TaskID = uuidutil.NewUUID()
	}
	if err := c.DB.Where("eid = ? and cluster_id=? and task_id=?", ck.EnterpriseID, ck.ClusterID, ck.TaskID).Take(&old).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
		// not found error, create new
		if err := c.DB.Save(ck).Error; err != nil {
			if isDuplicateEntry(err) {
				return errors.WithStack(bcode.ErrDuplicateKubernetesUpdateTask)
			}
			return err
		}
		return nil
	}
	return errors.WithStack(bcode.ErrKubernetesTaskExists)
}

// GetLastOneByClusterID -
func (c *CreateKubernetesTaskDao) GetLastOneByClusterID(clusterID string) (*model.CreateKubernetesTask, error) {
	var task model.CreateKubernetesTask
	if err := c.DB.Where("cluster_id = ?", clusterID).Last(&task).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.WithStack(bcode.ErrKubernetesTaskNotFound)
		}
		return nil, err
	}
	return &task, nil
}

//GetLastTask get last task
func (c *CreateKubernetesTaskDao) GetLastTask(eid string, providerName string) (*model.CreateKubernetesTask, error) {
	var old model.CreateKubernetesTask
	if err := c.DB.Where("eid = ? and provider_name=?", eid, providerName).Order("created_at desc").Limit(1).Take(&old).Error; err != nil {
		return nil, err
	}
	return &old, nil
}

//UpdateStatus update status
func (c *CreateKubernetesTaskDao) UpdateStatus(eid string, taskID string, status string) error {
	var old model.CreateKubernetesTask
	if err := c.DB.Model(&old).Where("eid = ? and task_id=?", eid, taskID).Update("status", status).Error; err != nil {
		return err
	}
	return nil
}

//GetTask get task
func (c *CreateKubernetesTaskDao) GetTask(eid string, taskID string) (*model.CreateKubernetesTask, error) {
	var old model.CreateKubernetesTask
	if err := c.DB.Where("eid = ? and task_id=?", eid, taskID).Take(&old).Error; err != nil {
		return nil, err
	}
	return &old, nil
}
