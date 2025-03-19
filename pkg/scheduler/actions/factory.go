// Copyright 2025 NVIDIA CORPORATION
// SPDX-License-Identifier: Apache-2.0

package actions

import (
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/actions/allocate"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/actions/consolidation"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/actions/preempt"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/actions/reclaim"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/actions/stalegangeviction"
	"github.com/NVIDIA/KAI-scheduler/pkg/scheduler/framework"
)

func InitDefaultActions() {
	framework.RegisterAction(reclaim.New())
	framework.RegisterAction(allocate.New())
	framework.RegisterAction(preempt.New())
	framework.RegisterAction(consolidation.New())
	framework.RegisterAction(stalegangeviction.New())
}
