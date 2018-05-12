package main

import (
	"context"
	"log"
	"time"

	"github.com/jackwakefield/cgscheduler"
)

func main() {
	scheduler := cgscheduler.New()

	// create some example tasks
	a := scheduler.AddTask(exampleTask("A"))
	b := scheduler.AddTask(exampleTask("B"))
	c := scheduler.AddTask(exampleTask("C"))
	d := scheduler.AddTask(exampleTask("D"))
	e := scheduler.AddTask(exampleTask("E"))
	f := scheduler.AddTask(exampleTask("F"))
	g := scheduler.AddTask(exampleTask("G"))
	h := scheduler.AddTask(exampleTask("H"))

	// register dependencies between the tasks
	h.DependsOn(g)
	h.DependsOn(f)
	g.DependsOn(e)
	f.DependsOn(e)
	e.DependsOn(d)
	d.DependsOn(c)
	c.DependsOn(b)
	c.DependsOn(a)

	// run the tasks
	if err := scheduler.Run(context.Background()); err != nil {
		log.Fatalln(err)
	}

	log.Println("Done")
}

func exampleTask(label string) cgscheduler.TaskFunc {
	return func(ctx context.Context) error {
		log.Println("Running task", label)
		time.Sleep(1 * time.Second)

		return nil
	}
}
