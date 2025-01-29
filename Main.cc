#include <cstdlib>
#include <iostream>
#include "Sushi.hh"

int main(int argc, char *argv[]) {
    UNUSED(argc);
    UNUSED(argv);

    Sushi sushi;

    // Read configuration file from $HOME/sushi.conf
    const char *home = std::getenv("HOME");
    if (home) {
        std::string configPath = std::string(home) + "/sushi.conf";
        sushi.read_config(configPath.c_str(), true);
    }

    while (true) {
        std::cout << Sushi::DEFAULT_PROMPT;
        
        std::string command = sushi.read_line(std::cin);
        if (command == nullptr) continue;

        sushi.store_to_history(command);
        sushi.show_history();
    }

    return EXIT_SUCCESS;
}
