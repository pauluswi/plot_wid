package app

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	dao "github.com/pauluswi/plot_wid/dao"
)

var db dao.DataStore

func createParkingLot(maxSlots int) (string, error) {
	var err error
	db, err = dao.NewInMemoryDB(maxSlots)
	var out string
	if err == nil {
		out = fmt.Sprintf("Created a parking lot with %d slots", maxSlots)
	}
	return out, err
}

func park(regNo, colour string) (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	c := dao.NewCar(regNo, colour)
	slot, err := db.Park(c)
	var out string
	if err != nil {
		out = err.Error()
	} else {
		out = fmt.Sprintf("Allocated slot number: %d", slot.GetNo())
	}
	return out, err
}

func leave(no int) (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	s := dao.NewSlot(no)
	err := db.Leave(s)
	var out string
	if err != nil {
		out = err.Error()
	} else {
		out = fmt.Sprintf("Slot number %d is free", no)
	}
	return out, err
}

func status() (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	slotsList, err := db.GetAll()
	var out string
	if err != nil {
		out = err.Error()
	} else {
		out = fmt.Sprintf("Slot No.    Registration No    Colour")
		for _, slot := range slotsList {
			if slot != nil && slot.GetVehicle() != nil {
				out += fmt.Sprintf("\n%d           %s      %s", slot.GetNo(), slot.GetVehicle().GetRegNo(), slot.GetVehicle().GetColour())
			}
		}
	}
	return out, err
}

func getAllRegNoBycolour(colour string) (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	slotsList, err := db.GetAllSlotsByColour(colour)
	var out string
	if err != nil {
		out = err.Error()
	} else {
		first := true
		for _, slot := range slotsList {
			if slot != nil && slot.GetVehicle() != nil {
				if first {
					out += fmt.Sprintf("%s", slot.GetVehicle().GetRegNo())
				} else {
					out += fmt.Sprintf(", %s", slot.GetVehicle().GetRegNo())
				}
				first = false
			}
		}
	}
	return out, err

}

func getAllSlotNoBycolour(colour string) (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	slotsList, err := db.GetAllSlotsByColour(colour)
	var out string
	if err != nil {
		out = err.Error()
	} else {
		first := true
		for _, slot := range slotsList {
			if slot != nil {
				if first {
					out += fmt.Sprintf("%d", slot.GetNo())
				} else {
					out += fmt.Sprintf(", %d", slot.GetNo())
				}
				first = false
			}
		}
	}
	return out, err
}

func getSlotNoByRegNo(regNo string) (string, error) {
	if db == nil {
		return "", ErrNoParkingLotCreated
	}
	slot, err := db.GetSlotByRegNo(regNo)
	var out string
	if err != nil {
		out = err.Error()
	} else {
		out = fmt.Sprintf("%d", slot.GetNo())
	}
	return out, err

}

func runCommand(command []string) (string, error) {
	switch command[0] {
	case CmdCreateParkingLot:
		maxSlots, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err.Error())
		}
		return createParkingLot(maxSlots)
	case CmdPark:
		return park(command[1], command[2])
	case CmdLeave:
		no, err := strconv.Atoi(command[1])
		if err != nil {
			panic(err.Error())
		}
		return leave(no)
	case CmdStatus:
		return status()
	case CmdRegistrationNumberByColour:
		return getAllRegNoBycolour(command[1])
	case CmdSlotnoByCarColour:
		return getAllSlotNoBycolour(command[1])
	case CmdSlotnoByRegNumber:
		return getSlotNoByRegNo(command[1])
	default:
	}
	return "", ErrCommandNotSupported
}

// TODO create a struct for command with args and parse all strings to that command,
// just like commmand pattern
func splitCommand(command string) []string {
	splitCommand := []string{}

	for _, s := range strings.Split(command, " ") {
		if s != "" {
			splitCommand = append(splitCommand, s)
		}
	}
	return splitCommand
}

func ExecuteFile(path string) error {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		command := splitCommand(scanner.Text())
		//fmt.Println(command)
		output, err := runCommand(command)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(output)
		}
	}
	return nil
}

func ExecuteInlineCommands() error {
	exitCommand := false
	var buffReader *bufio.Reader
	buffReader = bufio.NewReader(os.Stdin)

	for !exitCommand {
		inputText, _ := buffReader.ReadString('\n')
		inputText = strings.TrimRight(inputText, "\r\n")
		if inputText == "exit" {
			break
		}
		runCmdInput(inputText)
	}
	return nil
}

func runCmdInput(inputText string) {
	command := splitCommand(inputText)
	//fmt.Println(command)
	output, err := runCommand(command)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(output)
	}
}
