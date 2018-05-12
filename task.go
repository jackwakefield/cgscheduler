package cgscheduler

import (
	"context"

	"golang.org/x/sync/errgroup"
)

// TaskFunc describes the signature of a task function.
type TaskFunc = func(ctx context.Context) error

// Task wraps a task function.
type Task struct {
	scheduler *Scheduler
	function  TaskFunc
}

func newTask(scheduler *Scheduler, function TaskFunc) *Task {
	return &Task{
		scheduler: scheduler,
		function:  function,
	}
}

// DependsOn creates a dependency between this task and the dependency task.
// When ran, the scheduler ensures the dependency task is executed first.
func (t *Task) DependsOn(dependency *Task) {
	t.scheduler.AddDependency(t, dependency)
}

// Run executes the task.
func (t *Task) Run(ctx context.Context) error {
	return t.function(ctx)
}

type taskRunner struct {
}

func newTaskRunner() *taskRunner {
	return &taskRunner{}
}

func (r *taskRunner) Run(ctx context.Context, tasks []*Task) error {
	g, ctx := errgroup.WithContext(ctx)

	for i := range tasks {
		task := tasks[i]

		g.Go(func() error {
			return task.Run(ctx)
		})
	}

	return g.Wait()
}
