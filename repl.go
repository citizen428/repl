package repl

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"

	"github.com/peterh/liner"
)

var (
	cmd      string
	debug    bool
	compDir  string
	histDir  string
	histFile string
)

func Run() {
	parseFlags()
	if flag.NArg() == 0 {
		printUsage()
	}

	readLine := liner.NewLiner()
	defer readLine.Close()

	readLine.SetCtrlCAborts(true)
	loadCompletions(readLine)
	loadHistory(readLine)

	for {
		input, err := readLine.Prompt(cmd + ">> ")
		if err == io.EOF {
			saveHistory(readLine)
			os.Exit(0)
		}

		if input == "exit" || input == "quit" {
			fmt.Println("Use Ctrl-D (i.e. EOF) to exit")
			continue
		}

		if debug {
			fmt.Printf("EXECUTING: %s %s\n\n", cmd, input)
		}

		args := regexp.MustCompile(`\s+`).Split(input, -1)
		if cmdOut, err := exec.Command(cmd, args...).Output(); err == nil {
			fmt.Println(string(cmdOut))
			readLine.AppendHistory(input)
		}
	}
}

func parseFlags() {
	homeDir := os.Getenv(("HOME"))
	flag.BoolVar(&debug, "debug", false, "Enable debug output")
	flag.StringVar(&compDir, "compdir", homeDir+"/.repl", "Directory for completion files")
	flag.StringVar(&histDir, "histdir", homeDir, "Directory for history file")
	flag.Parse()

	cmd = flag.Arg(0)
	histFile = path.Join(histDir, ".repl_history")
}

func printUsage() {
	prog := path.Base(cmd)
	fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n  %s cmd [options]\n\nOptions:\n", prog)
	flag.PrintDefaults()
	os.Exit(0)
}

func loadCompletions(line *liner.State) {
	compFile := path.Join(compDir, cmd)
	if f, err := os.Open(compFile); err == nil {
		defer f.Close()

		buf := new(bytes.Buffer)
		buf.ReadFrom(f)
		comps := buf.String()

		line.SetCompleter(func(line string) (c []string) {
			for _, comp := range strings.Split(comps, " ") {
				if strings.HasPrefix(comp, strings.ToLower(line)) {
					c = append(c, comp)
				}
			}
			return
		})
	}
}

func loadHistory(line *liner.State) {
	if f, err := os.Open(histFile); err == nil {
		line.ReadHistory(f)
		f.Close()
	}
}

func saveHistory(line *liner.State) {
	if f, err := os.Create(histFile); err != nil {
		fmt.Println("Error writing history file: ", err)
	} else {
		line.WriteHistory(f)
		f.Close()
	}
}
