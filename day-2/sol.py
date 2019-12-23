with open("input", "r") as f:
    int_op_code = f.read()


int_op_code = [int(x) for x in int_op_code.split(",")]
ans = -1
ans_found = False

# handle 1202 program alarm state
for i in range(100):
    
    if ans_found:
        break

    for j in range(100):
        
        # print(i, j)

        iter_int_op_code = list(int_op_code)

        iter_int_op_code[1] = i
        iter_int_op_code[2] = j

        have_to_exit = False
        curr_index = 0

        while not have_to_exit:
            if max(iter_int_op_code[curr_index+3], iter_int_op_code[curr_index+2], iter_int_op_code[curr_index+1]) > len(int_op_code):
                break
            
            curr_int_code = iter_int_op_code[curr_index]

            if curr_int_code == 1:
                iter_int_op_code[iter_int_op_code[curr_index+3]] = iter_int_op_code[iter_int_op_code[curr_index+1]] + iter_int_op_code[iter_int_op_code[curr_index+2]]

            elif curr_int_code == 2:
                iter_int_op_code[iter_int_op_code[curr_index+3]] = iter_int_op_code[iter_int_op_code[curr_index+1]] * iter_int_op_code[iter_int_op_code[curr_index+2]]

            elif curr_int_code == 99:
                have_to_exit = True

            else:
                curr_index -= 3

            curr_index += 4

        if iter_int_op_code[0] == 19690720:
            ans = 100 * i + j
            ans_found = True
            break

print(ans)