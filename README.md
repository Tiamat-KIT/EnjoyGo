# EnjoyGo
Go言語をエンジョイするコーディングをしよう

### まずAPI作ろうぜ
[Iris](https://www.iris-go.com/)を使います
とりあえず、REST APIを作りましょう


### Webアプリケーションフレームワークを使おう
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