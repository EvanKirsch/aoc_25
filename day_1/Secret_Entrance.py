#!/usr/bin/env python3
import sys

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
    count_zeros = (current_index == 0 and abs(change) > MAX_INDEX) or current_index != 0
    while (new_index > MAX_INDEX):
        if (count_zeros and new_index != (MAX_INDEX + 1)):
            print(f"pass 0 [{new_index}]", end='')
            zero_count = zero_count + 1
        new_index = new_index - (MAX_INDEX + 1)
        count_zeros = True

    while (new_index < MIN_INDEX):
        if (count_zeros and new_index != abs(new_index) + 1):
            print(f"pass 0 [{new_index}]", end='')
            zero_count = zero_count + 1
        new_index = MAX_INDEX - abs(new_index) + 1
        count_zeros = True

    if (new_index == MIN_INDEX):
      zero_count = zero_count + 1
    print(f"points at {new_index} ")

    return Index_Info(new_index, zero_count)

def convert_to_change(str_change):
    str_direction = str_change[:1]
    magnatude = int(str_change[1:])
    int_direction = 1
    if ("L" == str_direction):
        int_direction = -1
    int_change = (int_direction * magnatude)
    return int_change

def loop_impl(name):
    my_file = open(name)
    my_index = int(START_INDEX)
    zero_count = int(0)
    int_change = int(0)
    for line in my_file:
        print(f"{line.replace("\n","")} ", end='')
        int_change = convert_to_change(line)
        my_index_info = convert_to_index(my_index, int_change)
        my_index = my_index_info.getIndex()
        change_zero_count = my_index_info.getZeroCount()
        zero_count = zero_count + change_zero_count
        print(f"[{zero_count}, {change_zero_count}]")
    
    print("zero count: ", zero_count)

def recursive_impl(name):
    my_file = open(name)
    my_indexes = []
    for line in my_file:
        int_change = convert_to_change(line)
        my_indexes.append(int_change)
    zero_count = get_zero_passes(0, my_indexes, START_INDEX)
    print("zero count: ", zero_count)

def get_zero_passes(zero_count, my_indexes, current_index):
    change_index = my_indexes.pop(0)
    max_possible_location = current_index + change_index
    if (current_index == 0 and max_possible_location < 100 and max_possible_location > -100):
        updated_index = current_index + change_index
        if(updated_index < 0):
            updated_index = 100 + updated_index
        partial_change_amount = change_index

    elif (max_possible_location >= 100):
        partial_change_amount = (100-current_index)
        partial_remainder = change_index - (100 - current_index) 
        zero_count = zero_count + 1
        updated_index = 0
        if (partial_remainder != 0):
            my_indexes.insert(0, partial_remainder)

    elif (max_possible_location <= 0) :
        partial_change_amount = current_index
        partial_remainder = change_index + current_index 
        zero_count = zero_count + 1
        updated_index = 0
        if (partial_remainder != 0):
            my_indexes.insert(0, partial_remainder)

    else:
        updated_index = current_index + change_index
        if(updated_index < 0):
            updated_index = 100 + updated_index
        partial_change_amount = change_index

    print("rotated", partial_change_amount, "to point at", updated_index, "zero count:", zero_count)
    if(len(my_indexes) > 0): 
        zero_count = get_zero_passes(zero_count, my_indexes, updated_index)

    return zero_count

current_limit = sys.getrecursionlimit()
print("Current recursion limit:", current_limit)
sys.setrecursionlimit(10000)
current_limit = sys.getrecursionlimit()
print("Current recursion limit:", current_limit)
recursive_impl("rotations.txt")
# loop_impl("test.txt")
