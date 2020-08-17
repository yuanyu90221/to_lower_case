# to-lower-case

## 題目解讀：

### 題目來源:

[to-lower-case](https://leetcode.com/problems/to-lower-case/)
### 原文:

Implement function ToLowerCase() that has a string parameter str, and returns the same string in lowercase.


### 解讀：
實做ToLowerCase這個function

把所有輸入的字串 轉成小寫的形式輸出

## 初步解法:
### 初步觀察:
每個character都是一個hex code的形式儲存
因此 只要對於每個輸入的 character檢查範圍

是否在 'A' ~ 'Z'的range之內

就知道是否 是Capitcal character

然後只要 把每個 大寫的character 加上 offset of 'a'-'A'

就可以轉換成大寫

### 初步設計:
Given a string s

Step 0: let a integer index = 0, a string result = ""

Step 1: if index > length of s go to step 4

Step 2: if s[index] >= 'A' and s['index'] <= Z  set result += string(s[index]+'a'-'A') else set result += string(s[index])

Step 3: index = index + 1 go to step 1

Step 4: return result

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package to_lower_case

func toLowerCase(str string) string {
	offset := 'a' - 'A'
	result := ""
	for _, r := range str {
		if r >= 'A' && r <= 'Z' {
			result += string(r + offset)
		} else {
			result += string(r)
		}
	}
	return result
}

```
## 測資的撰寫

```golang
package to_lower_case

import "testing"

func Test_toLowerCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example1",
			args: args{
				str: "Hello",
			},
			want: "hello",
		},
		{
			name: "Example2",
			args: args{
				str: "here",
			},
			want: "here",
		},
		{
			name: "Example3",
			args: args{
				str: "LOVELY",
			},
			want: "lovely",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toLowerCase(tt.args.str); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang leetcode 30day 21th day](https://hackmd.io/3Z72DVlTTtq_y8zSj6jHeQ?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)