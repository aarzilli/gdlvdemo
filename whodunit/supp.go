package main

func theBaroness(idx int) (*Thing, *Thing) {
	return &Thing{}, &Thing{}
}

func theWealthyCapitalist(idx int) *Thing {
	return &Thing{X:idx, Y:idx}
}

func mrReadHarrings(*Thing, int) (*Thing, *Thing) {
	return &Thing{}, &Thing{}
}

func thePainter(in *Thing) *Thing {
	if in.X == 48 {
		return nil
	}
	return &Thing{}
}

func theButler() *Thing {
	return &Thing{}
}

func P(idx int) bool {
	return (idx-1)%3==0
}

func execute(int, int) {
}
