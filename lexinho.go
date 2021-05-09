package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	token_error                = -1 // ERRO
	token_undefined            = 0  // UNDEFINED
	token_word                 = 1  // WORD
	token_floating_number      = 2  // 6.6
	token_integer_number       = 3  // 6
	token_void                 = 4  // void
	token_int                  = 5  // int
	token_float                = 6  // float
	token_double               = 7  // double
	token_char                 = 8  // char
	token_for                  = 9  // for
	token_while                = 10 // while
	token_do                   = 11 // do
	token_if                   = 12 // if
	token_else                 = 13 // else
	token_switch               = 14 // switch
	token_case                 = 15 // case
	token_default              = 16 // default
	token_continue             = 17 // continue
	token_break                = 18 // break
	token_plus                 = 19 // +
	token_minus                = 20 // -
	token_product              = 21 // *
	token_division             = 22 // /
	token_module               = 23 // %
	token_greater              = 24 // >
	token_less                 = 25 // <
	token_equals               = 26 // ==
	token_not_equals           = 27 // !=
	token_greater_or_equals    = 28 // >=
	token_less_or_equals       = 29 // <=
	token_bitwise_and          = 30 // &
	token_bitwise_or           = 31 // |
	token_bitwise_xor          = 32 // ^
	token_bitwise_rigth        = 33 // >>
	token_bitwise_left         = 34 // <<
	token_bitwise_invert       = 35 // ~
	token_logical_and          = 36 // &&
	token_logical_or           = 37 // ||
	token_logical_denied       = 38 // !
	token_assignment           = 39 // =
	token_assignment_increment = 40 // +=
	token_assignment_decrement = 42 // -=
	token_assignment_multiple  = 43 // *=
	token_assignment_divide    = 44 // /=
	token_assignment_module    = 45 // %=
	token_assignment_right     = 46 // >>=
	token_assignment_left      = 47 // <<=
	token_increment            = 49 // ++
	token_decrement            = 50 // --
	token_open_parentheses     = 51 // (
	token_close_parentheses    = 52 // )
	token_open_brace           = 53 // {
	token_close_brace          = 54 // }
	token_open_bracket         = 55 // [
	token_close_bracket        = 56 // ]
	token_comma                = 57 // ,
	token_point                = 58 // .
	token_semicolon            = 59 // ;
) //tokens

const (
	initial = 0
	word    = 1
	number  = 2
	end     = 3
) //states

var keyWords = map[string]int{
	"void":     token_void,
	"int":      token_int,
	"float":    token_float,
	"double":   token_double,
	"char":     token_char,
	"for":      token_for,
	"while":    token_while,
	"do":       token_do,
	"if":       token_if,
	"else":     token_else,
	"switch":   token_switch,
	"case":     token_case,
	"default":  token_default,
	"continue": token_continue,
	"break":    token_break,
}

var _char rune
var _err error

var _reader *bufio.Reader
var _writer *bufio.Writer

var _line int = 0
var _column int = 0

func undo_read() {
	_reader.UnreadRune()
	_column--
}

func undo_last_lex(lex string) string {
	return lex[:len(lex)-1]
}

func undo_lex_size(lex string) {
	_column -= len(lex)
}

func read() {
	if _char, _, _err = _reader.ReadRune(); _err != nil {
		if _err != io.EOF {
			fmt.Printf("Erro na leitura de caracter :( \n %s\n", _err.Error())
		}
	} else {
		_column++
	}
}

func charIsLetter() bool {
	return _char >= 'a' && _char <= 'z' ||
		_char >= 'A' && _char <= 'Z' ||
		_char == '_'
}

func charIsNumber() bool {
	return _char >= '0' && _char <= '9'
}

