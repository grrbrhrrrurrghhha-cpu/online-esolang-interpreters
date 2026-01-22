package main

import (
  "time"
  "strconv"
  "github.com/gin-gonic/gin"
)

func ExecuteHQ9Plus(code string) string {
  var output string
  var acc, ops int
  start := time.Now()
  for i := 0; i < len(code); i++ {
    ops++
    if ops % 1000 == 0 {
      if time.Since(start) > 5 * time.Second {
        output += "\nTimed out\n"
        return output
      }
    }
    switch code[i] {
    case 'H':
      output += "Hello, world!\n"
    case 'Q':
      output += code + "\n"
    case '9':
      for j := 99; j > 1; j-- {
        output += strconv.Itoa(j) + " bottles of beer on the wall,\n" + strconv.Itoa(j) + " bottles of beer.\nTake one down, pass it around,\n"
        if j > 2 {
          output += strconv.Itoa(j - 1) + " bottles of beer on the wall.\n"
        } else {
          output += "1 bottle of beer on the wall.\n"
        }
      }
      output += "1 bottle of beer on the wall,\n1 bottle of beer.\nTake one down, pass it around,\nNo bottles of beer on the wall.\n"
    case '+':
      acc++
    }
  }
  return output
}

func ExecuteBrainfuck(code string, input string) string {
  var output string
  var tape = [30000]uint8{}
  var index, pointer, ops int
  start := time.Now()
  for i := 0; i < len(code); i++ {
    ops++
    if ops % 1000 == 0 {
      if time.Since(start) > 5 * time.Second {
        output += "\nTimed out\n"
        return output
      }
    }
    switch code[i] {
    case '>':
      pointer++
      if pointer > 29999 {
        pointer = 0
      }
    case '<':
      pointer--
      if pointer < 0 {
        pointer = 29999
      }
    case '+':
      tape[pointer]++
    case '-':
      tape[pointer]--
    case '.':
      output += string(tape[pointer])
    case ',':
      if index < len(input) {
        tape[pointer] = uint8(input[index])
        index++
      } else {
        tape[pointer] = 0
      }
    case '[':
      if tape[pointer] == 0 {
        balance := 1
        for j := i + 1; j < len(code); j++ {
          switch code[j] {
          case '[':
            balance++
          case ']':
            balance--
          }
          if balance == 0 {
            i = j
            break
          }
        }
      }
    case ']':
      if tape[pointer] != 0 {
        balance := 1
        for j := i - 1; j >= 0; j-- {
          switch code[j] {
          case '[':
            balance--
          case ']':
            balance++
          }
          if balance == 0 {
            i = j
            break
          }
        }
      }
    }
  }
  return output
}

func ExecuteDeadfish(code string) string {
  var output string
  var acc, ops int
  start := time.Now()
  for i := 0; i < len(code); i++ {
    ops++
    if ops % 1000 == 0 {
      if time.Since(start) > 5 * time.Second {
        output += "\nTimed out\n"
        return output
      }
    }
    switch code[i] {
    case 'i':
      acc++
      if acc == 256 || acc == -1 {
        acc = 0
      }
    case 'd':
      acc--
      if acc == 256 || acc == -1 {
        acc = 0
      }
    case 's':
      acc *= acc
      if acc == 256 || acc == -1 {
        acc = 0
      }
    case 'o':
      output += strconv.Itoa(acc)
    }
  }
  return output
}

func main() {
  router := gin.Default()
  router.SetTrustedProxies(nil)
  router.LoadHTMLGlob("templates/*")
  router.Static("/assets", "./assets")
  router.POST("/", func(c *gin.Context) {
    lang := c.PostForm("lang")
    code := c.PostForm("code")
    input := c.PostForm("input")
    var output string
    switch lang {
    case "hq9plus":
      output = ExecuteHQ9Plus(code)
    case "brainfuck":
      output = ExecuteBrainfuck(code, input)
    case "deadfish":
      output = ExecuteDeadfish(code)
    default:
      output = "Unknown esolang: " + lang
    }
    c.HTML(200, "index.html", gin.H{
      "output": output,
      "code": code,
      "input": input,
      "lang": lang,
    })
  })
  router.GET("/", func(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{
      "output": "",
      "input": "",
      "code": "",
    })
  })
  router.Run(":4269")
}
