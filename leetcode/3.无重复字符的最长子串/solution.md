# solution
[3. 无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

滑动窗口：
- 两层循环
  - 第一层移动左指针，移动后删掉存数
  - 第二层移动右指针，边界是不越界并且没存过数
- 右指针移动完记录当前最大长度