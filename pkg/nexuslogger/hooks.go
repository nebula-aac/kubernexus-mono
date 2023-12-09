package nexuslogger

import "github.com/sirupsen/logrus"

type Hook interface {
	Levels() []logrus.Level
	Fire(*logrus.Entry) error
}

type LevelHooks map[logrus.Level][]Hook

func (hooks LevelHooks) Add(hook Hook) {
	for _, level := range hook.Levels() {
		hooks[level] = append(hooks[level], hook)
	}
}

func (hooks LevelHooks) Fire(level logrus.Level, entry *logrus.Entry) error {
	for _, hook := range hooks[level] {
		if err := hook.Fire(entry); err != nil {
			return err
		}
	}
	return nil
}
