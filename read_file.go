package main
import ("fmt"
		"os"		
)

func main(){

	data, err := os.ReadFile("greetings.txt")
	if err != nil{
		panic(err)
	}

	// os.ReadFile returns data as a byte slice ([]byte). so we converted into strings
	fmt.Println(string(data)) 
	
}