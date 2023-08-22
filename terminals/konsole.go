package terminals

import (
	"github.com/srlehn/termimg/internal/consts"
	"github.com/srlehn/termimg/internal/environ"
	"github.com/srlehn/termimg/internal/propkeys"
	"github.com/srlehn/termimg/term"
)

////////////////////////////////////////////////////////////////////////////////
// Konsole
////////////////////////////////////////////////////////////////////////////////

func init() {
	term.RegisterTermChecker(&termCheckerKonsole{term.NewTermCheckerCore(termNameKonsole)})
}

const termNameKonsole = `konsole`

var _ term.TermChecker = (*termCheckerKonsole)(nil)

type termCheckerKonsole struct{ term.TermChecker }

func (t *termCheckerKonsole) CheckIsQuery(qu term.Querier, tty term.TTY, pr environ.Proprietor) (is bool, p environ.Proprietor) {
	p = environ.NewProprietor()
	if t == nil || pr == nil {
		p.SetProperty(propkeys.CheckTermQueryIsPrefix+termNameKonsole, consts.CheckTermFailed)
		return false, p
	}
	term.QueryDeviceAttributes(qu, tty, pr, pr)
	da3ID, _ := pr.Property(propkeys.DA3ID)
	var konsoleDA3ID = `~KDE` // hex encoded: `7E4B4445`
	if da3ID != konsoleDA3ID {
		p.SetProperty(propkeys.CheckTermQueryIsPrefix+termNameKonsole, consts.CheckTermFailed)
		return false, p
	}
	p.SetProperty(propkeys.CheckTermQueryIsPrefix+termNameKonsole, consts.CheckTermPassed)
	return true, p
}
