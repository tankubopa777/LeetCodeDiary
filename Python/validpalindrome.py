from typing import Optional
class Solution:
    def isPalindrome(self, s: str) -> bool:
        y = ""
        n = ""
        for i in s.lower() :
            if i.isalnum() :
                y += i

        for i in s[::-1].lower() :
            if i.isalnum() :
                n += i

        if y == n :
            return True
        else : 
            return False

    def isNumPalindrome(self,  head: Optional[ListNode]) -> bool:
        print(5555)

Solution.isNumPalindrome()
