def decode_path(message, curr_location):
    direction = message[0]
    steps = int(message[1:])
    path = []
    if direction == "U":
        path = [(curr_location[0], i) for i in range(curr_location[1]+1, curr_location[1]+steps+1)]
    elif direction == "D":
        path = [(curr_location[0], i) for i in range(curr_location[1]-1, curr_location[1]-steps-1, -1)]
    elif direction == "R":
        path = [(i, curr_location[1]) for i in range(curr_location[0]+1, curr_location[0]+steps+1)]
    elif direction == "L":
        path = [(i, curr_location[1]) for i in range(curr_location[0]-1, curr_location[0]-steps-1, -1)]
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
for i in w1_input:
    w1_path += decode_path(i, w1_curr)
    w1_curr = w1_path[-1]

w2_curr = (0, 0)
for i in w2_input:
    w2_path += decode_path(i, w2_curr)
    w2_curr = w2_path[-1]

w1_path = set(w1_path)
w2_path = set(w2_path)

intersection_points = w1_path.intersection(w2_path)

intersection_points = list(intersection_points)

intersection_distances = []
for i in intersection_points:
    intersection_distances.append(abs(i[0]) + abs(i[1]))

intersection_distances = sorted(intersection_distances)

print(intersection_distances[0])
