
def pingjie(m, c):
    n = len(c)
    dp = [[0 for j in range(m)] for i in range(n)] # dp: (m, n)
    for i in range(n):
        dp[i][0] = 1
    for i in range(n):
        for j in range(m):
            if j-c[i-1] >= 0:
                dp[i][j] = dp[i-1][j] + dp[i][j-c[i-1]]
            
    
    return dp[n-1][m-1]
    
print(pingjie(4, [1, 2]))