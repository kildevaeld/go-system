package system

import (
	"os"
	"os/signal"
	"syscall"
)

func Run(fn func(kill <-chan struct{}) error) error {

	sigs := make(chan os.Signal, 1)
	done := make(chan error, 1)
	kill := make(chan struct{}, 1)

	defer close(sigs)
	defer close(done)
	defer close(kill)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer signal.Stop(sigs)

	go func() {
		done <- fn(kill)
	}()

	select {
	case err := <-done:
		return err
	case <-sigs:
		kill <- struct{}{}
		break
	}
	return nil


}
