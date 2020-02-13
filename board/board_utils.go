package board

func (b *Board) GetPositionValue(position Position) int8 {
	return b.Tab[position.Y][position.X]
}

func (b *Board) GetNPositionsSequence(start *Position, direction int8, len int) []Position {
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
	return sequence
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

func int8IntoInt32(i1 int8, i2 int8, i3 int8, i4 int8) int8 {
	i := i1 << 6
	i |= i2 << 4
	i |= i3 << 2
	i |= i4
	return i
}
