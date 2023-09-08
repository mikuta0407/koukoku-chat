# koukoku-chat-tui

mattn氏の電子公告のチャット機能を利用できるチャットクライアント…のTUIクライアント

![image](https://github.com/mikuta0407/koukoku-chat-tui/assets/13357430/8a2cf2c8-988e-4771-becc-1868ee2fa066)

## 用途

See https://twitter.com/dnobori/status/1699339202104889474

## Installation

```
go install github.com/mikuta0407/koukoku-chat-tui@latest
```

## 使い方

- `koukoku-chat-tui`で↑のスクショみたいな画面が出てきます。
- 入力欄と表示欄が分離して見やすいです。(入力訂正がしやすい)

## 既知の不具合

- Ctrl+Cを2回叩かないと完全には泊まらない
  - 治す気はあります。
- 稀にエスケープシーケンスの文字列が出てくる

## License

MIT

## Author

- mikuta0407
- Yasuhiro Matsumoto (a.k.a. mattn)
