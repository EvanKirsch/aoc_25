
def get_banks(fileName):
    my_file = open(fileName)
    my_banks = []
    for line in my_file:
        clean_line = line.rstrip('\r\n')
        my_banks.append(clean_line)
    return my_banks

def find_largest_battery(str_largest_battery, size, partial_banks):
    if (size == 0):
       return str_largest_battery
    avaliable_bank = partial_banks[:len(partial_banks) - (size -1 )] # reserving high power batteries deep in bank
    first_max_number_index = find_largest_avaliable_battery(avaliable_bank)
    str_largest_battery = str_largest_battery + partial_banks[first_max_number_index]
    remaining_banks = partial_banks[first_max_number_index + 1:]
    size = size - 1
    return find_largest_battery(str_largest_battery, size, remaining_banks)

def find_largest_avaliable_battery(avaliable_bank):
    current_max = 0
    current_max_index = 0
    i = 0
    for battery in avaliable_bank:
        if (int(battery) > current_max):
            current_max = int(battery)
            current_max_index = i
        i = i + 1
    return current_max_index

banks = get_banks("test.txt")
max_jg = []
for bank in banks:
    max_battery = int(find_largest_battery("", 12, bank))
    max_jg.append(max_battery)
    print(bank + " " + str(max_battery))

print(sum(max_jg))