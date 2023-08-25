package main

import (
	"os"

	x "github.com/bjornpagen/carbon-go"
)

const myStyle = `
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
* .bx--link {
	font-size: inherit;
}
`

func main() {
	head := []any{
		`<meta charset="utf-8">`,
		`<meta name="viewport" content="width=device-width, initial-scale=1">`,
		`<title>Carbon Design Go</title>`,
		x.Style(x.BaseCss),
		x.Style(x.FontFamilyCss),
		x.Style(myStyle),
	}
	body := []any{
		x.Header().Company("Carbon").PlatformName("Go"),
		x.Content(
			x.Section(
				`<h1>carbon-go</h1>`,
				x.P(
					`A conformant implementation of IBM's `,
					x.Link("Carbon Design").Href("https://carbondesignsystem.com/").Inline(true),
					`, built for HTMX/SSR first.`,
				),
				x.Div(
					x.Checkbox().LabelText("Store all application state server side").Checked(true),
					x.Checkbox().LabelText("Write Go instead of JavaScript (noscript works)").Checked(true),
					x.Checkbox().LabelText("Fully CSP friendly: no inline JS or inline styles").Checked(true),
				),
				x.ButtonSet(
					x.Button("Get started").Href("https://github.com/bjornpagen/carbon-go").Icon(x.Checkmark()),
					x.Button("See Demo").Kind("secondary").Href("https://github.com/bjornpagen/carbon-go-example").Icon(x.GitHub()),
				),
			).Attr("id", "hero"),

			x.Section(
				`<h3>Why does this exist?</h3>`,
				`<p>Basically, it's React Server components but not shit.</p>`,
				`<p>If you wanted to write an entire web app from scratch in Go, you had to write every single component by yourself.</p>`,
				`<p><strong>Solution—</strong>why not take an existing design system that already looks good, and just port components from the React codebase, into Go?`,
			).Attr("id", "rationale"),

			x.Section(
				`<h3>FAQ</h3>`,
				x.Accordion(
					x.AccordionItem().Title("<strong>Why would you want to use this instead of React?</strong>").Content("<p>Because hypermedia is the engine of application state.</p>").Open(true),
					x.AccordionItem().Title("<strong>Why do you bundle 600KB of CSS?</strong>").Content("<p>Currently we bundle all of Carbon Design's upstream styles directly, but we really should be smarter about generating this from only the components we use. This is on the TODO list.</p>").Open(true),
					x.AccordionItem().Title("<strong>This is wrong! The markup looks ugly!</strong>").Content("<p>I think it looks great. Go back to your JavaScript frameworks.</p>").Open(true),
				),
			).Attr("id", "faq"),

			x.Section(
				`<h3>What's next?</h3>`,
				x.UnorderedList(
					x.ListItem("Add the rest of the components"),
					x.ListItem("Implement the optional client side JS for accordion"),
					x.ListItem("Comb through the API and make it stable/idiomatic"),
					x.ListItem("Add comprehensise tests and documentation"),
					x.ListItem("Add a smart process to generate the CSS from only the components we use"),
				),
			),
		),
	}

	os.Mkdir("out", 0755)
	f, _ := os.Create("out/index.html")
	x.Html().Head(head).Body(body).Render(f)
	f.Sync()
	f.Close()
}
