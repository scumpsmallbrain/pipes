package main
import (
	"fmt"
	col "github.com/gookit/color"
	kb "atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)	

var graphics = [...]rune {
	'░',
	'▄',
	'█',
}

type pipe struct {
	state uint8;
	next *pipe;
	last *pipe;
}

type pipe_list struct {
	head *pipe;
	tail *pipe;
}

func main() {
	states := []uint8{1, 2, 1, 0, 1, 0, 2};
	pipes := new_pipe_list(states[0]);
	for i := (uint8)(1); i < (uint8)(len(states)); i++ {
		pipes.append(pipe{state: states[i]});
	}
	cursor_focus := pipes.head;
	pipes.display(cursor_focus)
}

func (pi *pipe_list) append(p pipe) {
	last_tail := pi.tail;
	pi.tail = &p;
	last_tail.next = &p;
	p.last = last_tail;
}

func (pi *pipe_list) display(cursor *pipe) bool {
	//show cursor again when function is over
	defer func() {
		fmt.Printf("\033[?25h");
	}();
	//hide cursor
	fmt.Printf("\033[?25l");
	pi.render(false, cursor)
	escaped := false
	kb.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.Right {
			cursor = pi.move_cursor(cursor, true);
			pi.render(true, cursor);
			return false, nil;
		} else if key.Code == keys.Left {
			cursor = pi.move_cursor(cursor, false);
			pi.render(true, cursor);
			return false, nil;
		} else if key.Code == keys.Escape {
			escaped = true;
			return true, nil;
		} else if key.Code == keys.Up {
			pi.render(!cursor.plus(), cursor);
			return false, nil;
		} else if key.Code == keys.Down {
			pi.render(!cursor.minus(), cursor);
			return false, nil;
		}
		return false, nil;
	});
	return escaped;
}

func (pi *pipe_list) move_cursor(cursor *pipe, right bool) *pipe {
	if right {
		if cursor.next == nil {
			cursor = pi.head;
		} else {
			cursor = cursor.next;
		}
	} else {
		if cursor.last == nil {
			cursor = pi.tail;
		} else {
			cursor = cursor.last;
		}
	}
	return cursor;
}

func (pi *pipe_list) render(redraw bool, cursor *pipe) {
	if redraw {
		fmt.Printf("\r");
	} else {
		fmt.Printf("\r\n\n");
	}

	cursor_style := col.HiGreen;
	adjacent_style := col.Cyan;
	
	var i int = 0;
	for p:= pi.head;; p = p.next {
		if p == cursor.last || p == cursor.next {
			adjacent_style.Printf("%c", graphics[p.state]);
		} else if p == cursor {
			cursor_style.Printf("%c", graphics[p.state]);
		} else {
			fmt.Printf("%c", graphics[p.state]);
		}
		i++;
		if p == pi.tail { break }
	}
}

func new_pipe_list(s uint8) pipe_list {
	firstelement := pipe{
		state: s,
		next: nil,
		last: nil,
	}
	return pipe_list{
		head: &firstelement,
		tail: &firstelement,
	}
}

func (p *pipe) plus() bool {
	dummy := pipe{state: 1}
	var next *pipe;
	var last *pipe;
	if p.next != nil {
		next = p.next;
	} else { next = &dummy }
	if p.last != nil {
		last = p.last;
	} else { last = &dummy }
	if p.state < 2 && next.state > 0 && last.state > 0 {
		p.state++;
		next.state--;
		last.state--;
		return true;
	} else {
		return false;
	}
}

func (p *pipe) minus() bool {
	dummy := pipe{state: 1}
	var next *pipe;
	var last *pipe;
	if p.next != nil {
		next = p.next;
	} else { next = &dummy }
	if p.last != nil {
		last = p.last;
	} else { last = &dummy }
	if p.state > 0 && next.state < 2 && last.state < 2 {
		p.state--;
		next.state++;
		last.state++;
		return true;
	} else {
		return false;
	}
}