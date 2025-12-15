#!/usr/bin/env python3
import sys

START_INDEX = 100050

def convert_to_change(str_change):
    str_direction = str_change[:1]
    magnatude = int(str_change[1:])
    int_direction = 1
    if ("L" == str_direction):
        int_direction = -1
    int_change = (int_direction * magnatude)
    return int_change

def floor_impl(name):
    my_file = open(name)
    running_total = START_INDEX
    zero_count = 0
    prev_roations = 1000
    skip_adj = 0
    for line in my_file:
        int_change = convert_to_change(line)
        print(line.strip('\r\n'), running_total, (running_total + int_change),  end=' ')
        total_passes = get_total_passes_for_change(running_total, int_change)
        total_zero_lands = get_end_on_zero_for_change(running_total, int_change)
        zero_count = zero_count + total_passes + total_zero_lands
        if(total_passes > 0):
            print("Passes zero", total_passes, end=" ")
        if(total_zero_lands > 0):
            print("Lands on Zero" , total_zero_lands, end=" ")
        running_total = running_total + int_change
        print("")

    print(zero_count)
    
def get_total_passes_for_change(og_running_total, int_change):
    running_total = og_running_total + int_change
    prev_roations = og_running_total // 100
    if((og_running_total % 100) == 0):
        if(int_change < 0):
            # print("adj prev_rotations", end=" ")
            prev_roations = prev_roations - 1
    
    new_roations = running_total // 100
    if((running_total % 100) == 0):
        if(int_change > 0):
            # print("adj new_rotations", end=" ")
            new_roations = new_roations - 1

    print(prev_roations, new_roations, end=" ")
    return abs(prev_roations - new_roations)

def get_end_on_zero_for_change(og_running_total, int_change):
    if(((og_running_total + int_change) % 100) == 0):
        return 1
    return 0

floor_impl("rotations.txt")