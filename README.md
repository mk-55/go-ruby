# go-ruby

<ruby>日本語<rp>(</rp><rt>にほんご</rt><rp>)</rp></ruby>の<ruby>文章<rp>(</rp><rt>ぶんしょう</rt><rp>)</rp></ruby>にルビを<ruby>振る<rp>(</rp><rt>ふる</rt><rp>)</rp></ruby>ためのプログラムです。このREADMEのルビも<ruby>本<rp>(</rp><rt>ほん</rt><rp>)</rp></ruby>プログラムでルビを<ruby>振り<rp>(</rp><rt>ふり</rt><rp>)</rp></ruby>ました。

<ruby>内部<rp>(</rp><rt>ないぶ</rt><rp>)</rp></ruby>では[kagome](https://github.com/ikawaha/kagome)を<ruby>利用<rp>(</rp><rt>りよう</rt><rp>)</rp></ruby>した<ruby>形態素<rp>(</rp><rt>けいたいそ</rt><rp>)</rp></ruby><ruby>解析<rp>(</rp><rt>かいせき</rt><rp>)</rp></ruby>を<ruby>実施<rp>(</rp><rt>じっし</rt><rp>)</rp></ruby>して<ruby>読み<rp>(</rp><rt>よみ</rt><rp>)</rp></ruby><ruby>仮名<rp>(</rp><rt>かめい</rt><rp>)</rp></ruby>を<ruby>取得<rp>(</rp><rt>しゅとく</rt><rp>)</rp></ruby>し、ルビ<ruby>記法<rp>(</rp><rt>きほう</rt><rp>)</rp></ruby>に<ruby>変換<rp>(</rp><rt>へんかん</rt><rp>)</rp></ruby>しています。

<ruby>現時点<rp>(</rp><rt>げんじてん</rt><rp>)</rp></ruby>のプログラムでルビを<ruby>振っ<rp>(</rp><rt>ふっ</rt><rp>)</rp></ruby>た<ruby>上記<rp>(</rp><rt>じょうき</rt><rp>)</rp></ruby>の<ruby>文章<rp>(</rp><rt>ぶんしょう</rt><rp>)</rp></ruby>では`よみがな`を`よみかめい`と<ruby>正しく<rp>(</rp><rt>ただしく</rt><rp>)</rp></ruby>ないルビを<ruby>振っ<rp>(</rp><rt>ふっ</rt><rp>)</rp></ruby>ています。

## コマンド

<ruby>現時点<rp>(</rp><rt>げんじてん</rt><rp>)</rp></ruby>では、<ruby>以下<rp>(</rp><rt>いか</rt><rp>)</rp></ruby>のように<ruby>引数<rp>(</rp><rt>ひきすう</rt><rp>)</rp></ruby>の<ruby>文字<rp>(</rp><rt>もじ</rt><rp>)</rp></ruby><ruby>列<rp>(</rp><rt>れつ</rt><rp>)</rp></ruby>にHTMLのrubyタグを<ruby>追加<rp>(</rp><rt>ついか</rt><rp>)</rp></ruby>した<ruby>文字<rp>(</rp><rt>もじ</rt><rp>)</rp></ruby><ruby>列<rp>(</rp><rt>れつ</rt><rp>)</rp></ruby>を<ruby>出力<rp>(</rp><rt>しゅつりょく</rt><rp>)</rp></ruby>します。

```bash
$ go-ruby 吾輩は猫である
args : []string{"go-ruby", "吾輩は猫である"}
<ruby>吾輩<rp>(</rp><rt>わがはい</rt><rp>)</rp></ruby>は<ruby>猫<rp>(</rp><rt>ねこ</rt><rp>)</rp></ruby>である

```