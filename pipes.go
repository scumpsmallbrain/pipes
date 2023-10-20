package main
import ("fmt")

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
	for {
		pipes.print()
		var command string;
		fmt.Scanln(&command);
		if command[0] == 'q' { break }
		p := look_for_pipe(command[0], &pipes);
		if p == nil { continue }

		switch (command[1]) {
		case '+':
			p.plus();
		case '-':
			p.minus();
		}
	}
}

func look_for_pipe(s byte, pl *pipe_list) *pipe {
	t := (int)(s) - 97;
	p := pl.head;
	for i := 0; i != t && p != nil; i++ {
		p = p.next;
	}
	return p;
}

func (pi *pipe_list) append(p pipe) {
	last_tail := pi.tail;
	pi.tail = &p;
	last_tail.next = &p;
	p.last = last_tail;
}

func (pi *pipe_list) print() {
	fmt.Println();
	var i int = 0;
	for p:= *pi.head;; p = *p.next {
		fmt.Printf(" %c \033[1B\033[3D", (rune)(i + 97))
		fmt.Printf(" %d \033[1A", p.state);
		i++;
		if p == *pi.tail { break }
	}
	fmt.Println("\n\n");
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

func (p *pipe) plus() {
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
	}
}

func (p *pipe) minus() {
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
	}
}