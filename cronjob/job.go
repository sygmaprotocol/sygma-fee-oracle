// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package cronjob

import "github.com/sirupsen/logrus"

type ScheduledJob interface {
	IsEnabled() bool
	JobName() string
	GetJobSchedule() string
	SetOperationFunc(func())
	GetOperationFunc() func()
}

var _ ScheduledJob = (*Job)(nil)

type Job struct {
	cronBase  *Cronjob
	name      string
	schedule  string
	enable    bool
	operation func()

	log *logrus.Entry
}

func NewJob(c *Cronjob, name, schedule string, enable bool) *Job {
	return &Job{
		cronBase: c,
		name:     name,
		schedule: schedule,
		enable:   enable,
		log:      c.log.WithField("job", name),
	}
}

func (c *Job) IsEnabled() bool {
	return c.enable
}

func (c *Job) JobName() string {
	return c.name
}

func (c *Job) GetJobSchedule() string {
	return c.schedule
}

func (c *Job) SetOperationFunc(operationFn func()) {
	c.operation = operationFn
}

func (c *Job) GetOperationFunc() func() {
	return c.operation
}
