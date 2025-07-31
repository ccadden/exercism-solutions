//
// This is only a SKELETON file for the 'Linked List' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class LinkedList {
  head = undefined;
  tail = undefined;
  length = 0;

  //end of list
  push(value) {
    this.length++;
    const newNode = new Node(value);

    if (!this.tail || !this.head) {
      this.tail = newNode;
      this.head = newNode;
      return;
    }

    this.tail.next = newNode;
    newNode.prev = this.tail;
    this.tail = newNode;
    return;
  }

  //end of list
  pop() {
    if (!this.tail) return;

    this.length--;

    const deadNode = this.tail;

    if (this.tail.prev) {
      this.tail.prev.next = undefined;
    }

    this.tail = this.tail.prev;

    deadNode.prev = undefined;

    return deadNode.value;
  }

  //remove from beginning
  shift() {
    if (!this.head) return;

    this.length--;

    const deadNode = this.head;
    if (this.head.next) {
      this.head.next.prev = undefined;
    }

    this.head = this.head.next;

    deadNode.next = undefined;

    return deadNode.value;
  }

  //add to beginning
  unshift(value) {
    this.length++;
    const newNode = new Node(value);

    if (!this.head || !this.tail) {
      this.head = newNode;
      this.tail = newNode;
      return;
    }

    this.head.prev = newNode;
    newNode.next = this.head;

    this.head = newNode;
  }

  delete(value) {
    let current = this.head;

    while (current) {
      if (current.value === value) {
        this.length--;

        if (this.length === 0) {
          this.head = undefined;
          this.tail = undefined;
          return;
        }

        if (current === this.head) {
          this.head = current.next;
          if (this.head) {
            this.head.prev = undefined;
          }

          return;
        }

        if (current === this.tail) {
          this.tail = current.prev;

          if (this.tail) {
            this.tail.next = undefined;
          }

          return;
        }

        current.prev.next = current.next;
        current.next.prev = current.prev;

        return;
      }

      current = current.next;
    }
  }

  count() {
    return this.length;
  }
}

class Node {
  value;
  next;
  prev;

  constructor(value) {
    this.value = value;
    this.next = undefined;
    this.prev = undefined;
  }
}
