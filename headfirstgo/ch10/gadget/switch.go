package main

type Toggleable interface {
	toggle()
}

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
}

func main() {
	s := Switch("off")
	// cannot use s (variable of type Switch) as Toggleable value in variable declaration:
	// missing method toggle (toggle has pointer receiver)compiler
	// 下面一行直接把s赋值给Toggleable类型的t时会报上面的错误
	// var t Toggleable = s
	var t Toggleable = &s
	t.toggle()
	t.toggle()
}
