class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        dp = [[0] * n for i in range(m)]
        print(dp)
        for i in range(m):
            for j in range(n):
                if i == 0 or j == 0:
                    dp[i][j] = 1
                else:
                    dp[i][j] += (dp[i-1][j] + dp[i][j-1])
        print(dp)
        return dp[m-1][n-1]

def main():
    s = Solution()
    case1 = s.uniquePaths(3,7)
    case1_answer = 28

    print(case1)
    print(case1_answer)
    print(case1 == case1_answer)

    case2 = s.uniquePaths(3,2)
    case2_answer = 3

    print(case2 == case2_answer)



if __name__ == "__main__":
    main() 