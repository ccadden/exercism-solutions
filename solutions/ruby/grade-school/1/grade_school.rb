class School

  def initialize
    @class_list = Hash.new { |hash, key| hash[key] = [] }
    @student_set = Set.new
  end

  def add(student, grade_num)
    if @student_set.add?(student)
      @class_list[grade_num] << student
      @class_list[grade_num].sort!
      true
    else
      false
    end
  end

  def roster
    res = []
    @class_list.keys.sort.each { |key| res.concat(@class_list[key].sort) }
    res
  end

  def grade(number)
    @class_list[number]
  end
end