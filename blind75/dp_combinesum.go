package blind

func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1 // 注意这个初始化的值，考虑case []int{2},2,如果target为0说明有一种路可以通
	for i := 1; i < target+1; i++ {
		for _, n := range nums {
			if i >= n {
				dp[i] += dp[i-n]
			}
		}
	}

	return dp[target]
}

// 同爬楼梯类似
// 同找零钱类似,但是找零钱需要求min

// class Solution {
//     public int combinationSum4(int[] nums, int target) {
//         // 定义dp[i]为爬上i层的方案数
//         // 最后一次可以决策爬 nums[j]层，则dp[i] = sum(dp[i - nums[j]])
//         int[] dp = new int[target + 1];
//         dp[0] = 1;
//         for (int i = 1; i <= target; i++) {
//             for (int n : nums) {
//                 if (i >= n) {
//                     dp[i] += dp[i - n];
//                 }
//             }
//         }
//         return dp[target];
//     }
// }
