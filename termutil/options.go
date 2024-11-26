package termutil

import "os"

type Option func(t *Terminal)

func WithLogFile(path string) Option {
	return func(t *Terminal) {
		t.logFile, _ = os.Create(path)
	}
}

func WithTheme(theme *Theme) Option {
	return func(t *Terminal) {
		t.theme = theme
	}
}

func WithCommand(command string, args ...string) Option {
	return func(t *Terminal) {
		t.command = command
		t.args = args
	}
}
