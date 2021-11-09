package stories

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/pojntfx/keygaen/pkg/components"
)

type NavbarStory struct {
	Story
}

func (c *NavbarStory) Render() app.UI {
	return c.WithRoot(
		&components.Navbar{},
	)
}
