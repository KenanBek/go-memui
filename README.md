# Go MemUI

*go-memu* - is in-memory object storage which can generate an automated JSON API for added values.

📚 It is an experimental Open Source Go package for learning purposes.

## ℹ️ Learning Materials

The whole development process is available as a step-by-step tutorial series. You can check the video course here:

#### [⚓️ Building Go-MemUI Project](https://www.codervlogger.com/build-memui-creating-opensource-go-project-full-tutorial/)

## 📜 How to use?

```go
p1 := Person{"John", 30}
p2 := Person{"Mary", 25}

mui := memui.New()
mui.Register(p1, p2)
```
