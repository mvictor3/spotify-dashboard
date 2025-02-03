#include <cstdlib>
#include <iostream>
#include "Sushi.hh"

Sushi my_shell; // New global var

int main(int argc, char *argv[]) {
    UNUSED(argc);
    UNUSED(argv);

    // DZ: Moved to globals (not an error)
    // Sushi sushi;

    // Read configuration file from $HOME/sushi.conf
    const char *home = std::getenv("HOME");
    if (home) {
        std::string configPath = std::string(home) + "/sushi.conf";
	// Must check the return value and exit of necessary
        my_shell.read_config(configPath.c_str(), true);
    }

    while (true) {
        std::cout << Sushi::DEFAULT_PROMPT;
        
        std::string command = my_shell.read_line(std::cin);
	// DZ: Wrong comparison
        if (command == "") continue;
        //if (command == nullptr) continue;

        my_shell.store_to_history(command);
        my_shell.show_history();
    }

    return EXIT_SUCCESS;
}
