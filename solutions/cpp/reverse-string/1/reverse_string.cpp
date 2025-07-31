#include "reverse_string.h"

namespace reverse_string {
    std::string reverse_string(std::string input) {
        for(unsigned int i=0; i < input.size()/2; i++) {
            unsigned int j = input.size() - 1 - i;
            std::swap(input[i], input[j]);
        }

        return input;
    }
}  // namespace reverse_string
