export class List {
  constructor(list) {
    this.values = list || []
  }

  myPush(list, el) {
    return [...list, el]
  }

  foldLeft(f, start, [head, ...tail]) {
    return head ? this.foldLeft(f, f(start, head), tail) : start
  }
  
  foldRight(f, start, [head, ...tail]) {
    return head ? f(this.foldRight(f, start, tail), head) : start
  }

  append(otherList) {
    return new List(this.foldLeft(
      (s, h) => this.myPush(s, h),
      this.values,
      otherList.values
    ))
  }

  concat(list) {
    return new List(this.foldLeft(
      (s, h) => [...s, ...h.values],
      this.values,
      list.values
    ))
  }

  filter(f) {
    return new List(this.foldLeft(
      (s, h) => f(h) ? this.myPush(s, h) : s,
      [],
      this.values
    ))
  }

  map(f) {
    return new List(this.foldLeft(
      (s, h) => this.myPush(s, f(h)),
      [],
      this.values
    ))
  }

  length() {
    return this.foldLeft(
      (s, h) => 1 + s,
      0,
      this.values
    )
  }

  foldl(f, start) {
    return this.foldLeft(f, start, this.values)
  }

  foldr(f, start) {
    return this.foldRight(f, start, this.values)
  }

  reverse() {
    return new List(
      this.foldRight(
        (s, h) => this.myPush(s, h),
        [],
        this.values
      )
    )
  }
}
