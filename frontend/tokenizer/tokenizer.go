package tokenizer

import (
  "strconv"
)

func Tokenize(rawData []byte) ([]Token) {
  // Creates an empty Slice of Tokens, that will hold everything
  tokens := make([]Token, 0)

  // Turns the byte-Array to a string, so its easier to work with
  data := string(rawData)

  line := 1
  lastEnd := 0
  // Goes through the entire Content
  for i, tmp := range data {
    // If it's not the end of a token, simply skip it
    if !(tmp == ' ' || tmp == ';' || tmp == '(' || tmp == ')' || tmp == '\n') {
      continue
    }

    // Check if the Token is not empty and is not ended because of a Line-Break
    if (i - lastEnd) != 0 && tmp != '\n' {
      // Get the content of the Token
      part := data[lastEnd:i]

      // Check if the content matches any known keywords
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
        // If it does not match any known Tokens, its either an ID or a Value
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

    // Checks for any "key-Symbols" that are needed for the Syntax-related stuff
    if tmp == ';' {
      tokens = append(tokens, Token{
        Type: "SEMICOLON",
        Content: ";",
        Line: line,
      })
    } else if tmp == '(' {
      tokens = append(tokens, Token{
        Type: "LPAREN",
        Content: "(",
        Line: line,
      })
    } else if tmp == ')' {
      tokens = append(tokens, Token{
        Type: "RPAREN",
        Content: ")",
        Line: line,
      })
    } else if tmp == '\n' {
      line++
    }

    lastEnd = i + 1
  }

  // Returns the slice of Tokens at the End
  return tokens
}
