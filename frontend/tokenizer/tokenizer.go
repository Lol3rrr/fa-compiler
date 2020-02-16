package tokenizer

import (
  "strconv"
)

func Tokenize(rawData []byte) ([]Token) {
  tokens := make([]Token, 0)

  data := string(rawData)

  line := 1
  lastEnd := 0
  for i, tmp := range data {
    if tmp == ' ' || tmp == ';' || tmp == '(' || tmp == ')' || tmp == '\n' {
      if (i - lastEnd) != 0 && tmp != '\n' {
        part := data[lastEnd:i]

        if part == "number" {
          tokens = append(tokens, Token{
            Type: "DATATYPE",
            Content: "int",
            Line: line,
          })
        } else if part == "=" {
          tokens = append(tokens, Token{
            Type: "EQUALS",
            Content: part,
            Line: line,
          })
        } else if part == "+" {
          tokens = append(tokens, Token{
            Type: "ADDITION",
            Content: part,
            Line: line,
          })
        } else if part == "-" {
          tokens = append(tokens, Token{
            Type: "SUBTRACTION",
            Content: part,
            Line: line,
          })
        }else {
          _, err := strconv.Atoi(part)
          if err == nil {
            tokens = append(tokens, Token{
              Type: "INT",
              Content: part,
              Line: line,
            })
          }else {
            tokens = append(tokens, Token{
              Type: "ID",
              Content: part,
              Line: line,
            })
          }
        }
      }

      if tmp == ';' {
        tokens = append(tokens, Token{
          Type: "SEMICOLON",
          Content: ";",
          Line: line,
        })
      }

      if tmp == '(' {
        tokens = append(tokens, Token{
          Type: "LPAREN",
          Content: "(",
          Line: line,
        })
      }
      if tmp == ')' {
        tokens = append(tokens, Token{
          Type: "RPAREN",
          Content: ")",
          Line: line,
        })
      }
      if tmp == '\n' {
        line++
      }

      lastEnd = i + 1
    }
  }

  return tokens
}
