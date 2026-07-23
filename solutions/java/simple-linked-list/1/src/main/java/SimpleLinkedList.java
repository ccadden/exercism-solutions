import java.util.List;
import java.util.ArrayList;
import java.util.NoSuchElementException;

class SimpleLinkedList<T> {
    private int size;
    private Node<T> head;

    SimpleLinkedList() {
        this.size = 0;
        this.head = null;
    }

    SimpleLinkedList(T[] values) {
        for(T value:values) {
            Node<T> newNode = new Node(value);

            if(head != null) {
                newNode.next = this.head;
            }

            this.head = newNode;
            this.size++;
        }
    }

    Node<T> getHead() {
        return this.head;
    }

    void push(T value) {
        Node<T> newNode = new Node<T>(value);

        if(this.head == null) {
            this.head = newNode;
        } else {
            newNode.next = this.head;
            this.head = newNode;
        }

        this.size++;
    }

    T pop() {
        if(size == 0) {
            throw new NoSuchElementException();
        }

        Node<T> deadNode = this.head;
        this.head = this.head.next;
        deadNode.next = null;


        this.size--;

        return deadNode.value;
    }

    T peek() {
        if(size == 0) {
            throw new NoSuchElementException();
        }

        return this.head.value;
    }

    void reverse() {
        Node<T> lastNode = null;
        Node<T> currNode = this.head;

        while(currNode != null) {
            Node<T> nextNode = currNode.next;
            currNode.next = lastNode;
            lastNode = currNode;
            currNode = nextNode;
        }

        this.head = lastNode;
    }

    List<T> toList() {
        ArrayList<T> list = new ArrayList<T>(size);

        Node<T> currNode = this.head;

        for(int i = 0; i < size; i++) {
            list.add(currNode.value);
            currNode = currNode.next;
        }

        return list;
    }

    int size() {
        return size;
    }
}

class Node<T> {
    public T value;
    public Node<T> next;

    Node(T value) {
        this.value = value;
        this.next = null;
    }

    void setNext(Node<T> node) {
        this.next = node;
    }
}
