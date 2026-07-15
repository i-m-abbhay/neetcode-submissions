class Solution {
public:
    bool hasDuplicate(vector<int>& nums) {
        unordered_set<int> et;
        for (auto it = nums.begin();it!=nums.end(); it++){
            if(et.find(*it)!=et.end()){
                return true;
            }
            et.insert(*it);
        }
        return false;
    }
};
