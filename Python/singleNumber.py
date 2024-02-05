nums = [1]

onecount = 0
listpow = []
for i in nums :
    if i not in listpow:
        listpow.append(i)

for i in listpow :
    if nums.count(i) == 1 :
        print (f"{i}")

