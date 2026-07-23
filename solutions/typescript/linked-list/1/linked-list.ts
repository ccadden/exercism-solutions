export class LinkedList<T> {
  private head: Node<T> | null = null;
  private tail: Node<T> | null = null;
  private size: number = 0;

  public push(element: T) {
    const newNode = new Node(element);

    const oldHead = this.head;

    newNode.next = oldHead;

    if (oldHead) {
      oldHead.prev = newNode;
    }

    this.head = newNode;

    if (this.tail === null) {
      this.tail = newNode;
    }

    this.size++;
  }

  public pop(): T | null {
    const deadNode = this.head;

    if (deadNode === null) {
      return null;
    }

    this.head = deadNode.next;
    deadNode.next = null;

    if (this.head === null) {
      this.tail = null;
    }

    this.size--;

    return deadNode.value;
  }

  // remove last node
  public shift(): T | null {
    if (this.tail === null) {
      return null;
    }

    const deadNode = this.tail;

    this.tail = deadNode.prev;

    if (deadNode.prev !== null) {
      deadNode.prev.next = null;
    }

    if (this.head === deadNode) {
      this.head = null;
    }

    this.size--;
    return deadNode.value;
  }

  // append to end of list
  public unshift(element: T): void {
    const newNode = new Node<T>(element);

    if (this.tail === null) {
      this.head = this.tail = newNode;
      this.size++;
      return;
    }

    newNode.prev = this.tail;
    this.tail.next = newNode;
    this.tail = this.tail.next;

    this.size++;

    return;
  }

  public delete(element: T) {
    if (this.size === 0) {
      return;
    }

    let currNode = this.head;

    while (currNode !== null) {
      if (currNode.value === element) {
        if (currNode.prev) {
          currNode.prev = currNode.next;
        }

        if (currNode.next) {
          currNode.next = currNode.prev;
        }

        if (currNode === this.head) {
          this.head = null;
        }

        if (currNode === this.tail) {
          this.tail = null;
        }

        this.size--;
      }

      currNode = currNode.next;
    }
  }

  public count(): number {
    return this.size;
  }
}

class Node<T> {
  readonly value: T;
  next: Node<T> | null = null;
  prev: Node<T> | null = null;

  constructor(value: T) {
    this.value = value;
  }
}
