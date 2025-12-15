class My_Range:
    def __init__(self, min, max):
        self.min = min
        self.max = max

    def __repr__(self):
        return f"({self.min}, {self.max})"

    def getMin(self):
        return self.min

    def getMax(self):
        return self.max

def getRanges(fileName):
    my_file = open(fileName)
    my_range_strings = []
    for line in my_file:
        clean_line = line.rstrip('\r\n')
        my_range_strings.extend(clean_line.split(","))

    my_ranges = []
    for my_range_string in my_range_strings:
        values_string = my_range_string.split("-")
        my_ranges.append(My_Range(int(values_string[0]), int(values_string[1])))
    
    return my_ranges

def getInvalidIdsInRange(my_range):
    invalid_numbers = []
    for n in range(my_range.getMin(), my_range.getMax() + 1):
        split_point = int(len(str(n)) / 2)
        n_len = len(str(n))
        first_n = str(n)[split_point:n_len]
        last_n = str(n)[0:split_point]
        if first_n == last_n:
            invalid_numbers.append(n)

    return invalid_numbers

def _getInvalidIdsInRange(my_range):
    invalid_numbers = []
    for n in range(my_range.getMin(), my_range.getMax() + 1):
        # print("Number", n)
        partial_string = str(n)[:1]
        remaining_string = str(n)[1:]
        if(isRepeatedString(partial_string, remaining_string)):
            print("Invalid Number", n)
            invalid_numbers.append(n)

    return invalid_numbers

def isRepeatedString(partial_string, remaining_string):
    if len(partial_string) > len(remaining_string):
        return False

    if(isOnlyPattern(partial_string, remaining_string)):
        return True
    
    new_partial = partial_string + remaining_string[0:1]
    new_remaining = remaining_string[1:len(remaining_string)]
    return isRepeatedString(new_partial, new_remaining)

def isOnlyPattern(partial_string, remaining_string):
    if(len(remaining_string) == 0):
        return True
    if(partial_string == remaining_string[0:len(partial_string)]):
        return isOnlyPattern(partial_string, remaining_string[len(partial_string):])

    return False
    
ranges = getRanges("my_input.txt")
invalid_numbers = []
for my_range in ranges:
    invalid_numbers.extend(_getInvalidIdsInRange(my_range))

#print(invalid_numbers)
print(sum(invalid_numbers))