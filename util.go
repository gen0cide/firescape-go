package firescape

import "fmt"

type GoPlayer struct {
	Username string
}

func (g *GoPlayer) SetUsername(s string) {
	g.Username = s
}

func (g *GoPlayer) SayHello() {
	fmt.Printf("Hello %s!\n", g.Username)
}

func (g *GoPlayer) GetUsername() string {
	return g.Username
}
