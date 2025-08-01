const foldLeft = (f, start, list) => {
    const [head, ...tail] = list
  
    if (head === undefined) { return start }
    
    return foldLeft(f, f(start, head), tail)
}

const foldRight = (f, start, list) => {
  const [head, ...tail] = list

  if (head === undefined) { return start }
    
  return f(foldRight(f, start, tail), head)
}

export class List {
  constructor(list) {
    this.values = list || []
  }

  append(otherList) {
    return new List(foldLeft(
      (s, h) => { s.push(h); return s },
      this.values,
      otherList.values
    ))
  }

  concat(list) {
    return new List(foldLeft(
      (s, h) => [...s, ...h.values],
      this.values,
      list.values
    ))
  }

  filter(f) {
    return new List(foldLeft(
      (s, h) => {
        if (f(h)) {
          s.push(h)
        }
        return s
      },
      [],
      this.values
    ))
  }

  map(f) {
    return new List(foldLeft(
      (s, h) => { s.push(f(h)); return s },
      [],
      this.values
    ))
  }

  length() {
    return foldLeft(
      (s, h) => 1 + s,
      0,
      this.values
    )
  }

  foldl(f, start) {
    return foldLeft(f, start, this.values)
  }

  foldr(f, start) {
    return foldRight(f, start, this.values)
  }

  reverse() {
    return new List(
      foldRight(
        (s, h) => { s.push(h); return s },
        [],
        this.values
      )
    )
  }
}
