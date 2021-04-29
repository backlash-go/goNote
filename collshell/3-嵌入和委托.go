package main

import "fmt"

func main() {
	label := Label{Widget{10, 10}, "State1:"}

	button1 := Button{Label{Widget{10, 70}, "OK"}}
	button2 := Button{Label{Widget{50, 70}, "cancel"}}
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	for _, Widget:= range []interface{}{label}{
		if c, ok := Widget.(Painter); ok{
			fmt.Printf("%p\n",&Widget)

			c.Paint()
			fmt.Printf("%p\n",&c)
		}
	}

}






type Widget struct {
	X,Y int
}

type Label struct {
	Widget
	Text string
}

type Button struct {
	Label // Embedding (delegation)
}

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}


type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}


func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

//所以，可以在 Button 中可以重载这个接口方法以
func (button Button) Paint() { // Override
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}


func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}
func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}