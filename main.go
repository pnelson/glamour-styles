package main

import (
	"flag"
	"log"
	"os"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
)

func main() {
	flag.Parse()
	for _, color := range []string{"#ff0000", "#00ff00", "#0000ff"} {
		err := render(color)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func render(color string) error {
	styles := glamour.WithStyles(ansi.StyleConfig{
		Code: ansi.StyleBlock{
			StylePrimitive: ansi.StylePrimitive{
				Color: stringPtr(color),
			},
		},
		CodeBlock: ansi.StyleCodeBlock{
			Chroma: &ansi.Chroma{
				Text: ansi.StylePrimitive{
					Color: stringPtr(color),
				},
			},
		},
	})
	renderer, err := glamour.NewTermRenderer(styles)
	if err != nil {
		return err
	}
	b, err := renderer.RenderBytes([]byte(color + "\n\n`code`\n\n```\ncode block\n```\n"))
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(append(b, '\n'))
	return err
}

func stringPtr(s string) *string { return &s }
