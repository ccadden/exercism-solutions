class BinarySearch
  attr_accessor :list
  
  def initialize(list)
    @list = list
  end

  def search_for(val)
    left = 0
    right = list.length
    
    while left <= right
      i = ((left + right)/2).floor
      
      case
        when list[i].nil?
          return nil
        when list[i] > val
          right = i - 1
        when list[i] < val
          left = i + 1
        else
          return i
      end
    end
  end
end