// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kube

import (
	"context"

	"yunion.io/x/jsonutils"
	"yunion.io/x/pkg/errors"

	"yunion.io/x/onecloud/pkg/apis"
	"yunion.io/x/onecloud/pkg/cloudcommon/db"
	"yunion.io/x/onecloud/pkg/cloudcommon/db/taskman"
	"yunion.io/x/onecloud/pkg/compute/models"
	"yunion.io/x/onecloud/pkg/util/logclient"
)

type KubeClusterSyncstatusTask struct {
	taskman.STask
}

func init() {
	taskman.RegisterTask(KubeClusterSyncstatusTask{})
}

func (self *KubeClusterSyncstatusTask) taskFailed(ctx context.Context, cluster *models.SKubeCluster, err error) {
	cluster.SetStatus(ctx, self.GetUserCred(), apis.STATUS_UNKNOWN, err.Error())
	db.OpsLog.LogEvent(cluster, db.ACT_SYNC_STATUS, cluster.GetShortDesc(ctx), self.GetUserCred())
	logclient.AddActionLogWithContext(ctx, cluster, logclient.ACT_SYNC_STATUS, err, self.UserCred, false)
	self.SetStageFailed(ctx, jsonutils.NewString(err.Error()))
}

func (self *KubeClusterSyncstatusTask) OnInit(ctx context.Context, obj db.IStandaloneModel, data jsonutils.JSONObject) {
	cluster := obj.(*models.SKubeCluster)

	iCluster, err := cluster.GetIKubeCluster(ctx)
	if err != nil {
		self.taskFailed(ctx, cluster, errors.Wrapf(err, "GetIKubeCluster"))
		return
	}

	err = cluster.SyncAllWithCloudKubeCluster(ctx, self.GetUserCred(), iCluster, nil)
	if err != nil {
		self.taskFailed(ctx, cluster, errors.Wrapf(err, "SyncAllWithCloudKubeCluster"))
		return
	}
	self.SetStageComplete(ctx, nil)
}
