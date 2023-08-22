package urxvt

import (
	"fmt"
	"image"
	"path/filepath"
	"strings"

	"github.com/srlehn/termimg/internal/consts"
	"github.com/srlehn/termimg/internal/encoder/encpng"
	"github.com/srlehn/termimg/internal/errors"
	"github.com/srlehn/termimg/mux"
	"github.com/srlehn/termimg/term"
)

func init() { term.RegisterDrawer(&drawerURXVT{}) }

var _ term.Drawer = (*drawerURXVT)(nil)

type drawerURXVT struct{}

func (d *drawerURXVT) Name() string     { return `urxvt` }
func (d *drawerURXVT) New() term.Drawer { return &drawerURXVT{} }

func (d *drawerURXVT) IsApplicable(inp term.DrawerCheckerInput) bool {
	return inp != nil && inp.Name() == `urxvt`
}

// TODO write ' ' over area (image is in a layer below text)
// TODO replace urxvt graphic with persistent block graphic when cleared

func (d *drawerURXVT) Draw(img image.Image, bounds image.Rectangle, tm *term.Terminal) error {
	if d == nil {
		return errors.New(consts.ErrNilReceiver)
	}
	if tm == nil || img == nil {
		return errors.New(consts.ErrNilParam)
	}
	timg, ok := img.(*term.Image)
	if !ok {
		timg = term.NewImage(img)
	}
	if timg == nil {
		return errors.New(consts.ErrNilImage)
	}

	rsz := tm.Resizer()
	if rsz == nil {
		return errors.New(`nil resizer`)
	}
	err := timg.Fit(bounds, rsz, tm)
	if err != nil {
		return err
	}

	urxvtString, err := d.getInbandString(timg, bounds, tm)
	if err != nil {
		return err
	}
	tm.Write([]byte(urxvtString))

	return nil
}

func (d *drawerURXVT) getInbandString(timg *term.Image, bounds image.Rectangle, tm *term.Terminal) (string, error) {
	if timg == nil {
		return ``, errors.New(consts.ErrNilImage)
	}
	urxvtString, err := timg.GetInband(bounds, d, tm)
	if err == nil {
		return urxvtString, nil
	}
	if timg.Cropped == nil {
		return ``, errors.New(consts.ErrNilImage)
	}

	// TODO uses unscaled original image
	// rm, err := timg.File(term)
	_, err = timg.SaveAsFile(tm, `png`, &encpng.PngEncoder{})
	if err != nil {
		return ``, err
	}
	fileName, err := filepath.Abs(timg.FileName)
	if err != nil {
		return ``, err
	}

	tcw, tch, err := tm.SizeInCells()
	if err != nil {
		return ``, err
	}
	if tcw == 0 || tch == 0 {
		return ``, errors.New("could not query terminal dimensions")
	}

	var cleanCanvasStr string
	for y := 0; y <= bounds.Dy(); y++ {
		cleanCanvasStr += fmt.Sprintf("\033[%d;%dH%s", bounds.Min.Y+y+1, bounds.Min.X+1, strings.Repeat(` `, bounds.Dx()))
	}

	widthPercentage := uint(100*bounds.Dx()) / tcw
	heightPercentage := uint(100*bounds.Dy()) / tch
	maxX := uint(bounds.Max.X)
	maxY := uint(bounds.Max.Y)
	if tcw < maxX {
		maxX = tcw
	}
	if tch < maxY {
		maxY = tch
	}
	CenterPosXPercentage := 100 * (uint(bounds.Min.X)) / (tcw - uint(bounds.Dx()))
	CenterPosYPercentage := 100 * (uint(bounds.Min.Y)) / (tch - uint(bounds.Dy()))

	// urxvtString = mux.Wrap(fmt.Sprintf("\033]20;%s;%dx%d+%d+%d:op=keep-aspect\a", fileName, widthPercentage, heightPercentage, CenterPosXPercentage, CenterPosYPercentage), term)
	urxvtString = cleanCanvasStr + mux.Wrap(fmt.Sprintf("\033]20;%s;%dx%d+%d+%d\a", fileName, widthPercentage, heightPercentage, CenterPosXPercentage, CenterPosYPercentage), tm)
	timg.SetInband(bounds, urxvtString, d, tm)

	return urxvtString, nil
}

func (d *drawerURXVT) Clear(term *term.Terminal) error {
	// TODO doesn't clear but upscales to terminal size
	clearStr := mux.Wrap("\033]20;;100x100+1000+1000\a", term)
	_, err := term.WriteString(clearStr)
	return err
}

/*
	executableName, ok1 := t.ExtraProperty(`urxvt_executableName`)
	versionFirstChar, ok2 := t.ExtraProperty(`urxvt_versionFirstChar`)
	versionThirdChar, ok3 := t.ExtraProperty(`urxvt_versionThirdChar`)
*/
