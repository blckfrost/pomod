package waybar

import (
	"encoding/json"
	"os"

	"github.com/blckfrost/pomod.git/internal/core"
)

type Output struct {
	Text    string `json:"text"`
	Class   string `json:"class"`
	Tooltip string `json:"tooltip"`
}

func Print(p *core.Pomodoro) {
	out := Output{
		Text:    "🍅 00:00",
		Class:   "idle",
		Tooltip: "Pomodoro",
	}

	if p.State != core.Idle {
		out.Text = "🍅 " + format(p.Remaining())
		out.Class = string(p.State)
	}

	json.NewEncoder(os.Stdout).Encode(out)
}
