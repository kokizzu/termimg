package propkeys

const (
	QueryCachePrefix            = `queryCache_`
	CheckPrefix                 = `check`
	CheckTermPrefix             = CheckPrefix + `Term`
	CheckTermEnvExclPrefix      = CheckTermPrefix + `EnvExclude_`
	CheckTermQueryIsPrefix      = CheckTermPrefix + `EnvIs_`
	CheckTermWindowIsPrefix     = CheckTermPrefix + `WindowIs_`
	CheckTermCompletePrefix     = CheckTermPrefix + `Complete_`
	EnvPrefix                   = `env_`
	XResourcesPrefix            = `xResources_`
	PreservedOuterEnvPrefix     = GeneralPrefix + `envOuterPreserved_`
	GeneralPrefix               = `general_`
	EnvIsLoaded                 = GeneralPrefix + `envIsLoaded`
	Mode                        = GeneralPrefix + `mode` /// "tui", "cli" (default)
	ManualComposition           = GeneralPrefix + `manualComposition`
	TerminalName                = GeneralPrefix + `termName`
	TerminalPID                 = GeneralPrefix + `termPID`
	TerminalTTY                 = GeneralPrefix + `termTTY` // directly provided tty by the terminal
	PTYName                     = GeneralPrefix + `ptyName` // opened tty
	Executable                  = GeneralPrefix + `executable`
	TempDir                     = GeneralPrefix + `tempDir`
	Passages                    = GeneralPrefix + `passages`
	IsRemote                    = GeneralPrefix + `isRemote`
	AvoidANSI                   = GeneralPrefix + `avoidANSI`
	AvoidDA1                    = GeneralPrefix + `avoidDA1`
	AvoidDA2                    = GeneralPrefix + `avoidDA2`
	AvoidDA3                    = GeneralPrefix + `avoidDA3`
	DeviceAttributesWereQueried = GeneralPrefix + `deviceAttributesWereQueried`
	DeviceAttributes            = GeneralPrefix + `deviceAttributes`
	DeviceClass                 = GeneralPrefix + `deviceClass`
	ReGISCapable                = GeneralPrefix + `regisCapable`
	SixelCapable                = GeneralPrefix + `sixelCapable`
	WindowingCapable            = GeneralPrefix + `windowingCapable`
	DA3ID                       = GeneralPrefix + `DA3ID`
	DA3IDHex                    = GeneralPrefix + `DA3IDHex`
	XTVERSION                   = GeneralPrefix + `XTVERSION`
	XTGETTCAPPrefix             = GeneralPrefix + `XTGETTCAP_`
	XTGETTCAPKeyNamePrefix      = XTGETTCAPPrefix + `keyName_`
	XTGETTCAPSpecialPrefix      = XTGETTCAPPrefix + `special_`
	XTGETTCAPSpecialTN          = XTGETTCAPSpecialPrefix + `TN`
	XTGETTCAPSpecialCo          = XTGETTCAPSpecialPrefix + `Co`
	XTGETTCAPSpecialRGB         = XTGETTCAPSpecialPrefix + `RGB`
	XTGETTCAPInvalidPrefix      = XTGETTCAPPrefix + `invalid_`
	// TODO positioning of uncropped image is dependent on terminal size (e.g. urxvt)
	FullImgPosDepOnSize = GeneralPrefix + `fullImagePositionDependentOnSize`

	// TermChecker might set WindowBorder…Estimated+`_termname`
	WindowPrefix                = `window_`
	WindowBorderPrefix          = WindowPrefix + `border`
	WindowBorderEstimated       = WindowBorderPrefix + `Estimated`
	WindowBorderLeft            = WindowBorderPrefix + `Left`
	WindowBorderRight           = WindowBorderPrefix + `Right`
	WindowBorderTop             = WindowBorderPrefix + `Top`
	WindowBorderBottom          = WindowBorderPrefix + `Bottom`
	WindowBorderLeftEstimated   = WindowBorderLeft + `Estimated`
	WindowBorderRightEstimated  = WindowBorderRight + `Estimated`
	WindowBorderTopEstimated    = WindowBorderTop + `Estimated`
	WindowBorderBottomEstimated = WindowBorderBottom + `Estimated`

	PlatformPrefix = `platform_`
	SystemD        = PlatformPrefix + `systemd`
	RunsOnWine     = PlatformPrefix + `wine`

	// Terminal type properties
	TerminalPrefix        = `terminal_`
	AppleTermPrefix       = TerminalPrefix + `apple_`
	AppleTermVersion      = AppleTermPrefix + `version` // CFBundleVersion of Terminal.app
	ContourPrefix         = TerminalPrefix + `contour_`
	ContourVersion        = ContourPrefix + `version`
	DomTermPrefix         = TerminalPrefix + `domterm_`
	DomTermLibWebSockets  = DomTermPrefix + `libwebsockets`
	DomTermSession        = DomTermPrefix + `session`
	DomTermTTY            = DomTermPrefix + `tty`
	DomTermPID            = DomTermPrefix + `pid`
	DomTermVersion        = DomTermPrefix + `version`
	DomTermWindowName     = DomTermPrefix + `windowName`
	DomTermWindowInstance = DomTermPrefix + `windowInstance`
	ITerm2Prefix          = TerminalPrefix + `iterm2_`
	ITerm2Version         = ITerm2Prefix + `version`
	KittyPrefix           = TerminalPrefix + `kitty_`
	KittyWindowID         = KittyPrefix + `windowID` // tab id
	MacTermPrefix         = TerminalPrefix + `macterm_`
	MacTermBuildNr        = MacTermPrefix + `buildNumber` // YYYYMMDD
	MltermPrefix          = TerminalPrefix + `mlterm_`
	MltermVersion         = MltermPrefix + `version`
	MinttyPrefix          = TerminalPrefix + `mintty_`
	MinttyShortcut        = MinttyPrefix + `shortcut`
	MinttyVersion         = MinttyPrefix + `version`
	URXVTPrefix           = TerminalPrefix + `urxvt_`
	URXVTExeName          = URXVTPrefix + `executableName`
	URXVTVerFirstChar     = URXVTPrefix + `versionFirstChar`
	URXVTVerThirdChar     = URXVTPrefix + `versionThirdChar`
	VSCodePrefix          = TerminalPrefix + `vscode_`
	VSCodeVersion         = VSCodePrefix + `version`
	VSCodeVersionMajor    = VSCodeVersion + `Major`
	VSCodeVersionMinor    = VSCodeVersion + `Minor`
	VSCodeVersionPatch    = VSCodeVersion + `Patch`
	HyperPrefix           = TerminalPrefix + `hyper_`
	HyperVersion          = HyperPrefix + `version`
	HyperVersionMajor     = HyperVersion + `Major`
	HyperVersionMinor     = HyperVersion + `Minor`
	HyperVersionPatch     = HyperVersion + `Patch`
	HyperVersionCanary    = HyperVersion + `Canary`
	VTEPrefix             = TerminalPrefix + `vte_`
	VTEVersion            = VTEPrefix + `version`
	VTEVersionMajor       = VTEVersion + `Major`
	VTEVersionMinor       = VTEVersion + `Minor`
	VTEVersionPatch       = VTEVersion + `Patch`
	WezTermPrefix         = TerminalPrefix + `wezterm_`
	WezTermExe            = WezTermPrefix + `executable`
	WezTermExeDir         = WezTermPrefix + `executableDir`
	WezTermPane           = WezTermPrefix + `pane`
	WezTermUnixSocket     = WezTermPrefix + `unixSocket`
	XTermPrefix           = TerminalPrefix + `xterm_`
	XTermVersion          = XTermPrefix + `version`
)
