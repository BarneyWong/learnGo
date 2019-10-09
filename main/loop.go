package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBinary(n int) string{
	result:=""
	if n ==0 {
		result="0"
	}
	for ;n>0;n/=2{
		lsb:=n%2
		result=strconv.Itoa(lsb)+result
	}
	return result
}
func printFile(filename string){
	file,err:=os.Open(filename)
	if err!= nil{
		fmt.Println(err)
	}

	scanner:= bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}
func main( ) {
	fmt.Println(
		convertToBinary(5),
		convertToBinary(10),
		convertToBinary(64),
		convertToBinary(0),
		)

	printFile("abc.txt")
}