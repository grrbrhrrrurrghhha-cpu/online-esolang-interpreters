package main

import (
  "math/rand"
  "os"
  "strconv"
  "strings"
  "time"
  "github.com/gin-gonic/gin"
)

func mod(a int, b int) int {
  return ((a % b) + b) % b
}

func push(stack *[]int, a int) {
  *stack = append(*stack, a)
}

func pop(stack *[]int) int {
  if len(*stack) > 0 {
    top := (*stack)[len(*stack) - 1]
    *stack = (*stack)[:len(*stack) - 1]
    return top
  } else {
    return 0
  }
}

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

func ExecutePercentCaretAnd78(code string) string {
  var output string
  var grid = [256][256]bool{}
  var ops int
  var x, y uint8
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
      x++
    case '<':
      x--
    case 'v':
      y++
    case '^':
      y--
    case '@':
      grid[y][x] = !grid[y][x]
    case '{':
      if !grid[y][x] {
        balance := 1
        for j := i + 1; j < len(code); j++ {
          switch code[j] {
          case '{':
            balance++
          case '}':
            balance--
          }
          if balance == 0 {
            i = j
            break
          }
        }
      }
    case '}':
      if grid[y][x] {
        balance := 1
        for j := i - 1; j >= 0; j-- {
          switch code[j] {
          case '{':
            balance--
          case '}':
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
  for i := 0; i < 256; i++ {
    for j := 0; j < 256; j++ {
      if grid[i][j] {
        output += "1"
      } else {
        output += "0"
      }
    }
    output += "\n"
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
    case 'd':
      acc--
    case 's':
      acc *= acc
    case 'o':
      output += strconv.Itoa(acc)
    }
    if acc == 256 || acc == -1 {
      acc = 0
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
    if a >= 0 && b >= 0 && a < 30000 && b < 30000 {
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
          stack = append(stack, mod(a, b))
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
    output = strconv.Itoa(stack[len(stack) - 1])
  }
  return output
}

func ExecuteCuteCats(code string) string {
  var output string
  var acc, ops int
  start := time.Now()
  for _, c := range code {
    ops++
    if ops % 1000 == 0 {
      if time.Since(start) > 5 * time.Second {
        output += "\nTimed out\n"
        return output
      }
    }
    switch c {
     case '🐱':
       acc = 2
     case '🐈':
       if acc == 2 {
         output += "4"
        } else {
          output += "31"
        }
    }
  }

  return output
}

func Execute67machine(code string) string {
  var output string
  var ip, ops int
  codeArray := []rune(code)
  start := time.Now()
  for {
    ops++
    if ops % 1000 == 0 {
      if time.Since(start) > 5 * time.Second {
        output += "\nTimed out\n"
        return output
      }
    }
    output += string(codeArray)
    output += "\n"
    c := codeArray[ip]
    switch c {
    case '6':
      if codeArray[mod(ip + 1, len(codeArray))] == '6' {
        codeArray[mod(ip + 1, len(codeArray))] = '7'
      } else {
        codeArray[mod(ip + 1, len(codeArray))] = '6'
      }
      ip = mod(ip + 7, len(codeArray))
    case '7':
      codeArray = append(codeArray, codeArray[mod(ip + 1, len(codeArray))])
      ip = mod(ip - 6, len(codeArray))
    }
  }

  return output
}

func ExecuteBefunge93(code string, input string) string {
  var output string
  var index, x, y, dx, dy, ops int
  var stringMode bool
  width := 1
  height := 1
  running := true
  codeGrid := [25][80]rune{}
  stack := []int{}
  start := time.Now()
  
  for _, c := range code {
    if c != '\n' {
      codeGrid[y][x] = c
      x++
      if x > width {
        width = x
      }
      if x >= 80 {
        return "Out of bounds\n"
      }
    } else {
      y++
      if y + 1 > height {
        height = y + 1
      }
      x = 0
      if y >= 25 {
        return "Out of bounds\n"
      }
    }
  }
  x = 0
  y = 0
  
  dx = 1
  dy = 0
  for running {
    ops++
    if ops % 1000 == 0 {
      if time.Since(start) > 5 * time.Second {
        output += "\nTimed out\n"
        return output
      }
    }
    
    if stringMode && codeGrid[y][x] != '"' {
      push(&stack, int(codeGrid[y][x]))
      x = mod(x + dx, width)
      y = mod(y + dy, height)
      continue
    }
    
    switch codeGrid[y][x] {
      case '+':
        a := pop(&stack)
        b := pop(&stack)
        push(&stack, a + b)
      case '-':
        a := pop(&stack)
        b := pop(&stack)
        push(&stack, b - a)
      case '*':
        a := pop(&stack)
        b := pop(&stack)
        push(&stack, a * b)
      case '/':
        a := pop(&stack)
        b := pop(&stack)
        if a == 0 {
          var negative bool
          for index < len(input) && (input[index] == ' ' || input[index] == '\n' || input[index] == '\r') {
            index++
          }
          if index < len(input) && input[index] == '-' {
            negative = true
            index++
          }
          for index < len(input) && (input[index] >= '0' && input[index] <= '9') {
            a *= 10
            a += int(input[index] - '0')
            index++
          }
          if negative {
            a = -a
          }
        }
        push(&stack, b / a)
      case '%':
        a := pop(&stack)
        b := pop(&stack)
        push(&stack, b % a)
      case '!':
        a := pop(&stack)
        if a == 0 {
          push(&stack, 1)
        } else {
          push(&stack, 0)
        }
      case '`':
        a := pop(&stack)
        b := pop(&stack)
        if b > a {
          push(&stack, 1)
        } else {
          push(&stack, 0)
        }
      case '>':
        dx = 1
        dy = 0
      case '<':
        dx = -1
        dy = 0
      case '^':
        dy = -1
        dx = 0
      case 'v':
        dy = 1
        dx = 0
      case '?':
        dir := rand.Intn(4)
        switch dir {
          case 0:
            dx = 1
            dy = 0
          case 1:
            dx = -1
            dy = 0
          case 2:
            dy = -1
            dx = 0
          case 3:
            dy = 1
            dx = 0
        }
      case '_':
        a := pop(&stack)
        if a == 0 {
          dx = 1
        } else {
          dx = -1
        }
        dy = 0
      case '|':
        a := pop(&stack)
        if a == 0 {
          dy = 1
        } else {
          dy = -1
        }
        dx = 0
      case '"':
        stringMode = !stringMode
      case ':':
        a := pop(&stack)
        push(&stack, a)
        push(&stack, a)
      case '\\':
        a := pop(&stack)
        b := pop(&stack)
        push(&stack, a)
        push(&stack, b)
      case '$':
        pop(&stack)
      case '.':
        a := pop(&stack)
        output += strconv.Itoa(a) + " "
      case ',':
        a := pop(&stack)
        output += string(a)
      case '#':
        x = mod(x + dx, width)
        y = mod(y + dy, height)
      case 'g':
        y := pop(&stack)
        x := pop(&stack)
        if x < 80 && x >= 0 && y < 25 && y >= 0 {
          push(&stack, int(codeGrid[y][x]))
        } else {
          push(&stack, 0)
        }
      case 'p':
        y := pop(&stack)
        x := pop(&stack)
        v := pop(&stack)
        if x < 80 && x >= 0 && y < 25 && y >= 0 {
          codeGrid[y][x] = rune(v)
        }
      case '&':
        var negative bool
        n := 0
        for index < len(input) && (input[index] == ' ' || input[index] == '\n' || input[index] == '\r') {
          index++
        }
        if index < len(input) && input[index] == '-' {
          negative = true
          index++
        }
        for index < len(input) && (input[index] >= '0' && input[index] <= '9') {
          n *= 10
          n += int(input[index] - '0')
          index++
        }
        if negative {
          n = -n
        }
        push(&stack, n)
      case '~':
        if index < len(input) {
          push(&stack, int(input[index]))
          index++
        }
      case '@':
        running = false
      case '0':
         push(&stack, 0)
      case '1':
         push(&stack, 1)
      case '2':
         push(&stack, 2)
      case '3':
         push(&stack, 3)
      case '4':
         push(&stack, 4)
      case '5':
         push(&stack, 5)
      case '6':
         push(&stack, 6)
      case '7':
         push(&stack, 7)
      case '8':
         push(&stack, 8)
      case '9':
        push(&stack, 9)
    }
    
    x = mod(x + dx, width)
    y = mod(y + dy, height)
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
    case "cutecats":
      output = ExecuteCuteCats(code)
    case "67machine":
      output = Execute67machine(code)
    case "percentcaretand78":
      output = ExecutePercentCaretAnd78(code)
    case "befunge93":
      output = ExecuteBefunge93(code, input)
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
