package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode/utf8"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"gopkg.in/ini.v1"
)

// Struct for organizing all ANSI escape sequences
type AnsiEscapes struct {
	EraseScreen        string
	CursorHide         string
	CursorShow         string
	CursorBackward     string
	CursorPrevLine     string
	CursorLeft         string
	CursorTop          string
	CursorTopLeft      string
	CursorBlinkEnable  string
	CursorBlinkDisable string
	ScrollUp           string
	ScrollDown         string
	TextInsertChar     string
	TextDeleteChar     string
	TextEraseChar      string
	TextInsertLine     string
	TextDeleteLine     string
	EraseRight         string
	EraseLeft          string
	EraseLine          string
	EraseDown          string
	EraseUp            string
	Black              string
	Red                string
	Green              string
	Yellow             string
	Blue               string
	Magenta            string
	Cyan               string
	White              string
	BlackHi            string
	RedHi              string
	GreenHi            string
	YellowHi           string
	BlueHi             string
	MagentaHi          string
	CyanHi             string
	WhiteHi            string
	BgBlack            string
	BgRed              string
	BgGreen            string
	BgYellow           string
	BgBlue             string
	BgMagenta          string
	BgCyan             string
	BgWhite            string
	BgBlackHi          string
	BgRedHi            string
	BgGreenHi          string
	BgYellowHi         string
	BgBlueHi           string
	BgMagentaHi        string
	BgCyanHi           string
	BgWhiteHi          string
	Reset              string
	// Add more as needed
}

// Struct for organizing all font ANSI sequences
type Fonts struct {
	Mosoul          string
	Potnoodle       string
	Microknight     string
	MicroknightPlus string
	Topaz           string
	TopazPlus       string
	Ibm             string
	IbmThin         string
}

// Initialize ansi with ANSI escape sequences
var ansi = AnsiEscapes{
	EraseScreen:        "\033[2J",
	CursorHide:         "\033[?25l",
	CursorShow:         "\033[?25h",
	CursorBackward:     "\033[D",
	CursorPrevLine:     "\033[F",
	CursorLeft:         "\033[G",
	CursorTop:          "\033[d",
	CursorTopLeft:      "\033[H",
	CursorBlinkEnable:  "\033[?12h",
	CursorBlinkDisable: "\033[?12l",
	ScrollUp:           "\033[S",
	ScrollDown:         "\033[T",
	TextInsertChar:     "\033[@",
	TextDeleteChar:     "\033[P",
	TextEraseChar:      "\033[X",
	TextInsertLine:     "\033[L",
	TextDeleteLine:     "\033[M",
	EraseRight:         "\033[K",
	EraseLeft:          "\033[1K",
	EraseLine:          "\033[2K",
	EraseDown:          "\033[J",
	EraseUp:            "\033[1J",
	Black:              "\033[30m",
	Red:                "\033[31m",
	Green:              "\033[32m",
	Yellow:             "\033[33m",
	Blue:               "\033[34m",
	Magenta:            "\033[35m",
	Cyan:               "\033[36m",
	White:              "\033[37m",
	BlackHi:            "\033[30;1m",
	RedHi:              "\033[31;1m",
	GreenHi:            "\033[32;1m",
	YellowHi:           "\033[33;1m",
	BlueHi:             "\033[34;1m",
	MagentaHi:          "\033[35;1m",
	CyanHi:             "\033[36;1m",
	WhiteHi:            "\033[37;1m",
	BgBlack:            "\033[40m",
	BgRed:              "\033[41m",
	BgGreen:            "\033[42m",
	BgYellow:           "\033[43m",
	BgBlue:             "\033[44m",
	BgMagenta:          "\033[45m",
	BgCyan:             "\033[46m",
	BgWhite:            "\033[47m",
	BgBlackHi:          "\033[40;1m",
	BgRedHi:            "\033[41;1m",
	BgGreenHi:          "\033[42;1m",
	BgYellowHi:         "\033[43;1m",
	BgBlueHi:           "\033[44;1m",
	BgMagentaHi:        "\033[45;1m",
	BgCyanHi:           "\033[46;1m",
	BgWhiteHi:          "\033[47;1m",
	Reset:              "\033[0m",
}

// Initialize fonts with ANSI escape sequences for SyncTerm-supported fonts
var fonts = Fonts{
	Mosoul:          "\033[0;38 D",
	Potnoodle:       "\033[0;37 D",
	Microknight:     "\033[0;41 D",
	MicroknightPlus: "\033[0;39 D",
	Topaz:           "\033[0;42 D",
	TopazPlus:       "\033[0;40 D",
	Ibm:             "\033[0;0 D",
	IbmThin:         "\033[0;26 D",
}

const (
	Esc = "\u001B["
	Osc = "\u001B]"
	Bel = "\u0007"
)

