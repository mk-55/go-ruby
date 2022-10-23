# go-ruby

<ruby>日本語<rp>(</rp><rt>にほんご</rt><rp>)</rp></ruby>の<ruby>文章<rp>(</rp><rt>ぶんしょう</rt><rp>)</rp></ruby>にルビを<ruby>振る<rp>(</rp><rt>ふる</rt><rp>)</rp></ruby>ためのプログラムです。このREADMEのルビも<ruby>本<rp>(</rp><rt>ほん</rt><rp>)</rp></ruby>プログラムでルビを<ruby>振り<rp>(</rp><rt>ふり</rt><rp>)</rp></ruby>ました。
<ruby>内部<rp>(</rp><rt>ないぶ</rt><rp>)</rp></ruby>ではkagomeを<ruby>利用<rp>(</rp><rt>りよう</rt><rp>)</rp></ruby>した<ruby>形態素<rp>(</rp><rt>けいたいそ</rt><rp>)</rp></ruby><ruby>解析<rp>(</rp><rt>かいせき</rt><rp>)</rp></ruby>を<ruby>実施<rp>(</rp><rt>じっし</rt><rp>)</rp></ruby>して<ruby>読み<rp>(</rp><rt>よみ</rt><rp>)</rp></ruby><ruby>仮名<rp>(</rp><rt>かめい</rt><rp>)</rp></ruby>を<ruby>取得<rp>(</rp><rt>しゅとく</rt><rp>)</rp></ruby>し、ルビ<ruby>記法<rp>(</rp><rt>きほう</rt><rp>)</rp></ruby>に<ruby>変換<rp>(</rp><rt>へんかん</rt><rp>)</rp></ruby>しています

※<ruby>現時点<rp>(</rp><rt>げんじてん</rt><rp>)</rp></ruby>のプログラムでルビを<ruby>振っ<rp>(</rp><rt>ふっ</rt><rp>)</rp></ruby>た<ruby>上記<rp>(</rp><rt>じょうき</rt><rp>)</rp></ruby>の<ruby>文章<rp>(</rp><rt>ぶんしょう</rt><rp>)</rp></ruby>では`よみかな`を`よみかめい`と<ruby>正しく<rp>(</rp><rt>ただしく</rt><rp>)</rp></ruby>ないルビを<ruby>振っ<rp>(</rp><rt>ふっ</rt><rp>)</rp></ruby>ています。

## コマンド

<ruby>現時点<rp>(</rp><rt>げんじてん</rt><rp>)</rp></ruby>では、<ruby>以下<rp>(</rp><rt>いか</rt><rp>)</rp></ruby>のように<ruby>引数<rp>(</rp><rt>ひきすう</rt><rp>)</rp></ruby>の<ruby>文字<rp>(</rp><rt>もじ</rt><rp>)</rp></ruby><ruby>列<rp>(</rp><rt>れつ</rt><rp>)</rp></ruby>にHTMLのrubyタグを<ruby>追加<rp>(</rp><rt>ついか</rt><rp>)</rp></ruby>した<ruby>文字<rp>(</rp><rt>もじ</rt><rp>)</rp></ruby><ruby>列<rp>(</rp><rt>れつ</rt><rp>)</rp></ruby>を<ruby>標準<rp>(</rp><rt>ひょうじゅん</rt><rp>)</rp></ruby><ruby>出力<rp>(</rp><rt>しゅつりょく</rt><rp>)</rp></ruby>に<ruby>出力<rp>(</rp><rt>しゅつりょく</rt><rp>)</rp></ruby>します。

```bash
$ go-ruby 吾輩は猫である
args : []string{"go-ruby", "吾輩は猫である"}
<ruby>吾輩<rp>(</rp><rt>わがはい</rt><rp>)</rp></ruby>は<ruby>猫<rp>(</rp><rt>ねこ</rt><rp>)</rp></ruby>である

```

また<ruby>標準<rp>(</rp><rt>ひょうじゅん</rt><rp>)</rp></ruby><ruby>入力<rp>(</rp><rt>にゅうりょく</rt><rp>)</rp></ruby>を<ruby>受け取る<rp>(</rp><rt>うけとる</rt><rp>)</rp></ruby>こともできます

```bash
$ cat ./README_base.md | go-ruby >> README.md

```

## <ruby>既知<rp>(</rp><rt>きち</rt><rp>)</rp></ruby>の<ruby>問題<rp>(</rp><rt>もんだい</rt><rp>)</rp></ruby>

<ruby>現時点<rp>(</rp><rt>げんじてん</rt><rp>)</rp></ruby>では<ruby>以下<rp>(</rp><rt>いか</rt><rp>)</rp></ruby>の<ruby>問題<rp>(</rp><rt>もんだい</rt><rp>)</rp></ruby>があります。

* すでにルビを<ruby>振っ<rp>(</rp><rt>ふっ</rt><rp>)</rp></ruby>た<ruby>文書<rp>(</rp><rt>ぶんしょ</rt><rp>)</rp></ruby>を<ruby>入力<rp>(</rp><rt>にゅうりょく</rt><rp>)</rp></ruby>として<ruby>与え<rp>(</rp><rt>あたえ</rt><rp>)</rp></ruby>た<ruby>場合<rp>(</rp><rt>ばあい</rt><rp>)</rp></ruby><ruby>正常<rp>(</rp><rt>せいじょう</rt><rp>)</rp></ruby>に<ruby>動作<rp>(</rp><rt>どうさ</rt><rp>)</rp></ruby>しません
* <ruby>入力<rp>(</rp><rt>にゅうりょく</rt><rp>)</rp></ruby>した<ruby>文字<rp>(</rp><rt>もじ</rt><rp>)</rp></ruby><ruby>列<rp>(</rp><rt>れつ</rt><rp>)</rp></ruby>すべてにルビを<ruby>振る<rp>(</rp><rt>ふる</rt><rp>)</rp></ruby>ため、<ruby>細かい<rp>(</rp><rt>こまかい</rt><rp>)</rp></ruby><ruby>制御<rp>(</rp><rt>せいぎょ</rt><rp>)</rp></ruby>ができません(preタグの<ruby>中<rp>(</rp><rt>なか</rt><rp>)</rp></ruby>はルビの<ruby>対象<rp>(</rp><rt>たいしょう</rt><rp>)</rp></ruby><ruby>外<rp>(</rp><rt>がい</rt><rp>)</rp></ruby>にする)
