package main

import (
	"fmt"
	"learngocui/model"
	"learngocui/tui"
	"learngocui/tui/events"
	"learngocui/tui/logger"
	"learngocui/tui/vm"
	"net/mail"
	"os"
	"os/exec"
	"strings"
)

func main() {

	e := events.NewEventManager()
	seed := model.SeedData()

	accVm1 := vm.NewAccountVM(e, &seed[0])
	accVm2 := vm.NewAccountVM(e, &seed[1])
	accVm3 := vm.NewAccountVM(e, &seed[2])
	accVm4 := vm.NewAccountVM(e, &seed[3])

	err, notmuchAcc := parseShellCommand("notmuch search --output=files --format=text folder:support@essayclip.com/INBOX")
	if err != nil {
		panic(err)
	}
	accVm5 := vm.NewAccountVM(e, notmuchAcc)

	logger := logger.NewLogger("log.txt")
	viewModel := vm.NewVM(e, []*vm.AccountVM{accVm1, accVm2, accVm3, accVm4, accVm5}, logger)

	tui.Init(e, viewModel, logger)
}

func parseShellCommand(command string) (error, *model.EmailAccount) {
	executable := strings.Split(command, " ")[0]
	args := strings.Split(command, " ")[1:]
	fmt.Println("executable: " + executable)
	fmt.Println("arguments: " + strings.Join(args, " "))

	raw_output, err := exec.Command(executable, args...).Output()

	if err != nil {
		panic(err.Error())
	}
	fmt.Print(raw_output)
	output := strings.TrimSpace(string(raw_output))
	files := strings.Split(string(output), "\n")

	fmt.Print(output)

	acc := &model.EmailAccount{
		ShortName: "Notmuch inbox",
	}

	for _, path := range files {
		// email := model.Email{
		// 	Path:    path,
		// 	Subject: path,
		// }
		email := parseMaildirFile(path)
		acc.Emails = append(acc.Emails, email)
	}

	return nil, acc
}

func parseMaildirFile(path string) model.Email {
	fmt.Println(path)
	raw_msg_byte, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	raw_msg := string(raw_msg_byte)

	r := strings.NewReader(raw_msg)
	m, err := mail.ReadMessage(r)
	if err != nil {
		panic(err)
	}
	//body, err := io.ReadAll(m.Body)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("mu", "view", path)
	byte_output, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}
	string_output := strings.TrimSpace(string(byte_output))

	return model.Email{
		Date:    m.Header.Get("Date"),
		From:    m.Header.Get("From"),
		To:      m.Header.Get("To"),
		Subject: m.Header.Get("Subject"),
		Text:    string_output,
	}
}
