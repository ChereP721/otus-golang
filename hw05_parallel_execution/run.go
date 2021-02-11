package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorNegativeZeroN = errors.New("count goroutines is negative or zero")

var wg = sync.WaitGroup{}

type Task func() error

func generator(tasks []Task, n int, doneCh <-chan struct{}) <-chan Task {
	taskStream := make(chan Task, n)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(taskStream)

		for _, task := range tasks {
			select {
			case <-doneCh:
				return
			default:
			}

			select {
			case <-doneCh:
				return
			case taskStream <- task:
			}
		}
	}()

	return taskStream
}

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks.
func Run(tasks []Task, n int, m int) error {
	if len(tasks) < n {
		n = len(tasks)
	}

	if n <= 0 {
		return ErrErrorNegativeZeroN
	}

	doneCh := make(chan struct{})
	taskStream := generator(tasks, n, doneCh)
	errCounter := newErrCounter(m, doneCh)

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-doneCh:
					return
				default:
				}

				select {
				case <-doneCh:
					return
				case task, ok := <-taskStream:
					if !ok {
						return
					}

					if errCounter.check(task()) {
						return
					}
				}
			}
		}()
	}

	wg.Wait()

	return errCounter.err
}
