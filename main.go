package main

import (
    "fmt"
    "io/ioutil"
    "math/rand"
    "time"
)

func wordToGuess() string {
    var elements []string
    fileName := "resources.txt"
    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Println(err.Error())
    }
    start := 0
    end := 0
    for i := 0; i < len(data); i++ {
        if string(data[i]) == "\n" {
            end = i
            elements = append(elements, string(data[start:end]))
            start = i + 1
        }
        if i == len(data)-1 {
            elements = append(elements, string(data[end+1:]))
        }
    }
    rand.Seed(time.Now().Unix())
    random_number := rand.Intn(len(elements))
    return elements[random_number]
}

func revealHint(s string) []string{
    var sToDisplay []string
    rand.Seed(time.Now().Unix())
    lettre := string(s[rand.Intn(len(s)-1)])
    for i := 0; i < len(s); i++ {
        if string(s[i]) == lettre {
            sToDisplay = append(sToDisplay, lettre)
        }else if i != len(s)-1 {
            sToDisplay = append(sToDisplay, " _ ")
        }
    }
    return sToDisplay
}
func ContainAny( s string, given string) bool{
    fmt.Println("You have choosen the letter :", given)
    for _, g:= range s{
        if given==string(g){
         return true
        }else{
            return false
        }
    }
}

func ShowResult() string{
    mistakes:=0
    if ContainAny(s string, given string ){
		for i,v := range s {
			if v==given{
				sToDisplay[i]=given
			}
		} 
        return sToDisplay, ShowHangmansPositions()
    }else if !ContainAny(s string, given string){
        mistakes++
       return  ShowError(mistakes), ShowHangmansPositions(mistakes)
    }
}
func ShowError(msitakes int) string{
attempts:=10-mistakes
fmt.Println("You have only %v attempts remaining!!",attempts)
ShowHangmansPositions(mistakes)
}
func ShowHangmansPositions(mistakes int) string{
	switch mistakes {
	case 0:
	fmt.Printf("  +---+\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("========\n")
	case 1:
	fmt.Printf("  +---+\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("  O   |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("========\n")
	case 2:
	fmt.Printf("  +---+\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("  O   |\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("========\n")
	case 3:
	fmt.Printf("  +---+\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("  O   |\n")
	fmt.Printf(" /|   |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("========\n")
	case 4:
	fmt.Printf("  +---+\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("  O   |\n")
	fmt.Printf("_/|   |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("========\n")
	case 5:
	fmt.Printf("  +---+\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("  O   |\n")
	fmt.Printf("_/|\\  |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("========\n")
	case 6:
	fmt.Printf("  +---+\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("  O   |\n")
	fmt.Printf("_/|\\_ |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("========\n")
	case 7:
	fmt.Printf("  +---+\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("  O   |\n")
	fmt.Printf("_/|\\_ |\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("========\n")
	case 8:
	fmt.Printf("  +---+\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("  O   |\n")
	fmt.Printf("_/|\\_ |\n")
	fmt.Printf("  |   |\n")
	fmt.Printf(" /    |\n")
	fmt.Printf("      |\n")
	fmt.Printf("      |\n")
	fmt.Printf("========\n")
	case 9:
	fmt.Printf("  +---+\n")
	fmt.Printf("  |   |\n")
	fmt.Printf("  O   |\n")
	fmt.Printf("_/|\\_ |\n")
	fmt.Printf("  |   |\n")
	fmt.Printf(" / \\  |\n")
	fmt.Printf("      |\n")
	fmt.Printf("R.I.P |\n")
	fmt.Printf("========\n")
	}
}

func main() {
    fmt.Println(wordToGuess())
    result := revealHint(wordToGuess())
    for _,v := range result{
        fmt.Printf(v)
    }
}
