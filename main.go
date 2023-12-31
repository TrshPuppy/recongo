package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/trshpuppy/recongo/projects/exit"
	"github.com/trshpuppy/recongo/projects/exit/errors"
	"github.com/trshpuppy/recongo/projects/misc/words"
	
	// keyvalue "github.com/sam-caldwell/go/v2/projects/KeyValue"
	// "github.com/sam-caldwell/go/v2/projects/convert"
	// "github.com/sam-caldwell/go/v2/projects/exit"
	// "github.com/sam-caldwell/go/v2/projects/exit/errors"
	// "github.com/sam-caldwell/go/v2/projects/misc/words"
	// cpu "github.com/sam-caldwell/go/v2/projects/systemrecon/cpu"
	// memory "github.com/sam-caldwell/go/v2/projects/systemrecon/memory"
	// opsys "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
	// "github.com/sam-caldwell/go/v2/projects/version"
)

const (
	usage = `
what <option> [--printError
	
options:
  --help          : display usage information
  --version       : display the current version
  --printError    : indicates whether error messages are printed to stdout.

  --cpus          : Return number of CPU cores
  --cpu-arch      : Return cpu architecture (e.g. arm64, amd64)
  --cpu-cache     : Return the L1,L2,L3 CPU cache size in KB (e.g. 32:128:256 for 32 L1, 128 L2 and 256 L3)
  --cpu-info      : Return CPU specifications
  --os            : Return the operating system (e.g. darwin, linux, windows)
  --os-family     : Return an operating system family (e.g. Windows 10)
  --os-version    : Return the operating system version
  --ram           : Return the amount of memory (in KB)qq
  --software-list : Return a list of software installed on the system`
)

func main(){
	var output string
	var err error

	//*
	exit.OnCondition(len(os.Args) < 2, exit.GeneralError, errors.MissingArguments, usage)
	
	var command string = strings.TrimLeft(strings.ToLower(strings.TrimSpace(os.Args[1])), words.Hyphen)
	
	switch command {
	case "tiddies":
		err = fmt.Errorf(errors.InvalidCommandWithDetail, command)
		output = ""

	case "tats":
		output = "acceptable"
		err = nil
	
	default:
		output = ""
		err = fmt.Errorf(errors.MissingArgWithDetail, command)
	}

	if err != nil {
		if (len(os.Args) == 3) && (os.Args[2] == "--printError") {
			fmt.Println(err)
		}
		fmt.Println(err) //*
		os.Exit(exit.GeneralError)
	}

	fmt.Println(output)

	os.Exit(0)

	//switch command := strings.TrimLeft(strings.ToLower(strings.TrimSpace(os.Args[1])), words.Hyphen); command {
	// /*
	//  * CPU-related commands
	//  */
	// case "cpus":
	// 	output, err = convert.IntToStringFuncWrapper(cpu.CpuCores)

	// case "cpu-arch":
	// 	output, err = cpu.CpuArch()

	// case "cpu-cache":
	// 	output, err = cpu.CpuCache()

	// case "cpu-info":
	// 	output, err = keyvalue.Interceptor(cpu.CpuInfo)
	// /*
	//  * operating system stuff
	//  */
	// case "os":
	// 	output, err = opsys.OpSys()

	// case "os-family":
	// 	output, err = opsys.OpSysFamily()

	// case "os-version":
	// 	output, err = opsys.OpSysVersion()
	// /*
	//  * Memory related
	//  */
	// case "ram":
	// 	output, err = convert.IntToStringFuncWrapper(memory.RamSize)

	/*
	* Networking
	*/
	// case "ip":
	// 	output, err = 



	// /*
	//  * General user stuff (help, version, etc)
	//  */
	// case "help":
	// 	fmt.Printf("\n%s", usage)

	// case "version":
	// 	fmt.Printf("%s", version.Version)

	// default:
	// 	output = ""
	// 	err = fmt.Errorf(errors.MissingArgWithDetail, command)
	// }

	// if err != nil {
	// 	if (len(os.Args) == 3) && (os.Args[2] == "--printError") {
	// 		fmt.Println(err)
	// 	}
	// 	os.Exit(exit.GeneralError)
	// }

	// fmt.Println(output)

	// // os.Exit(0)
}