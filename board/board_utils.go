package board

func (b *Board) GetPositionValue(position Position) int8 {
	if position.X < 0 || position.X >= SIZE || position.Y < 0 || position.Y >= SIZE {
		return -1
	}
	return b.Tab[position.Y][position.X]
}

func (b *Board) GetSequenceValues(sequence *[]Position) *[]int8 {
	values := []int8{}
	for _, elem := range *sequence {
		values = append(values, b.Tab[elem.Y][elem.X])
	}
	return &values
}

func (b *Board) GetNPositionsSequence(start *Position, direction int8, len int) *[]Position {
	tmp := *start
	sequence := []Position{}
	for i := 0; i < len; i++ {
		if tmp.X < 0 || tmp.X > SIZE-1 ||
			tmp.Y < 0 || tmp.Y > SIZE-1 {
			break
		}
		sequence = append(sequence, tmp)
		tmp = b.GetNextPosition(tmp, direction)
	}
	return &sequence
}

// Build 4 positions sequences (NW-SE, N - S, NE-SW, W-E)
// Sequence starts at POS-len and ends at POS+len for every direction
func (b *Board) GetNSurroundingPositionsSequence(start *Position, length int) *[][]Position {
	sequences := [][]Position{}
	for dir := int8(0); dir < 4; dir++ {
		sequence := (*b.GetNPositionsSequence(start, dir, length+1))[1:]
		for i, j := 0, len(sequence)-1; i < j; i, j = i+1, j-1 {
			sequence[i], sequence[j] = sequence[j], sequence[i]
		}
		sequence = append(sequence, *b.GetNPositionsSequence(start, 7-dir, length+1)...)
		sequences = append(sequences, sequence)
	}
	return &sequences
}

func (b *Board) GetNextPosition(position Position, direction int8) Position {
	switch direction {
	case NW:
		return Position{position.X - 1, position.Y - 1}
	case N:
		return Position{position.X, position.Y - 1}
	case NE:
		return Position{position.X + 1, position.Y - 1}
	case W:
		return Position{position.X - 1, position.Y}
	case E:
		return Position{position.X + 1, position.Y}
	case SW:
		return Position{position.X - 1, position.Y + 1}
	case S:
		return Position{position.X, position.Y + 1}
	case SE:
		return Position{position.X + 1, position.Y + 1}
	}
	return Position{}
}

func int8ToInt32(toSquash []int8) uint32 {
	res := uint32(0)
	size := len(toSquash)
	for i, i8 := range toSquash {
		res |= uint32(i8) << uint((size-i)*2)
	}
	return res
}
