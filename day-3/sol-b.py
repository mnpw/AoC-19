def update_timing(pos, wire, steps):
    if wire == 1:
        if pos not in w1_timing.keys():
            w1_timing[pos] = steps
        else:
            w1_timing[pos] = min(w1_timing[pos], steps) 
        # print("pos is", pos, "timing is", w1_timing)
    elif wire == 2:
        if pos not in w2_timing.keys():
            w2_timing[pos] = steps
        else:
            w2_timing[pos] = min(w2_timing[pos], steps) 
        # print("pos is", pos, "timing is", w2_timing)

def decode_path(message, curr_location, wire, steps_so_far):
    direction = message[0]
    steps = int(message[1:])
    path = []
    counter = 0
    if direction == "U":
        for i in range(curr_location[1]+1, curr_location[1]+steps+1):
            counter += 1
            path.append((curr_location[0], i))
            update_timing(path[-1], wire, steps_so_far+counter)
    elif direction == "D":
        for i in range(curr_location[1]-1, curr_location[1]-steps-1, -1):
            counter += 1
            path.append((curr_location[0], i))
            update_timing(path[-1], wire, steps_so_far+counter)
    elif direction == "R":
        for i in range(curr_location[0]+1, curr_location[0]+steps+1):
            counter += 1
            path.append((i, curr_location[1]))
            update_timing(path[-1], wire, steps_so_far+counter)
    elif direction == "L":
        for i in range(curr_location[0]-1, curr_location[0]-steps-1, -1):
            counter += 1
            path.append((i, curr_location[1]))
            update_timing(path[-1], wire, steps_so_far+counter)
    else:
        print("invalid message, cannot decode")
    return path


f = open("input", "r")
input_info = f.read().split()
f.close()

w1_path = []
w2_path = []

w1_input = input_info[0].split(',')
w2_input = input_info[1].split(',')

w1_curr = (0, 0)
w1_steps = 0
w1_timing = {}
for i in w1_input:
    w1_path += decode_path(i, w1_curr, 1, w1_steps)
    # print(w1_path)
    w1_curr = w1_path[-1]
    w1_steps += int(i[1:])

w2_curr = (0, 0)
w2_steps = 0
w2_timing = {}
for i in w2_input:
    w2_path += decode_path(i, w2_curr, 2, w2_steps)
    # print(w2_path)
    w2_curr = w2_path[-1]
    w2_steps += int(i[1:])

w1_path = set(w1_path)
w2_path = set(w2_path)

intersection_points = w1_path.intersection(w2_path)

intersection_points = list(intersection_points)

intersection_distances = []
for i in intersection_points:
    intersection_distances.append(w1_timing[i] + w2_timing[i])

intersection_distances = sorted(intersection_distances)

print(intersection_distances[0])
