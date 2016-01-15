#include <iostream>
#include <vector>
#include <algorithm>
#include <cstdlib>

using namespace std;

class FilipTheFrog {

public:
    int countReachableIslands(vector<int> pos, int l) {
        int reachable = 1;
        int first = pos[0];
        
        sort(pos.begin(), pos.end());
        int pos_length = pos.size();

        vector<int>::iterator it;
        int st = find(pos.begin(), pos.end(), first) - pos.begin();
        int st_bu = st;

        int d;
        while (true) {
            if ( ! (st >= 1)) {
                break;
            }
            d = abs(pos[st-1] - pos[st]);
            if (d <= l) {
                reachable += 1;
                st -= 1;
            } else {
                break;
            }
        }

        st = st_bu;
        while (true) {
            if ( ! (st <= pos_length-2)) {
                break;
            }

            d = abs(pos[st] - pos[st+1]);
            if (d <= l) {
                reachable += 1;
                st += 1;
            } else {
                break;
            }
        }

        return reachable;

    }
} frog;

int main() {
    vector<int> words { 0 };
    int reachable = frog.countReachableIslands(words, 2);

    cout << reachable << endl;
    return 0;
}
