// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || zos

package utils

import (
	"os/exec"
	"os/user"
	"strconv"
	"syscall"
)

func lookupUser(username string) (int, int, error) {
	u, err := user.Lookup(username)
	if err != nil {
		return 0, 0, err
	}

	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return 0, 0, err
	}

	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return 0, 0, err
	}

	return uid, gid, nil
}

// setProcessExecByUsername 当提供的用户名不正确时,不做任何操作
func setProcessExecByUsername(cmd *exec.Cmd, username string) error {
	uid, gid, err := lookupUser(username)
	if err != nil {
		return err
	}
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Credential: &syscall.Credential{
			Uid: uint32(uid),
			Gid: uint32(gid),
		},
	}
	return nil
}

func setProcessGroupID(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}

func processGroupKill(cmd *exec.Cmd) error {
	if cmd.Process.Pid != 0 {
		err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		if err != nil {
			return err
		}
	}
	return nil
}
