package arraysandhashing

func isValidSudoku(board [][]byte) bool {
    cols := make(map[int]map[int]bool)
    rows := make(map[int]map[int]bool)
    squares := make(map[int]map[int]bool)

    for i := 0; i < 81; i++{
        x, y := DivideIntoTwoNumbers(i)
        if _, ok := cols[x]; !ok{
            cols[x] = make(map[int]bool)
        }

        if _, ok := cols[y]; !ok{
            cols[y] = make(map[int]bool)
        }

        if _, ok := rows[x]; !ok{
            rows[x] = make(map[int]bool)
        }

        if _, ok := rows[y]; !ok{
            rows[y] = make(map[int]bool)
        }

        if _, ok := squares[x]; !ok{
            squares[x] = make(map[int]bool)
        }

        if _, ok := squares[y]; !ok{
            squares[y] = make(map[int]bool)
        }
        if board[x][y] == '.'{
            continue
        }
        num := int(board[x][y] - '0')
        if cols[y][num] == true{
            return false
        }else{
            cols[y][num] = true
        }
        if rows[x][num] == true{

            return false
        }else{
            rows[x][num] = true
        }
        if x <=2 {
            if y <=2 {
                if squares[0][num] == true{

                    return false
                }else{
                    squares[0][num] = true
                }
    
            } else if y <=5{
                // square2
                  if squares[1][num] == true{
                    return false
                }else{
                    squares[1][num] = true
                }
            }else if y <= 8{
                // square3
                  if squares[2][num] == true{

                    return false
                }else{
                    squares[2][num] = true
                }
            }
        } else if x <=5 {
            if y <=2{
                // square 4

                  if squares[3][num] == true{

                    return false
                }else{
                    squares[3][num] = true
                }
            } else if y <=5 {
                // square 5

                  if squares[4][num] == true{

                    return false
                }else{
                    squares[4][num] = true
                }
            }else if y <= 8{
                // square 6

                  if squares[5][num] == true{

                    return false
                }else{
                    squares[5][num] = true
                }
            }
        }else if x <= 8{

            if y <=2{
                // square 7

                  if squares[6][num] == true{

                    return false
                }else{
                    squares[6][num] = true
                }
            } else if y <=5{
                // square 8
                  if squares[7][num] == true{
                    return false
                }else{
                    squares[7][num] = true
                }
            }else if y <= 8{
                // square 9

                  if squares[8][num] == true{

                    return false
                }else{
                    squares[8][num] = true
                }
            }
        }

    }

    return true
}

func DivideIntoTwoNumbers(x int) (int, int) {
	y := x / 9
	z := x % 9
	return y, z
}
