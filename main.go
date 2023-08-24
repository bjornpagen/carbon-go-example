package main

import (
	"os"

	"github.com/bjornpagen/carbon-go"
)

const styles = `
body {
	max-width: 672px;
	margin: 0 auto;
	padding-bottom: 8rem;
}
main {
	display: flex;
	flex-direction: column;
	gap: 3rem;
}
section {
	display: flex;
	flex-direction: column;
	gap: 1.5rem;
}
`

func main() {
	head := []any{
		`<meta charset="utf-8">`,
		`<meta name="viewport" content="width=device-width, initial-scale=1">`,
		`<title>Carbon Design Go</title>`,
		carbon.Style().Custom(styles),
	}
	body := []any{
		carbon.Header().Company("Carbon").PlatformName("Go"),
		carbon.Content(
			carbon.Section(
				`<h1>carbon-go</h1>`,
				`<p>A conformant implementation of IBM's Carbon Design, built for HTMX/SSR first.</p>`,
				carbon.Div(
					carbon.Checkbox().LabelText("Store all application state server side").Checked(true),
					carbon.Checkbox().LabelText("Write Go instead of JavaScript (noscript works)").Checked(true),
					carbon.Checkbox().LabelText("Fully CSP friendly: no inline JS or inline styles").Checked(true),
				),
				carbon.ButtonSet(
					carbon.Button("Get started").Href("https://github.com/bjornpagen/carbon-go").Icon(carbon.Checkmark()),
					carbon.Button("See Demo").Kind("secondary").Href("https://github.com/bjornpagen/carbon-go-example").Icon(carbon.GitHub()),
				),
			).Attr("id", "hero"),

			carbon.Section(
				`<h3>Why does this exist?</h3>`,
				`<p>Basically, it's React Server components but not shit.</p>`,
				`<p>If you wanted to write an entire web app from scratch in Go, you had to write every single component by yourself.</p>`,
				`<p><strong>Solution—</strong>why not take an existing design system that already looks good, and just port components from the React codebase, into Go?`,
			).Attr("id", "rationale"),

			carbon.Section(
				`<h3>FAQ</h3>`,
				carbon.Accordion(
					carbon.AccordionItem().Title("<strong>Why would you want to use this instead of React?</strong>").Content("<p>Because I prefer Go and hypermedia as the engine of application state.</p>").Open(true),
					carbon.AccordionItem().Title("<strong>Why do you bundle 600KB of CSS?</strong>").Content("<p>Currently we bundle all of Carbon Design's upstream styles directly, but we really should be smarter about generating this from only the components we use. This is on the TODO list.</p>").Open(true),
					carbon.AccordionItem().Title("<strong>This is wrong! The markup looks ugly!</strong>").Content("<p>I think it looks great. Go back to your JavaScript frameworks.</p>").Open(true),
				),
			).Attr("id", "faq"),

			carbon.Section(
				`<h3>What's next?</h3>`,
				carbon.UnorderedList(
					carbon.ListItem("Add the rest of the components"),
					carbon.ListItem("Implement the optional client side JS for accordion"),
					carbon.ListItem("Comb through the API and make it stable/idiomatic"),
					carbon.ListItem("Add comprehensise tests and documentation"),
					carbon.ListItem("Add a smart process to generate the CSS from only the components we use"),
				),
			),
		),
	}

	os.Mkdir("out", 0755)
	f, _ := os.Create("out/index.html")
	carbon.Html().Head(head).Body(body).Render(f)
	f.Sync()
	f.Close()
}
