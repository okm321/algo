# ABC301 C問題 解説 - AtCoder Cards

## 問題の理解

2つの文字列S、Tが与えられ、以下の操作で同じにできるか判定する問題です：
1. 文字列内の文字は自由に並び替え可能
2. `@`は`a,t,c,o,d,e,r`のいずれかに置き換え可能

## 解法の考え方

### ステップ1: 文字をカウント
各文字列の文字の出現回数を数えます（既に実装済み！）

### ステップ2: 判定ロジック
1. 各文字について、SとTでの出現回数の差を調べる
2. 差がある文字が`atcoder`に含まれる文字なら、`@`で補える
3. 必要な`@`の数が、使える`@`の数以下ならYes

## 実装方法

```go
func main() {
    defer w.Flush()
    
    S := readString(r)
    T := readString(r)
    cs := count(S)
    ct := count(T)
    
    // @で置き換え可能な文字
    atcoder := map[string]bool{
        "a": true, "t": true, "c": true, "o": true,
        "d": true, "e": true, "r": true,
    }
    
    // Sで必要な@の数、Tで必要な@の数
    needAtS := 0
    needAtT := 0
    
    // 全ての文字をチェック
    allChars := make(map[string]bool)
    for c := range cs {
        allChars[c] = true
    }
    for c := range ct {
        allChars[c] = true
    }
    
    for c := range allChars {
        if c == "@" {
            continue
        }
        
        diff := cs[c] - ct[c]
        
        if diff > 0 {
            // Sの方が多い → Tで@を使う必要がある
            if atcoder[c] {
                needAtT += diff
            } else {
                // @で置き換えられない文字で差がある
                fmt.Fprintln(w, "No")
                return
            }
        } else if diff < 0 {
            // Tの方が多い → Sで@を使う必要がある
            if atcoder[c] {
                needAtS += -diff
            } else {
                // @で置き換えられない文字で差がある
                fmt.Fprintln(w, "No")
                return
            }
        }
    }
    
    // @が足りるかチェック
    if needAtS <= cs["@"] && needAtT <= ct["@"] {
        fmt.Fprintln(w, "Yes")
    } else {
        fmt.Fprintln(w, "No")
    }
}
```

## ポイント

1. **文字の差を計算**：各文字でSとTの出現回数の差を見る
2. **@で補えるか確認**：差がある文字が`atcoder`に含まれるかチェック
3. **@の数が足りるか確認**：必要な@の数 ≤ 使える@の数

## 例

S = "ch@ku@ai", T = "choku@@i"の場合：
- Sは`a`が1個多い → Tで@を1個使って`a`にする
- Tは`o`が1個多い → Sで@を1個使って`o`にする
- Sの@は2個、Tの@は2個あるので、両方とも足りる → Yes!