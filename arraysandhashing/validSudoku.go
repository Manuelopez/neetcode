package arraysandhashing

func isValidSudoku(board [][]byte) bool {
    cols := make(map[int]map[int]bool)
    rows := make(map[int]map[int]bool)
    squares := make(map[int]map[int]bool)
    cols[0] = make(map[int]bool)
    cols[1] = make(map[int]bool)
    cols[2] = make(map[int]bool)
    cols[3] = make(map[int]bool)
    cols[4] = make(map[int]bool)
    cols[5] = make(map[int]bool)
    cols[6] = make(map[int]bool)
    cols[7] = make(map[int]bool)
    cols[8] = make(map[int]bool)



    rows[0] = make(map[int]bool)
    rows[1] = make(map[int]bool)
    rows[2] = make(map[int]bool)
    rows[3] = make(map[int]bool)
    rows[4] = make(map[int]bool)
    rows[5] = make(map[int]bool)
    rows[6] = make(map[int]bool)
    rows[7] = make(map[int]bool)
    rows[8] = make(map[int]bool)

    squares[1] = make(map[int]bool)
    squares[2] = make(map[int]bool)
    squares[3] = make(map[int]bool)
    squares[4] = make(map[int]bool)
    squares[5] = make(map[int]bool)
    squares[6] = make(map[int]bool)
    squares[7] = make(map[int]bool)
    squares[8] = make(map[int]bool)
    squares[9] = make(map[int]bool)
    for i := 0; i < 81; i++{
        x, y := DivideIntoTwoNumbers(i)
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
                if squares[1][num] == true{

                    return false
                }else{
                    squares[1][num] = true
                }
    
            } else if y <=5{
                // square2
                  if squares[2][num] == true{

                    return false
                }else{
                    squares[2][num] = true
                }
            }else if y <= 8{
                // square3
                  if squares[3][num] == true{

                    return false
                }else{
                    squares[3][num] = true
                }
            }
        } else if x <=5 {
            if y <=2{
                // square 4

                  if squares[4][num] == true{

                    return false
                }else{
                    squares[4][num] = true
                }
            } else if y <=5 {
                // square 5

                  if squares[5][num] == true{

                    return false
                }else{
                    squares[5][num] = true
                }
            }else if y <= 8{
                // square 6

                  if squares[6][num] == true{

                    return false
                }else{
                    squares[6][num] = true
                }
            }
        }else if x <= 8{

            if y <=2{
                // square 7

                  if squares[7][num] == true{

                    return false
                }else{
                    squares[7][num] = true
                }
            } else if y <=5{
                // square 8
                  if squares[8][num] == true{
                    return false
                }else{
                    squares[8][num] = true
                }
            }else if y <= 8{
                // square 9

                  if squares[9][num] == true{

                    return false
                }else{
                    squares[9][num] = true
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




