package command

import "fmt"

func Execute() {

	fmt.Println("START:==>")
	position := &Position{12, 14}
	drawCommand := DrawCommand{position}

	position2 := &Position{122, 144}
	drawCommand2 := DrawCommand{position2}
	macroCommand := &MacroCommand{}

	macroCommand.Append(&drawCommand)
	macroCommand.Append(&drawCommand2)
	fmt.Println("macroCommand:==>", macroCommand.commands)

	res := macroCommand.Execute()
	fmt.Println("RESULT:==>", res)
}
