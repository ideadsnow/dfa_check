# Go 实现的基于 DFA 的敏感词过滤库

[DFA 确定有限自动状态机](https://zh.wikipedia.org/wiki/%E7%A1%AE%E5%AE%9A%E6%9C%89%E9%99%90%E7%8A%B6%E6%80%81%E8%87%AA%E5%8A%A8%E6%9C%BA)

时间复杂度 O(n)，n 是匹配文本的长度

Example：

```go
// 敏感词列表
invalidWords := []string{"黑词", "敏感词"}

f := New()
f.AddInvalidWords(invalidWords)

// 检查文本中包含的敏感词
found, exist := f.Check("有敏感词文本，黑#wefwef词xxxx，违禁 ooohhg词，hh$%fppp")

fmt.Println(found, exist)
```

Output：

```go
[黑#wefwef词 违禁 ooohhg词] true
```

其它能力：

- 过滤敏感词
- 替换敏感词为指定词
