package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/eiannone/keyboard"
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

// GetKeyboardInput reads a single key press
func GetKeyboardInput() (string, error) {
	err := keyboard.Open()
	if err != nil {
		return "", err
	}
	defer keyboard.Close()

	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		return "", err
	}

	return string(char), nil
}

// Pause waits for the user to press any key
func Pause() {
	fmt.Print("[ Press any key ]")
	err := keyboard.Open()
	if err != nil {
		fmt.Printf("Error opening keyboard: %v\n", err)
		return
	}
	defer keyboard.Close()

	// Wait for a single key press
	_, _, err = keyboard.GetSingleKey()
	if err != nil {
		fmt.Printf("Error reading key: %v\n", err)
	}
}

// Move cursor to X, Y location
func MoveCursor(x int, y int) {
	fmt.Printf(Esc+"%d;%df", y, x)
}

// WaitForAnyKey waits for a user to press any key to continue.
func WaitForAnyKey() error {
	// Open the keyboard listener
	err := keyboard.Open()
	if err != nil {
		return err
	}
	defer keyboard.Close() // Ensure that the keyboard listener is closed when done

	// Wait for a single key press
	_, _, err = keyboard.GetSingleKey()
	if err != nil {
		return err
	}

	return nil
}

// Improved function to display ANSI file content
func DisplayAnsiFile(filePath string, delay int, localDisplay bool) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file %s: %v", filePath, err)
	}
	PrintAnsi(string(content), delay, localDisplay)
}

func PrintAnsi(content string, delay int, localDisplay bool) {
	lines := strings.Split(content, "\r\n")
	for _, line := range lines {
		fmt.Println(line)
		time.Sleep(time.Millisecond * time.Duration(delay))
	}
}

// TrimStringFromSauce trims SAUCE metadata from a string.
func TrimStringFromSauce(s string) string {
	return trimMetadata(s, "COMNT", "SAUCE00")
}

// trimMetadata trims metadata based on delimiters.
func trimMetadata(s string, delimiters ...string) string {
	for _, delimiter := range delimiters {
		if idx := strings.Index(s, delimiter); idx != -1 {
			return trimLastChar(s[:idx])
		}
	}
	return s
}

// trimLastChar trims the last character from a string.
func trimLastChar(s string) string {
	if len(s) > 0 {
		_, size := utf8.DecodeLastRuneInString(s)
		return s[:len(s)-size]
	}
	return s
}

// Print text at an X, Y location
func PrintStringLoc(text string, x int, y int) {
	fmt.Fprintf(os.Stdout, Esc+strconv.Itoa(y)+";"+strconv.Itoa(x)+"f"+text)
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
