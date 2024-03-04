package main

import (
    "fmt"
    "strings"
    //"reflect"
    "strconv"
)

func textToLines(text string) []string {
    var lines []string
    var line string
    for _, c := range text {
        if c == '\n' {
            lines = append(lines, line)
            line = ""
        } else {
            line += string(c)
        }
    }
    return lines
}

func main() {

    const text string = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`
    lines := textToLines(text)

    res:= 0

    for gameIndex, line := range lines {

        tmp := strings.Split(line, ":")
        bagsList := strings.TrimSpace(tmp[1])
        bags := strings.Split(bagsList, ";")

        test := true
        for _, bag := range bags {
            bag = strings.TrimSpace(bag)
            for _, cube := range strings.Split(bag, ",") {
                cube = strings.TrimSpace(cube)
                tmp2 := strings.Split(cube, " ")
                cubeCount, _ := strconv.Atoi(tmp2[0])
                cubeColor := tmp2[1]

                if cubeColor == "red" && cubeCount > 12 {
                    test = test && false
                }
                if cubeColor == "blue" && cubeCount > 14 {
                    test = test && false
                }
                if cubeColor == "green" && cubeCount > 13 {
                    test = test && false
                }
            }
        }
        if test {
            fmt.Println("Game", gameIndex+1, "is valid")
            res+= gameIndex+1
        }
    }

    fmt.Println("Result:", res)
}
