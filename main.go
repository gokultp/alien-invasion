package main

import (
	"flag"
	"os"

	"github.com/gokultp/aliegn-invasion/pkg"
)

func main() {
	input := flag.String("input", "./input.txt", "use --input=<file path> to specify map input file")
	output := flag.String("output", "./out.txt", "use --output=<file path> to specify map output file")
	aliens := flag.Int("aliens", 2, "--aliens=<value> to specify number of aliens")

	flag.Parse()
	inf, err := os.Open(*input)
	if err != nil {
		panic(err)
	}
	outf, err := os.Create(*output)
	if err != nil {
		panic(err)
	}
	iv := pkg.NewInvader(inf)
	iv.InitAliens(*aliens)
	iv.Exec()
	iv.PrintCities(outf)
}
