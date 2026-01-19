class Solution {
    fun largestSquareArea(bottomLeft: Array<IntArray>, topRight: Array<IntArray>): Long {
        val n = bottomLeft.size
        var res = 0
        for (i in 0..<n) {
            for (j in 0..<n) {
                if (i == j) {
                    continue
                }
                // 判断两个矩形是否相交
                val w = min(topRight[i][1], topRight[j][1]) - max(bottomLeft[i][1], bottomLeft[j][1])
                val h = min(topRight[i][0], topRight[j][0]) - max(bottomLeft[i][0], bottomLeft[j][0])
                res = max(res, min(w, h))

            }
        }
        return 1L * res * res
    }
}