// Example usage of fonts
// fmt.Println(fonts.Mosoul + "This text is in Mosoul font" + ansi.Reset)
// fmt.Println(fonts.IbmThin + "This text is in IBM Thin font" + ansi.Reset)

// Example usage of the structured ANSI sequences
// fmt.Println(ansi.EraseScreen)                           // Clears the screen
// fmt.Println(ansi.Red + "This text is red" + ansi.Reset) // Print red text
// fmt.Println(ansi.CursorHide)                            // Hides the cursor

// GetKeyboardInput reads a single key press using atomicgo keyboard
func GetKeyboardInput() (string, error) {
	var input string
	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		input = key.String()
		return true, nil
	})
	return input, err
}

// Pause waits for any key press
func Pause() {
	fmt.Print("\r\n[ Press any key to continue ]")
	_, _ = GetKeyboardInput()
}

// Prompt user for input using keyboard.Listen()
func Prompt(label string) string {
	fmt.Printf("%s: ", label)
	var input strings.Builder

	err := keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch {
		case key.Code == keys.Enter:
			return true, nil
		case key.Code == keys.Backspace:
			if input.Len() > 0 {
				inputStr := input.String()
				input.Reset()
				input.WriteString(inputStr[:len(inputStr)-1])
				fmt.Print("\b \b")
			}
		case key.Code == keys.Space:
			input.WriteString(" ")
			fmt.Print(" ")
		default:
			input.WriteString(key.String())
			fmt.Print(key.String())
		}
		return false, nil
	})

	if err != nil {
		fmt.Printf("\r\nError reading input: %v\n", err)
		return ""
	}

	return strings.TrimSpace(input.String())
}

// PromptInt prompts the user for integer input
func PromptInt(label string) int {
	for {
		inputStr := Prompt(label)
		num, err := strconv.Atoi(inputStr)
		if err == nil {
			return num
		}
		fmt.Println("\r\nInvalid input. Please enter a valid number.")
	}
}

// LoadConfig loads configuration from an ini file
func LoadConfig(filePath string) (*Config, error) {
	cfg, err := ini.Load(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read ini file: %w", err)
	}

	config := &Config{
		AdminSecurityLevel: cfg.Section("Settings").Key("AdminSecurityLevel").MustInt(255),
		WWIVnet:            cfg.Section("Settings").Key("WWIVnet").MustBool(false),
		FTN:                cfg.Section("Settings").Key("FTN").MustBool(false),
	}

	return config, nil
}

// Finds the drop file in a case-insensitive way
func FindDropFile(path string) (string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return "", fmt.Errorf("error reading directory: %w", err)
	}

	for _, file := range files {
		if strings.EqualFold(file.Name(), "door32.sys") {
			return filepath.Join(path, file.Name()), nil
		}
	}

	return "", errors.New("door32.sys file not found")
}

// Reads and parses the drop file data
func GetDropFileData(path string) (DropFileData, error) {
	filePath, err := FindDropFile(path)
	if err != nil {
		return DropFileData{}, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return DropFileData{}, fmt.Errorf("error opening drop file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	if len(text) < 11 {
		return DropFileData{}, errors.New("drop file has insufficient lines")
	}

	// Parse each line
	commType, _ := strconv.Atoi(text[0])
	commHandle, _ := strconv.Atoi(text[1])
	baudRate, _ := strconv.Atoi(text[2])
	bbsID := text[3]
	userRecordPos, _ := strconv.Atoi(text[4])
	realName := text[5]
	alias := text[6]
	securityLevel, _ := strconv.Atoi(text[7])
	timeLeft, _ := strconv.Atoi(text[8])
	emulation, _ := strconv.Atoi(text[9])
	nodeNum, _ := strconv.Atoi(text[10])

	return DropFileData{
		CommType:      commType,
		CommHandle:    commHandle,
		BaudRate:      baudRate,
		BBSID:         bbsID,
		UserRecordPos: userRecordPos,
		RealName:      realName,
		Alias:         alias,
		SecurityLevel: securityLevel,
		TimeLeft:      timeLeft,
		Emulation:     emulation,
		NodeNum:       nodeNum,
	}, nil
}

// Print text at a specific X, Y location (optional)
func PrintStringLoc(text string, x int, y int) {
	fmt.Fprintf(os.Stdout, "\033[%d;%df%s", y, x, text)
}

// TrimStringFromSauce removes metadata from strings
func TrimStringFromSauce(s string) string {
	return trimMetadata(s, "COMNT", "SAUCE00")
}

// Helper to trim metadata based on delimiters
func trimMetadata(s string, delimiters ...string) string {
	for _, delimiter := range delimiters {
		if idx := strings.Index(s, delimiter); idx != -1 {
			return trimLastChar(s[:idx])
		}
	}
	return s
}

// trimLastChar trims the last character from a string
func trimLastChar(s string) string {
	if len(s) > 0 {
		_, size := utf8.DecodeLastRuneInString(s)
		return s[:len(s)-size]
	}
	return s
}
