package dictionary

import "fmt"

type QwertyKey string

const (
	Escape    QwertyKey = "Escape"
	Space     QwertyKey = "Space"
	Tab       QwertyKey = "Tab"
	Return    QwertyKey = "Return"
	Home      QwertyKey = "Home"
	PageUp    QwertyKey = "Page_Up"
	PageDown  QwertyKey = "Page_Down"
	End       QwertyKey = "End"
	Backspace QwertyKey = "BackSpace"
	Delete    QwertyKey = "Delete"
	Left      QwertyKey = "Left"
	Up        QwertyKey = "Up"
	Down      QwertyKey = "Down"
	Right     QwertyKey = "Right"
	QwertyA   QwertyKey = "a"
	QwertyB   QwertyKey = "b"
	QwertyC   QwertyKey = "c"
	QwertyD   QwertyKey = "d"
	QwertyE   QwertyKey = "e"
	QwertyF   QwertyKey = "f"
	QwertyG   QwertyKey = "g"
	QwertyH   QwertyKey = "h"
	QwertyI   QwertyKey = "i"
	QwertyJ   QwertyKey = "j"
	QwertyK   QwertyKey = "k"
	QwertyL   QwertyKey = "l"
	QwertyM   QwertyKey = "m"
	QwertyN   QwertyKey = "n"
	QwertyO   QwertyKey = "o"
	QwertyP   QwertyKey = "p"
	QwertyQ   QwertyKey = "q"
	QwertyR   QwertyKey = "r"
	QwertyS   QwertyKey = "s"
	QwertyT   QwertyKey = "t"
	QwertyU   QwertyKey = "u"
	QwertyV   QwertyKey = "v"
	QwertyW   QwertyKey = "w"
	QwertyX   QwertyKey = "x"
	QwertyY   QwertyKey = "y"
	QwertyZ   QwertyKey = "z"
	F1        QwertyKey = "F1"
	F2        QwertyKey = "F2"
	F3        QwertyKey = "F3"
	F4        QwertyKey = "F4"
	F5        QwertyKey = "F5"
	F6        QwertyKey = "F6"
	F7        QwertyKey = "F7"
	F8        QwertyKey = "F8"
	F9        QwertyKey = "F9"
	F10       QwertyKey = "F10"
	F11       QwertyKey = "F11"
	F12       QwertyKey = "F12"
	N1        QwertyKey = "1"
	N2        QwertyKey = "2"
	N3        QwertyKey = "3"
	N4        QwertyKey = "4"
	N5        QwertyKey = "5"
	N6        QwertyKey = "6"
	N7        QwertyKey = "7"
	N8        QwertyKey = "8"
	N9        QwertyKey = "9"
	N0        QwertyKey = "0"
)

type QwertyMod string

const (
	Shift QwertyMod = "Shift_L(%s)"
	Ctrl  QwertyMod = "Control_L(%s)"
	Alt   QwertyMod = "Alt_L(%s)"
	Gui   QwertyMod = "Super_L(%s)"
)

func (qm QwertyMod) apply(s string) QwertyMod {
	return QwertyMod(fmt.Sprintf(string(qm), s))
}
