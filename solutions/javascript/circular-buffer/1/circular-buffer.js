//
// This is only a SKELETON file for the 'Circular Buffer' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

class CircularBuffer {
  constructor(maxSize) {
    this.maxSize = maxSize
    this.buffer = []
  }
  
  write(val) {
    if (this.buffer.length == this.maxSize) {
      throw new BufferFullError()
    }

    this.buffer.push(val)
  }

  read() {
    if (this.buffer.length == 0) {
      throw new BufferEmptyError()
    }
    return this.buffer.shift()
  }

  forceWrite(val) {
    if (this.buffer.length == this.maxSize) {
      this.buffer = this.buffer.slice(1)
    }
    this.write(val)
  }

  clear() {
    this.buffer = []
  }
}

export default CircularBuffer;

export class BufferFullError extends Error {
  constructor() {
    super()
  }
}

export class BufferEmptyError extends Error {
  constructor() {
    super()
  }
}
