# PE 16 Power Digit Sum

if __name__ == '__main__':
    num = pow(2, 1000)
    num = str(num)
    result = 0
    for d in num:
        result += int(d)
    print(result)