func getToken(state int, lex string) (int, int, string) {
	lex += string(_char)

	switch state {
	case initial:
		if _char == ' ' || _char == '\t' {
			read()
			return token_undefined, initial, ""
		}
		if _char == '\r' {
			read()

			if _char == '\n' {
				read()
				_line++
				_column = 0
				return token_undefined, initial, ""
			}

			return token_error, initial, ""
		}
		if _char == '\n' {
			read()
			_line++
			_column = 0
			return token_undefined, initial, ""
		}
		if charIsLetter() {
			read()
			return token_undefined, word, lex
		}
		if charIsNumber() {
			read()
			return token_undefined, number, lex
		}
		if _char == '+' {
			read()

			if _char == '+' {
				return token_increment, initial, lex + string(_char)
			}
			if _char == '=' {
				return token_assignment_increment, initial, lex + string(_char)
			}

			undo_read()
			return token_plus, initial, lex
		}
		if _char == '-' {
			read()

			if _char == '-' {
				return token_decrement, initial, lex + string(_char)
			}
			if _char == '=' {
				return token_assignment_decrement, initial, lex + string(_char)
			}

			undo_read()
			return token_minus, initial, lex
		}
		if _char == '*' {
			read()

			if _char == '=' {
				return token_assignment_multiple, initial, lex + string(_char)
			}

			undo_read()
			return token_product, initial, lex
		}
		if _char == '/' {
			read()

			if _char == '=' {
				return token_assignment_divide, initial, lex + string(_char)
			}

			undo_read()
			return token_division, initial, lex
		}
		if _char == '%' {
			read()

			if _char == '=' {
				return token_assignment_module, initial, lex + string(_char)
			}

			undo_read()
			return token_module, initial, lex
		}
		if _char == '>' {
			read()

			if _char == '=' {
				return token_greater_or_equals, initial, lex + string(_char)
			}

			if _char == '>' {
				lex += string(_char)

				read()

				if _char == '=' {
					return token_assignment_right, initial, lex + string(_char)
				}

				undo_read()
				return token_bitwise_rigth, initial, lex
			}

			undo_read()
			return token_greater, initial, lex
		}
		if _char == '<' {
			read()

			if _char == '=' {
				return token_less_or_equals, initial, lex + string(_char)
			}

			if _char == '<' {
				lex += string(_char)

				read()

				if _char == '=' {
					return token_assignment_left, initial, lex + string(_char)
				}

				undo_read()
				return token_bitwise_left, initial, lex
			}

			undo_read()
			return token_less, initial, lex
		}
		if _char == '=' {
			read()

			if _char == '=' {
				return token_equals, initial, lex + string(_char)
			}

			undo_read()
			return token_assignment, initial, lex
		}
		if _char == '!' {
			read()

			if _char == '=' {
				return token_not_equals, initial, lex + string(_char)
			}

			undo_read()
			return token_logical_denied, initial, lex
		}
		if _char == '&' {
			read()

			if _char == '&' {
				return token_logical_and, initial, lex + string(_char)
			}

			undo_read()
			return token_bitwise_and, initial, lex
		}
		if _char == '|' {
			read()

			if _char == '|' {
				return token_logical_or, initial, lex + string(_char)
			}

			undo_read()
			return token_bitwise_or, initial, lex
		}
		if _char == '|' {
			read()

			if _char == '|' {
				return token_logical_or, initial, lex + string(_char)
			}

			undo_read()
			return token_bitwise_or, initial, lex
		}
		if _char == '^' {
			return token_bitwise_xor, initial, lex
		}
		if _char == '~' {
			return token_bitwise_invert, initial, lex
		}
		if _char == '(' {
			return token_open_parentheses, initial, lex
		}
		if _char == ')' {
			return token_close_parentheses, initial, lex
		}
		if _char == '{' {
			return token_open_brace, initial, lex
		}
		if _char == '}' {
			return token_close_brace, initial, lex
		}
		if _char == '[' {
			return token_open_bracket, initial, lex
		}
		if _char == ']' {
			return token_close_bracket, initial, lex
		}
		if _char == ',' {
			return token_comma, initial, lex
		}
		if _char == '.' {
			return token_point, initial, lex
		}
		if _char == ';' {
			return token_semicolon, initial, lex
		}
		if _err == io.EOF {
			return token_undefined, end, ""
		} else {
			return token_error, initial, lex
		}
	case word:
		if charIsLetter() || charIsNumber() {
			read()
			return token_undefined, word, lex
		} else {
			undo_read()
			lex = undo_last_lex(lex)
			undo_lex_size(lex)

			if keyWork, ok := keyWords[lex]; ok {
				return keyWork, initial, lex
			} else {
				return token_word, initial, lex
			}
		}
	case number:
		if charIsNumber() {
			read()
			return token_undefined, number, lex
		} else if _char == '.' {
			read()
			return token_undefined, number, lex
		} else {
			undo_read()
			lex = undo_last_lex(lex)
			undo_lex_size(lex)

			if _, err := strconv.ParseInt(lex, 10, 64); err == nil {
				return token_integer_number, initial, lex
			} else if _, err := strconv.ParseFloat(lex, 64); err == nil {
				return token_floating_number, initial, lex
			} else {
				return token_error, initial, lex
			}
		}
	}

	return token_undefined, initial, ""
}

func analyse() {
	read()

	tk, state, lex := getToken(initial, "")

	for ; state != end; tk, state, lex = getToken(state, lex) {
		switch tk {
		case token_undefined:
			continue
		case token_error:
			output := fmt.Sprintf("Erro léxico: encontrou o caracter %c (linha: %d coluna: %d)\n", _char, _line, _column)

			fmt.Printf(output)
			_writer.WriteString(output)

			return
		default:
			output := fmt.Sprintf("%d %s (linha: %d coluna: %d)\n", tk, lex, _line, _column)

			fmt.Printf(output)
			_writer.WriteString(output)

			lex = ""
			read()
		}
	}
}

func main() {
	filePath := "teste.c"
	outputPath := "lexinho.txt"

	if inputFile, err := os.Open(filePath); err != nil {
		fmt.Println("Caminho do arquivo de entrada é inválido, não foi possível abrir o arquivo.")
		fmt.Println(err.Error())
	} else if outputFile, err := os.Create(outputPath); err != nil {
		fmt.Println("Caminho do arquivo de saída é inválido, não foi possível criar o arquivo.")
		fmt.Println(err.Error())
	} else {
		defer inputFile.Close()
		defer outputFile.Close()

		_reader = bufio.NewReader(inputFile)
		_writer = bufio.NewWriter(outputFile)
		defer _writer.Flush()

		analyse()
	}
}
