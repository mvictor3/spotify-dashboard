#include <sstream> // DZ: Needed by std::istringstream
#include <iostream>
#include "Sushi.hh"

const size_t Sushi::MAX_INPUT = 256;
const size_t Sushi::HISTORY_LENGTH = 10;
const std::string Sushi::DEFAULT_PROMPT = "sushi> ";

// DZ: `history` is a member of the Sushi class
// History storage using a basic array-like approach
// std::string history[Sushi::HISTORY_LENGTH];
// DZ: Make it a member variable
size_t history_size = 0;

// Reads a line from input, trims whitespace, and handles errors
std::string Sushi::read_line(std::istream &in) {
    char buffer[MAX_INPUT + 1];  // Buffer to hold input
    if (!in.getline(buffer, MAX_INPUT + 1)) { 
        std::perror("Error reading input");
	// DZ: Wrong return
	return "";
        // return nullptr;
    }

    // Handle long input
    // DZ: signedness issue fixed
    if (in.gcount() > static_cast<std::streamsize>(MAX_INPUT)) {
        std::cerr << "Line too long, truncated." << std::endl;
        in.ignore(10000, '\n');  // Clear excess characters
	// DZ: Why 10000?
    }

    std::string line(buffer);

    // DZ: The spaces whould not be trimmed
    // Trim leading spaces
    size_t start = line.find_first_not_of(" \t\r\n");
    if (start == std::string::npos) {
      // DZ: Wrong return
      return "";
      //return nullptr;  // Only whitespace, discard
    }

    // Trim trailing spaces
    size_t end = line.find_last_not_of(" \t\r\n");
    line = line.substr(start, end - start + 1);

    return line;
}

// Stores commands in history (simple shifting array approach)
void Sushi::store_to_history(std::string line) {
    if (line.empty()) return;

    // Shift history down if full
    if (history_size == HISTORY_LENGTH) {
      // DZ: VERY inefficient; history.pop_front();
        for (size_t i = HISTORY_LENGTH - 1; i > 0; --i) {
            history[i] = history[i - 1];
        }
    } else {
       // DZ: VERY inefficient and not needed
       history_size++;
        for (size_t i = history_size - 1; i > 0; --i) {
            history[i] = history[i - 1];
        }
    }

    // DZ: history.emplace_back(line);
    history[0] = line;  // Insert new entry at the front
}

// Reads a config file and stores lines in history
bool Sushi::read_config(const char *fname, bool ok_if_missing) {
  // DZ: This is a C way, not C++ way
    FILE *file = fopen(fname, "r");
    if (!file) {
      if (!ok_if_missing) {
	// DZ: Wrong use of perror
	// std::perror("Error opening config file");
	std::perror(fname);
      }
        return ok_if_missing;
    }

    char buffer[MAX_INPUT + 1];
    // DZ: Very inefficient
    while (fgets(buffer, sizeof(buffer), file)) {
        std::string line(buffer);
	// DZ: The original code cannot be compiled
	// DZ: error: cannot bind non-const lvalue reference of type ‘std::istream&’
	std::istringstream iss(line);
        std::string trimmed = read_line(iss);
	// DZ: Wrong comparison
	// DZ: read_config whould not store to history
        if (trimmed != "") store_to_history(trimmed);
        // if (trimmed != nullptr) store_to_history(trimmed);
    }

    fclose(file);
    return true;
}

// Displays stored history
void Sushi::show_history() {
  for (size_t i = 0; i < history_size; ++i) {
        std::cout << i + 1 << "  " << history[i] << std::endl;
    }
}

void Sushi::set_exit_flag()
{
  // To be implemented
}

bool Sushi::get_exit_flag() const
{
  return false; // To be fixed
}
