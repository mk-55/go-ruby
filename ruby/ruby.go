package ruby

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"unicode"

	"github.com/ikawaha/kagome-dict/ipa"
	"github.com/ikawaha/kagome/v2/tokenizer"
)

// Configuration
type Configuration struct {
	Pattern *template.Template
}

// ルビ
// TODO命名がくそ
type Ruby struct {
	Parent string
	Ruby   string
}

// RubyWriter
type RubyWriter struct {
	config    *Configuration
	tokenizer *tokenizer.Tokenizer
}

// Configuration Constructor
func NewConfiguration(pattern string) (*Configuration, error) {

	err := validatePattern(pattern)
	if err != nil {
		return nil, err
	}

	t, err := template.New("pattern").Parse(pattern)
	if err != nil {
		return nil, err
	}

	return &Configuration{
		Pattern: t,
	}, nil
}

// Pattern validation
// Valid pattern contain {{.Parent}} and {{.Ruby}}
func validatePattern(pattern string) error {

	if strings.Contains(pattern, "{{.Parent}}") && strings.Contains(pattern, "{{.Ruby}}") {
		return nil
	}
	return fmt.Errorf("invalid pattern : %s", pattern)
}

// RubyWriter Constructor
func NewRubyWriter(config *Configuration) (*RubyWriter, error) {
	//TODO Validate　ARG
	//TODO tokenizer setting
	t, err := tokenizer.New(ipa.Dict(), tokenizer.OmitBosEos())
	if err != nil {
		return nil, err
	}

	return &RubyWriter{
		config:    config,
		tokenizer: t,
	}, nil
}

//引数の文字列にルビを追加する
func (rw *RubyWriter) AddRuby(str string) (string, error) {
	tokens := rw.tokenizer.Tokenize(str)

	rubinized_string := ""

	for _, token := range tokens {

		chunk := token.Surface
		if rw.isRubinizable(token) {
			c, err := rw.rubinizeToken(token)
			if err != nil {
				return "", err
			}

			chunk = c
		}
		rubinized_string += chunk
	}

	return rubinized_string, nil
}

//TODO引数の文字列からルビを除去する
func (rw *RubyWriter) RemoveRuby(string) (string, error) {
	return "", nil
}

//TODO引数の文字列のルビを除去してルビを降りなおす
func (rw *RubyWriter) MakeRuby(string) (string, error) {
	return "", nil
}

//ルビを振る対象か判定する
//ルビを振る条件
//1. Tokenに読み仮名が存在している
//2. Tokenをカタカナに変換して、読み仮名と一致しない
func (rw *RubyWriter) isRubinizable(token tokenizer.Token) bool {

	reading, _ := token.Reading()
	return reading != hiragana2katakana(token.Surface) && reading != ""
}

//tokenをルビ記法に変換する
func (rw *RubyWriter) rubinizeToken(token tokenizer.Token) (string, error) {

	surface := token.Surface
	reading, _ := token.Reading()
	//TODO 将来的にはルビのひらがな/カタカナの選択も設定で切り替えられるとよい
	return rw.rubinize(surface, katakana2hiragana(reading))
}

//ルビ記法を出力する
func (rw *RubyWriter) rubinize(parent string, ruby string) (string, error) {
	var msg_result bytes.Buffer

	replace_to := Ruby{
		Parent: parent,
		Ruby:   ruby,
	}

	err := rw.config.Pattern.Execute(&msg_result, replace_to)
	if err != nil {
		return "", err
	}

	return msg_result.String(), nil
}

const (
	//ひらがなとカタカナの文字コードの差
	KANADIFF          = 96
	RUBY_PATTERN_HTML = "<ruby>{{.Parent}}<rp>(</rp><rt>{{.Ruby}}</rt><rp>)</rp></ruby>"
)

//カタカナをひらがなに変換する
//https://qiita.com/yoheimuta/items/35c8dfb36cddfe19a64e
func katakana2hiragana(str string) string {
	src := []rune(str)
	dst := make([]rune, len(src))

	for i, r := range src {
		if unicode.In(r, unicode.Katakana) {
			dst[i] = r - KANADIFF
		} else {
			dst[i] = r
		}
	}

	return string(dst)
}

//stringに含まれるひらがなをカタカナに変換する
func hiragana2katakana(str string) string {
	src := []rune(str)
	dst := make([]rune, len(src))

	for i, r := range src {
		if unicode.In(r, unicode.Hiragana) {
			dst[i] = r + KANADIFF
		} else {
			dst[i] = r
		}
	}
	return string(dst)
}
