// 1.1 from CTCI.
// Implemented it using a Unordered map from C++11

#include <unordered_map>
#include <iostream>
#include <string>

using namespace std;

bool is_unique(string s) {
    unordered_map<char, bool> uniques;
    for(int i = 0; i < (int)s.length(); i++) {
        if (uniques[s[i]] == true) {
            return false;
        }
        uniques[s[i]] = true;
    }

    return true;
}

int main() {
    string s = "so_uniq";

    if (is_unique(s)) {
        cout << "Unique!" << endl;
    } else {
        cout << "NOT Unique!" << endl;
    }

    string s2 = "omgqtfbbq";

    if (is_unique(s2)) {
        cout << "Unique!" << endl;
    } else {
        cout << "NOT Unique!" << endl;
    }

    return 0;
}
