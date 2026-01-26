package main

import (
  "os"
  "strconv"
  "strings"
  "time"
  "github.com/gin-gonic/gin"
)

func ExecuteHQ9Plus(code string) string {
  var output string
  var acc, ops int
  start := time.Now()
  for _, instruction := range code {
    ops++
    if ops % 1000 == 0 {
      if time.Since(start) > 5 * time.Second {
        output += "\nTimed out\n"
        return output
      }
    }
    switch instruction {
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
  for _, instruction := range code {
    ops++
    if ops % 1000 == 0 {
      if time.Since(start) > 5 * time.Second {
        output += "\nTimed out\n"
        return output
      }
    }
    switch instruction {
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

func ExecuteSubleq(code string, input string) string {
  var ip, ops, index int
  var output string
  memory := [30000]int{}
  start := time.Now()
  fields := strings.Fields(code)
  for i, field := range fields {
    num, err := strconv.Atoi(field)
    if err != nil {
      continue
    }
    if i < len(memory) {
      memory[i] = num
    } else {
      break
    }
  }
  for ip >= 0 && ip + 2 < len(memory) {
    ops++
    if ops % 1000 == 0 {
      if time.Since(start) > 5 * time.Second {
        output += "\nTimed out\n"
        return output
      }
    }
    a := memory[ip]
    b := memory[ip + 1]
    c := memory[ip + 2]
    if a >= 0 && b >= 0 {
      memory[b] = memory[b] - memory[a]
      if memory[b] <= 0 {
        ip = c
      } else {
        ip += 3
      }
    } else if a == -1 && b >= 0 {
      if index < len(input) {
        memory[b] = int(input[index])
        index++
      } else {
        memory[b] = 0
      }
      ip += 3
    } else if a >= 0 && b == -1 {
      output += string(byte(memory[a]))
      ip += 3
    } else {
      break
    }
  }
  return output
}

func ExecuteRPN(code string, input string) string {
  stack := []int{}
  inputs := []int{}
  var output string
  inputFields := strings.Fields(input)
  for _, in := range inputFields {
    num, err := strconv.Atoi(in)
    if err != nil {
      continue
    }
    inputs = append(inputs, num)
  }
  fields := strings.Fields(code)
  for _, field := range fields {
    num, err := strconv.Atoi(field)
    if err != nil {
      if len(stack) > 1 {
        b := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        a := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        switch field {
        case "+":
          stack = append(stack, a + b)
        case "-":
          stack = append(stack, a - b)
        case "/":
          stack = append(stack, a / b)
        case "*":
          stack = append(stack, a * b)
        case "%":
          stack = append(stack, a % b)
        }
      }
      if len(field) > 1 && strings.HasPrefix(field, "$") {
        index, err := strconv.Atoi(field[1:])
        if err == nil && index < len(inputs) {
          stack = append(stack, inputs[index])
        }
      }
      continue
    }
    stack = append(stack, num)
  }
  if len(stack) > 0 {
    output = strconv.Itoa(stack[0])
  }
  return output
}

func main() {
  var count int
  content, err := os.ReadFile("count")
  if err == nil {
    count, err = strconv.Atoi(string(content))
    if err != nil {
      count = 0
    }
  }
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
    case "subleq":
      output = ExecuteSubleq(code, input)
    case "rpn":
      output = ExecuteRPN(code, input)
    case "text":
      output = code
    default:
      output = "Unknown esolang: " + lang
    }
    c.HTML(200, "index.html", gin.H{
      "output": output,
      "code": code,
      "input": input,
      "lang": lang,
      "count": count,
    })
  })
  router.GET("/", func(c *gin.Context) {
    count++
    _ = os.WriteFile("count", []byte(strconv.Itoa(count)), 600)
    c.HTML(200, "index.html", gin.H{
      "output": "",
      "input": "",
      "code": "+[-->-[>>+>-----<<]<--<---]>-.>>>+.>>..+++[.>]<<<<.+++.------.<<-.>>>>+.",
      "lang": "brainfuck",
      "count": count,
    })
  })
  router.Run(":4269")
}
