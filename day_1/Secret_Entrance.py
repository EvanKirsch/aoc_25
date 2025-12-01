#!/usr/bin/env python3

START_INDEX = 50
MIN_INDEX = 0
MAX_INDEX = 99

class Index_Info:
    def __init__(self, index, zero_count):
        self.index = index
        self.zero_count = zero_count

    def getIndex(self):
        return self.index

    def getZeroCount(self):
        return self.zero_count

def convert_to_index(current_index, change):
    new_index = int(current_index) + int(change)
    zero_count = 0
    changed_index = False
    while (new_index > MAX_INDEX):
        new_index = new_index - (MAX_INDEX + 1)
        if ((current_index != MIN_INDEX or changed_index) and new_index != MIN_INDEX):
            print("pass 0")
            zero_count = zero_count + 1
        changed_index = True

    while (new_index < MIN_INDEX):
        new_index = MAX_INDEX - abs(new_index) + 1
        if ((current_index != MIN_INDEX or changed_index) and new_index != MIN_INDEX):
            print("pass 0")
            zero_count = zero_count + 1
        changed_index = True

    return Index_Info(new_index, zero_count)

def convert_to_change(str_change):
    str_direction = str_change[:1]
    magnatude = int(str_change[1:])
    int_direction = 1
    if ("L" == str_direction):
        int_direction = -1
    int_change = (int_direction * magnatude)
    return int_change

my_file = open("test.txt")
#my_file = open("rotations.txt")
my_index = int(START_INDEX)
zero_count = int(0)
int_change = int(0)
for line in my_file:
    print(line.replace("\n",""))
    int_change = convert_to_change(line)
    my_index_info = convert_to_index(my_index, int_change)
    my_index = my_index_info.getIndex()
    change_zero_count = my_index_info.getZeroCount()
    zero_count = zero_count + change_zero_count
    if (my_index == MIN_INDEX):
       print("points at 0")
       zero_count = zero_count + 1

print("zero count: ", zero_count)
