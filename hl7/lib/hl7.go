package hl7

const (
    LEVEL_0 int = 0
    SEGMENT_SEPARATOR = LEVEL_0
    LEVEL_1 = 1
    FIELD_SEPARATOR = LEVEL_1
    LEVEL_2 = 2
    COMPONENT_SEPARATOR = LEVEL_2
    LEVEL_3 = 3
    SUB_COMPONENT_SEPARATOR = LEVEL_3
    LEVEL_4 = 4
    REPEAT_SEPARATOR = -1
    ESCAPE_CHARACTER = -2
    ENCODING_FIELD_IDX = 0
    ENCODING_COMPONENT_IDX = 1
    ENCODING_REPEAT_IDX = 2
    ENCODING_ESCAPE_IDX = 3
    ENCODING_SUB_COMPONENT_IDX = 4
)

type Delimiters map[int]string

var defaultEncoding = "|^~\\&"

const (
    CARRIAGE_RETURN_SEQ string = ".br"
    LINE_FEED_SEQ = "X0D"
    FIELD_SEPARATOR_SEQ = "F"
    COMPONENT_SEPARATOR_SEQ = "S"
    SUB_COMPONENT_SEPARATOR_SEQ = "T"
    REPEAT_SEPARATOR_SEQ = "R"
    ESCAPE_CHARACTER_SEQ = "E"
)

type EscapeSequences map[string]string

func newDelimitersAndEscapeSequences(encoding []byte) (Delimiters, EscapeSequences) {
    delimiters := Delimiters{
                 LEVEL_0 : "\x0D",
                 LEVEL_1 : string([]byte{encoding[ENCODING_FIELD_IDX]}),
                 LEVEL_2 : string([]byte{encoding[ENCODING_COMPONENT_IDX]}),
                 LEVEL_3 : string([]byte{encoding[ENCODING_SUB_COMPONENT_IDX]}),
                 REPEAT_SEPARATOR : string([]byte{encoding[ENCODING_REPEAT_IDX]}),
                 ESCAPE_CHARACTER : string([]byte{encoding[ENCODING_ESCAPE_IDX]}),
            }
    escape := delimiters[ESCAPE_CHARACTER]
    escapeSequences := EscapeSequences{
                 escape + CARRIAGE_RETURN_SEQ + escape : "\x0D",
                 escape + LINE_FEED_SEQ + escape : "\x0A",
                 escape + FIELD_SEPARATOR_SEQ + escape : delimiters[FIELD_SEPARATOR],
                 escape + COMPONENT_SEPARATOR_SEQ + escape : delimiters[COMPONENT_SEPARATOR],
                 escape + SUB_COMPONENT_SEPARATOR_SEQ + escape : delimiters[SUB_COMPONENT_SEPARATOR],
                 escape + REPEAT_SEPARATOR_SEQ + escape : delimiters[REPEAT_SEPARATOR],
                 escape + ESCAPE_CHARACTER_SEQ + escape : delimiters[ESCAPE_CHARACTER],
            }
    return delimiters, escapeSequences
}
