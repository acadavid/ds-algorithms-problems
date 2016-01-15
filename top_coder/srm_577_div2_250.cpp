#include <vector>
#include <list>
#include <map>
#include <set>
#include <queue>
#include <deque>
#include <stack>
#include <bitset>
#include <algorithm>
#include <functional>
#include <numeric>
#include <utility>
#include <sstream>
#include <iostream>
#include <iomanip>
#include <cstdio>
#include <cmath>
#include <cstdlib>
#include <ctime>

using namespace std;


class EllysNewNickname {

    bool isVowel(char c) {
        return (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'y');
    }

public:
    int getLength(string nickname) {
        int length = 0;
        int i;
        for (i=0;i<(int)nickname.size();i++) {
            if (isVowel(nickname[i])) {
                length += 1;
                for (int j=i+1; j<(int)nickname.size();j++) {
                    if (!isVowel(nickname[j])) {
                        i=j-1;
                        break;
                    }
                    i=j;
                }
            } else {
                length += 1;
            }
        }

        return length;
    }
} elly;

int main() {
    cout << elly.getLength("eagaeoppooaaa") << endl;
    cout << elly.getLength("esprit") << endl;
    cout << elly.getLength("naaaaaaaanaaaanaananaaaaabaaaaaaaatmaaaaan") << endl;
    cout << elly.getLength("ayayayayayayayaya") << endl;
    return 0;
}
