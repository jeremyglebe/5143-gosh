package main

import(
    "fmt"
    "os"
    //"path/filepath"
)

// A command grouped with arguments for calling it
type CommandLine struct {
    comm string
    args []string
}

// PipeLine
// Special version of the execute function which takes a list of Commands, then
// for each Command, sends it's output to a file and uses that file as the first
// arg of the next command.
func PipeLine(commands []CommandLine){

    fmt.Println("PipeLine starting!\nfmt imported for debugging...")

    // stdout backup
    stdout := os.Stdout;
    // Path to the pipe file
    //pipeFilePath := filepath.Join(os.TempDir(), "gosh.pipe.tmp")
    pipeFilePath := "./gosh.pipe.tmp"
    // Create the actual pipe file
    pipeFile, _ := os.Create(pipeFilePath)
    // Send output to the pipe file
    os.Stdout = pipeFile

    // For each command in the array
    for i, pipe := range commands {

        // If the Command is valid
        if com, valid := ComMap[pipe.comm]; valid{
            
            // If this isn't the first command
            if i > 0{
                // We need to add the pipe file to the args (at the front)
                frAddStr(pipe.args, pipeFilePath)
            }

            // If the command is the last command in the PipeLine
            if i == len(commands) - 1 {
                // Restore stdout
                os.Stdout = stdout
            }

            // Execute the command with its arguments
            com(pipe.args)
        }

    // After processing each command, close the pipe file
    pipeFile.Close()

    }

}

func frAddStr(argList []string, arg string) []string {
    argList = append(argList, "")
    copy(argList[1:], argList)
    argList[0] = arg
    return argList
}
