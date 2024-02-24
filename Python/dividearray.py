nums = [1,3,4,8,7,9,3,5,20]
list = []
list2 = []
k = 2


print(len(nums)%3)
nums.sort()

if len(nums)%3 == 0 :
    n = 0

    for i in nums :
        n += 1
        list2.append(i)
        if n == 3 :
            list.append(list2)
            list2 = []
            n = 0

print(list)
for i in range(len(list)) :
    if list[i][2] - list[i][0] > k :
        print("False")
    else : 
        print("true")

print(list)
