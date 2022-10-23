package ruby

import (
	"fmt"
	"testing"
)

func TestValidatePattern_Valid(t *testing.T) {
	valid_pattern := "{{.Parent}} | {{.Ruby}}"
	err := validatePattern(valid_pattern)
	if err != nil {
		t.Error(`wants no error but error:`, err)
	}
}

func TestValidatePattern_NotIncludeRuby(t *testing.T) {
	invalid_pattern := "{{.Parent}} | Ruby"
	err := validatePattern(invalid_pattern)
	if err == nil {
		t.Error(`wants error but no error`)
	}
}

func TestValidatePattern_NoMatch(t *testing.T) {
	invalid_pattern := "RUBY_PATTERN_HTML"
	err := validatePattern(invalid_pattern)
	if err == nil {
		t.Error(`wants error but no error`)
	}
}

func TestValidatePattern_NotIncludeParent(t *testing.T) {
	invalid_pattern := "Parent | {{.Ruby}}"
	err := validatePattern(invalid_pattern)
	if err == nil {
		t.Error(`wants error but no error`)
	}
}

func TestRubyWriter_rubinize(t *testing.T) {
	parent := "日本語"
	ruby := "にほんご"
	wants := "<ruby>日本語<rp>(</rp><rt>にほんご</rt><rp>)</rp></ruby>"

	conf, err := NewConfiguration(RUBY_PATTERN_HTML)
	if err != nil {
		t.Error(`miss`)
	}
	writer, err := NewRubyWriter(conf)
	if err != nil {
		t.Error(`miss`)
	}

	output, err := writer.rubinize(parent, ruby)
	if err != nil {
		t.Error(`miss`)
	}

	if output != wants {
		t.Error(fmt.Printf("wants %s but %s", wants, output))
	}
}

func Test_hiragana2katakana(t *testing.T) {
	hiragana := "あいうえおわおん"
	katakana := "アイウエオワオン"

	output := hiragana2katakana(hiragana)

	if output != katakana {
		t.Error(fmt.Printf("wants %s but %s", katakana, output))
	}
}

func Test_katakana2hiragana(t *testing.T) {
	hiragana := "あいうえおわおん"
	katakana := "アイウエオワオン"

	output := katakana2hiragana(katakana)

	if output != hiragana {
		t.Error(fmt.Printf("wants %s but %s", hiragana, output))
	}
}

func TestAddRuby_Normal(t *testing.T) {
	src := "吾輩は猫である"
	wants := "<ruby>吾輩<rp>(</rp><rt>わがはい</rt><rp>)</rp></ruby>は<ruby>猫<rp>(</rp><rt>ねこ</rt><rp>)</rp></ruby>である"

	conf, err := NewConfiguration(RUBY_PATTERN_HTML)
	if err != nil {
		t.Error(`miss`)
	}
	writer, err := NewRubyWriter(conf)
	if err != nil {
		t.Error(`miss`)
	}

	output, err := writer.AddRuby(src)
	if err != nil {
		t.Error(`miss`)
	}

	if output != wants {
		t.Error(fmt.Printf("wants %s but %s", wants, output))
	}
}

func TestAddRuby_English(t *testing.T) {
	src := "吾輩はCatである"
	wants := "<ruby>吾輩<rp>(</rp><rt>わがはい</rt><rp>)</rp></ruby>はCatである"

	conf, err := NewConfiguration(RUBY_PATTERN_HTML)
	if err != nil {
		t.Error(`miss`)
	}
	writer, err := NewRubyWriter(conf)
	if err != nil {
		t.Error(`miss`)
	}

	output, err := writer.AddRuby(src)
	if err != nil {
		t.Error(`miss`)
	}

	if output != wants {
		t.Error(fmt.Printf("wants %s but %s", wants, output))
	}
}

func TestAddRuby_MultiLine(t *testing.T) {
	src := `吾輩は猫である。
	名前はまだない。
	`
	wants := `<ruby>吾輩<rp>(</rp><rt>わがはい</rt><rp>)</rp></ruby>は<ruby>猫<rp>(</rp><rt>ねこ</rt><rp>)</rp></ruby>である。
	<ruby>名前<rp>(</rp><rt>なまえ</rt><rp>)</rp></ruby>はまだない。
	`

	conf, err := NewConfiguration(RUBY_PATTERN_HTML)
	if err != nil {
		t.Error(`miss`)
	}
	writer, err := NewRubyWriter(conf)
	if err != nil {
		t.Error(`miss`)
	}

	output, err := writer.AddRuby(src)
	if err != nil {
		t.Error(`miss`)
	}

	if output != wants {
		t.Error(fmt.Printf("wants %s but %s", wants, output))
	}
}

//TODO Markdownのパターン追加(Pre記法はルビを振らない、とか)
func TestAddRuby_Markdown(t *testing.T) {
	src := `# タイトル1
	## タイトル2

	* 箇条書き1
	* 箇条書き2

	[リンク先](http://example.com)
	`
	wants := `# タイトル1
	## タイトル2

	* <ruby>箇条<rp>(</rp><rt>かじょう</rt><rp>)</rp></ruby><ruby>書き<rp>(</rp><rt>がき</rt><rp>)</rp></ruby>1
	* <ruby>箇条<rp>(</rp><rt>かじょう</rt><rp>)</rp></ruby><ruby>書き<rp>(</rp><rt>がき</rt><rp>)</rp></ruby>2

	[リンク<ruby>先<rp>(</rp><rt>さき</rt><rp>)</rp></ruby>](http://example.com)
	`

	conf, err := NewConfiguration(RUBY_PATTERN_HTML)
	if err != nil {
		t.Error(`miss`)
	}
	writer, err := NewRubyWriter(conf)
	if err != nil {
		t.Error(`miss`)
	}

	output, err := writer.AddRuby(src)
	if err != nil {
		t.Error(`miss`)
	}

	if output != wants {
		t.Error(fmt.Printf("wants %s but %s", wants, output))
	}
}
