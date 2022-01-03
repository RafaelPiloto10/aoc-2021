#include <fstream>
#include <iostream>
#include <string>

using namespace std;

struct Node {
public:
  int data;
  Node *next;

  Node(int data) {
    this->data = data;
    next = NULL;
  }

  void add(int data);
};

void Node::add(int data) {
  Node *n = new Node(data);
  if (next == NULL) {
	next = n;
	return;
  }

  Node *current = next;
  while (current->next != NULL) {
    current = current->next;
  }
  current->next = n;
}

int main() {

  ifstream myfile("./input.txt");
  string line;

  Node *head = NULL;

  if (myfile.is_open()) {
    while (getline(myfile, line)) {
      if (line != "") {
        int data = stoi(line);
		  if (head == NULL)
			  head = new Node(data);
		  else
        	head->add(data);
      }
    }

    myfile.close();
  } else {
    cout << "Unable to open file\n";
  }

  Node *current = head;
  int prev = INT_MAX; 
  int count = 0;
  while (current != NULL) {
	  if (current->data > prev) {
		  count++;
	  }
	prev = current->data;
    current = current->next;
  }

  cout << "Total Count: " << count << "\n";

  return 0;
}
