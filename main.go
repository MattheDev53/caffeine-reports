package main

import (
	"fmt"
	"os"
	"strings"
	typst "github.com/Dadido3/go-typst"
)

func NoImp() {
	fmt.Println("!! Not Implemented : Nothing will be done")
}

// TODO: Figure out how to read the Git Tag
func PrintVersion () {
	fmt.Printf("CaffeineReports %s\n", "v0.1.1")
	fmt.Printf("")
}

func PrintCredits() {
	PrintLogo()
	fmt.Println("Author & Maintainer: MattheDev53")
	fmt.Println("GitHub: https://github.com/MattheDev53/CaffeineReports")
	fmt.Println("")
	fmt.Println("CaffeineReports is Open Source under the GNU AGPLv3")
	// fmt.Println("                                                                               ")
}

func GetSelfFilename() string {
	argZeroList := strings.Split(os.Args[0], string(os.PathSeparator))
	filename := argZeroList[len(argZeroList)-1]
	return filename
}

func PrintUsage() {
	filename := GetSelfFilename()
	fmt.Printf("Usage of %s:\n", filename)
	fmt.Printf("%s help : Prints this help message\n", filename)
	fmt.Printf("%s generate [template] [data] [output] : generates a report from the provided data and template\n", filename)
	fmt.Printf("%s version : Prints Version Information\n", filename)
	fmt.Printf("%s credits : Prints Credits\n", filename)
}

func GenerateReport(templateFile string, dataFile string, outputFile string) {
	fmt.Printf("Template: %s\n", templateFile)
	fmt.Printf("Data:     %s\n", dataFile)
	fmt.Printf("Output:   %s\n", outputFile)

	inputs := map[string]string{
		"data": dataFile,
	}
	options := typst.CLIOptions{
		Input: inputs,
	}

	typstCLI := typst.CLI{}

	in, err := os.Open(templateFile);

	if err != nil {
		fmt.Printf("Error reading %s: %v\n", templateFile, err)
		os.Exit(1)
	}
	defer in.Close()

	out, err := os.Create(outputFile)

	if err != nil {
		fmt.Printf("Error creating %s: %v\n", outputFile, err)
		os.Exit(1)
	}
	defer out.Close()

	fmt.Println("Compiling Report...")
	if err := typstCLI.Compile(in, out, &options); err != nil {
		fmt.Printf("Error compiling report: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Report Generated Successfully!")
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No arguments provided")
		PrintUsage()
		os.Exit(3)
	}
	switch os.Args[1] {
		default:
			PrintUsage()
			os.Exit(3)
		case "version":
			PrintVersion()
			os.Exit(0)
		case "credits":
			PrintCredits()
			os.Exit(0)
		case "help":
			PrintUsage()
			os.Exit(0)
		case "generate":
			if len(os.Args) != 5 {
				fmt.Printf("Invalid number of arguments (%d provided, 3 expected)\n", len(os.Args)-2)
				os.Exit(3)
			}
			GenerateReport(os.Args[2], os.Args[3], os.Args[4])
			os.Exit(0)
	}
}

func PrintLogo() {
	fmt.Println("   ______      ________     _            ____                        __      ")
	fmt.Println(`  / ____/___ _/ __/ __/__  (_)___  ___  / __ \___  ____  ____  _____/ /______`)
	fmt.Println(" / /   / __ `/ /_/ /_/ _ \\/ / __ \\/ _ \\/ /_/ / _ \\/ __ \\/ __ \\/ ___/ __/ ___/")
	fmt.Println("/ /___/ /_/ / __/ __/  __/ / / / /  __/ _, _/  __/ /_/ / /_/ / /  / /_(__  ) ")
	fmt.Println(`\____/\__,_/_/ /_/  \___/_/_/ /_/\___/_/ |_|\___/ .___/\____/_/   \__/____/  `)
	fmt.Println("                                               /_/                           ")
}
