# ABC302 B問題 解説

## 問題概要
H×Wのマス目から「snuke」という文字列を見つける問題。文字は一直線上に等間隔で並んでいる必要がある。

## 解法

### 考え方
各マスを起点として、8方向に「snuke」があるか確認します。

### 8方向の移動
```
方向の定義:
(-1,-1) (-1,0) (-1,+1)
( 0,-1)   ●    ( 0,+1)
(+1,-1) (+1,0) (+1,+1)
```

### アルゴリズム
1. 全てのマス(i,j)を起点として
2. 8方向それぞれについて
3. 「snuke」の5文字が順番に並んでいるか確認
4. 見つかったら、その5マスの座標を出力

## 実装例（Go言語）

```go
package main

import (
    "fmt"
)

func main() {
    var H, W int
    fmt.Scan(&H, &W)
    
    // グリッドを読み込む
    grid := make([]string, H)
    for i := 0; i < H; i++ {
        fmt.Scan(&grid[i])
    }
    
    target := "snuke"
    
    // 8方向の移動量
    dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
    dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
    
    // 全マスを起点として探索
    for i := 0; i < H; i++ {
        for j := 0; j < W; j++ {
            // 8方向それぞれを試す
            for dir := 0; dir < 8; dir++ {
                // この方向に"snuke"があるか確認
                found := true
                positions := make([][2]int, 5)
                
                for k := 0; k < 5; k++ {
                    ni := i + dx[dir]*k
                    nj := j + dy[dir]*k
                    
                    // 範囲外チェック
                    if ni < 0 || ni >= H || nj < 0 || nj >= W {
                        found = false
                        break
                    }
                    
                    // 文字チェック
                    if grid[ni][nj] != target[k] {
                        found = false
                        break
                    }
                    
                    positions[k] = [2]int{ni, nj}
                }
                
                // 見つかったら出力
                if found {
                    for k := 0; k < 5; k++ {
                        fmt.Printf("%d %d\n", positions[k][0]+1, positions[k][1]+1)
                    }
                    return
                }
            }
        }
    }
}
```

## ポイント

### 1. 8方向の表現
```go
dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}
```
これで8方向への移動を表現できます。

### 2. 境界チェック
```go
if ni < 0 || ni >= H || nj < 0 || nj >= W
```
グリッドの外に出ないようにチェックが必要です。

### 3. 座標の出力
問題では1-indexedで出力するので、`+1`を忘れずに！

## 計算量
- 時間計算量: O(H × W × 8 × 5) = O(H × W)
- 空間計算量: O(H × W)（グリッドの保存）

制約が小さい（最大100×100）ので、全探索で十分間に合います！