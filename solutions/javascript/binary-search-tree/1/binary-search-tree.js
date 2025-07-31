//
// This is only a SKELETON file for the 'Binary Search Tree' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class BinarySearchTree {
  constructor(value) {
    this.value = value;
    this._left = undefined;
    this._right = undefined;
  }

  get data() {
    return this.value;
  }
  get right() {
    return this._right;
  }

  get left() {
    return this._left;
  }

  insert(n) {
    if (n <= this.value) {
      if (this._left) {
        this._left.insert(n);
      } else {
        this._left = new BinarySearchTree(n);
      }
    } else {
      if (this._right) {
        this._right.insert(n);
      } else {
        this._right = new BinarySearchTree(n);
      }
    }
  }

  each(fn) {
    if (this._left) {
      this._left.each(fn);
    }

    fn(this.value);

    if (this._right) {
      this._right.each(fn);
    }
  }

  ceil(n) {
    let ceilValue = undefined;

    function walk(node) {
      if (!node) return;

      if (node.value <= n) {
        walk(node._right);
      } else {
        ceilValue = node.value;
        walk(node._left);
      }
    }

    walk(this);

    return ceilValue;
  }

  floor(n) {
    let floorValue = undefined;

    function walk(node) {
      if (!node) return;

      if (node.value >= n) {
        walk(node._left);
      } else {
        floorValue = node.value;
        walk(node._right);
      }
    }

    walk(this);

    return floorValue;
  }
}
