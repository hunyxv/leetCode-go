## 最长回文子串

给你一个字符串 s，找到 s 中最长的回文子串。

 

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-palindromic-substring
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


## 中心扩展算法
边界情况即为子串长度为 11 或 22 的情况。我们枚举每一种边界情况，并从对应的子串开始不断地向两边扩展。如果两边的字母相同，我们就可以继续扩展，例如从 `P(i+1,j-1)P(i+1,j−1)` 扩展到 `P(i,j)P(i,j)`；如果两边的字母不同，我们就可以停止扩展，因为在这之后的子串都不能是回文串了。

聪明的读者此时应该可以发现，「边界情况」对应的子串实际上就是我们「扩展」出的回文串的「回文中心」。方法二的本质即为：我们枚举所有的「回文中心」并尝试「扩展」，直到无法扩展为止，此时的回文串长度即为此「回文中心」下的最长回文串长度。我们对所有的长度求出最大值，即可得到最终的答案。


## 动态规划
用动态规划的思想来分析，先拆分子问题：

定义s为一个字符串，`0...j` 为字符串中各字符的索引，如果从 `s[0....j]` 是一个回文子串则 `s[1...j-1]` 必定是一个回文子串，且 `s[0]==s[j-1]`.
定义 `p[i][j]` 为“`s[i...j]` 是否为回文字符串”如果是 `p[i][j]=true`，否则 `p[i][j]=false`。 则 `p[i][j]=p[i+1][j-1] && s[i]==s[j]` 。
考虑边界情况，上述递推式子中用到了 `i+1` 及 `j-1` ,则 `i,j` 中间至少有一个值时 `i+1,j-1` 正好向中间这个值，也就是 `s[i...j]` 这个字符串中至少有3个字符，现在我们考虑小于三个字符的情况，这也就是问题的边界。
如果有 1 个字符，`i==j` 此时 `p[i][j]=true` .
如果有 2 个字符，i+1=j此时如果 `s[i]==s[j]` 则 `p[i][j]=true` 否则 `p[i][j]=false`.
至此用动态规划的方法问题分析完毕。
