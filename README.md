# EnjoyGo
Go言語をエンジョイするコーディングをしよう

## まずAPI作ろうぜ
~~[Iris](https://www.iris-go.com/)を使います~~

技術自体に罪はないと考えていますが、どうやらGo言語界隈では
- [**Irisの作者**に対して良いイメージがない](https://github.com/the-benchmarker/web-frameworks/issues/1129)こと
- 動作が最速であるというベンチマークの結果も信用に足るものではないこと
上記を理由としてGinに差し替えます

[Gin](https://gin-gonic.com/ja/)

とりあえず、REST APIを作りましょう

### 参考リンク
[golangフレームワークginを使ってみる](https://zenn.dev/ajapa/articles/6471ac0c612fda)

## Webアプリケーションフレームワークを使おう
[Revel](https://revel.github.io/)を使います
ToDoアプリとか作ります

ToDoアプリの単一ToDoアイテムの構造は以下を定義する。以下JSON構造のイメージ
|プロパティ名|型|説明|
|----|----|----|
|title|文字列|タイトル|
|id|UUID|アイテムID|
|description|文字列|説明文|
|limit|文字列または日付型|期限日|

これを構造体で定義したい。
データベースはせっかくGoogle謹製の言語使ってるから
FirebaseのFireStoreにしちゃおう、言語から何からGoogleにしちゃおう