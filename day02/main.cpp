#include <iostream>
#include <fstream>
#include <sstream>
#include <vector>
#include <cmath>
#include <algorithm>
#include <iterator>

using VecInt = std::vector<int>;
using VecVecInt = std::vector<VecInt>;

VecVecInt readFileByLines(const std::string& filename = "input.txt") {
    VecVecInt data;
    std::ifstream file(filename);
    std::string line;

    while (std::getline(file, line)) {
        std::istringstream iss(line);
        data.emplace_back(VecInt{std::istream_iterator<int>{iss}, {}});
    }
    return data;
}

bool isSafe(const VecInt& row) {
    bool increasing = true, decreasing = true;
    for (size_t i = 1; i < row.size(); ++i) {
        int diff = std::abs(row[i] - row[i - 1]);
        if (diff < 1 || diff > 3) return false;
        increasing &= row[i] >= row[i - 1];
        decreasing &= row[i] <= row[i - 1];
    }
    return increasing || decreasing;
}

bool isSafeWithRemoval(const VecInt& row) {
    if (isSafe(row)) return true;
    for (size_t i = 0; i < row.size(); ++i) {
        VecInt temp(row.begin(), row.end());
        temp.erase(temp.begin() + i);
        if (isSafe(temp)) return true;
    }
    return false;
}

void solve(const VecVecInt& data, bool (*checkFunc)(const VecInt&), const std::string& partName) {
    int count = std::count_if(data.begin(), data.end(), checkFunc);
    std::cout << partName << ": " << count << '\n';
}

int main() {
    VecVecInt data = readFileByLines();
    solve(data, isSafe, "Part 1");
    solve(data, isSafeWithRemoval, "Part 2");
    return 0;
}