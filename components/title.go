package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type AppControl struct {
	app.Compo
	tabTitle string
}

func (uc *AppControl) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}

// implement Navigator in the component, so as to use SetTitle()
// https://github.com/maxence-charriere/go-app/issues/578
func (uc *AppControl) OnNav(ctx app.Context) {
	return
}

func (uc *AppControl) Render() app.UI {
	return app.Div().
		Body(
			app.If(uc.tabTitle == "",
				app.P().Body(
					app.Input().Size(30).
						ID("tabTitle").
						Value(uc.tabTitle).
						Placeholder("What is the title for the tab?").
						AutoFocus(true).
						Spellcheck(true).
						OnChange(uc.OnChange),
				),
			).Else(
				app.H1().Body(
					app.Text(uc.tabTitle),
				),
			))
}

func (uc *AppControl) OnChange(ctx app.Context, e app.Event) {
	uc.tabTitle = ctx.JSSrc().Get("value").String()
	ctx.Page().SetTitle(uc.tabTitle)
}
