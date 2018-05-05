package mdopen

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"

	"github.com/pkg/errors"
	"github.com/romanyx/mdopen/internal/templates/github"
	"github.com/tink-ab/tempfile"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// Option for initializer.
type Option func(*Opener)

// GithubTemplate option sets layout as
// github.com template.
func GithubTemplate() Option {
	return func(opnr *Opener) {
		opnr.layout = template.Must(template.New("layout").Parse(github.Template))
	}
}

// Opener holds layout and command name
// to open default browser. Use New function
// to initialize corrent one.
type Opener struct {
	cmdName string
	layout  *template.Template
}

// New returns initialized Opener.
func New(options ...Option) *Opener {
	opnr := Opener{
		cmdName: cmdName(),
		layout:  template.Must(template.New("layout").Parse(github.Template)),
	}

	for _, option := range options {
		option(&opnr)
	}

	return &opnr
}

// Open will create a tmp file, execute layout
// template with given markdown into it and then
// open it in browser.
func (opnr *Opener) Open(f io.Reader) error {
	tmpfile, err := tmpFile()
	if err != nil {
		return errors.Wrap(err, "tempfile creation failed")
	}
	defer tmpfile.Close()

	if err := opnr.prepareFile(tmpfile, f); err != nil {
		return errors.Wrap(err, "tmp file perpare")
	}

	url := fmt.Sprintf("file:///%s", tmpfile.Name())
	cmd := exec.Command(opnr.cmdName, url)
	if err := cmd.Run(); err != nil {
		return errors.Wrap(err, "open letter in the browser failed")
	}

	return nil
}

func (opnr *Opener) prepareFile(w io.Writer, f io.Reader) error {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "read file failed")
	}

	templateData := struct {
		Body template.HTML
	}{
		Body: template.HTML(blackfriday.Run(data)),
	}

	if err := opnr.layout.Execute(w, templateData); err != nil {
		return errors.Wrap(err, "template execution failed")
	}

	return nil
}

func tmpFile() (*os.File, error) {
	tmpfile, err := tempfile.TempFile("", "mdopen", ".html")
	if err != nil {
		return tmpfile, err
	}

	return tmpfile, nil
}

func cmdName() string {
	switch runtime.GOOS {
	case "darwin":
		return "open"
	case "windows":
		return "cmd /c start"
	default:
		return "xdg-open"
	}
}
