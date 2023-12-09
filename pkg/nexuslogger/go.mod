module github.com/nebula-aac/kubernexus-mono/pkg/nexuslogger

replace github.com/nebula-aac/kubernexus-mono => ../..

go 1.21.0

require (
	github.com/charmbracelet/lipgloss v0.9.1
	github.com/nebula-aac/kubernexus-mono v0.0.0-00010101000000-000000000000
	github.com/sirupsen/logrus v1.9.3
	gorm.io/gorm v1.25.5
	xorm.io/core v0.7.3
)

require (
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/muesli/reflow v0.3.0 // indirect
	github.com/muesli/termenv v0.15.2 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	golang.org/x/sys v0.15.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
)
