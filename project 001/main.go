package main
import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
	"runtime"
)

//average marks calculator
func calculateAverage(grade map[string]int) float64 {
	var total int
	for _, mark := range grade {
		total += mark
	}
	return float64(total) / float64(len(grade))
}
//formatter function to format the output
func formatter(grade map[string]int,average float64,name string) string{
	var result string
	result += "Name: "+name+"\n"
	result += "subject\tMarks\n"
	for sub,mark :=range(grade) {
		result+= fmt.Sprintf("%v\t%v\n", sub, mark)
	}
	result += fmt.Sprintf("Average: %.3f\n", average)
	return result
}
func ClearScreen() {
	var cmd *exec.Cmd 

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	case "linux", "darwin": 
		cmd = exec.Command("clear")
	default:
		fmt.Println("Warning: Clear screen not supported on this OS.")
		return 
	}


	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error clearing screen: %v\n", err)
	}
}
func main(){
	reader:= bufio.NewReader(os.Stdin)
	var name string
	grade:=make(map[string]int)
	var num_sub int
	fmt.Print("Enter your name: ")
	name, _ = reader.ReadString('\n')
	name = name[:len(name)-1]

	fmt.Print("Enter the number of subjects: ")
	fmt.Scanln(&num_sub)

	for i:=0;i<num_sub;i++{
		var sub string
		var mark int

		fmt.Print("Enter subject name: ")
		fmt.Scanln(&sub)
		if _, exists := grade[sub]; exists {
			fmt.Println("Subject already exists. Please enter a different subject.")
			i--
			continue
		}
		fmt.Print("Enter marks for ", sub, ": ")
		fmt.Scanln(&mark)
		
		if mark<0 ||mark>100{
			fmt.Println("Invalid marks. Please enter a value between 0 and 100.")
			i--
			continue
		}
		grade[sub] = mark
	}
		average := calculateAverage(grade)
		ClearScreen()
		fmt.Printf(formatter(grade, average, name))
	}
