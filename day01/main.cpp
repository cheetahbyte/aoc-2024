#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <sstream>

#define IntVectorReference std::vector<int>

bool isDigit(const std::string &str)
{
    return !str.empty() && std::all_of(str.begin(), str.end(), ::isdigit);
}

std::pair<IntVectorReference, IntVectorReference> readFileLines()
{
    std::vector<int> left, right;
    //
    std::ifstream inputFile;
    inputFile.open("input.txt");
    std::string fileContent;
    std::string line;
    while (std::getline(inputFile, line))
    {
        std::istringstream iss(line);
        int num1, num2;

        if (iss >> num1 >> num2) 
        {
            left.push_back(num1);
            right.push_back(num2);
        }
        else
        {
            std::cerr << "Invalid line skipped: " << line << std::endl;
        }
    }
    std::sort(left.begin(), left.end());
    std::sort(right.begin(), right.end());
    return {left, right};
}

void part2(IntVectorReference left, IntVectorReference right)
{
    auto similarityScore = 0;
    for (const auto element: left) {
        similarityScore += std::count(right.begin(), right.end(), element) * element;
    }
    std::cout << "Part 2: " << similarityScore << std::endl;
}

void part1(IntVectorReference left, IntVectorReference right)
{
    auto distance = 0;
    for (std::size_t i = 0; i < left.size(); i++)
    {
        const int lvalue = left[i];
        const int rvalue = right[i];
        
        if (lvalue < rvalue)
            distance += rvalue - lvalue;
        else
            distance += lvalue - rvalue;
    }

    std::cout << "Part 1: " << distance << std::endl;
}


int main()
{
    const auto [left, right] = readFileLines();
    part1(left, right);
    part2(left, right);
    return 0;
}