package cgscheduler

// Options holds the scheduler's configuration.
type Options struct {
	ConcurrentTasks int
}

// Option describes a function which mutates the scheduler's configuration.
type Option func(*Options)

// ConcurrentTasks sets the maximum number of tasks to run at any given time.
func ConcurrentTasks(maximum int) Option {
	return func(options *Options) {
		options.ConcurrentTasks = maximum
	}
}
