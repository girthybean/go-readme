package main

import (
    "fmt"
    "strings"
    "reflect"
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
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

    lines := textToLines(text)

    for _, line := range lines {
        tmp := strings.Split(line, ":")
        lineIndexString := strings.Split(tmp[0], " ")

        lineIndexInt, err := strconv.Atoi(lineIndexString[1])
        if err != nil {
            // ... handle error
            panic(err)
        }

        fmt.Println(lineIndexInt,reflect.TypeOf(lineIndexInt))
    }
}
