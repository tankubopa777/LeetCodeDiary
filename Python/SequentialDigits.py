def SequentialDigits(low, hight) :
    listnumber = []
    numbers = len(str(low))
    numbersinloop = ""

    
    for xsum in range(len(str(low))) :
        numbersinloop += str(xsum)

    lownumber = int(numbersinloop)
    while not lownumber + int(numbers * "1") > low and lownumber + int(numbers * "1") < hight :
        lownumber += int(numbers * "1")
        
            
    print(lownumber)

    while lownumber + int(numbers * "1") < hight :
        if str(lownumber + int(numbers * "1"))[-1::-1][0] == "9" :
            lownumber += int(numbers * "1")
            listnumber.append(lownumber)
            numbers += 1
        else :
            lownumber += int(numbers * "1")
            listnumber.append(lownumber)



    print(listnumber)

