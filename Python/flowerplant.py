def canPlaceFlowers(flower, nums) :
    havePlant = len(flower) - flower.count(1)
    canPlant = nums * 3
    if havePlant > canPlant :
        return True
    else :
        return False


def moveZeroes(number) : 
    print(sorted(number))
    for numbers in number :
        if numbers == 0 :
            number.remove(numbers)
            number.append(0)
    print(number)

def isSubsequence(s, t) :
    s2 = ""
    for alphabeth2 in range(len(t)) :
        if s[alphabeth2] in s :
            s2 += s[alphabeth2]
    if s == s2 :
        print(s)
        print(s2)
    else : 
        print(s)
        print(s2)

        
print(isSubsequence("ab", "baab"))