def collatzSequence(n: int) -> int:
    data = {1: [4, 1]}
    result = 0
    maxChainLen = 0
    for x in range(2, n+1, 1):
        if x not in data:
            # stack: list[list[int]] where [key, ptr, len]
            # we wanna have stack empty each time
            stack = []
            ptr = x
            while ptr not in data:
                if x & 1:  # if odd
                    ptr = (3*x)+1
                else:
                    ptr = x//2
                stack.append([x, ptr])
                x = ptr
            chainLength = data[ptr][1]

            while stack:
                a, b = stack.pop()
                chainLength += 1
                data.update({a: [b, chainLength]})
                # because we are calculating chain lengths here,
                # the one at the very bottom of the stack will be the largest anyway
                # this should retrieve the last value popped outside of the while loop
                # if while loops are not locally scoped
            if data[a][1] > maxChainLen:
                maxChainLen = data[a][1]
                result = a

    # print(data)
    # print(len(data))      n=999999 -> 2_168_611 keys in data
    # print(data[837799])   chainLength for 837799 is 525
    return result


if __name__ == '__main__':
    print(collatzSequence(999999))


# given integer, n, find the integer with the longest Collatz Sequence.
# end when you hit "1"
# you include the number you started with
# 1's length is 4: [1, 4, 2, 1]
# 2's length is 2: [2, 1]
# 3's length is 8: [3, 10, 5, 16, 8, 4, 2, 1]
# 4's length is 3: [4, 2, 1]
# as we fill our memory with key value pairs. we will check if we have already found where that number lies in the seq
# for example, we already know now that 4 points to 2.
#       A) k: [ptr, len] vs
#       B) k: len
#       A would let me stop recalculating where each k points to, because then i only retrieve ptr
#       B would save more space on memory but ill have to calculate ptr each time i come across K
#       k: len
#       key is current number (this way no duplicates)
#       len is length from "1" node
#       using dict because i need to both check it often AND retrieve from it.
#       first term may not be necessary but will try later with an empty dict
#                               x isn't in our dict yet, so we need to start calculating some new nodes
#                               however we dont know how far away x is from 1 in the sequence,
#                               so we dont want to update the dict twice
#                               we could create a list of tuples to keep track of the data..?
#                               order doesnt matter here, but it can make things convenient,
#                               but once a number already in data is found,
#                               we know how far away it is based off that number's len (data[k][1] position)
#                               all prior nodes will just be += 1 (if we ordered our temp data),
#                               then we know each new nodes' len from the 1 node
#                               **using a stack because utilizing FILO property.
#               example to test if my flow works:
#               1 and 2 are already in data by now, so 4 will have also been logged as well too
#               data = {1:[4,4], 2:[1,2],4:[2,3]}
#               x = 3, 3 is not in data
#               stack = [[3, 10]]
#               x = ptr (10) so x is now 10.
#               10 is still not in data
#               stack = [[3, 10],[10,5]]
#               x -> 5 ... -> 16 -> 8
#               stack = [[3,10],[10,5],[5,16],[16,8],[8,4]]
#               ptr is currently 4 so break out of the while loop
#               chainLength = data[4][1] # retrieve the chainLength value from 4
#
#               while stack:
#                   x, y = stack.pop()
#                   chainLength += 1
#                   data.update({x: [y, chainLength]})
#
#               so as we keep building up our dictionary, where do we check for the max value to what key?
# result = 0
# maxChainLen = 0
# for x in range(n):
#     if x not in data:
#         # stack: list[list[int]] where [key, ptr, len]
#         # we wanna have queue empty each time
#         stack = []
#         while ptr not in data:
#             if x & 1: # if odd
#                 ptr = (3*x)+1
#             else:
#                 ptr = x//2
#             stack.push([x,ptr,None])
#             x = ptr
#         chainLength = data[ptr][1]
#         while stack:
#             a, b = stack.pop()
#             chainLength += 1
#             data.update({a: [b, chainLength]})
#             # because we are calculating chain lengths here,
#             # the one at the very bottom of the stack will be the largest anyway
#             # this should retrieve the last value popped outside of the while loop if while loops are not locally scoped
#         if data[a][1] > maxChainLen:
#         maxChainLen = data[a][1]
#         result = a
# we will need to keep track of the data[k][1] value and recall which one Key was the maximum (not the maximum seq itself)
# result = 0
# maxChainLen = 0
# if data[k][1] > maxChainLen:
#     maxChainLen = data[k][1]
#     result = k
