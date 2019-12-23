begin = 138241 
end = 674034

def valid_pass(password):
    adj_rule = False
    ord_rule = True
    for i in range(1,len(password)):
        if password[i] < password[i-1]:
            ord_rule = False
            break
        elif password[i] == password[i-1]:
            if not ((i-2 >= 0 and password[i-2] == password[i-1]) or (i+1 < len(password) and password[i] == password [i+1])):
                adj_rule = True
    return adj_rule and ord_rule

count = 0
for i in range(begin, end+1):
    if valid_pass(str(i)):
        count += 1

print(count)
