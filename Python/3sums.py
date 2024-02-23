# nums = [-1,0,1,2,-1,-4]

# nums.sort()

# result = []

# for i in range(len(nums)-2):
#     if i > 0 and nums[i] == nums[i-1]:
#         continue
#     l, r = i+1, len(nums)-1
#     while l < r:
#         s = nums[i] + nums[l] + nums[r]
#         if s < 0:
#             l +=1 
#         elif s > 0:
#             r -= 1
#         else:
#             result.append((nums[i], nums[l], nums[r]))
#             while l < r and nums[l] == nums[l+1]:
#                 l += 1
#             while l < r and nums[r] == nums[r-1]:
#                 r -= 1
#             l += 1; r -= 1

# print(result)

from typing import List
class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        print(nums)

nums = [-1,0,1,2,-1,-4]
Solution().threeSum(nums)

# abs(a[i] - x) + abs(a[1] - x) + ... abs(a[a.length - 1] - x)
# a = [2, 4, 7] the result is 4
def minAbsSumnums(nums):
    min_sum = float('inf')  # Initialize with a large value
    closest_element = None

    for listNumber in range(len(a)):
        current_sum = 0
        for sumNumber in range(len(a)):
            difference = abs(a[sumNumber] - a[listNumber])
            current_sum += difference

        if current_sum < min_sum:
            min_sum = current_sum
            closest_element = a[listNumber]
        elif current_sum == min_sum and a[listNumber] < closest_element:  # Tiebreaker for smaller element
            closest_element = a[listNumber]

    return closest_element


a = [2, 3]
print(minAbsSumnums(a))  # 2