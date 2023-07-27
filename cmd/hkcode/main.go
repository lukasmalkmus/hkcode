package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/lukasmalkmus/hkcode/hk"
	"github.com/lukasmalkmus/hkcode/hk/qr"
	"github.com/lukasmalkmus/hkcode/hk/text"
)

const usage = `Usage:
    hkcode --text [-o OUTPUT] [SETUP_CODE]
    hkcode --qr [-b BOOL] [-i SETUP_ID] [-f SETUP_FLAG]... [-c CATEGORY]
           [-o OUTPUT] [SETUP_CODE]

Options:
    -t, --text               Create a text based Apple HomeKit® setup code.
    -q, --qr                 Create a QR code based Apple HomeKit® setup code.
    -o, --output OUTPUT      Write the result to the file at path OUTPUT.
    -b, --box BOOL           Box the QR code with a text code and the Apple
                             HomeKit® logo. Optional.
    -i, --id SETUP_ID        Four character setup id.
    -f, --flag SETUP_FLAG    Describes the accessories supported pairing
                             methods. Optional.
    -c, --category CATEGORY  Category of the accessory. Optional.
    
If SETUP_CODE is ommited as an argument, it will default to standard input. Must
be a number between 0 and 99999999, no padding required. Trivial setup codes are
accepted but not recommended.

If OUTPUT exists, it will be overwritten. OUTPUT is png encoded.

SETUP_FLAG is one of "nfc", "ip" or "btle".

CATEGORY is one of the following: other, bridge, fan, garage_door_opener,
lightbulb, door_lock, outlet, switch, thermostat, sensor, security_system, door,
window, window_covering, programmable_switch, ip_camera, video_doorbell,
air_purifier, heater, air_conditioner, humidifier, dehumidifier, sprinklers,
faucets, shower_systems.

Example:
    $ hkcode --text -o=code.png 12344321
    $ hkcode --qr -o=code.png -i=MHKA -f=ip -f=btle -c=outlet 12344321
    $ shuf -i 1-99999999 -n 1 | hkcode --qr -b -o=code.png -i=MHKA -f=ip -c=switch
`

type multiFlag []string

func (f *multiFlag) String() string { return fmt.Sprint(*f) }

func (f *multiFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}

type categoryFlag struct {
	hk.Category
}

func (f categoryFlag) String() string { return f.Category.String() }

func (f *categoryFlag) Set(value string) error {
	for c := hk.CategoryUnknown; c <= hk.CategoryShowerSystems; c++ {
		if strings.EqualFold(c.String(), value) {
			*f = categoryFlag{c}
			return nil
		}
	}
	return fmt.Errorf("unknown category %q", value)
}

var version string

