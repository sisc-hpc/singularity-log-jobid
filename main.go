// Copyright (c) 2020, Sylabs Inc. All rights reserved.
// Copyright (c) 2020, Ward Poelmans (Vrije Universiteit Brussel).
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package main

import (
	"fmt"
	"log/syslog"
	"path/filepath"
	"os/user"
        "io/ioutil"
	"regexp"

	pluginapi "github.com/sylabs/singularity/pkg/plugin"
	singularitycallback "github.com/sylabs/singularity/pkg/plugin/callback/runtime/engine/singularity"
	"github.com/sylabs/singularity/pkg/runtime/engine/config"
	singularityConfig "github.com/sylabs/singularity/pkg/runtime/engine/singularity/config"
)

// Plugin is the only variable which a plugin MUST export.
// This symbol is accessed by the plugin framework to initialize the plugin
var Plugin = pluginapi.Plugin{
	Manifest: pluginapi.Manifest{
                Name:        "github.com/wpoely86/singularity-log-jobid",
		Author:      "Ward Poelmans <wpoely86@gmail.com>",
		Version:     "0.1.0",
		Description: "Log executed commands to syslog along with username and jobid (both slurm and torque are supported)",
	},
	Callbacks: []pluginapi.Callback{
		(singularitycallback.PostStartProcess)(logCommand),
	},
}

func logCommand(common *config.Common, pid int) error {
	cfg := common.EngineConfig.(*singularityConfig.EngineConfig)

	command := "unknown"
	if cfg.OciConfig != nil && cfg.OciConfig.Process != nil {
		if len(cfg.OciConfig.Process.Args) > 0 {
			command = filepath.Base(cfg.OciConfig.Process.Args[0])
		}
	}

	image := cfg.GetImage()
	w, err := syslog.New(syslog.LOG_INFO, "singularity")
	if err != nil {
		return err
	}
	defer w.Close()

        user, err := user.Current()
        if err != nil {
           return err
        }

        // to figure out the jobid (if any), we look at the cgroups the current process is part of.
        cgroups, err := ioutil.ReadFile("/proc/self/cgroup")
        if err != nil {
           return err
        }

        // this regex should work for both slurm and torque
	r := regexp.MustCompile(`[0-9]+:cpuacct,cpu:/(?:slurm/uid_[0-9]+/job_(?P<slurm>[0-9_]+)/|torque/(?P<torque>[0-9]+.*))`)
        res := r.FindStringSubmatch(string(cgroups))

        var jobid string = ""
        if len(res) == 3 {
            if res[1] != "" {
                jobid = res[1]
            } else {
                jobid = res[2]
            }
        }

	msg := fmt.Sprintf("USER=%s IMAGE=%s JOBID=%s COMMAND=%s", user.Username, image, jobid, command)
	return w.Info(msg)
}
