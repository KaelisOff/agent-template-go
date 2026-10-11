// Command mon-agent est un point de départ pour écrire un agent Kaelis :
// remplacez la fonction run par votre logique.
package main

import (
	"fmt"
	"strings"

	"github.com/KaelisOff/agent-sdk-go/agent"
)

func main() { agent.Main(run) }

// state est ce que l'agent garde d'un passage à l'autre (capacité storage).
type state struct {
	Runs int `json:"runs"`
}

// run fait le travail d'un passage et renvoie son résumé, affiché dans
// l'historique de l'agent.
func run(box *agent.Box, r agent.Run) (string, error) {
	var st state
	if _, err := box.GetJSON("state", &st); err != nil {
		return "", err
	}
	st.Runs++

	text, err := box.ChatMax(150,
		agent.System("Tu écris en français, en deux phrases maximum, sans titre."),
		agent.User(fmt.Sprintf("Écris un mot d'encouragement pour la journée du %s.", r.Now.Format("02/01/2006"))))
	if err != nil {
		return "", err
	}
	if err := box.Notify("Mon agent", strings.TrimSpace(text)); err != nil {
		return "", err
	}
	if err := box.SetJSON("state", st); err != nil {
		return "", err
	}
	box.Logf("passage n°%d (%s)", st.Runs, r.Trigger.Kind)
	return fmt.Sprintf("message n°%d envoyé", st.Runs), nil
}
