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

package repo

import (
	"github.com/pkg/errors"
	"goodrain.com/cloud-adaptor/internal/domain"
	"goodrain.com/cloud-adaptor/internal/model"
	"goodrain.com/cloud-adaptor/internal/repo/dao"
	"goodrain.com/cloud-adaptor/pkg/bcode"
	"goodrain.com/cloud-adaptor/pkg/util/mapper"
	"gorm.io/gorm"
)

// KubernetesTaskRepo -
type KubernetesTaskRepo interface {
	Transaction(tx *gorm.DB) KubernetesTaskRepo
	Create(kubernetesTask *domain.KubernetesTask) error
	UpdateStatus(eid string, taskID string, status string) error
	GetLastTask(eid string, providerName string) (*domain.KubernetesTask, error)
	GetTask(eid string, taskID string) (*domain.KubernetesTask, error)
}

// NewKubernetesTaskRepo creates a new KubernetesTaskRepo.
func NewKubernetesTaskRepo(kubernetesTaskDao dao.KubernetesTaskDao) KubernetesTaskRepo {
	return &kubernetesTaskRepo{
		kubernetesTaskDao: kubernetesTaskDao,
	}
}

type kubernetesTaskRepo struct {
	kubernetesTaskDao dao.KubernetesTaskDao
}

func (k *kubernetesTaskRepo) Transaction(tx *gorm.DB) KubernetesTaskRepo {
	kubernetesTaskDao := k.kubernetesTaskDao.Transaction(tx)
	return &kubernetesTaskRepo{
		kubernetesTaskDao: kubernetesTaskDao,
	}
}

func (k *kubernetesTaskRepo) UpdateStatus(eid string, taskID string, status string) error {
	return k.kubernetesTaskDao.UpdateStatus(eid, taskID, status)
}

func (k *kubernetesTaskRepo) GetLastTask(eid string, providerName string) (*domain.KubernetesTask, error) {
	task, err := k.kubernetesTaskDao.GetLastTask(eid, providerName)
	if err != nil {
		return nil, err
	}

	res := &domain.KubernetesTask{}
	mapper.Mapper(task, res)
	return res, nil
}

func (k *kubernetesTaskRepo) GetTask(eid string, taskID string) (*domain.KubernetesTask, error) {
	task, err := k.kubernetesTaskDao.GetTask(eid, taskID)
	if err != nil {
		return nil, err
	}

	res := &domain.KubernetesTask{}
	mapper.Mapper(task, res)
	return res, nil
}

func (k *kubernetesTaskRepo) Create(kubernetesTask *domain.KubernetesTask) error {
	lastOne, err := k.kubernetesTaskDao.GetLastOneByClusterID(kubernetesTask.ClusterID)
	if err != nil && !errors.Is(err, bcode.ErrKubernetesTaskNotFound) {
		return nil
	}
	if lastOne != nil {
		if lastOne.Status != "complete" {
			return errors.WithStack(bcode.ErrLastKubernetesTaskNotComplete)
		}
		kubernetesTask.Version = lastOne.Version + 1
	}

	task := &model.CreateKubernetesTask{}
	mapper.Mapper(kubernetesTask, task)
	return k.kubernetesTaskDao.Create(task)
}