func main() {
	log.SetFlags(0)

	flag.Usage = func() { fmt.Fprint(os.Stderr, usage) }

	var (
		versionFlag   bool
		textFlag      bool
		qrFlag        bool
		outFlag       string
		boxFlag       bool
		setupIDFlag   string
		setupFlagFlag multiFlag
		categoryFlag  categoryFlag
	)

	flag.BoolVar(&versionFlag, "version", false, "print the version")
	flag.BoolVar(&textFlag, "t", false, "create text code")
	flag.BoolVar(&textFlag, "text", false, "create text code")
	flag.BoolVar(&qrFlag, "q", false, "create qr code")
	flag.BoolVar(&qrFlag, "qr", false, "create qr code")
	flag.StringVar(&outFlag, "o", "", "output to `FILE`")
	flag.StringVar(&outFlag, "output", "", "output to `FILE`")
	flag.BoolVar(&boxFlag, "b", false, "create boxed qr code")
	flag.BoolVar(&boxFlag, "box", false, "create boxed qr code")
	flag.StringVar(&setupIDFlag, "i", "", "setup id")
	flag.StringVar(&setupIDFlag, "id", "", "setup id")
	flag.Var(&setupFlagFlag, "f", "supported pairing methods")
	flag.Var(&setupFlagFlag, "flag", "supported pairing methods")
	flag.Var(&categoryFlag, "c", "accessory category")
	flag.Var(&categoryFlag, "category", "accessory category")

	flag.Parse()

	if versionFlag {
		if version != "" {
			fmt.Println(version)
			return
		} else if buildInfo, ok := debug.ReadBuildInfo(); ok {
			fmt.Println(buildInfo.Main.Version)
			return
		} else {
			fmt.Println("(unknown)")
		}
		return
	}

	if flag.NArg() > 1 {
		errorWithHint(fmt.Sprintf("too many arguments: %q", flag.Args()),
			"note that the setup code must be specified after all flags")
	}

	switch {
	case textFlag:
		if qrFlag {
			errorf("-q/--qr can't be used with -t/--text")
		}
		if boxFlag {
			errorf("-b/--box can't be used with -t/--text")
		}
		if len(setupIDFlag) > 0 {
			errorf("-i/--id can't be used with -t/--text")
		}
		if len(setupFlagFlag) > 0 {
			errorf("-f/--flag can't be used with -t/--text")
		}
		if categoryFlag.Category > 0 {
			errorf("-c/--category can't be used with -t/--text")
		}
	case qrFlag:
		if textFlag {
			errorf("-t/--text can't be used with -q/--qr")
		}
	default:
		errorWithHint("missing mode",
			"did you forget to specify one of -t/--text or -q/--qr?")
	}

	if len(outFlag) == 0 {
		errorWithHint("missing output file",
			"did you forget to specify -o/--output?")
	}

	var setupCodeStr string
	if setupCodeStr = flag.Arg(0); setupCodeStr == "" || setupCodeStr == "-" {
		setupCode, err := io.ReadAll(os.Stdin)
		if err != nil {
			errorf("failed to read setup code from stdin: %v", err)
		}
		setupCodeStr = strings.TrimSuffix(string(setupCode), "\n")
	}

	if setupCodeStr == "" {
		errorWithHint("setup code is empty",
			"set it as the last argument after all flags or pipe it via stdin")
	}

	var setupCode hk.Code
	setupCode64, err := strconv.ParseUint(setupCodeStr, 10, 32)
	if err != nil {
		errorf("failed to parse setup code: %v", err)
	}
	setupCode = hk.Code(setupCode64)

	out := newLazyOpener(outFlag)
	defer func() {
		if err := out.Close(); err != nil {
			errorf("failed to close output file %q: %v", outFlag, err)
		}
	}()

	var outImg image.Image
	switch {
	case textFlag:
		outImg, err = text.CreateCode(setupCode)
	case qrFlag && !boxFlag:
		outImg, err = qr.CreateCode(setupCode, hk.ID(setupIDFlag), hk.FlagNone, categoryFlag.Category)
	case qrFlag && boxFlag:
		outImg, err = qr.CreateBoxedCode(setupCode, hk.ID(setupIDFlag), hk.FlagNone, categoryFlag.Category)
	}
	if err != nil {
		errorf("failed to create code: %v", err)
	}

	if err := png.Encode(out, outImg); err != nil {
		errorf("failed to encode code: %v", err)
	}
}

type lazyOpener struct {
	name string
	f    *os.File
	err  error
}

func newLazyOpener(name string) io.WriteCloser {
	return &lazyOpener{name: name}
}

func (l *lazyOpener) Write(p []byte) (n int, err error) {
	if l.f == nil && l.err == nil {
		l.f, l.err = os.Create(l.name)
	}
	if l.err != nil {
		return 0, l.err
	}
	return l.f.Write(p)
}

func (l *lazyOpener) Close() error {
	if l.f != nil {
		return l.f.Close()
	}
	return nil
}

func warnWithHint(msg string, hints ...string) {
	log.Printf("hkcode: warning: %s", msg)
	for _, hint := range hints {
		log.Printf("hkcode: hint: %s", hint)
	}
}

func errorf(format string, v ...any) {
	log.Fatalf("hkcode: error: "+format, v...)
}

func errorWithHint(msg string, hints ...string) {
	log.Printf("hkcode: error: %s", msg)
	for _, hint := range hints {
		log.Printf("hkcode: hint: %s", hint)
	}
	os.Exit(1)
}
