package main

import (
	"strings"
	"testing"

	"github.com/KaelisOff/agent-sdk-go/agent"
	"github.com/KaelisOff/agent-sdk-go/agenttest"
)

func TestRun(t *testing.T) {
	box, err := agenttest.Load("manifest.toml")
	if err != nil {
		t.Fatal(err)
	}
	// Ce que répond le modèle pendant le test.
	box.Model = func(msgs []agent.Message, maxTokens int) (string, error) {
		return " Belle journée à vous ! ", nil
	}

	res := box.Run(run, agenttest.Scheduled("0 9 * * 1-5"))
	if res.Status != "ok" {
		t.Fatalf("%+v\n%s", res, res.Logs)
	}
	if len(box.Notifications) != 1 || box.Notifications[0].Body != "Belle journée à vous !" {
		t.Errorf("notifications = %+v", box.Notifications)
	}

	// La mémoire est gardée d'un passage à l'autre.
	res = box.Run(run, agenttest.Manual())
	if !strings.Contains(res.Summary, "n°2") {
		t.Errorf("summary = %q", res.Summary)
	}
}
