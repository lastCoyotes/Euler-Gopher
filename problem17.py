def tripleDigitInterpretor(num: str) -> str:
    table = {
        "0": "",
        "1": "One",
        "2": "Two",
        "3": "Three",
        "4": "Four",
        "5": "Five",
        "6": "Six",
        "7": "Seven",
        "8": "Eight",
        "9": "Nine",
        "10": "Ten",
        "11": "Eleven",
        "12": "Twelve",
        "13": "Thirteen",
        "14": "Fourteen",
        "15": "Fifteen",
        "16": "Sixteen",
        "17": "Seventeen",
        "18": "Eighteen",
        "19": "Nineteen",
        "20": "Twenty",
        "30": "Thirty",
        "40": "Forty",
        "50": "Fifty",
        "60": "Sixty",
        "70": "Seventy",
        "80": "Eighty",
        "90": "Ninety",
    }

    result = ""
    numlist = []
    numlist.extend(num)
    if numlist[1] == "1":
        numlist = [numlist[0], numlist[1] + numlist[2]]
    else:
        numlist[1] = numlist[1] + "0"

    do_once = True
    for x in numlist:
        if int(x) == 0:
            pass
        else:
            if do_once and int(numlist[0]) > 0:
                result += table[x] + " Hundred "
                do_once = False
            else:
                result += table[x] + " "

    return result


def numberToWords(num: int) -> str:
    table = {
        1: "Thousand",
        2: "Million",
        3: "Billion"
    }

    if num == 0:
        return "Zero"

    test_str = str(num)
    remainder = len(test_str) % 3
    if remainder == 2:
        test_str = "0" + test_str
    elif remainder == 1:
        test_str = "00" + test_str
    else:
        pass

    split_string = [test_str[idx: idx + 3] for idx in range(0, len(test_str), 3)]
    string_list = []

    for x in split_string:
        string_list.append(tripleDigitInterpretor(x))

    string_list.reverse()

    digit_grouped = []
    index_ctr = 0
    for y in string_list:
        if index_ctr > 0 and y != '':
            digit_grouped.append(y + table[index_ctr] + " ")
        else:
            digit_grouped.append(y.strip())
        index_ctr += 1

    digit_grouped.reverse()
    solution = ""
    for z in digit_grouped:
        solution += z

    return solution.strip()

if __name__ == '__main__':
    result = 0
    for x in range(1, 1001, 1):
        print(numberToWords(x))
        result += len(numberToWords(x).replace(" ", ""))
    print(result + 2673)
