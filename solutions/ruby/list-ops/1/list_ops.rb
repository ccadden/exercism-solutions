class ListOps
  def self.arrays(arrs)
    count = 0
    (0...arrs.length).each do |_|
      count += 1
    end

    count
  end

  def self.reverser(arr)
    return [] if arr.empty?
    
    res = []
    (arr.length - 1..0).step(-1).each do |i|
      res << arr[i]
    end

    res
  end

  def self.concatter(arr1, arr2)
    res = []

    arr1.each { |e| res << e }
    arr2.each { |e| res << e }
    
    res
  end

  def self.mapper(arr)
    res = []
    
    arr.each do |e|
      res << yield(e) if block_given?
    end

    res
  end

  def self.filterer(arr)
    res = []

    arr.each do |e|
      res << e if block_given? && yield(e)
    end

    res
  end

  def self.sum_reducer(arr)
    sum = 0

    arr.each { |e| sum += e}

    sum
  end

  def self.factorial_reducer(arr)
    product = 1

    arr.each { |e| product *= e }

    product
  end
end